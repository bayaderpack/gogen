package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/job_order"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createJobOrderHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	job_orderModel := models.JobOrder{}
	c.Bind(&job_orderModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newJobOrder := models.NewJobOrderRepository(db.DB)
	newJobOrder.Create(&job_orderModel)

setFlashmessages(c, "success", "JobOrder created successfully!!")
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
	return renderView(c, job_order.JobOrderIndex(
		"| Create JobOrder",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order.CreateJobOrder(),
	))

}

func job_orderListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newJobOrder := models.NewJobOrderRepository(db.DB)
	allJobOrder, err := newJobOrder.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's JobOrder List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, job_order.JobOrderIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order.JobOrderList(titlePage, allJobOrder),
	))
}

func updateJobOrderHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newJobOrder := models.NewJobOrderRepository(db.DB)
	
	job_orderModel , err := newJobOrder.GetSingle(idParams)

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
		//job_orderModel := models.JobOrder{}
	c.Bind(job_orderModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newJobOrder := models.NewJobOrderRepository(db.DB)
	err = newJobOrder.Update(job_orderModel)

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


		setFlashmessages(c, "success", "JobOrder successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/job_order/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, job_order.JobOrderIndex(
		fmt.Sprintf("| Edit JobOrder #%d", job_orderModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		job_order.UpdateJobOrder(*job_orderModel, tz.(string)),
	))

}

func deleteJobOrderHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.JobOrder{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newJobOrder := models.NewJobOrderRepository(db.DB)
	err = newJobOrder.Delete(idParams)
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

	setFlashmessages(c, "success", "JobOrder successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/job_order/list")
}
