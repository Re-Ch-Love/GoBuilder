package commands

import "errors"

//goland:noinspection GoUnusedParameter
func RunCommand(projectPath, commandName string) (retErr error) {
	/*
		先读取并解析 gobuilder/cmd.gb ，这个文件用 yaml 编写，具体格式如下。解析后查看是否有对应命令，如有，执行，如没有，返回 error 。
		cmd1:
		  - task1
		  - task2
		  - task3
		cmd2:
		  - task3
		  - task2
		  - task1
	*/
	return errors.New("this command has not yet developed")
}

//goland:noinspection GoUnusedParameter
func AddCommand(projectPath, commandName string, tasks []string) error {
	return errors.New("this command has not yet developed")
}
