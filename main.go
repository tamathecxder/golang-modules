package main

import (
	"fmt"

	"github.com/tamathecxder/randomail/v2"
)

func main() {
	privateDomain := "dev.id"

	testMail := randomail.GenerateRandomEmail(50)

	emails := randomail.GenerateRandomEmails(10)

	privateEmail, err := randomail.GenerateRandomEmailWithCustomDomain(privateDomain)

	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("Private Email:", privateEmail)
	fmt.Println("Test:", testMail)

	for _, email := range emails {
		fmt.Println("Generate Email:", email)
	}

}
