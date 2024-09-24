package main

import (
    "github.com/go-sql-driver/mysql"
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"
)

var db *sql.DB

type Old_record struct {
    ID     int64
    Title  string
    Sound_File Sound_File
    Result string
}

type Sound_File {
    ID_sound int64
    Date time.Time
}


func main() {
    // Capture connection properties.
    cfg := mysql.Config{
        User:   os.Getenv("DBUSER"),
        Passwd: os.Getenv("DBPASS"),
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "recordings",
    }
    // Get a database handle.
    var err error

    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}
