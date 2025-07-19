package server

type Capabilities struct {
	Prompts   Prompts   `json:"prompts"`
	Resources Resources `json:"resources"`
	Tools     Tools     `json:"tools"`
}

type Resources struct {
	Subscribe   bool `json:"subscrible"`
	ListChanged bool `json:"listChanged"`
}

type Prompts struct {
	ListChanged bool `json:"listChanged"`
}

type Tools struct {
	ListChanged bool `json:"listChanged"`
}
