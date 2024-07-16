package auth

import (
	"fmt"
	"time"

	"github.com/anfen93/goPasetoV4"
)

// Function for generating authentication token
func GenerateToken() {
	maker := goPasetoV4.NewPasetoMaker()

	// Create a token for a specific username with a 1-hour duration
	a, _, err := maker.CreateToken("username", time.Hour)
	if err != nil {
		fmt.Printf("Error creating token: %v", err)
	}

	// fmt.Printf("Generated Token: %v", a)

	_, err = maker.VerifyToken(a)
	if err != nil {
		fmt.Printf("Error verifying token: %v", err)
	}

	// fmt.Printf("Token Payload: %+v", payload.Username)
}
