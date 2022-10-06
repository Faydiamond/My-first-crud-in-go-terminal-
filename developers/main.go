package main

import (
	"bufio"
	"fmt"
	"godevps/db"
	"godevps/db/models"
	"os"
	"strconv"

	"github.com/TwiN/go-color"
)

func main() {

	creditos := `==========================================================
	                CRUD de MySQL y GO
                                           
                                __ __          __         
	.-----.---.-.----.-----|__|  |--.--.--|  |_.-----.
	|  _  |  _  |   _|-- __|  |  _  |  |  |   _|  -__|
	|   __|___._|__| |_____|__|_____|___  |____|_____|
	|__|                            |_____|           
    ==========================================================`

	fmt.Println(color.Ize(color.Yellow, creditos))

	menu := "¿Qué deseas hacer?\n[1] -- Ciudades\n[2] -- Lenguajes\n[3] -- Usuarios\n[5] -- Salir\n------------------------------>"

	submenu := "¿Qué deseas hacer?\n[1] -- Crear\n[2] -- Editar\n[3] -- Listar\n[4] -- Eliminar\n------------------------------>\n"

	var option uint8
	var suboption uint8
	fmt.Println(color.Ize(color.Cyan, menu))
	fmt.Scanln(&option)

	switch option {
	case 1:
		fmt.Println("Haz ingresado al menu de ciudades ")
		fmt.Println(color.Ize(color.Cyan, submenu))
		fmt.Scanln(&suboption)
		if suboption == 1 {
			fmt.Println(`Vamos a crear una Ciudad ,
            a continuacion observe el listado de paises e identifique su respectivo valor numerico`)

			countries := models.ListConts()

			for _, e := range countries {
				fmt.Println(e.Id_pais, "-", e.Pais)
			}
			fmt.Println(`Digite el numero del pais identificado, 
            al cual relacionara con la respectiva ciudad: `)
			var numberCountry uint
			fmt.Scanln(&numberCountry)
			fmt.Println(`Digite el nombre de la ciudad: `)
			var nameCity string
			fmt.Scanln(&nameCity)

			if models.ExistCity(nameCity) == true {
				fmt.Println("La ciudad ya existe y no se puede insertar de nuevo.")
			} else {
				insCity := models.CreateCity(int(numberCountry), nameCity)
				//fmt.Println()
				fmt.Println(color.Ize(color.Green, insCity))
			}
		} else if suboption == 2 {
			fmt.Println(color.Ize(color.Gray, "Vamos a editar una ciudad"))
			db.Connect()

			fmt.Println("A continuacion digita la ciudad que deseas editar ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			cityfound := scaner.Text()

			founCity := models.GetCityName(cityfound)
			if founCity.Ciudad != "" {
				//fmt.Println("LA ciudad existe")
				fmt.Println("Por favor digite el nuevo nombre de la ciudad ")
				scaner := bufio.NewScanner(os.Stdin)
				scaner.Scan()
				newcity := scaner.Text()
				founCity.Ciudad = newcity
				founCity.Update()

			} else {
				fmt.Println(color.Ize(color.Red, "La ciudad no existe, por favor digite otra ciudad que exista en la base de datos"))
			}

		} else if suboption == 3 {
			fmt.Println(color.Ize(color.Gray, "Vamos a ver las ciudades que se encuentran en la base de datos"))
			db.Connect()
			cities := models.ListCities()

			for _, citi := range cities {
				cityBycountry := fmt.Sprintf("%s - %s ", citi.Pais, citi.Ciudad)
				fmt.Println(color.Ize(color.Bold, cityBycountry))
			}
		} else if suboption == 4 {
			fmt.Println(color.Ize(color.Gray, "Vamos a Eliminar una ciudad"))

			db.Connect()

			fmt.Println("A continuacion digita la ciudad que deseas eliminar ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			cityfound := scaner.Text()

			founCity := models.GetCityName(cityfound)
			if founCity.Ciudad != "" {
				founCity.Delete()
			} else {
				fmt.Println(color.Ize(color.Red, "La ciudad no existe, por favor digite otra ciudad que exista en la base de datos"))
			}
		}
	//Languages
	case 2:
		fmt.Println("Haz ingresado al menu de lenguajes ")
		fmt.Println(color.Ize(color.Cyan, submenu))
		fmt.Scanln(&suboption)
		if suboption == 1 {
			fmt.Println(`Vamos a crear una Lenguaje`)

			fmt.Println("A continuacion digita el lenguaje de programacion que deseas crear")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			languagee := scaner.Text()
			if languagee != "" {
				db.Connect()
				existLenguage := models.GetLanguageName(languagee)
				if existLenguage.Language != "" {
					fmt.Println(color.Ize(color.Red, "Ya existe el lenguaje, por favor digite otra ciudad que exista en la base de datos"))

				} else {
					newLanguage := models.CreateLanguages(languagee)
					fmt.Println(color.Ize(color.Green, newLanguage))
				}
			}
		} else if suboption == 2 {
			fmt.Println(color.Ize(color.Gray, "Vamos a editar un lenguaje"))
			db.Connect()

			fmt.Println("A continuacion digita el lenguaje que deseas editar ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			languagefound := scaner.Text()

			foundLang := models.GetLanguageName(languagefound)

			if foundLang.Language != "" {
				//fmt.Println("LA ciudad existe")
				fmt.Println("Por favor digite el nuevo nombre del lenguaje ")
				scaner := bufio.NewScanner(os.Stdin)
				scaner.Scan()
				newlang := scaner.Text()
				foundLang.Language = newlang
				foundLang.Update()

			} else {
				fmt.Println(color.Ize(color.Red, "El lenguaje no existe, por favor digite otra lenguaje que exista en la base de datos"))
			}

		} else if suboption == 3 {
			fmt.Println(color.Ize(color.Gray, "Vamos a ver los lenguajes que se encuentran en la base de datos"))
			db.Connect()
			alllanguages := models.ListLanguages()
			for _, language := range alllanguages {
				fmt.Println(color.Ize(color.Bold, language.Language))
			}
		} else if suboption == 4 {
			fmt.Println(color.Ize(color.Gray, "Vamos a eliminar un lenguaje"))
			db.Connect()

			fmt.Println("A continuacion digita el lenguaje que deseas eliminar ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			languagefound := scaner.Text()

			foundLang := models.GetLanguageName(languagefound)

			if foundLang.Language != "" {
				//fmt.Println("LA ciudad existe")
				foundLang.Delete()
			} else {
				fmt.Println(color.Ize(color.Red, "El lenguaje no existe, por favor digite otra lenguaje que exista en la base de datos"))
			}
		}
	case 3:
		fmt.Println("Haz ingresado al menu de usuarios ")
		db.Connect()

		fmt.Println(color.Ize(color.Cyan, submenu))
		fmt.Scanln(&suboption)
		if suboption == 1 {
			fmt.Println(`Vamos a crear una Usuario`)
			//Nombre usuarui
			fmt.Println("A continuacion digite los nombres del usuario ")
			scanername := bufio.NewScanner(os.Stdin)
			scanername.Scan()
			namee := scanername.Text()
			//apellidos usuario
			fmt.Println("A continuacion digite los apellidos del usuario  ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			lastaname := scaner.Text()

			if namee != "" {
				{
					existUser := models.GetUserName(namee, lastaname)
					if existUser.Nombres != "" {
						fmt.Println(color.Ize(color.Red, "El usuario ya existe  en la base de datos"))
					} else {
						//age, number city number gender.
						fmt.Println("A continuacion digita la edad del usuario ")
						scaneragee := bufio.NewScanner(os.Stdin)
						scaneragee.Scan()
						agee := scaneragee.Text()
						fmt.Println("Agee ", agee)
						ageeint, err := strconv.Atoi(agee)
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println("A continuacion digita el numero de la ciudad del usuarioss ")
						scanercity := bufio.NewScanner(os.Stdin)
						scanercity.Scan()
						cityy := scanercity.Text()
						cityint, err := strconv.Atoi(cityy)
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println("A continuacion digita el numero del genero del usuario, recuerda que el 1 representa a las mujeres y el 2 para hombresss ")

						scanergender := bufio.NewScanner(os.Stdin)
						scanergender.Scan()
						gender := scanergender.Text()
						genderint, err := strconv.Atoi(gender)
						if err != nil {
							fmt.Println(err)
						}

						newuserr := models.CreateUser(namee, lastaname, ageeint, cityint, genderint)
						fmt.Println(newuserr)

					}
				}
			}

		} else if suboption == 2 {
			fmt.Println(`Vamos a editar un Usuario`)

			fmt.Println("A continuacion digite los nombres del usuario ")
			scanername := bufio.NewScanner(os.Stdin)
			scanername.Scan()
			namee := scanername.Text()
			//apellidos usuario
			fmt.Println("A continuacion digite los apellidos del usuario  ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			lastaname := scaner.Text()

			if namee != "" {
				{
					existUser := models.GetUserName(namee, lastaname)
					if existUser.Nombres != "" {

						fmt.Println("A continuacion digita el nuevo nombre")
						editname := bufio.NewScanner(os.Stdin)
						editname.Scan()
						nametxt := editname.Text()

						if nametxt != "" {
							existUser.Nombres = nametxt

						}

						fmt.Println("A continuacion digita los nuevos apellidos")
						editlastname := bufio.NewScanner(os.Stdin)
						editlastname.Scan()
						lastnametxt := editlastname.Text()

						if lastnametxt != "" {
							existUser.Apellidos = lastnametxt
						}

						fmt.Println("A continuacion digita la edad del usuario")
						editage := bufio.NewScanner(os.Stdin)
						editage.Scan()
						ageeint := editage.Text()
						ageinteger, err := strconv.Atoi(ageeint)
						if err != nil {
							fmt.Println(err)
						}

						if ageinteger > 0 {
							existUser.Age = ageinteger
						}
						existUser.Update()

					} else {
						fmt.Println(color.Ize(color.Red, "El usuario no  existe  en la base de datos"))
					}
				}
			}
		} else if suboption == 3 {
			fmt.Println("Listar usuarios")
			users := models.ListUsers()
			fmt.Println(users)

			for _, user := range users {
				usercomplete := fmt.Sprintf("%s  %s  %d", user.Nombres, user.Apellidos, user.Age)
				fmt.Println(color.Ize(color.Bold, usercomplete))
			}
		} else if suboption == 4 {
			fmt.Println("Vamos a eliminar un usuario")
			fmt.Println("A continuacion digite los nombres del usuario ")
			scanername := bufio.NewScanner(os.Stdin)
			scanername.Scan()
			namee := scanername.Text()
			//apellidos usuario
			fmt.Println("A continuacion digite los apellidos del usuario  ")
			scaner := bufio.NewScanner(os.Stdin)
			scaner.Scan()
			lastaname := scaner.Text()

			if namee != "" {
				existUser := models.GetUserName(namee, lastaname)
				if existUser.Nombres != "" {
					existUser.DeleteUser()
				} else {
					fmt.Println(color.Ize(color.Red, "El usuario no  existe  en la base de datos"))
				}
			}
		}
	} //closed case
}
