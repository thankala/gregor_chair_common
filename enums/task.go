package enums

type Task string

const (
	NoneTask      Task = "NoneTask"
	AssemblyTask1 Task = "AssemblyTask1"
	AssemblyTask2 Task = "AssemblyTask2"
	AssemblyTask3 Task = "AssemblyTask3"
	AssemblyTask4 Task = "AssemblyTask4"
	AssemblyTask5 Task = "AssemblyTask5"
	AssemblyTask6 Task = "AssemblyTask6"
	AssemblyTask7 Task = "AssemblyTask7"
	AssemblyTask8 Task = "AssemblyTask8"
	Orchestrator  Task = "Orchestrator"
)

func (a Task) String() string {
	return string(a)
}
