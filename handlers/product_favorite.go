package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_favorite"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductFavoriteHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_favoriteModel := models.ProductFavorite{}
	c.Bind(&product_favoriteModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductFavorite := models.NewProductFavoriteRepository(db.DB)
	newProductFavorite.Create(&product_favoriteModel)

setFlashmessages(c, "success", "ProductFavorite created successfully!!")
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
	return renderView(c, product_favorite.ProductFavoriteIndex(
		"| Create ProductFavorite",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_favorite.CreateProductFavorite(),
	))

}

func product_favoriteListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductFavorite := models.NewProductFavoriteRepository(db.DB)
	allProductFavorite, err := newProductFavorite.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductFavorite List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_favorite.ProductFavoriteIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_favorite.ProductFavoriteList(titlePage, allProductFavorite),
	))
}

func updateProductFavoriteHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductFavorite := models.NewProductFavoriteRepository(db.DB)
	
	product_favoriteModel , err := newProductFavorite.GetSingle(idParams)

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
		//product_favoriteModel := models.ProductFavorite{}
	c.Bind(product_favoriteModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductFavorite := models.NewProductFavoriteRepository(db.DB)
	err = newProductFavorite.Update(product_favoriteModel)

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


		setFlashmessages(c, "success", "ProductFavorite successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_favorite/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_favorite.ProductFavoriteIndex(
		fmt.Sprintf("| Edit ProductFavorite #%d", product_favoriteModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_favorite.UpdateProductFavorite(*product_favoriteModel, tz.(string)),
	))

}

func deleteProductFavoriteHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductFavorite{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductFavorite := models.NewProductFavoriteRepository(db.DB)
	err = newProductFavorite.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductFavorite successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_favorite/list")
}
