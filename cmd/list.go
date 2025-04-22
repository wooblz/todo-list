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
    "github.com/mergestat/timediff"
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
                now := timediff.TimeDiff(v.CreatedAt)
                //var done bool
                fmt.Fprintf(w,"%d\t%s\t%s\t",v.ID,v.Name, now)
                if v.CompletedAt != nil  {
                    fmt.Fprintf(w,timediff.TimeDiff(*v.CompletedAt))
                }  else  {
                    fmt.Fprintf(w,"NOT DONE")
                }
                fmt.Fprintf(w,"\n")
            }
            w.Flush()
            return nil
        }
        fmt.Fprintln(w, "ID\tTask\tCreated")
        for _, v := range tasks  {
            now := timediff.TimeDiff(v.CreatedAt)
            fmt.Fprintf(w,"%d\t%s\t%s\n",v.ID, v.Name, now)
        }
        w.Flush()
        return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
    showAll  = listCmd.Flags().BoolP("all", "a", false, "Show all tasks, including completed ones")
}
