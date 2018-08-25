# go-imap

An interactive command-line tool for learn [IMAP4rev1](https://tools.ietf.org/html/rfc3501) protocol details.

## Why?
当我开始学习IMAP协议时，我发现，github上现有的关于imap的golang项目都是封装成`library`以便快速搭建imap的客户端和服务器程序。这对于快速构建相关产品是有用的，但对于那些想要探索imap底层通信细节的同学，可能并不太适合。

因此，本项目的目的就是提供一个可交互的命令行工具，可直接使用[IMAP4rev1](https://tools.ietf.org/html/rfc3501)规定的命令与服务器通信，并且可以非常详细的查看通信细节。(也就是说，本项目实际上只是一个imap client)

> 后来我才知道，原来可以直接使用`telnet`程序进行交互([用telnet验证imap](http://blog.51cto.com/linuxroad/1000530))，但`telnet`只能连接到未使用ssl加密的服务，也就是143端口。而想要连接993端口，就无法通过`telnet`程序实现了。而本程序一开始就是测试的993端口。
>
> 也就是说，本程序其实只不过是一个`telnet`程序而已。但我发现，用`telnet`程序的返回报文中，对于中文没有进行解码，也就是无法显示中文。所以本项目也是有意义的:)

## go-imap VS telnet with imap
todo

## Install
```cmd
go get github.com/champkeh/go-imap
```

## Usage
```
go-imap -addr [imap-server-address]
```

## Example
### 1.startup
```shell
$ go-imap -addr imap.qq.com:993
connecting to imap.qq.com:993
connected
S[84]:* OK [CAPABILITY IMAP4 IMAP4rev1 ID AUTH=LOGIN NAMESPACE] QQMail IMAP4Server ready
imap>
```

### 2. login
```shell
$ go-imap -addr imap.qq.com:993
connecting to imap.qq.com:993
connected
S[84]:* OK [CAPABILITY IMAP4 IMAP4rev1 ID AUTH=LOGIN NAMESPACE] QQMail IMAP4Server ready
imap>LOGIN 1481536930 ***(此处是密码)
C[41]:A0001 LOGIN 1481536930 ***(此处是密码)
S[27]:A0001 OK Success login ok
imap>
```

### 3. list
```
imap>LIST "" %
C[17]:A0002 LIST "" %
S[324]:* LIST (\NoSelect \HasChildren) "/" "&UXZO1mWHTvZZOQ-"
* LIST (\HasNoChildren) "/" "INBOX"
* LIST (\HasNoChildren) "/" "Sent Messages"
* LIST (\HasNoChildren) "/" "Drafts"
* LIST (\HasNoChildren) "/" "Deleted Messages"
* LIST (\HasNoChildren) "/" "Junk"
* LIST (\HasNoChildren) "/" "Archive"
A0002 OK LIST completed
imap>
```

### 4. logout
```
imap>LOGOUT
C[14]:A0003 LOGOUT
S[50]:* BYE LOGOUT received
A0003 OK LOGOUT Completed
imap>2018/08/21 00:04:12 EOF
```


## Todo
- [ ] 添加命令历史记录，以便可以通过上下箭头键进行快速输入
- [ ] 自动检测 imap-server-address 是否是ssl
- [ ] 实现读取邮件内容的命令
