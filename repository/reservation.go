package repository

import (
	"context"
	reser "reservation_service/genproto/reservation_service"

	"github.com/jmoiron/sqlx"
)

type ReservationRespoitory struct {
	db *sqlx.DB
}

func NewReservationRepo(db *sqlx.DB) *ReservationRespoitory {
	return &ReservationRespoitory{db: db}
}

func (r *ReservationRespoitory) CreateReservation(ctx context.Context, reservRes *reser.AddReservationRequest) (*reser.AddReservationOrderResponse, error) {
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
	return nil, nil
}
