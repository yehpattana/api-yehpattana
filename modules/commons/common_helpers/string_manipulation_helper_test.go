package commonhelpers

import "testing"

func TestReplacePercent20WithSpace(t *testing.T) {
	t.Run("when word included %20, should return replace %20 with space", func(t *testing.T) {
		testCases := []struct {
			wordInput string
			expected  string
		}{
			{
				wordInput: "DORAEMON123",
				expected:  "DORAEMON123",
			},
			{
				wordInput: "DORAEMON123%20",
				expected:  "DORAEMON123",
			},
			{
				wordInput: "DORAEMON123%20%20",
				expected:  "DORAEMON123",
			},
			{
				wordInput: "%20DORAEMON123%20%20",
				expected:  "DORAEMON123",
			},
			{
				wordInput: "DORAE%20MON123",
				expected:  "DORAE MON123",
			},
		}
		for _, tt := range testCases {
			if got := ReplacePercent20WithSpace(tt.wordInput); got != tt.expected {
				t.Errorf("ReplacePercent20WithSpace() = %v, want %v", got, tt.expected)
			}
		}
	})
}
