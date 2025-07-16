package helpers

import (
	"fmt"
	"strings"


	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Fungsi untuk translate error validasi ke map
func TranslateErrorMessage(err error) map[string]string {
	errorsMap := make(map[string]string)

	if validatorErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldError := range validatorErrors {
			field := fieldError.Field()
			switch fieldError.Tag() {
			case "required":
				errorsMap[field] = fmt.Sprintf("%s is required", field)
			case "email":
				errorsMap[field] = "Invalid email format"
			case "unique":
				errorsMap[field] = fmt.Sprintf("%s already exists", field)
			case "min":
				errorsMap[field] = fmt.Sprintf("%s must be at least %s characters", field, fieldError.Param())
			case "max":
				errorsMap[field] = fmt.Sprintf("%s must be at most %s characters", field, fieldError.Param())
			case "numeric":
				errorsMap[field] = fmt.Sprintf("%s must be a number", field)
			default:
				errorsMap[field] = "Invalid value"
			}
		}
	}
if err != nil {
	if strings.Contains(err.Error(), "Duplicate entry") {
		if strings.Contains(err.Error(), "username") {
			errorsMap["Username"] = "Username already exists"
		}
		if strings.Contains(err.Error(), "email") {
			errorsMap["Email"] = "Email already exists"
		} 
	}else if err == gorm.ErrRecordNotFound {
		errorsMap["Error"] = "Record not found"
	}
}
	return errorsMap
}

func IsDuplicateEntryError(err error) bool {
	return err != nil && strings.Contains(err.Error(), "Duplicate entry")
}