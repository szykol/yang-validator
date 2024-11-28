package validator

import (
	"fmt"

	"github.com/openconfig/goyang/pkg/yang"
	"github.com/szykol/yang-validator/pkg/domain"
)

func ValidateYang(input domain.ValidationInput) error {
	_, err := yang.Parse(string(input.Contents), input.FileName)
	if err != nil {
		return fmt.Errorf("Invalid YANG: %w", err)
	}

	return nil
}
