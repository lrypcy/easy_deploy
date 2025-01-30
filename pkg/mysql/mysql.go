package mysql

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

type containerInfo struct {
	container_name string
	container_id   string
}
type MysqlDeployConfig struct {
	Start_method string
}
type MysqlContext struct {
	config         MysqlDeployConfig
	pid            int
	container_info containerInfo
	cmd            *exec.Cmd
}

func (mysql_ctx *MysqlContext) Init(cfg MysqlDeployConfig) error {
	mysql_ctx.config = cfg
	return nil
}

func (mysql_ctx *MysqlContext) Start() error {
	var start_cmd []string
	if mysql_ctx.config.Start_method == "docker" {
		cmd_str := "docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=Congyuan#678 -d mysql:latest"
		start_cmd = strings.Split(cmd_str, " ")
	}

	cmd := exec.Command(start_cmd[0], start_cmd[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	mysql_ctx.cmd = cmd

	fmt.Printf("start to execute: %s\n", cmd.String())
	err := cmd.Run()
	if err != nil {
		fmt.Printf("Fail to execute cmd")
		outStr, errStr := string(stdout.Bytes()), string(stderr.Bytes())
		fmt.Printf("out:\n%s\nerr:\n%s\n", outStr, errStr)
	} else {
		fmt.Printf("execute success!!!")
	}

	return err
}

func (mysql_ctx *MysqlContext) Remove() error {
	return nil
}
