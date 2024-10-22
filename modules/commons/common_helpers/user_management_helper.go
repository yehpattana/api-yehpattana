package commonhelpers

func RoleIdConverter(role int) string {
	switch role {
	case 1:
		return "Customer"
	case 2:
		return "Admin"
	}
	return "Unknown"
}
