package aasdk

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math"
	"math/big"
	"net/http"
	"sync/atomic"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/genefriendway/aa-sdk-go/bindings/account"
	"github.com/genefriendway/aa-sdk-go/bindings/entrypoint"
)

var _ Bundler = &Client{}
var _ BaseAccount = &Client{}

type Client struct {
	id               atomic.Uint64 // unique id for the client
	chainId          *big.Int
	config           *Config
	eth              *ethclient.Client
	http             *http.Client
	simpleFactory    *account.SimpleAccountFactory
	entrypoint       *entrypoint.EntryPoint
	simpleAccountABI *abi.ABI
	simpleFactoryABI *abi.ABI
	lruCache         LRUCache
}

// NewClient creates a new Client instance with given config.
func NewClient(config *Config, cache LRUCache) (*Client, error) {
	eth, err := ethclient.Dial(config.NodeUrl)
	if err != nil {
		return nil, fmt.Errorf("error creating eth client: %v", err)
	}
	chainId, err := eth.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting chain id: %v", err)
	}
	entrypoint, err := entrypoint.NewEntryPoint(config.Entrypoint, eth)
	if err != nil {
		return nil, fmt.Errorf("error creating entrypoint client: %v", err)
	}
	simpleFactory, err := account.NewSimpleAccountFactory(config.AccountFactory, eth)
	if err != nil {
		return nil, fmt.Errorf("error creating account factory client: %v", err)
	}
	simpleFactoryABI, err := account.SimpleAccountFactoryMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("error getting account factory ABI: %v", err)
	}
	simpleAccountABI, err := account.SimpleAccountMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("error getting simple account ABI: %v", err)
	}

	c := &Client{
		id:               atomic.Uint64{},
		chainId:          chainId,
		eth:              eth,
		http:             http.DefaultClient,
		config:           config,
		lruCache:         cache,
		entrypoint:       entrypoint,
		simpleFactory:    simpleFactory,
		simpleAccountABI: simpleAccountABI,
		simpleFactoryABI: simpleFactoryABI,
	}
	return c, nil
}

// GetAccount returns the smart account address for the given owner and salt.
func (c *Client) GetAccount(ctx context.Context, owner common.Address, salt *big.Int) (common.Address, error) {
	if c.lruCache == nil {
		return c.simpleFactory.GetAddress(&bind.CallOpts{}, owner, salt)
	}
	key := fmt.Sprintf("%s-%s", owner.Hex(), salt.String())
	if addr, ok := c.lruCache.Get(key); ok {
		return addr.(common.Address), nil
	}
	addr, err := c.simpleFactory.GetAddress(&bind.CallOpts{}, owner, salt)
	if err != nil {
		return common.Address{}, fmt.Errorf("error getting account address: %v", err)
	}
	c.lruCache.Set(key, addr)
	return addr, nil
}

// GetAccountBalance returns the balance of the given account.
func (c *Client) GetAccountBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	balance, err := c.eth.BalanceAt(ctx, account, nil)
	if err != nil {
		return nil, fmt.Errorf("error getting account balance: %v", err)
	}
	return balance, nil
}

// FillAndSign fills the user operation with default values and signs it.
func (c *Client) FillAndSign(ctx context.Context, userOp *UserOperation, signer *ecdsa.PrivateKey) (*UserOperation, common.Hash, error) {
	if userOp.Sender == (common.Address{}) {
		return nil, common.Hash{}, fmt.Errorf("sender address is empty")
	}
	if userOp.Nonce == nil {
		nonce, err := c.entrypoint.GetNonce(&bind.CallOpts{}, userOp.Sender, userOp.Salt)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("error getting nonce: %v", err)
		}
		userOp.Nonce = nonce
	}

	initCode, data, err := c.getInitCodeData(ctx, userOp.Sender, crypto.PubkeyToAddress(signer.PublicKey), userOp.Salt)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("error getting account init code: %v", err)
	}

	if len(initCode) != 0 {
		userOp.InitCode = initCode
		userOp.Factory = c.config.AccountFactory
		userOp.FactoryData = data
	} else {
		userOp.InitCode = []byte{}
	}

	if c.config.PaymasterAddress != nil {
		// Using paymaster default validation time
		validAfter := big.NewInt(0)
		validUntil := big.NewInt(math.MaxInt32)

		paymasterData, err := EncodePaymasterData(validUntil, validAfter, EmptySignature)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("error encoding paymaster data: %v", err)
		}

		userOp.Paymaster = *c.config.PaymasterAddress

		paymasterHash, err := GetPaymasterHash(&entrypoint.PackedUserOperation{
			Sender:             userOp.Sender,
			Nonce:              userOp.Nonce,
			InitCode:           userOp.InitCode,
			CallData:           userOp.CallData,
			AccountGasLimits:   PackInt(userOp.VerificationGasLimit, userOp.CallGasLimit),
			PreVerificationGas: userOp.PreVerificationGas,
			GasFees:            PackInt(userOp.MaxPriorityFeePerGas, userOp.MaxFeePerGas),
			PaymasterAndData:   PackPaymasterAndData(userOp.Paymaster, userOp.PaymasterVerificationGasLimit, userOp.PaymasterPostOpGasLimit, paymasterData),
			Signature:          []byte{},
		}, c.chainId, validUntil, validAfter)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("error getting paymaster data: %v", err)
		}
		paymasterSig, err := SignMessage(c.config.VerifyingSigner, paymasterHash.Bytes())
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("error signing paymaster data: %v", err)
		}
		paymasterData, err = EncodePaymasterData(validUntil, validAfter, paymasterSig)
		if err != nil {
			return nil, common.Hash{}, fmt.Errorf("error encoding paymaster data: %v", err)
		}
		userOp.PaymasterData = paymasterData
	}

	packed := PackUserOperation(userOp)

	sig, hash, err := c.SignUserOp(&packed, signer)
	if err != nil {
		return nil, common.Hash{}, fmt.Errorf("error signing user operation: %v", err)
	}
	userOp.Signature = sig

	return userOp, hash, nil
}

