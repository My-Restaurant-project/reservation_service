package main

import (
	"log"
	"net"
	c "reservation_service/configs"
	pb "reservation_service/genproto/reservation_service"
	"reservation_service/pkg"
	"reservation_service/services"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection" // Import the reflection package
)

func main() {
	// Load configuration
	config := c.Load()

	// Connect to the database
	db, err := pkg.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create a TCP listener on specified port
	listener, err := net.Listen("tcp", ":"+config.URL_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	log.Println("Server started on port " + config.URL_PORT)

	// Initialize repository and services
	rs := services.NewReservationService(db)

	// Create a new gRPC server
	s := grpc.NewServer()

	// Register the ReservationServiceServer with the gRPC server
	pb.RegisterReservationServiceServer(s, rs)

	// Register reflection service on gRPC server
	reflection.Register(s)

	// Start the gRPC server
	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
