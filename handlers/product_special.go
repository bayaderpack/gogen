package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_special"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductSpecialHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_specialModel := models.ProductSpecial{}
	c.Bind(&product_specialModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductSpecial := models.NewProductSpecialRepository(db.DB)
	newProductSpecial.Create(&product_specialModel)

setFlashmessages(c, "success", "ProductSpecial created successfully!!")
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
	return renderView(c, product_special.ProductSpecialIndex(
		"| Create ProductSpecial",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_special.CreateProductSpecial(),
	))

}

func product_specialListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductSpecial := models.NewProductSpecialRepository(db.DB)
	allProductSpecial, err := newProductSpecial.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductSpecial List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_special.ProductSpecialIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_special.ProductSpecialList(titlePage, allProductSpecial),
	))
}

func updateProductSpecialHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductSpecial := models.NewProductSpecialRepository(db.DB)
	
	product_specialModel , err := newProductSpecial.GetSingle(idParams)

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
		//product_specialModel := models.ProductSpecial{}
	c.Bind(product_specialModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductSpecial := models.NewProductSpecialRepository(db.DB)
	err = newProductSpecial.Update(product_specialModel)

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


		setFlashmessages(c, "success", "ProductSpecial successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_special/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_special.ProductSpecialIndex(
		fmt.Sprintf("| Edit ProductSpecial #%d", product_specialModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_special.UpdateProductSpecial(*product_specialModel, tz.(string)),
	))

}

func deleteProductSpecialHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductSpecial{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductSpecial := models.NewProductSpecialRepository(db.DB)
	err = newProductSpecial.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductSpecial successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_special/list")
}
