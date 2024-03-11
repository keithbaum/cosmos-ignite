package keeper

func (k Keeper) calculateTxPositionInMerklePath(transactionIndex int64, merklePath [][]byte) (int64, error) {
	// ToDo
	return 0, nil
}

func (k Keeper) calculateMerkleRoot(fullMerklePath [][]byte) ([]byte, error) {
	// ToDo
	return []byte("0xdca3326ad7e8121bf9cf9c12333e6b2271abe823ec9edfe42f813b1e768fa57b"), nil
}
