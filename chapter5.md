# mgo.bson使用指南

用godoc查看：

手写需要导入

```
import "gopkg.in/mgo.v2/bson" 
```
然后它说这个库是为go语言的一个BSON规范的实现

```
IsObjectIdHex(s string) bool
Marshal(in interface{}) (out []byte, err error)
Now() time.Time 
Unmarshal(in []byte, out interface{}) (err error)

```
bson的基本类型

1.Binary    
2.D BsonDocument
```
bson.D{{"a", 1}, {"b", true}}
并自带Map()方法
func Map() (m M)

```
3.DocElem 
4.M map
```
bson.M{"a": 1, "b": true}
```
5.MongoTimestamp Int64

6.func (id ObjectId) Time() time.Time 
时间返回ID的时间戳部分

7.Raw
原始类型
