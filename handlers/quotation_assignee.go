package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/quotation_assignee"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createQuotationAssigneeHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	quotation_assigneeModel := models.QuotationAssignee{}
	c.Bind(&quotation_assigneeModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newQuotationAssignee := models.NewQuotationAssigneeRepository(db.DB)
	newQuotationAssignee.Create(&quotation_assigneeModel)

setFlashmessages(c, "success", "QuotationAssignee created successfully!!")
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
	return renderView(c, quotation_assignee.QuotationAssigneeIndex(
		"| Create QuotationAssignee",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_assignee.CreateQuotationAssignee(),
	))

}

func quotation_assigneeListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationAssignee := models.NewQuotationAssigneeRepository(db.DB)
	allQuotationAssignee, err := newQuotationAssignee.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's QuotationAssignee List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_assignee.QuotationAssigneeIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_assignee.QuotationAssigneeList(titlePage, allQuotationAssignee),
	))
}

func updateQuotationAssigneeHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationAssignee := models.NewQuotationAssigneeRepository(db.DB)
	
	quotation_assigneeModel , err := newQuotationAssignee.GetSingle(idParams)

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
		//quotation_assigneeModel := models.QuotationAssignee{}
	c.Bind(quotation_assigneeModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationAssignee := models.NewQuotationAssigneeRepository(db.DB)
	err = newQuotationAssignee.Update(quotation_assigneeModel)

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


		setFlashmessages(c, "success", "QuotationAssignee successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_assignee/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, quotation_assignee.QuotationAssigneeIndex(
		fmt.Sprintf("| Edit QuotationAssignee #%d", quotation_assigneeModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_assignee.UpdateQuotationAssignee(*quotation_assigneeModel, tz.(string)),
	))

}

func deleteQuotationAssigneeHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationAssignee{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationAssignee := models.NewQuotationAssigneeRepository(db.DB)
	err = newQuotationAssignee.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationAssignee successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_assignee/list")
}
