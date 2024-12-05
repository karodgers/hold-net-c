package server

// IsUsernameTaken checks if a username is already taken
func isUsernameTaken(name string) bool {
	for _, v := range clients {
		if v == name {
			return true
		}
	}
	return false
}
