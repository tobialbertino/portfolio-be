package usecase

import (
	"context"
	"errors"
	"time"
	"tobialbertino/portfolio-be/exception"
	"tobialbertino/portfolio-be/internal/notes/models/domain"
	"tobialbertino/portfolio-be/internal/notes/models/entity"
	"tobialbertino/portfolio-be/internal/notes/repository/postgres"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type NotesUseCaseImpl struct {
	NotesRepository postgres.NotesRepository
	DB              *pgxpool.Pool
	Validate        *validator.Validate
}

func NewNotesUseCase(NotesRepo postgres.NotesRepository, DB *pgxpool.Pool, validate *validator.Validate) NotesUseCase {
	return &NotesUseCaseImpl{
		NotesRepository: NotesRepo,
		DB:              DB,
		Validate:        validate,
	}
}

// Add implements NotesUseCase
func (useCase *NotesUseCaseImpl) Add(req *domain.ReqAddNote) (*domain.NoteId, error) {
	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, exception.NewClientError(err.Error(), 400)
	}

	request := &entity.Notes{
		Id:        uuid.New().String(),
		Title:     req.Title,
		Body:      req.Body,
		Tags:      req.Tags,
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
	}
	i, err := useCase.NotesRepository.Add(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	if i <= 0 {
		err = exception.Wrap("not found", 404, errors.New("rows affected: 0"))
		return nil, err
	}

	response := &domain.NoteId{
		NoteId: request.Id,
	}
	return response, err
}

func (useCase *NotesUseCaseImpl) GetAll() (*[]domain.Notes, error) {
	var listResult *entity.ListNotes = new(entity.ListNotes)

	listResult, err := useCase.NotesRepository.GetAll(context.Background(), useCase.DB)
	if err != nil {
		return nil, err
	}

	result := listResult.ToDomain()
	return result, nil
}

func (useCase *NotesUseCaseImpl) GetById(id string) (*domain.Notes, error) {
	var listResult *entity.Notes = new(entity.Notes)

	listResult, err := useCase.NotesRepository.GetById(context.Background(), useCase.DB, id)
	if err != nil {
		return nil, err
	}
	if listResult.Id == "" {
		return nil, exception.Wrap("Catatan tidak ditemukan", 404, errors.New("fail"))
	}

	result := listResult.ToDomain()
	return result, nil
}

func (useCase *NotesUseCaseImpl) Update(req *domain.ReqAddNote, id string) (*domain.Notes, error) {
	err := useCase.Validate.Struct(req)
	if err != nil {
		return nil, exception.NewClientError(err.Error(), 400)
	}

	request := &entity.Notes{
		Id:        id,
		Title:     req.Title,
		Body:      req.Body,
		Tags:      req.Tags,
		UpdatedAt: time.Now().Unix(),
	}
	_, err = useCase.NotesRepository.Update(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response := request.ToDomain()
	return response, err
}

func (useCase *NotesUseCaseImpl) Delete(id string) (*domain.RowsAffected, error) {
	request := &entity.Notes{
		Id: id,
	}

	i, err := useCase.NotesRepository.Delete(context.Background(), useCase.DB, request)
	if err != nil {
		return nil, err
	}

	response := &domain.RowsAffected{
		RowsAffected: i,
	}
	return response, err
}
