package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"bytes"

	"github.com/github-v3-rest-api/mcp-server/config"
	"github.com/github-v3-rest-api/mcp-server/models"
	"github.com/mark3labs/mcp-go/mcp"
)

func Projects_add_collaboratorHandler(cfg *config.APIConfig) func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		args, ok := request.Params.Arguments.(map[string]any)
		if !ok {
			return mcp.NewToolResultError("Invalid arguments object"), nil
		}
		project_idVal, ok := args["project_id"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: project_id"), nil
		}
		project_id, ok := project_idVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: project_id"), nil
		}
		usernameVal, ok := args["username"]
		if !ok {
			return mcp.NewToolResultError("Missing required path parameter: username"), nil
		}
		username, ok := usernameVal.(string)
		if !ok {
			return mcp.NewToolResultError("Invalid path parameter: username"), nil
		}
		queryParams := make([]string, 0)
		// Handle multiple authentication parameters
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("key=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("last_used_after=%s", cfg.BearerToken))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("last_used_before=%s", cfg.BearerToken))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("owner=%s", cfg.BearerToken))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("permission=%s", cfg.BearerToken))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("repository=%s", cfg.BearerToken))
		}
		if cfg.APIKey != "" {
			queryParams = append(queryParams, fmt.Sprintf("secret_type=%s", cfg.APIKey))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("sort=%s", cfg.BearerToken))
		}
		if cfg.BearerToken != "" {
			queryParams = append(queryParams, fmt.Sprintf("token_id=%s", cfg.BearerToken))
		}
		queryString := ""
		if len(queryParams) > 0 {
			queryString = "?" + strings.Join(queryParams, "&")
		}
		// Create properly typed request body using the generated schema
		var requestBody map[string]interface{}
		
		// Optimized: Single marshal/unmarshal with JSON tags handling field mapping
		if argsJSON, err := json.Marshal(args); err == nil {
			if err := json.Unmarshal(argsJSON, &requestBody); err != nil {
				return mcp.NewToolResultError(fmt.Sprintf("Failed to convert arguments to request type: %v", err)), nil
			}
		} else {
			return mcp.NewToolResultError(fmt.Sprintf("Failed to marshal arguments: %v", err)), nil
		}
		
		bodyBytes, err := json.Marshal(requestBody)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to encode request body", err), nil
		}
		url := fmt.Sprintf("%s/projects/%s/collaborators/%s%s", cfg.BaseURL, project_id, username, queryString)
		req, err := http.NewRequest("PUT", url, bytes.NewBuffer(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to create request", err), nil
		}
		// Set authentication based on auth type
		// Handle multiple authentication parameters
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		// API key already added to query string
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Request failed", err), nil
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to read response body", err), nil
		}

		if resp.StatusCode >= 400 {
			return mcp.NewToolResultError(fmt.Sprintf("API error: %s", body)), nil
		}
		// Use properly typed response
		var result map[string]interface{}
		if err := json.Unmarshal(body, &result); err != nil {
			// Fallback to raw text if unmarshaling fails
			return mcp.NewToolResultText(string(body)), nil
		}

		prettyJSON, err := json.MarshalIndent(result, "", "  ")
		if err != nil {
			return mcp.NewToolResultErrorFromErr("Failed to format JSON", err), nil
		}

		return mcp.NewToolResultText(string(prettyJSON)), nil
	}
}

func CreateProjects_add_collaboratorTool(cfg *config.APIConfig) models.Tool {
	tool := mcp.NewTool("put_projects_project_id_collaborators_username",
		mcp.WithDescription("Add project collaborator"),
		mcp.WithNumber("project_id", mcp.Required(), mcp.Description("The unique identifier of the project.")),
		mcp.WithString("username", mcp.Required(), mcp.Description("The handle for the GitHub user account.")),
		mcp.WithString("permission", mcp.Description("Input parameter: The permission to grant the collaborator.")),
	)

	return models.Tool{
		Definition: tool,
		Handler:    Projects_add_collaboratorHandler(cfg),
	}
}
