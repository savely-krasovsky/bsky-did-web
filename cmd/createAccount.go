/*
Copyright © 2024 Savely Krasovsky <savely@krasovs.ky>
*/
package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strings"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	"github.com/bluesky-social/indigo/util/cliutil"
	"github.com/bluesky-social/indigo/xrpc"
	"github.com/spf13/cobra"
)

// createAccountCmd represents the createAccount command
var createAccountCmd = &cobra.Command{
	Use:   "createAccount",
	Short: "Allows to create account at chosen PDS",
	RunE:  createAccount,
}

var (
	pds        string
	invideCode string
	email      string
	password   string
)

func init() {
	rootCmd.AddCommand(createAccountCmd)

	createAccountCmd.Flags().StringVar(&pds, "pds", "", "PDS URL")
	createAccountCmd.MarkFlagRequired("pds")
	createAccountCmd.Flags().StringVar(&handle, "handle", "", "Your handle")
	createAccountCmd.MarkFlagRequired("handle")
	createAccountCmd.Flags().StringVar(&invideCode, "invite", "", "Invite code")
	createAccountCmd.MarkFlagRequired("invite")
	createAccountCmd.Flags().StringVar(&email, "email", "", "Initial email")
	createAccountCmd.MarkFlagRequired("email")
	createAccountCmd.Flags().StringVar(&password, "password", "", "Initial password")
	createAccountCmd.MarkFlagRequired("password")
}

func createAccount(cmd *cobra.Command, args []string) error {
	rawStr, err := io.ReadAll(cmd.InOrStdin())
	if err != nil {
		return err
	}

	jwtToken := strings.TrimSpace(string(rawStr))

	client := &xrpc.Client{
		Client: cliutil.NewHttpClient(),
		Host:   hostname,
		Auth: &xrpc.AuthInfo{
			AccessJwt: jwtToken,
			Handle:    handle,
			Did:       "did:web:" + handle,
		},
	}

	did := "did:web:" + handle
	acc, err := comatproto.ServerCreateAccount(context.TODO(), client, &comatproto.ServerCreateAccount_Input{
		Did:        &did,
		Email:      &email,
		Handle:     handle,
		InviteCode: &invideCode,
		Password:   &password,
	})
	if err != nil {
		return err
	}

	b, err := json.MarshalIndent(acc, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(b))
	return nil
}
