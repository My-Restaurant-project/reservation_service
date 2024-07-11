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

func NewRestaurantRepository(db *sqlx.DB) *RestaurantRepository {
	return &RestaurantRepository{db: db}
}

func (r *RestaurantRepository) CreateRestaurant(ctx context.Context, resReq *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	var addRes *reser.AddRestaurantResponse
	newId := uuid.NewString()
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
	ID := req.Id
	query := `
		SELECT name, description, address, phone_number, created_at, updated_at
        FROM restaurants
        WHERE id = $1 AND deleted_at IS NULL
    `
	row := r.db.QueryRowContext(ctx, query, ID)
	var res reser.Restaurant
	err := row.Scan(&res.Name, &res.Description, &res.Address, &res.PhoneNumber, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return &reser.GetRestaurantResponse{Restaurant: &reser.Restaurant{}}, err
	}

	return &reser.GetRestaurantResponse{Restaurant: &res}, nil
}

func (r *RestaurantRepository) UpdateRestaurant(ctx context.Context, req *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error) {

	query := `
		UPDATE restaurants
        SET name = $1, description = $2, address = $3, phone_number = $4, updated_at = now()
        WHERE id = $5
		RETURNING id, name, description, address, phone_number, created_at, updated_at
    `
	row := r.db.QueryRowContext(ctx, query, req.Name, req.Description, req.Address, req.PhoneNumber, req.Id)
	var updRestRes reser.UpdateRestaurantResponse
	var res reser.Restaurant

	err := row.Scan(&res.Id, &res.Name, &res.Description, &req.Address, &res.PhoneNumber, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		log.Println(err)
		return &updRestRes, err
	}
	updRestRes.Restaurant = &res

	return &updRestRes, nil
}

func (r *RestaurantRepository) DeleteRestaurant(ctx context.Context, req *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error) {
	Id := req.Id
	query := `
		UPDATE restaurants
        SET deleted_at = now()
        WHERE id = $1
		
    `
	_, err := r.db.ExecContext(ctx, query, Id)
	if err != nil {
		log.Println(err)
		return &reser.DeleteRestaurantResponse{Deleted: false}, err
	}

	return &reser.DeleteRestaurantResponse{Deleted: true}, nil
}

func (r *RestaurantRepository) GetAllRestaurants(ctx context.Context, req *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error) {

	query := `
		SELECT id, name, description, address, phone_number, created_at, updated_at
        FROM restaurants
        WHERE deleted_at IS NULL
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return &reser.GetRestaurantsResponse{Restaurant: []*reser.Restaurant{}}, err
	}
	defer rows.Close()
	var restaurants []*reser.Restaurant
	for rows.Next() {
		var res reser.Restaurant
		err := rows.Scan(&res.Id, &res.Name, &res.Description, &res.Address, &res.PhoneNumber, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			log.Println(err)
			return &reser.GetRestaurantsResponse{Restaurant: []*reser.Restaurant{}}, err
		}
		restaurants = append(restaurants, &res)
	}

	return &reser.GetRestaurantsResponse{Restaurant: restaurants}, nil
}
