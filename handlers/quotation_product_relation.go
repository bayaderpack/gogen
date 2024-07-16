package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/quotation_product_relation"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createQuotationProductRelationHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	quotation_product_relationModel := models.QuotationProductRelation{}
	c.Bind(&quotation_product_relationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newQuotationProductRelation := models.NewQuotationProductRelationRepository(db.DB)
	newQuotationProductRelation.Create(&quotation_product_relationModel)

setFlashmessages(c, "success", "QuotationProductRelation created successfully!!")
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
	return renderView(c, quotation_product_relation.QuotationProductRelationIndex(
		"| Create QuotationProductRelation",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_relation.CreateQuotationProductRelation(),
	))

}

func quotation_product_relationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newQuotationProductRelation := models.NewQuotationProductRelationRepository(db.DB)
	allQuotationProductRelation, err := newQuotationProductRelation.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's QuotationProductRelation List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, quotation_product_relation.QuotationProductRelationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		quotation_product_relation.QuotationProductRelationList(titlePage, allQuotationProductRelation),
	))
}

func updateQuotationProductRelationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newQuotationProductRelation := models.NewQuotationProductRelationRepository(db.DB)
	
	quotation_product_relationModel , err := newQuotationProductRelation.GetSingle(idParams)

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
		//quotation_product_relationModel := models.QuotationProductRelation{}
	c.Bind(quotation_product_relationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newQuotationProductRelation := models.NewQuotationProductRelationRepository(db.DB)
	err = newQuotationProductRelation.Update(quotation_product_relationModel)

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


		setFlashmessages(c, "success", "QuotationProductRelation successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/quotation_product_relation/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, quotation_product_relation.QuotationProductRelationIndex(
		fmt.Sprintf("| Edit QuotationProductRelation #%d", quotation_product_relationModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		quotation_product_relation.UpdateQuotationProductRelation(*quotation_product_relationModel, tz.(string)),
	))

}

func deleteQuotationProductRelationHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.QuotationProductRelation{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newQuotationProductRelation := models.NewQuotationProductRelationRepository(db.DB)
	err = newQuotationProductRelation.Delete(idParams)
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

	setFlashmessages(c, "success", "QuotationProductRelation successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/quotation_product_relation/list")
}
