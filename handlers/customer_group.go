package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/customer_group"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createCustomerGroupHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		customer_groupModel := models.CustomerGroup{}
		c.Bind(&customer_groupModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newCustomerGroup := models.NewCustomerGroupRepository(db.DB)
		newCustomerGroup.Create(&customer_groupModel)

		setFlashmessages(c, "success", "CustomerGroup created successfully!!")
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
	return renderView(c, customer_group.CustomerGroupIndex(
		"| Create CustomerGroup",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_group.CreateCustomerGroup(),
	))

}

func customer_groupListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCustomerGroup := models.NewCustomerGroupRepository(db.DB)
	allCustomerGroup, err := newCustomerGroup.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's CustomerGroup List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, customer_group.CustomerGroupIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_group.CustomerGroupList(titlePage, allCustomerGroup),
	))
}

func updateCustomerGroupHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCustomerGroup := models.NewCustomerGroupRepository(db.DB)

	customer_groupModel, err := newCustomerGroup.GetSingle(idParams)

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
	//customer_groupModel := models.CustomerGroup{}
	c.Bind(customer_groupModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCustomerGroup := models.NewCustomerGroupRepository(db.DB)
	err = newCustomerGroup.Update(customer_groupModel)

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

	setFlashmessages(c, "success", "CustomerGroup successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/customer_group/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, customer_group.CustomerGroupIndex(
		fmt.Sprintf("| Edit CustomerGroup #%d", customer_groupModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		customer_group.UpdateCustomerGroup(*customer_groupModel, tz.(string)),
	))

}

func deleteCustomerGroupHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CustomerGroup{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCustomerGroup := models.NewCustomerGroupRepository(db.DB)
	err = newCustomerGroup.Delete(idParams)
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

	setFlashmessages(c, "success", "CustomerGroup successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/customer_group/list")
}
