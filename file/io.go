package file

import (
	"encoding/json"
    "os"
	"github.com/wooblz/todo-list/model"
)

func LoadTasks(fileName string) ([]model.Task, error)  {
    file,err := os.ReadFile(fileName)
    if err != nil  {
        return nil, err
    }

    var tasks []model.Task
    err = json.Unmarshal(file, &tasks)
    if err != nil  {
        return nil, err
    }

    return tasks, nil
}
func SaveTasks(fileName string, tasks []model.Task) error  {
    jsonData, err := json.Marshal(tasks)
    if err != nil  {
        return err
    }

    err = os.WriteFile(fileName, jsonData, 0644)
    if err != nil  {
        return err
    }

    return nil
}
