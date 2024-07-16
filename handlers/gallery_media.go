package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/gallery_media"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createGalleryMediaHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		gallery_mediaModel := models.GalleryMedia{}
		c.Bind(&gallery_mediaModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newGalleryMedia := models.NewGalleryMediaRepository(db.DB)
		newGalleryMedia.Create(&gallery_mediaModel)

		setFlashmessages(c, "success", "GalleryMedia created successfully!!")
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
	return renderView(c, gallery_media.GalleryMediaIndex(
		"| Create GalleryMedia",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery_media.CreateGalleryMedia(),
	))

}

func gallery_mediaListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newGalleryMedia := models.NewGalleryMediaRepository(db.DB)
	allGalleryMedia, err := newGalleryMedia.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's GalleryMedia List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, gallery_media.GalleryMediaIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery_media.GalleryMediaList(titlePage, allGalleryMedia),
	))
}

func updateGalleryMediaHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newGalleryMedia := models.NewGalleryMediaRepository(db.DB)

	gallery_mediaModel, err := newGalleryMedia.GetSingle(idParams)

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
	//gallery_mediaModel := models.GalleryMedia{}
	c.Bind(gallery_mediaModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newGalleryMedia := models.NewGalleryMediaRepository(db.DB)
	err = newGalleryMedia.Update(gallery_mediaModel)

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

	setFlashmessages(c, "success", "GalleryMedia successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/gallery_media/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, gallery_media.GalleryMediaIndex(
		fmt.Sprintf("| Edit GalleryMedia #%d", gallery_mediaModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		gallery_media.UpdateGalleryMedia(*gallery_mediaModel, tz.(string)),
	))

}

func deleteGalleryMediaHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.GalleryMedia{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newGalleryMedia := models.NewGalleryMediaRepository(db.DB)
	err = newGalleryMedia.Delete(idParams)
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

	setFlashmessages(c, "success", "GalleryMedia successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/gallery_media/list")
}
