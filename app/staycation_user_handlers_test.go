package app

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"staycation/dto"
	"staycation/errs"
	"staycation/mocks/service"
	"testing"
)

var router *mux.Router
var uh StaycationUserHandler
var mockService *service.MockStaycationUserService

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)
	mockService = service.NewMockStaycationUserService(ctrl)

	uh = StaycationUserHandler{service: mockService}

	router = mux.NewRouter()
	router.HandleFunc("/users", uh.getAllStaycationUsers)

	return func() {
		router = nil
		defer ctrl.Finish()
	}
}

func TestShouldReturnUsersWithStatusOk(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	dummyUsers := []dto.StaycationUserResponse{
		{
			Id:          "1",
			Name:        "John",
			City:        "Zagreb",
			Zipcode:     "10000",
			DateOfBirth: "2020-01-01",
			Status:      "1",
		},
		{
			Id:          "2",
			Name:        "Doe",
			City:        "Split",
			Zipcode:     "10002",
			DateOfBirth: "2020-01-01",
			Status:      "1",
		},
	}
	mockService.EXPECT().GetAllStaycationUsers().Return(dummyUsers, nil)

	request, _ := http.NewRequest(http.MethodGet, "/users", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusOK {
		t.Error("Returned wrong status code")
	}
}

func TestShouldReturnInternalServerError(t *testing.T) {
	// Arrange
	teardown := setup(t)
	defer teardown()

	mockService.EXPECT().GetAllStaycationUsers().Return(nil, errs.NewInternalServerError("Error"))

	request, _ := http.NewRequest(http.MethodGet, "/users", nil)

	// Act
	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, request)

	// Assert
	if recorder.Code != http.StatusInternalServerError {
		t.Error("Returned wrong status code")
	}
}
