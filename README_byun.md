# UFO

## 목차
[1. 프로젝트 개요](#프로젝트-개요)      
[2. 개발 파트](#개발-파트)  
[3. 개발 환경](#개발-환경)    
[4. 개발 내용](#개발-내용)  
[5. 파트별 Git Repository](#파트별-Git-Repository)  


## 프로젝트 개요 (2020.7 - )
* **명칭**  
**U**niversity **F**estival in **O**ne  
* **개요**  
전국 대학 축제를 하나의 플랫폼에서 확인하고, 고유 화폐로 결제
* **소개**   
각 대학의 축제를 하나의 어플로 확인하고 각종 정보를 확인할 수 있으며, 블록체인 *Hyperledger Fabric*으로 거래 네트워크를 구성해 안전하고 신빙성있는 거래환경 제공

## 개발 파트
 *BlockChain Network*, *Server*, *Application* 파트로 구성
### 내가 맡은 파트
* PM(Project Manager) 역할로 프로젝트 전반적인 일정 계획 및 각 파트 개발 진행
* Application
    * 주요 기능 개발 진행
* BlockChain
    * 초기 네트워크 구성시 설정 부분 맡아서 진행
* Server
    * *AWS 이관 작업 맡아서 진행할 예정*

*전 파트에 걸쳐 피드백 및 개발 진행, 공식적으로 맡은 파트는 Application*

## 개발 환경
https://github.com/byun618/UFO#%EA%B0%9C%EB%B0%9C-%ED%99%98%EA%B2%BD

## 개발 내용
* Mobile
    * 주요 기능
        * 카카오 로그인
        * 스토어, 메뉴 띄우기
        * Map 띄우기
        * QR코드 Scanner
        * SlideOver 

* BlockChain
    * Network 설정 부분
        ```
        configtx.yaml
        config-crypto.yaml
        ```
    * BlockChain Network 구동 부분
        ```
        docker-compose.yaml
        start.sh
        ```
    개발 초기 단계에 블록체인 개발 진행할 때, Network의 설정부분을 진행하고 구동 시킬때 Script파일로 동작할 수 있게 개발

* Server
    * AWS S3 구현
      * EC2에서 서버가 구동 되겠지만, 로컬환경에서 개발은 진행 될것이고, S3는 계속 쓰여야 하기에 각 팀원에게 AWS IAM 유저를 부여하고 Role과 Policy를 이용하여 관리

## 파트별 Git Repository
* Mobile    
[https://github.com/myungsworld/UFO_Mobile](https://github.com/myungsworld/UFO_Mobile)
* Fabric    
[https://github.com/yuminee/UFO_FabricNet](https://github.com/yuminee/UFO_FabricNet)
* Server    
[https://github.com/Cozy-Ho/UFO_Server](https://github.com/Cozy-Ho/UFO_Server)
