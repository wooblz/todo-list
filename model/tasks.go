package model

import  (
    "time"
)

type Task struct  {
    ID int
    Name string
    CreatedAt time.Time
    CompletedAt *time.Time
}

func NewTask(id int, name string) Task  {
    return Task  {
        ID: id,
        Name: name,
        CreatedAt: time.Now(),
        CompletedAt: nil,
    }
}

func (t *Task) CompleteTask()  {
    now := time.Now()
    t.CompletedAt = &now
} 
func (t *Task) IsComplete() bool {
    return t.CompletedAt != nil
}
