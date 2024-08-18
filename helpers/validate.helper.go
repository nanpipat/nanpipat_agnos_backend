package helpers

import (
	"regexp"
	"strings"
)

var prefixFiles = []string{
	"jpg",
	"jpeg",
	"png",
	"gif",
	"bmp",
	"tif",
	"tiff",
	"svg",
	"doc",
	"docx",
	"xls",
	"xlsx",
	"csv",
	"ppt",
	"pptx",
	"pdf",
	"txt",
	"rtf",
	"vc",
	"crt",
	"key",
	"pub",
	"pem",
	"asc",
	"gpg",
	"ssh",
	"p12",
	"pfx",
}

// func IsValueValid[T comparable](key string, value T) core.IError {
// 	if reflect.ValueOf(value).IsZero() {
// 		return emsgs.ValueIsNotValid(key)
// 	}

// 	return nil
// }

// func IsPasswordValid(password string) (bool, string) {
// 	if len(password) < 8 {
// 		// return false, "Password must be at least 8 characters long."
// 		return false, consts.PasswordTooShort
// 	}

// 	if strings.Contains(`$'"`, password) {
// 		// return false, "Password must not contains $, ', or \"."
// 		return false, "Password must not contains $, ', or \"."
// 	}

// 	hasLower := false
// 	hasUpper := false
// 	hasNumber := false
// 	hasSymbol := false
// 	hasQuote := false
// 	hasSequential := false

// 	for _, r := range password {
// 		if unicode.IsLower(r) {
// 			hasLower = true
// 		} else if unicode.IsUpper(r) {
// 			hasUpper = true
// 		} else if unicode.IsDigit(r) {
// 			hasNumber = true
// 		} else if strings.ContainsRune("!@#%^&*()_+-=[]{}|;:,.<>?/", r) {
// 			hasSymbol = true
// 		}

// 		if strings.ContainsRune("'\"", r) {
// 			hasQuote = true
// 			break
// 		}

// 		// if i > 0 && (r == rune(password[i-1])+1 || r == rune(password[i-1])-1) {
// 		// 	if i > 1 && (rune(password[i-1]) == rune(password[i-2])+1 || rune(password[i-1]) == rune(password[i-2])-1) {
// 		// 		hasSequential = true
// 		// 		break
// 		// 	}
// 		// }

// 		// if i > 0 && r == rune(password[i-1]) {
// 		// 	if i > 1 && r == rune(password[i-2]) {
// 		// 		hasSequential = true
// 		// 		break
// 		// 	}
// 		// }
// 	}

// 	seqPass := PasswordSeq(password)
// 	if !seqPass {
// 		hasSequential = true
// 	}

// 	if !hasLower {
// 		// return false, "Password must contain at least 1 lowercase letter."
// 		return false, consts.PasswordRequiredLowercase

// 	}
// 	if !hasUpper {
// 		// return false, "Password must contain at least 1 uppercase letter."
// 		return false, consts.PasswordRequiredUppercase
// 	}
// 	if !hasNumber {
// 		// return false, "Password must contain at least 1 number."
// 		return false, consts.PasswordRequiredNumber
// 	}
// 	if !hasSymbol {
// 		// return false, "Password must contain at least 1 symbol."
// 		return false, consts.PasswordRequiredSymbol
// 	}
// 	if hasQuote {
// 		return false, "Password must not contain ' or \"."
// 	}
// 	if hasSequential {
// 		return false, "Password must not contain sequential or repeating characters."
// 	}

// 	return true, "Password is valid."
// }

// func PasswordSeq(password string) bool {
// 	// Check if password contains repeated characters more than four times
// 	for i, r := range password {

// 		if i > 0 && r == rune(password[i-1]) {
// 			if i > 1 && r == rune(password[i-2]) {
// 				if i > 2 && r == rune(password[i-3]) {
// 					return false
// 				}

// 			}
// 		}
// 	}

// 	// Check if password is in sequential order
// 	for i := 0; i < len(password)-3; i++ {
// 		if isSequential(password[i : i+4]) {
// 			return false
// 		}
// 	}

// 	return true
// }

func isSequential(substring string) bool {
	// Check if substring is in sequential order
	for i := 1; i < len(substring); i++ {
		if substring[i] != substring[i-1]+1 {
			return false
		}
	}

	return true
}

func IDCardValid(s string) bool {
	match, _ := regexp.MatchString("^[0-9]{13}$", s)
	return match
}

func EmailValid(s string) bool {
	emailRegex := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	return emailRegex.MatchString(s)
}

func FileValid(s string) bool {
	file := strings.Split(s, ".")
	if len(file) < 2 {
		return false
	}

	filePrefix := file[len(file)-1]

	valid := false
	for _, prefix := range prefixFiles {
		if prefix == filePrefix {
			valid = true
			break
		}
	}

	return valid
}
