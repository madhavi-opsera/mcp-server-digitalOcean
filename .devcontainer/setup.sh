#!/bin/bash
set -e

echo "🔧 Initializing JavaScript MCP Server..."

# Navigate to the JavaScript MCP server directory
if [ -d "MCP/McpServer/javascript" ]; then
    cd MCP/McpServer/javascript
    echo "📦 Installing npm dependencies..."
    npm install
    echo "✅ Dependencies installed successfully"
else
    echo "❌ No MCP/McpServer/javascript directory found. Cannot start server."
    exit 1
fi

# Check for AUTH_BEARER env variable
if [ -z "$AUTH_BEARER" ]; then
  echo "❌ AUTH_BEARER environment variable is not set!"
  echo "👉 Add it in GitHub Codespaces: Repository > Settings > Codespaces > Secrets."
  exit 1
fi

# Start server
if [ -f "main.js" ]; then
    echo "▶️ Running main.js..."
    AUTH_BEARER="$AUTH_BEARER" node main.js
else
    echo "❌ No main.js found in MCP/McpServer/javascript/"
    exit 1
fi
