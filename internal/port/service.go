package port

import "github.com/marianozunino/go-scrap-challenge/internal/core/domain"

type Service interface {
	Login(username string, password string) bool
	GetOrders(page int, pageSize int) ([]domain.Order, int64)
	GetOrderCSV(orderID int) ([]byte, error)
}
