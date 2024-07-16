package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/contact_entry"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createContactEntryHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		contact_entryModel := models.ContactEntry{}
		c.Bind(&contact_entryModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newContactEntry := models.NewContactEntryRepository(db.DB)
		newContactEntry.Create(&contact_entryModel)

		setFlashmessages(c, "success", "ContactEntry created successfully!!")
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
	return renderView(c, contact_entry.ContactEntryIndex(
		"| Create ContactEntry",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		contact_entry.CreateContactEntry(),
	))

}

func contact_entryListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newContactEntry := models.NewContactEntryRepository(db.DB)
	allContactEntry, err := newContactEntry.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's ContactEntry List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, contact_entry.ContactEntryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		contact_entry.ContactEntryList(titlePage, allContactEntry),
	))
}

func updateContactEntryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newContactEntry := models.NewContactEntryRepository(db.DB)

	contact_entryModel, err := newContactEntry.GetSingle(idParams)

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
	//contact_entryModel := models.ContactEntry{}
	c.Bind(contact_entryModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newContactEntry := models.NewContactEntryRepository(db.DB)
	err = newContactEntry.Update(contact_entryModel)

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

	setFlashmessages(c, "success", "ContactEntry successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/contact_entry/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, contact_entry.ContactEntryIndex(
		fmt.Sprintf("| Edit ContactEntry #%d", contact_entryModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		contact_entry.UpdateContactEntry(*contact_entryModel, tz.(string)),
	))

}

func deleteContactEntryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ContactEntry{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newContactEntry := models.NewContactEntryRepository(db.DB)
	err = newContactEntry.Delete(idParams)
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

	setFlashmessages(c, "success", "ContactEntry successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/contact_entry/list")
}
