package app

import (
	wasmkeeper "wasmd.mleku.dev/x/wasm/keeper"
)

// Deprecated: Use BuiltInCapabilities from github.com/CosmWasm/wasmd/x/wasm/keeper
func AllCapabilities() []string {
	return wasmkeeper.BuiltInCapabilities()
}
