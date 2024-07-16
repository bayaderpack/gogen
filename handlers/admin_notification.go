package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"errors"

	"bajscheme/models"
	"bajscheme/db"
	"bajscheme/views/admin_notification"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)


func createAdminNotificationHandler(c *gin.Context) error {
isError = false
if c.Request.Method == http.MethodPost {
	admin_notificationModel := models.AdminNotification{}
	c.Bind(&admin_notificationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	newAdminNotification := models.NewAdminNotificationRepository(db.DB)
	newAdminNotification.Create(&admin_notificationModel)

setFlashmessages(c, "success", "AdminNotification created successfully!!")
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
	return renderView(c, admin_notification.AdminNotificationIndex(
		"| Create AdminNotification",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		admin_notification.CreateAdminNotification(),
	))

}

func admin_notificationListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newAdminNotification := models.NewAdminNotificationRepository(db.DB)
	allAdminNotification, err := newAdminNotification.GetAll()

	if err != nil {
	fmt.Println(err)
	}


	titlePage := fmt.Sprintf(
		"| %s's AdminNotification List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, admin_notification.AdminNotificationIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		admin_notification.AdminNotificationList(titlePage, allAdminNotification),
	))
}

func updateAdminNotificationHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newAdminNotification := models.NewAdminNotificationRepository(db.DB)
	
	admin_notificationModel , err := newAdminNotification.GetSingle(idParams)

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
		//admin_notificationModel := models.AdminNotification{}
	c.Bind(admin_notificationModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newAdminNotification := models.NewAdminNotificationRepository(db.DB)
	err = newAdminNotification.Update(admin_notificationModel)

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


		setFlashmessages(c, "success", "AdminNotification successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/admin_notification/list")
	//}

username, _ := c.Get(username_key)
tz, _ := c.Get(tzone_key)
		return renderView(c, admin_notification.AdminNotificationIndex(
		fmt.Sprintf("| Edit AdminNotification #%d", admin_notificationModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		admin_notification.UpdateAdminNotification(*admin_notificationModel, tz.(string)),
	))

}

func deleteAdminNotificationHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.AdminNotification{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newAdminNotification := models.NewAdminNotificationRepository(db.DB)
	err = newAdminNotification.Delete(idParams)
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

	setFlashmessages(c, "success", "AdminNotification successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/admin_notification/list")
}
