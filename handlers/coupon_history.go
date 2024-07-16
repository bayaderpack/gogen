package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/coupon_history"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createCouponHistoryHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		coupon_historyModel := models.CouponHistory{}
		c.Bind(&coupon_historyModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newCouponHistory := models.NewCouponHistoryRepository(db.DB)
		newCouponHistory.Create(&coupon_historyModel)

		setFlashmessages(c, "success", "CouponHistory created successfully!!")
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
	return renderView(c, coupon_history.CouponHistoryIndex(
		"| Create CouponHistory",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_history.CreateCouponHistory(),
	))

}

func coupon_historyListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newCouponHistory := models.NewCouponHistoryRepository(db.DB)
	allCouponHistory, err := newCouponHistory.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's CouponHistory List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, coupon_history.CouponHistoryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		coupon_history.CouponHistoryList(titlePage, allCouponHistory),
	))
}

func updateCouponHistoryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newCouponHistory := models.NewCouponHistoryRepository(db.DB)

	coupon_historyModel, err := newCouponHistory.GetSingle(idParams)

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
	//coupon_historyModel := models.CouponHistory{}
	c.Bind(coupon_historyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newCouponHistory := models.NewCouponHistoryRepository(db.DB)
	err = newCouponHistory.Update(coupon_historyModel)

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

	setFlashmessages(c, "success", "CouponHistory successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/coupon_history/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, coupon_history.CouponHistoryIndex(
		fmt.Sprintf("| Edit CouponHistory #%d", coupon_historyModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		coupon_history.UpdateCouponHistory(*coupon_historyModel, tz.(string)),
	))

}

func deleteCouponHistoryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.CouponHistory{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newCouponHistory := models.NewCouponHistoryRepository(db.DB)
	err = newCouponHistory.Delete(idParams)
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

	setFlashmessages(c, "success", "CouponHistory successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/coupon_history/list")
}
