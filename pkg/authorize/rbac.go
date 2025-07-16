package authorize

func HasRole(role string, allowedRoles ...string) bool {
	for _, r := range allowedRoles {
		if r == role {
			return true
		}
	}
	return false
}
