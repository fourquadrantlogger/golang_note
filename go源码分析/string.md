#string

string类型的底层是一个C struct.

成员str为字符数组，len为字符数组长度。golang的字符串是不可变类型，对string类型的变量初始化意味着会对底层结构的初始化

```
struct String
{
        byte*   str;
        intgo   len;
};
```