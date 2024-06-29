package enums

type Robot int

const (
	Robot1 Robot = iota
	Robot2
	Robot3
)

func (r Robot) String() string {
	return [...]string{"Robot1", "Robot2", "Robot3"}[r]
}
