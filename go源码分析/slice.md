# slice

slice类型的底层同样是一个C struct
```
struct	Slice
{				// must not move anything
	byte*	array;		// actual data
	uintgo	len;		// number of elements
	uintgo	cap;		// allocated number of elements
};
```

slice1和slice2可以共用一个底层数组