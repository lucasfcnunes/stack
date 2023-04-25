package wallet

import (
	"context"

	sdk "github.com/formancehq/formance-sdk-go"
	"github.com/formancehq/formance-sdk-go/pkg/models/operations"
	"github.com/formancehq/formance-sdk-go/pkg/models/shared"
	"github.com/formancehq/stack/libs/go-libs/metadata"
)

type ListAccountsQuery struct {
	Cursor   string
	Limit    int
	Metadata map[string]any
}

type ListTransactionsQuery struct {
	Cursor      string
	Limit       int
	Metadata    map[string]any
	Destination string
	Source      string
	Account     string
}

type Ledger interface {
	AddMetadataToAccount(ctx context.Context, ledger, account string, metadata metadata.Metadata) error
	GetAccount(ctx context.Context, ledger, account string) (*sdk.Ledger.AccountWithVolumesAndBalances, error)
	ListAccounts(ctx context.Context, ledger string, query ListAccountsQuery) (*sdk.AccountsCursorResponseCursor, error)
	ListTransactions(ctx context.Context, ledger string, query ListTransactionsQuery) (*sdk.TransactionsCursorResponseCursor, error)
	RunScript(ctx context.Context, ledger string, script sdk.Script) (*operations.RunScriptResponse, error)
}

type DefaultLedger struct {
	client *sdk.Formance
}

func (d DefaultLedger) ListTransactions(ctx context.Context, ledger string, query ListTransactionsQuery) (*sdk.TransactionsCursorResponseCursor, error) {
	var (
		ret *sdk.TransactionsCursorResponse
		err error
	)
	if query.Cursor == "" {
		//nolint:bodyclose
		ret, _, err = d.client.TransactionsApi.ListTransactions(ctx, ledger).
			Metadata(query.Metadata).
			PageSize(int64(query.Limit)).
			Destination(query.Destination).
			Account(query.Account).
			Source(query.Source).
			Execute()
	} else {
		//nolint:bodyclose
		ret, _, err = d.client.TransactionsApi.ListTransactions(ctx, ledger).
			Cursor(query.Cursor).
			Execute()
	}
	if err != nil {
		return nil, err
	}

	return &ret.Cursor, nil
}

func (d DefaultLedger) AddMetadataToAccount(ctx context.Context, ledger, account string, metadata metadata.Metadata) error {
	//nolint:bodyclose
	_, err := d.client.AccountsApi.AddMetadataToAccount(ctx, ledger, account).RequestBody(metadata).Execute()
	return err
}

func (d DefaultLedger) GetAccount(ctx context.Context, ledger, account string) (*sdk.AccountWithVolumesAndBalances, error) {
	//nolint:bodyclose
	ret, _, err := d.client.AccountsApi.GetAccount(ctx, ledger, account).Execute()
	return &ret.Data, err
}

func (d DefaultLedger) ListAccounts(ctx context.Context, ledger string, query ListAccountsQuery) (*sdk.AccountsCursorResponseCursor, error) {
	var (
		ret *sdk.AccountsCursorResponse
		err error
	)
	if query.Cursor == "" {
		//nolint:bodyclose
		ret, _, err = d.client.AccountsApi.ListAccounts(ctx, ledger).
			Metadata(query.Metadata).
			PageSize(int64(query.Limit)).
			Execute()
	} else {
		//nolint:bodyclose
		ret, _, err = d.client.AccountsApi.ListAccounts(ctx, ledger).
			Cursor(query.Cursor).
			Execute()
	}
	if err != nil {
		return nil, err
	}

	return &ret.Cursor, nil
}

func (d DefaultLedger) CreateTransaction(ctx context.Context, ledger string, transaction shared.PostTransaction) error {
	//nolint:bodyclose
	tx := operations.CreateTransactionRequest{
		Ledger: ledger,
		PostTransaction: transaction,
	}
	_, err := d.client.Ledger.CreateTransaction(ctx, tx)
	return err
}

func (d DefaultLedger) RunScript(ctx context.Context, ledger string, script shared.Script) (*operations.RunScriptResponse, error) {
	//nolint:bodyclose
	request := operations.RunScriptRequest{
		Ledger: ledger,
		Script: script,
	}
	runScript, err := d.client.Ledger.RunScript(ctx, request)
	if err != nil {
		return nil, err
	}
	return runScript, err
}

var _ Ledger = &DefaultLedger{}

func NewDefaultLedger(client *sdk.Formance) *DefaultLedger {
	return &DefaultLedger{
		client: client,
	}
}
