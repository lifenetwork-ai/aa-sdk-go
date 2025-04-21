package main

import (
	"context"
	"log"
	"log/slog"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	aasdk "github.com/genefriendway/aa-sdk-go"
)

const (
	privateKeyFile = "./examples/private_key.pem"
)

func main() {
	// Example usage of the functions defined in the aasdk package
	// This is just a placeholder to show where the main function would be
	// In a real application, you would call the functions defined in the aasdk package here
	// and handle any errors or results as needed.

	// Load private keys from .env file or environment variables
	verifyingKey := os.Getenv("PAYMASTER_VERIFYING_KEY")
	if verifyingKey == "" {
		log.Fatal("PAYMASTER_VERIFYING_KEY not set")
	}
	verifyingSigner, err := crypto.HexToECDSA(verifyingKey)
	if err != nil {
		log.Fatalf("Failed to parse verifying key: %v", err)
	}
	executorKey := os.Getenv("EXECUTOR_KEY")
	if executorKey == "" {
		log.Fatal("EXECUTOR_KEY not set")
	}
	executorSigner, err := crypto.HexToECDSA(executorKey)
	if err != nil {
		log.Fatalf("Failed to parse executor key: %v", err)
	}

	paymasterAddress := common.HexToAddress("0xe7db0C105Ac75A493B0413046417e48594360542")
	config := &aasdk.Config{
		NodeUrl:             "https://testnet-lifeaiv1-c648f.avax-test.network/ext/bc/62fkxYTWbGBfXoHNXcGJbq2dTXba2uoCFySzdHy87iovJj2F4/rpc?token=25e957a027b09bb006da7e9fc981100ce25f333cd998a76eb36a842fcb5ba63a",
		BundlerUrl:          "http://34.126.118.65:8080/rpc",
		WaitReceiptInterval: 2 * time.Second,
		Entrypoint:          common.HexToAddress("0xd308aE59cb31932E8D9305BAda32Fa782d3D5d42"),
		AccountFactory:      common.HexToAddress("0xD421D8470b577f6A64992132D04906EfE51F1dE3"),
		PaymasterAddress:    &paymasterAddress,
		VerifyingSigner:     verifyingSigner, // optional, only used for verifying paymaster data
		ExecutorSigner:      executorSigner,  // optional, only used for sending atomic transactions
	}

	client, err := aasdk.NewClient(config, aasdk.NewLRUCache(10000))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	entrypoints, err := client.SupportedEntryPoints(context.Background())
	if err != nil {
		log.Fatalf("Failed to get supported entry points: %v", err)
	}
	slog.Info("Supported entry points", "entrypoints", entrypoints)

	signer, err := crypto.LoadECDSA(privateKeyFile)
	if err != nil {
		log.Fatalf("Failed to load signer: %v", err)
	}

	salt := big.NewInt(0)

	// Example of creating a user operation
	sender, err := client.GetAccount(context.Background(), crypto.PubkeyToAddress(signer.PublicKey), salt)
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}

	target := common.HexToAddress("0x6B7A05ED423D339743ae6d2090f5AFC148344566")
	amount := big.NewInt(1e18)

	balance, err := client.GetAccountBalance(context.Background(), sender)
	if err != nil {
		log.Fatalf("Failed to get balance: %v", err)
	}
	slog.Info("Account balance", "balance", balance)
	if balance.Cmp(amount) < 0 {
		log.Fatalf("Insufficient balance: %s < %s", balance.String(), amount.String())
	}

	calldata, err := aasdk.PackTransferData(client.AccountABI(), target, amount)
	if err != nil {
		log.Fatalf("Failed to pack transfer data: %v", err)
	}
	userOp := aasdk.NewUserOpWithDefault(sender, calldata, salt)

	tx, err := client.SendUserOp(context.Background(), userOp, signer)
	if err != nil {
		log.Fatalf("Failed to send user operation: %v", err)
	}

	slog.Info("User operation sent", "tx", tx.Hex())

	// Wait for the transaction to be mined
	receipt, err := client.WaitForUserOperation(context.Background(), tx)
	if err != nil {
		log.Fatalf("Failed to wait for user operation: %v", err)
	}
	if receipt == nil {
		log.Fatalf("User operation receipt is nil")
	}
	slog.Info("User operation receipt", "receipt", receipt)
	slog.Info("User operation status", "success", receipt.Success)
	slog.Info("Actual gas cost", "gas", receipt.ActualGasCost)
	slog.Info("Actual gas used", "gasUsed", receipt.ActualGasUsed)
}
