package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/event_history"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createEventHistoryHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		event_historyModel := models.EventHistory{}
		c.Bind(&event_historyModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newEventHistory := models.NewEventHistoryRepository(db.DB)
		newEventHistory.Create(&event_historyModel)

		setFlashmessages(c, "success", "EventHistory created successfully!!")
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
	return renderView(c, event_history.EventHistoryIndex(
		"| Create EventHistory",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		event_history.CreateEventHistory(),
	))

}

func event_historyListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newEventHistory := models.NewEventHistoryRepository(db.DB)
	allEventHistory, err := newEventHistory.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's EventHistory List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, event_history.EventHistoryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		event_history.EventHistoryList(titlePage, allEventHistory),
	))
}

func updateEventHistoryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newEventHistory := models.NewEventHistoryRepository(db.DB)

	event_historyModel, err := newEventHistory.GetSingle(idParams)

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
	//event_historyModel := models.EventHistory{}
	c.Bind(event_historyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newEventHistory := models.NewEventHistoryRepository(db.DB)
	err = newEventHistory.Update(event_historyModel)

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

	setFlashmessages(c, "success", "EventHistory successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/event_history/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, event_history.EventHistoryIndex(
		fmt.Sprintf("| Edit EventHistory #%d", event_historyModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		event_history.UpdateEventHistory(*event_historyModel, tz.(string)),
	))

}

func deleteEventHistoryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.EventHistory{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newEventHistory := models.NewEventHistoryRepository(db.DB)
	err = newEventHistory.Delete(idParams)
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

	setFlashmessages(c, "success", "EventHistory successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/event_history/list")
}
