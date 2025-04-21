package aasdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	DefaultCallGasLimit                  = int64(2000000)
	DefaultVerificationGasLimit          = int64(200000)
	DefaultPreVerificationGas            = int64(20000)
	DefaultMaxFeePerGas                  = int64(25e9)
	DefaultMaxPriorityFeePerGas          = int64(1000000)
	DefaultPaymasterVerificationGasLimit = int64(3e5)
	DefaultPaymasterPostOpGasLimit       = int64(100)
)

type Config struct {
	// The url of node.
	NodeUrl string
	// The url of bundler.
	BundlerUrl string
	// The interval to query the receipt.
	WaitReceiptInterval time.Duration
	// The entrypoint address.
	// Currently, it supports Entrypoint V0.7.0
	Entrypoint common.Address
	// The simple account factory address.
	AccountFactory common.Address
	// The verifying paymaster address.
	PaymasterAddress *common.Address
	// The account verifying Paymaster requests.
	VerifyingSigner *ecdsa.PrivateKey
	// The account that will sign the user operation.
	// It's needed when call directly to Entrypoint contract.
	ExecutorSigner *ecdsa.PrivateKey
}

func NewUserOpWithDefault(sender common.Address, calldata []byte, salt *big.Int) *UserOperation {
	return &UserOperation{
		Sender:                        sender,
		CallData:                      calldata,
		CallGasLimit:                  big.NewInt(DefaultCallGasLimit),
		VerificationGasLimit:          big.NewInt(DefaultVerificationGasLimit),
		PreVerificationGas:            big.NewInt(DefaultPreVerificationGas),
		MaxFeePerGas:                  big.NewInt(DefaultMaxFeePerGas),
		MaxPriorityFeePerGas:          big.NewInt(DefaultMaxPriorityFeePerGas),
		PaymasterVerificationGasLimit: big.NewInt(DefaultPaymasterVerificationGasLimit),
		PaymasterPostOpGasLimit:       big.NewInt(DefaultPaymasterPostOpGasLimit),
		Salt:                          salt,
	}
}

// UserOperation represents the base structure for operations by ERC-4337
// Supported EntryPoint V0.7.0
type UserOperation struct {
	Sender                        common.Address `json:"sender"`
	Nonce                         *big.Int       `json:"nonce"`
	CallData                      []byte         `json:"callData"`
	CallGasLimit                  *big.Int       `json:"callGasLimit"`
	VerificationGasLimit          *big.Int       `json:"verificationGasLimit"`
	PreVerificationGas            *big.Int       `json:"preVerificationGas"`
	MaxFeePerGas                  *big.Int       `json:"maxFeePerGas"`
	MaxPriorityFeePerGas          *big.Int       `json:"maxPriorityFeePerGas"`
	Signature                     []byte         `json:"signature"`
	Paymaster                     common.Address `json:"paymaster"`
	PaymasterData                 []byte         `json:"paymasterData"`
	PaymasterVerificationGasLimit *big.Int       `json:"paymasterVerificationGasLimit"`
	PaymasterPostOpGasLimit       *big.Int       `json:"paymasterPostOpGasLimit"`
	Factory                       common.Address `json:"factory"`
	FactoryData                   []byte         `json:"factoryData"`
	InitCode                      []byte         `json:"initCode"`
	Salt                          *big.Int
}

