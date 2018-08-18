# go-imap

An interactive command-line tool for learn [IMAP4rev1](https://tools.ietf.org/html/rfc3501) protocol details.

## Why?
当我开始学习IMAP协议时，我发现，github上现有的关于imap的golang项目都是封装成`library`以便快速搭建imap的客户端和服务器程序。这对于快速构建相关产品是有用的，但对于那些想要探索imap底层通信细节的同学，可能并不太适合。

因此，本项目的目的就是提供一个可交互的命令行工具，可直接使用[IMAP4rev1](https://tools.ietf.org/html/rfc3501)规定的命令与服务器通信，并且可以非常详细的查看通信细节。(因此，此项目实际上只是一个imap client)

## Install
```cmd
go get github.com/champkeh/go-imap
```

## Usage
```
go-imap -addr [imap-server-address]
```

## Example
1. login


2. list


3. logout