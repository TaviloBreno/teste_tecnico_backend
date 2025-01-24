package http

import (
	"crud-clientes/internal/domain/entities"
	"crud-clientes/internal/usecase/client"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ClientHandler define os handlers para as rotas de cliente
type ClientHandler struct {
	CreateUC      *client.CreateClientUseCase
	GetUC         *client.GetClientUseCase
	GetAllClients *client.GetAllClientsUseCase
	UpdateUC      *client.UpdateClientUseCase
	DeleteUC      *client.DeleteClientUseCase
}

// CreateClient lida com a criação de um cliente
func (h *ClientHandler) CreateClient(w http.ResponseWriter, r *http.Request) {
	var client entities.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	id, err := h.CreateUC.Execute(client)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{"id": id})
}

// GetClient lida com a obtenção de um cliente pelo ID
func (h *ClientHandler) GetClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	client, err := h.GetUC.Execute(id)
	if err != nil {
		http.Error(w, "Client not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(client)
}

// GetAllClientsHandler lida com a obtenção de todos os clientes
func (h *ClientHandler) GetAllClientsHandler(w http.ResponseWriter, r *http.Request) {
	clients, err := h.GetAllClients.Execute()
	if err != nil {
		log.Printf("GetAllClients Error: %v", err) // Log do erro
		http.Error(w, "Failed to fetch clients", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(clients)
}

func (h *ClientHandler) UpdateClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	var client entities.Client
	if err := json.NewDecoder(r.Body).Decode(&client); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	client.ID = id // Certifique-se de que o ID está atribuído ao cliente

	if err := h.UpdateUC.Execute(client); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client updated successfully"})
}

func (h *ClientHandler) DeleteClient(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid client ID", http.StatusBadRequest)
		return
	}

	if err := h.DeleteUC.Execute(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Client deleted successfully"})
}
