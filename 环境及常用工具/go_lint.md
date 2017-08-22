1. 点击Edit Configuration, 添加一个 `Go Application`，设置Before launch: External tool
2. 点击+，选择`Run External tool`,再点击+，create tool
3. 设置好Name: golint, 选择运行的Progam: $GOPATH/bin/golint,设置Parameters: ., 设置Working directory: $FileDir$, 勾选上显示信息在console

作者：maiyang
链接：http://www.jianshu.com/p/71bcea150a4b
來源：简书
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。