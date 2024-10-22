package commonhelpers

import "testing"

func TestGenerateUUID(t *testing.T) {
	t.Run("uuid must be return correct value", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if result == "" {
			t.Errorf("expected non-empty string, got empty string")
		}
	},
	)

	t.Run("returns a string of length 36", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if len(result) != 36 {
			t.Errorf("expected string of length 36, got string of length %d", len(result))
		}
	},
	)

	t.Run("returns a string in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GenerateUUID()

		// Assert
		if len(result) != 36 {
			t.Errorf("expected string of length 36, got string of length %d", len(result))
		}
		if result[8] != '-' || result[13] != '-' || result[18] != '-' || result[23] != '-' {
			t.Errorf("expected string in the format xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx, got %s", result)
		}
	},
	)
}

func TestCheckIsValidUUID(t *testing.T) {
	t.Run("valid uuid must return true", func(t *testing.T) {
		// Arrange
		testCases := []struct {
			uuid     string
			expected bool
		}{
			{
				uuid:     "d1b42c6e-888d-4bf3-8df7-ad27140858fb",
				expected: true,
			},
			{
				uuid:     "2b7e0e32-500e-4398-a821-2d969458abeb",
				expected: true,
			},
			{
				uuid:     "8baa25e6-dad5-4227-ad63-2b02d6f46c52",
				expected: true,
			},
			{
				uuid:     "27064f52-86d6-43c4-b265-805b8bf08928",
				expected: true,
			},
			{
				uuid:     "27064f52-86d6-43c4-b265-805b8bf0892",
				expected: false,
			},
		}
		// Act & Assert
		for _, tt := range testCases {
			if got := CheckIsValidUUID(tt.uuid); got != tt.expected {
				t.Errorf("CheckIsValidUUID() = %v, want %v", got, tt.expected)
			}
		}
	},
	)
}
