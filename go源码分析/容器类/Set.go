package 容器类



type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Equal(other Set) bool
	Elements() []interface{}
	String() string
}