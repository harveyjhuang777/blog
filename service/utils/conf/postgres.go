package conf

import (
	"net"
	"net/url"
	"os"
	"strings"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ConnectionInfo struct {
	Scheme   string
	Host     string
	Port     string
	Username string
	Password string
	DBname   string
}

func GenDatabaseUrl() (*ConnectionInfo, error) {
	databaseUrl, ok := os.LookupEnv("DATABASE_URL")
	if !ok {
		panic("DATABASE_URL not set")
	}
	dbUrl, err := url.Parse(databaseUrl)
	if err != nil {
		panic("DATABASE_URL wrong format")
	}
	scheme := dbUrl.Scheme
	host, port, _ := net.SplitHostPort(dbUrl.Host)
	user := dbUrl.User.Username()
	pass, _ := dbUrl.User.Password()
	dbname := strings.Trim(dbUrl.Path, "/")
	connInfo := &ConnectionInfo{scheme, host, port, user, pass, dbname}
	return connInfo, nil
}
