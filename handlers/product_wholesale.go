package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/product_wholesale"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createProductWholesaleHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		product_wholesaleModel := models.ProductWholesale{}
		c.Bind(&product_wholesaleModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newProductWholesale := models.NewProductWholesaleRepository(db.DB)
		newProductWholesale.Create(&product_wholesaleModel)

		setFlashmessages(c, "success", "ProductWholesale created successfully!!")
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
	return renderView(c, product_wholesale.ProductWholesaleIndex(
		"| Create ProductWholesale",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_wholesale.CreateProductWholesale(),
	))

}

func product_wholesaleListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductWholesale := models.NewProductWholesaleRepository(db.DB)
	allProductWholesale, err := newProductWholesale.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's ProductWholesale List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_wholesale.ProductWholesaleIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_wholesale.ProductWholesaleList(titlePage, allProductWholesale),
	))
}

func updateProductWholesaleHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductWholesale := models.NewProductWholesaleRepository(db.DB)

	product_wholesaleModel, err := newProductWholesale.GetSingle(idParams)

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
	//product_wholesaleModel := models.ProductWholesale{}
	c.Bind(product_wholesaleModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductWholesale := models.NewProductWholesaleRepository(db.DB)
	err = newProductWholesale.Update(product_wholesaleModel)

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

	setFlashmessages(c, "success", "ProductWholesale successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_wholesale/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, product_wholesale.ProductWholesaleIndex(
		fmt.Sprintf("| Edit ProductWholesale #%d", product_wholesaleModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_wholesale.UpdateProductWholesale(*product_wholesaleModel, tz.(string)),
	))

}

func deleteProductWholesaleHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductWholesale{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductWholesale := models.NewProductWholesaleRepository(db.DB)
	err = newProductWholesale.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductWholesale successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_wholesale/list")
}
