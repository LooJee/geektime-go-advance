package mockDB

import "errors"

var (
	ErrNoRows = errors.New("not found")
)

func ErrQuery() ([]int, error) {
	return nil, ErrNoRows
}
