package datasources

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
	// Import mysql driver
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DBMain for main DB connection
	DBMain *gorm.DB
	// DBMainConnStr for main DB connection
	DBMainConnStr string
)

// DateTimeFormat Long date time mysql format
const DateTimeFormat = "2006-02-01 15:04:05"

// DateFormat Date mysql format
const DateFormat = "2006-02-01"

// TimeFormat Time mysql format
const TimeFormat = "15:04:05"

// mariadbConfig for init connection
type mariadbConfig struct {
	// Optional.
	Username, Password string

	// Host of the mariadb instance.
	//
	// If set, UnixSocket should be unset.
	Host string

	// Port of the mariadb instance.
	//
	// If set, UnixSocket should be unset.
	Port int

	// UnixSocket is the filepath to a unix socket.
	//
	// If set, Host and Port should be unset.
	UnixSocket string
}

// mariadbDStoreString returns a connection string suitable for sql.Open.
func (c mariadbConfig) mariadbDStoreString(databaseName string) string {
	var cred string
	// [username[:password]@]
	if c.Username != "" {
		cred = c.Username
		if c.Password != "" {
			cred = cred + ":" + c.Password
		}
		cred = cred + "@"
	}

	if c.UnixSocket != "" {
		return fmt.Sprintf("%sunix(%s)/%s", cred, c.UnixSocket, databaseName)
	}
	return fmt.Sprintf("%stcp([%s]:%d)/%s", cred, c.Host, c.Port, databaseName)
}

// NewMariadbDB creates a new database connection backed by a given mariadb server.
func NewMariadbDB(dbname string, username string, password string, host string, portStr string, socket string, maxIdleConns int, maxOpenConns int, parseTime bool) (conn *gorm.DB, connStr string, err error) {

	// Use system default database if empty
	if len(dbname) == 0 {
		dbname = os.Getenv("DB_NAME")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, connStr, fmt.Errorf("MariaDB: port must be number string: %v", err)
	}

	dbConnOption := mariadbConfig{
		Username:   username,
		Password:   password,
		Host:       host,
		Port:       port,
		UnixSocket: socket,
	}

	connStr = dbConnOption.mariadbDStoreString(dbname)
	connStr = connStr + "?loc=Asia%2FBangkok&time_zone=%27Asia%2FBangkok%27"
	connStr = connStr + "&charset=utf8mb4,utf8"
	if parseTime {
		connStr = connStr + "&parseTime=true"
	}

	// Use system default database if empty
	if len(connStr) == 0 {
		return nil, connStr, fmt.Errorf("MariaDB: connStr needed")
	}
	// Open connection to database
	conn, err = gorm.Open("mysql", connStr)
	if err != nil {
		return nil, connStr, fmt.Errorf("MariaDB: could not get a connection: %v", err)
	}

	// Set max open connection at time
	if maxOpenConns > 0 {
		conn.DB().SetMaxOpenConns(maxOpenConns)
	} else {
		// Default value follow mariadb.js pool
		conn.DB().SetMaxOpenConns(10)
	}

	// Set max idle connection at time
	if maxIdleConns > 0 {
		conn.DB().SetMaxIdleConns(maxIdleConns)
	} else {
		// Default value follow mariadb.js pool
		conn.DB().SetMaxIdleConns(5)
	}

	// Time out for long connection
	// follow mariadb.js pool
	conn.DB().SetConnMaxLifetime(1800 * time.Second)

	return
}

// GetDBMainConn get db connection
func GetDBMainConn() (conn *gorm.DB, err error) {
	checkPing := DBMain.DB().Ping()
	if checkPing != nil {
		log.Printf("MariaDB: %+v", checkPing)
		DBMain, err = gorm.Open("mysql", DBMainConnStr)
		if err != nil {
			log.Printf("MariaDB: %+v", err)
			return nil, fmt.Errorf("MariaDB: could not get a connection: %v", err)
		}
	}
	return DBMain, err
}
