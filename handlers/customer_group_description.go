package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/customer_group_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createCustomerGroupDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		customer_group_descriptionModel := models.CustomerGroupDescription{}
		c.Bind(&customer_group_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newCustomerGroupDescription := models.NewCustomerGroupDescriptionRepository(db.DB)
		newCustomerGroupDescription.Create(&customer_group_descriptionModel)

		setFlashmessages(c, "success", "CustomerGroupDescription created successfully!!")
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
	return renderView(c, customer_group_description.CustomerGroupDescriptionIndex(
		"| Create CustomerGroupDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_group_description.CreateCustomerGroupDescription(),
	))

}

func customer_group_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCustomerGroupDescription := models.NewCustomerGroupDescriptionRepository(db.DB)
	allCustomerGroupDescription, err := newCustomerGroupDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's CustomerGroupDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, customer_group_description.CustomerGroupDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		customer_group_description.CustomerGroupDescriptionList(titlePage, allCustomerGroupDescription),
	))
}

func updateCustomerGroupDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCustomerGroupDescription := models.NewCustomerGroupDescriptionRepository(db.DB)

	customer_group_descriptionModel, err := newCustomerGroupDescription.GetSingle(idParams)

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
	//customer_group_descriptionModel := models.CustomerGroupDescription{}
	c.Bind(customer_group_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCustomerGroupDescription := models.NewCustomerGroupDescriptionRepository(db.DB)
	err = newCustomerGroupDescription.Update(customer_group_descriptionModel)

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

	setFlashmessages(c, "success", "CustomerGroupDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/customer_group_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, customer_group_description.CustomerGroupDescriptionIndex(
		fmt.Sprintf("| Edit CustomerGroupDescription #%d", customer_group_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		customer_group_description.UpdateCustomerGroupDescription(*customer_group_descriptionModel, tz.(string)),
	))

}

func deleteCustomerGroupDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CustomerGroupDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCustomerGroupDescription := models.NewCustomerGroupDescriptionRepository(db.DB)
	err = newCustomerGroupDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "CustomerGroupDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/customer_group_description/list")
}