// SignUserOp signs a user operation using the provided private key.
func (c *Client) SignUserOp(packed *entrypoint.PackedUserOperation, privateKey *ecdsa.PrivateKey) ([]byte, common.Hash, error) {
	hash, err := GetUserOpHash(packed, c.config.Entrypoint, c.chainId)
	if err != nil {
		return nil, common.Hash{}, err
	}
	sig, err := SignMessage(privateKey, hash.Bytes())
	if err != nil {
		return nil, common.Hash{}, err
	}
	return sig, hash, nil
}

func (c *Client) getInitCodeData(ctx context.Context, account common.Address, owner common.Address, salt *big.Int) ([]byte, []byte, error) {
	isDeployed, err := IsAccountDeployed(ctx, c.eth, account)
	if err != nil {
		return nil, nil, fmt.Errorf("error checking if account is deployed: %v", err)
	}
	if isDeployed {
		return nil, nil, nil
	}

	data, err := c.simpleFactoryABI.Pack("createAccount", owner, salt)
	if err != nil {
		return nil, nil, fmt.Errorf("error packing account init code: %v", err)
	}
	initCode := append(c.config.AccountFactory.Bytes(), data...)
	return initCode, data, nil
}

// HandleOps handles the user operations by calling the entrypoint contract directly.
func (c *Client) HandleOps(ctx context.Context, ops []entrypoint.PackedUserOperation) ([]common.Hash, common.Hash, error) {
	if c.config.ExecutorSigner == nil {
		panic("executor signer is nil")
	}

	txOpts, err := bind.NewKeyedTransactorWithChainID(c.config.ExecutorSigner, c.chainId)
	if err != nil {
		return []common.Hash{}, common.Hash{}, fmt.Errorf("error creating transaction options: %v", err)
	}
	tx, err := c.entrypoint.HandleOps(txOpts, ops, crypto.PubkeyToAddress(c.config.ExecutorSigner.PublicKey))
	if err != nil {
		return []common.Hash{}, common.Hash{}, fmt.Errorf("error handling ops: %v", err)
	}
	var opHashes []common.Hash
	for _, op := range ops {
		hashed, err := HashedUserOp(&op)
		if err != nil {
			return []common.Hash{}, common.Hash{}, fmt.Errorf("error hashing user operation: %v", err)
		}
		opHashes = append(opHashes, hashed)
	}
	return opHashes, tx.Hash(), nil
}

// HandleAtomicOps handles the user operations with atomic mode by calling the entrypoint contract directly.
func (c *Client) HandleAtomicOps(ctx context.Context, ops []entrypoint.PackedUserOperation) ([]common.Hash, common.Hash, error) {
	if c.config.ExecutorSigner == nil {
		panic("executor signer is nil")
	}
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.config.ExecutorSigner, c.chainId)
	if err != nil {
		return []common.Hash{}, common.Hash{}, fmt.Errorf("error creating transaction options: %v", err)
	}
	tx, err := c.entrypoint.HandleAtomicOps(txOpts, ops, crypto.PubkeyToAddress(c.config.ExecutorSigner.PublicKey))
	if err != nil {
		return []common.Hash{}, common.Hash{}, fmt.Errorf("error handling atomic ops: %v", err)
	}
	var opHashes []common.Hash
	for _, op := range ops {
		hashed, err := HashedUserOp(&op)
		if err != nil {
			return []common.Hash{}, common.Hash{}, fmt.Errorf("error hashing user operation: %v", err)
		}
		opHashes = append(opHashes, hashed)
	}
	return opHashes, tx.Hash(), nil
}

// Prefund deposits to entrypoint and waits for the transaction to be mined.
func (c *Client) Prefund(ctx context.Context, to common.Address, amount *big.Int) (*types.Receipt, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(c.config.VerifyingSigner, c.chainId)
	if err != nil {
		return nil, fmt.Errorf("error creating transactor: %v", err)
	}
	txOpts.Value = amount
	tx, err := c.entrypoint.DepositTo(txOpts, to)
	if err != nil {
		return nil, fmt.Errorf("error depositing fund: %v", err)
	}
	return bind.WaitMined(ctx, c.eth, tx)
}

// DeployAccount deploys the smart account and waits for the transaction to be mined.
func (c *Client) DeployAccount(ctx context.Context, signer *ecdsa.PrivateKey, account common.Address, salt *big.Int) (*types.Receipt, error) {
	txOpts, err := bind.NewKeyedTransactorWithChainID(signer, c.chainId)
	if err != nil {
		return nil, fmt.Errorf("error creating transactor: %v", err)
	}
	tx, err := c.simpleFactory.CreateAccount(txOpts, account, salt)
	if err != nil {
		return nil, fmt.Errorf("error creating account: %v", err)
	}
	return bind.WaitMined(ctx, c.eth, tx)
}

// ChainId returns the chain ID of the node.
func (c *Client) ChainId() *big.Int {
	return c.chainId
}

// FactoryABI returns the ABI of the account factory.
func (c *Client) FactoryABI() *abi.ABI {
	return c.simpleFactoryABI
}

// ACcountABI returns the ABI of the account contract.
func (c *Client) AccountABI() *abi.ABI {
	return c.simpleAccountABI
}
