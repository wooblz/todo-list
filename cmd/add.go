/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "github.com/wooblz/todo-list/file"
	"github.com/wooblz/todo-list/model"
	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [task]",
	Short: "Adds a task ",
    Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error{
        tasks, err := file.LoadTasks(dataFile)
        if err !=nil  {
            return fmt.Errorf("failed to load tasks %w", err)
        }

        var newID int
        for _, v := range tasks  {
            if v.ID > newID  {
                newID = v.ID
            }
        }
        newID++
        newTask := model.NewTask(newID, args[0])
        tasks = append(tasks,newTask)

        err = file.SaveTasks(dataFile, tasks)
        if err != nil  {
            return fmt.Errorf("failed to save tasks %w", err)
        }

        return nil; 
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
