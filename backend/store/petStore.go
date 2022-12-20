package store

import (
	"github.com/forceattack012/petAppApi/entities"
	"gorm.io/gorm"
)

type PetStore struct {
	*gorm.DB
}

func NewPetStore(db *gorm.DB) *PetStore {
	return &PetStore{db}
}

func (s *PetStore) ExceuteSql(sql string, v interface{}) error {
	return s.DB.Raw(sql).Scan(v).Error
}

func (s *PetStore) Raw(sql string, param interface{}, dest interface{}) error {
	return s.DB.Raw(sql, param).Scan(dest).Error
}

func (s *PetStore) Paginate(page int, pageSize int, preload string, sqlRaw string, where string, pet *[]entities.Pet) error {
	return s.DB.Scopes(paginate(page, pageSize)).
		Preload(preload).
		Joins(sqlRaw).
		Where(where).Find(pet).Error
}

func (s *PetStore) Save(pet *entities.Pet) error {
	return s.DB.Create(pet).Error
}

func (s *PetStore) GetWithFiles(pets *[]entities.Pet) error {
	return s.DB.Preload("Files").Find(&pets).Error
}

func (s *PetStore) GetPet(pet *entities.Pet, id int) error {
	return s.DB.Preload("Files").Find(&pet, id).Error
}

func (s *PetStore) Delete(pet *entities.Pet, id int) error {
	return s.DB.Preload("Files").Delete(pet, id).Error
}

func (s *PetStore) Update(pet *entities.Pet) error {
	return s.DB.Updates(pet).Error
}

func paginate(page int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
