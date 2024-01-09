package models

type ProjectResource struct {
	ID         string `json:"id"`
	ProjectId  string `json:"project_id"`
	ResourceId string `json:"resource:id"`
}
