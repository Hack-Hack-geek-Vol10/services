package domain

type UploadImageParam struct {
	Key         string `json:"key"`
	ContentType string `json:"content_type"`
	Body        []byte `json:"body"`
}
