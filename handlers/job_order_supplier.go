package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/job_order_supplier"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)


func createJobOrderSupplierHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	job_order_supplierModel := models.JobOrderSupplier{}
	c.Bind(&job_order_supplierModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newJobOrderSupplier := models.NewJobOrderSupplierRepository(db.DB)
	newJobOrderSupplier.Create(&job_order_supplierModel)

setFlashmessages(c, "success", "JobOrderSupplier created successfully!!")
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
	return renderView(c, job_order_supplier.JobOrderSupplierIndex(
		"| Create JobOrderSupplier",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_supplier.CreateJobOrderSupplier(),
	))

}

func job_order_supplierListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newJobOrderSupplier := models.NewJobOrderSupplierRepository(db.DB)
	allJobOrderSupplier, err := newJobOrderSupplier.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's JobOrderSupplier List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, job_order_supplier.JobOrderSupplierIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_supplier.JobOrderSupplierList(titlePage, allJobOrderSupplier),
	))
}

func updateJobOrderSupplierHandler(c *gin.Context) error {
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
