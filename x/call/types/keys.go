package types

const (
	// ModuleName defines the module name
	ModuleName = "call"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_call"
)

var (
	ParamsKey = []byte("p_call")
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
