package file

import  ( 
    "testing"
    "os"
    "fmt"
    "time"
)

func ReadFileTest(t *testing.T)  {
    t.Run("Read File", func(t *testing.T)  {
        tmpFile, err := os.CreateTemp("", "test_tasks_*.json")
        if err != nil {
            t.Fatalf("Could not create file %v",err)
        }
        defer os.Remove(tmpFil.Name)
        
        start := time.Now()
        now := time.Now()
        end := time.Now()
        jsonContent := fmt.Sprintf(`[
            {
                "ID": 1,
                "Name": Test Task,
                "CreatedAt"; %s
                "CompletedAt": null
            }
        ]`, now)

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
        if tasks[0].Name != "Test task" {
            t.Errorf("expected task name 'Test task', got %q", tasks[0].Name)
        }
        if tasks[0].CreatedAt.Before(start) || tasks[0].CreatedAt.Before(end)  {
            t.Errorf("expected task between %v and %v, but got %v",start,end,tasks[0].CreatedAt)
        }
        if tasks[0].CompletedAt != nil {
            t.Errorf("expected nil CompletedAt, got %v", tasks[0].CompletedAt)
        }
    })
}


