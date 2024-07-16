package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/coupon_product"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createCouponProductHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	coupon_productModel := models.CouponProduct{}
	c.Bind(&coupon_productModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newCouponProduct := models.NewCouponProductRepository(db.DB)
	newCouponProduct.Create(&coupon_productModel)

setFlashmessages(c, "success", "CouponProduct created successfully!!")
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
	return renderView(c, coupon_product.CouponProductIndex(
		"| Create CouponProduct",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_product.CreateCouponProduct(),
	))

}

func coupon_productListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCouponProduct := models.NewCouponProductRepository(db.DB)
	allCouponProduct, err := newCouponProduct.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's CouponProduct List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, coupon_product.CouponProductIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_product.CouponProductList(titlePage, allCouponProduct),
	))
}

func updateCouponProductHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCouponProduct := models.NewCouponProductRepository(db.DB)
	
	coupon_productModel , err := newCouponProduct.GetSingle(idParams)

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
		//coupon_productModel := models.CouponProduct{}
	c.Bind(coupon_productModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCouponProduct := models.NewCouponProductRepository(db.DB)
	err = newCouponProduct.Update(coupon_productModel)

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


		setFlashmessages(c, "success", "CouponProduct successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/coupon_product/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, coupon_product.CouponProductIndex(
		fmt.Sprintf("| Edit CouponProduct #%d", coupon_productModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		coupon_product.UpdateCouponProduct(*coupon_productModel, tz.(string)),
	))

}

func deleteCouponProductHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CouponProduct{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCouponProduct := models.NewCouponProductRepository(db.DB)
	err = newCouponProduct.Delete(idParams)
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

	setFlashmessages(c, "success", "CouponProduct successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/coupon_product/list")
}
