package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"testedocker/db"
)

func main() {
	db := db.ConectaComBancoDeDados()

	res, error := db.Exec("create table if not exists meubanco.teste (id int auto_increment not null primary key, name varchar(255));")
	if error != nil {
		log.Fatal("deu ruim ao criar tabela ", error)
	} else {
		fmt.Println("deu bom ", res)
	}

	insereDadosNoBanco, err := db.Prepare("insert into meubanco.teste(name) values(?)")
	if err != nil {
		log.Fatal("deu ruim no prepare ", err)
	}

	insert, errorInsert := insereDadosNoBanco.Exec("testando")

	if errorInsert != nil {
		log.Fatal("deu ruim no exec ", errorInsert)
	} else {
		fmt.Println("deu bom no insert ", insert)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("{'hello' : 'world!'}")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))

}
