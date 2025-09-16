package prompts

import (
	"context"
	"fmt"

	"github.com/OlaIsaac/tech-talk/consts"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func CreatePR() (mcp.Prompt, server.PromptHandlerFunc) {
	prompt := mcp.Prompt{
		Name:        "create-pr",
		Description: "Create a pull request template",
		Arguments: []mcp.PromptArgument{
			{
				Name:        "title",
				Description: "The title of the pull request",
				Required:    true,
			},
			{
				Name:        "description",
				Description: "The description of the pull request",
				Required:    false,
			},
		},
	}

	handler := func(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
		title, ok := request.Params.Arguments["title"]
		if !ok || title == "" {
			return nil, fmt.Errorf("title argument is required")
		}

		description := request.Params.Arguments["description"]

		prTemplate := fmt.Sprintf(consts.TemplatePR, title, description)

		return &mcp.GetPromptResult{
			Description: "Pull request template",
			Messages: []mcp.PromptMessage{
				{
					Role:    "user",
					Content: mcp.TextContent{Type: "text", Text: prTemplate},
				},
			},
		}, nil
	}

	return prompt, handler
}
