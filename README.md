简体中文 | [English](./README-en.md)

# BSN-SDK-GO-alpha
## 缘起
* [BSN-SDK-GO](http://kb.bsnbase.com/webdoc/view/PubFile2c908ad371c6396b01725ea21e1b2832.html)是[BSN官方](https://www.bsnbase.com/)提供的基于golang语言开发[Hyperledger-Fabric](https://www.hyperledger.org/)区块链项目的SDK。
* 由于官方未提供基于github下载的SDK，需要在项目代码中包含SDK的源码，使用稍有麻烦。因此在原有代码基础上,仅仅使用go mod处理依赖组件，未修改源代码，经过测试后上载github，供个人/团队测试开发使用。
* 欢迎对本SDK提出Issues进行改进。
* 如果BSN官方后期发布基于github的SDK，本SDK则停止使用。
* 本SDK代码版权属于BSN。

## 使用方法
* 在项目目录执行如下语句即可：
```
    cd yourproject
    go mod init 
    go get github.com/chcp/bsn-sdk-go
```
* 生成的go.mod内容如下：
```
    module yourproject
    
    go 1.14
    
    require github.com/chcp/bsn-sdk-go v1.1.1

```