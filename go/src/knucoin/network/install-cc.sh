#!/bin/bash
set -ev

# install chaincode for channelsales1
docker exec cli peer chaincode install -n test-cc -v 1.2 -p chaincode/go
sleep 1
# instantiate chaincode for channelsales1
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -v 1.2 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"
sleep 10
# invoke chaincode for channelsales1
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -c '{"function":"initWallet","Args":["byun618"]}'
docker exec cli peer chaincode invoke -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -c '{"function":"setMusic","Args":["Fabric", "Hyper", "20", "1Q2W3E4R"]}'
sleep 3
# query chaincode for channelsales1
docker exec cli peer chaincode query -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -c '{"function":"getWallet","Args":["byun618"]}'

docker exec cli peer chaincode install -n test-cc -v 1.2 -p chaincode/go
docker exec cli peer chaincode instantiate -o orderer.knucoin.com:7050 -C channelsales1 -n test-cc -v 1.2 -c '{"Args":[""]}' -P "OR ('SalesOrg.member','CustomerOrg.member')"