package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/job_order_supplier"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
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
	newJobOrderSupplier := models.NewJobOrderSupplierRepository(db.DB)
	
	job_order_supplierModel , err := newJobOrderSupplier.GetSingle(idParams)

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
		//job_order_supplierModel := models.JobOrderSupplier{}
	c.Bind(job_order_supplierModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newJobOrderSupplier := models.NewJobOrderSupplierRepository(db.DB)
	err = newJobOrderSupplier.Update(job_order_supplierModel)

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


		setFlashmessages(c, "success", "JobOrderSupplier successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/job_order_supplier/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, job_order_supplier.JobOrderSupplierIndex(
		fmt.Sprintf("| Edit JobOrderSupplier #%d", job_order_supplierModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		job_order_supplier.UpdateJobOrderSupplier(*job_order_supplierModel, tz.(string)),
	))

}

func deleteJobOrderSupplierHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.JobOrderSupplier{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newJobOrderSupplier := models.NewJobOrderSupplierRepository(db.DB)
	err = newJobOrderSupplier.Delete(idParams)
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

	setFlashmessages(c, "success", "JobOrderSupplier successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/job_order_supplier/list")
}
