package middleware

// FiberMiddleware provide Fiber's built-in middlewares.
// See: https://docs.gofiber.io/api/middleware

type FindAccessToken struct {
	userId      string `json:"user_id"`
	accessToken string `json:"access_token"`
}
