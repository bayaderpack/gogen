package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/supplier_group_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createSupplierGroupDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		supplier_group_descriptionModel := models.SupplierGroupDescription{}
		c.Bind(&supplier_group_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newSupplierGroupDescription := models.NewSupplierGroupDescriptionRepository(db.DB)
		newSupplierGroupDescription.Create(&supplier_group_descriptionModel)

		setFlashmessages(c, "success", "SupplierGroupDescription created successfully!!")
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
	return renderView(c, supplier_group_description.SupplierGroupDescriptionIndex(
		"| Create SupplierGroupDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_group_description.CreateSupplierGroupDescription(),
	))

}

func supplier_group_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSupplierGroupDescription := models.NewSupplierGroupDescriptionRepository(db.DB)
	allSupplierGroupDescription, err := newSupplierGroupDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's SupplierGroupDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, supplier_group_description.SupplierGroupDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_group_description.SupplierGroupDescriptionList(titlePage, allSupplierGroupDescription),
	))
}

func updateSupplierGroupDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSupplierGroupDescription := models.NewSupplierGroupDescriptionRepository(db.DB)

	supplier_group_descriptionModel, err := newSupplierGroupDescription.GetSingle(idParams)

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
	//supplier_group_descriptionModel := models.SupplierGroupDescription{}
	c.Bind(supplier_group_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSupplierGroupDescription := models.NewSupplierGroupDescriptionRepository(db.DB)
	err = newSupplierGroupDescription.Update(supplier_group_descriptionModel)

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

	setFlashmessages(c, "success", "SupplierGroupDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/supplier_group_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, supplier_group_description.SupplierGroupDescriptionIndex(
		fmt.Sprintf("| Edit SupplierGroupDescription #%d", supplier_group_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		supplier_group_description.UpdateSupplierGroupDescription(*supplier_group_descriptionModel, tz.(string)),
	))

}

func deleteSupplierGroupDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SupplierGroupDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSupplierGroupDescription := models.NewSupplierGroupDescriptionRepository(db.DB)
	err = newSupplierGroupDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "SupplierGroupDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/supplier_group_description/list")
}
