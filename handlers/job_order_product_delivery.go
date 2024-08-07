package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/job_order_product_delivery"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createJobOrderProductDeliveryHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		job_order_product_deliveryModel := models.JobOrderProductDelivery{}
		c.Bind(&job_order_product_deliveryModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newJobOrderProductDelivery := models.NewJobOrderProductDeliveryRepository(db.DB)
		newJobOrderProductDelivery.Create(&job_order_product_deliveryModel)

		setFlashmessages(c, "success", "JobOrderProductDelivery created successfully!!")
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
	return renderView(c, job_order_product_delivery.JobOrderProductDeliveryIndex(
		"| Create JobOrderProductDelivery",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_product_delivery.CreateJobOrderProductDelivery(),
	))

}

func job_order_product_deliveryListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newJobOrderProductDelivery := models.NewJobOrderProductDeliveryRepository(db.DB)
	allJobOrderProductDelivery, err := newJobOrderProductDelivery.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's JobOrderProductDelivery List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, job_order_product_delivery.JobOrderProductDeliveryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_product_delivery.JobOrderProductDeliveryList(titlePage, allJobOrderProductDelivery),
	))
}

func updateJobOrderProductDeliveryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newJobOrderProductDelivery := models.NewJobOrderProductDeliveryRepository(db.DB)

	job_order_product_deliveryModel, err := newJobOrderProductDelivery.GetSingle(idParams)

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
	//job_order_product_deliveryModel := models.JobOrderProductDelivery{}
	c.Bind(job_order_product_deliveryModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newJobOrderProductDelivery := models.NewJobOrderProductDeliveryRepository(db.DB)
	err = newJobOrderProductDelivery.Update(job_order_product_deliveryModel)

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

	setFlashmessages(c, "success", "JobOrderProductDelivery successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/job_order_product_delivery/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, job_order_product_delivery.JobOrderProductDeliveryIndex(
		fmt.Sprintf("| Edit JobOrderProductDelivery #%d", job_order_product_deliveryModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		job_order_product_delivery.UpdateJobOrderProductDelivery(*job_order_product_deliveryModel, tz.(string)),
	))

}

func deleteJobOrderProductDeliveryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.JobOrderProductDelivery{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newJobOrderProductDelivery := models.NewJobOrderProductDeliveryRepository(db.DB)
	err = newJobOrderProductDelivery.Delete(idParams)
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

	setFlashmessages(c, "success", "JobOrderProductDelivery successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/job_order_product_delivery/list")
}
