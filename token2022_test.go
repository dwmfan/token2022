package token2022

import (
	"testing"

	solana "github.com/gagliardetto/solana-go"
)

func TestCreate2022Instruction(t *testing.T) {

	var (
		wallet = solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun")
		payer  = solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun")
		mint   = solana.MustPublicKeyFromBase58("D8zFabAK4Jt2Wi1TZJvMnr6EeD9K4qpiGhya1NQpyrZn")
	)

	instruction := NewCreate2022Instruction(payer, wallet, mint)

	err := instruction.Validate()
	if err != nil {
		t.Fatalf("Error validating instruction: %v", err)
	}

	built := instruction.Build()

	if len(built.Accounts()) != 7 {
		t.Errorf("Expected 7 accounts, got %d", len(built.Accounts()))
	}

	if built.Accounts()[5].PublicKey != solana.Token2022ProgramID {
		t.Errorf("Expected Token 2022 program ID, got %s", built.Accounts()[5].PublicKey)
	}
}

func TestFindAssociatedTokenAddress2022(t *testing.T) {

	var (
		wallet          = solana.MustPublicKeyFromBase58("nrw1b6stoyvm3QPsh78iWoJwsjM1b7KfcvxYT3LbFun")
		mint            = solana.MustPublicKeyFromBase58("D8zFabAK4Jt2Wi1TZJvMnr6EeD9K4qpiGhya1NQpyrZn")
		expectedAddress = solana.MustPublicKeyFromBase58("83mctxW8BCh6nPGjxx4jmyaEfbpcMZpLQiv7tXVSAV7a")
	)

	address, _, err := FindAssociatedTokenAddress2022(wallet, mint)
	if err != nil {
		t.Fatalf("Error finding associated token address: %v", err)
	}

	if address != expectedAddress {
		t.Errorf("Expected address %s, got %s", expectedAddress, address)
	}
}
