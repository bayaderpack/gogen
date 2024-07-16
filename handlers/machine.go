package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"bajscheme/db"
	"bajscheme/models"
	"bajscheme/views/machine"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func createMachineHandler(c *gin.Context) error {
	isError = false
	if c.Request.Method == http.MethodPost {
		machineModel := models.Machine{}
		c.Bind(&machineModel)
		//err := ctx.ShouldBindJSON(&createTagRequest)
		//helper.ErrorPanic(err)

		newMachine := models.NewMachineRepository(db.DB)
		newMachine.Create(&machineModel)

		setFlashmessages(c, "success", "Machine created successfully!!")
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
	return renderView(c, machine.MachineIndex(
		"| Create Machine",
		username_key_value.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		machine.CreateMachine(),
	))

}

func machineListHandler(c *gin.Context) error {
	isError = false
	userId, ok := c.Get(user_id_key)
	if !ok {
		fmt.Println("Some error")
	}
	newMachine := models.NewMachineRepository(db.DB)
	allMachine, err := newMachine.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	titlePage := fmt.Sprintf(
		"| %s's Machine List",
		cases.Title(language.English).String(userId.(string)),
	)

	return renderView(c, machine.MachineIndex(
		titlePage,
		userId.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		machine.MachineList(titlePage, allMachine),
	))
}

func updateMachineHandler(c *gin.Context) error {
	isError = false

	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	newMachine := models.NewMachineRepository(db.DB)

	machineModel, err := newMachine.GetSingle(idParams)

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
	//machineModel := models.Machine{}
	c.Bind(machineModel)
	//err := ctx.ShouldBindJSON(&createTagRequest)
	//helper.ErrorPanic(err)

	//newMachine := models.NewMachineRepository(db.DB)
	err = newMachine.Update(machineModel)

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

	setFlashmessages(c, "success", "Machine successfully updated!!")

	//	return c.Redirect(http.StatusSeeOther, "/machine/list")
	//}

	username, _ := c.Get(username_key)
	tz, _ := c.Get(tzone_key)
	return renderView(c, machine.MachineIndex(
		fmt.Sprintf("| Edit Machine #%d", machineModel),
		username.(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"), // ↓ getting time zone from context ↓
		machine.UpdateMachine(*machineModel, tz.(string)),
	))

}

func deleteMachineHandler(c *gin.Context) {
	idParams, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		fmt.Println(err)
	}

	//t := models.Machine{
	//	CreatedBy: c.Get(user_id_key).(int),
	//	ID:        idParams,
	//}

	newMachine := models.NewMachineRepository(db.DB)
	err = newMachine.Delete(idParams)
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

	setFlashmessages(c, "success", "Machine successfully deleted!!")

	c.Redirect(http.StatusSeeOther, "/machine/list")
}
