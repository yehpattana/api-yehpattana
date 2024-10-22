package userhelpers

import "testing"

func TestCheckIsValidPassword(t *testing.T) {
	t.Run("password must be valid", func(t *testing.T) {
		// Arrange
		testCases := []struct {
			testCase string
			value    string
			expected bool
		}{
			{
				testCase: "valid password - no special characters 11 characters",
				value:    "password123",
				expected: true,
			},
			{
				testCase: "valid password - strongpassword 12 characters",
				value:    "k2ljF*kjPr)P",
				expected: true,
			},
			{
				testCase: "valid password - simplepassword and no number inclueded 8 characters",
				value:    "password",
				expected: true,
			},
			{
				testCase: "invalid password - less than 8 characters",
				value:    "passwor",
				expected: true,
			},
			{
				testCase: "invalid password - more than 20 characters",
				value:    "passwordpasswordpasswordpassword",
				expected: false,
			},
			{
				testCase: "invalid password - less than 6 characters",
				value:    "pass",
				expected: false,
			},
		}

		// Act
		for _, tc := range testCases {
			t.Run(tc.testCase, func(t *testing.T) {
				// Act
				password := CheckIsValidPassword(tc.value)

				// Assert
				if password != tc.expected {
					t.Errorf("Expected value does not match actual value for test case %s. Expected: %t, Got: %t", tc.testCase, tc.expected, password)
				}

				// Print password for visual inspection
				t.Logf("Password: %s, Expected: %t, Got: %t", tc.value, tc.expected, password)
			})
		}
	})
}
