package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/supplier_group"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSupplierGroupHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	supplier_groupModel := models.SupplierGroup{}
	c.Bind(&supplier_groupModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSupplierGroup := models.NewSupplierGroupRepository(db.DB)
	newSupplierGroup.Create(&supplier_groupModel)

setFlashmessages(c, "success", "SupplierGroup created successfully!!")
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
	return renderView(c, supplier_group.SupplierGroupIndex(
		"| Create SupplierGroup",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_group.CreateSupplierGroup(),
	))

}

func supplier_groupListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSupplierGroup := models.NewSupplierGroupRepository(db.DB)
	allSupplierGroup, err := newSupplierGroup.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's SupplierGroup List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, supplier_group.SupplierGroupIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		supplier_group.SupplierGroupList(titlePage, allSupplierGroup),
	))
}

func updateSupplierGroupHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSupplierGroup := models.NewSupplierGroupRepository(db.DB)
	
	supplier_groupModel , err := newSupplierGroup.GetSingle(idParams)

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
		//supplier_groupModel := models.SupplierGroup{}
	c.Bind(supplier_groupModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSupplierGroup := models.NewSupplierGroupRepository(db.DB)
	err = newSupplierGroup.Update(supplier_groupModel)

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


		setFlashmessages(c, "success", "SupplierGroup successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/supplier_group/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, supplier_group.SupplierGroupIndex(
		fmt.Sprintf("| Edit SupplierGroup #%d", supplier_groupModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		supplier_group.UpdateSupplierGroup(*supplier_groupModel, tz.(string)),
	))

}

func deleteSupplierGroupHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SupplierGroup{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSupplierGroup := models.NewSupplierGroupRepository(db.DB)
	err = newSupplierGroup.Delete(idParams)
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

	setFlashmessages(c, "success", "SupplierGroup successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/supplier_group/list")
}
