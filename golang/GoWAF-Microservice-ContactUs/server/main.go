package main

import (
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/configuration"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/handler"
	service "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/internal"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/model"
	ormGopg "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/persistance/orm-go-pg"
	"github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/pkg/gopg"

	contact "github.com/NlaakStudiosLLC/GoWAF-Microservice-ContactUs/server/proto"

	"github.com/micro/go-grpc"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
)

const (
	pathToCfg = "./configuration/config.json"
)

func main() {
	cfg, err := configuration.Extract(pathToCfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gopg.New(cfg.Database.DbType, cfg.Database.Name, cfg.Database.Name, &model.Contact{}, nil)
	if err != nil {
		log.Fatal(err)
	}

	dbI := ormGopg.New(db)

	srv := service.New(dbI)

	contactS := grpc.NewService(
		micro.Name(cfg.General.ServiceName),
		micro.Version(cfg.General.Version),
	)

	// Register Handler
	c := handler.New(srv)
	err = contact.RegisterContactMicroHandler(contactS.Server(), c)
	if err != nil {
		log.Fatal(err)
	}

	contactS.Init()

	// Run service
	if err := contactS.Run(); err != nil {
		log.Fatal(err)
	}
}
