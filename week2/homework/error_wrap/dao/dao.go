package dao

import (
	"geektime-go-advance/week2/homework/error_wrap/mockDB"
	"github.com/pkg/errors"
)

func QueryUser() ([]int, error) {
	data, err := mockDB.ErrQuery()
	if err != nil {
		err = errors.Wrap(err, "query user failed")
	}
	return data, err
}
