package xset

type Set interface {
	Insert(item interface{}) bool
	Contains(item interface{}) bool
	Size() int
}
