package db

import (
	"fmt"

	"github.com/lrypcy/easy_deploy/pkg/mysql"
	"github.com/spf13/cobra"
)

// db.initCmd represents the db.init command
var mysqlCmd = &cobra.Command{
	Use:   "mysql deploy",
	Short: "Used for mysql deploy",
	// 	Long: `A longer description that spans multiple lines and likely contains examples
	// and usage of using your command. For example:

	// Cobra is a CLI library for Go that empowers applications.
	// This application is a tool to generate the needed files
	// to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s, %s\n", start_method, name)
		ctx := mysql.MysqlContext{}
		ctx.Init(mysql.MysqlDeployConfig{
			Start_method: start_method,
		})
		ctx.Start()
		fmt.Println(args)
	},
}

type commonFlag struct {
	flag_type    string
	flag_name    string
	flag_usage   string
	flag_default string
	flag_value   string
	flag_short   string
}

// 添加构造函数设置默认值
func newCommonFlag(flagType, name string) commonFlag {
	return commonFlag{
		flag_type:    flagType,
		flag_name:    name,
		flag_usage:   fmt.Sprintf("Specify %s value", name), // 默认usage
		flag_default: "",                                    // 默认空值
		flag_short:   "",                                    // 默认无短标记
	}
}

var flagInfoMap = map[string]commonFlag{
	"test_string": newCommonFlag("string", "test").withDefault("default_value"),
}

// 添加链式方法设置特定值
func (f commonFlag) withDefault(def string) commonFlag {
	f.flag_default = def
	return f
}

var processTable = map[string]string{
	// "string": ["StringP",]
}

var start_method string
var name string

func addCommonFlag(cmd *cobra.Command) {
	cmd.Flags().StringVar(&start_method, "start_method", "docker", "start mysql in docker")
	cmd.Flags().StringVar(&name, "name", "", "container name")
}

func init() {
	addCommonFlag(mysqlCmd)
	db_cmds[mysqlCmd] = void_element
}
