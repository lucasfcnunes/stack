package install

import (
	"github.com/formancehq/fctl/cmd/payments/connectors/internal"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pkg/errors"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewModulrCommand() *cobra.Command {
	const (
		endpointFlag = "endpoint"

		defaultEndpoint = "https://api-sandbox.modulrfinance.com"
	)
	return fctl.NewCommand(internal.ModulrConnector+" <api-key> <api-secret>",
		fctl.WithShortDescription("Install a Modulr connector"),
		fctl.WithArgs(cobra.ExactArgs(2)),
		fctl.WithStringFlag(endpointFlag, defaultEndpoint, "API endpoint"),
		fctl.WithRunE(func(cmd *cobra.Command, args []string) error {
			cfg, err := fctl.GetConfig(cmd)
			if err != nil {
				return err
			}

			organizationID, err := fctl.ResolveOrganizationID(cmd, cfg)
			if err != nil {
				return err
			}

			stack, err := fctl.ResolveStack(cmd, cfg, organizationID)
			if err != nil {
				return err
			}

			if !fctl.CheckStackApprobation(cmd, stack, "You are about to install connector '%s'", internal.ModulrConnector) {
				return fctl.ErrMissingApproval
			}

			paymentsClient, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return err
			}

			var endpoint *string
			if e := fctl.GetString(cmd, endpointFlag); e != "" {
				endpoint = &e
			}

			_, err = paymentsClient.Payments.InstallConnector(cmd.Context(), operations.InstallConnectorRequest{
				ConnectorConfig: shared.ConnectorConfig{
					ModulrConfig: &shared.ModulrConfig{
						APIKey:    args[0],
						APISecret: args[1],
						Endpoint:  endpoint,
					},
				},
				Connector: internal.ModulrConnector,
			})

			pterm.Success.WithWriter(cmd.OutOrStdout()).Printfln("Connector installed!")

			return errors.Wrap(err, "installing connector")
		}),
	)
}
