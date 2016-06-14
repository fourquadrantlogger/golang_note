``` shell
~$ sudo apt-get install git
~$ sudo apt-get install gcc
下载最新版
~$ wget http://www.golangtc.com/static/go/go1.6beta1/go1.6beta1.linux-amd64.tar.gz
用tar 命令来解压压缩包
tar -zxvf
接着我们要添加环境变量
sudo gedit /etc/profile

export GOROOT=/usr/local/go
export GOBIN=$GOROOT/bin
export PATH=$PATH:$GOBIN
export GOPATH=/media/timeloveboy/WinData/GOPATH


更新一下环境变量
source /ect/profile
```