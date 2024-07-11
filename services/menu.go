package services

import (
	"context"
	reser "reservation_service/genproto/reservation_service"
	"reservation_service/repository"
)

type MenuService interface {
	AddMenu(context.Context, *reser.AddMenuRequest) (*reser.AddMenuResponse, error)
	GetMenu(context.Context, *reser.GetMenuRequest) (*reser.GetMenuResponse, error)
	UpdateMenu(context.Context, *reser.UpdateMenuRequest) (*reser.UpdateMenuResponse, error)
	DeleteMenu(context.Context, *reser.DeleteMenuRequest) (*reser.DeleteMenuResponse, error)
	GetMenus(context.Context, *reser.GetMenusRequest) (*reser.GetMenusResponse, error)
}

type menuServiceImpl struct {
	repo *repository.MenuRepository
}

func NewMenuService(repo *repository.MenuRepository) MenuService {
	return &menuServiceImpl{repo: repo}
}

func (ms *menuServiceImpl) AddMenu(ctx context.Context, req *reser.AddMenuRequest) (*reser.AddMenuResponse, error) {
	return ms.repo.CreateMenu(ctx, req)
}

func (ms *menuServiceImpl) GetMenu(ctx context.Context, req *reser.GetMenuRequest) (*reser.GetMenuResponse, error) {
	return ms.repo.GetMenuById(ctx, req)
}

func (ms *menuServiceImpl) UpdateMenu(ctx context.Context, req *reser.UpdateMenuRequest) (*reser.UpdateMenuResponse, error) {
	return ms.repo.UpdateMenu(ctx, req)
}

func (ms *menuServiceImpl) DeleteMenu(ctx context.Context, req *reser.DeleteMenuRequest) (*reser.DeleteMenuResponse, error) {
	return ms.repo.DeleteMenu(ctx, req)
}

func (ms *menuServiceImpl) GetMenus(ctx context.Context, req *reser.GetMenusRequest) (*reser.GetMenusResponse, error) {
	return ms.repo.GetAllMenus(ctx, req)
}
