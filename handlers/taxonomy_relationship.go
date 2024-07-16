package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/taxonomy_relationship"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createTaxonomyRelationshipHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		taxonomy_relationshipModel := models.TaxonomyRelationship{}
		c.Bind(&taxonomy_relationshipModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newTaxonomyRelationship := models.NewTaxonomyRelationshipRepository(db.DB)
		newTaxonomyRelationship.Create(&taxonomy_relationshipModel)

		setFlashmessages(c, "success", "TaxonomyRelationship created successfully!!")
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
	return renderView(c, taxonomy_relationship.TaxonomyRelationshipIndex(
		"| Create TaxonomyRelationship",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy_relationship.CreateTaxonomyRelationship(),
	))

}

func taxonomy_relationshipListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newTaxonomyRelationship := models.NewTaxonomyRelationshipRepository(db.DB)
	allTaxonomyRelationship, err := newTaxonomyRelationship.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's TaxonomyRelationship List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, taxonomy_relationship.TaxonomyRelationshipIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		taxonomy_relationship.TaxonomyRelationshipList(titlePage, allTaxonomyRelationship),
	))
}

func updateTaxonomyRelationshipHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newTaxonomyRelationship := models.NewTaxonomyRelationshipRepository(db.DB)

	taxonomy_relationshipModel, err := newTaxonomyRelationship.GetSingle(idParams)

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
	//taxonomy_relationshipModel := models.TaxonomyRelationship{}
	c.Bind(taxonomy_relationshipModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newTaxonomyRelationship := models.NewTaxonomyRelationshipRepository(db.DB)
	err = newTaxonomyRelationship.Update(taxonomy_relationshipModel)

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

	setFlashmessages(c, "success", "TaxonomyRelationship successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/taxonomy_relationship/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, taxonomy_relationship.TaxonomyRelationshipIndex(
		fmt.Sprintf("| Edit TaxonomyRelationship #%d", taxonomy_relationshipModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		taxonomy_relationship.UpdateTaxonomyRelationship(*taxonomy_relationshipModel, tz.(string)),
	))

}

func deleteTaxonomyRelationshipHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.TaxonomyRelationship{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newTaxonomyRelationship := models.NewTaxonomyRelationshipRepository(db.DB)
	err = newTaxonomyRelationship.Delete(idParams)
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

	setFlashmessages(c, "success", "TaxonomyRelationship successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/taxonomy_relationship/list")
}
