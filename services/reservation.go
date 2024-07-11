package services

import (
	"context"
	reser "reservation_service/genproto/reservation_service"
	"reservation_service/repository"
)

type ReservationService interface {
	AddReservation(context.Context, *reser.AddReservationRequest) (*reser.AddReservationResponse, error)
	GetReservation(context.Context, *reser.GetReservationRequest) (*reser.GetReservationResponse, error)
	UpdateReservation(context.Context, *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error)
	DeleteReservation(context.Context, *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error)
	GetReservations(context.Context, *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error)
}

type reservationServiceImpl struct {
	repo *repository.ReservationRespoitory
}

func NewReservationService(repo *repository.ReservationRespoitory) ReservationService {
	return &reservationServiceImpl{repo: repo}
}

func (r *reservationServiceImpl) GetReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {
	return r.repo.GetAllReservations(ctx, req)
}

func (r *reservationServiceImpl) AddReservation(ctx context.Context, req *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	return r.repo.CreateReservation(ctx, req)
}

func (r *reservationServiceImpl) GetReservation(ctx context.Context, req *reser.GetReservationRequest) (*reser.GetReservationResponse, error) {
	return r.repo.GetReservationById(ctx, req)
}

func (r *reservationServiceImpl) UpdateReservation(ctx context.Context, req *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error) {
	return r.repo.UpdateReservation(ctx, req)
}

func (r *reservationServiceImpl) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	return r.repo.DeleteReservation(ctx, req)
}
