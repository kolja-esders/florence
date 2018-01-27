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

func PutQuestion(db *sql.DB, content string, answer string) (int64, error) {
    sql := "INSERT INTO questions(content, answer) VALUES(?, ?)"

    stmt, err := db.Prepare(sql)
    if err != nil {
        panic(err)
    }
    defer stmt.Close()

    result, err2 := stmt.Exec(content, answer)
    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}
