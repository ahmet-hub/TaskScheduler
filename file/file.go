package file

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"taskScheduler/entity"
)

var (
	FilePath     = "C:\\Ahmet\\TaskScheduler\\task.json"
	ErrTaskExist = errors.New("task exist")
)

func LoadFile() ([]entity.Task, error) {

	jsonData, err := os.ReadFile(FilePath)
	if err != nil {
		fmt.Println("ReadErr", err)
		return nil, err
	}

	var tasks []entity.Task
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		fmt.Println("ConvertErr", err)
		return nil, err
	}

	return tasks, nil
}

func AppendFile(newTask entity.Task) error {
	jsonData, err := os.ReadFile(FilePath)
	if err != nil {
		fmt.Println("ReadErr", err)
		return err
	}

	var tasks []entity.Task
	err = json.Unmarshal(jsonData, &tasks)
	if err != nil {
		fmt.Println("ConvertErr", err)
		return err
	}

	for _, task := range tasks {
		if task.Id == newTask.Id {
			return ErrTaskExist
		}
	}

	tasks = append(tasks, newTask)

	newJsonData, err := json.Marshal(tasks)
	if err != nil {
		return err
	}

	err = os.WriteFile(FilePath, newJsonData, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
