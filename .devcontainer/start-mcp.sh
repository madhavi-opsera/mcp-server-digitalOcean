#!/bin/bash

# Universal start MCP Server script
echo "Starting MCP Server..."

# Function to start Node.js MCP server
start_node_mcp() {
    # Check for nested JavaScript MCP server first
    if [ -f "MCP/McpServer/javascript/package.json" ]; then
        echo "Starting nested JavaScript MCP server..."
        cd MCP/McpServer/javascript
        if [ -f "dist/index.js" ]; then
            node dist/index.js &
        elif [ -f "src/index.ts" ]; then
            npx tsx src/index.ts &
        elif [ -f "main.js" ]; then
            node main.js &
        else
            echo "No valid entry point found for nested JavaScript MCP server"
            return 1
        fi
        cd ../../..
        return 0
    # Check for root-level JavaScript MCP server
    elif [ -f "MCP/package.json" ]; then
        echo "Starting Node.js MCP server..."
        cd MCP
        if [ -f "dist/index.js" ]; then
            node dist/index.js &
        elif [ -f "src/index.ts" ]; then
            npx tsx src/index.ts &
        else
            echo "No valid entry point found for Node.js MCP server"
            return 1
        fi
        cd ..
        return 0
    fi
    return 1
}

# Function to start Python MCP server
start_python_mcp() {
    # Check for nested Python MCP server first
    if [ -f "MCP/McpServer/python/requirements.txt" ]; then
        echo "Starting nested Python MCP server..."
        cd MCP/McpServer/python
        if [ -f "main.py" ]; then
            python main.py &
        elif [ -f "mcp_server.py" ]; then
            python mcp_server.py &
        else
            echo "No valid entry point found for nested Python MCP server"
            return 1
        fi
        cd ../../..
        return 0
    # Check for root-level Python MCP server
    elif [ -f "MCP/requirements.txt" ]; then
        echo "Starting Python MCP server..."
        cd MCP
        if [ -f "main.py" ]; then
            python main.py &
        elif [ -f "mcp_server.py" ]; then
            python mcp_server.py &
        else
            echo "No valid entry point found for Python MCP server"
            return 1
        fi
        cd ..
        return 0
    fi
    return 1
}

# Function to start Go MCP server
start_go_mcp() {
    # Check for Go MCP server in various locations
    if [ -f "MCP/go.mod" ]; then
        echo "Starting Go MCP server from MCP/..."
        cd MCP
        if [ -f "mcp-server" ]; then
            ./mcp-server &
        elif [ -f "main.go" ]; then
            go run main.go &
        else
            echo "No valid entry point found for Go MCP server in MCP/"
            return 1
        fi
        cd ..
        return 0
    elif [ -f "MCP/McpServer/go/go.mod" ]; then
        echo "Starting Go MCP server from MCP/McpServer/go/..."
        cd MCP/McpServer/go
        if [ -f "mcp-server" ]; then
            ./mcp-server &
        elif [ -f "main.go" ]; then
            go run main.go &
        else
            echo "No valid entry point found for Go MCP server in MCP/McpServer/go/"
            return 1
        fi
        cd ../../..
        return 0
    fi
    return 1
}

# Try to start MCP server based on detected type
if start_node_mcp; then
    echo "Node.js MCP server started successfully"
elif start_python_mcp; then
    echo "Python MCP server started successfully"
elif start_go_mcp; then
    echo "Go MCP server started successfully"
else
    echo "No MCP server found or failed to start"
    echo "Make sure you have an MCP directory with a valid MCP server"
fi

# Keep the script running
wait
