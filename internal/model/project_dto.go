package model

type ProjectCreate struct {
	ProjectName        string `json:"project_name" binding:"required"`
	Pinned             bool   `json:"pinned"`
	ProjectDescription string `json:"project_description" binding:"required"`
	Tags               string `json:"tags"`
	LogoPicture        string `json:"logo_picture"`
	GithubURL          string `json:"github_url"`
	SiteURL            string `json:"site_url"`
	OwnerEmail         string `json:"owner_email"`
}
