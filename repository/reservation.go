package repository

import (
	"context"
	"database/sql"
	"fmt"
	reser "reservation_service/genproto/reservation_service"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReservationRespoitory struct {
	db *sqlx.DB
}

func NewReservationRepo(db *sqlx.DB) *ReservationRespoitory {
	return &ReservationRespoitory{db: db}
}

func (r *ReservationRespoitory) CreateReservation(ctx context.Context, reservRes *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	var addRes *reser.AddReservationResponse
	newId := uuid.NewString()

	qry := `
		SELECT id
        FROM restaurants
        WHERE id = $1 AND deleted_at IS NULL

	`
	row := r.db.QueryRow(qry, reservRes.RestaurantId)
	var resId string
	err := row.Scan(&resId)
	if err != nil {
		if err == sql.ErrNoRows {
			return addRes, fmt.Errorf("restaurant not found")
		}
		return addRes, err
	}

	query := `
		INSERT INTO restaurants (id, user_id, restaurant_id, reservation_time, status)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err = r.db.ExecContext(ctx, query, newId, reservRes.UserId, reservRes.RestaurantId, reservRes.ReservationTime, reservRes.Status)
	if err != nil {
		return addRes, err
	}

	addRes = &reser.AddReservationResponse{
		Id:              newId,
		UserId:          reservRes.UserId,
		RestaurantId:    reservRes.RestaurantId,
		ReservationTime: reservRes.ReservationTime,
		Status:          reservRes.Status,
	}
	
	return addRes, nil
}

func (r *ReservationRespoitory) GetReservationById(ctx context.Context, req *reser.GetReservationRequest) (*reser.GetReservationResponse, error) {
	ID := req.Id
	query := `
		SELECT user_id, restaurant_id, reservation_time, status, created_at, updated_at
        FROM reservations
        WHERE id = $1 AND deleted_at IS NULL
    `
	row := r.db.QueryRowContext(ctx, query, ID)

	var res reser.Reservation

	err := row.Scan(&res.UserId, &res.RestaurantId, &res.ReservationTime, &res.Status, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &reser.GetReservationResponse{Reservation: &reser.Reservation{}}, err
	}

	return &reser.GetReservationResponse{Reservation: &res}, nil
}



func (r *ReservationRespoitory) UpdateReservation(ctx context.Context, req *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error) {
	query := `
		UPDATE reservations
        SET restaurant_id = $1, reservation_time = $2, status = $3, updated_at = now()
        WHERE id = $4
		RETURNING id, user_id, restaurant_id, reservation_time, status, created_at, updated_at
    `
	row := r.db.QueryRowContext(ctx, query, req.RestaurantId, req.ReservationTime, req.Status, req.Id)
	var updResRes *reser.UpdateReservationResponse
	var res reser.Reservation
	
	err := row.Scan(&res.Id, &res.UserId, &res.RestaurantId, &req.ReservationTime, &res.Status, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return updResRes, err
	}
	updResRes.Reservation = &res

	return updResRes, nil
}



func (r *ReservationRespoitory) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	ID := req.Id
	query := `
		UPDATE reservations
        SET deleted_at = now()
        WHERE id = $1
		
    `
	_, err := r.db.ExecContext(ctx, query, ID)
	if err != nil {
		return &reser.DeleteReservationResponse{Deleted: false}, err
	}

	return &reser.DeleteReservationResponse{Deleted: true}, nil
}


func (r *ReservationRespoitory) GetAllReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {
	query := `
		SELECT id, user_id, restaurant_id, reservation_time, status, created_at, updated_at
        FROM reservations
        WHERE deleted_at IS NULL
    `
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return  &reser.GetReservationsResponse{Reservations: []*reser.Reservation{}}, err
	}
	defer rows.Close()
	var reservations []*reser.Reservation
	for rows.Next() {
		var res reser.Reservation
		err := rows.Scan(&res.Id, &res.UserId, &res.RestaurantId, &res.ReservationTime, &res.Status, &res.CreatedAt, &res.UpdatedAt)
		if err != nil {
			return &reser.GetReservationsResponse{Reservations: []*reser.Reservation{}}, err
		}
		reservations = append(reservations, &res)
	}

	return &reser.GetReservationsResponse{Reservations: reservations}, nil
}


