package aasdk

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/genefriendway/aa-sdk-go/bindings/entrypoint"
)

const (
	PaymasterValidationGasOffset = 20
	PaymasterPostOpGasOffset     = 36
	PaymasterDataOffset          = 52
)

var (
	EmptySignature = make([]byte, 65)
)

// GetPaymasterHash returns the hash to sign for a user operation
func GetPaymasterHash(
	packedUserOp *entrypoint.PackedUserOperation,
	chainId *big.Int,
	validUntil *big.Int,
	validAfter *big.Int,
) (common.Hash, error) {
	if len(packedUserOp.PaymasterAndData) < PaymasterDataOffset {
		return common.Hash{}, fmt.Errorf("PaymasterAndData too short")
	}
	paymaster := common.BytesToAddress(packedUserOp.PaymasterAndData[:PaymasterValidationGasOffset])
	args := abi.Arguments{
		{Type: abi.Type{T: abi.AddressTy}},              //	sender
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      //	nonce
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, //	initCode
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, //	callData
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, //	accountGasLimits
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      //	paymasterValidationGas
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      //	preVerificationGas
		{Type: abi.Type{T: abi.FixedBytesTy, Size: 32}}, //	gasFees
		{Type: abi.Type{T: abi.UintTy, Size: 256}},      //	chainId
		{Type: abi.Type{T: abi.AddressTy}},              //	paymaster's address
		{Type: abi.Type{T: abi.UintTy, Size: 48}},       //	validUntil
		{Type: abi.Type{T: abi.UintTy, Size: 48}},       //	validAfter
	}

	packed, err := args.Pack(
		packedUserOp.Sender,
		packedUserOp.Nonce,
		crypto.Keccak256Hash(packedUserOp.InitCode),
		crypto.Keccak256Hash(packedUserOp.CallData),
		packedUserOp.AccountGasLimits,
		new(big.Int).SetBytes(packedUserOp.PaymasterAndData[PaymasterValidationGasOffset:PaymasterDataOffset]),
		packedUserOp.PreVerificationGas,
		packedUserOp.GasFees,
		chainId,
		paymaster,
		validUntil,
		validAfter,
	)
	if err != nil {
		return common.Hash{}, fmt.Errorf("pack error in GetPaymasterHash: %v", err)
	}

	return crypto.Keccak256Hash(packed), nil
}

// PackPaymasterAndData constructs paymasterAndData field
func PackPaymasterAndData(paymaster common.Address, verGasLimit, postOpGasLimit *big.Int, data []byte) []byte {
	// Convert gas limits to 16-byte padded slices
	verGasBytes := common.LeftPadBytes(verGasLimit.Bytes(), 16)
	postOpGasBytes := common.LeftPadBytes(postOpGasLimit.Bytes(), 16)

	length := len(paymaster) + len(verGasBytes) + len(postOpGasBytes) + len(data)
	result := make([]byte, 0, length)
	result = append(result, paymaster[:]...)   // 20 bytes
	result = append(result, verGasBytes...)    // 16 bytes
	result = append(result, postOpGasBytes...) // 16 bytes
	result = append(result, data...)           // variable length

	return result
}

// EncodePaymasterData encodes validUntil, validAfter, and signature into a byte array
func EncodePaymasterData(validUntil, validAfter *big.Int, signature []byte) ([]byte, error) {
	data, err := abi.Arguments{
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
		{Type: abi.Type{T: abi.UintTy, Size: 48}},
	}.Pack(
		validUntil,
		validAfter,
	)
	if err != nil {
		return nil, err
	}
	data = append(data, signature...)
	return data, nil
}
