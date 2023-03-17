package ledger

import (
	"context"
	"fmt"

	"github.com/formancehq/ledger/pkg/cache"
	"github.com/formancehq/ledger/pkg/core"
	"github.com/formancehq/ledger/pkg/storage"
	"github.com/formancehq/stack/libs/go-libs/api"
	"github.com/pkg/errors"
)

type Ledger struct {
	store               storage.LedgerStore
	allowPastTimestamps bool
	locker              Locker
	dbCache             *cache.Ledger
}

type Option = func(*Ledger)

func WithPastTimestamps(l *Ledger) {
	l.allowPastTimestamps = true
}

func WithLocker(locker Locker) Option {
	return func(ledger *Ledger) {
		ledger.locker = locker
	}
}

var defaultLedgerOptions = []Option{
	WithLocker(NewInMemoryLocker()),
}

func NewLedger(store storage.LedgerStore, dbCache *cache.Ledger, options ...Option) (*Ledger, error) {
	l := &Ledger{
		store:   store,
		dbCache: dbCache,
	}

	for _, option := range append(
		defaultLedgerOptions,
		options...,
	) {
		option(l)
	}

	return l, nil
}

func (l *Ledger) GetDBCache() *cache.Ledger {
	return l.dbCache
}

func (l *Ledger) Close(ctx context.Context) error {
	if err := l.store.Close(ctx); err != nil {
		return errors.Wrap(err, "closing store")
	}
	return nil
}

func (l *Ledger) GetLedgerStore() storage.LedgerStore {
	return l.store
}

func (l *Ledger) GetTransactions(ctx context.Context, q storage.TransactionsQuery) (api.Cursor[core.ExpandedTransaction], error) {
	return l.store.GetTransactions(ctx, q)
}

func (l *Ledger) CountTransactions(ctx context.Context, q storage.TransactionsQuery) (uint64, error) {
	return l.store.CountTransactions(ctx, q)
}

func (l *Ledger) GetTransaction(ctx context.Context, id uint64) (*core.ExpandedTransaction, error) {
	tx, err := l.store.GetTransaction(ctx, id)
	if err != nil {
		return nil, err
	}
	if tx == nil {
		return nil, NewNotFoundError("transaction not found")
	}

	return tx, nil
}

func (l *Ledger) RevertTransaction(ctx context.Context, id uint64) (*core.ExpandedTransaction, *LogHandler, error) {

	revertedTx, err := l.store.GetTransaction(ctx, id)
	if err != nil {
		return nil, nil, errors.Wrap(err, fmt.Sprintf("getting transaction %d", id))
	}
	if revertedTx == nil {
		return nil, nil, NewNotFoundError(fmt.Sprintf("transaction %d not found", id))
	}
	if revertedTx.IsReverted() {
		return nil, nil, NewValidationError(fmt.Sprintf("transaction %d already reverted", id))
	}

	rt := revertedTx.Reverse()
	rt.Metadata = core.Metadata{}
	rt.Metadata.MarkReverts(revertedTx.ID)

	scriptData := core.TxsToScriptsData(core.TransactionData{
		Postings:  rt.Postings,
		Reference: rt.Reference,
		Metadata:  rt.Metadata,
	})

	return l.runScript(ctx, scriptData[0], false, func(expandedTx core.ExpandedTransaction, accountMetadata map[string]core.Metadata) core.Log {
		return core.NewRevertedTransactionLog(expandedTx.Timestamp, revertedTx.ID, expandedTx.Transaction)
	})
}

func (l *Ledger) CountAccounts(ctx context.Context, a storage.AccountsQuery) (uint64, error) {
	return l.store.CountAccounts(ctx, a)
}

func (l *Ledger) GetAccounts(ctx context.Context, a storage.AccountsQuery) (api.Cursor[core.Account], error) {
	return l.store.GetAccounts(ctx, a)
}

func (l *Ledger) GetAccount(ctx context.Context, address string) (*core.AccountWithVolumes, error) {
	return l.store.GetAccountWithVolumes(ctx, address)
}

func (l *Ledger) GetBalances(ctx context.Context, q storage.BalancesQuery) (api.Cursor[core.AccountsBalances], error) {
	return l.store.GetBalances(ctx, q)
}

func (l *Ledger) GetBalancesAggregated(ctx context.Context, q storage.BalancesQuery) (core.AssetsBalances, error) {
	return l.store.GetBalancesAggregated(ctx, q)
}

// TODO(gfyrag): maybe we should check transaction exists before set a metadata ? (accounts always exists even if never used)
func (l *Ledger) SaveMeta(ctx context.Context, targetType string, targetID interface{}, m core.Metadata) (*LogHandler, error) {

	if targetType == "" {
		return nil, NewValidationError("empty target type")
	}

	if targetID == "" {
		return nil, NewValidationError("empty target id")
	}

	// TODO(gfyrag): Stop round dates for v2
	at := core.Now()

	var (
		err error
		log core.Log
	)
	switch targetType {
	case core.MetaTargetTypeTransaction:
		log = core.NewSetMetadataLog(at, core.SetMetadataLogPayload{
			TargetType: core.MetaTargetTypeTransaction,
			TargetID:   targetID.(uint64),
			Metadata:   m,
		})
	case core.MetaTargetTypeAccount:
		// Machine can access account metadata, so store the metadata until CQRS compute final of the account
		// The cache can still evict the account entry before CQRS part compute the view
		err = l.dbCache.UpdateAccountMetadata(ctx, targetID.(string), m)
		if err != nil {
			return nil, err
		}

		log = core.NewSetMetadataLog(at, core.SetMetadataLogPayload{
			TargetType: core.MetaTargetTypeAccount,
			TargetID:   targetID.(string),
			Metadata:   m,
		})
	default:
		return nil, NewValidationError(fmt.Sprintf("unknown target type '%s'", targetType))
	}
	if err != nil {
		return nil, err
	}

	return writeLog(ctx, l.store.AppendLogs, log)
}

func (l *Ledger) GetLogs(ctx context.Context, q *storage.LogsQuery) (api.Cursor[core.Log], error) {
	return l.store.GetLogs(ctx, q)
}
