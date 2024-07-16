package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/gallery_description"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createGalleryDescriptionHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		gallery_descriptionModel := models.GalleryDescription{}
		c.Bind(&gallery_descriptionModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newGalleryDescription := models.NewGalleryDescriptionRepository(db.DB)
		newGalleryDescription.Create(&gallery_descriptionModel)

		setFlashmessages(c, "success", "GalleryDescription created successfully!!")
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
	return renderView(c, gallery_description.GalleryDescriptionIndex(
		"| Create GalleryDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery_description.CreateGalleryDescription(),
	))

}

func gallery_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newGalleryDescription := models.NewGalleryDescriptionRepository(db.DB)
	allGalleryDescription, err := newGalleryDescription.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's GalleryDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, gallery_description.GalleryDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery_description.GalleryDescriptionList(titlePage, allGalleryDescription),
	))
}

func updateGalleryDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newGalleryDescription := models.NewGalleryDescriptionRepository(db.DB)

	gallery_descriptionModel, err := newGalleryDescription.GetSingle(idParams)

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
	//gallery_descriptionModel := models.GalleryDescription{}
	c.Bind(gallery_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newGalleryDescription := models.NewGalleryDescriptionRepository(db.DB)
	err = newGalleryDescription.Update(gallery_descriptionModel)

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

	setFlashmessages(c, "success", "GalleryDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/gallery_description/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, gallery_description.GalleryDescriptionIndex(
		fmt.Sprintf("| Edit GalleryDescription #%d", gallery_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		gallery_description.UpdateGalleryDescription(*gallery_descriptionModel, tz.(string)),
	))

}

func deleteGalleryDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.GalleryDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newGalleryDescription := models.NewGalleryDescriptionRepository(db.DB)
	err = newGalleryDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "GalleryDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/gallery_description/list")
}
