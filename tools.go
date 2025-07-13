package server

import(
	"fmt"
)

type ToolParams struct{
	Name string `json:"name"`
	Arguments map[string]string `json:"arguments"`
}
type Tool struct {
	Name string `json:"name"`
	Title string `json:"title"`
	Description string `json:"description"`
	InputSchema map[string]any `json:"inputSchema"`
}

type ListToolResponse struct {
	Tools []Tool `json:"tools"`
	NextCursor string `json:"nextCursor"`
}

func ToolsError(err error) map[string]any {
	return map[string]any{
		"content": []map[string]string{
			{
				"type": "text",
				"text": fmt.Sprint(err),
			},
		},
		"isError": true,
	}
}

func ToolsResponse(text string) map[string]any{
	return map[string]any{
		"content": []map[string]string{
			{
				"type": "text",
				"text": text,
			},
		},
	}
}
