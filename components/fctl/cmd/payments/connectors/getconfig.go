package connectors

import (
	"github.com/formancehq/fctl/cmd/payments/connectors/internal"
	fctl "github.com/formancehq/fctl/pkg"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func NewGetConfigCommand() *cobra.Command {
	return fctl.NewCommand("get-config <connector-name>",
		fctl.WithAliases("getconfig", "getconf", "gc", "get", "g"),
		fctl.WithArgs(cobra.ExactArgs(1)),
		fctl.WithValidArgs("stripe"),
		fctl.WithShortDescription("Read a connector config"),
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

			client, err := fctl.NewStackClient(cmd, cfg, stack)
			if err != nil {
				return err
			}

			request := operations.ReadConnectorConfigRequest{
				Connector: shared.ConnectorEnum(args[0]),
			}
			connectorConfig, err := client.Payments.ReadConnectorConfig(cmd.Context(), request)
			if err != nil {
				return fctl.WrapError(err, "reading connector config")
			}
			switch args[0] {
			case internal.StripeConnector:
				err = displayStripeConfig(cmd, connectorConfig.ConnectorConfigResponse.Data)
			case internal.ModulrConnector:
				err = displayModulrConfig(cmd, connectorConfig.ConnectorConfigResponse.Data)
			case internal.BankingCircleConnector:
				err = displayBankingCircleConfig(cmd, connectorConfig.ConnectorConfigResponse.Data)
			case internal.CurrencyCloudConnector:
				err = displayCurrencyCloudConfig(cmd, connectorConfig.ConnectorConfigResponse.Data)
			case internal.WiseConnector:
				err = displayWiseConfig(cmd, connectorConfig.ConnectorConfigResponse.Data)
			default:
				pterm.Error.WithWriter(cmd.OutOrStderr()).Printfln("Connection unknown.")
			}
			return err
		}),
	)
}

func displayStripeConfig(cmd *cobra.Command, connectorConfig shared.ConnectorConfig) error {
	config := connectorConfig.StripeConfig

	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("API key:"), config.APIKey})

	if err := pterm.DefaultTable.
		WithWriter(cmd.OutOrStdout()).
		WithData(tableData).
		Render(); err != nil {
		return err
	}
	return nil
}

func displayModulrConfig(cmd *cobra.Command, connectorConfig shared.ConnectorConfig) error {
	config := connectorConfig.ModulrConfig

	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("API key:"), config.APIKey})
	tableData = append(tableData, []string{pterm.LightCyan("API secret:"), config.APISecret})
	tableData = append(tableData, []string{pterm.LightCyan("Endpoint:"), func() string {
		if config.Endpoint == nil {
			return ""
		}
		return *config.Endpoint
	}()})

	if err := pterm.DefaultTable.
		WithWriter(cmd.OutOrStdout()).
		WithData(tableData).
		Render(); err != nil {
		return err
	}
	return nil
}

func displayWiseConfig(cmd *cobra.Command, connectorConfig shared.ConnectorConfig) error {
	config := connectorConfig.WiseConfig

	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("API key:"), config.APIKey})

	if err := pterm.DefaultTable.
		WithWriter(cmd.OutOrStdout()).
		WithData(tableData).
		Render(); err != nil {
		return err
	}
	return nil
}

func displayBankingCircleConfig(cmd *cobra.Command, connectorConfig shared.ConnectorConfig) error {
	config := connectorConfig.BankingCircleConfig

	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("Username:"), config.Username})
	tableData = append(tableData, []string{pterm.LightCyan("Password:"), config.Password})
	tableData = append(tableData, []string{pterm.LightCyan("Endpoint:"), config.Endpoint})
	tableData = append(tableData, []string{pterm.LightCyan("Authorization endpoint:"), config.AuthorizationEndpoint})

	if err := pterm.DefaultTable.
		WithWriter(cmd.OutOrStdout()).
		WithData(tableData).
		Render(); err != nil {
		return err
	}
	return nil
}

func displayCurrencyCloudConfig(cmd *cobra.Command, connectorConfig shared.ConnectorConfig) error {
	config := connectorConfig.CurrencyCloudConfig

	tableData := pterm.TableData{}
	tableData = append(tableData, []string{pterm.LightCyan("API key:"), config.APIKey})
	tableData = append(tableData, []string{pterm.LightCyan("Login ID:"), config.LoginID})
	tableData = append(tableData, []string{pterm.LightCyan("Endpoint:"), func() string {
		if config.Endpoint == nil {
			return ""
		}
		return *config.Endpoint
	}()})

	if err := pterm.DefaultTable.
		WithWriter(cmd.OutOrStdout()).
		WithData(tableData).
		Render(); err != nil {
		return err
	}
	return nil
}
