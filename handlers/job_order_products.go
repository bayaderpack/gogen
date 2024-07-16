package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/job_order_products"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createJobOrderProductsHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		job_order_productsModel := models.JobOrderProducts{}
		c.Bind(&job_order_productsModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newJobOrderProducts := models.NewJobOrderProductsRepository(db.DB)
		newJobOrderProducts.Create(&job_order_productsModel)

		setFlashmessages(c, "success", "JobOrderProducts created successfully!!")
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
	return renderView(c, job_order_products.JobOrderProductsIndex(
		"| Create JobOrderProducts",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_products.CreateJobOrderProducts(),
	))

}

func job_order_productsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newJobOrderProducts := models.NewJobOrderProductsRepository(db.DB)
	allJobOrderProducts, err := newJobOrderProducts.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's JobOrderProducts List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, job_order_products.JobOrderProductsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		job_order_products.JobOrderProductsList(titlePage, allJobOrderProducts),
	))
}

func updateJobOrderProductsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newJobOrderProducts := models.NewJobOrderProductsRepository(db.DB)

	job_order_productsModel, err := newJobOrderProducts.GetSingle(idParams)

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
	//job_order_productsModel := models.JobOrderProducts{}
	c.Bind(job_order_productsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newJobOrderProducts := models.NewJobOrderProductsRepository(db.DB)
	err = newJobOrderProducts.Update(job_order_productsModel)

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

	setFlashmessages(c, "success", "JobOrderProducts successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/job_order_products/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, job_order_products.JobOrderProductsIndex(
		fmt.Sprintf("| Edit JobOrderProducts #%d", job_order_productsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		job_order_products.UpdateJobOrderProducts(*job_order_productsModel, tz.(string)),
	))

}

func deleteJobOrderProductsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.JobOrderProducts{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newJobOrderProducts := models.NewJobOrderProductsRepository(db.DB)
	err = newJobOrderProducts.Delete(idParams)
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

	setFlashmessages(c, "success", "JobOrderProducts successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/job_order_products/list")
}
