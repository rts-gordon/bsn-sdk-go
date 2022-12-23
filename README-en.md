English | [简体中文](./README.md)

# BSN-SDK-GO-alpha

## Why
* [BSN-SKD-GO](http://kb.bsnbase.com/webdoc/view/PubFile2c908ad371c6396b01725ea21e1b2832.html) is an SDK for [Hyperledger-Fabric](https://www.hyperledger.org/) blockchain project based on golang language provided by [BSN official](https://www.bsnbase.com/).
* BSN does not provide SDK on github for download, so SDK source code needs to be included in the project code, which is a little troublesome to use. Below is the solution to the problem:
 based on the original source code , only the go mod was used to process dependent components, and source code was not modified. It has been tested, and then uploaded to github for personal/team development.
* Issues are welcome to be improved on this SDK.
* If the official BSN later releases the github-based SDK, this SDK will be discontinued.
* The copyright of this SDK code belongs to BSN.

## Usage
* in your project folder:
```
    cd yourproject
    go mod init 
    go get github.com/rts-gordon/bsn-sdk-go
```
* generate file go.mod:
```
    module yourproject
    
    go 1.14
    
    require github.com/rts-gordon/bsn-sdk-go v1.1.1
```