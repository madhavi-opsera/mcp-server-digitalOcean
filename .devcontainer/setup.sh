#!/bin/bash
set -e

echo "🔧 Initializing JavaScript MCP Server..."

# Ensure we have Node.js deps
if [ -d "MCP" ] && [ -f "MCP/package.json" ]; then
    cd MCP
    echo "📦 Installing npm dependencies..."
    npm install

    if [ -f "tsconfig.json" ]; then
        echo "🛠️ Building TypeScript project..."
        npm run build
    fi
else
    echo "❌ No MCP/package.json found. Cannot start server."
    exit 1
fi

# Check for AUTH_BEARER env variable
if [ -z "$AUTH_BEARER" ]; then
  echo "❌ AUTH_BEARER environment variable is not set!"
  echo "👉 Add it in GitHub Codespaces: Repository > Settings > Codespaces > Secrets."
  exit 1
fi

# Start server
if [ -f "dist/index.js" ]; then
    echo "▶️ Running dist/index.js..."
    AUTH_BEARER="$AUTH_BEARER" node dist/index.js
elif [ -f "src/index.ts" ]; then
    echo "▶️ Running src/index.ts..."
    AUTH_BEARER="$AUTH_BEARER" npx tsx src/index.ts
elif [ -f "main.js" ]; then
    echo "▶️ Running main.js..."
    AUTH_BEARER="$AUTH_BEARER" node main.js
else
    echo "❌ No valid entry point found in MCP/"
    exit 1
fi
