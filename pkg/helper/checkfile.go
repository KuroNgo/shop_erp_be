package helper

import (
	"path/filepath"
	"strings"
)

// IsMP3 is used for checked format file mp3
func IsMP3(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".mp3"
}

// IsMP4 is used for checked format file mp4
func IsMP4(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".mp4"
}

// IsExcel is used for checked format file excel
func IsExcel(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".xls" || ext == ".xlsx"
}

func IsImage(filename string) bool {
	extensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".svg"}
	ext := strings.ToLower(filename[len(filename)-4:])
	for _, e := range extensions {
		if ext == e {
			return true
		}
	}
	return false
}
