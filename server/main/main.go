package main

import (
	context "context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"server/config"
	"server/implementation"
	proto "server/protos"
)

// A rectangle that doesn't contain anything (for use as the default oldRect for new ViewStreams).
var NegativeRect = &proto.Rectangle{
	StartX: -100000,
	StartY: -100000,
	Width:  1,
	Height: 1,
}

type factoryProvider struct {
	configs                                     config.ConfigProvider
	tileSize, worldWidthTiles, worldHeightTiles int
	factory                                     implementation.Factory
}

func (f factoryProvider) RequestViewStream(server proto.FactoryService_RequestViewStreamServer) error {
	<-f.factory.AddViewStream(server)
	return nil
}

func (f factoryProvider) Interact(ctx context.Context, interaction *proto.Interaction) (*proto.Empty, error) {
	f.factory.SendInteraction(interaction)
	return &proto.Empty{}, nil
}

func (f factoryProvider) RequestView(ctx context.Context, rectangle *proto.Rectangle) (*proto.ScreenResponse, error) {
	return f.factory.RequestViewSquares(rectangle, NegativeRect)
}

// todo make this part of configurations rather than hardcoding it. A little bit annoying to do but probably a good idea
func makeInteractions() map[string]implementation.Interaction {
	res := make(map[string]implementation.Interaction)
	res["0"] = implementation.NewPaintInteraction(0)
	res["1"] = implementation.NewPaintInteraction(1)
	res["2"] = implementation.NewPaintInteraction(2)
	res["b"] = implementation.NewSpawnInteraction(func(x,y int32) implementation.EntityTicker{
		return implementation.BEETLE
	})
	res["a"] = implementation.NewSpawnInteraction(func(x,y int32) implementation.EntityTicker{
		return implementation.NewTimeAnimator(100,[]int32{0,1,2,3,4,3,2,1})
	})
	return res
}

func makeFactoryProvider() *factoryProvider {
	res := factoryProvider{}
	res.configs = config.MakeConfigs()
	res.tileSize = res.configs.GetConfigI("world-configs", "tile-size")
	res.worldWidthTiles = res.configs.GetConfigI("world-configs", "world-width-tiles")
	res.worldHeightTiles = res.configs.GetConfigI("world-configs", "world-height-tiles")
	waitsForSync := res.configs.GetConfig("world-configs", "tile-sync").(bool)
	msBetweenTicks := 1000 / res.configs.GetConfigI64("world-configs", "max-tile-tick-rate")
	res.factory = implementation.MakeFactory(res.tileSize, res.worldWidthTiles, res.worldHeightTiles, waitsForSync, msBetweenTicks, res.configs, makeInteractions())
	res.factory.StartRunning()
	return &res
}

func main() {
	provider := makeFactoryProvider()
	serverOptions := grpc.WriteBufferSize(32 * 1024 * 10)
	grpcServer := grpc.NewServer(serverOptions)

	proto.RegisterFactoryServiceServer(grpcServer, provider)
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d",
		provider.configs.GetConfig("net-configs", "hostname"),
		int(provider.configs.GetConfig("net-configs", "service-port").(float64))))
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}
