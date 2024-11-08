package password

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
)

// HashPassword Done
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", fmt.Errorf("could not hash password %w", err)
	}
	return string(hashedPassword), nil
}

func VerifyPassword(hashedPassword string, candidatePassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(candidatePassword))
}

// Sanitize để tránh bị SQL injection
// SQL injection là một loại tấn công về cơ sở dữ liệu
// Hacker sẽ gửi một loạt các thông tin giả về cơ sở dữ liệu nhằm hủy hoại môi trường server
func Sanitize(data string) string {
	data = html.EscapeString(strings.TrimSpace(data))
	return data
}
