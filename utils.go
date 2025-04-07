package aasdk

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/genefriendway/aa-sdk-go/bindings/entrypoint"
)

// IsAccountDeployed checks if the account is deployed by querying its bytecode
func IsAccountDeployed(ctx context.Context, client *ethclient.Client, address common.Address) (bool, error) {
	code, err := client.CodeAt(ctx, address, nil)
	if err != nil {
		return false, err
	}
	return len(code) > 0, nil
}

func GetUserOpHash(packed *entrypoint.PackedUserOperation, entrypoint common.Address, chainId *big.Int) (common.Hash, error) {
	hashed, err := HashedUserOp(packed)
	if err != nil {
		return common.Hash{}, err
	}
	hashArgs := abi.Arguments{
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // userOp.hash
		{Type: abi.Type{T: abi.AddressTy}},              // entrypoint address
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      // chainID
	}
	packedHash, err := hashArgs.Pack(hashed, entrypoint, chainId)
	if err != nil {
		return common.Hash{}, err
	}
	// Compute final Keccak-256 hash
	return crypto.Keccak256Hash(packedHash), nil
}

func HashedUserOp(userOp *entrypoint.PackedUserOperation) (common.Hash, error) {
	arguments := abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},              // sender
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      // nonce
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // hashInitCode
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // hashCallData
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // accountGasLimits
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      // preVerificationGas
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // gasFees
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, // hashPaymasterAndData
	}

	// Compute hashes for dynamic fields
	hashInitCode := crypto.Keccak256Hash(userOp.InitCode)
	hashCallData := crypto.Keccak256Hash(userOp.CallData)
	hashPaymasterAndData := crypto.Keccak256Hash(userOp.PaymasterAndData)

	packed, err := arguments.Pack(
		userOp.Sender,
		userOp.Nonce,
		hashInitCode,
		hashCallData,
		userOp.AccountGasLimits,
		userOp.PreVerificationGas,
		userOp.GasFees,
		hashPaymasterAndData,
	)
	if err != nil {
		return common.Hash{}, err
	}
	return crypto.Keccak256Hash(packed), nil
}

// HexToBigInt converts a hex string to a big.Int.
// If the hex string is prefixed with "0x", it will be removed.
func HexToBigInt(hex string) *big.Int {
	if strings.HasPrefix(hex, "0x") {
		return new(big.Int).SetBytes(common.Hex2Bytes(hex[2:]))
	} else {
		return new(big.Int).SetBytes(common.Hex2Bytes(hex))
	}
}

// PackTransferData packs the transfer data for a user operation.
func PackTransferData(accountABI *abi.ABI, target common.Address, value *big.Int) ([]byte, error) {
	packed, err := accountABI.Pack("execute", target, value, []byte{})
	if err != nil {
		return nil, err
	}
	return packed, nil
}

// PackTransferData packs the transfer data for a user operation.
func PackBatchTransferData(accountABI *abi.ABI, dest []common.Address, value []*big.Int) ([]byte, error) {
	if len(dest) != len(value) {
		return nil, fmt.Errorf("dest and value length mismatch: %d != %d", len(dest), len(value))
	}
	data := make([][]byte, len(dest))
	packed, err := accountABI.Pack("executeBatch", dest, value, data)
	if err != nil {
		return nil, err
	}
	return packed, nil
}

// PackInt packs two big.Ints into a common.Hash.
// The first 16 bytes are the first big.Int and the last 16 bytes are the second big.Int.
// If one of the big.Ints is nil, it will be padded with zeros.
// It panics if both big.Ints are nil.
func PackInt(a *big.Int, b *big.Int) common.Hash {
	if a == nil || b == nil {
		panic("nil data")
	}
	var result common.Hash
	copy(result[:], append(
		common.LeftPadBytes(a.Bytes(), 16),
		common.LeftPadBytes(b.Bytes(), 16)...,
	))
	return result
}

// PackUserOperation packs a user operation into a PackedUserOperation.
// It panics if the user operation is nil.
func PackUserOperation(userOp *UserOperation) entrypoint.PackedUserOperation {
	if userOp == nil {
		panic("nil user operation")
	}
	return entrypoint.PackedUserOperation{
		Sender:             userOp.Sender,
		Nonce:              userOp.Nonce,
		CallData:           userOp.CallData,
		AccountGasLimits:   PackInt(userOp.VerificationGasLimit, userOp.CallGasLimit),
		PreVerificationGas: userOp.PreVerificationGas,
		GasFees:            PackInt(userOp.MaxPriorityFeePerGas, userOp.MaxFeePerGas),
		PaymasterAndData:   PackPaymasterAndData(userOp.Paymaster, userOp.PaymasterVerificationGasLimit, userOp.PaymasterPostOpGasLimit, userOp.PaymasterData),
		Signature:          userOp.Signature,
		InitCode:           userOp.InitCode,
	}
}
