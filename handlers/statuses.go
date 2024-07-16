package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/statuses"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createStatusesHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		statusesModel := models.Statuses{}
		c.Bind(&statusesModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newStatuses := models.NewStatusesRepository(db.DB)
		newStatuses.Create(&statusesModel)

		setFlashmessages(c, "success", "Statuses created successfully!!")
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
	return renderView(c, statuses.StatusesIndex(
		"| Create Statuses",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		statuses.CreateStatuses(),
	))

}

func statusesListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newStatuses := models.NewStatusesRepository(db.DB)
	allStatuses, err := newStatuses.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's Statuses List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, statuses.StatusesIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		statuses.StatusesList(titlePage, allStatuses),
	))
}

func updateStatusesHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newStatuses := models.NewStatusesRepository(db.DB)

	statusesModel, err := newStatuses.GetSingle(idParams)

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
	//statusesModel := models.Statuses{}
	c.Bind(statusesModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newStatuses := models.NewStatusesRepository(db.DB)
	err = newStatuses.Update(statusesModel)

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

	setFlashmessages(c, "success", "Statuses successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/statuses/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, statuses.StatusesIndex(
		fmt.Sprintf("| Edit Statuses #%d", statusesModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		statuses.UpdateStatuses(*statusesModel, tz.(string)),
	))

}

func deleteStatusesHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Statuses{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newStatuses := models.NewStatusesRepository(db.DB)
	err = newStatuses.Delete(idParams)
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

	setFlashmessages(c, "success", "Statuses successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/statuses/list")
}
