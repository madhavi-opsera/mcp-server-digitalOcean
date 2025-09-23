# DigitalOcean MCP Server Codespace

This Codespace configuration automatically starts your DigitalOcean MCP server when you open the repository in GitHub Codespaces.

## Features

- **Automatic MCP Server Detection**: Automatically detects and starts MCP servers written in Node.js, Python, or Go
- **Multi-language Support**: Works with TypeScript/JavaScript, Python, and Go MCP servers
- **Auto-setup**: Installs dependencies and builds the project automatically
- **Port Forwarding**: Automatically forwards ports 3000 and 8080 for MCP server communication

## How to Use

1. **Open in Codespaces**: Click the "Open in Codespaces" badge or use the Code button
2. **Automatic Startup**: The MCP server will start automatically when the Codespace opens
3. **Ready to Use**: The MCP server will be running and accessible

## Supported MCP Server Types

### Node.js/TypeScript
- Looks for `MCP/package.json`
- Installs dependencies with `npm install`
- Builds TypeScript projects with `npm run build`
- Starts with `node dist/index.js` or `npx tsx src/index.ts`

### Python
- Looks for `MCP/requirements.txt`
- Installs dependencies with `pip install -r requirements.txt`
- Starts with `python main.py` or `python mcp_server.py`

### Go
- Looks for `MCP/go.mod`
- Installs dependencies with `go mod tidy`
- Builds with `go build -o mcp-server .`
- Starts with `./mcp-server` or `go run main.go`

## Manual Commands

If you need to manually start the MCP server:

```bash
# Start the MCP server manually
bash .devcontainer/start-mcp.sh

# Or restart the Codespace to trigger automatic startup
```

## Troubleshooting

- **No MCP server found**: Ensure you have an `MCP/` directory with a valid MCP server
- **Dependencies not installed**: The setup script runs automatically, but you can run it manually with `bash .devcontainer/setup.sh`
- **Port conflicts**: The configuration forwards ports 3000 and 8080. Change these in `.devcontainer/devcontainer.json` if needed
