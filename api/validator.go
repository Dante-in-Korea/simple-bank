package api

import (
	"github.com/Dante-in-Korea/simple-bank/util"
	"github.com/go-playground/validator/v10"
)

var validaCurrency validator.Func = func(filedLevel validator.FieldLevel) bool {
	if currency, ok := filedLevel.Field().Interface().(string); ok {
		// check currency is supported
		return util.IsSupportedCurrency(currency)
	}
	return false
}
