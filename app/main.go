package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// Embed htmlファイルに埋め込むデータ構造体
type Embed struct {
	Title   string
	Message string
	Users   map[int]User
	Time    time.Time
}

// User db users
type User struct {
	ID       int
	Name     string
	Password string
}

const (
	// DriverName ドライバ名(mysql固定)
	DriverName = "mysql"
	// DataSourceName user:password@tcp(container-name:port)/dbname
	DataSourceName = "root:golang@tcp(mysql-container:3306)/golang_db"
)

var usr = make(map[int]User)
var templates = make(map[string]*template.Template)