package operations

import (
	"fmt"
	"taskScheduler/entity"
	"taskScheduler/file"
	"time"
)

func AddTask(args CreateOp) {

	task := entity.Task{
		Id:        args.Id,
		StartDate: time.Now(),
	}

	err := file.AppendFile(task)

	if err != nil {
		fmt.Println(err)
	}
}
