package passwordEncoder

import "golang.org/x/crypto/bcrypt"

func MatchPassword(rawPassword, encodedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword))
	if err == nil {
		return true
	} else {
		return false
	}
}
