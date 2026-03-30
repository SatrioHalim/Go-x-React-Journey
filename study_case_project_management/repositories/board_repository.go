package repositories

import (
	"github.com/SatrioHalim/Go-x-React-Journey/config"
	"github.com/SatrioHalim/Go-x-React-Journey/models"
)

type BoardRepository interface {
	Create(board *models.Board) error
}

type boardRepository struct {
}

func NewBoardRepository() BoardRepository {
	return &boardRepository{}
}

func (r *boardRepository) Create(board *models.Board) error {
	return config.DB.Create(board).Error
}