package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// LetterOfCredit defines the structure for the Letter of Credit
type LetterOfCredit struct {
	LOCID            string   `json:"locId"`
	Buyer            string   `json:"buyer"`
	Seller           string   `json:"seller"`
	IssuingBank      string   `json:"issuingBank"`
	AdvisingBank     string   `json:"advisingBank"`
	Amount           string   `json:"amount"`
	Currency         string   `json:"currency"`
	ExpiryDate       string   `json:"expiryDate"`
	GoodsDescription string   `json:"goodsDescription"`
	Status           string   `json:"status"`
	DocumentHashes   []string `json:"documentHashes"`
	History          []string `json:"history"`
}

// SmartContract provides functions for managing the Letter of Credit
type SmartContract struct {
	contractapi.Contract
}

// InitLedger initializes the chaincode (optional)
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	// TODO: Initialization code if needed
	return nil
}

// RequestLOC creates a new LoC request
func (s *SmartContract) RequestLOC(ctx contractapi.TransactionContextInterface, locID string, buyer string, seller string, issuingBank string, advisingBank string, amount string, currency string, expiryDate string, goodsDescription string) error {

	clientmpsid, err := ctx.GetClientIdentity.GetMSPID()

	if err != nil {
		return fmt.Errorf("Error while fetching the MSP ID: %v", err)
	}

	if clientmpsid != "TataMotorsMSP" {
		return fmt.Errorf("Only tata motors have the privelege to request the LOC: %v", err)
	}

	letterofcredit := LetterOfCredit{
		LOCID:            locID,
		Buyer:            buyer,
		Seller:           seller,
		IssuingBank:      issuingBank,
		AdvisingBank:     advisingBank,
		Amount:           amount,
		Currency:         currency,
		ExpiryDate:       expiryDate,
		GoodsDescription: goodsDescription,
		Status:           "Requested",
		DocumentHashes:   []string{"123"},
		History:          []string{"hist1"},
	}

	exists, err := ctx.GetStub().GetState(locID)

	if err != nil {
		return fmt.Errorf("Error while fetching the loc id: %v", err)
	}

	if exists != nil {
		return fmt.Errorf("loc id already exists in the ledger: %v", err)
	}

	locjson, err := json.Marshal(letterofcredit)

	if err != nil {
		return fmt.Errorf("Error while marshal the loc data: %v", err)
	}

	if locjson == nil {
		return fmt.Errorf("Empty values passed for the loc data: %v", err)
	}

	err = ctx.GetStub.PutState(locID, locjson)

	if err != nil {
		return fmt.Errorf("Error while put state of the loc data: %v", err)
	}

	// Add to history
	return nil
}

func (s *SmartContract) IssueLOC(ctx contractapi.TransactionContextInterface, locID string, buyer string, seller string, issuingBank string, advisingBank string, amount string, currency string, expiryDate string, goodsDescription string) error {

	clientmpsid, err := ctx.GetClientIdentity.GetMSPID()

	if err != nil {
		return fmt.Errorf("Error while fetching the MSP ID: %v", err)
	}

	if clientmpsid != "TataMotorsMSP" {
		return fmt.Errorf("Only tata motors have the privelege to request the LOC: %v", err)
	}

	letterofcredit := LetterOfCredit{
		LOCID:            locID,
		Buyer:            buyer,
		Seller:           seller,
		IssuingBank:      issuingBank,
		AdvisingBank:     advisingBank,
		Amount:           amount,
		Currency:         currency,
		ExpiryDate:       expiryDate,
		GoodsDescription: goodsDescription,
		Status:           "Requested",
		DocumentHashes:   []string{"123"},
		History:          []string{"hist1"},
	}

	exists, err := ctx.GetStub().GetState(locID)

	if err != nil {
		return fmt.Errorf("Error while fetching the loc id: %v", err)
	}

	if exists != nil {
		return fmt.Errorf("loc id already exists in the ledger: %v", err)
	}

	locjson, err := json.Marshal(letterofcredit)

	if err != nil {
		return fmt.Errorf("Error while marshal the loc data: %v", err)
	}

	if locjson == nil {
		return fmt.Errorf("Empty values passed for the loc data: %v", err)
	}

	err = ctx.GetStub.PutState(locID, locjson)

	if err != nil {
		return fmt.Errorf("Error while put state of the loc data: %v", err)
	}

	// Add to history
	return nil
}

// TODO: Implement other functions here (IssueLOC, AcceptLOC, etc.)

func main() {
	chaincode, err := contractapi.NewChaincode(&SmartContract{})
	if err != nil {
		fmt.Printf("Error creating loc chaincode: %s", err.Error())
		return
	}

	if err := chaincode.Start(); err != nil {
		fmt.Printf("Error starting loc chaincode: %s", err.Error())
	}
}
