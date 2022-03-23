package model

import "regexp"

type Query struct {
	Tp    int
	Query string
}

func NewQuery(str string) *Query {
	tp, query := VerifyQueryFormat(str)
	return &Query{
		Tp:    tp,
		Query: query,
	}
}

func VerifyQueryFormat(query string) (int, string) {
	if checkEmail(query) {
		return Email, query
	}
	if checkPhone(query) {
		return Phone, query
	}
	if checkAddress(query) {
		return Address, query
	}
	return -1, ""
}

func checkEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

func checkPhone(phone string) bool {
	//regular := "^1(3\d|4[5-9]|5[0-35-9]|6[2567]|7[0-8]|8\d|9[0-35-9])\d{8}$"
	//
	//reg := regexp.MustCompile(regular)
	//return reg.MatchString(phone)
	return len(phone) == 11
}

func checkAddress(address string) bool {
	return len(address) == 34
}
