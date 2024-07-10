package repository

import (
	"context"
	reser "reservation_service/genproto/reservation_service"

	"github.com/jmoiron/sqlx"
)

type ReservationOrderRespository struct {
	db *sqlx.DB
}

func NewReservationOrderRepository(db *sqlx.DB) *ReservationOrderRespository {
	return &ReservationOrderRespository{db: db}
}

func (r *ReservationOrderRespository) CreateReservationOrder(ctx context.Context, reservOrderRes *reser.AddReservationOrderRequest) (*reser.AddReservationOrderResponse, error) {
	return nil, nil
}

func (r *ReservationOrderRespository) GetReservationOrderById(ctx context.Context, req *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error) {
	return nil, nil
}

func (r *ReservationOrderRespository) UpdateReservationOrder(ctx context.Context, req *reser.UpdateReservationOrderRequest) (*reser.UpdateReservationOrderResponse, error) {
	return nil, nil
}

func (r *ReservationOrderRespository) DeleteReservationOrder(ctx context.Context, req *reser.DeleteReservationOrderRequest) (*reser.DeleteReservationOrderResponse, error) {
	return nil, nil
}

func (r *ReservationOrderRespository) GetAllReservationOrders(ctx context.Context, req *reser.GetReservationOrdersRequest) (*reser.GetReservationOrdersResponse, error) {
	return nil, nil
}
