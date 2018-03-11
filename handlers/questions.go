package handlers

import (
    "database/sql"
    "net/http"
    "strconv"

    "florence/models"

    "github.com/labstack/echo"
)

type H map[string]interface{}

func GetQuestions(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetQuestions(db))
    }
}

func PostQuestion(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var question models.Question
        c.Bind(&question)
        id, err := models.PostQuestion(db, question.Content, question.Answer)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        } else {
            return err
        }
    }
}

func GetQuestion(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        return c.JSON(http.StatusOK, models.GetQuestion(db, id))
    }
}

func PutQuestion(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("id"))
        var question models.Question
        c.Bind(&question)

        var success = models.PutQuestion(db, id, question.Content, question.Answer, question.IsDeleted)
        return c.JSON(http.StatusOK, success)
    }
}
