package commonhelpers

import "testing"

func TestRoleIdConverter(t *testing.T) {
	type args struct {
		role int
	}
	tests := []struct {
		name     string
		args     args
		expected string
	}{
		{
			name: "Role ID 1 - Customer",
			args: args{
				role: 1,
			},
			expected: "Customer",
		},
		{
			name: "Role ID 2 - Admin",
			args: args{
				role: 2,
			},
			expected: "Admin",
		},
		{
			name: "Role Id 3 - Unknown",
			args: args{
				role: 3,
			},
			expected: "Unknown",
		},
		{
			name: "Role Id 0 - Unknown",
			args: args{
				role: 0,
			},
			expected: "Unknown",
		},
		{
			name: "Role Id 999 - Unknown",
			args: args{
				role: 999,
			},
			expected: "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RoleIdConverter(tt.args.role); got != tt.expected {
				t.Errorf("RoleIdConverter() = %v, want %v", got, tt.expected)
			}
		})
	}
}
