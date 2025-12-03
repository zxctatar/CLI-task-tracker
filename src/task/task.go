package task

type Status int

func (s Status) String() string {
	switch s {
	case 0: return "NotStarted "
	case 1: return "Started"
	case 2: return "Done"
	default: return "Unknown"
	}
}

const (
	NotStarted Status = iota
	Started
	Done
)

type Task struct {
	Id int `json:"id"`
	Title string `json:"title"`
	CreatedTime string `json:"dateTime"`
	LastUpdateTime string `json:"lastUpdate"`
	Stat Status `json:"status"`
}

func NewTask(id int, title string, createdTime string, lastUpdateTime string, stat Status) *Task {
	return &Task{id, title, createdTime, lastUpdateTime, stat}
}

func (t* Task) GetId() int {
	return t.Id
}

func (t* Task) GetTitle() string {
	return t.Title
}

func (t* Task) GetCreatedTime() string {
	return t.CreatedTime
}

func (t* Task) GetLastUpdateTime() string {
	return t.LastUpdateTime
}

func (t* Task) GetStatus() Status {
	return t.Stat
}

func (t* Task) SetTitle(newTitle string) {
	t.Title = newTitle
}

func (t* Task) SetLastUpdateTime(newLastUpdateTime string) {
	t.LastUpdateTime = newLastUpdateTime
}