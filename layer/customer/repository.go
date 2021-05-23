package customer

import (
	"project-individu-go-react/entities"

	"gorm.io/gorm"
)

type Repository interface {
	FindAll() ([]entities.Costumer, error)
	Create(customer entities.Costumer) (entities.Costumer, error)
	UpdateByID(CID string, dataUpdate map[string]interface{}) (entities.Costumer, error)
	DeleteByID(CID string) (string, error)
	FindByID(CID string) (entities.Costumer, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

//CUSTOMER

//READ CUSTOMERS
func (r *repository) FindAll() ([]entities.Costumer, error) {
	var customers []entities.Costumer

	if err := r.db.Find(&customers).Error; err != nil {
		return customers, err
	}

	return customers, nil
}

// Read by id
func (r *repository) FindByID(CID string) (entities.Costumer, error) {
	var costumers entities.Costumer

	if err := r.db.Where("c_id = ?", CID).Find(&costumers).Error; err != nil {
		return costumers, err
	}

	return costumers, nil
}

//CREATE
func (r *repository) Create(customer entities.Costumer) (entities.Costumer, error) {
	if err := r.db.Create(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

//UPDATE
func (r *repository) UpdateByID(CID string, dataUpdate map[string]interface{}) (entities.Costumer, error) {
	var customer entities.Costumer

	if err := r.db.Model(&customer).Where("c_id = ?", CID).Updates(dataUpdate).Error; err != nil {
		return customer, err
	}

	if err := r.db.Where("c_id = ?", CID).Find(&customer).Error; err != nil {
		return customer, err
	}

	return customer, nil
}

//DELETE
func (r *repository) DeleteByID(CID string) (string, error) {
	if err := r.db.Where("c_id = ?", CID).Delete(&entities.Costumer{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}
