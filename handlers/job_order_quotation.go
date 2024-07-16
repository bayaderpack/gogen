package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/job_order_quotation"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createJobOrderQuotationHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	job_order_quotationModel := models.JobOrderQuotation{}
	c.Bind(&job_order_quotationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newJobOrderQuotation := models.NewJobOrderQuotationRepository(db.DB)
	newJobOrderQuotation.Create(&job_order_quotationModel)

setFlashmessages(c, "success", "JobOrderQuotation created successfully!!")
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
	return renderView(c, job_order_quotation.JobOrderQuotationIndex(
		"| Create JobOrderQuotation",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_quotation.CreateJobOrderQuotation(),
	))

}

func job_order_quotationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newJobOrderQuotation := models.NewJobOrderQuotationRepository(db.DB)
	allJobOrderQuotation, err := newJobOrderQuotation.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's JobOrderQuotation List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, job_order_quotation.JobOrderQuotationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_quotation.JobOrderQuotationList(titlePage, allJobOrderQuotation),
	))
}

func updateJobOrderQuotationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newJobOrderQuotation := models.NewJobOrderQuotationRepository(db.DB)
	
	job_order_quotationModel , err := newJobOrderQuotation.GetSingle(idParams)

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
		//job_order_quotationModel := models.JobOrderQuotation{}
	c.Bind(job_order_quotationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newJobOrderQuotation := models.NewJobOrderQuotationRepository(db.DB)
	err = newJobOrderQuotation.Update(job_order_quotationModel)

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


		setFlashmessages(c, "success", "JobOrderQuotation successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/job_order_quotation/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, job_order_quotation.JobOrderQuotationIndex(
		fmt.Sprintf("| Edit JobOrderQuotation #%d", job_order_quotationModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		job_order_quotation.UpdateJobOrderQuotation(*job_order_quotationModel, tz.(string)),
	))

}

func deleteJobOrderQuotationHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.JobOrderQuotation{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newJobOrderQuotation := models.NewJobOrderQuotationRepository(db.DB)
	err = newJobOrderQuotation.Delete(idParams)
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

	setFlashmessages(c, "success", "JobOrderQuotation successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/job_order_quotation/list")
}
