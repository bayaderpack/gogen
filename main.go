package main

import (
	// "encoding/json"
	"bajscheme/handlers"
	"fmt"
	"os"
	"regexp"
	"strings"

	// "time"

	// "bajscheme/auth"

	pluralize "github.com/gertd/go-pluralize"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	// "gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "github.com/gin-gonic/gin"
	// "bajscheme/handlers"
)

var dbName = "bayaderpack"

type ColumnSchema struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Nullable string `json:"nullable,omitempty"`
	Default  string `json:"default,omitempty"`
}

type TableSchema struct {
	Name    string         `json:"name"`
	Columns []ColumnSchema `json:"columns"`
}

// var db *gorm.DB

func main() {

	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")
	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/index", handlers.HistoryListHandler)
	router.GET("/quotation/edit/:id", handlers.UpdateQuotationProductsHandler)
	router.Run(":9092")
	// // timeNow := time.Now()
	// dsn := "root:@tcp(127.0.0.1:3307)/bayaderpack?charset=utf8mb4&parseTime=True&loc=Local"
	// db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// // if err != nil {
	// // 	panic(err.Error())
	// // }
	// // r := gin.Default()
	// // r.GET("/ping", handlers.HistoryListHandler)
	// // r.Run(":9090")
	// var tableNames []string
	// db.Table("information_schema.tables").Where("table_schema = ?", dbName).Pluck("table_name", &tableNames)

	// for _, tableName := range tableNames {
	// 	columns := getTableColumnsWithDetails(db, tableName)
	// 	tableSchema := TableSchema{Name: tableName, Columns: []ColumnSchema{}}
	// 	columnsFileds := ""
	// 	headerFields := ""
	// 	tableFields := ""
	// 	formFields := ""
	// 	editFields := ""
	// 	firstField := ""
	// 	isItFirst := 0
	// 	needsTime := false
	// 	for _, column := range columns {

	// 		// fmt.Printf("%s ---- %s", column["column_name"].(string), column["column_type"].(string))
	// 		columnName, ok := column["column_name"].(string)
	// 		if ok {
	// 			tableSchema.Columns = append(tableSchema.Columns, ColumnSchema{
	// 				Name:     columnName,
	// 				Type:     column["column_type"].(string),
	// 				Nullable: column["is_nullable"].(string),
	// 			})
	// 			if columnName == "id" || columnName == "date_added" || columnName == "date_modified" || columnName == "created_at" || columnName == "updated_at" {
	// 				continue
	// 			}
	// 			columnType := column["column_type"].(string)

	// 			if strings.Contains(column["column_type"].(string), "tinyint") {
	// 				columnType = "bool"
	// 			} else if strings.Contains(column["column_type"].(string), "varchar") || strings.Contains(column["column_type"].(string), "text") {
	// 				columnType = "string"
	// 			} else if strings.Contains(column["column_type"].(string), "int") {
	// 				columnType = "int"
	// 			} else if strings.Contains(column["column_type"].(string), "datetime") || strings.Contains(column["column_type"].(string), "timestamp") {
	// 				columnType = "*time.Time"
	// 			} else if strings.Contains(column["column_type"].(string), "decimal") {
	// 				columnType = "float32"
	// 			}

	// 			if strings.Contains(column["column_type"].(string), "time") {
	// 				needsTime = true
	// 			}
	// 			if defaultVal, ok := column["column_default"]; ok {
	// 				if defaultVal != nil {
	// 					tableSchema.Columns[len(tableSchema.Columns)-1].Default = defaultVal.(string)
	// 					nameINeed := toCamelCase(columnName)
	// 					// fmt.Println(strings.Contains(nameINeed, "Id"))

	// 					if strings.Contains(nameINeed, "Id") {
	// 						nameINeed = strings.ReplaceAll(nameINeed, "Id", "ID")
	// 					}

	// 					// fmt.Println(defaultVal.(string))

	// 					columnsFileds += fmt.Sprintf("%s %s `gorm:\"column:%s;default:%s\" json:\"%s\"`\n", nameINeed, columnType, columnName, defaultVal.(string), columnName)
	// 					headerFields += fmt.Sprintf("<th class=\"text-center\">%s</th>\n", capitalizeWords(splitWords(columnName)))
	// 					tableFields += fmt.Sprintf(`<td>{fmt.Sprintf("%%s",%s.%s)} </td>\n`, tableName, relaceID(toCamelCase(columnName)))
	// 					if isItFirst == 0 {
	// 						firstField = relaceID(fmt.Sprintf("%s.%s",tableName, relaceID(toCamelCase(columnName))))
	// 					}
	// 					formFields += generateFieldsTypes(column, columnName, defaultVal.(string), column["is_nullable"].(string))

	// 					editFields += generateEditFields(tableName, column, columnName, defaultVal.(string), column["is_nullable"].(string))

	// 				} else {
	// 					nameINeed := toCamelCase(columnName)

	// 					// if strings.Contains(nameINeed, "Id") {
	// 					nameINeed = strings.Replace(nameINeed, "Id", "ID", -1)
	// 					// }
	// 					tableSchema.Columns[len(tableSchema.Columns)-1].Default = "nil"
	// 					if isItFirst == 0 {
	// 						firstField = relaceID(fmt.Sprintf("%s.%s",tableName, relaceID(toCamelCase(columnName))))
	// 					}
	// 					columnsFileds += fmt.Sprintf("%s %s `gorm:\"column:%s;default:%s\" json:\"%s\"`\n", nameINeed, columnType, columnName, "", columnName)

	// 					formFields += generateFieldsTypes(column, columnName, "", column["is_nullable"].(string))
	// 					editFields += generateEditFields(tableName, column, columnName, "", column["is_nullable"].(string))

	// 					headerFields += fmt.Sprintf("<th class=\"text-center\">%s</th>\n", capitalizeWords(splitWords(columnName)))
	// 					tableFields += fmt.Sprintf(`<td>{fmt.Sprintf("%%s",%s.%s)} </td>\n`, tableName, relaceID(toCamelCase(columnName)))

	// 				}
	// 			}
	// 			isItFirst += 1
	// 		} else {
	// 			continue
	// 		}

	// 	}

	// 	generateViewFolders(tableName)
	// 	generateModelTemplate(tableName, needsTime, columnsFileds)
	// 	// 	auth.GenerateToken()
	// 	generateIndexFiles(tableName, headerFields, firstField, tableFields)
	// 	generateCreateFiles(tableName, formFields)
	// 	generateEditFiles(tableName, editFields)
	// 	generateHandlersTemplate(tableName)

	// }

	// fmt.Println("Project generated it takes ", time.Since(timeNow))
}

