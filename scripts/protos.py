# generates the protobuf code for the project.
import os
os.chdir("../protos")
os.system("protoc -I . *.proto --go_out=plugins=grpc:../server/protos")
os.chdir("../")
os.system("python -m grpc_tools.protoc -I . protos/*.proto --python_out=frontend --grpc_python_out=frontend")