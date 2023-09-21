package operations

import (
	"fmt"
	"os"
)

func PrintHelp() {
	help := `USAGE:
	ls : list task
	ls {TaskId} filter for task ids
	ls {-s} filter starting tasks
	ls {asc/desc} filter create date asc/desc
	`
	fmt.Fprintf(os.Stdout, "%s\n", help)
}
