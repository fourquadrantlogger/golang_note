# map

{$go}/src/runtime/hashmap.go

golang的map实现并不是像c++一样使用红黑树，而是使用了hashmap，用数组来实现

golang的map使用了桶的概念，元素是被hash到桶存储，每个桶预设是存储八个kv,而且在头部有一个uint8 tophash[8]的结构，存储每个key的高八位（即hash(key) » (64 - 8)），如果该位置未被放置元素，则有一个特殊的标志Empty

利用了overflow指针，可以无限的增长，类似链表

