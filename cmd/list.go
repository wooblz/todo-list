/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
    "text/tabwriter"
    "os"
    "github.com/wooblz/todo-list/file"
	//"github.com/wooblz/todo-list/model"
	"github.com/spf13/cobra"
)

var showAll *bool;

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
        tasks, err := file.LoadTasks(dataFile)  
        if err != nil  {
            return fmt.Errorf("failed to load tasks %w", err)
        }
        w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
        if *showAll  {
            fmt.Fprintln(w, "ID\tTask\tCreated At\tCompleted At")
            for _, v := range tasks  {
                now := v.CreatedAt
                var done bool
                if v.CompletedAt != nil  {
                    done = true
                }
                fmt.Fprintln(w,"%d\t%s\t%s\t%t",now.Format("2006-01-02 15:04:0"),done)
            }
            return nil
        }
        fmt.Fprintln(w, "ID\tTask\tCreated")
        for _, v := range tasks  {
            now := v.CreatedAt
            fmt.Fprintln(w,"%d\t%s\t%s",now.Format("2006-01-02 15:04:0"))
        }
        return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    showAll  = listCmd.Flags().BoolP("all", "a", false, "Show all tasks, including completed ones")
}
