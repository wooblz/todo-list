package model_test

import  (
    "testing"
    "time"
    "github.com/wooblz/todo-list/model"
)

func TestTaskCreation(t *testing.T)  {
    t.Run("test time", func(t *testing.T)  {
        start := time.Now()
        want := model.Task{
            ID:          1,
            Name:        "test file",
            CreatedAt:   time.Now(),
            CompletedAt: nil,
        }

        got   := model.NewTask(1, "test file")
        end   := time.Now()

        if got.ID != want.ID || got.Name != want.Name  {
            t.Errorf("Expetced ID: %d and Name: %s, but got ID: %d and Name: %s",want.ID, want.Name, got.ID, got.Name)
        }

        if got.CreatedAt.Before(start) || got.CreatedAt.After(end)  {
            t.Errorf("Expected time between %v and %v, but got %v", start,end,got.CreatedAt)
        }

        if got.CompletedAt != nil  {
            t.Errorf("expected nil completion time, but got %v", got.CompletedAt)
        }  
    })
}
func TestTaskCompletion(t *testing.T)  {
    t.Run("Test Completion", func(t *testing.T)  {
        got := model.Task{
            ID:          1,
            Name:        "test file",
            CreatedAt:   time.Now(),
            CompletedAt: nil,
        }
        start := time.Now()
        got.CompleteTask()
        end := time.Now()

        if got.CompletedAt == nil  {
            t.Errorf("Expected %v, but got nil", got.CompletedAt)
        }

        if got.CompletedAt.Before(start) || got.CompletedAt.After(end)  {
            t.Errorf("Expected time between %v and %v, but got %v", start,end,got.CompletedAt)
        }
    })
}
