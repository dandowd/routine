package common

type Repository[T interface{}] interface {
	Insert(entity T) (*T, *RepositoryError)
	Get(id string) (*T, *RepositoryError)
}

type CollectionRepository[T interface{}] interface {
	Repository[T]
	GetPage(startId string, limit int) (*[]*T, *RepositoryError)
}

type RepositoryErrorType string

const (
	NotFound      RepositoryErrorType = "NotFound"
	DatabaseError RepositoryErrorType = "DatabaseError"
)

type RepositoryError struct {
	errorType RepositoryErrorType
	message   string
}

func (e *RepositoryError) Error() string {
	return e.message
}

func (e *RepositoryError) Type() RepositoryErrorType {
	return e.errorType
}

func NewRepositoryError(errorType RepositoryErrorType, message string) *RepositoryError {
	return &RepositoryError{errorType, message}
}

type DbTable[T interface{}] interface {
	Get(id string) (*T, error)
	Put(entity T) (*T, error)
	GetPage(limit int, startId *string) (*[]*T, error)
}
