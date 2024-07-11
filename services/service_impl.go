package services

import (
	"context"
	"fmt"
	user "reservation_service/genproto/authentication_service"
	reser "reservation_service/genproto/reservation_service"
	repo "reservation_service/repository"

	"github.com/jmoiron/sqlx"
)

type MainService interface {
	RestaurantService() RestaurantService
	ReservationService() ReservationService
	ReservationOrderService() ReservationOrderService
	MenuService() MenuService
}

type mainServiceImpl struct {
	reser.UnimplementedReservationServiceServer
	user.UnimplementedAuthenticationServiceServer
	restaurantService       RestaurantService
	reservatiionService     ReservationService
	reservationOrderService ReservationOrderService
	menuService             MenuService
}

func NewMainService(db *sqlx.DB) *mainServiceImpl {
	return &mainServiceImpl{
		restaurantService:       NewRestaurantService(repo.NewRestaurantRepository(db)),
		reservatiionService:     NewReservationService(repo.NewReservationRepository(db)),
		reservationOrderService: NewReservationOrderService(repo.NewReservationOrderRepository(db)),
		menuService:             NewMenuService(repo.NewMenuRepository(db)),
	}
}

func (rs *mainServiceImpl) RestaurantService() RestaurantService {
	return rs.restaurantService
}

func (rs *mainServiceImpl) ReservationService() ReservationService {
	return rs.reservatiionService
}

func (rs *mainServiceImpl) ReservationOrderService() ReservationOrderService {
	return rs.reservationOrderService
}

func (rs *mainServiceImpl) MenuService() MenuService {
	return rs.menuService
}

// Restaurant Service implementation
func (rs *mainServiceImpl) AddRestaurant(ctx context.Context, resReq *reser.AddRestaurantRequest) (*reser.AddRestaurantResponse, error) {
	resp, err := rs.RestaurantService().AddRestaurant(ctx, resReq)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetRestaurant(ctx context.Context, req *reser.GetRestaurantRequest) (*reser.GetRestaurantResponse, error) {
	resp, err := rs.RestaurantService().GetRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateRestaurant(ctx context.Context, req *reser.UpdateRestaurantRequest) (*reser.UpdateRestaurantResponse, error) {
	resp, err := rs.RestaurantService().UpdateRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteRestaurant(ctx context.Context, req *reser.DeleteRestaurantRequest) (*reser.DeleteRestaurantResponse, error) {
	resp, err := rs.RestaurantService().DeleteRestaurant(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetRestaurants(ctx context.Context, req *reser.GetRestaurantsRequest) (*reser.GetRestaurantsResponse, error) {
	resp, err := rs.RestaurantService().GetRestaurants(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Reservation service implementation
func (rs *mainServiceImpl) AddReservation(ctx context.Context, resReq *reser.AddReservationRequest) (*reser.AddReservationResponse, error) {
	userId := resReq.UserId
	resId := resReq.RestaurantId

	_, err := rs.GetRestaurant(ctx, &reser.GetRestaurantRequest{Id: resId})
	if err != nil {
		return nil, fmt.Errorf("restaurant not found for id: %s", resId)
	}

	_, err = rs.GetProfileById(ctx, &user.UserIdRequest{Id: userId})
	if err != nil {
		return nil, fmt.Errorf("user not found for id: %s", userId)
	}

	resp, err := rs.ReservationService().AddReservation(ctx, resReq)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetReservation(ctx context.Context, req *reser.GetReservationRequest) (*reser.GetReservationResponse, error) {
	resp, err := rs.ReservationService().GetReservation(ctx, req)

	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateReservation(ctx context.Context, req *reser.UpdateReservationRequest) (*reser.UpdateReservationResponse, error) {
	resp, err := rs.ReservationService().UpdateReservation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteReservation(ctx context.Context, req *reser.DeleteReservationRequest) (*reser.DeleteReservationResponse, error) {
	resp, err := rs.ReservationService().DeleteReservation(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// getall
func (rs *mainServiceImpl) GetReservations(ctx context.Context, req *reser.GetReservationsRequest) (*reser.GetReservationsResponse, error) {
	resp, err := rs.ReservationService().GetReservations(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Rerervaton Order Service

func (rs *mainServiceImpl) AddReservationOrder(ctx context.Context, resReq *reser.AddReservationOrderRequest) (*reser.AddReservationOrderResponse, error) {
	resp, err := rs.ReservationOrderService().AddReservationOrder(ctx, resReq)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetReservationOrder(ctx context.Context, req *reser.GetReservationOrderRequest) (*reser.GetReservationOrderResponse, error) {
	resp, err := rs.ReservationOrderService().GetReservationOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateReservationOrder(ctx context.Context, req *reser.UpdateReservationOrderRequest) (*reser.UpdateReservationOrderResponse, error) {
	resp, err := rs.ReservationOrderService().UpdateReservationOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteReservationOrder(ctx context.Context, req *reser.DeleteReservationOrderRequest) (*reser.DeleteReservationOrderResponse, error) {
	resp, err := rs.ReservationOrderService().DeleteReservationOrder(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil

}

func (rs *mainServiceImpl) GetReservationsOrders(ctx context.Context, req *reser.GetReservationOrdersRequest) (*reser.GetReservationOrdersResponse, error) {
	resp, err := rs.ReservationOrderService().GetReservationsOrders(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Menu service implementation

func (rs *mainServiceImpl) Menu(ctx context.Context, req *reser.AddMenuRequest) (*reser.AddMenuResponse, error) {
	resp, err := rs.MenuService().AddMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetMenu(ctx context.Context, req *reser.GetMenuRequest) (*reser.GetMenuResponse, error) {
	resp, err := rs.MenuService().GetMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) UpdateMenu(ctx context.Context, req *reser.UpdateMenuRequest) (*reser.UpdateMenuResponse, error) {
	resp, err := rs.MenuService().UpdateMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) DeleteMenu(ctx context.Context, req *reser.DeleteMenuRequest) (*reser.DeleteMenuResponse, error) {
	resp, err := rs.MenuService().DeleteMenu(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (rs *mainServiceImpl) GetMenus(ctx context.Context, req *reser.GetMenusRequest) (*reser.GetMenusResponse, error) {
	resp, err := rs.MenuService().GetMenus(ctx, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
