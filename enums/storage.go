package enums

type Storage string

const (
	NoneStorage Storage = "NoneStorage"
	StorageB1   Storage = "StorageB1"
	StorageB2   Storage = "StorageB2"
	StorageB3   Storage = "StorageB3"
	StorageB4   Storage = "StorageB4"
	StorageB5   Storage = "StorageB5"
	StorageB6L  Storage = "StorageB6L"
	StorageB6R  Storage = "StorageB6R"
	StorageB7A  Storage = "StorageB7A"
	StorageB7B  Storage = "StorageB7B"
	StorageB8C  Storage = "StorageB8C"
	StorageB8D  Storage = "StorageB8D"
	StorageB8E  Storage = "StorageB8E"
)

func (s Storage) String() string {
	return string(s)
}

func (s Storage) StringShort() string {
	switch s {
	case StorageB1:
		return "B1"
	case StorageB2:
		return "B2"
	case StorageB3:
		return "B3"
	case StorageB4:
		return "B4"
	case StorageB5:
		return "B5"
	case StorageB6L:
		return "B6L"
	case StorageB6R:
		return "B6R"
	case StorageB7A:
		return "B7A"
	case StorageB7B:
		return "B7B"
	case StorageB8C:
		return "B8C"
	case StorageB8D:
		return "B8D"
	case StorageB8E:
		return "B8E"
	default:
		return "NoneStorage"
	}
}
