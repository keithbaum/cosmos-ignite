package types

import "fmt"

const (
	// ModuleName defines the module name
	ModuleName = "lightclient"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_lightclient"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

var (
	// Keys for store prefixes
	ExternalChain1BlockHeightPrefix = KeyPrefix("ExternalChain1-BlockHeight-") // prefix for each blockheight in external chain 1
	ExternalChain2BlockHeightPrefix = KeyPrefix("ExternalChain2-BlockHeight-") // prefix for each blockheight in external chain 1
)

func GetExternalChain1BlockHeightPrefixKey(blockId int64) []byte {
	return append(ExternalChain1BlockHeightPrefix, []byte(fmt.Sprintf("%d", blockId))...)
}
