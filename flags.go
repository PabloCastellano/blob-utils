package main

import (
	"github.com/urfave/cli"
)

var (
	TxRPCURLFlag = cli.StringFlag{
		Name:  "rpc-url",
		Usage: "Address of exuection node JSON-RPC endpoint",
		Value: "http://127.0.0.1:8545",
	}
	TxBlobFileFlag = cli.StringFlag{
		Name:     "blob-file",
		Usage:    "Blob file data",
		Required: true,
	}
	TxToFlag = cli.StringFlag{
		Name:     "to",
		Usage:    "tx to address",
		Required: true,
	}
	TxValueFlag = cli.StringFlag{
		Name:  "value",
		Usage: "tx value (wei deonominated)",
		Value: "0x0",
	}
	TxPrivateKeyFlag = cli.StringFlag{
		Name:     "private-key",
		Usage:    "tx private key",
		Required: true,
	}
	TxNonceFlag = cli.Int64Flag{
		Name:  "nonce",
		Usage: "tx nonce",
		Value: -1,
	}
	TxGasLimitFlag = cli.Uint64Flag{
		Name:  "gas-limit",
		Usage: "tx gas limit",
		Value: 21000,
	}
	TxGasPriceFlag = cli.StringFlag{
		Name:  "gas-price",
		Usage: "sets the tx max_fee_per_gas",
	}
	TxPriorityGasPrice = cli.StringFlag{
		Name:  "priority-gas-price",
		Usage: "Sets the priority fee per gas",
		Value: "2000000000",
	}
	TxMaxFeePerBlobGas = cli.StringFlag{
		Name:  "max-fee-per-blob-gas",
		Usage: "Sets the max_fee_per_blob_gas",
		Value: "3000000000",
	}
	TxChainID = cli.StringFlag{
		Name:  "chain-id",
		Usage: "chain-id of the transaction",
		Value: "1332",
	}
	TxCalldata = cli.StringFlag{
		Name:  "calldata",
		Usage: "calldata of the transaction",
		Value: "0x",
	}

	// With 6 blobs per tx you can upload 768KB every 12 seconds
	MultiTxBlobsPerTx = cli.IntFlag{
		Name:  "blobs-per-tx",
		Usage: "Blobs per transaction",
		Value: 6,
	}

	DownloadBeaconRPCURLFlag = cli.StringFlag{
		Name:  "rpc-url",
		Usage: "Address of beacon node JSON-RPC endpoint",
		Value: "http://10.128.0.8:5052",
	}
	DownloadSlotFlag = cli.Int64Flag{
		Name:  "slot",
		Usage: "Slot to download blob from",
		Value: 125754,
	}

	ProofBlobFileFlag = cli.StringFlag{
		Name:     "blob-file",
		Usage:    "Blob file data",
		Required: true,
	}
	ProofBlobIndexFlag = cli.StringFlag{
		Name:     "blob-index",
		Usage:    "Blob index",
		Required: true,
	}
	ProofInputPointFlag = cli.StringFlag{
		Name:     "input-point",
		Usage:    "Input point of the proof",
		Required: true,
	}
)

var TxFlags = []cli.Flag{
	TxRPCURLFlag,
	TxBlobFileFlag,
	TxToFlag,
	TxValueFlag,
	TxPrivateKeyFlag,
	TxNonceFlag,
	TxGasLimitFlag,
	TxGasPriceFlag,
	TxPriorityGasPrice,
	TxMaxFeePerBlobGas,
	TxChainID,
	TxCalldata,
	MultiTxBlobsPerTx,
}

var DownloadFlags = []cli.Flag{
	DownloadBeaconRPCURLFlag,
	DownloadSlotFlag,
}

var WebserverFlags = []cli.Flag{
	TxRPCURLFlag,
	TxPrivateKeyFlag,
}

var ProofFlags = []cli.Flag{
	ProofBlobFileFlag,
	ProofBlobIndexFlag,
	ProofInputPointFlag,
}
