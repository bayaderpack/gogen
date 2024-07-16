package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/media_favorite"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createMediaFavoriteHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	media_favoriteModel := models.MediaFavorite{}
	c.Bind(&media_favoriteModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newMediaFavorite := models.NewMediaFavoriteRepository(db.DB)
	newMediaFavorite.Create(&media_favoriteModel)

setFlashmessages(c, "success", "MediaFavorite created successfully!!")
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
	return renderView(c, media_favorite.MediaFavoriteIndex(
		"| Create MediaFavorite",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_favorite.CreateMediaFavorite(),
	))

}

func media_favoriteListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMediaFavorite := models.NewMediaFavoriteRepository(db.DB)
	allMediaFavorite, err := newMediaFavorite.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's MediaFavorite List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, media_favorite.MediaFavoriteIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		media_favorite.MediaFavoriteList(titlePage, allMediaFavorite),
	))
}

func updateMediaFavoriteHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMediaFavorite := models.NewMediaFavoriteRepository(db.DB)
	
	media_favoriteModel , err := newMediaFavorite.GetSingle(idParams)

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
		//media_favoriteModel := models.MediaFavorite{}
	c.Bind(media_favoriteModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMediaFavorite := models.NewMediaFavoriteRepository(db.DB)
	err = newMediaFavorite.Update(media_favoriteModel)

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


		setFlashmessages(c, "success", "MediaFavorite successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/media_favorite/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, media_favorite.MediaFavoriteIndex(
		fmt.Sprintf("| Edit MediaFavorite #%d", media_favoriteModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		media_favorite.UpdateMediaFavorite(*media_favoriteModel, tz.(string)),
	))

}

func deleteMediaFavoriteHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.MediaFavorite{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMediaFavorite := models.NewMediaFavoriteRepository(db.DB)
	err = newMediaFavorite.Delete(idParams)
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

	setFlashmessages(c, "success", "MediaFavorite successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/media_favorite/list")
}
