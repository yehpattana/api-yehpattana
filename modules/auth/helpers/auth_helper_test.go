package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCheckIsValidEmailPattern(t *testing.T) {
	t.Run("email must be in valid format", func(t *testing.T) {
		// Arrange
		testCases := []struct {
			testCase string
			value    string
			expected bool
		}{
			{
				testCase: "abc@mail.com",
				value:    "abc@mail.com",
				expected: true,
			},
			{
				testCase: "abc",
				value:    "abc",
				expected: false,
			},
			{
				testCase: "abc@",
				value:    "abc@",
				expected: false,
			},
			{
				testCase: "abc@mail",
				value:    "abc@mail",
				expected: false,
			},
		}

		// Act
		for _, tc := range testCases {
			t.Run(tc.testCase, func(t *testing.T) {
				// Act
				email := CheckIsValidEmailPattern(tc.value)

				// Assert
				if email != tc.expected {
					t.Errorf("Expected value does not match actual value for test case %s. Expected: %t, Got: %t", tc.testCase, tc.expected, email)
				}

				// Print email for visual inspection
				t.Logf("Email: %s, Expected: %t, Got: %t", tc.value, tc.expected, email)
			})
		}
	})
}

func TestBcryptHashingPassword(t *testing.T) {
	t.Run("password must be hashed and return hashed password", func(t *testing.T) {
		// Arrange
		password := "password123"

		// Act
		hashedPassword, err := BcryptHashingPassword(password)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if hashedPassword == "" {
			t.Errorf("expected non-empty string, got empty string")
			t.Log("password:", password)
		}
	},
	)
}

func TestBcryptComparePassword(t *testing.T) {
	t.Run("password must be compared and return nil", func(t *testing.T) {
		// Arrange
		password := "password123"
		hashedPassword, err := BcryptHashingPassword(password)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Act
		err = BcryptComparePassword(hashedPassword, password)

		// Assert
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	},
	)
}

func TestGeneratePassword(t *testing.T) {
	t.Run("password must be return correct value", func(t *testing.T) {
		// Arrange
		testCases := []struct {
			testCase string
			value    int
			expected int
		}{
			{
				testCase: "Length -1",
				value:    -1,
				expected: 6,
			},
			{
				testCase: "Length 0",
				value:    0,
				expected: 6,
			},
			{
				testCase: "Length 1",
				value:    1,
				expected: 6,
			},
			{
				testCase: "Length 5",
				value:    5,
				expected: 6,
			},
			{
				testCase: "Length 6",
				value:    6,
				expected: 6,
			},
			{
				testCase: "Length 10",
				value:    10,
				expected: 10,
			},
			{
				testCase: "Length 16",
				value:    16,
				expected: 16,
			},
			{
				testCase: "Length 20",
				value:    20,
				expected: 16,
			},
		}

		for _, tc := range testCases {
			t.Run(tc.testCase, func(t *testing.T) {
				// Act
				password := GeneratePassword(tc.value)

				// Assert
				assert.Equal(t, tc.expected, len(password), "Expected length does not match actual length.")

				// Print password for visual inspection
				t.Logf("Generated password (%d chars): %s", len(password), password)
			})
		}
	})
}

func TestIsHashedPassword(t *testing.T) {
	t.Run("password must be hashed and return true", func(t *testing.T) {
		// Arrange
		password := "password123"
		hashedPassword, err := BcryptHashingPassword(password)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		// Act
		isHashed := IsHashedPassword(hashedPassword)

		// Assert
		if !isHashed {
			t.Errorf("expected true, got false")
			t.Log("password:", password)
		}
	},
	)
}
