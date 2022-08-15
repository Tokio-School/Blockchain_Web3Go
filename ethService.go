package main

import (
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	_ "github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	_ "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	_ "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	_ "github.com/ethereum/go-ethereum/params"
	"log"
	"math/big"
)

func main() {
	// Use Ganache node:
	cl, err := ethclient.Dial("http://127.0.0.1:7545")
	if err != nil {
		log.Fatalf("error dialing eth client: %v", err)
	}

	/*
	// Use infura:
	infura := "https://rinkeby.infura.io/v3/89cdc30a5a8c440abfa323d028e0c70f"
	cl, err := ethclient.Dial(infura)
	*/

	addr := common.HexToAddress("0x9412CbAd85F371CAa6ffC2A1956204d1d6362524")  //Ganache

	ctx := context.Background()

	// Retrieve a block by number
	block, err := cl.BlockByNumber(ctx, big.NewInt(11205568))
	if err != nil {
		log.Fatalf("error getting block number: %v", err)
	}
	log.Printf("Block: Transactaions: %v", block.Transactions())


	// Get Balance of an account (nil means at newest block)
	balance, err := cl.BalanceAt(ctx, addr, nil)
	if err != nil {
		log.Fatalf("error getting balance: %v", err)
	}
	log.Printf("Balance: %v", balance)

	// Get sync progress of the node. If nil, the node is not syncing
	progress, err := cl.SyncProgress(ctx)
	if err != nil {
		log.Fatalf("error getting balance: %v", err)
	}
	log.Printf("Progress: %v", progress)


	// Retrieve the pending nonce for an account
	nonce, err := cl.NonceAt(ctx, addr, nil)
	to := common.HexToAddress("0xABCD")
	amount := big.NewInt(10 * params.GWei)
	gasLimit := uint64(21000)
	gasPrice := big.NewInt(10 * params.GWei)
	var data []byte

	// Create a raw unsigned transaction
	tx := types.NewTransaction(nonce, to, amount, gasLimit, gasPrice, data)

	// Use secret key hex string to sign a raw transaction
	SK := "0x0fd57c00030c11b9a1f636dd0e065b78f2afb4eb7cdef9d88ed5eacf39b45930"
	sk := crypto.ToECDSAUnsafe(common.FromHex(SK)) // Sign the transaction
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(nil), sk)
	// You could also create a TransactOpts object
	opts := bind.NewKeyedTransactor(sk)
	// To get the address corresponding to your private key
	addr := crypto.PubkeyToAddress(sk.PublicKey)

    /*
	// Open Keystore
	ks := keystore.NewKeyStore("/home/matematik/keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	// Create an account
	acc, err := ks.NewAccount("password")
	// List all accounts
	accs := ks.Accounts()
	// Unlock an account
	ks.Unlock(accs[0], "password")
	// Create a TransactOpts object
	ksOpts, err := bind.NewKeyStoreTransactor(ks, accs[0])
    */


	// If you have a bind.TransactOpts object you can sign a transaction
	sigTx, err := opts.Signer(types.NewEIP155Signer(nil), addr, tx)

}



