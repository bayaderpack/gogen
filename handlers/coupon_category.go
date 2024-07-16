package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/coupon_category"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createCouponCategoryHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		coupon_categoryModel := models.CouponCategory{}
		c.Bind(&coupon_categoryModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newCouponCategory := models.NewCouponCategoryRepository(db.DB)
		newCouponCategory.Create(&coupon_categoryModel)

		setFlashmessages(c, "success", "CouponCategory created successfully!!")
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
	return renderView(c, coupon_category.CouponCategoryIndex(
		"| Create CouponCategory",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_category.CreateCouponCategory(),
	))

}

func coupon_categoryListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCouponCategory := models.NewCouponCategoryRepository(db.DB)
	allCouponCategory, err := newCouponCategory.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's CouponCategory List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, coupon_category.CouponCategoryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_category.CouponCategoryList(titlePage, allCouponCategory),
	))
}

func updateCouponCategoryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCouponCategory := models.NewCouponCategoryRepository(db.DB)

	coupon_categoryModel, err := newCouponCategory.GetSingle(idParams)

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
	//coupon_categoryModel := models.CouponCategory{}
	c.Bind(coupon_categoryModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCouponCategory := models.NewCouponCategoryRepository(db.DB)
	err = newCouponCategory.Update(coupon_categoryModel)

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

	setFlashmessages(c, "success", "CouponCategory successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/coupon_category/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, coupon_category.CouponCategoryIndex(
		fmt.Sprintf("| Edit CouponCategory #%d", coupon_categoryModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		coupon_category.UpdateCouponCategory(*coupon_categoryModel, tz.(string)),
	))

}

func deleteCouponCategoryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CouponCategory{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCouponCategory := models.NewCouponCategoryRepository(db.DB)
	err = newCouponCategory.Delete(idParams)
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

	setFlashmessages(c, "success", "CouponCategory successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/coupon_category/list")
}
