package validation

import (
	"regexp"

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
