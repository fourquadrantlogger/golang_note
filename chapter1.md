# vscode mac安装过程

##第一步
安装VSCODE

##第二步
```
cmd-shift-p
```

选择 最新版本的go-estension，一定要是最新版本，否则就会无法debug

![](C76B1353-75CB-4E77-BB0D-59D034CAD333.png)

##第三步
安装插件，这一步需要翻墙下载golang.org里的包
```
go get -u -v github.com/nsf/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v github.com/golang/lint/golint
go get -u -v github.com/lukehoban/go-find-references
go get -u -v github.com/lukehoban/go-outline
go get -u -v sourcegraph.com/sqs/goreturns
go get -u -v golang.org/x/tools/cmd/gorename
```

##第四步
安装编译delve

https://github.com/derekparker/delve

首先，通过很多次go get对这个包进行gobuild，的到一个二进制文件



