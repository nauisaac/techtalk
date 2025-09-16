package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"

	"github.com/OlaIsaac/tech-talk/gateways"
)

func GetCharacterTool() (tool mcp.Tool, handler server.ToolHandlerFunc) {
	tool = mcp.NewTool("get-character",
		mcp.WithDescription("Get a character from the Marvel API"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name of the character"),
		),
	)

	handler = func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		name, err := request.RequireString("name")
		if err != nil {
			return nil, fmt.Errorf("failed to get name: %w", err)
		}

		character, err := gateways.GetCharacter(name)
		if err != nil {
			return nil, fmt.Errorf("failed to get character: %w", err)
		}

		return mcp.NewToolResultText(fmt.Sprintf("Character: %s, Description: %s, Thumbnail: %s.%s", character.Name, character.Description, character.Thumbnail.Path, character.Thumbnail.Extension)), nil
	}

	return tool, handler
}
