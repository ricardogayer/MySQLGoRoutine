package main

import (
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	db, err := sql.Open("mysql", "root:senha@tcp(127.0.0.1:3306)/projeto")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, name, email FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Crie um canal para enviar linhas da tabela para goroutines processadoras.
	// A capacidade do canal determina quantas goroutines processadoras podem trabalhar
	// simultaneamente.
	rowsChan := make(chan *User, 10)

	// Crie uma WaitGroup para esperar que todas as goroutines processadoras concluam.
	var wg sync.WaitGroup

	// Inicie as goroutines processadoras. Aqui estamos iniciando 5 goroutines, mas você
	// pode ajustar esse número para o que for apropriado para o seu caso.
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go processRows(rowsChan, &wg)
	}

	// Itere sobre as linhas da tabela e envie cada uma para o canal.
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			panic(err)
		}
		rowsChan <- &user
	}

	// Feche o canal para informar às goroutines processadoras que não há mais linhas para processar.
	close(rowsChan)

	// Aguarde todas as goroutines processadoras concluírem.
	wg.Wait()
}

func processRows(rowsChan <-chan *User, wg *sync.WaitGroup) {
	defer wg.Done()

	// Processa cada linha da tabela recebida do canal.
	for user := range rowsChan {
		fmt.Printf("Processando usuário com ID %d -> %s\n", user.ID, user.Name)

		// Faça o que quiser com os dados do usuário aqui.
		// Simule algum processamento demorado para fins de demonstração.
		// Você pode remover essa linha em um uso real.
		// time.Sleep(1 * time.Second)
	}
}
