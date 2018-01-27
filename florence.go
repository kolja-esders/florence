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
        answer VARCHAR NOT NULL
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
    e.File("/", "leonardo/index.html")
    e.GET("/questions", handlers.GetQuestions(db))

    // Start as a web server
    e.Start(":8000")
}
