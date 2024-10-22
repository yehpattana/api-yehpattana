package commonhelpers

import "testing"

func TestGetCurrentTimeISO(t *testing.T) {
	t.Run("returns a string in the format 2006-01-02T15:04:05Z", func(t *testing.T) {
		// Arrange - Nothing to arrange
		// Act
		result := GetCurrentTimeISO()

		// Assert
		if len(result) != 20 {
			t.Errorf("expected string of length 20, got string of length %d", len(result))
		}
		if result[4] != '-' || result[7] != '-' || result[10] != 'T' || result[13] != ':' || result[16] != ':' || result[19] != 'Z' {
			t.Errorf("expected string in the format 2006-01-02T15:04:05Z, got %s", result)
		}
	},
	)
}
