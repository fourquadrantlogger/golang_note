# RWMutex读写锁

即是针对于读写操作的互斥锁。它与普通的互斥锁最大的不同就是，它可以分别针对读操作和写操作进行锁定和解锁操作

1、可以随便读。多个goroutin同时读。

2、写的时候，啥都不能干。不能读，也不能写。


 - 读锁定

func (*RWMutex) Lock

func (*RWMutex) Unlock

 - 写锁定

func (*RWMutex) RLock

func (*RWMutex) RUnlock