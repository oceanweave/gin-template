package conf

import (
	"github.com/spf13/viper"
)

var (
	ServerConfig *ServerConf
	MysqlConfig  *MysqlConf
	LoggerConfig *LoggerConf
)

type ServerConf struct {
	Host string
	Port int
	Name string
	Mode string
}

type MysqlConf struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
}

type LoggerConf struct {
	Level       string
	FilePath    string
	FileName    string
	MaxFileSize uint64
	ToFile      bool
}

func LoadConf() error {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	initServerConf()
	initMysqlConf()
	initLogConf()

	return nil
}

func initServerConf() {
	ServerConfig = &ServerConf{
		Host: viper.GetString("server.host"),
		Port: viper.GetInt("server.port"),
		Name: viper.GetString("server.name"),
		Mode: viper.GetString("server.mode"),
	}
}

func initMysqlConf() {
	MysqlConfig = &MysqlConf{
		DataSourceName: viper.GetString("mysql.dataSourceName"),
		MaxOpenConns:   viper.GetInt("mysql.maxOpenConns"),
		MaxIdleConns:   viper.GetInt("mysql.maxIdleConns"),
	}
}

func initLogConf() {
	LoggerConfig = &LoggerConf{
		Level:       viper.GetString("logger.level"),
		FilePath:    viper.GetString("logger.filePath"),
		FileName:    viper.GetString("logger.fileName"),
		MaxFileSize: viper.GetUint64("logger.maxFileSize"),
		ToFile:      viper.GetBool("logger.toFile"),
	}
}
