package service

import (
	"movie-app/cmd/metadata/internal/models"
	"movie-app/cmd/metadata/internal/repository"
)

type MetadataService struct {
	repo repository.MetadataRepository
}

func NewMetadataService(repo repository.MetadataRepository) MetadataService {
	return MetadataService{repo: repo}
}

func (s MetadataService) GetMetadata(id string) (models.Metadata, error) {
	return s.repo.GetMetadata(id)
}

func (s MetadataService) CreateMetadata(metadata models.Metadata) (models.Metadata, error) {
	return s.repo.CreateMetadata(metadata)
}

func (s MetadataService) UpdateMetadata(metadata models.Metadata) (models.Metadata, error) {
	return s.repo.UpdateMetadata(metadata)
}

func (s MetadataService) DeleteMetadata(id string) error {
	return s.repo.DeleteMetadata(id)
}
