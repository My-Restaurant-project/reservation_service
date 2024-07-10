package main

import (
	"log"
	"net"
	c "reservation_service/configs"
	pb "reservation_service/genproto/reservation_service"
	"reservation_service/pkg"
	"reservation_service/repository"
	"reservation_service/services"

	"google.golang.org/grpc"
)

func main() {
	config := c.Load()

	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on port  " + config.URL_PORT)

	reserve := repository.NewReservationRepo(db)

	rs := services.NewReservationService(reserve)

	s := grpc.NewServer()
	pb.RegisterReservationServiceServer(s, rs)

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
