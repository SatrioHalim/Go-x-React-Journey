package repositories

import (
	"strings"

	"github.com/SatrioHalim/Go-x-React-Journey/config"
	"github.com/SatrioHalim/Go-x-React-Journey/models"
)

// Interface biar bisa dibaca dari luar
type UserRepository interface {
	Create(user *models.User) error
	FindByEmail(email string) (*models.User,error)
	FindByID(id uint)(*models.User,error)
	FindByPublicID(publicID string)(*models.User,error)
	FindAllPagination(filter,sort string, limit,offset int)([]models.User,int64,error)
}

type userRepository struct {}

func NewUserRepository()UserRepository{
	return &userRepository{}
}

func (r *userRepository) Create(user *models.User) error{
	return config.DB.Create(user).Error
}

func (r *userRepository) FindByEmail(email string)(*models.User,error){
	var user models.User
	err := config.DB.Where("email = ?",email).First(&user).Error
	return &user,err
}

func (r *userRepository) FindByID(id uint)(*models.User,error){
	var user models.User
	err := config.DB.First(&user, id).Error
	return &user, err
}

func (r *userRepository) FindByPublicID(publicID string)(*models.User,error){
	var user models.User
	err := config.DB.Where("public_id = ?",publicID).First(&user).Error
	return &user, err
}

func (r *userRepository) FindAllPagination(filter,sort string, limit,offset int)([]models.User,int64,error){
	var users []models.User
	var total int64

	db := config.DB.Model(&models.User{})

	// Filtering
	if filter != ""{
		filterPattern := "%" + filter + "%"
		db = db.Where("name Ilike ? OR email Ilike ?",filterPattern,filterPattern) // Ilike => Case Insensitive, LIKE => Case sensitive
	}
	// Count total data
	if err := db.Count(&total).Error ; err != nil {
		return nil,0,err
	}

	// Sorting
	if sort != "" {
		// Ex, sort=name (ASC ascending), sort=-name(DSC descending)
		if sort == "-id" {
			sort = "-internal_id"
		} else if sort == "id"{
			sort = "internal_id"
		}

		if strings.HasPrefix(sort,"-"){
			sort = strings.TrimPrefix(sort,"-") + " DESC"
		} else {
			sort += "ASC"
		}
		db = db.Order(sort)
	}

	err := db.Limit(limit).Offset(offset).Find(&users).Error
	return users,total,err
}