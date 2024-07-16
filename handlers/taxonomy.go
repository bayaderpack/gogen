package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/taxonomy"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createTaxonomyHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	taxonomyModel := models.Taxonomy{}
	c.Bind(&taxonomyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newTaxonomy := models.NewTaxonomyRepository(db.DB)
	newTaxonomy.Create(&taxonomyModel)

setFlashmessages(c, "success", "Taxonomy created successfully!!")
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
	return renderView(c, taxonomy.TaxonomyIndex(
		"| Create Taxonomy",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy.CreateTaxonomy(),
	))

}

func taxonomyListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newTaxonomy := models.NewTaxonomyRepository(db.DB)
	allTaxonomy, err := newTaxonomy.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's Taxonomy List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, taxonomy.TaxonomyIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy.TaxonomyList(titlePage, allTaxonomy),
	))
}

func updateTaxonomyHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newTaxonomy := models.NewTaxonomyRepository(db.DB)
	
	taxonomyModel , err := newTaxonomy.GetSingle(idParams)

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
		//taxonomyModel := models.Taxonomy{}
	c.Bind(taxonomyModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newTaxonomy := models.NewTaxonomyRepository(db.DB)
	err = newTaxonomy.Update(taxonomyModel)

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


		setFlashmessages(c, "success", "Taxonomy successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/taxonomy/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, taxonomy.TaxonomyIndex(
		fmt.Sprintf("| Edit Taxonomy #%d", taxonomyModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		taxonomy.UpdateTaxonomy(*taxonomyModel, tz.(string)),
	))

}

func deleteTaxonomyHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Taxonomy{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newTaxonomy := models.NewTaxonomyRepository(db.DB)
	err = newTaxonomy.Delete(idParams)
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

	setFlashmessages(c, "success", "Taxonomy successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/taxonomy/list")
}
