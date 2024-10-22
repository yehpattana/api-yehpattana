package configs

import (
	"time"
)

type ServiceConfigInterface interface {
	Url() string // host:port
	Name() string
	Version() string
	ReadTimeout() time.Duration
	WriteTimeout() time.Duration
	BodyLimit() int
	FileLimit() int
	Host() string
	Port() int
}

type DbConfigInterface interface {
	Url() string
	MaxOpenConnections() int
}

type JwtConfigInterface interface {
	SecretKey() []byte
	AdminKey() []byte
	ApiKey() []byte
	AccessExpiresAt() int
	RefreshExpiresAt() int
	SetJwtAccessExpires(t int)
	SetJwtRefreshExpires(t int)
}

type CloudinaryConfigInterface interface {
	CloudName() string
	ApiKey() string
	ApiSecret() string
	CloudinaryBaseURL() string
}
