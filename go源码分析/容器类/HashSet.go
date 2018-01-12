package 容器类

import "bytes"

type HashSet struct {
	items map[interface{}]struct{}
}

func NewHashSet() *HashSet{
	var this *HashSet =&HashSet{
		items:make(map[interface{}]struct{}),
	}
	return this
}
func (this *HashSet)Add(a interface{}) bool{
	_,has:=this.items[a]
	if has{
		return false
	}else {
		this.items[a]=struct{}{}
		return true
	}
}
func (this *HashSet)String(){
	sb:=bytes.Buffer{}
	sb.WriteString()
}