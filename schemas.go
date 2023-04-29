package main

type OptionSchema struct {
	Files      []string               `json:"files"`
	Data       map[string]interface{} `json:"data"`
	ID         int                    `json:"id"`
	EntityName string                 `json:"entity_name"`
}

type FileSchema struct {
	Name string `json:"name"`
}
