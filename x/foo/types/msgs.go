package types

import (
	lightclientmoduletypes "foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ sdk.Msg = &MsgVerifyTxRequest{}
)

const (
	TypeMsgVerifyTxRequest = "verifyTxRequest"
)

func NewMsgVerifyTxRequest(txHash string, block_height int64, proof lightclientmoduletypes.Proof, txData lightclientmoduletypes.TxData) *MsgVerifyTxRequest {
	return &MsgVerifyTxRequest{
		TxHash:      txHash,
		BlockHeight: block_height,
		Proof:       &proof,
		TxData:      &txData,
	}
}

func (msg *MsgVerifyTxRequest) Route() string {
	return RouterKey
}

func (msg *MsgVerifyTxRequest) Type() string {
	return TypeMsgVerifyTxRequest
}

func (msg *MsgVerifyTxRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgVerifyTxRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgVerifyTxRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}
