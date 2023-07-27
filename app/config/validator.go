package config

import (
	"fmt"
	"reflect"
	"strings"
	"sync"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
)

var (
	vOnce    sync.Once
	validate *validator.Validate
	transEn  ut.Translator
	v        *Validator
	uni      *ut.UniversalTranslator
)

type (
	CustomValidator interface {
		Validate(validate *validator.Validate) error
	}

	Validator struct {
		validator *validator.Validate
	}
)

func (v *Validator) Validate(i any) error {
	err := v.validator.Struct(i)
	if err != nil {
		return fmt.Errorf(buildErrorMessage(err))
	}

	return nil
}

func buildErrorMessage(err error) string {
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}

	var errMsg strings.Builder
	for i, ve := range ves {
		if i > 0 {
			errMsg.WriteString("; ")
		}
		errMsg.WriteString(ve.Translate(transEn))
	}

	return errMsg.String()
}

func GetValidator() *Validator {
	vOnce.Do(func() {
		uni = ut.New(en.New())
		transEn, _ = uni.GetTranslator("en")
		validate = validator.New()

		validate.RegisterTagNameFunc(func(field reflect.StructField) string {
			return strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
		})

		_ = enTrans.RegisterDefaultTranslations(validate, transEn)
		_ = registerCustomTagValidation()

		v = &Validator{validator: validate}
	})

	return v
}

func registerCustomTagValidation() error {
	customTagValidation := []struct {
		tagName             string
		validationFunc      func(fl validator.FieldLevel) bool
		validateEvenNull    bool
		translationTemplate string
		translationOverride bool
	}{
		// override required tag
		{
			tagName:             "required",
			translationTemplate: "{0} is a required field",
			translationOverride: true,
		},
	}

	for _, s := range customTagValidation {
		if s.validationFunc != nil {
			err := validate.RegisterValidation(s.tagName, s.validationFunc, s.validateEvenNull)
			if err != nil {
				return err
			}
		}

		err := validate.RegisterTranslation(s.tagName, transEn, registrationFunc(s.tagName, s.translationTemplate, s.translationOverride), translateFunc)
		if err != nil {
			return err
		}
	}

	return nil
}

func registrationFunc(tag string, translation string, override bool) validator.RegisterTranslationsFunc {
	return func(ut ut.Translator) error {
		return ut.Add(tag, translation, override)
	}
}

func translateFunc(ut ut.Translator, fe validator.FieldError) string {
	i := strings.Index(fe.Namespace(), ".")
	t, err := ut.T(fe.Tag(), fe.Namespace()[i+1:])
	if err != nil {
		return fe.(error).Error()
	}
	return t
}
