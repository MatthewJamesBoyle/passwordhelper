package passwordhelper

import (
	"errors"
	"strings"
)

type Service struct {
}

var (
	ErrEmptyPass    = errors.New("err_password_empty")
	ErrInvalidIndex = errors.New("err_invalid_index")
)

func (s *Service) CharsAt(password string, index ...int) (*string, error) {
	passLength := len(password)
	if passLength == 0 {
		return nil, ErrEmptyPass
	}

	var r []byte
	for _, v := range index {
		if v <= 0 || v > passLength {
			return nil, ErrInvalidIndex
		}
		r = append(r, password[v-1])
		r = append(r, ',')
	}

	res := strings.Trim(string(r), ",")
	return &res, nil
}
