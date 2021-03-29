package models

import (
	"errors"
	"time"
)

// TaskInfoRepo
type TaskInfoRepo struct {
	tasks []TaskInfo
	//use a map, use the ID as the key...
	taskmap map[string]TaskInfo
}

// NewTaskInfoRepo ...
func NewTaskInfoRepo() *TaskInfoRepo {
	var repo TaskInfoRepo
	repo.taskmap = make(map[string]TaskInfo)

	// add some fake data:
	repo.taskmap["1"] = TaskInfo{
		Name:        "Redis",
		Tag:         "3.2.1",
		State:       "idle",
		ContainerID: "deadbeef007",
		Updated:     time.Now(),
	}

	return &repo
}

// GetTaskInfo ...
func (r *TaskInfoRepo) GetTaskInfo(id string) (*TaskInfo, error) {
	ti, prs := r.taskmap[id]
	if !prs {
		return nil, errors.New("Not found")
	}

	//fmt.Println("Repo: id:", id, ti)
	// set a timer to update the task info
	timer3 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer3.C
		//fmt.Println("timer!")
		ti.State = "stopped"
		ti.NotifyAll()
	}()

	return &ti, nil
}

// GetTaskInfoX ...
func (r *TaskInfoRepo) GetTaskInfoX(id string, cb func(*TaskInfo, error)) {
	//check to see if we have the data locally, otherwise go get it (from?)
	ti, prs := r.taskmap[id]
	if !prs {
		cb(nil, errors.New("Not found"))
	}

	cb(&ti, nil) //go?
}

// what about the situation where the task info is updated?
