package dto

import "math/big"

type User struct {
	Id               int
	Username         string `db:"username"`
	Email            string `db:"email"`
	PhoneNumber      string `db:"phone"`
	RegistrationData big.Int
	Password         string
}
