package main

import (
	api "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/http"

	contactGrpc "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/grpc"
	protoC "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/client/grpc/proto"

	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/micro/go-web"

	"github.com/micro/go-grpc"
)

func main() {
	//create grpc service
	serviceGrpc := grpc.NewService()

	//init gprc service
	serviceGrpc.Init(
		micro.Name("go.micro.srv.server"),
		micro.Version("latest"),
	)

	contactClient := protoC.NewContactMicroService("go.micro.srv.contact", serviceGrpc.Client())
	userMicro := contactGrpc.New(contactClient)
	contactController := api.NewHandler(userMicro, nil)

	routers := api.InitRouters(contactController)
	routers.UseEncodedPath()

	// create new web service
	service := web.NewService(
		web.Name("my.microservice.webclient"),
		web.Version("latest"),
		web.Address(":8000"),
		web.Handler(routers),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
