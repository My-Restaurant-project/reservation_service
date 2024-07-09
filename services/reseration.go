package services

import (
	pb "Github.com/Project-2/Reservation-Service/genproto/reservation_service"
	"Github.com/Project-2/Reservation-Service/repository"
	"context"
)

type ReservationService struct {
	pb.UnimplementedReservationServiceServer
	repo *repository.ReservationRepo
}

func NewReservationService(reservationRepo *repository.ReservationRepo) *ReservationService {
	return &ReservationService{repo: reservationRepo}
}

func (r *ReservationService) CreateRestaurant(ctx context.Context, req *pb.AddRestaurantRequest) (*pb.AddRestaurantResponse, error) {
	res, err := r.repo.CreateRestaurant(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) GetRestaurantById(ctx context.Context, req *pb.GetRestaurantRequest) (*pb.GetRestaurantResponse, error) {
	res, err := r.repo.GetRestaurantById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) UpdateRestaurant(ctx context.Context, req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error) {
	res, err := r.repo.UpdateRestaurant(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) DeleteRestaurant(ctx context.Context, req *pb.DeleteRestaurantRequest) (*pb.DeleteRestaurantResponse, error) {
	res, err := r.repo.DeleteRestaurant(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (r *ReservationService) GetAllRestaurants(ctx context.Context, req *pb.GetRestaurantsRequest) (*pb.GetRestaurantsResponse, error) {
	res, err := r.repo.GetAllRestaurants(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
