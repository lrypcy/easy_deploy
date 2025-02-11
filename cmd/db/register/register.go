package register

import "github.com/spf13/cobra"

type void struct{}

var void_element void
var db_cmds = make(map[*cobra.Command]void)

func RegisteredCmd() map[*cobra.Command]void {
	return db_cmds
}

func AddCommand(cmd * cobra.Command) {
	db_cmds[cmd] = void_element
}