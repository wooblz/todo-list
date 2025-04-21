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
	RunE: func(cmd *cobra.Command, args []string) Error{

        newTask := NewTask
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
