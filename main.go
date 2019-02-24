package main

import (
	"log"
	"os"

	"github.com/nilorg/pkg/logger"

	"github.com/cargoboat/cli/client"

	"github.com/cargoboat/cli/command"

	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

// appBefore app 开始
func appBefore(ctx *cli.Context) error {
	// 初始化配置文件
	viper.SetConfigFile("./config.toml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	} else {
		viper.WatchConfig()
	}
	logger.Init()
	command.ManagementClient = client.NewManagementClient(viper.GetString("server.addr"), viper.GetString("server.basic_auth.username"), viper.GetString("server.basic_auth.password"))
	return nil
}

// appAfter app 结束
func appAfter(ctx *cli.Context) error {

	return nil
}

func main() {
	app := cli.NewApp()
	app.Name = "cargoboat-cli"
	app.Usage = ""
	app.Description = "cargoboat管理工具"
	app.Version = "1.0.0"
	app.Author = "德意洋洋"
	app.Copyright = "cargoboat"

	app.Before = appBefore
	app.After = appAfter

	app.Commands = []cli.Command{
		{
			Name:        "list",
			Aliases:     []string{"l"},
			Usage:       "显示所有配置",
			Description: "查询Cargoboat服务器上的所有配置项",
			Subcommands: []cli.Command{
				{
					Name:        "keys",
					Aliases:     []string{"k"},
					Usage:       "显示所有分组或配置项Key",
					Description: "查询Cargoboat服务器上的所有分组或者分组下面的所有Key",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name,n",
							Usage: "group name",
						},
					},
					Action: command.GetAllKeys,
				},
				{
					Name:        "group",
					Aliases:     []string{"g"},
					Usage:       "显示所有分组或配置项、内容",
					Description: "查询Cargoboat服务器上的所有分组或者分组下面的所有数据",
					Flags: []cli.Flag{
						cli.StringFlag{
							Name:  "name,n",
							Usage: "group name",
						},
					},
					Action: command.GetAllGroupList,
				},
				{
					Name:        "env",
					Aliases:     []string{"e"},
					Usage:       "显示所有环境变量",
					Description: "查询Cargoboat服务器上的环境变量和数据",
					Action:      command.GetAllEnvList,
				},
			},
		},
		{
			Name:        "set",
			Aliases:     []string{"s"},
			Usage:       "设置配置",
			Description: "设置Cargoboat服务器配置项",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "group,g",
					Usage: "分组名称",
				},
				cli.StringFlag{
					Name:  "key,k",
					Usage: "Key",
				},
				cli.StringFlag{
					Name:  "value,v",
					Usage: "value",
				},
			},
			Action: command.Set,
		},
		{
			Name:        "set-env",
			Aliases:     []string{"se"},
			Usage:       "设置环境变量",
			Description: "设置Cargoboat服务器全局环境变量",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "key,k",
					Usage: "Key",
				},
				cli.StringFlag{
					Name:  "value,v",
					Usage: "value",
				},
			},
			Action: command.SetEnv,
		},
		{
			Name:        "delete",
			Aliases:     []string{"d"},
			Usage:       "删除配置",
			Description: "删除Cargoboat服务器配置项",
			Subcommands: []cli.Command{
				{
					Name:        "keys",
					Aliases:     []string{"k"},
					Usage:       "删除配置项Key",
					Description: "删除Cargoboat服务器上的Key",
					Action:      command.DeleteKeys,
				},
				{
					Name:        "group",
					Aliases:     []string{"g"},
					Usage:       "删除分组",
					Description: "删除rgoboat服务器上的分组和",
					Action:      command.DeleteGroup,
				},
				{
					Name:        "env",
					Aliases:     []string{"e"},
					Usage:       "显示所有环境变量",
					Description: "查询Cargoboat服务器上的环境变量和数据",
					Action:      command.GetAllEnvList,
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
