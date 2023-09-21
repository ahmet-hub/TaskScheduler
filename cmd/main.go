package main

import (
	"fmt"
	"os"
	"taskScheduler/operations"
)

func main() {

	op := operations.GetOperation(os.Args[1:])

	switch args := op.(type) {
	case operations.HelpOp:
		operations.PrintHelp()
	case operations.ListOp:
		operations.ListTasks(args)
	case operations.CreateOp:
		operations.AddTask(args)
		operations.ListTasks(operations.ListOp{Search: args.Id})
	case operations.UpdateOp:
		panic("not implemented")
	case operations.DeleteOp:
		panic("not implemented")
	case operations.UnknownOp:
		panic("not implemented")
	default:
		fmt.Println("asdasasd")
	}

}
