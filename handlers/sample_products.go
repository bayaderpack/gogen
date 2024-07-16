package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/sample_products"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSampleProductsHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	sample_productsModel := models.SampleProducts{}
	c.Bind(&sample_productsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSampleProducts := models.NewSampleProductsRepository(db.DB)
	newSampleProducts.Create(&sample_productsModel)

setFlashmessages(c, "success", "SampleProducts created successfully!!")
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
	return renderView(c, sample_products.SampleProductsIndex(
		"| Create SampleProducts",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_products.CreateSampleProducts(),
	))

}

func sample_productsListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSampleProducts := models.NewSampleProductsRepository(db.DB)
	allSampleProducts, err := newSampleProducts.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's SampleProducts List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, sample_products.SampleProductsIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_products.SampleProductsList(titlePage, allSampleProducts),
	))
}

func updateSampleProductsHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSampleProducts := models.NewSampleProductsRepository(db.DB)
	
	sample_productsModel , err := newSampleProducts.GetSingle(idParams)

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
		//sample_productsModel := models.SampleProducts{}
	c.Bind(sample_productsModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSampleProducts := models.NewSampleProductsRepository(db.DB)
	err = newSampleProducts.Update(sample_productsModel)

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


		setFlashmessages(c, "success", "SampleProducts successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/sample_products/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, sample_products.SampleProductsIndex(
		fmt.Sprintf("| Edit SampleProducts #%d", sample_productsModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		sample_products.UpdateSampleProducts(*sample_productsModel, tz.(string)),
	))

}

func deleteSampleProductsHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SampleProducts{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSampleProducts := models.NewSampleProductsRepository(db.DB)
	err = newSampleProducts.Delete(idParams)
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

	setFlashmessages(c, "success", "SampleProducts successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/sample_products/list")
}
