package task

type Task struct {
	ID        uint32 `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}
