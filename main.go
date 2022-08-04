package main

import (
	"database/sql"
   "fmt"
   "log"
   "os"
   _ "github.com/lib/pq"
   "github.com/gofiber/fiber"
   "github.com/gofiber/template/html"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "161092"
	dbname   = "dbtasking"
  )

func main() {

	//    Konek ke Database
	// connStr := "postgresql://postgres:161092@localhost:8888/dbtasking?sslmode=disable"
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
	// Connect to database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	// end Konek ke Database

   engine := html.New("./views", ".html")
   app := fiber.New(fiber.Config{
	Views: engine,
   })
   

   app.Get("/", func(c *fiber.Ctx) error {
		return indexHandler(c, db)
	})

	// app.Post("/", func(c *fiber.Ctx) error {
	// 	return postHandler(c, db)
	// })

	// app.Put("/update", func(c *fiber.Ctx) error {
	// 	return putHandler(c, db)
	// })

	// app.Delete("/delete", func(c *fiber.Ctx) error {
	// 	return deleteHandler(c, db)
	// })

   port := os.Getenv("PORT")
   if port == "" {
       port = "8888"
   }
   app.Static("/", "./public") // add this before starting the app

   log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	var res string
	var task []string
	rows, err := db.Query("SELECT * FROM task")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		rows.Scan(&res)
		task = append(task, res)
	}
	return c.Render("index", fiber.Map{
		"Task": task,
	})
 }

//  func postHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello")
//  }
//  func putHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello")
//  }
//  func deleteHandler(c *fiber.Ctx) error {
// 	return c.SendString("Hello")
//  }