package repository

import (
	"context"
	"log"
	reser "reservation_service/genproto/reservation_service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type RestaurantRepository struct {
	db *sqlx.DB
}

func NewRestaurantRepo(db *sqlx.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) CreateRestaurant(ctx context.Context, resReq *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	var addRes *reser.AddRestaurantResponse
	newId := uuid.NewString()
	log.Println(resReq)
	query := `
		INSERT INTO restaurants (id, name, description, address, phone_number)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.db.ExecContext(ctx, query, newId, resReq.Name, resReq.Description, resReq.Address, resReq.PhoneNumber)
	if err != nil {
		return addRes, err
	}

	addRes = &reser.AddRestaurantResponse{
		Id:          newId,
		Name:        resReq.Name,
		Description: resReq.Description,
		Address:     resReq.Address,
		PhoneNumber: resReq.PhoneNumber,
	}

	return addRes, nil
}

func (r *RestaurantRepository) GetRestaurantById(ctx context.Context, req *reser.GetRestaurantRequest) (*reser.GetRestaurantResponse, error) {
	return nil, nil
}

func (r *RestaurantRepository) UpdateRestaurant(ctx context.Context, req *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error) {
	return nil, nil
}

func (r *RestaurantRepository) DeleteRestaurant(ctx context.Context, req *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error) {
	return nil, nil
}

func (r *RestaurantRepository) GetAllRestaurants(ctx context.Context, req *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error) {
	return nil, nil
}
