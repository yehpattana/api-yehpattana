package userhelpers

import "regexp"

type UserHelper interface {
}

type userHelperImpl struct {
}

func UserHelperImpl() UserHelper {
	return &userHelperImpl{}
}

func CheckIsValidPassword(password string) bool {
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!#$%&*()_`"
	pattern := "[" + regexp.QuoteMeta(charset) + "]"
	regex := regexp.MustCompile(pattern)

	if len(password) >= 6 && len(password) <= 16 && regex.MatchString(password) {
		return true
	}
	return false
}
