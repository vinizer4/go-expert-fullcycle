package cep

import (
	"fmt"
	"regexp"
)

func IsValid(cep string) (bool, error) {
	matched, err := regexp.MatchString(`^\d{8}$`, cep)
	if err != nil {
		return false, fmt.Errorf("erro enquanto validava o CEP :: %v", err)
	}
	return matched, nil
}
