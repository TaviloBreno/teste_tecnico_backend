package client

import (
	"crud-clientes/internal/domain/repositories"
	"fmt"
)

type DeleteClientUseCase struct {
	Repo *repositories.ClientRepository
}

func (uc *DeleteClientUseCase) Execute(id int64) error {
	if id <= 0 {
		return fmt.Errorf("invalid client ID")
	}

	return uc.Repo.Delete(id)
}
