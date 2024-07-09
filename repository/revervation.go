package repository

import (
	"database/sql"
	"github.com/google/uuid"
	"github.com/spf13/cast"
	"time"

	pb "Github.com/Project-2/Reservation-Service/genproto/reservation_service"
)

type ReservationRepo struct {
	DB *sql.DB
}

func NewReservationRepo(db *sql.DB) *ReservationRepo {
	return &ReservationRepo{
		DB: db,
	}
}

func (r *ReservationRepo) CreateRestaurant(req *pb.AddRestaurantRequest) (*pb.AddRestaurantResponse, error) {
	query := `insert into restaurants (id, name, address, phone_number, description) values ($1, $2, $3, $4, $5)`

	id := uuid.NewString()
	_, err := r.DB.Exec(query, id, req.Name, req.Address, req.PhoneNumber, req.Description)
	if err != nil {
		return nil, err
	}
	return &pb.AddRestaurantResponse{
		Id:          id,
		Name:        req.Name,
		Address:     req.Address,
		PhoneNumber: req.PhoneNumber,
		Description: req.Description,
		CreatedAt:   cast.ToString(time.Now()),
	}, nil
}

func (r *ReservationRepo) GetRestaurantById(req *pb.GetRestaurantRequest) (*pb.GetRestaurantResponse, error) {
	query := `select name, address,phone_number, description from restaurants where id = $1`

	row := r.DB.QueryRow(query, req.Id)
	var res pb.GetRestaurantResponse
	err := row.Scan(&res.Restaurant.Name, &res.Restaurant.Address, &res.Restaurant.PhoneNumber, &res.Restaurant.Description)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *ReservationRepo) GetAllRestaurants(req *pb.GetRestaurantsRequest) (*pb.GetRestaurantsResponse, error) {
	query := `select id, name, address, phone_number, description from restaurantsw where deleted_at IS NULL`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var res pb.GetRestaurantsResponse
	var restaurants []*pb.Restaurant
	for rows.Next() {
		var row *pb.Restaurant

		err := rows.Scan(&row.Id, &row.Name, &row.Address, &row.PhoneNumber, &row.Description)

		if err != nil {
			return &pb.GetRestaurantsResponse{}, err
		}
		restaurants = append(restaurants, row)
	}
	res.Restaurant = restaurants
	return &res, nil
}

func (r *ReservationRepo) UpdateRestaurant(req *pb.UpdateRestaurantRequest) (*pb.UpdateRestaurantResponse, error) {
	query := `update restaurants set name, address,phone_number, description where id = $1`
	_, err := r.DB.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	var res pb.UpdateRestaurantResponse
	res.Restaurant.Name = req.Name
	res.Restaurant.Address = req.Address
	res.Restaurant.PhoneNumber = req.PhoneNumber
	res.Restaurant.Description = req.Description

	return &res, nil
}

func (r *ReservationRepo) DeleteRestaurant(req *pb.DeleteRestaurantRequest) (*pb.DeleteRestaurantResponse, error) {
	query := `update restaurants set deleted_at = $1 where id = $2`
	deletedAt := time.Now().Unix()
	_, err := r.DB.Exec(query, deletedAt, req.Id)
	if err != nil {
		return &pb.DeleteRestaurantResponse{Deleted: false}, err
	}
	return &pb.DeleteRestaurantResponse{Deleted: true}, nil
}
