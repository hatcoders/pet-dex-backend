package db

import (
	"database/sql"
	"fmt"
	"pet-dex-backend/v2/entity"
	"pet-dex-backend/v2/interfaces"
	"strings"
)

type BreedRepository struct {
	dbconnection *sql.DB
}

func NewBreedRepository(db *sql.DB) interfaces.BreedRepository {
	return &BreedRepository{
		dbconnection: db,
	}
}

func (br *BreedRepository) Save(entity.Breed) error {
	return nil
}

func (br *BreedRepository) FindById(id int) (breed *entity.Breed, err error) {

	return nil, nil
}

func (br *BreedRepository) List() (breeds []*entity.Breed, err error) {
	rows, err := br.dbconnection.Query(`
		SELECT 
			id, 						
			name, 					
			specie, 				
			size,					
			description, 	
			height,		
			weight,		
			physicalChar,	
			disposition,	
			idealFor,		
			fur,			
			imgUrl,		
			weather,		
			dressage,		
			orgId,	 		
			lifeExpectancy
		FROM breeds`)
	if err != nil {
		return nil, fmt.Errorf("error listing breeds: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var breed entity.Breed
		err := rows.Scan(
			&breed.ID,
			&breed.Name,
			&breed.Specie,
			&breed.Size,
			&breed.Description,
			&breed.Height,
			&breed.Weight,
			&breed.PhysicalChar,
			&breed.Disposition,
			&breed.IdealFor,
			&breed.Fur,
			&breed.ImgUrl,
			&breed.Weather,
			&breed.Dressage,
			&breed.OrgID,
			&breed.LifeExpectancy,
		)
		if err != nil {
			return nil, fmt.Errorf("error scanning breeds: %w", err)
		}
		breeds = append(breeds, &breed)
	}

	return breeds, nil
}

func (br *BreedRepository) Update(breedID string, updatePayload map[string]interface{}) error {
	query := "UPDATE breeds SET "
	values := []interface{}{}

	for key, value := range updatePayload {
		query += key + "=?, "
		values = append(values, value)
	}

	query = strings.TrimSuffix(query, ", ")
	query += " WHERE id=?"
	values = append(values, breedID)

	_, err := br.dbconnection.Exec(query, values...)
	if err != nil {
		return fmt.Errorf("error updating breed: %w \\n", err)
	}

	return nil
}
