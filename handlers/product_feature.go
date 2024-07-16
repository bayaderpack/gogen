package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/product_feature"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createProductFeatureHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	product_featureModel := models.ProductFeature{}
	c.Bind(&product_featureModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newProductFeature := models.NewProductFeatureRepository(db.DB)
	newProductFeature.Create(&product_featureModel)

setFlashmessages(c, "success", "ProductFeature created successfully!!")
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
	return renderView(c, product_feature.ProductFeatureIndex(
		"| Create ProductFeature",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_feature.CreateProductFeature(),
	))

}

func product_featureListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newProductFeature := models.NewProductFeatureRepository(db.DB)
	allProductFeature, err := newProductFeature.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's ProductFeature List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, product_feature.ProductFeatureIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		product_feature.ProductFeatureList(titlePage, allProductFeature),
	))
}

func updateProductFeatureHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newProductFeature := models.NewProductFeatureRepository(db.DB)
	
	product_featureModel , err := newProductFeature.GetSingle(idParams)

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
		//product_featureModel := models.ProductFeature{}
	c.Bind(product_featureModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newProductFeature := models.NewProductFeatureRepository(db.DB)
	err = newProductFeature.Update(product_featureModel)

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


		setFlashmessages(c, "success", "ProductFeature successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/product_feature/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, product_feature.ProductFeatureIndex(
		fmt.Sprintf("| Edit ProductFeature #%d", product_featureModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		product_feature.UpdateProductFeature(*product_featureModel, tz.(string)),
	))

}

func deleteProductFeatureHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.ProductFeature{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newProductFeature := models.NewProductFeatureRepository(db.DB)
	err = newProductFeature.Delete(idParams)
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

	setFlashmessages(c, "success", "ProductFeature successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/product_feature/list")
}
