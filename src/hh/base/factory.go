package base

type Factory interface {
	New()(interface{},error)
}