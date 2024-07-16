package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/purchase"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createPurchaseHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	purchaseModel := models.Purchase{}
	c.Bind(&purchaseModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newPurchase := models.NewPurchaseRepository(db.DB)
	newPurchase.Create(&purchaseModel)

setFlashmessages(c, "success", "Purchase created successfully!!")
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
	return renderView(c, purchase.PurchaseIndex(
		"| Create Purchase",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		purchase.CreatePurchase(),
	))

}

func purchaseListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newPurchase := models.NewPurchaseRepository(db.DB)
	allPurchase, err := newPurchase.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's Purchase List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, purchase.PurchaseIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		purchase.PurchaseList(titlePage, allPurchase),
	))
}

func updatePurchaseHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newPurchase := models.NewPurchaseRepository(db.DB)
	
	purchaseModel , err := newPurchase.GetSingle(idParams)

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
		//purchaseModel := models.Purchase{}
	c.Bind(purchaseModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newPurchase := models.NewPurchaseRepository(db.DB)
	err = newPurchase.Update(purchaseModel)

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


		setFlashmessages(c, "success", "Purchase successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/purchase/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, purchase.PurchaseIndex(
		fmt.Sprintf("| Edit Purchase #%d", purchaseModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		purchase.UpdatePurchase(*purchaseModel, tz.(string)),
	))

}

func deletePurchaseHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Purchase{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newPurchase := models.NewPurchaseRepository(db.DB)
	err = newPurchase.Delete(idParams)
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

	setFlashmessages(c, "success", "Purchase successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/purchase/list")
}
