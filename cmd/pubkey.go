/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"fmt"
	"io"
	"math/big"
	"slices"
	"strings"

	"github.com/multiformats/go-multibase"
	"github.com/spf13/cobra"
)

// pubkeyCmd represents the pubkey command
var pubkeyCmd = &cobra.Command{
	Use:   "pubkey",
	Short: "Generates secp256r1 public key from private key passed to stdin and prints it as base58btc",
	RunE:  getPublicKey,
}

func init() {
	rootCmd.AddCommand(pubkeyCmd)
}

func getPublicKey(cmd *cobra.Command, args []string) error {
	rawStr, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return err
	}

	privkey, err := getPublicFromPrivate(string(rawStr))
	if err != nil {
		return err
	}

	b := slices.Concat[[]byte]([]byte{0x80, 0x24}, elliptic.MarshalCompressed(elliptic.P256(), privkey.PublicKey.X, privkey.PublicKey.Y))
	pubkey, err := multibase.Encode(multibase.Base58BTC, b)
	if err != nil {
		return err
	}

	fmt.Println(pubkey)
	return nil
}

func getPublicFromPrivate(str string) (*ecdsa.PrivateKey, error) {
	crv := elliptic.P256()
	privkey := &ecdsa.PrivateKey{
		PublicKey: ecdsa.PublicKey{
			Curve: crv,
		},
		D: new(big.Int),
	}

	b, err := hex.DecodeString(strings.TrimSpace(str))
	if err != nil {
		return nil, err
	}

	privkey.D = new(big.Int).SetBytes(b)
	privkey.PublicKey.X, privkey.PublicKey.Y = privkey.PublicKey.Curve.ScalarBaseMult(privkey.D.Bytes())

	return privkey, nil
}
