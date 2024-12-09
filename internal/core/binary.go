package core

// BinaryStorager - хранилище бинарных файлов
type BinaryStorager interface {
	AddBinaryByUser(*Binary) error
	SelectBinaryByUser(*User) ([]Binary, error)
	UpdateBinaryByUser(*Binary, *User) error
	DeleteBinaryByUser(*Binary, *User) error
}

// Binary - произвольный бинарный файл
type Binary struct {
	Value []byte
	Meta  string
	ID    int
}
