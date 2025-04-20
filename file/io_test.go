package file

import  ( 
    "testing"
    "os"
    "fmt"
    "time"
    "reflect"
    "encoding/json"
    "github.com/wooblz/todo-list/model"
)

func TestReadFile(t *testing.T)  {
    t.Run("Read File", func(t *testing.T)  {
        tmpFile, err := os.CreateTemp("", "test_tasks_*.json")
        if err != nil {
            t.Fatalf("Could not create file %v",err)
        }
        defer os.Remove(tmpFile.Name())
        
        start := time.Now()
        now := time.Now()
        end := time.Now()
        jsonContent := fmt.Sprintf(`[
            {
                "ID": 1,
                "Name": "Test Task",
                "CreatedAt": "%s",
                "CompletedAt": null
            }
        ]`, now.Format(time.RFC3339Nano))

        if _, err := tmpFile.Write([]byte(jsonContent)); err != nil {
            t.Fatalf("could not write to temp file: %v", err)
        }
        tmpFile.Close()

        tasks, err := LoadTasks(tmpFile.Name())
        if err != nil {
            t.Fatalf("LoadTasks failed: %v", err)
        }

        if len(tasks) != 1 {
            t.Errorf("expected 1 task, got %d", len(tasks))
        }
        if tasks[0].Name != "Test Task" {
            t.Errorf("expected task name 'Test task', got %q", tasks[0].Name)
        }
        if tasks[0].CreatedAt.Before(start) || tasks[0].CreatedAt.After(end)  {
            t.Errorf("expected task between %v and %v, but got %v",start,end,tasks[0].CreatedAt)
        }
        if tasks[0].CompletedAt != nil {
            t.Errorf("expected nil CompletedAt, got %v", tasks[0].CompletedAt)
        }
    })
}
func TestSaveFile(t *testing.T)  {
    t.Run("default", func(t* testing.T)  {
        tmpFile, err := os.CreateTemp("", "test_tasks_*.json")
        if err != nil {
            t.Fatalf("Could not create file %v",err)
        }
        defer os.Remove(tmpFile.Name())
        
        now := time.Now().Round(0)
        tasks := []model.Task{
            {ID:1, Name: "Task 1", CreatedAt: now,},
            {ID:2, Name: "Task 2", CreatedAt: now, CompletedAt: ptr(now),},
        }

        err = SaveTasks(tmpFile.Name(), tasks)
        if err != nil  {
            t.Fatalf("Could not save tasks %v", err)
        }

        data, err := os.ReadFile(tmpFile.Name())
        if err != nil  {
            t.Fatalf("Could not read file %v", err)
        }

        var loaded []model.Task
        if err := json.Unmarshal(data, &loaded); err != nil  {
            t.Fatalf("invalid JSON %v",err)
        }

       // got := json.MarshalIndent
        if !reflect.DeepEqual(loaded, tasks)  {
            t.Errorf("Expected: \n%+v\n Got: \n%+v", loaded, tasks)
        }
    })
}

func ptr(t time.Time) *time.Time  {
    return &t
}


