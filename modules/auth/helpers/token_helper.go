package helpers

import (
	"math"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func JwtTimeDurationCalculator(t int) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Now().Add(time.Duration(int64(t) * int64(math.Pow10(9)))))
}

func JwtTimeRepeatAdapter(t int64) *jwt.NumericDate {
	return jwt.NewNumericDate(time.Unix(t, 0))
}
