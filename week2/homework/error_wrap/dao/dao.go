package dao

import (
	"geektime-go-advance/week2/homework/error_wrap/mockDB"
	"github.com/pkg/errors"
)

func QueryUser() ([]int, error) {
	data, err := mockDB.ErrQuery()

	return data, errors.Wrap(err, "query user failed")
}
