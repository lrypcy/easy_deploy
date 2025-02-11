package postgresql

import (
	"fmt"

	"github.com/lrypcy/easy_deploy/cmd/db/register"
	"github.com/lrypcy/easy_deploy/internal/db/postgresql"
	"github.com/spf13/cobra"
)

// db.initCmd represents the db.init command
var postgresqlCmd = &cobra.Command{
	Use:   "postgresql deploy",
	Short: "Used for postgresql deploy",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s\n", start_method, name)
		ctx := postgresql.PostgresqlContext{}
		ctx.Init(postgresql.PostgresqlDeployConfig{
			Start_method: start_method,
		})
		ctx.Start()
		fmt.Println(args)
	},
}

var start_method string
var name string

func addCommonFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&start_method, "start_method", "docker", "start mysql in docker")
	cmd.Flags().StringVar(&name, "name", "", "container name")
}

func init() {
	addCommonFlag(postgresqlCmd)
	register.AddCommand(postgresqlCmd)
}