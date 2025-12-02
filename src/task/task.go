package task

type Status int

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