package command

import (
	"errors"

	"github.com/nilorg/pkg/logger"

	"github.com/urfave/cli"
)

// Set ...
func Set(ctx *cli.Context) error {
	groupName := ctx.String("group")
	if groupName == "" {
		return errors.New("group name cannot be empty")
	}
	key := ctx.String("key")
	if key == "" {
		return errors.New("key cannot be empty")
	}
	value := ctx.String("value")
	return ManagementClient.SetValue(groupName, key, value)
}

// SetEnv ...
func SetEnv(ctx *cli.Context) error {
	key := ctx.String("key")
	if key == "" {
		return errors.New("key cannot be empty")
	}
	value := ctx.String("value")
	return ManagementClient.SetValue("env", key, value)
}

// DeleteKeys ...
func DeleteKeys(ctx *cli.Context) error {
	keys := ctx.Args()
	if len(keys) == 0 {
		return errors.New("keys cannot be empty")
	}
	for _, value := range keys {
		err := ManagementClient.Delete(value)
		if err != nil {
			logger.Errorf("delete key %s err:%s", value, err)
		}
	}
	return nil
}

//DeleteGroup ...
func DeleteGroup(ctx *cli.Context) error {
	groupName := ctx.Args().First()
	if groupName == "" {
		return errors.New("group cannot be empty")
	}
	keys, err := ManagementClient.GetKeysList(groupName)
	if err != nil {
		return err
	}
	for _, value := range keys {
		err := ManagementClient.Delete(value)
		if err != nil {
			logger.Errorf("delete key %s from %s err:%s", value, groupName, err)
		}
	}
	return nil
}
