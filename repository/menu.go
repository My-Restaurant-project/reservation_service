package repository

import (
	"context"
	reser "reservation_service/genproto/reservation_service"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cast"
)

type MenuRepository struct {
	db *sqlx.DB
}

func NewMenuRepository(db *sqlx.DB) *MenuRepository {
	return &MenuRepository{db: db}
}

func (r *MenuRepository) CreateMenu(ctx context.Context, mReq *reser.AddMenuRequest) (*reser.AddMenuResponse, error) {
	id := uuid.NewString()

	query := `insert into menu(id, restaurant_id, name, description, price)values($1,$2,$3,$4,$5)`

	_, err := r.db.ExecContext(ctx, query, id, mReq.RestaurantId, mReq.Name, mReq.Description, mReq.Price)
	if err != nil {
		return nil, err
	}

	return &reser.AddMenuResponse{
		Id:           id,
		RestaurantId: mReq.RestaurantId,
		Name:         mReq.Name,
		Description:  mReq.Description,
		Price:        mReq.Price,
		CreatedAt:    cast.ToString(time.Now()),
	}, nil
}

func (r *MenuRepository) GetMenuById(ctx context.Context, req *reser.GetMenuRequest) (*reser.GetMenuResponse, error) {
	query := `select restaurant_id, name, description, price from menu where id=$1 and deleted_at is null`

	row := r.db.QueryRowContext(ctx, query, req.Id)

	var resp = &reser.GetMenuResponse{}
	var menu = &reser.Menu{}

	err := row.Scan(&menu.RestaurantId, &menu.Name, &menu.Description, &menu.Price)
	if err != nil {
		return nil, err
	}

	resp.Menu = menu

	return resp, nil
}

func (r *MenuRepository) UpdateMenu(ctx context.Context, req *reser.UpdateMenuRequest) (*reser.UpdateMenuResponse, error) {
	query := `
	UPDATE menu
	SET restaurant_id=$1, name=$2, description=$3, price=$4, updated_at=now()
	WHERE id=$5 and deleted_at is null
	RETURNING id, restaurant_id, name, description, price, created_at, updated_at
	`
	row := r.db.QueryRowContext(ctx, query, req.RestaurantId, req.Name, req.Description, req.Price, req.Id)
	var updMenuRes reser.UpdateMenuResponse
	var menu reser.Menu

	err := row.Scan(&menu.Id, &menu.RestaurantId, &menu.Name, &menu.Description, &menu.Price, &menu.CreatedAt, &menu.UpdatedAt)
	if err != nil {
		return &updMenuRes, err
	}
	updMenuRes.Menu = &menu

	

	return &updMenuRes, nil
}

func (r *MenuRepository) DeleteMenu(ctx context.Context, req *reser.DeleteMenuRequest) (*reser.DeleteMenuResponse, error) {
	query := `
	UPDATE menu
	SET deleted_at=now()
	WHERE id=$1 and deleted_at is null
	`

	_, err := r.db.ExecContext(ctx, query, req.Id)
	if err != nil {
		return &reser.DeleteMenuResponse{Deleted: false}, err
	}

	return &reser.DeleteMenuResponse{Deleted: true}, nil
}

func (r *MenuRepository) GetAllMenus(ctx context.Context, req *reser.GetMenusRequest) (*reser.GetMenusResponse, error) {
	query := `select id, restaurant_id, name, description, price from menu where deleted_at is null`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}

	var resp = &reser.GetMenusResponse{}
	var menus = []*reser.Menu{}

	for rows.Next() {
		var menu = &reser.Menu{}

		err := rows.Scan(&menu.Id, &menu.RestaurantId, &menu.Name, &menu.Description, &menu.Price)
		if err != nil {
			return nil, err
		}
		menus = append(menus, menu)
	}

	resp.Menus = menus

	return resp, nil

}
