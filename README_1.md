# UFO

## 목차
[1. 프로젝트 개요](#프로젝트-개요)  
[2. 개발 환경](#개발환경)    
[3. 개발 과정](#개발과정)    
[4. 개발 계획](#개발계획)       
[5. 개발 파트](#개발파트) 

## 프로젝트 개요
* **명칭**  
**U**niversity **F**estival in **O**ne  
* **개요**  
전국 대학 축제를 하나의 플랫폼에서 확인하고, 고유 화폐로 결제
* **소개**   
각 대학의 축제를 하나의 어플로 확인하고 각종 정보를 확인할 수 있으며, 블록체인 __Hyperledger Fabric__으로 거래 네트워크를 구성해 안전하고 신빙성있는 거래환경 제공

## 개발 환경
* NodeJs    
    * 서버 구성
* GO
    * Hyperledger Fabric의 Chaincode 개발 언어
* Docker & Docker Compose
    * Hyperledger Fabric 네트워크를 도커 컨테이너로 구성
* Swift
    * iOS mobile application 개발 언어
* AWS
    * EC2, LoadBalancer를 이용하여 NodeJS 서버와 Hyperledger Fabric 네트워크를 구성
    * RDS MySQL을 통해 축제 관련 정보, 유저 데이터 관리
    * S3를 이용해 Hyperledger Fabric의 인증서 및 축제 관련 이미지 등 정적데이터 관리
    * IAM 정책을 통해 각종 액세스 관리

## 개발 과정
* 기본적인 블록체인, Hyperledger Fabric 관련 공부 및 실습
* Server, Fabirc 파트로 나누어 각 파트 진행
* Fabric 파트 chaincode 및 solo 방식으로 구현
* Server, Mobile, Fabric 파트로 나누어 각각 개발 진행(2020.09.21)
* 각 파트의 repo를 새로 생성, sub module로 나누어 관리(2020.10.13)
* Fabric 파트 개발 완료로 Server, Mobile로 파트 재편성(2020.11.03)
* 자세한 사항은 각 파트 repo 확인.

## 개발 파트
* __PM__ : `변상현`    
* ~~__ChainCode + Network__ : `송동명 박유민 이상훈`~~
* ~~__Server__ : `허강주`~~ 
* ~~__Fabric__ : `박유민 송동명`~~
* ~~__Server__ : `허강주 이상훈`~~
* ~~__Mobile__ : `변상현 송동명`~~
* __Server__ : `허강주 이상훈 변상현`
* __Mobile__ : `변상현 송동명 박유민`




