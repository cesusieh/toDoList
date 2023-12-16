package main

import (
	"database/sql"
	"fmt"
)

type task struct {
	id	int
	nome string
	descricao string
	concluido bool
	idUsuario int
}


func consultTaskTable(db *sql.DB, id int, opcao int) []task{

	rows, err := db.Query(`select * from TAREFA where idUsuario = ? and concluido = ?`, id, opcao)
	checkErr(err)
	defer rows.Close()

	var	tasks []task
	for rows.Next() {
		var t task
		err := rows.Scan(&t.id, &t.nome, &t.descricao, &t.concluido, &t.idUsuario)
		checkErr(err)
		tasks = append(tasks, t)
	}
	fmt.Println("----------------------------------")
	for x,y := range(tasks){
		fmt.Println(x, " - ", y.nome)
		fmt.Println("----------------------------------")
	}
	return tasks
}

func addTask(db *sql.DB, id int){
	var(
		taskName string
		taskDesc string
	)

	clear()
	fmt.Print("Nome da tarefa: ")
	taskName = scan()
	fmt.Print(("Descrição da tarefa: "))
	taskDesc = scan()

	query := `insert into TAREFA(nome, descricao, concluido, idUsuario) values (?, ?, false, ?);`

	_, err := db.Exec(query, taskName, taskDesc, id)

	checkErr(err)

	clear()
	fmt.Println("Tarefa adicionada!")

}

func completeTask(db *sql.DB, idTask int) {

	query := `update TAREFA set concluido = true where id = ?`

	_, err := db.Exec(query, idTask)
	checkErr(err)

	clear()
	fmt.Println("Tarefa completa, muito bem.")

}

func deleteTask(db *sql.DB, idTask int){

	query := `delete from TAREFA where id = ?`
	
	_, err := db.Exec(query, idTask)
	checkErr(err)

	clear()
	fmt.Println("Tarefa deletada.")

}