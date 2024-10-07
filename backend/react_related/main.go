
package main

import (
	"html/template"
    "database/sql"
    "fmt"
    "log"
    "net/http"
    // "os"
    "time"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gofiber/fiber/v2"
)


type db_creator struct {
   creator_id int
   creator_name string
};

type db_sound_file struct {
	sound_title string
	sound_created_time time.Time
	sound_updated_time time.Time
	sound_file_path string
	sound_file_type string
	sound_file_size int
	sound_text_result string
};


func insert_sound_file (db *sql.DB,creator_id int64, sound_title string, sound_file_path string, sound_file_type string, sound_file_size int, sound_file_text_result string) (  int64){
	insert_stmt, err := db.Prepare("INSERT INTO Sound_Files(creator_id, title, file_path, file_type, file_size, text_result ) VALUES (?, ?, ?, ?, ?, ? )")
    if err != nil {
    	log.Fatal(err)
    }
	defer insert_stmt.Close()

	result_msg, err := insert_stmt.Exec(creator_id, sound_title, sound_file_path, sound_file_type, sound_file_size, sound_file_text_result)
	if err != nil{
		log.Fatal(err)
	}

	last_sound_file_id, err := result_msg.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("product already insert with id : %d\n", last_sound_file_id)
	return last_sound_file_id

}


func query_all_creator (db *sql.DB){
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
        fmt.Printf ("ID: %d,Name: %s \n", id, name )
    }

}

func main() {
	var db *sql.DB
	dsn := "username:password@tcp(localhost:3306)/records"

    // Capture connection properties.
    db, err := sql.Open("mysql", dsn)

    if (err != nil){
    	panic(err.Error());
    }
    // fmt.Fprintln(os.Stderr, "Success");
    defer db.Close();
	query_all_creator(db);

	test_sound_file := db_sound_file {
		sound_title: "sound_title",
		sound_created_time: time.Now(),
		sound_updated_time: time.Now(),
		sound_file_path: "file_path",
		sound_file_type: "file_type",
		sound_file_size: 0,
		sound_text_result: "Nothing yet",
	};


	last_sound_id := insert_sound_file(db, 01 , test_sound_file.sound_title , test_sound_file.sound_file_path , test_sound_file.sound_file_type , test_sound_file.sound_file_size , test_sound_file.sound_text_result )
	fmt.Printf("This is the sound file id: %d \n", last_sound_id)


	router := fiber.New()
	router.Engine.SetHTMLTemplate(template.Must(template.ParseGlob("templates/*.html")))

	api := router.Group("/api")
	{
		api.Get("/hello", func(ctx *fiber.Ctx) error {
			return ctx.Status(200).JSON( map[string]interface{}{"msg": "world"})
		})
		api.Get("/react", func(ctx *fiber.Ctx) error {
			return ctx.Status(200).Render("index.html", nil)
		})

	}
	fmt.Println ("Type of x:", (fiber.Map{"msg": "world"}))


	router.Get("/", func (ctx *fiber.Ctx) error{
		return ctx.Status(404).JSON(http.StatusNotFound) })
	log.Fatal(router.Listen(":8080"))
}