func relaceID(field string) string {
	if strings.Contains(field, "Id") {
		field = strings.ReplaceAll(field, "Id", "ID")
		return field
	}
	return field
}

func getTableColumnsWithDetails(db *gorm.DB, tableName string) []map[string]interface{} {
	var columns []map[string]interface{}
	db.Table("information_schema.columns").
		Where("table_schema = ? AND table_name = ?", dbName, tableName).
		Select("column_name, column_type, is_nullable, column_default").
		Scan(&columns)
	return columns
}

func generateFieldsTypes(column map[string]interface{}, columnName string, defaultValue string, is_nullable string) string {

	isItRequired := "required"

	// fmt.Println("default value: ",defaultValue, "----- is nullable: ", is_nullable, " is condition true", (defaultValue == "NULL" || is_nullable == "YES"))
	formFields := ""
	if defaultValue == "NULL" || is_nullable == "YES" {
		isItRequired = ""
	}
	if strings.Contains(column["column_type"].(string), "tinyint") {

		formFields = fmt.Sprintf(`<div class="form-control">
	  <label class="label cursor-pointer">
		<span class="label-text">%s</span> 
		<input type="checkbox" name="%s" class="checkbox" %s />
	  </label>
	</div>
		`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
	} else if strings.Contains(column["column_type"].(string), "varchar") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input
			class="input input-bordered input-primary"
			type="text"
			name="%s"
			%s
			autofocus
			minlength="3"
			maxlength="64"
		/>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
	} else if strings.Contains(column["column_type"].(string), "text") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<textarea
			class="textarea textarea-primary h-36 max-h-36"
			%s
			name="%s"
			maxlength="255"
		></textarea>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
	} else if strings.Contains(column["column_type"].(string), "int") || strings.Contains(column["column_type"].(string), "float32") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input
			class="input input-bordered input-primary"
			type="number"
			name="%s"
			%s
			autofocus
			minlength="3"
			maxlength="64"
		/>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
	} else if strings.Contains(column["column_type"].(string), "enum") {
		// fmt.Printf("%s", )
		re := regexp.MustCompile(`'([^']*)'`)
		matches := re.FindAllString(column["column_type"].(string), -1)
		formFields += fmt.Sprintf(`<label class="form-control w-full max-w-xs">
		<div class="label">
		  <span class="label-text">%s</span>

		</div>
		<select class="select select-bordered" name="%s" %s>
		  <option disabled selected>Select one</option>

`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
		for _, match := range matches {
			// fmt.Println(match[1 : len(match)-1])
			formFields += fmt.Sprintf(`<option>%s</option>
			`, capitalizeWords(splitWords(match[1:len(match)-1])))
		}
		formFields += `</select>

		</label>`

	} else if strings.Contains(column["column_type"].(string), "datetime") || strings.Contains(column["column_type"].(string), "date") || strings.Contains(column["column_type"].(string), "time") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input class="input" name="%s" type="datetime-local" %s />
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired)
	}

	return formFields
}

