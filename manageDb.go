package main

import (
	"database/sql"
	"fmt"
)

func login(db *sql.DB) (int, string){

	var(
		username string
		password string
		usernameReturn string
		passwordReturn string
		idReturn int 
	)

	for true{
		fmt.Println("Nome de usuário: ")
		username = scan()
			
		query := `select id, username, password from USER where username = ?`
		err := db.QueryRow(query, username).Scan(&idReturn, &usernameReturn, &passwordReturn)

		if err == nil {
			fmt.Println("Senha: ")
			fmt.Scan(&password)
			
			if password == passwordReturn{
				clear()
				fmt.Println("Logado!")
				break
			} else {
				fmt.Println("Senha incorreta.")
			}
		} else {
			fmt.Println("Usuário não encontrado.")
		}

	}

	return idReturn, usernameReturn

}


func registerNewUser(db *sql.DB) {

	var(
		username string
		password string
		usernameReturn string
	)
	for true{
		fmt.Print("Nome de usuario: ")
		fmt.Scan(&username)

		query := `select username, password from USER where username = ?`
		err := db.QueryRow(query, username).Scan(&usernameReturn)

		if err.Error() == "sql: no rows in result set"{
			fmt.Print(("senha: "))
			fmt.Scan(&password)

			query := `insert into USER(username, password) values (?, ?);`

			_, err := db.Exec(query, username, password)
			checkErr(err)
			clear()
			fmt.Println("Usuario cadastrado")
			break
		} else {
			fmt.Println("Nome de usuário indisponível")
		}
	}
}

	

func deleteUser(db *sql.DB, id int) {

	query := `DELETE FROM USER WHERE id = ?`

	_, err := db.Exec(query, id)

	checkErr(err)

}
