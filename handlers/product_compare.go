package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_compare"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductCompareHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_compareModel := models.ProductCompare{}
	c.Bind(&product_compareModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductCompare := models.NewProductCompareRepository(db.DB)
	newProductCompare.Create(&product_compareModel)

setFlashmessages(c, "success", "ProductCompare created successfully!!")
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
	return renderView(c, product_compare.ProductCompareIndex(
		"| Create ProductCompare",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_compare.CreateProductCompare(),
	))

}

func product_compareListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductCompare := models.NewProductCompareRepository(db.DB)
	allProductCompare, err := newProductCompare.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductCompare List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_compare.ProductCompareIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_compare.ProductCompareList(titlePage, allProductCompare),
	))
}

func updateProductCompareHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductCompare := models.NewProductCompareRepository(db.DB)
	
	product_compareModel , err := newProductCompare.GetSingle(idParams)

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
		//product_compareModel := models.ProductCompare{}
	c.Bind(product_compareModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductCompare := models.NewProductCompareRepository(db.DB)
	err = newProductCompare.Update(product_compareModel)

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


		setFlashmessages(c, "success", "ProductCompare successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_compare/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_compare.ProductCompareIndex(
		fmt.Sprintf("| Edit ProductCompare #%d", product_compareModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_compare.UpdateProductCompare(*product_compareModel, tz.(string)),
	))

}

func deleteProductCompareHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductCompare{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductCompare := models.NewProductCompareRepository(db.DB)
	err = newProductCompare.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductCompare successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_compare/list")
}
