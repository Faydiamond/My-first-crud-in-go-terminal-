package models

import (
	"fmt"
	"godevps/db"

	"github.com/TwiN/go-color"
)

type Language struct {
	Id_language int
	Language    string
	Date        string
}

type Languages []Language

/* Insertar */
func newlenguages(language string) *Language {
	lenguagecreated := &Language{Language: language}
	return lenguagecreated
}

func (language *Language) insertLanguage() {
	sql := "INSERT INTO lenguajes SET lenguaje=?"
	db.Exec(sql, language.Language)

}

func CreateLanguages(language string) string {
	languagee := newlenguages(language)
	languagee.insertLanguage()
	res := fmt.Sprintf(" Se creo el lenguaje'%s' con exito. ", language)
	return res
}

/* Fin insertar */

/* Editar */

/* Fin Editar */

/* Lenguaje por nombre */
func GetLanguageName(language string) *Language {
	lang := newlenguages("")
	sql := "select id_lenguaje,lenguaje,fecha_creacion from lenguajes where lenguaje=?"
	row, _ := db.Query(sql, language)
	for row.Next() {
		row.Scan(&lang.Id_language, &lang.Language, &lang.Date)
	}
	return lang
}

/* Fin Lenguaje por nombre */

/* Editar lenguaje */
func (lang *Language) Update() {
	sql := " Update lenguajes set lenguaje=? where  lenguaje=? "
	db.Exec(sql, lang.Language, lang.Language)
	res := fmt.Sprintf("El lenguaje  %s se actualizo de manera correcta!", lang.Language)
	fmt.Println(color.Ize(color.Green, res))
}

/* Fin Editar lenguaje */

/* Eliminar Lenguaje */
func (lang *Language) Delete() {
	sql := "DELETE from lenguajes where lenguaje=?"
	db.Exec(sql, lang.Language)
	res := fmt.Sprintf("El lenguaje %s se elimino de manera correcta!", lang.Language)
	fmt.Println(color.Ize(color.Green, res))
}

/* Fin Eliminar Lenguaje */

/* Listar lenguajes */
func ListLanguages() Languages {
	langus := Languages{}
	sql := "SELECT * FROM developers.lenguajes"
	rows, _ := db.Query(sql)
	for rows.Next() {
		lang := Language{}
		rows.Scan(&lang.Id_language, &lang.Language, &lang.Date)
		langus = append(langus, lang)
	}
	return langus
}

/* Fin Listar lenguajes */
