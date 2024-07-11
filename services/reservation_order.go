package services

import (
	"context"
	reser "reservation_service/genproto/reservation_service"
	"reservation_service/repository"
)

type ReservationOrderService interface {
	AddReservationOrder(context.Context, *reser.AddReservationOrderRequest) (*reser.AddReservationOrderResponse, error)
	GetReservationOrder(context.Context, *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error)
	UpdateReservationOrder(context.Context, *reser.UpdateReservationOrderRequest) (*reser.UpdateReservationOrderResponse, error)
	DeleteReservationOrder(context.Context, *reser.DeleteReservationOrderRequest) (*reser.DeleteReservationOrderResponse, error)
	GetReservationsOrders(context.Context, *reser.GetReservationOrdersRequest) (*reser.GetReservationOrdersResponse, error)
}

type reservationOrderImpl struct {
	repo *repository.ReservationOrderRespository
}

func NewReservationOrderService(repo *repository.ReservationOrderRespository) ReservationOrderService {
	return &reservationOrderImpl{repo: repo}
}

func (rs *reservationOrderImpl) GetReservationOrder(ctx context.Context, req *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error) {
	return rs.repo.GetReservationOrderById(ctx, req)
}

func (rs *reservationOrderImpl) AddReservationOrder(ctx context.Context, req *reser.AddReservationOrderRequest) (*reser.AddReservationOrderResponse, error) {
	return rs.repo.CreateReservationOrder(ctx, req)
}

func (rs *reservationOrderImpl) UpdateReservationOrder(ctx context.Context, req *reser.UpdateReservationOrderRequest) (*reser.UpdateReservationOrderResponse, error) {
	return rs.repo.UpdateReservationOrder(ctx, req)
}

func (rs *reservationOrderImpl) DeleteReservationOrder(ctx context.Context, req *reser.DeleteReservationOrderRequest) (*reser.DeleteReservationOrderResponse, error) {
	return rs.repo.DeleteReservationOrder(ctx, req)
}

func (rs *reservationOrderImpl) GetReservationsOrders(ctx context.Context, req *reser.GetReservationOrdersRequest) (*reser.GetReservationOrdersResponse, error) {
	return rs.repo.GetAllReservationOrders(ctx, req)
}
