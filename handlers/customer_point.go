package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/customer_point"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createCustomerPointHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	customer_pointModel := models.CustomerPoint{}
	c.Bind(&customer_pointModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newCustomerPoint := models.NewCustomerPointRepository(db.DB)
	newCustomerPoint.Create(&customer_pointModel)

setFlashmessages(c, "success", "CustomerPoint created successfully!!")
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
	return renderView(c, customer_point.CustomerPointIndex(
		"| Create CustomerPoint",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_point.CreateCustomerPoint(),
	))

}

func customer_pointListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCustomerPoint := models.NewCustomerPointRepository(db.DB)
	allCustomerPoint, err := newCustomerPoint.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's CustomerPoint List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, customer_point.CustomerPointIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_point.CustomerPointList(titlePage, allCustomerPoint),
	))
}

func updateCustomerPointHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCustomerPoint := models.NewCustomerPointRepository(db.DB)
	
	customer_pointModel , err := newCustomerPoint.GetSingle(idParams)

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
		//customer_pointModel := models.CustomerPoint{}
	c.Bind(customer_pointModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCustomerPoint := models.NewCustomerPointRepository(db.DB)
	err = newCustomerPoint.Update(customer_pointModel)

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


		setFlashmessages(c, "success", "CustomerPoint successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/customer_point/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, customer_point.CustomerPointIndex(
		fmt.Sprintf("| Edit CustomerPoint #%d", customer_pointModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		customer_point.UpdateCustomerPoint(*customer_pointModel, tz.(string)),
	))

}

func deleteCustomerPointHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CustomerPoint{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCustomerPoint := models.NewCustomerPointRepository(db.DB)
	err = newCustomerPoint.Delete(idParams)
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

	setFlashmessages(c, "success", "CustomerPoint successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/customer_point/list")
}
