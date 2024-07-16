package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/supplier_contact"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createSupplierContactHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		supplier_contactModel := models.SupplierContact{}
		c.Bind(&supplier_contactModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newSupplierContact := models.NewSupplierContactRepository(db.DB)
		newSupplierContact.Create(&supplier_contactModel)

		setFlashmessages(c, "success", "SupplierContact created successfully!!")
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
	return renderView(c, supplier_contact.SupplierContactIndex(
		"| Create SupplierContact",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_contact.CreateSupplierContact(),
	))

}

func supplier_contactListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSupplierContact := models.NewSupplierContactRepository(db.DB)
	allSupplierContact, err := newSupplierContact.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's SupplierContact List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, supplier_contact.SupplierContactIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_contact.SupplierContactList(titlePage, allSupplierContact),
	))
}

func updateSupplierContactHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSupplierContact := models.NewSupplierContactRepository(db.DB)

	supplier_contactModel, err := newSupplierContact.GetSingle(idParams)

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
	//supplier_contactModel := models.SupplierContact{}
	c.Bind(supplier_contactModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSupplierContact := models.NewSupplierContactRepository(db.DB)
	err = newSupplierContact.Update(supplier_contactModel)

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

	setFlashmessages(c, "success", "SupplierContact successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/supplier_contact/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, supplier_contact.SupplierContactIndex(
		fmt.Sprintf("| Edit SupplierContact #%d", supplier_contactModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		supplier_contact.UpdateSupplierContact(*supplier_contactModel, tz.(string)),
	))

}

func deleteSupplierContactHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SupplierContact{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSupplierContact := models.NewSupplierContactRepository(db.DB)
	err = newSupplierContact.Delete(idParams)
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

	setFlashmessages(c, "success", "SupplierContact successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/supplier_contact/list")
}
