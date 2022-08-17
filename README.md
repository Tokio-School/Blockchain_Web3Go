# Web3Go
This is a project to test how the Web3Go (Go-Ethereum) is working to integrate a Smart Contract into a common application. 

In this case, the Smart contract used is based in Hola Mundo project, but the process to integrate any of this kind of applications is exactly the same.

Once the Smart Contract is ready in Solidity language, the process to transform the code and integrate in any code language is the following:

```shell
solc --abi HolaMundo.sol

abigen --abi=HolaMundo.abi --pkg=contract --out=HolaMundo.go

solc --bin HolaMundo.sol

abigen --bin=HolaMundo.bin --abi=HolaMundo.abi --pkg=contract --out=HolaMundo.go
```

At the end of this process a new file is extracted named HolaMundo.go. Inside the file, there is the SmartContract code based on GoLang language and a binary section where the Smart Contract is encoded to be deployed into the Blockchain network.
