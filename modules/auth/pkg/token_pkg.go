package pkg

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yehpattana/api-yehpattana/configs"
	"github.com/yehpattana/api-yehpattana/modules/auth/helpers"
	commonstring "github.com/yehpattana/api-yehpattana/modules/commons/common_string"
	"github.com/yehpattana/api-yehpattana/modules/data/entities"
)

type TokenType string

const (
	AccessToken  TokenType = "access_token"
	RefreshToken TokenType = "refresh_token"
	Admin        TokenType = "admin"
	ApiKey       TokenType = "api_key"
)

type YptAuthInterface interface {
	SignToken() string
}

type yptAuth struct {
	mapClaims *yptMapClaims
	cfg       configs.JwtConfigInterface
}

type yptAdmin struct {
	*yptAuth
}

type yptApiKey struct {
	*yptAuth
}

type yptMapClaims struct {
	Claims *entities.UserClaims `json:"claims"`
	jwt.RegisteredClaims
}

func NewYptAuth(cfg configs.JwtConfigInterface, tokenType TokenType, claims *entities.UserClaims) (YptAuthInterface, error) {
	switch tokenType {
	case AccessToken:
		return newAccessToken(cfg, claims), nil
	case RefreshToken:
		return newRefreshToken(cfg, claims), nil
	case Admin:
		return newAdminToken(cfg), nil
	case ApiKey:
		return newApiKey(cfg), nil
	default:
		return nil, fmt.Errorf(commonstring.UnknowTokenType)
	}
}

func newAccessToken(cfg configs.JwtConfigInterface, claims *entities.UserClaims) YptAuthInterface {
	return &yptAuth{
		cfg: cfg,
		mapClaims: &yptMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    commonstring.ServiceName,
				Subject:   commonstring.AccessToken,
				Audience:  []string{"customer", "admin"},
				ExpiresAt: helpers.JwtTimeDurationCalculator(cfg.AccessExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func newRefreshToken(cfg configs.JwtConfigInterface, claims *entities.UserClaims) YptAuthInterface {
	return &yptAuth{
		cfg: cfg,
		mapClaims: &yptMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    commonstring.ServiceName,
				Subject:   commonstring.RefreshToken,
				Audience:  []string{"customer", "admin"},
				ExpiresAt: helpers.JwtTimeDurationCalculator(cfg.RefreshExpiresAt()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
}

func (a *yptAuth) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	signedToken, _ := token.SignedString(a.cfg.SecretKey())
	return signedToken
}

func (a *yptAdmin) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	signedToken, _ := token.SignedString(a.cfg.AdminKey())
	return signedToken
}

func (a *yptApiKey) SignToken() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, a.mapClaims)
	signedToken, _ := token.SignedString(a.cfg.ApiKey())
	return signedToken
}

func ParseToken(cfg configs.JwtConfigInterface, tokenString string) (*yptMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &yptMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(commonstring.SigningMedthodIsInvalid)
		}
		return cfg.SecretKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf(commonstring.TokenFormatIsInvalid)
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf(commonstring.TokenIsExpired)
		} else {
			return nil, fmt.Errorf(commonstring.ParseTokenFailed)
		}
	}

	if claims, ok := token.Claims.(*yptMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf(commonstring.ClaimsTypeIsInvalid)
	}
}

func ParseAdminToken(cfg configs.JwtConfigInterface, tokenString string) (*yptMapClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &yptMapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf(commonstring.SigningMedthodIsInvalid)
		}
		return cfg.AdminKey(), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, fmt.Errorf(commonstring.TokenFormatIsInvalid)
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, fmt.Errorf(commonstring.TokenIsExpired)
		} else {
			return nil, fmt.Errorf(commonstring.ParseTokenFailed)
		}
	}

	if claims, ok := token.Claims.(*yptMapClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf(commonstring.ClaimsTypeIsInvalid)
	}
}

func RepeatToken(cfg configs.JwtConfigInterface, claims *entities.UserClaims, exp int64) string {
	obj := &yptAuth{
		cfg: cfg,
		mapClaims: &yptMapClaims{
			Claims: claims,
			RegisteredClaims: jwt.RegisteredClaims{
				Issuer:    commonstring.ServiceName,
				Subject:   commonstring.RefreshToken,
				Audience:  []string{"customer", "admin"},
				ExpiresAt: helpers.JwtTimeRepeatAdapter(exp),
				NotBefore: jwt.NewNumericDate(time.Now()),
				IssuedAt:  jwt.NewNumericDate(time.Now()),
			},
		},
	}
	return obj.SignToken()
}

func newAdminToken(cfg configs.JwtConfigInterface) YptAuthInterface {
	return &yptAdmin{
		&yptAuth{
			cfg: cfg,
			mapClaims: &yptMapClaims{
				Claims: nil,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    commonstring.ServiceName,
					Subject:   commonstring.AccessToken,
					Audience:  []string{"admin"},
					ExpiresAt: helpers.JwtTimeDurationCalculator(300),
					NotBefore: jwt.NewNumericDate(time.Now()),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				},
			},
		},
	}
}

func newApiKey(cfg configs.JwtConfigInterface) YptAuthInterface {
	return &yptApiKey{
		&yptAuth{
			cfg: cfg,
			mapClaims: &yptMapClaims{
				Claims: nil,
				RegisteredClaims: jwt.RegisteredClaims{
					Issuer:    commonstring.ServiceName,
					Subject:   commonstring.AccessToken,
					Audience:  []string{"admin", "customer"},
					ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(2, 0, 0)),
					NotBefore: jwt.NewNumericDate(time.Now()),
					IssuedAt:  jwt.NewNumericDate(time.Now()),
				},
			},
		},
	}
}
