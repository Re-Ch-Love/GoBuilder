package main

import (
	internalApp "gitee.com/KongchengPro/GoBuilder/internal/app"
	"gitee.com/KongchengPro/GoBuilder/internal/app/commands"
	. "gitee.com/KongchengPro/GoBuilder/pkg/log"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"os"
)

var app *cli.App

//goland:noinspection GoBoolExpressions
func init() {
	// init app
	app = cli.NewApp()
	app.Name = "GoBuilder"
	app.Usage = "build Golang application"
	app.Authors = []*cli.Author{
		{
			Name:  "kongchengpro",
			Email: "kongchengpro@163.com",
		},
	}
	app.Version = "0.0.1"
	app.ExitErrHandler = func(ctx *cli.Context, err error) {
		if err != nil {
			log.WithError(err).Error("error handler")
		}
	}
	app.Flags = []cli.Flag{
		&cli.BoolFlag{
			Name:  "debug-mode",
			Usage: "enable debug mode",
		},
	}
	app.Commands = []*cli.Command{
		{
			Name:  "init",
			Usage: "initialize golang project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				return commands.InitializeProject(ctx.String("path"))
			},
		},
		{
			Name:  "addTask",
			Usage: "initialize golang project",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				return commands.AddTask(ctx.String("path"), ctx.Args().First())
			},
		},
		{
			Name:  "runTask",
			Usage: "run GoBuilder task",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				return commands.RunTask(ctx.String("path"), ctx.Args().First(), ctx.Args().Slice()[1:])
			},
		},
		{
			Name:  "addTaskAndRun",
			Usage: "add GoBuilder task then run it",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				err := commands.AddTask(ctx.String("path"), ctx.Args().First())
				if err != nil {
					return err
				}
				return commands.RunTask(ctx.String("path"), ctx.Args().First(), ctx.Args().Slice()[1:])
			},
		},
		{
			Name:  "addCmd",
			Usage: "run GoBuilder command",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				return commands.AddCommand(ctx.String("path"), ctx.Args().First(), ctx.Args().Slice()[1:])
			},
		},
		{
			Name:  "runCmd",
			Usage: "run GoBuilder command",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:    "path",
					Aliases: []string{"p"},
					Usage:   "specify project path",
					Value:   internalApp.DefaultProjectPath,
				},
			},
			Action: func(ctx *cli.Context) error {
				return commands.RunCommand(ctx.String("path"), ctx.Args().First())
			},
		},
		{
			Name:  "info",
			Usage: "print info",
			Action: func(ctx *cli.Context) error {
				log.WithFields(log.Fields{
					"version": app.Version,
					"authors": app.Authors,
				}).Info("GoBuilder info")
				return nil
			},
		},
	}
	app.Before = func(ctx *cli.Context) error {
		// init logger
		log.SetFormatter(&SimpleFormatter{EnableDebug: ctx.Bool("debug-mode")})
		log.SetOutput(os.Stdout)
		log.SetReportCaller(true)
		if ctx.Bool("debug-mode") {
			log.SetLevel(log.DebugLevel)
			log.WithField("level", "debug").Info("set log level")
		} else {
			log.SetLevel(log.InfoLevel)
		}
		return nil
	}
}

//goland:noinspection GoBoolExpressions
func main() {
	_ = app.Run(os.Args)
}
