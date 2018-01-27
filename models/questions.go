package models

import (
    "database/sql"

    _ "github.com/mattn/go-sqlite3"
)

type Question struct {
    ID int `json:"id"`
    Content string `json:"content"`
    Answer string `json:"answer"`
}

type QuestionCollection struct {
    Questions []Question `json:"items"`
}

func GetQuestions(db *sql.DB) QuestionCollection {
    sql := "SELECT * FROM questions"
    rows, err := db.Query(sql)
    if err != nil {
        panic(err)
    }
    defer rows.Close()

    result := QuestionCollection{}
    for rows.Next() {
        question := Question{}
        err2 := rows.Scan(&question.ID, &question.Content, &question.Answer)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }
        result.Questions = append(result.Questions, question)
    }
    return result
}
