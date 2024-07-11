package services

import (
	"context"
	user "reservation_service/genproto/authentication_service"
	reser "reservation_service/genproto/reservation_service"
	repo "reservation_service/repository"

	"github.com/jmoiron/sqlx"
)

type MainService interface {
	RestaurantService() RestaurantService
	ReservationService() ReservationService
}

type mainServiceImpl struct {
	reser.UnimplementedReservationServiceServer
	user.UnimplementedAuthenticationServiceServer
	restaurantService   RestaurantService
	reservatiionService ReservationService
}

func NewMainService(db *sqlx.DB) *mainServiceImpl {
	return &mainServiceImpl{
		restaurantService:   NewRestaurantService(repo.NewRestaurantRepo(db)),
		reservatiionService: NewReservationService(repo.NewReservationRepo(db)),
	}
}

func (rs *mainServiceImpl) RestaurantService() RestaurantService {
	return rs.restaurantService
}

func (rs *mainServiceImpl) ReservationService() ReservationService {
	return rs.reservatiionService
}

func (rs *mainServiceImpl) AddRestaurant(ctx context.Context, resReq *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	resp, err := rs.RestaurantService().AddRestaurant(ctx, resReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetRestaurant(ctx context.Context, req *reser.GetRestaurantRequest) (*reser.GetRestaurantResponse, error) {
	resp, err := rs.RestaurantService().GetRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateRestaurant(ctx context.Context, req *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error) {
	resp, err := rs.RestaurantService().UpdateRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteRestaurant(ctx context.Context, req *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error) {
	resp, err := rs.RestaurantService().DeleteRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetRestaurants(ctx context.Context, req *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error) {
	resp, err := rs.RestaurantService().GetRestaurants(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// /////////////////////////////////////////
func (rs *mainServiceImpl) AddReservation(ctx context.Context, resReq *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	resp, err := rs.ReservationService().AddReservation(ctx, resReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetReservation(ctx context.Context, req *reser.GetReservationRequest) (*reser.GetReservationResponse, error) {
	resp, err := rs.ReservationService().GetReservation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateReservation(ctx context.Context, req *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error) {
	resp, err := rs.ReservationService().UpdateReservation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	resp, err := rs.ReservationService().DeleteReservation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getall
func (rs *mainServiceImpl) GetReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {
	resp, err := rs.ReservationService().GetReservations(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
