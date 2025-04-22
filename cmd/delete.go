/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "github.com/wooblz/todo-list/file"
	//"github.com/wooblz/todo-list/model"
    "strconv"
	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete <task id>",
	Short: "deletes tasks",

	RunE: func(cmd *cobra.Command, args []string) error{
        tasks, err := file.LoadTasks(dataFile)
        if err != nil  {
            return fmt.Errorf("failed to load tasks %w", err)
        }

        var flag bool
        id, err := strconv.Atoi(args[0])
        if err != nil  {
            return fmt.Errorf("failed to convert id to int %w", err)
        }
        for i, v := range tasks  {
            if v.ID == id  {
                flag = true
                tasks = append(tasks[:i],tasks[i+1:]...)
            }
        }
        if !flag  {
            return fmt.Errorf("No task with id: %s", args[0])
        }
        err = file.SaveTasks(dataFile, tasks)
        if err != nil  {
            return fmt.Errorf("failed to save tasks %w", err)
        }

        return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
