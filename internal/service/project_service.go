package service

import (
	"context"
	"github.com/google/uuid"
	"spt/internal/gorm_gen/model"
	dto "spt/internal/model"
	"spt/internal/repository"
)

type ProjectService interface {
	CreateProject(ctx context.Context, project *dto.ProjectCreate) error
	GetProjectList(ctx context.Context) ([]*model.Project, error)
	GetProjectByID(ctx context.Context, projectID string) (*model.Project, error)
}

type projectService struct {
	repo repository.ProjectRepository
}

func NewProjectService(repo repository.ProjectRepository) ProjectService {
	return &projectService{
		repo: repo,
	}
}
func (p *projectService) CreateProject(ctx context.Context, project *dto.ProjectCreate) error {

	newProject := &model.Project{
		ID:                 uuid.New().String(),
		ProjectName:        project.ProjectName,
		Pinned:             project.Pinned,
		ProjectDescription: project.ProjectDescription,
		Tags:               project.Tags,
		LogoPicture:        project.LogoPicture,
		GithubURL:          project.GithubURL,
		SiteURL:            project.SiteURL,
		OwnerEmail:         project.OwnerEmail,
	}

	return p.repo.CreateProject(ctx, newProject)
}

func (p *projectService) GetProjectList(ctx context.Context) ([]*model.Project, error) {
	return p.repo.GetProjectList(ctx)
}

func (p *projectService) GetProjectByID(ctx context.Context, projectID string) (*model.Project, error) {
	return p.repo.GetProjectByID(ctx, projectID)
}
