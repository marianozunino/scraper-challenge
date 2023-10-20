package adapter

import (
	"errors"
	"fmt"
	"time"

	"github.com/jaswdr/faker"
	"github.com/marianozunino/go-scrap-challenge/internal/core/domain"
	"github.com/marianozunino/go-scrap-challenge/internal/port"
)

var _ port.Service = &service{}

type service struct {
	users            []domain.User
	orders           []domain.Order
	ordersWithoutCSV []int
}

func NewService() port.Service {
	s := service{}

	s.users = []domain.User{
		{Username: "jdoe", Password: "jdoe", ID: 1},
	}

	s.orders = []domain.Order{}

	fake := faker.New()

	// Create 100 orders
	for i := 1; i <= 100; i++ {
		order := domain.Order{
			ID:      i,
			Buyer:   fake.Person().FirstName(),
			Date:    fake.Time().TimeBetween(time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC), time.Now()),
			Details: []domain.ItemLine{},
		}

		// Create 1 to 10 items per order
		for j := 1; j <= fake.IntBetween(1, 10); j++ {
			itemLine := domain.ItemLine{
				ID:          j,
				Quantity:    fake.IntBetween(1, 100),
				Description: fake.Lorem().Sentence(5),
				Price:       fake.IntBetween(1, 100),
			}

			order.Details = append(order.Details, itemLine)
		}

		s.orders = append(s.orders, order)
	}

	s.ordersWithoutCSV = []int{7, 13, 25, 31, 42, 54, 61, 77, 86, 99}

	return &s
}

// GetOrders implements port.Service.
func (s *service) GetOrders(page int, pageSize int) ([]domain.Order, int64) {
	start := (page - 1) * pageSize
	end := start + pageSize

	if start > len(s.orders) {
		return []domain.Order{}, 0
	}

	if end > len(s.orders) {
		end = len(s.orders)
	}

	return s.orders[start:end], int64(len(s.orders))
}

// Login implements port.Service.
func (s *service) Login(username string, password string) bool {
	for _, user := range s.users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}

// GetOrder implements port.Service.
func (s *service) GetOrderCSV(id int) ([]byte, error) {
	if !s.idHasCSV(id) {
		return []byte{}, errors.New("This order has no Details")
	}

	var order domain.Order

	for _, o := range s.orders {
		if o.ID == id {
			order = o
		}
	}

	if order.ID == 0 {
		return []byte{}, errors.New("order not found")
	}

	csv := "id,description,quantity,price\n"
	for _, item := range order.Details {
		csv += fmt.Sprintf("%d,%s,%d,%d\n", item.ID, item.Description, item.Quantity, item.Price)
	}

	return []byte(csv), nil
}

func (s *service) idHasCSV(id int) bool {
	for _, i := range s.ordersWithoutCSV {
		if i == id {
			return false
		}
	}
	return true
}
