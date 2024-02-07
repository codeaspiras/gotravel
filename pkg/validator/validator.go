package validator

import (
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	translations "github.com/go-playground/validator/v10/translations/pt_BR"
)

type Validator interface {
	Validate(any interface{}) validator.ValidationErrorsTranslations
}

func New() Validator {
	ptBR := pt_BR.New()
	ut := ut.New(ptBR, ptBR)
	translator, _ := ut.GetTranslator("pt_BR")

	validator := validator.New()
	translations.RegisterDefaultTranslations(validator, translator)

	return &vr{
		v: validator,
		t: translator,
	}
}

type vr struct {
	v *validator.Validate
	t ut.Translator
}

func (v *vr) Validate(any interface{}) validator.ValidationErrorsTranslations {
	err := v.v.Struct(any)
	if err == nil {
		return nil
	}

	errs := err.(validator.ValidationErrors)
	return errs.Translate(v.t)
}
