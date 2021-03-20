package app

import "os"

const (
	GoBuilderDataDir                           = "gobuilder"
	GoBuilderExecutableDir                     = "gobuilder/.executable"
	GoBuilderExecutableDirKeepFile             = "gobuilder/.executable/.keep"
	GoBuilderTasksDir                          = "gobuilder/tasks"
	GoBuilderTasksDirKeepFile                  = "gobuilder/tasks/.keep"
	GoBuilderCommandFile                       = "gobuilder/cmd.yml"
	DefaultProjectPath                         = "./"
	DefaultPerm                    os.FileMode = 777
)
