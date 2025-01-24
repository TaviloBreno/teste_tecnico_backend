package main

import (
	"crud-clientes/internal/config"
	"crud-clientes/internal/domain/repositories"
	handlerHttp "crud-clientes/internal/handler/http"
	"crud-clientes/internal/usecase/client"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Conexão com o banco de dados
	db, err := config.ConnectMySQL()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}
	defer db.Close()

	// Injeção de dependências
	clientRepo := repositories.NewClientRepository(db)
	createClientUC := &client.CreateClientUseCase{Repo: clientRepo}
	getClientUC := &client.GetClientUseCase{Repo: clientRepo}
	getAllClientsUC := &client.GetAllClientsUseCase{Repo: clientRepo}
	updateClientUC := &client.UpdateClientUseCase{Repo: clientRepo}
	deleteClientUC := &client.DeleteClientUseCase{Repo: clientRepo}

	clientHandler := &handlerHttp.ClientHandler{
		CreateUC:      createClientUC,
		GetUC:         getClientUC,
		GetAllClients: getAllClientsUC,
		UpdateUC:      updateClientUC,
		DeleteUC:      deleteClientUC,
	}

	// Configuração do roteador
	r := mux.NewRouter()
	r.HandleFunc("/clients", clientHandler.CreateClient).Methods("POST")
	r.HandleFunc("/clients/{id:[0-9]+}", clientHandler.GetClient).Methods("GET")
	r.HandleFunc("/clients", clientHandler.GetAllClientsHandler).Methods("GET")
	r.HandleFunc("/clients/{id:[0-9]+}", clientHandler.UpdateClient).Methods("PUT")
	r.HandleFunc("/clients/{id:[0-9]+}", clientHandler.DeleteClient).Methods("DELETE")

	// Inicialização do servidor HTTP
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
