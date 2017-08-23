1. 点击Edit Configuration, 添加一个 `Go Application`，设置Before launch: External tool
2. 点击+，选择`Run External tool`,再点击+，create tool
3. 设置好Name: golint 选择运行的Progam: $GOPATH/bin/golint,设置Parameters: ., 设置Working directory: $FileDir$, 勾选上显示信息在console
 
