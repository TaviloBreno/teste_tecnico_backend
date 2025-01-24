package client

import (
	"crud-clientes/internal/domain/entities"
	"crud-clientes/internal/domain/repositories"
	"fmt"
)

type GetAllClientsUseCase struct {
	Repo *repositories.ClientRepository
}

func (uc *GetAllClientsUseCase) Execute() ([]entities.Client, error) {
	clients, err := uc.Repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all clients: %w", err)
	}
	return clients, nil
}
