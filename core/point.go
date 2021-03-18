package core

// 点必须实现的方法
type Point interface {
	Check(str string) (bool, error)
}