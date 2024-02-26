/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/cobra"
)

// signCmd represents the sign command
var signCmd = &cobra.Command{
	Use:   "sign",
	Short: "Sign allows to use your private key and create JWT token",
	RunE:  sign,
}

var (
	privkeyStr string
	iss        string
	aud        string
	exp        int
)

func init() {
	rootCmd.AddCommand(signCmd)

	signCmd.Flags().StringVar(&privkeyStr, "privkey", "", "Private key to sign JWT")
	signCmd.MarkFlagRequired("privkey")
	signCmd.Flags().StringVar(&iss, "iss", "", "Issuer; your did:web")
	signCmd.MarkFlagRequired("iss")
	signCmd.Flags().StringVar(&aud, "aud", "", "Audience; PDS on which you want to register")
	signCmd.MarkFlagRequired("aud")
	signCmd.Flags().IntVar(&exp, "exp", 60, "Expire at; amount of second token will be alive")
}

func sign(cmd *cobra.Command, args []string) error {
	privkey, err := getPublicFromPrivate(privkeyStr)
	if err != nil {
		return err
	}

	jwt.MarshalSingleStringAsArray = false
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.RegisteredClaims{
		Issuer:    iss,
		Audience:  jwt.ClaimStrings{aud},
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(exp) * time.Second)),
	})

	tokenString, err := token.SignedString(privkey)
	if err != nil {
		return err
	}

	fmt.Println(tokenString)
	return nil
}
