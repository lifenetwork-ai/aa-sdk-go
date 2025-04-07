package aasdk

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const (
	jsonrpcVersion     = "2.0"
	defaultWaitTimeout = 30 * time.Second
)

func (c *Client) GetUserOpReceipt(ctx context.Context, hash common.Hash) (*UserOpReceipt, error) {
	bytes, err := c.call("eth_getUserOperationReceipt", []any{hash})
	if err != nil {
		return nil, fmt.Errorf("error calling eth_getUserOperationReceipt: %v", err)
	}
	var response jsonRpcResponse[*UserOpReceipt]
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling user op receipt: %v", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("error from bundler: %s", response.Error.String())
	}
	return response.Result, nil
}

func (c *Client) EstimateUserOpGas(ctx context.Context, userOp *UserOperation) (*GasEstimates, error) {
	bytes, err := c.call("eth_estimateUserOperationGas", []any{userOp.ToBody(), c.config.Entrypoint})
	if err != nil {
		return nil, fmt.Errorf("error calling eth_estimateUserOperationGas: %v", err)
	}
	type gasEstimates struct {
		PreVerificationGas   *string `json:"preVerificationGas"`
		VerificationGasLimit *string `json:"verificationGasLimit"`
		CallGasLimit         *string `json:"callGasLimit"`
		VerificationGas      *string `json:"verificationGas"`
		MaxFeePerGas         *string `json:"maxFeePerGas"`
		MaxPriorityFeePerGas *string `json:"maxPriorityFeePerGas"`
	}
	var response jsonRpcResponse[*gasEstimates]
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling user gas estimates: %v", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("error from bundler: %s", response.Error.String())
	}
	if response.Result == nil {
		return nil, fmt.Errorf("no gas estimates response")
	}

	result := &GasEstimates{}
	if response.Result.PreVerificationGas != nil {
		result.PreVerificationGas = HexToBigInt(*response.Result.PreVerificationGas)
	}
	if response.Result.VerificationGasLimit != nil {
		result.VerificationGasLimit = HexToBigInt(*response.Result.VerificationGasLimit)
	}
	if response.Result.CallGasLimit != nil {
		result.CallGasLimit = HexToBigInt(*response.Result.CallGasLimit)
	}
	if response.Result.VerificationGas != nil {
		result.VerificationGas = HexToBigInt(*response.Result.VerificationGas)
	}
	if response.Result.MaxFeePerGas != nil {
		result.MaxFeePerGas = HexToBigInt(*response.Result.MaxFeePerGas)
	}
	if response.Result.MaxPriorityFeePerGas != nil {
		result.MaxPriorityFeePerGas = HexToBigInt(*response.Result.MaxPriorityFeePerGas)
	}
	return result, nil
}

func (c *Client) SupportedEntryPoints(ctx context.Context) ([]common.Address, error) {
	bytes, err := c.call("eth_supportedEntryPoints", nil)
	if err != nil {
		return nil, fmt.Errorf("error calling eth_supportedEntryPoints: %v", err)
	}
	var response jsonRpcResponse[[]common.Address]
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, fmt.Errorf("error unmarshalling entry points: %v", err)
	}
	if response.Error != nil {
		return nil, fmt.Errorf("error from bundler: %s", response.Error.String())
	}
	return response.Result, nil
}

func (c *Client) SendUserOp(ctx context.Context, userOp *UserOperation, signer *ecdsa.PrivateKey) (common.Hash, error) {
	signed, hash, err := c.FillAndSign(ctx, userOp, signer)
	if err != nil {
		return hash, fmt.Errorf("error fill and sign userop: %v", err)
	}

	bytes, err := c.call("eth_sendUserOperation", []any{signed.ToBody(), c.config.Entrypoint})
	if err != nil {
		return common.Hash{}, fmt.Errorf("error calling eth_sendUserOperation: %v", err)
	}

	var response jsonRpcResponse[common.Hash]
	if err = json.Unmarshal(bytes, &response); err != nil {
		return common.Hash{}, fmt.Errorf("error unmarshalling when sending user operation: %v", err)
	}
	if response.Error != nil {
		return common.Hash{}, fmt.Errorf("error from bundler: %s", response.Error.String())
	}
	return response.Result, nil
}

func (c *Client) GetUserOpHash(ctx context.Context, userOp *UserOperation, signer *ecdsa.PrivateKey) (common.Hash, error) {
	_, hash, err := c.FillAndSign(ctx, userOp, signer)
	if err != nil {
		return common.Hash{}, fmt.Errorf("error fill and sign userop: %v", err)
	}
	return hash, nil
}

func (c *Client) WaitForUserOperation(ctx context.Context, hash common.Hash) (*UserOpReceipt, error) {
	ticker := time.NewTicker(c.config.WaitReceiptInterval)
	defer ticker.Stop()
	ctx, cancel := context.WithTimeout(ctx, defaultWaitTimeout)
	defer cancel()
	for {
		select {
		case <-ticker.C:
			receipt, err := c.GetUserOpReceipt(ctx, hash)
			if err != nil {
				return nil, fmt.Errorf("error getting user operation receipt: %v", err)
			}
			if receipt != nil {
				return receipt, nil
			}
		case <-ctx.Done():
			return nil, fmt.Errorf("no receipt found for user operation %s", hash.Hex())
		}
	}
}

// call makes a JSON-RPC call to the bundler.
func (c *Client) call(method string, params []any) ([]byte, error) {
	if params == nil {
		params = []any{}
	}

	request := map[string]any{
		"jsonrpc": jsonrpcVersion,
		"id":      c.id.Add(1),
		"method":  method,
		"params":  params,
	}
	payloadBytes, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshalling payload: %v", err)
	}

	payload := strings.NewReader(string(payloadBytes))
	req, err := http.NewRequest("POST", c.config.BundlerUrl, payload)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}
	req.Header.Add("Content-Type", "application/json")
	res, err := c.http.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}
	return body, nil
}

type jsonRpcResponse[T any] struct {
	JsonRpc *string        `json:"jsonrpc"`
	Id      *int           `json:"id"`
	Result  T              `json:"result"`
	Error   *errorResponse `json:"error"`
}

type errorResponse struct {
	Code    *int    `json:"code"`
	Message *string `json:"message"`
}

// UnmarshalJSON implements custom unmarshaling for ErrorResponse
func (e *errorResponse) UnmarshalJSON(b []byte) error {
	// First, try to unmarshal as a simple string
	var errStr string
	if err := json.Unmarshal(b, &errStr); err == nil && errStr != "" {
		// If it's a string, set the message and leave code nil
		e.Message = &errStr
		e.Code = nil
		return nil
	}

	// Otherwise, try to unmarshal as an object
	type Alias struct {
		Code    *int    `json:"code"`
		Message *string `json:"message"`
	}
	var alias Alias
	if err := json.Unmarshal(b, &alias); err != nil {
		return err
	}

	// Populate the fields from the object
	e.Code = alias.Code
	e.Message = alias.Message
	return nil
}

func (e *errorResponse) String() string {
	result := ""
	if e.Code != nil {
		result += fmt.Sprintf("code: %d", *e.Code)
	}
	if e.Message != nil {
		if result != "" {
			result += ", "
		}
		result += fmt.Sprintf("message: %s", *e.Message)
	}
	return result
}
