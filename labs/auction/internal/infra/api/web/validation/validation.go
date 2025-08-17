package validation

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin/binding"
	en2 "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	validator_en "github.com/go-playground/validator/v10/translations/en"
	"vinizer4/go-expert-fullcycle/labs/auction/configuration/rest_err"
)

var (
	Validade = validator.New()
	transl   ut.Translator
)

func init() {
	if value, ok := binding.Validator.Engine().(*validator.Validate); ok {
		en := en2.New()
		enTransl := ut.New(en, en)
		transl, _ = enTransl.GetTranslator("en")
		validator_en.RegisterDefaultTranslations(value, transl)
	}
}

func ValidateErr(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidation validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewNotFoundError("Invalid type error")
	} else if errors.As(validation_err, &jsonValidation) {
		errorCauses := []rest_err.Causes{}

		for _, e := range validation_err.(validator.ValidationErrors) {
			errorCauses = append(errorCauses, rest_err.Causes{
				Field:   e.Field(),
				Message: e.Translate(transl),
			})
		}

		return rest_err.NewBadRequestError("Validation error", errorCauses...)
	} else {
		return rest_err.NewBadRequestError("Error trying to convert fields")
	}
}
