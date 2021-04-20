package main

import (
	"fmt"
	"geektime-go-advance/week2/homework/error_wrap/dao"
	"geektime-go-advance/week2/homework/error_wrap/mockDB"
	"github.com/pkg/errors"
)

func main() {
	_, err := dao.QueryUser()

	if errors.As(err, &mockDB.ErrNoRows) {
		fmt.Printf("%+v", err)
	}
}
