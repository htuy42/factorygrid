import grpc
import pygame

import protos.factori_pb2_grpc as pb2g

from config import Config
from view import View

conf = Config()

channel = grpc.insecure_channel(
    conf.getConfig("net-configs", "hostname") + ":" + str(conf.getConfig("net-configs", "service-port")))

stub = pb2g.FactoryServiceStub(channel)

pygame.display.init()
disp = pygame.display.set_mode(
    (conf.getConfig("frontend-configs", "default-width"), conf.getConfig("frontend-configs", "default-height")))
view = View(conf.getConfig("world-configs", "tile-size"), conf.getConfig("frontend-configs", "render-extra-border"),
            stub, disp, conf)

view.start()
