package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type SillyNephewError struct {
	c int
}

func (s *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", s.c)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	fodderAmount, err := weightFodder.FodderAmount()

	switch {
	case err == ErrScaleMalfunction && fodderAmount > 0:
		return fodderAmount * 2.0 / float64(cows), nil

	case fodderAmount < 0 && (err == ErrScaleMalfunction || err == nil):
		return 0, errors.New("negative fodder")

	case err != nil:
		return 0, err

	case cows == 0:
		return 0, errors.New("division by zero")

	case cows < 0:
		return 0, &SillyNephewError{c: cows}

	default:
		return fodderAmount / float64(cows), nil
	}
}
