package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/history"

	"github.com/gin-gonic/gin"

	// "golang.org/x/text/cases"
	// "golang.org/x/text/language"
	"gorm.io/gorm"
)

func CreateHistoryHandler(c *gin.Context) {
	isError = false
	if c.Request.Method == http.MethodPost {
		historyModel := models.History{}
		c.Bind(&historyModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newHistory := models.NewHistoryRepository(db.DB)
		newHistory.Create(&historyModel)

		setFlashmessages(c, "success", "History created successfully!!")
		c.JSON(http.StatusOK, gin.H{
			"Code":   200,
			"Status": "Ok",
			"Data":   nil,
		})
	}
	// username_key_value, ok  := c.Get(username_key)
	// if !ok {
	// 	fmt.Println("Some error")
	// }

	// c.String(200, history.CreateHistory(), gin.H{
	// 	"title": "Main website",
	// })
	Render(c, history.CreateHistory())
	//  renderView(c, history.HistoryIndex(
	// 	"| Create History",
	// 	username_key_value.(string),
	// 	fromProtected,
	// 	isError,
	// 	getFlashmessages(c, "error"),
	// 	getFlashmessages(c, "success"),
	// 	history.CreateHistory(),
	// ))

}

func HistoryListHandler(c *gin.Context) {
	isError = false
	// userId, ok := c.Get(user_id_key)
	// if !ok {
	// 	fmt.Println("Some error")
	// }
	newHistory := models.NewHistoryRepository(db.DB)
	fmt.Printf("%+v", newHistory)
	allHistory, err := newHistory.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's History List",
		"test",
		// cases.Title(language.English).String(userId.(string)),
	)
	// his := []models.History{}

	Render(c, history.HistoryIndex(
		titlePage,
		"bajro",
		// userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		history.HistoryList(titlePage, allHistory),
	))
}

func UpdateHistoryHandler(c *gin.Context) {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}
	newHistory := models.NewHistoryRepository(db.DB)

	historyModel, err := newHistory.GetSingle(idParams)

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
	if c.Request.Method == "POST" {
		//historyModel := models.History{}
		c.Bind(historyModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		//newHistory := models.NewHistoryRepository(db.DB)
		err = newHistory.Update(historyModel)

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

		setFlashmessages(c, "success", "History successfully updated!!")
	} else {

		//	return c.Redirect(http.StatusSeeOther, "/history/list")
		//}

		// username, _ := c.Get(username_key)
		// tz, _ := c.Get(tzone_key)
		Render(c, history.HistoryIndex(
			fmt.Sprintf("| Edit History #%d", historyModel),
			"najro",
			fromProtected,
			isError,
			getFlashmessages(c, "error"),
			getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
			history.UpdateHistory(*historyModel, "aa"),
		))
	}

}

func deleteHistoryHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.History{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newHistory := models.NewHistoryRepository(db.DB)
	err = newHistory.Delete(idParams)
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

	setFlashmessages(c, "success", "History successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/history/list")
}
