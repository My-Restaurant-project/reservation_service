package repository

import (
	"context"
	"fmt"
	reser "reservation_service/genproto/reservation_service"
	"strings"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ReservationRespoitory struct {
	db *sqlx.DB
}

func NewReservationRepository(db *sqlx.DB) *ReservationRespoitory {
	return &ReservationRespoitory{db: db}
}

func (r *ReservationRespoitory) CreateReservation(ctx context.Context, reservRes *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	var addRes *reser.AddReservationResponse
	newId := uuid.NewString()

	query := `
		INSERT INTO reservations (id, user_id, restaurant_id, reservation_time, status)
        VALUES ($1, $2, $3, $4, $5)
    `

	_, err := r.db.ExecContext(ctx, query, newId, reservRes.UserId, reservRes.RestaurantId, reservRes.ReservationTime, reservRes.Status)
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
        SET  reservation_time = $1, status = $2, updated_at = now()
        WHERE id = $3 AND deleted_at IS NULL
		RETURNING id, user_id, restaurant_id, reservation_time, status, created_at, updated_at
    `
	row := r.db.QueryRowContext(ctx, query, req.ReservationTime, req.Status, req.Id)
	var updResRes reser.UpdateReservationResponse
	var res reser.Reservation

	err := row.Scan(&res.Id, &res.UserId, &res.RestaurantId, &req.ReservationTime, &res.Status, &res.CreatedAt, &res.UpdatedAt)
	if err != nil {
		return &updResRes, err
	}
	updResRes.Reservation = &res

	return &updResRes, nil
}

func (r *ReservationRespoitory) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	ID := req.Id
	query := `
		UPDATE reservations
        SET deleted_at = now()
        WHERE id = $1 AND deleted_at IS NULL
		
    `
	_, err := r.db.ExecContext(ctx, query, ID)
	if err != nil {
		return &reser.DeleteReservationResponse{Deleted: false}, err
	}

	return &reser.DeleteReservationResponse{Deleted: true}, nil
}

func (r *ReservationRespoitory) GetAllReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {

	params := []string{}
	args := []interface{}{}

	query := `SELECT id, user_id, restaurant_id, reservation_time, status, created_at, updated_at FROM reservations `

	if req.GetRestaurantId() != "" {
		params = append(params, fmt.Sprintf("restaurant_id = $%d", len(args)+1))
		args = append(args, req.GetRestaurantId())
	}

	if req.GetUserId() != "" {
		params = append(params, fmt.Sprintf("user_id = $%d", len(args)+1))
		args = append(args, req.GetUserId())
	}

	if req.GetStatus() != "" {
		params = append(params, fmt.Sprintf("status = $%d", len(args)+1))
		args = append(args, req.GetStatus())
	}

	if len(params) > 0 {
		query += " WHERE " + strings.Join(params, " AND ") + " AND deleted_at IS NULL"
	}

	query += " ORDER BY created_at DESC"

	reservations := []*reser.Reservation{}
	fmt.Println(query, args)
	err := r.db.SelectContext(ctx, &reservations, query, args...)
	if err != nil {
		return nil, err
	}

	return &reser.GetReservationsResponse{Reservations: reservations}, nil
}
