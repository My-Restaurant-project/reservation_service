package repository

import (
	"context"
	reser "reservation_service/genproto/reservation_service"

	"github.com/jmoiron/sqlx"
)

type MenuRepository struct {
	db *sqlx.DB
}

func NewMenuRepo(db *sqlx.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) CreateMenu(ctx context.Context, mReq *reser.AddMenuRequest) (*reser.AddMenuResponse, error) {
	return nil, nil
}
g
func (r *MenuRepository) GetMenuById(ctx context.Context, req *reser.GetMenuRequest) (*reser.GetMenuResponse, error) {
	return nil, nil
}

func (r *MenuRepository) UpdateMenu(ctx context.Context, req *reser.UpdateMenuRequest) (*reser.UpdateMenuResponse, error) {
	return nil, nil
}

func (r *MenuRepository) DeleteMenu(ctx context.Context, req *reser.DeleteMenuRequest) (*reser.DeleteMenuResponse, error) {
	return nil, nil
}

func (r *MenuRepository) GetAllMenus(ctx context.Context, req *reser.GetMenusRequest) (*reser.GetMenusResponse, error) {
	return nil, nil
}
