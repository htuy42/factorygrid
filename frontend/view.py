import queue
import threading
from typing import Dict, Tuple

import pygame
from pygame.color import THECOLORS
from pygame.freetype import SysFont

import protos.factori_pb2 as pb2
import protos.factori_pb2_grpc as pb2g
from config import Config

# view client for the factori server

# repeatedly receive subviews and put them into the bufferedcontents of view
def handleIncomingSubviews(view):
    for viewset in view.responseIter:
        for subview in viewset.subViews:
            location = (subview.viewOf.startX, subview.viewOf.startY)
            view.bufferedContents[location] = subview

class View:
    def __init__(self, tileSize: int, renderExtraBorder: int, stub: pb2g.FactoryServiceStub, display: pygame.Surface,
                 config: Config):
        self.cellSize = config.getConfig("frontend-configs", "default-cell-size")
        self.x: int = 0
        self.y: int = 0
        self.xSize: int = display.get_width() / self.cellSize
        self.ySize: int = display.get_height() / self.cellSize
        self.tileSize: int = tileSize
        self.stub = stub
        self.bufferedContents: Dict[Tuple[int, int], pb2.ViewResponse] = {}
        self.renderExtraBoder = renderExtraBorder
        self.display = display
        self.running = True
        self.modeKey = 'q'
        self.colorCache = {}
        self.config = config
        pygame.freetype.init()
        self.myFont = SysFont('Comic Sans MS', 30)
        self.lastRect = None
        self.sendQueue = queue.Queue()
        self.lastInteractPos = (0,0)
        self.mouseDown = False


    def start(self):
        self.doRequestLoop()
        while self.running:
            self.step()

    def doRequestLoop(self):
        print("Requesting view stream")
        self.responseIter = self.stub.RequestViewStream(iter(self.sendQueue.get, None))
        threading.Thread(target=handleIncomingSubviews, args=(self,)).start()
        self.handleMove()

    def step(self):
        self.processInput()
        self.renderView()

    def getColor(self,colorNum):
        tileColor = self.colorCache.get(colorNum)
        if tileColor == None:
            tileColorList = self.config.getConfig("frontend-configs", "tile-colors")[colorNum]
            self.colorCache[colorNum] = (tileColorList[0], tileColorList[1], tileColorList[2])
            tileColor = self.colorCache[colorNum]
        return tileColor



    def drawTile(self, displayX, displayY, tile):
        pygame.draw.rect(self.display, self.getColor(tile.tileTypeId), (displayX, displayY, self.cellSize + 1, self.cellSize + 1))
        centerX = displayX + self.cellSize / 2
        centerY = displayY + self.cellSize / 2
        for entity in tile.entities:
            renderForm = self.config.getConfig("frontend-configs","entity-renders")[entity.typeId]
            size = renderForm["size"] * self.cellSize
            color = self.getColor(renderForm["color"])
            shape = renderForm["shape"]
            if shape == "circle":
                pygame.draw.circle(self.display,color,(int(centerX),int(centerY)),int(size / 2))
        # draw the entities in the cell

    def renderView(self):
        self.display.fill(THECOLORS["white"])
        for x in range(int(self.xSize)):
            for y in range(int(self.ySize)):
                realX = x + self.x
                realY = y + self.y
                tileX = realX / self.tileSize
                tileY = realY / self.tileSize
                if realX < 0 or realY < 0:
                    continue
                innerX = realX % self.tileSize
                innerY = realY % self.tileSize
                viewResp = self.bufferedContents.get((int(tileX), int(tileY)), None)
                if viewResp == None:
                    continue
                else:
                    tileIndex = innerY * self.tileSize + innerX
                    if tileIndex >= len(viewResp.tiles):
                        print(viewResp.tiles)
                        print(viewResp)
                        print("EROOR")
                        continue
                    tile = viewResp.tiles[tileIndex]
                    self.drawTile(x * self.cellSize, y * self.cellSize, tile)
        self.myFont.render_to(self.display, (0, 0), self.modeKey, (0, 0, 0))
        pygame.display.flip()

    def doInteraction(self, x, y, isMove=False):
        x /= self.cellSize
        y /= self.cellSize
        x = int(x)
        y = int(y)
        if isMove:
            if self.lastInteractPos == (x,y):
                return
        self.lastInteractPos = (x,y)
        self.stub.Interact(pb2.Interaction(x=int(x + self.x), y=int(y + self.y), interactionChar=self.modeKey))

    def processInput(self):
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                self.running = False
                return
            elif event.type == pygame.KEYDOWN:
                if event.key == pygame.K_i:
                    self.zoomIn()
                elif event.key == pygame.K_o:
                    self.zoomOut()
                elif event.key == pygame.K_DOWN:
                    self.moveScreen(0, 1)
                elif event.key == pygame.K_UP:
                    self.moveScreen(0, -1)
                elif event.key == pygame.K_LEFT:
                    self.moveScreen(-1, 0)
                elif event.key == pygame.K_RIGHT:
                    self.moveScreen(1, 0)
                else:
                    self.modeKey = pygame.key.name(event.key)
            elif event.type == pygame.MOUSEBUTTONDOWN and event.button == 1:
                x, y = event.pos
                self.doInteraction(x, y)
                self.mouseDown = True
            elif event.type == pygame.MOUSEBUTTONUP and event.button == 1:
                self.mouseDown = False
            elif event.type == pygame.MOUSEMOTION:
                if self.mouseDown:
                    x, y = event.pos
                    self.doInteraction(x,y,isMove=True)

    def handleMove(self):
        rect = pb2.Rectangle(startX=self.x - self.renderExtraBoder, startY=self.y - self.renderExtraBoder,
                                         width=int(self.xSize + 2 * self.renderExtraBoder), height=int(self.ySize + 2 * self.renderExtraBoder))
        # tempRect = pb2.Rectangle(startX=0,startY=0,width=10,height=25)
        if self.lastRect is None:
            self.lastRect = pb2.Rectangle(startX=-100000,startY=-100000,width=1,height=1)
        self.sendQueue.put(pb2.ViewRequest(fullView=rect,oldView=self.lastRect))
        self.lastRect = rect

    def zoomOut(self):
        self.cellSize -= max(self.cellSize * .1, 1)
        self.xSize: int = self.display.get_width() / self.cellSize
        self.ySize: int = self.display.get_height() / self.cellSize
        self.handleMove()

    def zoomIn(self):
        self.cellSize += max(1, self.cellSize * .1)
        self.xSize: int = self.display.get_width() / self.cellSize
        self.ySize: int = self.display.get_height() / self.cellSize
        self.handleMove()

    def moveScreen(self, xMul, yMul):
        self.x += 1 * xMul
        self.y += 1 * yMul
        self.handleMove()
