package service

import (
	"fmt"
)

func checkParameters(parameters interface{}, message string) error {
	if parameters == nil {
		return fmt.Errorf(message)
	}
	return nil
}
