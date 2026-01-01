package thefarm

import (
	"errors"
	"fmt"
)

type InvalidCowsError struct {
	cows    int
	details string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.details)
}

// TODO: define the 'DivideFood' function
func DivideFood(FodderCalculator FodderCalculator, cows int) (float64, error) {
	fooderAmount, err := FodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}

	fatteningFactor, err := FodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	food := float64(fooderAmount*fatteningFactor) / float64(cows)
	return food, nil
}

// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(FodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows < 1 {
		return 0, errors.New("invalid number of cows")
	}

	return DivideFood(FodderCalculator, cows)
}

// TODO: define the 'ValidateNumberOfCows' function
func ValidateNumberOfCows(cows int) error {
	var errMessage string
	if cows < 0 {
		return &InvalidCowsError{
			cows:    cows,
			details: "there are no negative cows",
		}
	} else if cows == 0 {
		return &InvalidCowsError{
			cows:    cows,
			details: "no cows don't need food",
		}
	} else {
		return nil
	}
	return errors.New(errMessage)
}

