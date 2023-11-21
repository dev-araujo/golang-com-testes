package main

import (
	"fmt"
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}
	return differentHellos(lang) + name

}

func differentHellos(lang string) (prefixo string) {
	switch lang {
	case spanish:
		prefixo = hola
	case french:
		prefixo = bonjour
	default:
		prefixo = hello
	}
	return
}

func main() {
	fmt.Println(Hello("world", "lang"))
}
