package db

type ImageMetadata struct {
	FileName string `json:"file_name"`
	Size     int64  `json:"size"`
	Note     string `json:"note"`

	// etc ...
}
