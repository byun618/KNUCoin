# UFO

## 목차
[1. 프로젝트 개요](#프로젝트-개요)  
[2. 개발 리뷰](#개발-리뷰)   
[3. 개발 환경](#개발-환경)    
[4. 개발 과정](#개발-과정)    
[5. 개발 계획](#개발-계획)       
[6. 개발 파트](#개발-파트) 

## 프로젝트 개요 (2020.7 - )
* **명칭**  
**U**niversity **F**estival in **O**ne  
* **개요**  
전국 대학 축제를 하나의 플랫폼에서 확인하고, 고유 화폐로 결제
* **소개**   
각 대학의 축제를 하나의 어플로 확인하고 각종 정보를 확인할 수 있으며, 블록체인 **Hyperledger Fabric**으로 거래 네트워크를 구성해 안전하고 신빙성있는 거래환경 제공

## 개발 리뷰
* 기존에 계획하였던 AWS로 REST서버와 Fabric서버를 올리려고 하였던것을 리소스의 제한으로 docker와 docker swarm을 활용하여 구성
* Mobile에서 기본적인 기능은 다 구현하였으나, 사용자 잔액 조회, 사용자의 Org 변경 등 미처 개발하지 못한 기능 존재
* 개발 로그를 꾸준히 작성하지 못한점이 아쉬움

## 개발 환경
* NodeJs    
    * 서버 개발
    * Hyperledger Fabric의 Chaincode 개발 언어
* ~~GO~~
    * ~~Hyperledger Fabric의 Chaincode 개발 언어~~
* Docker & Docker Compose
    * Hyperledger Fabric 네트워크를 도커 컨테이너로 구성
    * REST 서버 구성
* Swift
    * iOS mobile application 개발 언어
* AWS
    * ~~EC2, LoadBalancer를 이용하여 NodeJS 서버와 Hyperledger Fabric 네트워크를 구성~~
    * ~~RDS MySQL을 통해 축제 관련 정보, 유저 데이터 관리~~
    * ~~S3를 이용해 Hyperledger Fabric의 인증서 및 축제 관련 이미지 등 정적데이터 관리~~
    * S3를 이용해 애플리케이션을 위한 이미지 버킷
    * ~~IAM 정책을 통해 각종 액세스 관리~~

## 개발 파트
* __PM__ : `변상현`    
* ~~__ChainCode + Network__ : `송동명 박유민 이상훈`~~  
~~__Server__ : `허강주`~~
* ~~__Fabric__ : `박유민 송동명`~~  
~~__Server__ : `허강주 이상훈`~~   
~~__Mobile__ : `변상현 송동명`~~
* ~~__Server__ : `허강주 이상훈 변상현`~~   
 ~~__Mobile__ : `변상현 송동명 박유민`~~   
* __Server__ : `허강주 이상훈 변상현`   
 __Mobile__ : `변상현 송동명`   
 __Fabric__ : `박유민 송동명` 

## 개발 과정
* 기본적인 블록체인, Hyperledger Fabric 관련 공부 및 실습
* Server, Fabirc 파트로 나누어 각 파트 진행
* Fabric 파트 chaincode 및 solo 방식으로 구현
* Server, Mobile, Fabric 파트로 나누어 각각 개발 진행(2020.09.21)
* 각 파트의 repo를 새로 생성, sub module로 나누어 관리(2020.10.13)
* Fabric 파트 개발 완료로 Server, Mobile로 파트 재편성(2020.11.03)
* 파트를 재편성 하였으나 Fabric 업데이트로 인해 다시 편성
* 프로젝트 마무리(2020.12.10)
* 자세한 사항은 각 파트 repo 확인.

## 파트별 Git Repository
* Fabric    
[https://github.com/yuminee/UFO_FabricNet_raft.git](https://github.com/yuminee/UFO_FabricNet_raft.git)
* Server    
[https://github.com/Cozy-Ho/UFO_Server](https://github.com/Cozy-Ho/UFO_Server)
* Mobile    
[https://github.com/myungsworld/UFO_Mobile](https://github.com/myungsworld/UFO_Mobile)