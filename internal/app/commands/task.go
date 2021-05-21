package commands

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	"gitee.com/KongchengPro/GoBuilder/internal/app"
	"gitee.com/KongchengPro/GoBuilder/pkg/tdk/caller"
	"gitee.com/KongchengPro/GoBuilder/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// AddTask compiles the task source code and moves the executable file to `app.GoBuilderTasksDir`
func AddTask(projectPath, taskName string) (retErr error) {
	defer utils.ReturnErrorFromPanic(&retErr, func(err error) {
		log.WithError(err).Error("an error during add task")
	})
	if !IsInitialized(projectPath) {
		panic(&ProjectNotInitializedError{projectPath})
	}
	if taskName == "" {
		panic(&InvalidTaskNameError{taskName, "it is blank"})
	}
	if _, err := exec.LookPath("go"); err != nil {
		log.WithError(err).Error("cannot find `go` in environment")
		return err
	}
	cmd := exec.Command("go", "version")
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	outputStr := strings.TrimSpace(string(outputBytes))
	log.WithField("go-build-info", outputStr).Info()
	var inputDirPath string
	// if the path not starts with `./`, `go build` command will think that it refers to a package path (should refer to the working directory).
	if projectPath == "./" {
		inputDirPath = "./" + filepath.Join(projectPath, app.GoBuilderTasksDir, taskName)
	} else {
		inputDirPath = filepath.Join(projectPath, app.GoBuilderTasksDir, taskName)
	}
	if !utils.IsExist(inputDirPath) {
		panic(&InvalidTaskNameError{taskName, "it is not exists"})
	}
	outputFilePath := getTaskExecutableFilePath(projectPath, taskName)
	cmd = exec.Command("go", "build", "-o", outputFilePath, inputDirPath)
	log.WithField("cmd", cmd.Args).Debug()
	outputBytes, err = cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	outputStr = strings.TrimSpace(string(outputBytes))
	if outputStr != "" {
		panic(&TaskBuildError{outputStr})
	}
	log.Info(fmt.Sprintf("Task `%s` added successfully.", taskName))
	return nil
}

func RunTask(projectPath, taskName string, taskArgs []string) (retErr error) {
	defer utils.ReturnErrorFromPanic(&retErr, func(err error) {
		log.WithError(err).Error("an error during run task")
	})
	if !IsInitialized(projectPath) {
		panic(&ProjectNotInitializedError{projectPath})
	}
	taskExecutableFilePath := getTaskExecutableFilePath(projectPath, taskName)
	if !utils.IsExist(taskExecutableFilePath) {
		panic(&TaskNotAddedError{taskName})
	}
	log.Debug(caller.MarshalArgs(&caller.TaskCaller{ProjectPath: projectPath}))
	cmd := exec.Command(taskExecutableFilePath, caller.MarshalArgs(&caller.TaskCaller{ProjectPath: projectPath, Args: taskArgs}))
	log.WithField("cmd", cmd.Args).Debug()
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		panic(err)
	}
	reader := bufio.NewReader(stdout)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				panic(err)
			}
		}
		fmt.Println(readString)
	}
	return nil
}

func getTaskExecutableFilePath(projectPath, taskName string) string {
	if runtime.GOOS == "windows" {
		return filepath.Join(projectPath, app.GoBuilderExecutableDir, taskName+".exe")
	} else {
		return filepath.Join(projectPath, app.GoBuilderExecutableDir, taskName)
	}
}
