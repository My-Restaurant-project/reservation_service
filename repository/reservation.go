package repository

import (
	"context"
	"fmt"
	reser "reservation_service/genproto/reservation_service"
	"strings"

	"github.com/jmoiron/sqlx"
)

type ReservationRespoitory struct {
	db *sqlx.DB
}

func NewReservationRepo(db *sqlx.DB) *ReservationRespoitory {
	return &ReservationRespoitory{db: db}
}

func (r *ReservationRespoitory) CreateReservation(ctx context.Context, reservRes *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	return nil, nil
}

func (r *ReservationRespoitory) GetReservationById(ctx context.Context, req *reser.GetReservationRequest) (*reser.GetReservationResponse, error) {
	return nil, nil
}

func (r *ReservationRespoitory) UpdateReservation(ctx context.Context, req *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error) {
	return nil, nil
}

func (r *ReservationRespoitory) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	return nil, nil
}

func (r *ReservationRespoitory) GetAllReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {

	params := []string{}
	args := []interface{}{}

	query := `SELECT id, user_id, restaurant_id, reservation_time, status, created_at, updated_at FROM reservations`

	if req.GetRestaurantId() != "" {
		params = append(params, fmt.Sprintf("restaurant_id =$%d", len(args)+1))
		args = append(args, req.GetRestaurantId())
	}

	if req.GetUserId() != "" {
		params = append(params, fmt.Sprintf("user_id =$%d", len(args)+1))
		args = append(args, req.GetUserId())
	}

	if req.GetStatus() != "" {
		params = append(params, fmt.Sprintf("status =$%d", len(args)+1))
		args = append(args, req.GetStatus())
	}

	if len(params) > 0 {
		query += " WHERE " + strings.Join(params, " AND ") + "WHERE deleted_at IS NULL"
	}
	reservations := []*reser.Reservation{}

	err := r.db.SelectContext(ctx, &reservations, query, args...)
	if err != nil {
		return nil, err
	}

	return &reser.GetReservationsResponse{Reservations: reservations}, nil
}
