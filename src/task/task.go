package task

type Status int

func (s Status) String() string {
	switch s {
	case 0:
		return "Not started "
	case 1:
		return "Started"
	case 2:
		return "Done"
	default:
		return "Unknown"
	}
}

const (
	NotStarted Status = iota
	Started
	Done
)

type Task struct {
	Id             int    `json:"id"`
	Description    string `json:"title"`
	CreatedTime    string `json:"dateTime"`
	LastUpdateTime string `json:"lastUpdate"`
	Stat           Status `json:"status"`
}

func NewTask(id int, title string, createdTime string, lastUpdateTime string, stat Status) *Task {
	return &Task{id, title, createdTime, lastUpdateTime, stat}
}

func (t *Task) GetId() int {
	return t.Id
}

func (t *Task) GetDescription() string {
	return t.Description
}

func (t *Task) GetCreatedTime() string {
	return t.CreatedTime
}

func (t *Task) GetLastUpdateTime() string {
	return t.LastUpdateTime
}

func (t *Task) GetStatus() Status {
	return t.Stat
}

func (t *Task) SetDescription(newDescription string) {
	t.Description = newDescription
}

func (t *Task) SetLastUpdateTime(newLastUpdateTime string) {
	t.LastUpdateTime = newLastUpdateTime
}

func (t *Task) StartTask() {
	t.Stat = Started
}

func (t *Task) DoneTask() {
	t.Stat = Done
}

func (t *Task) CancelTask() {
	t.Stat = NotStarted
}
