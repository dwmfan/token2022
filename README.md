# Token2022

A Go package for Solana Token 2022 operations.

## Installation

```bash
go get github.com/dwnfan/token2022
```

## Usage

### Finding Associated Token Address for Token 2022

```go
package main

import (
    "fmt"
    "github.com/dwnfan/token2022"
    "github.com/gagliardetto/solana-go"
)

func main() {
    wallet := solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun")
    mint := solana.MustPublicKeyFromBase58("D8zFabAK4Jt2Wi1TZJvMnr6EeD9K4qpiGhya1NQpyrZn")
    
    address, _, err := token2022.FindAssociatedTokenAddress2022(wallet, mint)
    if err != nil {
        panic(err)
    }
    
    fmt.Println("Associated Token Address:", address.String())
}
```

### Creating an Associated Token Account for Token 2022

```go
package main

import (
    "context"
    "fmt"
    "github.com/dwnfan/token2022"
    "github.com/gagliardetto/solana-go"
    "github.com/gagliardetto/solana-go/rpc"
)

func main() {
    // Initialize client
    client := rpc.New("https://api.mainnet-beta.solana.com")
    
    // Define wallet, payer, and mint
    wallet := solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun")
    payer := solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun") // Using same as wallet for example
    mint := solana.MustPublicKeyFromBase58("D8zFabAK4Jt2Wi1TZJvMnr6EeD9K4qpiGhya1NQpyrZn")
    
    // Create instruction
    instruction := token2022.NewCreate2022Instruction(payer, wallet, mint)
    
    // Build and validate instruction
    built, err := instruction.ValidateAndBuild()
    if err != nil {
        panic(err)
    }
    
    // Create transaction
    recent, err := client.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
    if err != nil {
        panic(err)
    }
    
    tx, err := solana.NewTransaction(
        []solana.Instruction{built},
        recent.Value.Blockhash,
        solana.TransactionPayer(payer),
    )
    if err != nil {
        panic(err)
    }
    
    // Sign and send transaction (not shown)
    // ...
    
    fmt.Println("Transaction created successfully")
}
```

## License

MIT
