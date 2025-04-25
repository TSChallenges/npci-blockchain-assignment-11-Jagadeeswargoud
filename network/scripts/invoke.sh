#!/bin/bash

# Set environment variables for TataMotors
source ./scripts/envVar.sh tataMotors

# Request LOC
echo "Requesting LOC..."
peer chaincode invoke -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C $CHANNEL_NAME -n $CC_NAME --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/tatamotors.com/peers/peer0.tatamotors.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/tesla.com/peers/peer0.tesla.com/tls/ca.crt -c '{"function":"RequestLOC","Args