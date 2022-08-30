package golibs

import (
	"log"
	"sync"
)

var taskManagers *TaskManager

type Task struct {
	Name     string
	TaskFunc func(t *TaskManager, taskName string)
	Msg      chan string
	Stoped   chan bool
}

type TaskManager struct {
	Tasks map[string]*Task
	Wg    *sync.WaitGroup
}

func NewTaskManagerInit() *TaskManager {
	taskManagers = new(TaskManager)
	taskManagers.Tasks = make(map[string]*Task)
	taskManagers.Wg = new(sync.WaitGroup)
	return taskManagers
}

func GetTaskManagers() *TaskManager {
	if taskManagers != nil {
		return taskManagers
	}
	return NewTaskManagerInit()
}

func (t *TaskManager) AddTask(taskName string, taskFunc func(t *TaskManager, taskName string)) {
	m := new(Task)
	m.Name = taskName
	m.TaskFunc = taskFunc
	m.Msg = make(chan string, 20)
	m.Stoped = make(chan bool, 1)
	t.Tasks[taskName] = m
}

func (t *TaskManager) Start() {
	for taskName, task := range t.Tasks {
		t.Wg.Add(1)
		go task.TaskFunc(t, taskName)
	}
}

func (t *TaskManager) Stop() {
	for _, task := range t.Tasks {
		log.Println(task.Name, "stop")
		task.Stoped <- true
	}
}
