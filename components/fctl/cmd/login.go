package cmd

import (
	"context"
	"fmt"
	"net/url"

	fctl "github.com/formancehq/fctl/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
	"github.com/zitadel/oidc/pkg/client/rp"
	"github.com/zitadel/oidc/pkg/oidc"
)

type Dialog interface {
	DisplayURIAndCode(uri, code string)
}
type DialogFn func(uri, code string)

func (fn DialogFn) DisplayURIAndCode(uri, code string) {
	fn(uri, code)
}

func LogIn(ctx context.Context, dialog Dialog, relyingParty rp.RelyingParty) (*oidc.Tokens, error) {
	deviceCode, err := rp.GetDeviceCode(ctx, relyingParty)
	if err != nil {
		return nil, err
	}

	uri, err := url.Parse(deviceCode.GetVerificationUri())
	if err != nil {
		panic(err)
	}
	query := uri.Query()
	query.Set("user_code", deviceCode.GetUserCode())
	uri.RawQuery = query.Encode()

	dialog.DisplayURIAndCode(deviceCode.GetVerificationUri(), deviceCode.GetUserCode())

	if err := fctl.Open(uri.String()); err != nil {
		return nil, err
	}

	return rp.PollDeviceCode(ctx, deviceCode.GetDeviceCode(), deviceCode.GetInterval(), relyingParty)
}

func NewLoginCommand() *cobra.Command {
	return fctl.NewCommand("login",
		fctl.WithStringFlag(fctl.MembershipURIFlag, "", "service url"),
		fctl.WithHiddenFlag(fctl.MembershipURIFlag),
		fctl.WithShortDescription("Login"),
		fctl.WithArgs(cobra.ExactArgs(0)),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {

			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return err
			}

			profile := fctl.GetCurrentProfile(cmd, cfg)
			membershipUri, err := cmd.Flags().GetString(fctl.MembershipURIFlag)
			if err != nil {
				return err
			}
			if membershipUri == "" {
				membershipUri = profile.GetMembershipURI()
			}

			relyingParty, err := fctl.GetAuthRelyingParty(cmd, membershipUri)
			if err != nil {
				return err
			}

			ret, err := LogIn(cmd.Context(), DialogFn(func(uri, code string) {
				fmt.Fprintln(cmd.OutOrStdout(), "Please enter the following code on your browser:", code)
				fmt.Fprintln(cmd.OutOrStdout(), "Link:", uri)
			}), relyingParty)
			if err != nil {
				return err
			}

			profile.SetMembershipURI(membershipUri)
			profile.UpdateToken(ret.Token)

			currentProfileName := fctl.GetCurrentProfileName(cmd, cfg)

			cfg.SetCurrentProfile(currentProfileName, profile)

			pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln("Logged!")
			return cfg.Persist()
		}),
	)
}
