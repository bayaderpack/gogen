package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/language"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	lang "golang.org/x/text/language"
	"gorm.io/gorm"
)


func createLanguageHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	languageModel := models.Language{}
	c.Bind(&languageModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newLanguage := models.NewLanguageRepository(db.DB)
	newLanguage.Create(&languageModel)

setFlashmessages(c, "success", "Language created successfully!!")
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
	return renderView(c, language.LanguageIndex(
		"| Create Language",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		language.CreateLanguage(),
	))

}

func languageListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newLanguage := models.NewLanguageRepository(db.DB)
	allLanguage, err := newLanguage.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's Language List",
		cases.Title(lang.English).String(userId.(string)),
	)

	return renderView(c, language.LanguageIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		language.LanguageList(titlePage, allLanguage),
	))
}

func updateLanguageHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newLanguage := models.NewLanguageRepository(db.DB)
	
	languageModel , err := newLanguage.GetSingle(idParams)

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
		//languageModel := models.Language{}
	c.Bind(languageModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newLanguage := models.NewLanguageRepository(db.DB)
	err = newLanguage.Update(languageModel)

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


		setFlashmessages(c, "success", "Language successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/language/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, language.LanguageIndex(
		fmt.Sprintf("| Edit Language #%d", languageModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		language.UpdateLanguage(*languageModel, tz.(string)),
	))

}

func deleteLanguageHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Language{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newLanguage := models.NewLanguageRepository(db.DB)
	err = newLanguage.Delete(idParams)
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

	setFlashmessages(c, "success", "Language successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/language/list")
}
