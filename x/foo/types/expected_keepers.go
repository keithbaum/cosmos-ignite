package types

import (
	lightclienttypes "foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

type LightclientKeeper interface {
	VerifyTx(
		ctx sdk.Context,
		txHash []byte,
		blockHeight int64,
		proof lightclienttypes.Proof,
		data lightclienttypes.TxData,
	) (bool, error)
}

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	// Methods imported from bank should be defined here
}
