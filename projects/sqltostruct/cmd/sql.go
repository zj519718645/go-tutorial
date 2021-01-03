package cmd

import (
	"log"

	"github.com/zj519718645/go-tutorial/projects/sqltostruct/internal/sql2struct"
	"github.com/spf13/cobra"
)

var username string
var password string
var host string
var charset string
var dbType string
var dbName string
var tableName string

var sqlCmd = &cobra.Command{
	Use:   "sql",
	Short: "sql转换和处理",
	Long:  "sql转换和处理",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var sql2structCmd = &cobra.Command{
	Use:   "struct",
	Short: "sql转换",
	Long:  "sql转换",
	Run: func(cmd *cobra.Command, args []string) {
		dbInfo := &sql2struct.DBInfo{
			DBType:   dbType,
			Host:     host,
			UserName: username,
			Password: password,
			Charset:  charset,
		}
		dbModel := sql2struct.NewDBModel(dbInfo)
		err := dbModel.Connect()
		if err != nil {
			log.Fatalf("dbModel.Connect err: %v", err)
		}
		columns, err := dbModel.GetColumns(dbName, tableName)
		if err != nil {
			log.Fatalf("dbModel.GetColumn err :%v", err)
		}
		tpl := sql2struct.NewStructTemplate()
		tplColumns := tpl.AssemblyColumns(columns)
		err = tpl.Generate(tableName, tplColumns)
		if err != nil {
			log.Fatalf("tpl.Generate err: %v", err)
		}
	},
}

func init() {
	sqlCmd.AddCommand(sql2structCmd)
	sql2structCmd.Flags().StringVarP(&username, "username", "u", "", "请输入用户名")
	sql2structCmd.Flags().StringVarP(&password, "password", "p", "", "请输入用户密码")
	sql2structCmd.Flags().StringVarP(&host, "host", "", "127.0.0.1:3306", "请输入主机地址")
	sql2structCmd.Flags().StringVarP(&charset, "charset", "", "utf8mb4", "请输入数据库编码")
	sql2structCmd.Flags().StringVarP(&dbType, "type", "", "mysql", "请输入数据库实例类型")
	sql2structCmd.Flags().StringVarP(&dbName, "db", "d", "", "请输入数据库名称")
	sql2structCmd.Flags().StringVarP(&tableName, "table", "t", "", "请输入表名")
}
