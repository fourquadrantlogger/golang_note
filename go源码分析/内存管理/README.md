## Stack内存管理

介绍

要启动go协程，必须要为其准备栈空间，用来存储函数内变量、执行函数调用等。在go协程退出后，其占用的stack内存也会随之一同释放。
实际应用中协程的启动、退出可能会比较频繁，runtime必须要做点什么来保证启动、销毁协程的代价尽量小。而申请、释放stack空间所需内存则是一个比较大的开销，因此，go runtime采用了stack cache pool来缓存一定数量的stack memory。申请时从stack cache pool分配，释放时首先归还给stack cache pool。
主要思想

stack cache pool的主要思想是按照固定大小划分成多级：每级别其实是一个空闲stack队列：最小的stack size是_FixedStack=2KB。级别之间的大小按照2倍的关系递增。
同时为了效率考虑，除了全局的stack cache pool外，每个线程m还有自己的stack cache。这样，线程在分配的时候可以优先从自己的stack cache中查找，提高效率。
初始时，stack cache pool为空：当有stack分配需求时，首先向os申请一块较大的内存，将其按照该级别的stack大小切割后放在空闲stack列表上，然后再从该空闲stack列表上分配。主要结构如下：


## 核心数据结构

Go内存管理模块的核心数据结构比较少:
mheap：管理全局的从os申请的虚拟内存空间；
mspan：将mheap按照固定大小切分而成的细粒度的内存区块，每个区块映射了虚拟内存中的若干连续页面，页大小由Go内部定义；
mcache：与线程相关缓存，该结构的存在是为了减少内存分配时的锁操作，优化内存分配性能。
mcentral：集中内存池，线程在本地分配失败后，尝试向mcentral申请，如果mcentral也没有资源，则尝试向mheap分配。