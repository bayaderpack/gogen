package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/gallery"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createGalleryHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	galleryModel := models.Gallery{}
	c.Bind(&galleryModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newGallery := models.NewGalleryRepository(db.DB)
	newGallery.Create(&galleryModel)

setFlashmessages(c, "success", "Gallery created successfully!!")
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
	return renderView(c, gallery.GalleryIndex(
		"| Create Gallery",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery.CreateGallery(),
	))

}

func galleryListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newGallery := models.NewGalleryRepository(db.DB)
	allGallery, err := newGallery.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's Gallery List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, gallery.GalleryIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		gallery.GalleryList(titlePage, allGallery),
	))
}

func updateGalleryHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newGallery := models.NewGalleryRepository(db.DB)
	
	galleryModel , err := newGallery.GetSingle(idParams)

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
		//galleryModel := models.Gallery{}
	c.Bind(galleryModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newGallery := models.NewGalleryRepository(db.DB)
	err = newGallery.Update(galleryModel)

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


		setFlashmessages(c, "success", "Gallery successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/gallery/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, gallery.GalleryIndex(
		fmt.Sprintf("| Edit Gallery #%d", galleryModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		gallery.UpdateGallery(*galleryModel, tz.(string)),
	))

}

func deleteGalleryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Gallery{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newGallery := models.NewGalleryRepository(db.DB)
	err = newGallery.Delete(idParams)
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

	setFlashmessages(c, "success", "Gallery successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/gallery/list")
}
