package common

type Repository[T interface{}] interface {
	Insert(entity T)
	Get(id any) T
}
