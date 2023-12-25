//go:generate rm -fr mocks
//go:generate mockery --all

package validator

import (
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	universaltranslator "github.com/go-playground/universal-translator"
	base "github.com/go-playground/validator/v10"
	translations_en "github.com/go-playground/validator/v10/translations/en"
)

func BindAndValidate[T any](c *gin.Context, dto *T) (bool, map[string][]string) {
	v, ok := binding.Validator.Engine().(*base.Validate)

	if ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}

	ut := universaltranslator.New(en.New())
	translator, _ := ut.GetTranslator("en")
	translations_en.RegisterDefaultTranslations(v, translator)

	valid := true
	errs := make(map[string][]string)

	if err := c.ShouldBind(&dto); err != nil {
		valid = false
		var verr base.ValidationErrors
		if errors.As(err, &verr) {
			for _, f := range verr {
				err := f.ActualTag()
				if f.Param() != "" {
					err = fmt.Sprintf("%s=%s", err, f.Param())
				}
				errs[f.Field()] = []string{f.Translate(translator)}
			}

			return valid, errs
		}

		return valid, errs
	}

	return valid, errs
}

func BindAndValidateWithAbort[T any](c *gin.Context, dto *T) bool {
	if valid, err := BindAndValidate(c, dto); !valid {
		c.JSON(http.StatusBadRequest, DefaultInvalidInputResponse(err))
		c.Abort()
		return false
	}

	return true
}

type response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func DefaultInvalidInputResponse(errs map[string][]string) response {
	return response{
		Success: false,
		Message: "invalid data",
		Data:    errs,
	}
}
