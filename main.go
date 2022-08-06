package main

import (
	"database/sql"
	"fmt"

	//    "strconv"
	"log"
	"os"

	"github.com/gofiber/fiber"
	"github.com/gofiber/template/html"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 8080
	user     = "postgres"
	password = "161092"
	dbname   = "dbtasking"
)

type Task struct {
	Id       int    `json: "id"`
	Task     string `json: "task"`
	Assignee int    `json: "assigne"`
	Deadline string `json: "deadline"`
	Status   int    `json: "status"`
}

type Tasks struct {
	Tasks []Task `json: "products"`
}

type Edit struct {
	Id       int    `json: "id"`
	Task     string `json: "task"`
	Assignee int    `json: "assigne"`
	Deadline string `json: "deadline"`
	Status   int    `json: "status"`
}

type task struct {
	Id       int
	Task     string
	Assignee int
	Deadline string
	Status   int
}

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

	app.Post("/", func(c *fiber.Ctx) error {
		return postHandler(c, db)
	})

	app.Get("/update", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}
	app.Static("/", "./public") // add this before starting the app

	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
	// res := task{}
	result := make([]*task, 0)
	rows, err := db.Query("SELECT id,task,assignee,deadline,status FROM task")
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		res := new(task)
		rows.Scan(&res.Id, &res.Task, &res.Assignee, &res.Deadline, &res.Status)
		result = append(result, res)
	}

	return c.Render("index", fiber.Map{
		"Task": result,
	})
}

//  Post simpan data

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTask := task{}
	if err := c.BodyParser(&newTask); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	// fmt.Printf("%v", newTask)
	if newTask.Task != "" {
		fmt.Println(newTask.Task)
		fmt.Println(newTask.Assignee)
		fmt.Println(newTask.Deadline)
		_, err := db.Exec("INSERT into task (task,assignee,deadline,status) VALUES ($1,$2,$3,$4)", newTask.Task, newTask.Assignee, newTask.Deadline, newTask.Status)
		if err != nil {
			log.Fatalf("Gagal simpan, query error : %v", err)
		}
	}

	return c.Redirect("/")
}

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	id := c.Query("id")
	// var id string = "1"
	data := Edit{}
	rows, err := db.Query("SELECT id,task,assignee,deadline,status FROM task WHERE id = $1", id)
	fmt.Println(id)
	defer rows.Close()
	if err != nil {
		log.Fatalln(err)
		c.JSON("An error occured")
	}
	for rows.Next() {
		// res := new(Edit)
		rows.Scan(&data.Id, &data.Task, &data.Assignee, &data.Deadline, &data.Status)
		// result = append(result, res)
	}

	return c.JSON(&fiber.Map{
		"success": false,
		"message": "Successfully fetched product",
		"product": data,
	})
	// return c.Render("update", fiber.Map{
	// 	"Task": data,
	// })
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from task WHERE id=$1", todoToDelete)
	return c.SendString("deleted")
}
