package repository

import (
	"context"
	reser "reservation_service/genproto/reservation_service"

	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type ReservationOrderRespository struct {
	db *sqlx.DB
}

func NewReservationOrderRepository(db *sqlx.DB) *ReservationOrderRespository {
	return &ReservationOrderRespository{db: db}
}

func (r *ReservationOrderRespository) CreateReservationOrder(ctx context.Context, req *reser.AddReservationOrderRequest) (*reser.AddReservationOrderResponse, error) {
	var resp *reser.AddReservationOrderResponse

	id := uuid.NewString()

	query := `insert into reservationOrders (id,reservation_id, menu_item_id, quantity) values($1,$2,$3,$4)`

	_, err := r.db.ExecContext(ctx, query, id, req.ReservationId, req.MenuItemId, req.Quantity)
	if err != nil {
		return resp, err
	}

	resp = &reser.AddReservationOrderResponse{
		Id:            id,
		ReservationId: req.ReservationId,
		MenuItemId:    req.MenuItemId,
		Quantity:      req.Quantity,
		CreatedAt:     cast.ToString(time.Now()),
	}

	return resp, nil
}

func (r *ReservationOrderRespository) GetReservationOrderById(ctx context.Context, req *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error) {
	query := `select restaurant_id, menu_item_id, quantity from reservationOrders where id = $1`

	row := r.db.QueryRowContext(ctx, query, req.Id)

	var resp *reser.GetReservationOrderResponse
	var resOrd = &reser.ReservationOrder{}

	err := row.Scan(&resOrd.ReservationId, &resOrd.MenuItemId, &resOrd.Quantity)
	if err != nil {
		return resp, err
	}

	resp.ReservationOrder = resOrd

	return resp, nil
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
