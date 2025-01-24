package repositories

import (
	"crud-clientes/internal/domain/entities"
	"database/sql"
	"fmt"
	"log"
)

type ClientRepository struct {
	DB *sql.DB
}

func NewClientRepository(db *sql.DB) *ClientRepository {
	return &ClientRepository{DB: db}
}

func (r *ClientRepository) Create(client entities.Client) (int64, error) {
	query := "INSERT INTO clients (name, email, phone) VALUES (?, ?, ?)"
	result, err := r.DB.Exec(query, client.Name, client.Email, client.Phone)
	if err != nil {
		return 0, fmt.Errorf("failed to create client: %w", err)
	}
	return result.LastInsertId()
}

func (r *ClientRepository) GetAll() ([]entities.Client, error) {
	query := "SELECT id, name, email, phone, created_at, updated_at FROM clients"
	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("SQL Query Error: %v", err) // Log do erro da consulta SQL
		return nil, fmt.Errorf("failed to fetch clients: %w", err)
	}
	defer rows.Close()

	var clients []entities.Client
	for rows.Next() {
		var client entities.Client
		if err := rows.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.CreatedAt, &client.UpdatedAt); err != nil {
			log.Printf("Row Scan Error: %v", err) // Log do erro ao escanear a linha
			return nil, fmt.Errorf("failed to scan client: %w", err)
		}
		clients = append(clients, client)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows Iteration Error: %v", err) // Log do erro na iteração das linhas
		return nil, fmt.Errorf("error iterating over rows: %w", err)
	}

	return clients, nil
}

func (r *ClientRepository) GetByID(id int64) (*entities.Client, error) {
	query := "SELECT id, name, email, phone, created_at, updated_at FROM clients WHERE id = ?"
	row := r.DB.QueryRow(query, id)

	var client entities.Client
	if err := row.Scan(&client.ID, &client.Name, &client.Email, &client.Phone, &client.CreatedAt, &client.UpdatedAt); err != nil {
		return nil, fmt.Errorf("client not found: %w", err)
	}
	return &client, nil
}

func (r *ClientRepository) Update(client entities.Client) error {
	query := "UPDATE clients SET name = ?, email = ?, phone = ? WHERE id = ?"
	_, err := r.DB.Exec(query, client.Name, client.Email, client.Phone, client.ID)
	if err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}
	return nil
}

func (r *ClientRepository) Delete(id int64) error {
	query := "DELETE FROM clients WHERE id = ?"
	_, err := r.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete client: %w", err)
	}
	return nil
}
