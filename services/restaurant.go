package services

import (
	"context"
	reser "reservation_service/genproto/reservation_service"
	"reservation_service/repository"
)

type RestaurantService interface {
	AddRestaurant(context.Context, *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error)
	GetRestaurant(context.Context, *reser.GetRestaurantRequest) (*reser.GetRestaurantResponse, error)
	UpdateRestaurant(context.Context, *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error)
	DeleteRestaurant(context.Context, *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error)
	GetRestaurants(context.Context, *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error)
}

type restaurantServiceImpl struct {
	repo *repository.RestaurantRepository
}

func NewRestaurantService(repo *repository.RestaurantRepository) RestaurantService {
	return &restaurantServiceImpl{repo: repo}
}

func (rs *restaurantServiceImpl) AddRestaurant(ctx context.Context, req *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	return rs.repo.CreateRestaurant(ctx, req)
}

func (rs *restaurantServiceImpl) GetRestaurant(ctx context.Context, req *reser.GetRestaurantRequest) (*reser.GetRestaurantResponse, error) {
	return rs.repo.GetRestaurantById(ctx, req)
}

func (rs *restaurantServiceImpl) UpdateRestaurant(ctx context.Context, req *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error) {
	return rs.repo.UpdateRestaurant(ctx, req)
}

func (rs *restaurantServiceImpl) DeleteRestaurant(ctx context.Context, req *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error) {
	return rs.repo.DeleteRestaurant(ctx, req)
}

func (rs *restaurantServiceImpl) GetRestaurants(ctx context.Context, req *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error) {
	return rs.repo.GetAllRestaurants(ctx, req)
}
