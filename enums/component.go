package enums

type Component int

const (
	NoneComponent Component = 0
	Legs          Component = 1
	Base          Component = 2
	Castors       Component = 4
	Lift          Component = 8
	Seat          Component = 16
	SeatPlate     Component = 32
	SeatScrews    Component = 64
	Screws        Component = 128
	Back          Component = 256
	LeftArm       Component = 512
	RightArm      Component = 1024
)

func (c *Component) Stage() Stage {
	return Stage(*c)
}
