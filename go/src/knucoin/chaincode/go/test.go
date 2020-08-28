package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	// "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct {}

type Wallet struct {
	ID   string `json:"id"`
	Token string `json:"token"`
}

type Music struct {
	Title    string `json:"title"`
	Singer   string `json:"singer"`
	Price    string `json:"price"`
	WalletID    string `json:"walletid"`
}

type MusicKey struct {
	Key string
	Idx int
}

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) pb.Response {
	function, args := APIstub.GetFunctionAndParameters()
	
	if function == "initWallet" {
		return s.initWallet(APIstub, args)
	} else if function == "getWallet" {
		return s.getWallet(APIstub, args)
	}
	fmt.Println("Please check your function : " + function)
	return shim.Error("Unknown function")
}

func (s *SmartContract) initWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	//Declare wallets
	user := Wallet{ID: args[0], Token: "0"}

	// Convert seller to []byte
	UserasJSONBytes, _ := json.Marshal(user)
	err := APIstub.PutState(user.ID, UserasJSONBytes)
	if err != nil {
		return shim.Error("Failed to create asset " + user.ID)
	}

	return shim.Success(UserasJSONBytes)
}

func (s *SmartContract) getWallet(APIstub shim.ChaincodeStubInterface, args []string) pb.Response {

	walletAsBytes, err := APIstub.GetState(args[0])
	if err != nil {
		fmt.Println(err.Error())
	}

	wallet := Wallet{}
	json.Unmarshal(walletAsBytes, &wallet)

	var buffer bytes.Buffer
	buffer.WriteString("[")
	bArrayMemberAlreadyWritten := false

	if bArrayMemberAlreadyWritten == true {
		buffer.WriteString(",")
	}

	buffer.WriteString(", \"ID\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.ID)
	buffer.WriteString("\"")

	buffer.WriteString(", \"Token\":")
	buffer.WriteString("\"")
	buffer.WriteString(wallet.Token)
	buffer.WriteString("\"")

	buffer.WriteString("}")
	bArrayMemberAlreadyWritten = true
	buffer.WriteString("]\n")

	return shim.Success(buffer.Bytes())

}

func main() {

	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}