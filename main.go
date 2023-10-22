package main

import (
	"fmt"

	"github.com/tamathecxder/randomail"
)

func main() {
	privateDomain := "dev,id"

	testMail := randomail.GenerateRandomEmail()

	emails := randomail.GenerateRandomEmails(10)

	privateEmail, err := randomail.GenerateRandomEmailWithCustomDomain(privateDomain)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Private Email:", privateEmail)

	for _, email := range emails {
		fmt.Println("Generate Email:", email)
	}

	fmt.Println(testMail)
}
