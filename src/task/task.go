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
	DateTime string `json:"dateTime"`
	Stat Status `json:"status"`
}

func NewTask(id int, title string, dateTime string, stat Status) *Task {
	return &Task{id, title, dateTime, stat}
}

func (t* Task) GetId() int {
	return t.Id
}

func (t* Task) GetTitle() string {
	return t.Title
}

func (t* Task) GetDateTime() string {
	return t.DateTime
}

func (t* Task) GetStatus() Status {
	return t.Stat
}