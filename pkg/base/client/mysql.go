package client

import (
	"MengGoods/pkg/constants"
	"MengGoods/pkg/merror"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
)

func NewMySQLClient(dbName string) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		viper.GetString("mysql.user"),
		viper.GetString("mysql.password"),
		viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"),
		dbName,
		viper.GetString("mysql.charset"),
	)
	// 连接数据库
	db, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			PrepareStmt:            true,  // 在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
			SkipDefaultTransaction: false, // 不禁用默认事务(即单个创建、更新、删除时使用事务)
			TranslateError:         true,  // 允许翻译错误
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名
			},
		})
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql connect error: %v", err))
	}
	if err := db.Use(tracing.NewPlugin()); err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql connect error: %v", err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql connect error: %v", err))
	}
	sqlDB.SetMaxIdleConns(constants.MaxIdleConns)
	sqlDB.SetMaxOpenConns(constants.MaxConnections)
	sqlDB.SetConnMaxLifetime(constants.ConnMaxLifetime)
	sqlDB.SetConnMaxIdleTime(constants.ConnMaxIdleTime)

	//进行测试连接
	if err := sqlDB.Ping(); err != nil {
		return nil, merror.NewMerror(merror.InternalDatabaseErrorCode, fmt.Sprintf("dal.InitMySQL mysql ping error: %v", err))
	}
	return db, nil
}
