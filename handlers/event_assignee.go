package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/event_assignee"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createEventAssigneeHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		event_assigneeModel := models.EventAssignee{}
		c.Bind(&event_assigneeModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newEventAssignee := models.NewEventAssigneeRepository(db.DB)
		newEventAssignee.Create(&event_assigneeModel)

		setFlashmessages(c, "success", "EventAssignee created successfully!!")
		c.JSON(http.StatusOK, gin.H{
			"Code":   200,
			"Status": "Ok",
			"Data":   nil,
		})
	}
	username_key_value, ok := c.Get(username_key)
	if !ok {
		fmt.Println("Some error")
	}
	return renderView(c, event_assignee.EventAssigneeIndex(
		"| Create EventAssignee",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		event_assignee.CreateEventAssignee(),
	))

}

func event_assigneeListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newEventAssignee := models.NewEventAssigneeRepository(db.DB)
	allEventAssignee, err := newEventAssignee.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's EventAssignee List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, event_assignee.EventAssigneeIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		event_assignee.EventAssigneeList(titlePage, allEventAssignee),
	))
}

func updateEventAssigneeHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newEventAssignee := models.NewEventAssigneeRepository(db.DB)

	event_assigneeModel, err := newEventAssignee.GetSingle(idParams)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))
		}
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
			"something went wrong: %s",
			err,
		))
	}
	//event_assigneeModel := models.EventAssignee{}
	c.Bind(event_assigneeModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newEventAssignee := models.NewEventAssigneeRepository(db.DB)
	err = newEventAssignee.Update(event_assigneeModel)

	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {

			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))

		}

		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
			"something went wrong: %s",
			err,
		))

	}

	setFlashmessages(c, "success", "EventAssignee successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/event_assignee/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, event_assignee.EventAssigneeIndex(
		fmt.Sprintf("| Edit EventAssignee #%d", event_assigneeModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		event_assignee.UpdateEventAssignee(*event_assigneeModel, tz.(string)),
	))

}

func deleteEventAssigneeHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.EventAssignee{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newEventAssignee := models.NewEventAssigneeRepository(db.DB)
	err = newEventAssignee.Delete(idParams)
	if err != nil {
		if strings.Contains(err.Error(), "an affected row was expected") {

			c.AbortWithError(http.StatusNotFound, fmt.Errorf(
				"something went wrong: %s",
				err,
			))

		}

		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf(
			"something went wrong: %s",
			err,
		))

	}

	setFlashmessages(c, "success", "EventAssignee successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/event_assignee/list")
}
