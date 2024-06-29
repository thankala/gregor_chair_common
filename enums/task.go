package enums

type AssemblyTask int

const (
	NoneAssemblyTask AssemblyTask = iota
	AssemblyTask1
	AssemblyTask2
	AssemblyTask3
	AssemblyTask4
	AssemblyTask5
	AssemblyTask6
	AssemblyTask7
	AssemblyTask8
)

func (a AssemblyTask) String() string {
	return [...]string{"NoneAssemblyTask", "AssemblyTask1", "AssemblyTask2", "AssemblyTask3", "AssemblyTask4", "AssemblyTask5", "AssemblyTask6", "AssemblyTask7", "AssemblyTask8"}[a]
}
