package repository

import "gorm.io/gorm"

type Evaluation struct {
	ID      uint64 `json:"ID" gorm:"primaryKey"`
	Comment string `json:"comment"`
	Value   int16  `json:"value" gorm:"check:value>0 and value<6"`
	MovieFk int64  `json:"movieFk"`
}

type EvaluationRepository struct {
	dbConnection *gorm.DB
}

func NewEvaluationRepository(db *gorm.DB) *EvaluationRepository {
	return &EvaluationRepository{
		dbConnection: db,
	}
}

func (evaluationRepo *EvaluationRepository) GetAll(idMovie uint64) ([]Evaluation, error) {
	var res []Evaluation
	tx := evaluationRepo.dbConnection.Where("movie_fk = ?", idMovie).Find(&res)
	if tx.Error != nil {
		return res, tx.Error
	}
	return res, nil
}

func (evaluationRepo *EvaluationRepository) CreateEvaluation(newEvaluation Evaluation) (Evaluation, error) {
	var evalError Evaluation
	tx := evaluationRepo.dbConnection.Create(&newEvaluation)
	if tx.Error != nil {
		return evalError, tx.Error
	}
	return newEvaluation, nil
}
