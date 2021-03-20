package main

import (
	"bufio"
	"fmt"
	. "gitee.com/KongchengPro/GoBuilder/pkg/log"
	caller2 "gitee.com/KongchengPro/GoBuilder/pkg/tdk/caller"
	"gitee.com/KongchengPro/GoBuilder/pkg/tdk/commands"
	"gitee.com/KongchengPro/GoBuilder/pkg/utils"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var caller *caller2.TaskCaller

func init() {
	caller = caller2.UnmarshalArgs(os.Args[1])
	log.SetLevel(log.InfoLevel)
	log.SetFormatter(&SimpleFormatter{})
	log.SetOutput(os.Stdout)
	log.SetReportCaller(true)
	if len(caller.Args) < 3 {
		log.WithFields(log.Fields{
			"arg1": "name",
			"arg2": "version",
			"arg3": "mainFilePath",
		}).Fatal("you should provide the arguments")
	}
}

func main() {
	name := caller.Args[0]
	version := caller.Args[1]
	inputFilePath := caller.Args[2]
	if !utils.IsExist(inputFilePath) {
		log.WithField("inputFilePath", inputFilePath).Fatal("not exist")
	}
	var outputBytes []byte
	err := commands.InvokeCommandAndOutputToBytesAtOnce(&outputBytes, "go", "tool", "dist", "list")
	if err != nil {
		log.WithError(err).Fatal()
	}
	outputStr := strings.TrimSpace(string(outputBytes))
	dists := strings.Split(outputStr, "\n")
	log.Infof("total %d dists\n", len(dists))
	var doneCount int
	var lock sync.Mutex
	for _, dist := range dists {
		tmp := strings.SplitN(dist, "/", 2)
		goos, goarch := tmp[0], tmp[1]
		log.WithFields(log.Fields{"GOOS": goos, "GOARCH": goarch}).Info("start build")
		outputDir := filepath.Join(caller.ProjectPath, "bin", version)
		var outputFileName string
		if goos == "windows" {
			outputFileName = fmt.Sprintf("%s-%s_%s_%s.exe", name, version, goos, goarch)
		} else {
			outputFileName = fmt.Sprintf("%s-%s_%s_%s", name, version, goos, goarch)
		}
		if !utils.IsExist(outputDir) {
			err = os.MkdirAll(outputDir, 777)
			if err != nil {
				log.WithError(err).Fatal()
			}
		}
		outputFilePath := filepath.Join(outputDir, outputFileName)
		go func() {
			defer func() {
				lock.Lock()
				doneCount++
				lock.Unlock()
			}()
			build(inputFilePath, outputFilePath, goos, goarch)
			if utils.IsExist(outputFilePath) {
				upx(outputFilePath)
			}
		}()
	}
	for {
		if doneCount == len(dists) {
			log.Info("all done")
			break
		}
		time.Sleep(time.Second)
	}
}

func upx(outputFilePath string) {
	cmd := exec.Command("upx", "-9", outputFilePath)
	stdout, err := cmd.StdoutPipe()
	defer func() {
		if err := stdout.Close(); err != nil {
			log.WithError(err).Fatal()
		}
	}()
	if err != nil {
		log.WithError(err).Fatal("error during build")
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		log.WithError(err).Fatal("error during build")
	}
	reader := bufio.NewReader(stdout)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.WithError(err).Fatal("error during upx")
			}
		}
		fmt.Printf(readString)
	}
}

func build(inputFilePath, outputFilePath, goos, goarch string) {
	logEntry := log.WithFields(log.Fields{"GOOS": goos, "GOARCH": goarch})
	cmd := exec.Command("go", "build", "-ldflags", "-s -w", "-o", outputFilePath, inputFilePath)
	logEntry.WithField("cmd", cmd.Args).Debug("run cmd")
	cmd.Env = append(cmd.Env, "GOOS="+goos)
	cmd.Env = append(cmd.Env, "GOARCH="+goarch)
	cmd.Env = append(cmd.Env, "CGO_ENABLED=0")
	cmd.Env = append(cmd.Env, "GOCACHE="+getGoEnv("GOCACHE", logEntry))
	cmd.Env = append(cmd.Env, "GOPROXY="+getGoEnv("GOPROXY", logEntry))
	cmd.Env = append(cmd.Env, "GOPATH="+getGoEnv("GOPATH", logEntry))
	stdout, err := cmd.StdoutPipe()
	defer func() {
		if err := stdout.Close(); err != nil {
			log.WithError(err).Fatal()
		}
	}()
	if err != nil {
		logEntry.WithError(err).Fatal("error during build")
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		logEntry.WithError(err).Fatal("error during build")
	}
	reader := bufio.NewReader(stdout)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				logEntry.WithError(err).Fatal("error during build")
			}
		}
		fmt.Printf("(%s, %s) %s\n", goos, goarch, readString)
	}
}

func getGoEnv(key string, logEntry *log.Entry) string {
	var outputBytes []byte
	err := commands.InvokeCommandAndOutputToBytesAtOnce(&outputBytes, "go", "env", "get", key)
	if err != nil {
		logEntry.WithError(err).Fatal()
	}
	return strings.TrimSpace(string(outputBytes))
}
