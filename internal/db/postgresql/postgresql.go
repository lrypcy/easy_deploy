package postgresql

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
type PostgresqlDeployConfig struct {
	Start_method string
}
type PostgresqlContext struct {
	config         PostgresqlDeployConfig
	pid            int
	container_info containerInfo
	cmd            *exec.Cmd
}

func (postgresql_ctx *PostgresqlContext) Init(cfg PostgresqlDeployConfig) error {
	postgresql_ctx.config = cfg
	return nil
}

func (postgresql_ctx *PostgresqlContext) Start() error {
	var start_cmd []string
	if postgresql_ctx.config.Start_method == "docker" {
		cmd_str := "docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -d postgres"
		start_cmd = strings.Split(cmd_str, " ")
	}

	cmd := exec.Command(start_cmd[0], start_cmd[1:]...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	postgresql_ctx.cmd = cmd

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

func (postgresql_ctx *PostgresqlContext) Remove() error {
	return nil
}
