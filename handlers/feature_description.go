package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/feature_description"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createFeatureDescriptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	feature_descriptionModel := models.FeatureDescription{}
	c.Bind(&feature_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newFeatureDescription := models.NewFeatureDescriptionRepository(db.DB)
	newFeatureDescription.Create(&feature_descriptionModel)

setFlashmessages(c, "success", "FeatureDescription created successfully!!")
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
	return renderView(c, feature_description.FeatureDescriptionIndex(
		"| Create FeatureDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		feature_description.CreateFeatureDescription(),
	))

}

func feature_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newFeatureDescription := models.NewFeatureDescriptionRepository(db.DB)
	allFeatureDescription, err := newFeatureDescription.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's FeatureDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, feature_description.FeatureDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		feature_description.FeatureDescriptionList(titlePage, allFeatureDescription),
	))
}

func updateFeatureDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newFeatureDescription := models.NewFeatureDescriptionRepository(db.DB)
	
	feature_descriptionModel , err := newFeatureDescription.GetSingle(idParams)

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
		//feature_descriptionModel := models.FeatureDescription{}
	c.Bind(feature_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newFeatureDescription := models.NewFeatureDescriptionRepository(db.DB)
	err = newFeatureDescription.Update(feature_descriptionModel)

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


		setFlashmessages(c, "success", "FeatureDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/feature_description/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, feature_description.FeatureDescriptionIndex(
		fmt.Sprintf("| Edit FeatureDescription #%d", feature_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		feature_description.UpdateFeatureDescription(*feature_descriptionModel, tz.(string)),
	))

}

func deleteFeatureDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.FeatureDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newFeatureDescription := models.NewFeatureDescriptionRepository(db.DB)
	err = newFeatureDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "FeatureDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/feature_description/list")
}
