package repository

import (
	"testing"

	pb "reservation_service/genproto/reservation_service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateRestaurant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReservationRepo(db)
	id := uuid.NewString()

	mock.ExpectExec("insert into restaurants").
		WithArgs(id, "Test Restaurant", "Test Address", "1234567890", "Test Description").
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.AddRestaurantRequest{
		Name:        "Test Restaurant",
		Address:     "Test Address",
		PhoneNumber: "1234567890",
		Description: "Test Description",
	}

	res, err := repo.CreateRestaurant(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, id, res.Id)
	assert.Equal(t, req.Name, res.Name)
	assert.Equal(t, req.Address, res.Address)
	assert.Equal(t, req.PhoneNumber, res.PhoneNumber)
	assert.Equal(t, req.Description, res.Description)
}

func TestGetRestaurantById(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReservationRepo(db)
	id := uuid.NewString()

	rows := sqlmock.NewRows([]string{"name", "address", "phone_number", "description"}).
		AddRow("Test Restaurant", "Test Address", "1234567890", "Test Description")
	mock.ExpectQuery("select name, address,phone_number, description from restaurants where id = ?").
		WithArgs(id).
		WillReturnRows(rows)

	req := &pb.GetRestaurantRequest{
		Id: id,
	}

	res, err := repo.GetRestaurantById(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, "Test Restaurant", res.Restaurant.Name)
	assert.Equal(t, "Test Address", res.Restaurant.Address)
	assert.Equal(t, "1234567890", res.Restaurant.PhoneNumber)
	assert.Equal(t, "Test Description", res.Restaurant.Description)
}

func TestGetAllRestaurants(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReservationRepo(db)

	rows := sqlmock.NewRows([]string{"id", "name", "address", "phone_number", "description"}).
		AddRow(uuid.NewString(), "Test Restaurant", "Test Address", "1234567890", "Test Description")
	mock.ExpectQuery("select id, name, address, phone_number, description from restaurants where deleted_at IS NULL").
		WillReturnRows(rows)

	req := &pb.GetRestaurantsRequest{}

	res, err := repo.GetAllRestaurants(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Len(t, res.Restaurant, 1)
}

func TestUpdateRestaurant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReservationRepo(db)
	id := uuid.NewString()

	mock.ExpectExec("update restaurants set name = ?, address = ?, phone_number = ?, description = ? where id = ?").
		WithArgs("Updated Restaurant", "Updated Address", "0987654321", "Updated Description", id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.UpdateRestaurantRequest{
		Id:          id,
		Name:        "Updated Restaurant",
		Address:     "Updated Address",
		PhoneNumber: "0987654321",
		Description: "Updated Description",
	}

	res, err := repo.UpdateRestaurant(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, req.Name, res.Restaurant.Name)
	assert.Equal(t, req.Address, res.Restaurant.Address)
	assert.Equal(t, req.PhoneNumber, res.Restaurant.PhoneNumber)
	assert.Equal(t, req.Description, res.Restaurant.Description)
}

func TestDeleteRestaurant(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	repo := NewReservationRepo(db)
	id := uuid.NewString()

	mock.ExpectExec("update restaurants set deleted_at = ? where id = ?").
		WithArgs(sqlmock.AnyArg(), id).
		WillReturnResult(sqlmock.NewResult(1, 1))

	req := &pb.DeleteRestaurantRequest{
		Id: id,
	}

	res, err := repo.DeleteRestaurant(req)
	assert.NoError(t, err)
	assert.NotNil(t, res)
	assert.True(t, res.Deleted)
}
