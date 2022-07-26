package main

import (
	"fmt"
	"klickhr-hris/pkg/config"
	"klickhr-hris/pkg/db"
	"klickhr-hris/pkg/pb"
	"klickhr-hris/pkg/services"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DbUser, c.DbPass, c.DbHost, c.DbPort, c.DbName)
	h := db.Init(dsn)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("HRIS Svc on", c.Port)

	s := services.Server{
		H: h,
	}

	grpcServer := grpc.NewServer()

	pb.RegisterHRISServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
