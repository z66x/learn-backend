package main

import (
	"encoding/json"
	"os"
)

const dataFile = "data.json"
func LoadTasks() ([]*Task, error) {
	data, err := os.ReadFile(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []*Task{}, nil
		}
		return nil, err
	}
	var tasks []*Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, err
}

func SaveTasks(tasks []*Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	tmpFile := dataFile + ".tmp"
	err = os.WriteFile(tmpFile, data, 0644)
	if err != nil {
		return err
	}
	return os.Rename(tmpFile, dataFile)
}
