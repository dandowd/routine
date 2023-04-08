package infrastructure

import "routine/common"

type DynamoDbRepo[T interface{}] struct {
}

func (*DynamoDbRepo[T]) Insert(entity T) {

}

func (*DynamoDbRepo[T]) Get(id any) T {
	panic("Not yet implemented")
}

func NewDynamoDbRepo[T interface{}]() common.Repository[T] {
	return &DynamoDbRepo[T]{}
}