// ToBody converts the UserOperation to a map of strings.
// It helps to perform json request.
func (u *UserOperation) ToBody() map[string]string {
	body := make(map[string]string)
	if u.Sender != (common.Address{}) {
		body["sender"] = u.Sender.Hex()
	}
	if u.Nonce != nil {
		body["nonce"] = "0x" + u.Nonce.Text(16)
	}
	if len(u.CallData) > 0 {
		body["callData"] = "0x" + hex.EncodeToString(u.CallData)
	}
	if u.CallGasLimit != nil {
		body["callGasLimit"] = "0x" + u.CallGasLimit.Text(16)
	}
	if u.VerificationGasLimit != nil {
		body["verificationGasLimit"] = "0x" + u.VerificationGasLimit.Text(16)
	}
	if u.PreVerificationGas != nil {
		body["preVerificationGas"] = "0x" + u.PreVerificationGas.Text(16)
	}
	if u.MaxFeePerGas != nil {
		body["maxFeePerGas"] = "0x" + u.MaxFeePerGas.Text(16)
	}
	if u.MaxPriorityFeePerGas != nil {
		body["maxPriorityFeePerGas"] = "0x" + u.MaxPriorityFeePerGas.Text(16)
	}
	if len(u.Signature) > 0 {
		body["signature"] = "0x" + hex.EncodeToString(u.Signature)
	}
	if u.Paymaster != (common.Address{}) {
		body["paymaster"] = u.Paymaster.Hex()
	}
	if len(u.PaymasterData) > 0 {
		body["paymasterData"] = "0x" + hex.EncodeToString(u.PaymasterData)
	}
	if u.PaymasterVerificationGasLimit != nil {
		body["paymasterVerificationGasLimit"] = "0x" + u.PaymasterVerificationGasLimit.Text(16)
	}
	if u.PaymasterPostOpGasLimit != nil {
		body["paymasterPostOpGasLimit"] = "0x" + u.PaymasterPostOpGasLimit.Text(16)
	}
	if u.Factory != (common.Address{}) {
		body["factory"] = u.Factory.Hex()
	}
	if len(u.FactoryData) > 0 {
		body["factoryData"] = "0x" + hex.EncodeToString(u.FactoryData)
	}
	return body
}

// TxDetail represents the details of a transaction for a user operation
type TxDetail struct {
	Target               common.Address // The target address of the transaction
	Data                 []byte         // The data to be sent with the transaction
	Value                *big.Int       // The value to be transferred (optional)
	GasLimit             *big.Int       // The gas limit for the transaction (optional)
	MaxFeePerGas         *big.Int       // The maximum fee per gas unit (optional)
	MaxPriorityFeePerGas *big.Int       // The maximum priority fee per gas unit (optional)
	Nonce                *big.Int       // The nonce for the transaction (optional)
}

type receipt struct {
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       string         `json:"blockNumber"`
	From              common.Address `json:"from"`
	CumulativeGasUsed string         `json:"cumulativeGasUsed"`
	GasUsed           string         `json:"gasUsed"`
	Logs              []*types.Log   `json:"logs"`
	LogsBloom         types.Bloom    `json:"logsBloom"`
	TransactionHash   common.Hash    `json:"transactionHash"`
	TransactionIndex  string         `json:"transactionIndex"`
	EffectiveGasPrice string         `json:"effectiveGasPrice"`
}

type UserOpReceipt struct {
	UserOpHash    common.Hash    `json:"userOpHash"`
	Sender        common.Address `json:"sender"`
	Paymaster     common.Address `json:"paymaster"`
	Nonce         string         `json:"nonce"`
	Success       bool           `json:"success"`
	ActualGasCost string         `json:"actualGasCost"`
	ActualGasUsed string         `json:"actualGasUsed"`
	From          common.Address `json:"from"`
	Receipt       *receipt       `json:"receipt"`
	Logs          []*types.Log   `json:"logs"`
	ReturnData    []byte         `json:"returnData"`
}

// GasEstimates provides estimate values for all gas fields in a UserOperation.
type GasEstimates struct {
	PreVerificationGas   *big.Int `json:"preVerificationGas"`
	VerificationGasLimit *big.Int `json:"verificationGasLimit"`
	CallGasLimit         *big.Int `json:"callGasLimit"`
	VerificationGas      *big.Int `json:"verificationGas"`
	MaxFeePerGas         *big.Int `json:"maxFeePerGas"`
	MaxPriorityFeePerGas *big.Int `json:"maxPriorityFeePerGas"`
}

type BaseAccount interface {
	// GetAccount gets the account address for the given owner.
	GetAccount(ctx context.Context, owner common.Address, salt *big.Int) (common.Address, error)
}

type Bundler interface {
	// SendUserOp sends the user operation to the bundler.
	SendUserOp(ctx context.Context, userOp *UserOperation, signer *ecdsa.PrivateKey) (common.Hash, error)

	// EstimateUserOpGas estimates the gas needed for the user operation.
	EstimateUserOpGas(ctx context.Context, userOp *UserOperation) (*GasEstimates, error)

	// GetUserOpReceipt returns the receipt of the user operation.
	GetUserOpReceipt(ctx context.Context, userOpHash common.Hash) (*UserOpReceipt, error)

	// SupportedEntryPoints returns the supported entry points for the bundler.
	SupportedEntryPoints(ctx context.Context) ([]common.Address, error)
}
