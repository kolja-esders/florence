package handlers

import (
    "database/sql"
    "net/http"

    "florence/models"

    "github.com/labstack/echo"
)

type H map[string]interface{}

func GetQuestions(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetQuestions(db))
    }
}

func PutQuestion(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var question models.Question
        c.Bind(&question)
        id, err := models.PutQuestion(db, question.Content, question.Answer)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        } else {
            return err
        }
    }
}
