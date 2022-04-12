package arbitrum

import (
	"time"

	"github.com/ethereum/go-ethereum/eth/ethconfig"
	"github.com/ethereum/go-ethereum/params"
	flag "github.com/spf13/pflag"
)

type Config struct {
	// RPCGasCap is the global gas cap for eth-call variants.
	RPCGasCap uint64 `koanf:"gas-cap"`

	// RPCTxFeeCap is the global transaction fee(price * gaslimit) cap for
	// send-transction variants. The unit is ether.
	RPCTxFeeCap float64 `koanf:"tx-fee-cap"`

	// RPCEVMTimeout is the global timeout for eth-call.
	RPCEVMTimeout time.Duration `koanf:"evm-timeout"`

	// Parameters for the bloom indexer
	BloomBitsBlocks uint64 `koanf:"bloom-bits-blocks"`
	BloomConfirms   uint64 `koanf:"bloom-confirms"`
}

func ConfigAddOptions(prefix string, f *flag.FlagSet) {
	f.Uint64(prefix+".gas-cap", DefaultConfig.RPCGasCap, "cap on computation gas that can be used in eth_call/estimateGas (0=infinite)")
	f.Float64(prefix+".tx-fee-cap", DefaultConfig.RPCTxFeeCap, "cap on transaction fee (in ether) that can be sent via the RPC APIs (0 = no cap)")
	f.Duration(prefix+".evm-timeout", DefaultConfig.RPCEVMTimeout, "timeout used for eth_call (0=infinite)")
	f.Uint64(prefix+".bloom-bits-blocks", DefaultConfig.BloomBitsBlocks, "number of blocks a single bloom bit section vector holds")
	f.Uint64(prefix+".bloom-confirms", DefaultConfig.BloomConfirms, "number of confirmations before indexing new bloom filters")
}

var DefaultConfig = Config{
	RPCGasCap:       ethconfig.Defaults.RPCGasCap,     // 50,000,000
	RPCTxFeeCap:     ethconfig.Defaults.RPCTxFeeCap,   // 1 ether
	RPCEVMTimeout:   ethconfig.Defaults.RPCEVMTimeout, // 5 seconds
	BloomBitsBlocks: params.BloomBitsBlocks * 4,       // we generally have smaller blocks
	BloomConfirms:   params.BloomConfirms,
}
