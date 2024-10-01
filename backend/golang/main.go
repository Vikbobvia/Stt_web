package main

import (
    "database/sql"
    "fmt"
    // "log"
    // "os"
    "time"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DB_Creator struct {
   creator_id int
   creator_name string
};

type DB_Sound_File struct {
	sound_title string
	sound_created_time time.Time
	sound_updated_time time.Time
	sound_file_path string
	sound_file_type string
	sound_file_size int
	sound_Text_result string
};


func main() {
    // Capture connection properties.
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/records");
    if (err != nil){
    	panic(err.Error());
    }
    // fmt.Fprintln(os.Stderr, "Success");
    defer db.Close();

    // Query data ( creators )
    rows, err := db.Query("SELECT id, name FROM Creators")
    if err != nil {
    	panic(err.Error())
    }
    defer rows.Close()

    for rows.Next() {
    	var id int;
     	var name string;
     	err = rows.Scan(&id, &name)
      	if err != nil {
          	panic(err.Error())
          }
        fmt.Printf ("ID: %d,Name: %s", id, name )
    }

}
