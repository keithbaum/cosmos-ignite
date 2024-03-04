package types

import sdk "github.com/cosmos/cosmos-sdk/types"

var (
	_ sdk.Msg = &MsgStoreExternalChain1TxHeadersRequest{}
)

const (
	TypeMsgStoreExternalChain1TxHeaders = "storeExternalChain1TxHeaders"
)

func NewMsgStoreExternalChain1TxHeaders(merkleRoot string, blockHeight int64, sender string) *MsgStoreExternalChain1TxHeadersRequest {
	return &MsgStoreExternalChain1TxHeadersRequest{
		MerkleRoot:  merkleRoot,
		BlockHeight: blockHeight,
		Sender:      sender,
	}
}

func (msg *MsgStoreExternalChain1TxHeadersRequest) Route() string {
	return RouterKey
}

func (msg *MsgStoreExternalChain1TxHeadersRequest) Type() string {
	return TypeMsgStoreExternalChain1TxHeaders
}

func (msg *MsgStoreExternalChain1TxHeadersRequest) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgStoreExternalChain1TxHeadersRequest) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgStoreExternalChain1TxHeadersRequest) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return err
	}
	return nil
}
