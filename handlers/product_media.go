package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/product_media"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createProductMediaHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		product_mediaModel := models.ProductMedia{}
		c.Bind(&product_mediaModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newProductMedia := models.NewProductMediaRepository(db.DB)
		newProductMedia.Create(&product_mediaModel)

		setFlashmessages(c, "success", "ProductMedia created successfully!!")
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
	return renderView(c, product_media.ProductMediaIndex(
		"| Create ProductMedia",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_media.CreateProductMedia(),
	))

}

func product_mediaListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductMedia := models.NewProductMediaRepository(db.DB)
	allProductMedia, err := newProductMedia.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's ProductMedia List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_media.ProductMediaIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_media.ProductMediaList(titlePage, allProductMedia),
	))
}

func updateProductMediaHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductMedia := models.NewProductMediaRepository(db.DB)

	product_mediaModel, err := newProductMedia.GetSingle(idParams)

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
	//product_mediaModel := models.ProductMedia{}
	c.Bind(product_mediaModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductMedia := models.NewProductMediaRepository(db.DB)
	err = newProductMedia.Update(product_mediaModel)

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

	setFlashmessages(c, "success", "ProductMedia successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_media/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, product_media.ProductMediaIndex(
		fmt.Sprintf("| Edit ProductMedia #%d", product_mediaModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_media.UpdateProductMedia(*product_mediaModel, tz.(string)),
	))

}

func deleteProductMediaHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductMedia{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductMedia := models.NewProductMediaRepository(db.DB)
	err = newProductMedia.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductMedia successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_media/list")
}
