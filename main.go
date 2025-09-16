package main

import (
	"fmt"

	"github.com/OlaIsaac/tech-talk/tools"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server

	fmt.Println("Creating MCP server")
	s := server.NewMCPServer(
		"Calculator Demo",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	calcTool, calcHandler := tools.GetCalculatorTool()
	s.AddTool(calcTool, calcHandler)

	marvelTool, marvelHandler := tools.GetCharacterTool()
	s.AddTool(marvelTool, marvelHandler)

	// Start the server
	httpServer := server.NewStreamableHTTPServer(s)
	fmt.Println("Server started on port 8885")
	err := httpServer.Start(":8885")
	if err != nil {
		fmt.Printf("Server error: %v\n", err)
	}

}
