package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/media_variation"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


func createMediaVariationHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	media_variationModel := models.MediaVariation{}
	c.Bind(&media_variationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newMediaVariation := models.NewMediaVariationRepository(db.DB)
	newMediaVariation.Create(&media_variationModel)

setFlashmessages(c, "success", "MediaVariation created successfully!!")
	c.JSON(http.StatusOK, gin.H{
			"Code":   200,
		"Status": "Ok",
		"Data":   nil,
	})
	}
	username_key_value, ok  := c.Get(username_key)
	if !ok {
		fmt.Println("Some error")
	}
	return renderView(c, media_variation.MediaVariationIndex(
		"| Create MediaVariation",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_variation.CreateMediaVariation(),
	))

}

func media_variationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMediaVariation := models.NewMediaVariationRepository(db.DB)
	allMediaVariation, err := newMediaVariation.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's MediaVariation List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, media_variation.MediaVariationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_variation.MediaVariationList(titlePage, allMediaVariation),
	))
}

func updateMediaVariationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	t := services.Todo{
		ID:        idParams,
		CreatedBy: c.Get(user_id_key).(int),
	}

	todo, err := th.TodoServices.GetTodoById(t)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {

			return echo.NewHTTPError(
				echo.ErrNotFound.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}

	if c.Request().Method == "POST" {
		var status bool
		if c.FormValue("status") == "on" {
			status = true
		} else {
			status = false
		}

		todo := services.Todo{
			Title:       strings.Trim(c.FormValue("title"), " "),
			Description: strings.Trim(c.FormValue("description"), " "),
			Status:      status,
			CreatedBy:   c.Get(user_id_key).(int),
			ID:          idParams,
		}

		_, err := th.TodoServices.UpdateTodo(todo)
		if err != nil {
			return err
		}

		setFlashmessages(c, "success", "Task successfully updated!!")

		return c.Redirect(http.StatusSeeOther, "/todo/list")
	}

	return renderView(c, todo_views.TodoIndex(
		fmt.Sprintf("| Edit Todo #%d", todo.ID),
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		todo_views.UpdateTodo(todo, c.Get(tzone_key).(string)),
	))
}

func deleteTodoHandler(c echo.Context) error {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
		return err
	}

	t := services.Todo{
		CreatedBy: c.Get(user_id_key).(int),
		ID:        idParams,
	}

	err = th.TodoServices.DeleteTodo(t)
	if err != nil {
		if strings.Contains(err.Error(), "an affected row was expected") {

			return echo.NewHTTPError(
				echo.ErrNotFound.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}

	setFlashmessages(c, "success", "Task successfully deleted!!")

	return c.Redirect(http.StatusSeeOther, "/todo/list")
}
