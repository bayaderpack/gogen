package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/admin_role"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createAdminRoleHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		admin_roleModel := models.AdminRole{}
		c.Bind(&admin_roleModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newAdminRole := models.NewAdminRoleRepository(db.DB)
		newAdminRole.Create(&admin_roleModel)

		setFlashmessages(c, "success", "AdminRole created successfully!!")
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
	return renderView(c, admin_role.AdminRoleIndex(
		"| Create AdminRole",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		admin_role.CreateAdminRole(),
	))

}

func admin_roleListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newAdminRole := models.NewAdminRoleRepository(db.DB)
	allAdminRole, err := newAdminRole.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's AdminRole List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, admin_role.AdminRoleIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		admin_role.AdminRoleList(titlePage, allAdminRole),
	))
}

func updateAdminRoleHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newAdminRole := models.NewAdminRoleRepository(db.DB)

	admin_roleModel, err := newAdminRole.GetSingle(idParams)

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
	//admin_roleModel := models.AdminRole{}
	c.Bind(admin_roleModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newAdminRole := models.NewAdminRoleRepository(db.DB)
	err = newAdminRole.Update(admin_roleModel)

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

	setFlashmessages(c, "success", "AdminRole successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/admin_role/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, admin_role.AdminRoleIndex(
		fmt.Sprintf("| Edit AdminRole #%d", admin_roleModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		admin_role.UpdateAdminRole(*admin_roleModel, tz.(string)),
	))

}

func deleteAdminRoleHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.AdminRole{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newAdminRole := models.NewAdminRoleRepository(db.DB)
	err = newAdminRole.Delete(idParams)
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

	setFlashmessages(c, "success", "AdminRole successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/admin_role/list")
}
