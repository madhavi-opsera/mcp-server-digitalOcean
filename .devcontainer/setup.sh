#!/bin/bash

# Universal setup script for MCP Server Codespace
echo "Setting up Universal MCP Server environment..."

# Update package manager
sudo apt-get update

# Install additional dependencies
sudo apt-get install -y curl wget git build-essential

# Check if we're in a repository with MCP server
if [ -d "MCP" ]; then
    echo "Found MCP directory, setting up MCP server..."
    
    # Check for nested JavaScript structure first
    if [ -d "MCP/McpServer/javascript" ]; then
        echo "Found nested JavaScript MCP server in MCP/McpServer/javascript/"
        cd MCP/McpServer/javascript
        
        if [ -f "package.json" ]; then
            echo "Installing Node.js dependencies..."
            npm install
            
            # Build if TypeScript
            if [ -f "tsconfig.json" ]; then
                echo "Building TypeScript project..."
                npm run build
            fi
            echo "JavaScript MCP server setup complete"
        fi
        
        cd ../../..
    # Check for nested Go structure
    elif [ -d "MCP/McpServer/go" ]; then
        echo "Found nested Go MCP server in MCP/McpServer/go/"
        cd MCP/McpServer/go
        
        if [ -f "go.mod" ]; then
            echo "Installing Go dependencies..."
            go mod tidy
            go build -o mcp-server .
            echo "Go MCP server built successfully"
        fi
        
        cd ../../..
    # Check for nested Python structure
    elif [ -d "MCP/McpServer/python" ]; then
        echo "Found nested Python MCP server in MCP/McpServer/python/"
        cd MCP/McpServer/python
        
        if [ -f "requirements.txt" ]; then
            echo "Installing Python dependencies..."
            pip install -r requirements.txt
            echo "Python MCP server setup complete"
        fi
        
        cd ../../..
    else
        # Navigate to MCP directory
        cd MCP
        
        # Check if it's a Node.js MCP server
        if [ -f "package.json" ]; then
            echo "Installing Node.js dependencies..."
            npm install
            
            # Build if TypeScript
            if [ -f "tsconfig.json" ]; then
                echo "Building TypeScript project..."
                npm run build
            fi
        fi
        
        # Check if it's a Python MCP server
        if [ -f "requirements.txt" ]; then
            echo "Installing Python dependencies..."
            pip install -r requirements.txt
        fi
        
        # Check if it's a Go MCP server
        if [ -f "go.mod" ]; then
            echo "Installing Go dependencies..."
            go mod tidy
            go build -o mcp-server .
        fi
        
        cd ..
    fi
fi

echo "Setup complete! MCP server will start automatically when Codespace opens."
