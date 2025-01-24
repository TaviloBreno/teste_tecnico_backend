package client

import (
	"crud-clientes/internal/domain/entities"
	"crud-clientes/internal/domain/repositories"
	"fmt"
)

type CreateClientUseCase struct {
	Repo *repositories.ClientRepository
}

func (uc *CreateClientUseCase) Execute(client entities.Client) (int64, error) {
	if client.Name == "" || client.Email == "" || client.Phone == "" {
		return 0, fmt.Errorf("all fields are required")
	}
	return uc.Repo.Create(client)
}
