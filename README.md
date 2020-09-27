# KNU Coin With Hyperledger Fabric
- 대학교 축제 등 행사에 사용 할 가상 화폐

# **Technologies**  

## **Hyperledger Fabric of Blockchain**     
    
  - **OS**
    - Linux Ubuntu   
     
  - **Default Env**  
    - Go  
    - Docker  
    - Docker Compose  
    - NodeJS + NPM  

---
## 개발 파트 및 구성 인원

__PM:__ `변상현`    
__ChainCode + Network__ = `송동명, 박유민, 이상훈`  
__Server :__ `허강주`  

---

### 회의 및 개발 log

[dev_log.md](https://github.com/byun618/KNUCoin/blob/master/dev_log.md)

---

## 2020-09-21

![KakaoTalk_Photo_2020-09-21-18-41-38](https://user-images.githubusercontent.com/56465854/93752838-2fe41180-fc3a-11ea-88ad-2195e65e3743.png)

---

### For Linux user

1. Creat a new directory. somewhere to keep both the fabric-samples, and the KNUCoin repo:

```shell
$ mkdir fabric-repo 
```


2. Go into your newly created repo, and download the latest production release from Fabric. If you have any errors downloading the binaries, you may need to install wget on your system. 


```Shell
$ cd fabric-repo
fabric-repo$ curl -sSL http://bit.ly/2ysbOFE | bash -s -- 1.4.4 1.4.4 0.4.18 
```


3. After the downloads are complete, you should see a newly created fabric-samples repo. Next, let's go ahead and clone the raft-fabric-sample repo and then cd into it.


```shell
fabric-repo$ git clone https://github.com/byun816/KNUCoin.git
fabric-repo$ cd KNUCoin 
```


4. Remove the bin folder from KNUCoin (since it assumes platform binaries that are made for MacOS).

```shell
fabric-repo/KNUCoin/go/src/knucoin/network$ rm -rf bin/
```

5. Copy the bin folder from fabric-samples, and paste it into the KNUCoin folder

```shell
fabric-repo/KNUCoin/go/src/knucoin/network$ mkdir bin && cp -r ../../../../../fabric-samples/bin . 
```

6. If you do a ls in your bin folder within your KNUCoin, you should see the following:

```shell
fabric-repo/KNUCoin/go/src/knucoin/network/bin$ ls
configtxgen  configtxlator  cryptogen  discover  fabric-ca-client  fabric-ca-server  idemixgen  orderer  peer 
```



