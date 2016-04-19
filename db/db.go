package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	config map[string]interface{}
)

func init() {
	config = make(map[string]interface{})
}

//Config 配置数据库
func Config(driver, user, password, database string, host string, port int, charset string) {
	config["driver"] = driver
	config["user"] = user
	config["password"] = password
	config["host"] = host
	config["port"] = port
	config["database"] = database
	config["charset"] = charset
}

//Connect 连接数据库
func Connect() error {
	var err error
	db, err = getConnect(config["driver"].(string),
		config["user"].(string), config["password"].(string),
		config["host"].(string), config["port"].(int),
		config["database"].(string), config["charset"].(string))
	if err != nil {
		return err
	}
	return nil
}

func getConnect(driver, user, password, host string, port int, database, charset string) (*sql.DB, error) {
	conn, err := sql.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		user, password, host, port, database, charset))
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func query(sql string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(sql, args...)
}

func queryRow(sql string, args ...interface{}) *sql.Row {
	return db.QueryRow(sql, args...)
}

func exec(sql string, args ...interface{}) (sql.Result, error) {
	return db.Exec(sql, args...)
}

func count(table string) int64 {
	row := queryRow("select count(*) as num from " + table)
	var n int64
	err := row.Scan(&n)
	if err != nil {
		return -1
	}
	return n
}

func countBy(table string, where string, args ...interface{}) int64 {
	row := queryRow(fmt.Sprintf("select count(*) as num from %s where %s", table, where), args...)
	var n int64
	err := row.Scan(&n)
	if err != nil {
		return -1
	}
	return n
}