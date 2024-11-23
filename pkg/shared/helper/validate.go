package helper

import (
	"regexp"
	"unicode"
)

func AlphabetOnlyRegex(s string) bool {
	// Tạo biểu thức chính quy để kiểm tra chỉ chứa ký tự chữ
	re := regexp.MustCompile(`^[A-Za-z]+$`)
	return re.MatchString(s)
}

// EmailValid Không trả về json
func EmailValid(email string) bool {
	// Biểu thức chính quy kiểm tra địa chỉ email theo định dạng thông thường
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(emailRegex).MatchString(email)
}

func PhoneValid(phone string) bool {
	// Biểu thức chính quy để kiểm tra số điện thoại hợp lệ với đúng 10 chữ số
	pattern := `^(84|0[35789])\d{8}$`
	regex := regexp.MustCompile(pattern)

	return regex.MatchString(phone)
}

// PasswordStrong Không trả về Json
func PasswordStrong(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasLower := false
	hasUpper := false
	hasDigit := false

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	return hasLower && hasUpper && hasDigit
}
