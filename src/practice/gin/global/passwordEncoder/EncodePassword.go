package passwordEncoder

import "golang.org/x/crypto/bcrypt"

func EncodePassword(rawPassword string) (string, error) {
	password, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 10)
	if err != nil {
		return "", err
	}
	return string(password), nil
}
