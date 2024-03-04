package keeper

import (
	"bytes"
	"foochain/x/lightclient/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// VerifyTx verifies a transaction against a header
func (k Keeper) VerifyTx(
	ctx sdk.Context,
	txHash []byte,
	blockHeight int64,
	proof types.Proof,
	data types.TxData,
) (bool, error) {
	store := k.getStore(ctx)
	merkleRoot := store.Get(types.GetExternalChain1BlockHeightPrefixKey(blockHeight))

	if merkleRoot == nil {
		return false, types.ErrInvalidBlockHeight
	}

	if !bytes.Equal(proof.MerkleRoot, merkleRoot) {
		return false, nil
	} // Do not consider it an error, just invalid tx

	return k.verifyMerkleProof(txHash, proof, data)
}

func (k Keeper) verifyMerkleProof(txHash []byte, proof types.Proof, data types.TxData) (bool, error) {
	// ToDo calculateTxPositionInMerklePath: It calculates in the binary merkle tree path, in which position should the txHash be located
	merklePathIndex, err := k.calculateTxPositionInMerklePath(data.TxIdx, proof.MerklePath)
	if err != nil {
		return false, err
	}

	var fullMerklePath [][]byte
	fullMerklePath = append(fullMerklePath, proof.MerklePath[:merklePathIndex]...)
	fullMerklePath = append(fullMerklePath, txHash)
	fullMerklePath = append(fullMerklePath, proof.MerklePath[merklePathIndex:]...)

	calculatedMerkleRoot, err := k.calculateMerkleRoot(fullMerklePath) // ToDo calculateMerkleRoot: Do the mamuschka of hashes calculation
	if err != nil {
		return false, err
	}
	return bytes.Equal(calculatedMerkleRoot, proof.MerkleRoot), nil
}

func (k Keeper) storeExternalChain1MerkleRoot(ctx sdk.Context, blockHeight int64, merkleRoot []byte) error {
	store := k.getStore(ctx)
	key := types.GetExternalChain1BlockHeightPrefixKey(blockHeight)
	if store.Has(key) {
		if !bytes.Equal(store.Get(key), merkleRoot) {
			return types.ErrInvalidMerkleRoot
		}
		return nil
	}
	store.Set(types.GetExternalChain1BlockHeightPrefixKey(blockHeight), merkleRoot)
	return nil
}
