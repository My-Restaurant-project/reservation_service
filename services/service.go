package services

import (
	"context"
	user "reservation_service/genproto/authentication_service"
	reser "reservation_service/genproto/reservation_service"
	repo "reservation_service/repository"

	"github.com/jmoiron/sqlx"
)

type ReservationService interface {
	RestaurantService() RestaurantService
}

type reservationServiceImpl struct {
	reser.UnimplementedReservationServiceServer
	user.UnimplementedAuthenticationServiceServer
	restaurantService RestaurantService
}

func NewReservationService(db *sqlx.DB) *reservationServiceImpl {
	return &reservationServiceImpl{restaurantService: NewRestaurantService(repo.NewRestaurantRepo(db))}
}

func (rs *reservationServiceImpl) RestaurantService() RestaurantService {
	return rs.restaurantService
}

func (rs *reservationServiceImpl) AddRestaurant(ctx context.Context, resReq *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	resp, err := rs.RestaurantService().AddRestaurant(ctx, resReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
