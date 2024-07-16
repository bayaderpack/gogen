package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_option_value"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductOptionValueHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_option_valueModel := models.ProductOptionValue{}
	c.Bind(&product_option_valueModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductOptionValue := models.NewProductOptionValueRepository(db.DB)
	newProductOptionValue.Create(&product_option_valueModel)

setFlashmessages(c, "success", "ProductOptionValue created successfully!!")
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
	return renderView(c, product_option_value.ProductOptionValueIndex(
		"| Create ProductOptionValue",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_option_value.CreateProductOptionValue(),
	))

}

func product_option_valueListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductOptionValue := models.NewProductOptionValueRepository(db.DB)
	allProductOptionValue, err := newProductOptionValue.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductOptionValue List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_option_value.ProductOptionValueIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_option_value.ProductOptionValueList(titlePage, allProductOptionValue),
	))
}

func updateProductOptionValueHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductOptionValue := models.NewProductOptionValueRepository(db.DB)
	
	product_option_valueModel , err := newProductOptionValue.GetSingle(idParams)

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
		//product_option_valueModel := models.ProductOptionValue{}
	c.Bind(product_option_valueModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductOptionValue := models.NewProductOptionValueRepository(db.DB)
	err = newProductOptionValue.Update(product_option_valueModel)

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


		setFlashmessages(c, "success", "ProductOptionValue successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_option_value/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_option_value.ProductOptionValueIndex(
		fmt.Sprintf("| Edit ProductOptionValue #%d", product_option_valueModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_option_value.UpdateProductOptionValue(*product_option_valueModel, tz.(string)),
	))

}

func deleteProductOptionValueHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductOptionValue{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductOptionValue := models.NewProductOptionValueRepository(db.DB)
	err = newProductOptionValue.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductOptionValue successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_option_value/list")
}
