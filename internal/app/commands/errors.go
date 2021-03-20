package commands

import (
	"fmt"
)

type InvalidTaskNameError struct {
	TaskName string
	Reason   string
}

func (e *InvalidTaskNameError) Error() string {
	return fmt.Sprintf("Task name `%s` is invalid since %s.", e.TaskName, e.Reason)
}

type TaskBuildError struct {
	Output string
}

func (e *TaskBuildError) Error() string {
	return e.Output
}

type ProjectHasBeenInitializedError struct {
	ProjectPath string
}

func (e *ProjectHasBeenInitializedError) Error() string {
	return fmt.Sprintf("Project in path `%s` has already been initialized.", e.ProjectPath)
}

type TaskNotAddedError struct {
	TaskName string
}

func (e *TaskNotAddedError) Error() string {
	return fmt.Sprintf("Task `%s` was not added, you should add before run.", e.TaskName)
}

type ProjectNotInitializedError struct {
	ProjectPath string
}

func (e *ProjectNotInitializedError) Error() string {
	return fmt.Sprintf("Project in path `%s` did not initialized.", e.ProjectPath)
}
