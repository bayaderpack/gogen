package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/sample_quotation"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createSampleQuotationHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	sample_quotationModel := models.SampleQuotation{}
	c.Bind(&sample_quotationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newSampleQuotation := models.NewSampleQuotationRepository(db.DB)
	newSampleQuotation.Create(&sample_quotationModel)

setFlashmessages(c, "success", "SampleQuotation created successfully!!")
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
	return renderView(c, sample_quotation.SampleQuotationIndex(
		"| Create SampleQuotation",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_quotation.CreateSampleQuotation(),
	))

}

func sample_quotationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newSampleQuotation := models.NewSampleQuotationRepository(db.DB)
	allSampleQuotation, err := newSampleQuotation.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's SampleQuotation List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, sample_quotation.SampleQuotationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		sample_quotation.SampleQuotationList(titlePage, allSampleQuotation),
	))
}

func updateSampleQuotationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newSampleQuotation := models.NewSampleQuotationRepository(db.DB)
	
	sample_quotationModel , err := newSampleQuotation.GetSingle(idParams)

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
		//sample_quotationModel := models.SampleQuotation{}
	c.Bind(sample_quotationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newSampleQuotation := models.NewSampleQuotationRepository(db.DB)
	err = newSampleQuotation.Update(sample_quotationModel)

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


		setFlashmessages(c, "success", "SampleQuotation successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/sample_quotation/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, sample_quotation.SampleQuotationIndex(
		fmt.Sprintf("| Edit SampleQuotation #%d", sample_quotationModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		sample_quotation.UpdateSampleQuotation(*sample_quotationModel, tz.(string)),
	))

}

func deleteSampleQuotationHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.SampleQuotation{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newSampleQuotation := models.NewSampleQuotationRepository(db.DB)
	err = newSampleQuotation.Delete(idParams)
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

	setFlashmessages(c, "success", "SampleQuotation successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/sample_quotation/list")
}
