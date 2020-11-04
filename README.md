# KITSquid

[![](http://img.shields.io/liberapay/receives/muety.svg?logo=liberapay&style=flat-square)](https://liberapay.com/muety/)
[![Say thanks](https://img.shields.io/badge/SayThanks.io-%E2%98%BC-1EAEDB.svg?style=flat-square)](https://saythanks.io/to/n1try)
![](https://img.shields.io/github/license/muety/kitsquid?style=flat-square)
[![Go Report Card](https://goreportcard.com/badge/github.com/muety/kitsquid?style=flat-square)](https://goreportcard.com/report/github.com/muety/kitsquid)
[![](https://img.shields.io/website.svg?url=https%3A%2F%2Fkitsquid.de%2Fhealth&style=flat-square)](https://kitsquid.de)

---

<img src="app/public/assets/images/squid_green.png" width="128">

🗒 An independent, private hobby project aiming to provide an alternative course catalog to students of [Karlsruhe Institute Of Technology](https://kit.edu). Students are given the ability to rate and comment lectures and provide valuable tips and advice to their fellow students. 

## ⚙️ Requirements
* Go >= 1.13
* NodeJS >= 9.8

## ▶️ Usage
* Get the code: `go get github.com/n1try/kitsquid`
* Change directory: `cd $GOPATH/src/github.com/n1try/kitsquid`
* Copy `config.example.yml` to `config.yml` and adapt it to your needs
* Go to `app/public` and run `yarn` and `yarn build` (or `yarn dev` for development purposes) 
* Go back to the project root and run `GO111MODULE=on go build`
* Run the application: `./kitsquid`

## 👩‍🎓 Community
This is a community project rather than commercial software. It certainly contains bugs or even security issues. We kindly ask you to report such as you find them, rather than trying to exploit them. Moreover, a community lives from its members, so please be kind. And of course, contribute as much as possible 😉.

## ⚠️ Disclaimer
As this is a non-official project and not related to the KIT by any means, please try to be responsible. Do not start scraping data and battering KIT's web services without permission. 

## 📓 License
MIT @ [Ferdinand Mütsch](https://muetsch.io)
