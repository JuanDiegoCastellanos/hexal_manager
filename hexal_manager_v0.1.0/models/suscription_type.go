package models

type SuscriptionType struct {
	ID                      string  `json:"id"`
	Name                    string  `json:"name"`
	Details                 string  `json:"details"`
	ProjectsCapacity        float64 `json:"projects_capacity"`
	EmployeesCapacity       float64 `json:"employee_capacity"`
	ClientsCapacity         float64 `json:"clients_capacity"`
	ReportsCapacity         float64 `json:"reports_capacity"`
	MeetingsCapacity        float64 `json:"meetings_capacity"`
	ProblemsCapacity        float64 `json:"problems_capacity"`
	HitsCapacity            float64 `json:"hits_capacity"`
	ManageReleasesCapacity  float64 `json:"manage_releases_capacity"`
	ProjectProgressCapacity float64 `json:"project_progress_capacity"`
}
