package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/media_variation"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createMediaVariationHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		media_variationModel := models.MediaVariation{}
		c.Bind(&media_variationModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newMediaVariation := models.NewMediaVariationRepository(db.DB)
		newMediaVariation.Create(&media_variationModel)

		setFlashmessages(c, "success", "MediaVariation created successfully!!")
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
	return renderView(c, media_variation.MediaVariationIndex(
		"| Create MediaVariation",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_variation.CreateMediaVariation(),
	))

}

func media_variationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMediaVariation := models.NewMediaVariationRepository(db.DB)
	allMediaVariation, err := newMediaVariation.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's MediaVariation List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, media_variation.MediaVariationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_variation.MediaVariationList(titlePage, allMediaVariation),
	))
}

func updateMediaVariationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMediaVariation := models.NewMediaVariationRepository(db.DB)

	media_variationModel, err := newMediaVariation.GetSingle(idParams)

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
	//media_variationModel := models.MediaVariation{}
	c.Bind(media_variationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMediaVariation := models.NewMediaVariationRepository(db.DB)
	err = newMediaVariation.Update(media_variationModel)

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

	setFlashmessages(c, "success", "MediaVariation successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/media_variation/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, media_variation.MediaVariationIndex(
		fmt.Sprintf("| Edit MediaVariation #%d", media_variationModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		media_variation.UpdateMediaVariation(*media_variationModel, tz.(string)),
	))

}

func deleteMediaVariationHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MediaVariation{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMediaVariation := models.NewMediaVariationRepository(db.DB)
	err = newMediaVariation.Delete(idParams)
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

	setFlashmessages(c, "success", "MediaVariation successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/media_variation/list")
}
