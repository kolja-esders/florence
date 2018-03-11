package main

import (
    "database/sql"
    "florence/handlers"

    "github.com/labstack/echo"
    _ "github.com/mattn/go-sqlite3"
)

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    if err != nil {
        panic(err)
    }
    if db == nil {
        panic("db nil")
    }
    return db
}

func migrate(db *sql.DB) {
    sql := `
    CREATE TABLE IF NOT EXISTS questions(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
        content VARCHAR NOT NULL,
        answer VARCHAR NOT NULL,
        is_deleted INT DEFAULT 0
    );
    `
    _, err := db.Exec(sql)
    if err != nil {
        panic(err)
    }
}

func main() {
    db := initDB("storage.db")
    migrate(db)

    e := echo.New()

    e.Static("/dist", "leonardo/dist")
    e.GET("/api/questions", handlers.GetQuestions(db))
    e.POST("/api/questions", handlers.PostQuestion(db))

    e.GET("/api/questions/:id", handlers.GetQuestion(db))
    e.PUT("/api/questions/:id", handlers.PutQuestion(db))

    e.File("/*", "leonardo/index.html")

    // Start as a web server
    e.Start(":8000")
}
