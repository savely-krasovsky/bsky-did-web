/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

// genkeyCmd represents the genkey command
var genkeyCmd = &cobra.Command{
	Use:   "genkey",
	Short: "Generates secp256r1 private key and prints it as hex",
	RunE:  generateKey,
}

func init() {
	rootCmd.AddCommand(genkeyCmd)
}

func generateKey(cmd *cobra.Command, args []string) error {
	crv := elliptic.P256()
	privkey, err := ecdsa.GenerateKey(crv, rand.Reader)
	if err != nil {
		return err
	}

	fmt.Println(hex.EncodeToString(privkey.D.Bytes()))
	return nil
}
