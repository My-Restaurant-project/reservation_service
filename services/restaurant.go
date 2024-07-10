package services

import (
	reser "reservation_service/genproto/reservation_service"
)

type RestaurantService interface {
	AddRestaurant(*reser.AddRestaurantRequest) *reser.AddRestaurantResponse
	GetRestaurantById(*reser.GetRestaurantRequest) *reser.GetRestaurantResponse
	UpdateRestaurant(*reser.UpdateRestaurantRequest) *reser.UpdateRestaurantResponse
	DeleteRestaurant(*reser.DeleteRestaurantRequest) *reser.DeleteRestaurantResponse
	GetAllRestaurants(*reser.GetRestaurantsRequest) *reser.GetRestaurantsResponse
}
