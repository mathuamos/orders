package common

import (
	"database/sql"
	"log"
	"github.com/spf13/viper"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Configutils int configs
func Configutils() {
	viper.SetDefault("host", "localhost")

	viper.SetConfigName("config")
	log.Println("running on development configs")

	viper.AddConfigPath("/Users/amosmathu/lenstech/orders/config/")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Println("config error should be logged ", err)
	}
	log.Println("viper initiated successfully ")
}

// Logutils -- logging
func Logutils() {
	log.SetOutput(&lumberjack.Logger{
		Filename:   viper.GetString("system_logging.log_url"),
		MaxSize:    viper.GetInt("system_logging.MaxSize"),
		MaxBackups: viper.GetInt("system_logging.MaxBackups"),
		MaxAge:     viper.GetInt("system_logging.MaxAge"),
		Compress:   viper.GetBool("system_logging.Compress"),
	})
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Println("microseconds added to prefix")
}

// DbConn DB CONNECTION
func DbConn() (db *sql.DB) {
	dbDriver := viper.GetString("db.driver")
	db, err := sql.Open(dbDriver, viper.GetString("db.dns"))
	if err != nil {
		log.Println(err)
		panic(err.Error())

	} else {

		log.Println("Database 1 created successfully")

	}
	return db
}
