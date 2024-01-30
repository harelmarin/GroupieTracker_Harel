package template

import (
	"fmt"
	"html/template"
)

var Temp *template.Template

func InitTemplate() {
	temp, errTemp := template.ParseGlob("./templates/*.html")
	if errTemp != nil {
		fmt.Println(fmt.Sprint("Erreur => %s", errTemp.Error()))
	}
	Temp = temp
}
