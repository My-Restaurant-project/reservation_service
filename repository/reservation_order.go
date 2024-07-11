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

	query := `insert into reservationOrders (id, reservation_id, menu_item_id, quantity) values($1, $2, $3, $4)`

	_, err := r.db.ExecContext(ctx, query, id, req.ReservationId, req.MenuItemId, req.Quantity)
	if err != nil {
		return resp, err
	}

	resp = &reser.AddReservationOrderResponse{
		Id:            id,
		ReservationId: req.ReservationId,
		MenuItemId:    req.MenuItemId,
		Quantity:      req.Quantity,
		CreatedAt:     time.Now().String(),
	}

	return resp, nil
}

func (r *ReservationOrderRespository) GetReservationOrderById(ctx context.Context, req *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error) {
	query := `select reservation_id, menu_item_id, quantity from reservationOrders where id = $1 deleted_at is null`

	row := r.db.QueryRowContext(ctx, query, req.Id)

	var resp = &reser.GetReservationOrderResponse{}
	var resOrd = &reser.ReservationOrder{}

	err := row.Scan(&resOrd.ReservationId, &resOrd.MenuItemId, &resOrd.Quantity)
	if err != nil {
		return resp, err
	}

	resp.ReservationOrder = resOrd

	return resp, nil
}

func (r *ReservationOrderRespository) UpdateReservationOrder(ctx context.Context, req *reser.UpdateReservationOrderRequest) (*reser.UpdateReservationOrderResponse, error) {
	var resp = &reser.UpdateReservationOrderResponse{}

	query := `update reservationOrders set reservation_id=$1, menu_item_id=$2, quantuty=$3 from reservationOrders where id=$4 and deleted_at is null`

	_, err := r.db.ExecContext(ctx, query, req.ReservationId, req.MenuItemId, req.Quantity, req.Id)
	if err != nil {
		return resp, err
	}

	var updOrd = &reser.ReservationOrder{
		Id:            req.Id,
		ReservationId: req.ReservationId,
		MenuItemId:    req.MenuItemId,
		Quantity:      req.Quantity,
		CreatedAt:     cast.ToString(time.Now()),
	}

	resp.ReservationOrder = updOrd

	return resp, nil
}

func (r *ReservationOrderRespository) DeleteReservationOrder(ctx context.Context, req *reser.DeleteReservationOrderRequest) (*reser.DeleteReservationOrderResponse, error) {
	query := `update reservationorders set deleted_at=$1 where id=$2 and deleted_at is null`

	deleted := time.Now()

	_, err := r.db.ExecContext(ctx, query, deleted, req.Id)
	if err != nil {
		return &reser.DeleteReservationOrderResponse{Deleted: false}, err
	}

	return &reser.DeleteReservationOrderResponse{Deleted: true}, nil
}

func (r *ReservationOrderRespository) GetAllReservationOrders(ctx context.Context, req *reser.GetReservationOrdersRequest) (*reser.GetReservationOrdersResponse, error) {
	query := `select id, reservation_id, menu_item_id, quantity from reservationorders where deleted_at is null`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var resp = &reser.GetReservationOrdersResponse{}
	var resReqs = []*reser.ReservationOrder{}
	for rows.Next() {
		var resReq = &reser.ReservationOrder{}

		err := rows.Scan(&resReq.Id, &resReq.ReservationId, &resReq.MenuItemId, &resReq.Quantity)
		if err != nil {
			return nil, err
		}
		resReqs = append(resReqs, resReq)
	}
	resp.ReservationOrders = resReqs

	return resp, nil
}
