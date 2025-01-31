package request

import (
	"github.com/go-playground/validator/v10"
	"math"
	"regexp"
	"strconv"
)

type Validator struct {
	Validator *validator.Validate
}

func (v *Validator) Validate(data interface{}) error {
	_ = v.Validator.RegisterValidation("valid_phone", validatePhoneNumber)
	_ = v.Validator.RegisterValidation("valid_store", validStoreID)
	_ = v.Validator.RegisterValidation("ge", ge)
	_ = v.Validator.RegisterValidation("eq", eq)
	_ = v.Validator.RegisterValidation("gt", gt)
	_ = v.Validator.RegisterValidation("city", func(fl validator.FieldLevel) bool {
		return validRangeCheck(fl, 1, 64)
	})
	_ = v.Validator.RegisterValidation("area", func(fl validator.FieldLevel) bool {
		return validRangeCheck(fl, 1, 100)
	})
	_ = v.Validator.RegisterValidation("zone", func(fl validator.FieldLevel) bool {
		return validRangeCheck(fl, 1, 100)
	})
	return v.Validator.Struct(data)
}
func validatePhoneNumber(field validator.FieldLevel) bool {
	return regexp.MustCompile(`^(01)[3-9]{1}[0-9]{8}$`).MatchString(field.Field().String())
}
func validRangeCheck(field validator.FieldLevel, l float64, u float64) bool {
	f := field.Field()
	var fieldValue float64
	if f.CanFloat() {
		fieldValue = f.Float()
	} else if f.CanInt() {
		fieldValue = float64(f.Int())
	} else {
		return false
	}
	return l <= fieldValue && fieldValue <= u
}
func validStoreID(field validator.FieldLevel) bool {
	f := field.Field()
	var fieldValue float64
	if f.CanFloat() {
		fieldValue = f.Float()
	} else if f.CanInt() {
		fieldValue = float64(f.Int())
	} else {
		return false
	}
	return fieldValue == 131172
}
func ge(field validator.FieldLevel) bool {
	v, err := strconv.ParseFloat(field.Param(), 64)
	if err != nil {
		return false
	}
	f := field.Field()
	var fieldValue float64
	if f.CanFloat() {
		fieldValue = f.Float()
	} else if f.CanInt() {
		fieldValue = float64(f.Int())
	} else {
		return false
	}
	return math.Abs(fieldValue) >= math.Abs(v)
}

func gt(field validator.FieldLevel) bool {
	v, err := strconv.ParseFloat(field.Param(), 64)
	if err != nil {
		return false
	}
	f := field.Field()
	var fieldValue float64
	if f.CanFloat() {
		fieldValue = f.Float()
	} else if f.CanInt() {
		fieldValue = float64(f.Int())
	} else {
		return false
	}
	return math.Abs(fieldValue) > math.Abs(v)
}

func eq(field validator.FieldLevel) bool {
	v, err := strconv.ParseFloat(field.Param(), 64)
	if err != nil {
		return false
	}
	f := field.Field()
	var fieldValue float64
	if f.CanFloat() {
		fieldValue = f.Float()
	} else if f.CanInt() {
		fieldValue = float64(f.Int())
	} else {
		return false
	}
	return math.Abs(fieldValue) == math.Abs(v)
}

func ValidationMsg(err error) map[string][]string {
	ve := make(map[string][]string)
	for _, e := range err.(validator.ValidationErrors) {
		if em := errorMsg(e); em != "" {
			ve[e.Field()] = []string{
				em,
			}
		}
	}
	return ve
}

func errorMsg(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "The " + e.Field() + " field is required."
	case "ge":
		return "The " + e.Field() + " must be greater than or equal " + e.Param() + "."
	case "city":
		return "The City must be between 1-64." // let's assume total 64 cities
	case "area":
		return "The Area must be between 1-100." // let's assume total 100 areas
	case "zone":
		return "The Zone must be between 1-100." // let's assume total 100 zones
	case "gt":
		return "The " + e.Field() + " must be greater than " + e.Param() + "."
	case "eq":
		return "The " + e.Field() + " must be " + e.Param() + "."
	case "valid_phone":
		return e.Field() + " is not valid phone."
	case "valid_store":
		return "Wrong Store selected."
	default:
		return ""
	}
}
