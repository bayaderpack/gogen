package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_option"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductOptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_optionModel := models.ProductOption{}
	c.Bind(&product_optionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductOption := models.NewProductOptionRepository(db.DB)
	newProductOption.Create(&product_optionModel)

setFlashmessages(c, "success", "ProductOption created successfully!!")
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
	return renderView(c, product_option.ProductOptionIndex(
		"| Create ProductOption",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_option.CreateProductOption(),
	))

}

func product_optionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductOption := models.NewProductOptionRepository(db.DB)
	allProductOption, err := newProductOption.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductOption List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_option.ProductOptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_option.ProductOptionList(titlePage, allProductOption),
	))
}

func updateProductOptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductOption := models.NewProductOptionRepository(db.DB)
	
	product_optionModel , err := newProductOption.GetSingle(idParams)

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
		//product_optionModel := models.ProductOption{}
	c.Bind(product_optionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductOption := models.NewProductOptionRepository(db.DB)
	err = newProductOption.Update(product_optionModel)

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


		setFlashmessages(c, "success", "ProductOption successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_option/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_option.ProductOptionIndex(
		fmt.Sprintf("| Edit ProductOption #%d", product_optionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_option.UpdateProductOption(*product_optionModel, tz.(string)),
	))

}

func deleteProductOptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductOption{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductOption := models.NewProductOptionRepository(db.DB)
	err = newProductOption.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductOption successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_option/list")
}
