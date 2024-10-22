package commonhelpers

import (
	"regexp"

	"github.com/google/uuid"
)

func GenerateUUID() string {
	uuid := uuid.New()
	return uuid.String()
}

func CheckIsValidUUID(uuid string) bool {
	// Regular expression for UUID v1, v2, v3, v4, and v5 (universally unique identifier)
	uuidRegex := regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)
	return uuidRegex.MatchString(uuid)
}
