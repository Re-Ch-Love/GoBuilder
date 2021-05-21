package caller

import (
	"encoding/json"
	"os"
)

var CurrentTaskCaller *TaskCaller

func init() {
	CurrentTaskCaller = UnmarshalArgs(os.Args[1])
}

// TaskCaller contains the arguments which pass to task.
// Using struct can make it more convenient to subsequent update.
// --------------------------------------------------
// The name of this struct may not be appropriate,
// but I really can't think of a better name to describe it.
// :(
// If you have a better name, please submit a pull request, thanks a lot.
type TaskCaller struct {
	ProjectPath string
	Args        []string
}

func UnmarshalArgs(jsonStr string) *TaskCaller {
	// There should be no error here, but TDK should make sure that it is no bugs, so panic it.
	t := &TaskCaller{}
	err := json.Unmarshal([]byte(jsonStr), t)
	if err != nil {
		panic(err)
	}
	return t
}

func MarshalArgs(args *TaskCaller) string {
	// There should be no error here, but TDK should make sure that it is no bugs, so panic it.
	bytes, err := json.Marshal(args)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
