#!/bin/bash

# Exit on first error
set -e

# Set channel name
CHANNEL_NAME="locchannel"

# Set chaincode details
CC_NAME="loccc"
CC_SRC_PATH="../chaincode/"
CC_LANG="golang"
CC_VERSION="1.0"
CC_SEQUENCE="1"

# Package the chaincode
echo "Packaging chaincode..."
peer lifecycle chaincode package ${CC_NAME}.tar.gz --path ${CC_SRC_PATH} --lang ${CC_LANG} --label ${CC_NAME}_${CC_VERSION}

# Install chaincode on all peers
echo "Installing chaincode on TataMotors peer..."
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/tatamotors.com/users/Admin@tatamotors.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_LOCALMSPID="TataMotorsMSP"
peer lifecycle chaincode install ${CC_NAME}.tar.gz

echo "Installing chaincode on Tesla peer..."
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/tesla.com/users/Admin@tesla.com/msp
export CORE_PEER_ADDRESS=localhost:9051
export CORE_PEER_LOCALMSPID="TeslaMSP"
peer lifecycle chaincode install ${CC_NAME}.tar.gz

# Similar for ICICIBank and ChaseBank

# Approve chaincode for each org
echo "Approving chaincode for TataMotors..."
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/tatamotors.com/users/Admin@tatamotors.com/msp
export CORE_PEER_ADDRESS=localhost:7051
export CORE_PEER_LOCALMSPID="TataMotorsMSP"
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL_NAME --name $CC_NAME --version $CC_VERSION --package-id $(peer lifecycle chaincode queryinstalled | grep -oP 'Package ID: \K.*(?=, Label:.*)') --sequence $CC_SEQUENCE --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem

# Similar approvals for other orgs

# Commit the chaincode
echo "Committing chaincode..."
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID $CHANNEL_NAME --name $CC_NAME --version $CC_VERSION --sequence $CC_SEQUENCE --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/tatamotors.com/peers/peer0.tatamotors.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/tesla.com/peers/peer0.tesla.com/tls/ca.crt

echo "Chaincode deployed successfully!"