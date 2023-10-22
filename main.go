package main

import (
	"fmt"

	"github.com/tamathecxder/randomail"
)

func main() {
	testMail := randomail.GenerateRandomEmail()

	emails := randomail.GenerateRandomEmails(10)

	for _, email := range emails {
		fmt.Println("Generate Email:", email)
	}

	fmt.Println(testMail)
}
