package interfaces

import "github.com/Micledson/sdkbanco/domain"

type IAccountFunctions interface {
	GetAccountById(baseURL string, id int) (*domain.Account, *error)
	Deposit(baseURL string, id int, balance float64) *error
	Withdraw(baseURL string, id int, balance float64) *error
	GetAllAccount(baseURL string) ([]domain.Account, *error)
}
