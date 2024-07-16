package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/taxonomy_description"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createTaxonomyDescriptionHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	taxonomy_descriptionModel := models.TaxonomyDescription{}
	c.Bind(&taxonomy_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newTaxonomyDescription := models.NewTaxonomyDescriptionRepository(db.DB)
	newTaxonomyDescription.Create(&taxonomy_descriptionModel)

setFlashmessages(c, "success", "TaxonomyDescription created successfully!!")
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
	return renderView(c, taxonomy_description.TaxonomyDescriptionIndex(
		"| Create TaxonomyDescription",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy_description.CreateTaxonomyDescription(),
	))

}

func taxonomy_descriptionListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newTaxonomyDescription := models.NewTaxonomyDescriptionRepository(db.DB)
	allTaxonomyDescription, err := newTaxonomyDescription.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's TaxonomyDescription List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, taxonomy_description.TaxonomyDescriptionIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy_description.TaxonomyDescriptionList(titlePage, allTaxonomyDescription),
	))
}

func updateTaxonomyDescriptionHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newTaxonomyDescription := models.NewTaxonomyDescriptionRepository(db.DB)
	
	taxonomy_descriptionModel , err := newTaxonomyDescription.GetSingle(idParams)

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
		//taxonomy_descriptionModel := models.TaxonomyDescription{}
	c.Bind(taxonomy_descriptionModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newTaxonomyDescription := models.NewTaxonomyDescriptionRepository(db.DB)
	err = newTaxonomyDescription.Update(taxonomy_descriptionModel)

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


		setFlashmessages(c, "success", "TaxonomyDescription successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/taxonomy_description/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, taxonomy_description.TaxonomyDescriptionIndex(
		fmt.Sprintf("| Edit TaxonomyDescription #%d", taxonomy_descriptionModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		taxonomy_description.UpdateTaxonomyDescription(*taxonomy_descriptionModel, tz.(string)),
	))

}

func deleteTaxonomyDescriptionHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.TaxonomyDescription{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newTaxonomyDescription := models.NewTaxonomyDescriptionRepository(db.DB)
	err = newTaxonomyDescription.Delete(idParams)
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

	setFlashmessages(c, "success", "TaxonomyDescription successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/taxonomy_description/list")
}
