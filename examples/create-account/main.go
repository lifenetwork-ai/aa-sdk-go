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
		VerifyingSigner:     verifyingSigner,
		ExecutorSigner:      executorSigner,
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

	signer, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate key: %v", err)
	}
	address := crypto.PubkeyToAddress(signer.PublicKey)
	slog.Info("Generated address", "address", address.Hex())

	// Save the private key to a file
	if err := crypto.SaveECDSA(privateKeyFile, signer); err != nil {
		log.Fatalf("Failed to save private key: %v", err)
	}

	// Create a new account
	account, err := client.GetAccount(context.Background(), address, big.NewInt(0))
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}
	slog.Info("Account created", "address", account.Hex())

}
