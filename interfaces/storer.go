package interfaces

type Storer interface {
	Ping() error
	Load(string) ([]byte, error)
	Store(string, []byte) error
	Remove(string) error
	AcquireLock(string, string) (bool, error)
	ReleaseLock(string) error
}
