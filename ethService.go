package main

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/samuelgoes/ethereum_test/contract"
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
	// Use Infura:
	infura := "https://rinkeby.infura.io/v3/89cdc30a5a8c440abfa323d028e0c70f"
	cl, err := ethclient.Dial(infura)
	*/

	defer cl.Close()

	hexPrivKey := "0fd57c00030c11b9a1f636dd0e065b78f2afb4eb7cdef9d88ed5eacf39b45930"
	key, err := crypto.HexToECDSA(hexPrivKey)
	if err != nil {
		log.Fatalf("Address is not OK. %v", err)
	}

	addr := common.HexToAddress("0x9412CbAd85F371CAa6ffC2A1956204d1d6362524")  //Ganache

	ctx := context.Background()

	// Retrieve a block by number
	block, err := cl.BlockByNumber(ctx, big.NewInt(37))
	if err != nil {
		log.Printf("error getting block number: %v", err)
	} else {
		log.Printf("Block: Transactaions: %v", block.Transactions())
	}

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

	// ****************** DEPLOY ******************

	publicKey := key.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("error reading next nonce")
	}

	gasPrice, err := cl.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("error reading suggest gas price")
	}

	chainID, err := cl.ChainID(context.Background())
	if err != nil {
		log.Fatalf("unable to get chainID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		log.Fatalf("unable to build new transactor: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	address, _, con, err := contract.DeployContract(auth, cl)
	if err != nil {
		log.Fatalf("unable to deploy smart contract. %v", err)
	}
	log.Printf("Smart Contract desplegado satisfactoriamente. Address: %s", address.Hex())

	/*
	// Load Smart Contract
	address = common.HexToAddress("0xB268907357B70F246ef1dB60c344fB15418F3efb")
	con, err = contract.NewContract(address, cl)
	if err != nil {
		log.Fatalf("unable to deploy smart contract")
	}
	log.Printf("Smart Contract cargado satisfactoriamente.")
	*/

	nonce, err = cl.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatalf("error reading next nonce")
	}
	auth.Nonce = big.NewInt(int64(nonce))

	_, err = con.StoreMessage(auth, "Hola Samu")
	if err != nil {
		log.Fatalf("unable to call store message function. %v", err)
	}

	log.Printf("Mensaje almacenado satisfactoriamente")

	message, err := con.RetrieveMessage(nil)
	if err != nil {
		log.Fatalf("unable to call store message function")
	}

	log.Printf("Este es el mensaje almacenado en el Smart Contract: %s", message)
}
