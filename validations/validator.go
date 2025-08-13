package validations

import (
	"regexp"
	"twitter-clone-go/request"

	"github.com/go-playground/validator/v10"
)

func HasKigou(fl validator.FieldLevel) bool {
	pw := fl.Field().String()
	hasKigou := regexp.MustCompile(`[-_!?]`).MatchString(pw)
	return hasKigou
}

func HasHanSu(fl validator.FieldLevel) bool {
	pw := fl.Field().String()
	hasSu := regexp.MustCompile(`[0-9]`).MatchString(pw)
	return hasSu
}

func HasLowerEi(fl validator.FieldLevel) bool {
	pw := fl.Field().String()
	hasLowerEi := regexp.MustCompile(`[a-z]`).MatchString(pw)
	return hasLowerEi
}

func HasUpperEi(fl validator.FieldLevel) bool {
	pw := fl.Field().String()
	hasUpperEi := regexp.MustCompile(`[A-Z]`).MatchString(pw)

	return hasUpperEi
}

func ValidateSignUpInfo(data request.SignUpInfo) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	validate.RegisterValidation("has_kigou", HasKigou)
	validate.RegisterValidation("has_han_su", HasHanSu)
	validate.RegisterValidation("has_lower_ei", HasLowerEi)
	validate.RegisterValidation("has_upper_ei", HasUpperEi)

	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
