package usecase

import (
	"context"
	"spt/internal/gorm_gen/model"
	dto "spt/internal/model"
	"spt/internal/service"
)

type ProjectUsecase interface {
	CreateNewProject(ctx context.Context, project *dto.ProjectCreate) error
	GetProjectList(ctx context.Context) ([]*model.Project, error)
	GetProjectByID(ctx context.Context, projectID string) (*model.Project, error)
}

type projectUsecase struct {
	service service.ProjectService
}

func NewProjectUsecase(service service.ProjectService) ProjectUsecase {
	return &projectUsecase{service: service}
}

func (p *projectUsecase) CreateNewProject(ctx context.Context, project *dto.ProjectCreate) error {
	return p.service.CreateProject(ctx, project)
}

func (p *projectUsecase) GetProjectList(ctx context.Context) ([]*model.Project, error) {
	return p.service.GetProjectList(ctx)
}

func (p *projectUsecase) GetProjectByID(ctx context.Context, projectID string) (*model.Project, error) {
	return p.service.GetProjectByID(ctx, projectID)
}