func generateEditFields(tableName string, column map[string]interface{}, columnName string, defaultValue string, is_nullable string) string {

	isItRequired := "required"
	nameINeed := strings.Replace(toCamelCase(columnName), "Id", "ID", -1)

	// fmt.Println("default value: ",defaultValue, "----- is nullable: ", is_nullable, " is condition true", (defaultValue == "NULL" || is_nullable == "YES"))
	formFields := ""
	if defaultValue == "NULL" || is_nullable == "YES" {
		isItRequired = ""
	}
	if strings.Contains(column["column_type"].(string), "tinyint") {

		formFields = fmt.Sprintf(`<div class="form-control">
	  <label class="label cursor-pointer">
		<span class="label-text">%s</span> 
		<input type="checkbox" name="%s" class="checkbox" value={strconv.FormatBool(%s.%s)} %s />
	  </label>
	</div>
		`, capitalizeWords(splitWords(columnName)), columnName, tableName, nameINeed, isItRequired)
	} else if strings.Contains(column["column_type"].(string), "varchar") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input
			class="input input-bordered input-primary"
			type="text"
			name="%s"
			%s
			autofocus
			minlength="3"
			maxlength="64"
			value={%s.%s}
		/>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired, tableName, nameINeed)
	} else if strings.Contains(column["column_type"].(string), "text") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<textarea
			class="textarea textarea-primary h-36 max-h-36"
			%s
			name="%s"
			maxlength="255"
		>{%s.%s}</textarea>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired, tableName, nameINeed)
	} else if strings.Contains(column["column_type"].(string), "int") || strings.Contains(column["column_type"].(string), "float32") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input
			class="input input-bordered input-primary"
			type="number"
			name="%s"
			%s
			autofocus
			minlength="3"
			maxlength="64"
			value={strconv.Itoa(int(%s.%s))}
		/>
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired, tableName, nameINeed)
	} else if strings.Contains(column["column_type"].(string), "enum") {
		// fmt.Printf("%s", )
		re := regexp.MustCompile(`'([^']*)'`)
		matches := re.FindAllString(column["column_type"].(string), -1)
		formFields += fmt.Sprintf(`<label class="form-control w-full max-w-xs">
		<div class="label">
		  <span class="label-text">%s</span>

		</div>
		<select class="select select-bordered" name="%s" %s value={%s.%s} >
		  <option disabled selected>Select one</option>

`, capitalizeWords(splitWords(columnName)), columnName, isItRequired, tableName, nameINeed)
		for _, match := range matches {
			// fmt.Println(match[1 : len(match)-1])
			formFields += fmt.Sprintf(`<option>%s</option>
			`, capitalizeWords(splitWords(match[1:len(match)-1])))
		}
		formFields += `</select>

		</label>`

	} else if strings.Contains(column["column_type"].(string), "datetime") || strings.Contains(column["column_type"].(string), "date") || strings.Contains(column["column_type"].(string), "time") {
		formFields = fmt.Sprintf(`<label class="flex flex-col justify-start gap-2">
		%s:
		<input class="input" name="%s" type="datetime-local" %s value={%s.%s.String()} />
	</label>`, capitalizeWords(splitWords(columnName)), columnName, isItRequired, tableName, nameINeed)
	}

	return formFields
}

//function to convert from snake case to camel case
//skip the first letter

func toCamelCase(s string) string {
	words := strings.Split(s, "_")

	if len(words) == 1 {
		return cases.Title(language.English).String(s)
	}

	if len(words) > 1 {
		for i := 0; i < len(words); i++ {
			words[i] = cases.Title(language.English).String(words[i])
		}
	}
	return strings.Join(words, "")
}

func splitWords(s string) string {
	words := strings.Split(s, "_")

	return strings.Join(words, " ")
}

func capitalizeWords(s string) string {
	words := strings.Split(s, " ")

	var final []string
	for _, word := range words {
		final = append(final, cases.Title(language.English).String(word))
		// final += cases.Title(language.English).String(word)
	}

	return strings.Join(final, " ")
}

func generateViewFolders(parent string) {
	err := os.Mkdir("./views/"+parent, 0755)
	if err != nil {
		fmt.Println(err)
	}

	// generateIndexFiles(parent)
	e, err := os.Create("./views/" + parent + "/create.templ")
	if err != nil {
		fmt.Println(err)
	}

	e.WriteString("package " + parent)
	defer e.Close()

	d, err := os.Create("./views/" + parent + "/edit.templ")
	if err != nil {
		fmt.Println(err)
	}

	d.WriteString("package " + parent)
	defer e.Close()
}

