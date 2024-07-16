package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/sample_assignee"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSampleAssigneeHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	sample_assigneeModel := models.SampleAssignee{}
	c.Bind(&sample_assigneeModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSampleAssignee := models.NewSampleAssigneeRepository(db.DB)
	newSampleAssignee.Create(&sample_assigneeModel)

setFlashmessages(c, "success", "SampleAssignee created successfully!!")
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
	return renderView(c, sample_assignee.SampleAssigneeIndex(
		"| Create SampleAssignee",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_assignee.CreateSampleAssignee(),
	))

}

func sample_assigneeListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSampleAssignee := models.NewSampleAssigneeRepository(db.DB)
	allSampleAssignee, err := newSampleAssignee.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's SampleAssignee List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, sample_assignee.SampleAssigneeIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_assignee.SampleAssigneeList(titlePage, allSampleAssignee),
	))
}

func updateSampleAssigneeHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSampleAssignee := models.NewSampleAssigneeRepository(db.DB)
	
	sample_assigneeModel , err := newSampleAssignee.GetSingle(idParams)

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
		//sample_assigneeModel := models.SampleAssignee{}
	c.Bind(sample_assigneeModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSampleAssignee := models.NewSampleAssigneeRepository(db.DB)
	err = newSampleAssignee.Update(sample_assigneeModel)

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


		setFlashmessages(c, "success", "SampleAssignee successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/sample_assignee/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, sample_assignee.SampleAssigneeIndex(
		fmt.Sprintf("| Edit SampleAssignee #%d", sample_assigneeModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		sample_assignee.UpdateSampleAssignee(*sample_assigneeModel, tz.(string)),
	))

}

func deleteSampleAssigneeHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SampleAssignee{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSampleAssignee := models.NewSampleAssigneeRepository(db.DB)
	err = newSampleAssignee.Delete(idParams)
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

	setFlashmessages(c, "success", "SampleAssignee successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/sample_assignee/list")
}
