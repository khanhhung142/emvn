package validator

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func InitValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("objectid", IsObjectId)
		return
	}
}

// validate does the field is a valid mongo object id
var IsObjectId validator.Func = func(fl validator.FieldLevel) bool {
	if fl.Field().String() == "" {
		return false
	}
	return IsMongoObjectId(fl.Field().String())
}

func IsMongoObjectId(s string) bool {
	_, err := primitive.ObjectIDFromHex(s)
	return err == nil
}
