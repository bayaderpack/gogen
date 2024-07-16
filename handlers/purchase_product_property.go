package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/purchase_product_property"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createPurchaseProductPropertyHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	purchase_product_propertyModel := models.PurchaseProductProperty{}
	c.Bind(&purchase_product_propertyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newPurchaseProductProperty := models.NewPurchaseProductPropertyRepository(db.DB)
	newPurchaseProductProperty.Create(&purchase_product_propertyModel)

setFlashmessages(c, "success", "PurchaseProductProperty created successfully!!")
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
	return renderView(c, purchase_product_property.PurchaseProductPropertyIndex(
		"| Create PurchaseProductProperty",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		purchase_product_property.CreatePurchaseProductProperty(),
	))

}

func purchase_product_propertyListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newPurchaseProductProperty := models.NewPurchaseProductPropertyRepository(db.DB)
	allPurchaseProductProperty, err := newPurchaseProductProperty.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's PurchaseProductProperty List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, purchase_product_property.PurchaseProductPropertyIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		purchase_product_property.PurchaseProductPropertyList(titlePage, allPurchaseProductProperty),
	))
}

func updatePurchaseProductPropertyHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newPurchaseProductProperty := models.NewPurchaseProductPropertyRepository(db.DB)
	
	purchase_product_propertyModel , err := newPurchaseProductProperty.GetSingle(idParams)

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
		//purchase_product_propertyModel := models.PurchaseProductProperty{}
	c.Bind(purchase_product_propertyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newPurchaseProductProperty := models.NewPurchaseProductPropertyRepository(db.DB)
	err = newPurchaseProductProperty.Update(purchase_product_propertyModel)

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


		setFlashmessages(c, "success", "PurchaseProductProperty successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/purchase_product_property/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, purchase_product_property.PurchaseProductPropertyIndex(
		fmt.Sprintf("| Edit PurchaseProductProperty #%d", purchase_product_propertyModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		purchase_product_property.UpdatePurchaseProductProperty(*purchase_product_propertyModel, tz.(string)),
	))

}

func deletePurchaseProductPropertyHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.PurchaseProductProperty{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newPurchaseProductProperty := models.NewPurchaseProductPropertyRepository(db.DB)
	err = newPurchaseProductProperty.Delete(idParams)
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

	setFlashmessages(c, "success", "PurchaseProductProperty successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/purchase_product_property/list")
}