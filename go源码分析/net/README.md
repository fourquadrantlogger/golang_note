## net

编程接口
```
func Listen(net, laddr string) (Listener, error)
func (*TCPListener) Accept (c Conn, err error)
func (c *conn) Read(b []byte) (int, error)
func (c *conn) Write(b []byte) (int, error)
```
内部数据结构

Listener 和 TCPListener

简单来说，一个是接口，一个是具体实现。因为golang支持tcp、udp等各种协议，天然使用golang的interface。
Listener定义了针对socket操作的各种接口：
```
type Listener interface {
    Accept() (c Conn, err error)
    Close() error
    Addr() Addr
}
```
TCPListener则定义了tcp协议需要使用的数据结构
```
type TCPListener struct {
    fd *netFD
}
```
netFD是golang网络库最核心数据结构，贯穿了golang网络库的所有API，它对底层的socket进行封装，屏蔽了不同os的网络实现。我们接下来会单独分析该数据结构。
TCPListener中的fd是与监听套接字关联的socket。
Conn和TCPConn

Conn与TCPConn的关系类似Listener与TCPListener。也是抽象与具象的关系。
type TCPConn struct {
    conn
}
type conn struct {
    fd *netFD
}
netFD

这个结构太关键了，说它是golang网络库最核心数据结构一点也不为过。所有golang网络接口最终都会转化为对该结构的方法。
```
// Network file descriptor.
type netFD struct {
    // 不太明白这个mutex的作用
    fdmu fdMutex
    // 该socket相关的fd
    sysfd       int
    family      int
    sotype      int
    isConnected bool
    net         string
    laddr       Addr
    raddr       Addr
    // wait server
    // 与epoll相关结构
    pd pollDesc
}
```
而这个结构中最重要的又是最后一个成员pd，golang网络库实现高并发全靠它了。
```
type pollDesc struct {
    runtimeCtx uintptr
}
```
大致梳理了golang网络库中的几个基本数据结构，接下来我们就可以深入到内部实现流程了。