package models

import (
    "github.com/astaxie/beego"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "gorm.io/gorm/logger"
    "gorm.io/gorm/schema"
    "log"
    "os"
    "time"
)
var (
    DB *gorm.DB
)
func InitMysql() (err error) {
    newLogger := logger.New(
        log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
        logger.Config{
            SlowThreshold: time.Second,   // 慢 SQL 阈值
            LogLevel:      logger.Info, // 日志级别
            IgnoreRecordNotFoundError: true,   // 忽略ErrRecordNotFound（记录未找到）错误
            Colorful:      false,         // 禁用彩色打印
        },
    )
    namingStrategy:= schema.NamingStrategy{
        //TablePrefix: "t_",   // 表名前缀，`User`表为`t_users`
        SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
    }

    user := beego.AppConfig.String("mysqluser")
    pwd := beego.AppConfig.String("mysqlpwd")
    host := beego.AppConfig.String("host")
    port := beego.AppConfig.String("port")
    dbname := beego.AppConfig.String("dbname")

    //dbConn := "root:yu271400@tcp(127.0.0.1:3306)/cmsproject?charset=utf8"
    dsn := user + ":" + pwd + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"

    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
        NamingStrategy: namingStrategy,
        Logger: newLogger,
    })
    if err != nil{
        return err
    }
    DB.AutoMigrate(&Todo{})
    return nil
}

