package main

import (
	"errors"
	"fmt"
	"unicode"
)

func main() {

	p1 := "123456"
	p2 := "abcdefg"
	p3 := "abcdefg123456"
	p4 := "Abcdefg123456"
	p5 := "abcdefg123456."
	p6 := "Abcdefg123456+"
	p7 := "Abcdefg123456#"
	p8 := "Abcd12+"
	passwords := []string{p1, p2, p3, p4, p5, p6, p7, p8}
	for i := range passwords {
		ok := verifyPassword(passwords[i])
		if ok != nil {
			fmt.Println("验证不通过: "+passwords[i], "err: ", ok)
		} else {
			fmt.Println("验证通过: " + passwords[i])
		}
	}
}

const (
	ErrPasswordTooShort     = "密码长度小于 8 位"
	ErrPasswordTooLong      = "密码长度大于 32 位"
	ErrPasswordRequirements = "密码必须包含英文大小写、数字和特殊符号"
)

func verifyPassword(s string) error {
	if len(s) < 8 {
		return errors.New(ErrPasswordTooShort)
	}
	if len(s) > 32 {
		return errors.New(ErrPasswordTooLong)
	}
	var hasNumber, hasUpperCase, hasLowercase, hasSpecial bool
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			hasNumber = true
		case unicode.IsUpper(c):
			hasUpperCase = true
		case unicode.IsLower(c):
			hasLowercase = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			hasSpecial = true
		}
	}
	if !(hasNumber && hasUpperCase && hasLowercase && hasSpecial) {
		return errors.New(ErrPasswordRequirements)
	}
	return nil
}

func PasswordCheck(passwd string) error {

	indNum := [4]int{0, 0, 0, 0}
	spCode := []byte{'!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '=', '\\', '`', '_', '+', '|', '~', '[', '{', ']', '}', ':', ';', '\'', '"', '<', '>', ',', '.', '?', '/'}

	if len(passwd) < 6 {
		return errors.New("password too short")
	}

	passwdByte := []byte(passwd)

	for _, i := range passwdByte {

		notEnd := 0
		for _, s := range spCode {
			if i == s {
				indNum[3] = 1
				notEnd = 1
				break
			}
		}

		if notEnd != 1 {
			return errors.New("没有特殊符号")
		}

		if i >= 'A' && i <= 'Z' {
			indNum[0] = 1
			continue
		}

		if i >= 'a' && i <= 'z' {
			indNum[1] = 1
			continue
		}

		if i >= '0' && i <= '9' {
			indNum[2] = 1
			continue
		}

	}

	codeCount := 0

	for _, i := range indNum {
		codeCount += i
	}

	if codeCount < 3 {
		return errors.New("Too simple password")
	}

	return nil
}
