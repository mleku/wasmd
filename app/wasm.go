package app

import (
	wasmkeeper "wasmd.mleku.dev/x/wasm/keeper"
)

// Deprecated: Use BuiltInCapabilities from wasmd.mleku.dev/x/wasm/keeper
func AllCapabilities() []string {
	return wasmkeeper.BuiltInCapabilities()
}
