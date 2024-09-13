package repository

import (
	"context"
	"spt/internal/db"
	"spt/internal/gorm_gen/model"
	"spt/internal/gorm_gen/models"
)

type ProjectRepository interface {
	CreateProject(ctx context.Context, project *model.Project) error
	GetProjectList(ctx context.Context) ([]*model.Project, error)
	GetProjectByID(ctx context.Context, id string) (*model.Project, error)
}

type projectRepository struct {
	dbService db.Service
}

func NewProjectRepository(dbService db.Service) ProjectRepository {
	return &projectRepository{
		dbService: dbService,
	}
}

func (p *projectRepository) CreateProject(ctx context.Context, project *model.Project) error {
	q := models.Use(p.dbService.GetDB(ctx))
	err := q.Project.Create(project)
	return err
}

func (p *projectRepository) GetProjectList(ctx context.Context) ([]*model.Project, error) {
	q := models.Use(p.dbService.GetDB(ctx))
	result, err := q.Project.ReadDB().Find()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (p *projectRepository) GetProjectByID(ctx context.Context, id string) (*model.Project, error) {
	q := models.Use(p.dbService.GetDB(ctx))
	result, err := q.Project.ReadDB().Where(q.Project.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return result, nil
}
