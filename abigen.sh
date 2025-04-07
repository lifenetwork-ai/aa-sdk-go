#!/bin/sh

rm ./bindings/*/*.go 2>/dev/null

abigen -abi ./abis/entrypoint_v7.json \
    -pkg entrypoint \
    -type EntryPoint \
    -out ./bindings/entrypoint/entrypoint.go

abigen -abi ./abis/simple_account_factory.json \
    -pkg account \
    -type SimpleAccountFactory \
    -out ./bindings/account/simple_account_factory.go

abigen -abi ./abis/simple_account.json \
    -pkg account \
    -type SimpleAccount \
    -out ./bindings/account/simple_account.go

abigen -abi ./abis/verifying_paymaster.json \
    -pkg paymaster \
    -type VerifyingPaymaster \
    -out ./bindings/paymaster/verifying_paymaster.go
