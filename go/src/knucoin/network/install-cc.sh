#!/bin/bash
set -ev

# install chaincode for channelsales1
docker exec cli peer chaincode install -n coin2-cc -v 1.0 -p chaincode/go
sleep 1
# instantiate chaincode for channelsales1
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n coin2-cc -v 1.0 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"
sleep 10
# invoke chaincode for channelsales1
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n coin2-cc -c '{"function":"initWallet","Args":[""]}'
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n coin2-cc -c '{"function":"sendMoney","Args":["myung","musk","1000"]}'
sleep 3
# query chaincode for channelsales1
docker exec cli peer chaincode query -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -c '{"function":"getWallet","Args":["myung"]}'
