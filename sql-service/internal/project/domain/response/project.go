package response

type ConvertDDL struct {
	ProjectID    string `json:"project_id"`
	ProjectTitle string `json:"project_title"`
	Data         string `json:"data"`
}
