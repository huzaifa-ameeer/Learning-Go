package greet

import "strings"

func Hello(s string) string {
	n := normalizeName(s)
	return  "Hello " + n
}

func normalizeName(s string) string {
	name := strings.TrimSpace(s)
	if(name ==""){
		return "Error"
	}
	return strings.ToUpper(name)
}