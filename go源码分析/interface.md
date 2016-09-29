# interface

interface实际上是一个结构体，包括两个成员，一个是指向数据的指针，一个包含了成员的类型信息
```
struct Iface
{
	Itab*	tab;
	void*	data;
};
```