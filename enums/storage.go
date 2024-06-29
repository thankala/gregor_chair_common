package enums

type Storage int

const (
	StorageB1 Storage = iota
	StorageB2
	StorageB3
	StorageB4
	StorageB5
	StorageB6L
	StorageB6R
	StorageB7A
	StorageB7B
	StorageB8C
	StorageB8D
	StorageB8E
)

func (s Storage) String() string {
	return [...]string{"StorageB1", "StorageB2", "StorageB3", "StorageB4", "StorageB5", "StorageB6L", "StorageB6R", "StorageB7A", "StorageB7B", "StorageB8C", "StorageB8D", "StorageB8E"}[s]
}
