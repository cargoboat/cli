package command

import (
	"github.com/cargoboat/cli/client"
	"github.com/nilorg/pkg/logger"
	"github.com/urfave/cli"
)

var ManagementClient *client.ManagementClient

// GetAllGroupList ...
func GetAllGroupList(ctx *cli.Context) error {
	groupName := ctx.String("name")
	list, err := ManagementClient.GetConfigList(groupName)
	if err != nil {
		return err
	}
	if groupName == "" {
		logger.Infoln("-----all config-----")
	} else {
		logger.Infof("-----group %s configs-----", groupName)
	}
	for key, value := range list {
		logger.Infof("%s:%s", key, value)
	}
	return nil
}

// GetAllKeys ...
func GetAllKeys(ctx *cli.Context) error {
	groupName := ctx.String("name")
	keys, err := ManagementClient.GetKeysList(groupName)
	if err != nil {
		return err
	}
	if groupName == "" {
		logger.Infoln("-----all key-----")
	} else {
		logger.Infof("-----group %s keys-----", groupName)
	}
	for _, value := range keys {
		logger.Infoln(value)
	}
	return nil
}

// GetAllEnvList ...
func GetAllEnvList(_ *cli.Context) error {
	list, err := ManagementClient.GetConfigList("env")
	if err != nil {
		return err
	}
	logger.Infoln("-----all env-----")
	for key, value := range list {
		logger.Infof("%s:%s", key, value)
	}
	return nil
}