func generateIndexFiles(tableName string, tableHeader string, firstField string, tableFields string) {
	indexTempl, err := os.ReadFile("./templates/index.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	indexTemplate := string(indexTempl)

	capitalTableName := toCamelCase(tableName)

	indexTemplate = strings.ReplaceAll(indexTemplate, "^CtableName^", capitalTableName)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableName^", tableName)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableNameSingular^", tableName)

	pluranForm := pluralize.NewClient().Plural(tableName)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableNamePlural^", pluranForm)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableFieldsHeader^", tableHeader)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableFields^", tableFields)
	indexTemplate = strings.ReplaceAll(indexTemplate, "^tableNameFirst^", firstField)
	f, err := os.Create("./views/" + tableName + "/index.templ")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(indexTemplate)
}

func generateCreateFiles(tableName string, formFields string) {
	indexTempl, err := os.ReadFile("./templates/create.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	createFile := string(indexTempl)

	capitalTableName := toCamelCase(tableName)

	createFile = strings.ReplaceAll(createFile, "^CtableName^", capitalTableName)
	createFile = strings.ReplaceAll(createFile, "^tableName^", tableName)
	createFile = strings.ReplaceAll(createFile, "^tableNameSingular^", tableName)

	pluranForm := pluralize.NewClient().Plural(tableName)
	createFile = strings.ReplaceAll(createFile, "^tableNamePlural^", pluranForm)
	// createFile = strings.ReplaceAll(createFile,"^tableFieldsHeader^", tableHeader )
	// createFile = strings.ReplaceAll(createFile,"^tableFields^", tableFields )
	createFile = strings.ReplaceAll(createFile, "^formFields^", formFields)
	createFile = strings.ReplaceAll(createFile, "^tableNameSpaces^", capitalizeWords(splitWords(tableName)))

	f, err := os.Create("./views/" + tableName + "/create.templ")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(createFile)
}

func generateEditFiles(tableName string, formFields string) {
	indexTempl, err := os.ReadFile("./templates/edit.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	editFile := string(indexTempl)

	capitalTableName := toCamelCase(tableName)

	editFile = strings.ReplaceAll(editFile, "^CtableName^", capitalTableName)
	editFile = strings.ReplaceAll(editFile, "^tableName^", tableName)
	editFile = strings.ReplaceAll(editFile, "^tableNameSingular^", tableName)

	pluranForm := pluralize.NewClient().Plural(tableName)
	editFile = strings.ReplaceAll(editFile, "^tableNamePlural^", pluranForm)
	// editFile = strings.ReplaceAll(editFile,"^tableFieldsHeader^", tableHeader )
	// editFile = strings.ReplaceAll(editFile,"^tableFields^", tableFields )
	editFile = strings.ReplaceAll(editFile, "^formFields^", formFields)
	editFile = strings.ReplaceAll(editFile, "^tableNameSpaces^", capitalizeWords(splitWords(tableName)))

	f, err := os.Create("./views/" + tableName + "/edit.templ")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	f.WriteString(editFile)
}

func generateModelTemplate(tableName string, needsTime bool, columnsFileds string) {
	f, err := os.Create("models/" + tableName + ".go")
	if err != nil {
		fmt.Println(err)
		return
	}
	modelTmpl, err := os.ReadFile("./templates/modelTemplate.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	modelTemplate := string(modelTmpl)

	if needsTime {
		modelTemplate = strings.Replace(modelTemplate, "^needsTime^", "\"time\"", 1)
	} else {
		modelTemplate = strings.Replace(modelTemplate, "^needsTime^", "", 1)
	}

	// uper := cases.Title(language.English).String(toCamelCase())

	// uper := caser.String(tableName)
	modelTemplate = strings.ReplaceAll(modelTemplate, "^CtableName^", toCamelCase(tableName))
	modelTemplate = strings.ReplaceAll(modelTemplate, "^tableName^", tableName)
	modelTemplate = strings.Replace(modelTemplate, "^columns^", columnsFileds, 1)
	modelTemplate = strings.ReplaceAll(modelTemplate, "^tableNameSingular^", pluralize.NewClient().Singular(tableName))

	_, err = f.WriteString(modelTemplate)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}

func generateHandlersTemplate(tableName string) {
	f, err := os.Create("handlers/" + tableName + ".go")
	if err != nil {
		fmt.Println(err)
		return
	}
	modelTmpl, err := os.ReadFile("./templates/handlerTemplate.tmpl")

	if err != nil {
		fmt.Println(err)
	}

	modelTemplate := string(modelTmpl)

	uper := toCamelCase(tableName)

	// uper := caser.String(tableName)
	modelTemplate = strings.ReplaceAll(modelTemplate, "^CtableName^", uper)
	modelTemplate = strings.ReplaceAll(modelTemplate, "^tableName^", tableName)
	modelTemplate = strings.ReplaceAll(modelTemplate, "^tableNameSingular^", pluralize.NewClient().Singular(tableName))

	_, err = f.WriteString(modelTemplate)
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
}
