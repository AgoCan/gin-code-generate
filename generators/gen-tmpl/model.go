package tmpl

// ModelContent model
var ModelContent = `package model
// https://gorm.io/zh_CN/
import (
	"fmt"
	
	// 导入mysql驱动
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	
	"{{ .ProjectName }}/config"
)
// DB db handler
var DB *sqlx.DB

// baseModel 继承使用，减少反复写此段代码
type baseModel struct { 
	ID        int        ` + "`db:" + `"id"` + "json:" + `"id"` + "`" + `
	CreatedAt time.Time  ` + "`db:" + `"created_at"` + "json:" + `"created_at"` + "`" + `
	UpdatedAt time.Time  ` + "`db:" + `"updated_at"` + "json:" + `"updated_at"` + "`" + `
	DeletedAt *time.Time ` + "`db:" + `"deleted_at"` + "json:" + `"deleted_at"` + "`" + `
}

// InitMysql 初始化数据库
func InitMysql() (err error) { 
	DB, err = sqlx.Open("mysql", config.MysqlConnect)
	if err != nil {
		return err
	}

	err = DB.Ping()
	if err != nil {
		return err
	}

	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)
	return nil
}
// Close 关闭数据库
func Close()(){
	err := DB.Close()
	if err != nil {
		fmt.Println(err)
	}
	return
}
`
