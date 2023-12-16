package main

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var (
		usernameSession string
		id = 0
	)
	
	clear()
	db := inicializeDB()

	for true{

		for id == 0 {

			x := 0
	
			fmt.Println("[1] - Login")
			fmt.Println("[2] - Cadastro")
			fmt.Println("[3] - Encerrar")
	
			fmt.Print("Opção: ")
			fmt.Scan(&x)
	
			switch x {
			case 1:
				idReturn, usernameReturn := login(db)
				id = idReturn
				usernameSession = usernameReturn
				break
			case 2:
				registerNewUser(db)
				break
	
			case 3:
				fmt.Println("Adeus!")
				os.Exit(1)
			default:
				clear()
				fmt.Println("Por favor escolha uma opção válida.")
				x = 0
				break
			}
		}

		fmt.Printf("Bem vindo, %v\n", usernameSession)

		for true{

			x := 0

			fmt.Println("[1] - Exibir tarefas")
			fmt.Println("[2] - Adicionar tarefa")
			fmt.Println("[3] - Deslogar")
			fmt.Println("[4] - Encerrar")
			fmt.Print("Opção: ")
			fmt.Scan(&x)

			switch x {
			case 1:
				clear()
				for true{
					tasks := consultTaskTable(db, id, 0)

					fmt.Println("[1] - Abrir tarefa")
					fmt.Println("[2] - Exibir tarefas completadas")
					fmt.Println("[3] - Voltar ao menu")

					x := 0
					fmt.Scan(&x)

					switch x{
					case 1:
						clear()
						tasks := tasks
						fmt.Println("ID da tarefa: ")
						idTask := 0
						fmt.Scan(&idTask)

						clear()

						fmt.Println(idTask, "-", tasks[idTask].nome)
						fmt.Println(tasks[idTask].descricao)
						fmt.Println("--------------------------------------")

						fmt.Println("[1] - Concluir tarefa")
						fmt.Println("[2] - Deletar tarefa")
						fmt.Println("[3] - Voltar ao menu de tarefas")
						y := 0
						fmt.Scan(&y)
						switch y {
						case 1:
							completeTask(db, tasks[idTask].id)
							break
						case 2:
							deleteTask(db, tasks[idTask].id)
							break
						case 3:
							break
						}
						break
					case 2:
						clear()
						 _ = consultTaskTable(db, id, 1)
						fmt.Println("[1] - Voltar ao menu de tarefas")
						fmt.Println("[2] - Voltar ao menu")
						z := 0
						fmt.Scan(&z)
						if z == 2{x = 3}
						break
					}
					if x == 3{
						clear()
						break
					}
				}
				break
			case 2:
				addTask(db, id)
				break
			case 3:
				id = 0
				break
			case 4:
				fmt.Println("Adeus!")
				os.Exit(1)
			}
			if id == 0{
				clear()
				break
			}
		}
	}
}