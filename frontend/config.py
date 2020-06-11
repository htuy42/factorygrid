import json

ROOTPATH = "../configs/"

class Config:
    def __init__(self):
        self.readFiles = {}

    def getConfig(self ,filename, configName):
        if filename not in self.readFiles:
            self.readFile(filename)
        return self.readFiles[filename][configName]

    # private
    def readFile(self,filename):
        with open(ROOTPATH + filename + ".json") as f:
            self.readFiles[filename] = json.load(f)

