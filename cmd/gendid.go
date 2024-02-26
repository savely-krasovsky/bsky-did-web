/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
)

// gendidCmd represents the gendid command
var gendidCmd = &cobra.Command{
	Use:   "gendid",
	Short: "Generates did.json for to serve",
	RunE:  generateDid,
}

var (
	handle   string
	pubkey   string
	hostname string
)

func init() {
	rootCmd.AddCommand(gendidCmd)

	gendidCmd.Flags().StringVar(&handle, "handle", "", "Pass your handle here")
	gendidCmd.MarkFlagRequired("handle")
	gendidCmd.Flags().StringVar(&pubkey, "pubkey", "", "Pass public key from pubkey command here")
	gendidCmd.MarkFlagRequired("pubkey")
	gendidCmd.Flags().StringVar(&hostname, "hostname", "", "Pass hostname of PDS you want to register")
	gendidCmd.MarkFlagRequired("hostname")
}

func generateDid(cmd *cobra.Command, args []string) error {
	did := &DID{
		Context: []string{
			"https://www.w3.org/ns/did/v1",
			"https://w3id.org/security/multikey/v1",
			"https://w3id.org/security/suites/secp256k1-2019/v1",
		},
		Id:          "did:web:" + handle,
		AlsoKnownAs: []string{"at://" + handle},
		VerificationMethod: []*VerificationMethod{{
			ID:                 "did:web:" + handle + "#athandle",
			Type:               "Multikey",
			Controller:         "did:web:" + handle,
			PublicKeyMultibase: pubkey,
		}},
		Service: []*Service{{
			ID:              "#atproto_pds",
			Type:            "AtprotoPersonalDataServer",
			ServiceEndpoint: "https://" + hostname,
		}},
	}

	b, err := json.MarshalIndent(did, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}

type DID struct {
	Context            []string              `json:"@context"`
	Id                 string                `json:"id"`
	AlsoKnownAs        []string              `json:"alsoKnownAs"`
	VerificationMethod []*VerificationMethod `json:"verificationMethod"`
	Service            []*Service            `json:"service"`
}

type VerificationMethod struct {
	ID                 string `json:"id"`
	Type               string `json:"type"`
	Controller         string `json:"controller"`
	PublicKeyMultibase string `json:"publicKeyMultibase"`
}

type Service struct {
	ID              string `json:"id"`
	Type            string `json:"type"`
	ServiceEndpoint string `json:"serviceEndpoint"`
}
