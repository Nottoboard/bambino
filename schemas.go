package main

type OptionSchema struct {
	Files      []string               `json:"files"`
	Fields     []string               `json:"fields"`
	Data       map[string]interface{} `json:"data"`
	ID         int                    `json:"id"`
	EntityName string                 `json:"entity_name"`
}

type FileSchema struct {
	Name string `json:"name"`
}

type FileRespSchema struct {
	ID         uint                   `json:"id"`
	FileName   string                 `json:"file_name"`
	Name       string                 `json:"name"`
	Data       map[string]interface{} `json:"data"`
	EntityName string                 `json:"entity_name"`
	Url        string                 `json:"url"`
	Path       string                 `json:"path"`
	CreatedOn  string                 `json:"created_on"`
}
