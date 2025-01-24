package client

import (
	"crud-clientes/internal/domain/entities"
	"crud-clientes/internal/domain/repositories"
	"fmt"
)

type GetClientUseCase struct {
	Repo *repositories.ClientRepository
}

func (uc *GetClientUseCase) Execute(id int64) (*entities.Client, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid client ID")
	}

	client, err := uc.Repo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get client: %w", err)
	}

	return client, nil
}
