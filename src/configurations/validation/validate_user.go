package validation

import (
	"errors"
	"peruccii/goapi/src/configurations/rest_err"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translation "github.com/go-playground/validator/v10/translations/en"
	"github.com/goccy/go-json"
)

// declaring a global variable
var (
    Validate = validator.New()
    transl     ut.Translator
)

// func `init` always run with de main
func init() {
    if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
        
        en := en.New()
        un := ut.New(en, en)
        transl, _ = un.GetTranslator("en")
        en_translation.RegisterDefaultTranslations(val, transl)
    }
}

func ValidateUserError(validation_err error) *rest_err.RestErr {
    
    var jsonErr             *json.UnmarshalTypeError    // jsonErr of type *json.UnmarshalTypeError 
    var jsonValidationError validator.ValidationErrors  


    if errors.As(validation_err, &jsonErr) { // <-- validate if field types is wrong
        return rest_err.NewBadRequestError("Invalid field type")

    } else if errors.As(validation_err, &jsonValidationError) { // <-- validate error based on binding binding struct
        errorCauses := []rest_err.Causes{} // array of Causes

        for _, e := range jsonValidationError {
            cause := rest_err.Causes{
                Field: e.Field(),
                Message: e.Translate(transl),
            }

            errorCauses = append(errorCauses, cause)
        }

        return rest_err.NewBadRequestValidationError("Some fields are invalid", errorCauses)
    } else {
        return rest_err.NewBadRequestError("Error trying to convert fields")
    }
}
