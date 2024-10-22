package configs

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

// func FiberConfig() fiber.Config {

// 	// Define server settings.
// 	readTimeoutSecondsCount, _ := strconv.Atoi(os.Getenv("SERVER_READ_TIMEOUT"))

// 	// Return Fiber configuration.
// 	return fiber.Config{
// 		ReadTimeout: time.Second * time.Duration(readTimeoutSecondsCount),
// 	}
// }

func ServerConfig(path string) ConfigInterface {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("Error loading .env file %v", err)
	}

	server := buildServiceConfig(envMap)
	db := buildDbConfig(envMap)
	jwt := buildJwtConfig(envMap)

	return &config{
		service: server,
		db:      db,
		jwt:     jwt,
	}
}

func buildServiceConfig(envMap map[string]string) *service {
	return &service{
		host:         envMap["SERVER_HOST"],
		port:         getIntEnv(envMap, "SERVER_PORT"),
		name:         envMap["SERVER_NAME"],
		version:      envMap["SERVER_VERSION"],
		readTimeout:  getDurationEnv(envMap, "SERVER_READ_TIMEOUT"),
		writeTimeout: getDurationEnv(envMap, "SERVER_WRITE_TIMEOUT"),
		bodyLimit:    getIntEnv(envMap, "SERVER_BODY_LIMIT"),
		fileLimit:    getIntEnv(envMap, "SERVER_FILE_LIMIT"),
	}
}

func buildDbConfig(envMap map[string]string) *db {
	return &db{
		host:           envMap["DB_HOST"],
		port:           getIntEnv(envMap, "DB_PORT"),
		protocol:       envMap["DB_PROTOCOL"],
		username:       envMap["DB_USERNAME"],
		password:       envMap["DB_PASSWORD"],
		database:       envMap["DB_DATABASE"],
		sslMode:        envMap["DB_SSL_MODE"],
		maxConnections: getIntEnv(envMap, "DB_MAX_CONNECTIONS"),
	}
}

func buildJwtConfig(envMap map[string]string) *jwt {
	return &jwt{
		secertKey:        envMap["JWT_SECRET_KEY"],
		adminKey:         envMap["JWT_ADMIN_KEY"],
		apiKey:           envMap["JWT_API_KEY"],
		accessExpiresAt:  getIntEnv(envMap, "JWT_ACCESS_TOKEN_EXPIRES"),
		refreshExpiresAt: getIntEnv(envMap, "JWT_REFRESH_TOKEN_EXPIRES"),
	}
}

func buildCloudinaryConfig(envMap map[string]string) *cdr {
	return &cdr{
		cloudName: envMap["CLOUDINARY_CLOUD_NAME"],
		apiKey:    envMap["CLOUDINARY_API_KEY"],
		apiSecret: envMap["CLOUDINARY_API_SECRET"],
		cloudinaryBaseUrl: envMap["CLOUDINARY_BASE_URL"],
	}
}

func getIntEnv(envMap map[string]string, key string) int {
	val, err := strconv.Atoi(envMap[key])
	if err != nil {
		msg := fmt.Sprintf("Failed to load %s from .env %v", key, err)
		log.Println(msg)
	}
	return val
}

func getDurationEnv(envMap map[string]string, key string) time.Duration {
	val := getIntEnv(envMap, key)
	return time.Duration(val) * time.Second
}

type ConfigInterface interface {
	Service() ServiceConfigInterface
	Db() DbConfigInterface
	Jwt() JwtConfigInterface
	Cdr() CloudinaryConfigInterface
}

type config struct {
	service *service
	db      *db
	jwt     *jwt
	cdr     *cdr
}

func (c *config) Service() ServiceConfigInterface {
	return c.service
}

func (c *config) Db() DbConfigInterface {
	return c.db
}

func (c *config) Jwt() JwtConfigInterface {
	return c.jwt
}

func (c *config) Cdr() CloudinaryConfigInterface {
	return c.cdr
}