package commands

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

//goland:noinspection GoUnusedExportedFunction
func InvokeCommandAndOutputToSTDOUTInRealTime(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		readString, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		fmt.Printf(readString)
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func InvokeCommandAndOutputPerByteToBytesInRealTime(output *[]byte, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		b, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		*output = append(*output, b)
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func InvokeCommandAndOutputPerLineToBytesInRealTime(output *[]byte, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	cmd.Stderr = cmd.Stdout
	if err := cmd.Start(); err != nil {
		return err
	}
	reader := bufio.NewReader(stdout)
	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		}
		*output = append(*output, []byte(str)...)
	}
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func InvokeCommandAndOutputToSTDOUTAtOnce(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	fmt.Print(string(outputBytes))
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func InvokeCommandAndOutputToBytesAtOnce(output *[]byte, name string, args ...string) error {
	cmd := exec.Command(name, args...)
	outputBytes, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	*output = outputBytes
	return nil
}

//goland:noinspection GoUnusedExportedFunction
func ParseArgs(str string) []string {
	var cmd []string
	var buffer bytes.Buffer
	var isInQuotes bool
	for i, r := range str {
		if r == ' ' && !isInQuotes {
			if buffer.Len() != 0 {
				cmd = append(cmd, buffer.String())
				buffer.Reset()
			}
		} else if r == '"' {
			isInQuotes = !isInQuotes
			if buffer.Len() != 0 || str[i+1] == '"' {
				cmd = append(cmd, buffer.String())
				buffer.Reset()
			}
		} else {
			buffer.WriteRune(r)
		}
		if i == len(str)-1 {
			cmd = append(cmd, buffer.String())
			buffer.Reset()
		}
	}
	return cmd
}

//goland:noinspection GoUnusedExportedFunction
func ParseCommand(str string) (string, []string) {
	s := strings.SplitN(str, " ", 2)
	if len(s) == 1 {
		return str, []string{}
	}
	cmd, args := s[0], s[1]
	return cmd, ParseArgs(args)
}
