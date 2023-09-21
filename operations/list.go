package operations

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"taskScheduler/entity"
	"taskScheduler/file"
	"text/tabwriter"
	"time"
)

func ListTasks(args ListOp) {
	tasks, err := file.LoadFile()
	if err != nil {
		fmt.Println("Opps !", err)
		return
	}

	var filteredTasks []entity.Task

	for _, task := range tasks {
		if taskMatchesFilters(task, args) {
			filteredTasks = append(filteredTasks, task)
		}
	}
	if args.Order == "-desc" {
		sort.SliceStable(filteredTasks, func(i, j int) bool {
			return filteredTasks[i].Id > filteredTasks[j].Id
		})
	} else {
		sort.SliceStable(filteredTasks, func(i, j int) bool {
			return filteredTasks[i].Id < filteredTasks[j].Id
		})
	}

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 8, 0, '\t', 0)
	for _, task := range filteredTasks {
		var endDate string
		var workingTime int
		if !task.EndDate.IsZero() {
			workingTime = CalculateWorkingHours(task.StartDate, task.EndDate)
			endDate = task.EndDate.Format("02-01 15:04:05")
		} else {
			workingTime = CalculateWorkingHours(task.StartDate, time.Now())
			endDate = " "
		}
		startDate := task.StartDate.Format("02-01 15:04:05")
		workingTimeString := fmt.Sprintf("%d hours", workingTime)
		fmt.Fprintf(w, "Id:\t%s\nStart_Date:\t%s\nEnd_Date:\t%s\nWorking_Time:\t%s\n", task.Id, startDate, endDate, workingTimeString)
		fmt.Fprintln(w, "--------------------------")
	}
	w.Flush()
}

func taskMatchesFilters(task entity.Task, args ListOp) bool {
	return (args.Search == "" || strings.HasPrefix(task.Id, args.Search)) &&
		(args.Type == "" ||
			(args.Type == "-s" && task.EndDate.IsZero()) ||
			(args.Type == "-f" && !task.EndDate.IsZero()))
}

// Calculate according to working hours (Mesai)
func CalculateWorkingHours(startDate time.Time, endDate time.Time) int {
	duration := endDate.Sub(startDate)
	workingHours := int(duration.Hours())
	return workingHours
}
