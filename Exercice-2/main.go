package main

import "fmt"

func convertion(args ...interface{}) string {
	var resulta string
	for i := 0; i < len(args); i++ {
		resulta += fmt.Sprintln(args[i])
	}
	return resulta
}

func main() {
	texte := convertion("Bonjour", 123, 1.2, true)
	fmt.Println(texte)
}
