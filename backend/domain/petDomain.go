package domain

import "github.com/forceattack012/petAppApi/entities"

type PetStore interface {
	ExceuteSql(sql string, v interface{}) error
	Raw(sql string, param interface{}, dest interface{}) error
	Paginate(page int, pageSize int, preload string, sqlRaw string, where string, pet *[]entities.Pet) error
	Save(pet *entities.Pet) error
	GetWithFiles(pets *[]entities.Pet) error
	GetPet(pet *entities.Pet, id int) error
	Delete(pet *entities.Pet, id int) error
	Update(pet *entities.Pet) error
}
