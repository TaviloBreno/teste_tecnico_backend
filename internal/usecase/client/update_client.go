package client

import (
	"crud-clientes/internal/domain/entities"
	"crud-clientes/internal/domain/repositories"
	"fmt"
)

type UpdateClientUseCase struct {
	Repo *repositories.ClientRepository
}

func (uc *UpdateClientUseCase) Execute(client entities.Client) error {
	if client.ID <= 0 {
		return fmt.Errorf("invalid client ID")
	}
	if client.Name == "" || client.Email == "" || client.Phone == "" {
		return fmt.Errorf("all fields are required")
	}

	err := uc.Repo.Update(client)
	if err != nil {
		return fmt.Errorf("failed to update client: %w", err)
	}

	return nil
}
