package main

import (
	c "Github.com/Project-2/Reservation-Service/configs"
	pb "Github.com/Project-2/Reservation-Service/genproto/reservation_service"
	"Github.com/Project-2/Reservation-Service/pkg"
	"Github.com/Project-2/Reservation-Service/repository"
	"Github.com/Project-2/Reservation-Service/services"
	"google.golang.org/grpc"
	"log"
	"net"
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
