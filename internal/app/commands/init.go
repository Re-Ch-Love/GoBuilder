package commands

import (
	"gitee.com/KongchengPro/GoBuilder/internal/app"
	"gitee.com/KongchengPro/GoBuilder/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// InitializeProject initializes a project
func InitializeProject(projectPath string) (retErr error) {
	defer utils.ReturnErrorFromPanic(&retErr, func(err error) {
		log.WithError(err).Error("an error during init project")
	})
	if IsInitialized(projectPath) {
		panic(&ProjectHasBeenInitializedError{projectPath})
	}
	log.WithField("projectPath", projectPath).Debug("init project")
	utils.MustMkDirAll(projectPath, app.GoBuilderDataDir)
	utils.MustMkDirAll(projectPath, app.GoBuilderExecutableDir)
	utils.MustMkFile(projectPath, app.GoBuilderExecutableDirKeepFile)
	utils.MustMkDirAll(projectPath, app.GoBuilderTasksDir)
	utils.MustMkFile(projectPath, app.GoBuilderTasksDirKeepFile)
	utils.MustMkFile(projectPath, app.GoBuilderCommandFile)
	log.Info("init successfully")
	return nil
}

func IsInitialized(projectPath string) bool {
	return utils.IsExist(projectPath+app.GoBuilderDataDir) &&
		utils.IsExist(projectPath+app.GoBuilderExecutableDir) &&
		utils.IsExist(projectPath+app.GoBuilderExecutableDirKeepFile) &&
		utils.IsExist(projectPath+app.GoBuilderTasksDir) &&
		utils.IsExist(projectPath+app.GoBuilderTasksDirKeepFile) &&
		utils.IsExist(projectPath+app.GoBuilderCommandFile)
}
