/**
 * MCP Server - JavaScript Implementation
 */

import { Server } from '@modelcontextprotocol/sdk/server/index.js';
import { StdioServerTransport } from '@modelcontextprotocol/sdk/server/stdio.js';
import { ListToolsRequestSchema, CallToolRequestSchema } from '@modelcontextprotocol/sdk/types.js';
import fs from 'fs';
import path from 'path';
import os from 'os';

import { post_profiles_username_follow, createPostProfilesUsernameFollowTool } from './tools/profile/followuserbyusername_converted.js';
import { delete_profiles_username_follow, createDeleteProfilesUsernameFollowTool } from './tools/profile/unfollowuserbyusername_converted.js';
import { get_profiles_username, createGetProfilesUsernameTool } from './tools/profile/getprofilebyusername_converted.js';
import { post_articles_slug_favorite, createPostArticlesSlugFavoriteTool } from './tools/favorites/createarticlefavorite_converted.js';
import { delete_articles_slug_favorite, createDeleteArticlesSlugFavoriteTool } from './tools/favorites/deletearticlefavorite_converted.js';
import { put_user, createPutUserTool } from './tools/user_and_authentication/updatecurrentuser_converted.js';
import { get_user, createGetUserTool } from './tools/user_and_authentication/getcurrentuser_converted.js';
import { post_users_login, createPostUsersLoginTool } from './tools/user_and_authentication/login_converted.js';
import { post_users, createPostUsersTool } from './tools/user_and_authentication/createuser_converted.js';
import { get_articles_slug_comments, createGetArticlesSlugCommentsTool } from './tools/comments/getarticlecomments_converted.js';
import { post_articles_slug_comments, createPostArticlesSlugCommentsTool } from './tools/comments/createarticlecomment_converted.js';
import { delete_articles_slug_comments_id, createDeleteArticlesSlugCommentsIdTool } from './tools/comments/deletearticlecomment_converted.js';
import { get_tags, createGetTagsTool } from './tools/tags/gettags_converted.js';
import { get_articles, createGetArticlesTool } from './tools/articles/getarticles_converted.js';
import { get_articles_feed, createGetArticlesFeedTool } from './tools/articles/getarticlesfeed_converted.js';
import { delete_articles_slug, createDeleteArticlesSlugTool } from './tools/articles/deletearticle_converted.js';
import { get_articles_slug, createGetArticlesSlugTool } from './tools/articles/getarticle_converted.js';
import { post_articles, createPostArticlesTool } from './tools/articles/createarticle_converted.js';
import { put_articles_slug, createPutArticlesSlugTool } from './tools/articles/updatearticle_converted.js';

// Create MCP server
const server = new Server({
  name: 'MCP Server',
  version: '1.0.0'
}, {
  capabilities: {
    tools: {}
  }
});

function getConfig() {
  const baseURL = process.env.API_BASE_URL;
  const bearerToken = process.env.API_BEARER_TOKEN;
  
  if (!baseURL || !bearerToken) {
    const configPath = path.join(os.homedir(), '.api', 'config.json');
    try {
      const configData = JSON.parse(fs.readFileSync(configPath, 'utf8'));
      return {
        baseURL: baseURL || configData.baseURL,
        bearerToken: bearerToken || configData.bearerToken
      };
    } catch (e) {
      throw new Error('Configuration not found. Please set API_BASE_URL and API_BEARER_TOKEN environment variables or create config file at ~/.api/config.json');
    }
  }
  
  return { baseURL, bearerToken };
}

// Register all tools
const tools = [
  createPostProfilesUsernameFollowTool(),
  createDeleteProfilesUsernameFollowTool(),
  createGetProfilesUsernameTool(),
  createPostArticlesSlugFavoriteTool(),
  createDeleteArticlesSlugFavoriteTool(),
  createPutUserTool(),
  createGetUserTool(),
  createPostUsersLoginTool(),
  createPostUsersTool(),
  createGetArticlesSlugCommentsTool(),
  createPostArticlesSlugCommentsTool(),
  createDeleteArticlesSlugCommentsIdTool(),
  createGetTagsTool(),
  createGetArticlesTool(),
  createGetArticlesFeedTool(),
  createDeleteArticlesSlugTool(),
  createGetArticlesSlugTool(),
  createPostArticlesTool(),
  createPutArticlesSlugTool()
];

// List all available tools
function listToolsHandler() {
  return { tools: tools.map(tool => tool.definition) };
}

// Handle tool calls
function createCallToolHandler(toolMap) {
  return async (request) => {
    const { name, arguments: args } = request.params;
    
    const tool = toolMap.find(t => t.definition.name === name);
    if (!tool) {
      throw new Error(`Unknown tool: ${name}`);
    }
    
    try {
      const result = await tool.handler(args);
      return {
        content: [{
          type: 'text',
          text: result
        }]
      };
    } catch (error) {
      throw new Error(`Tool execution failed: ${error.message}`);
    }
  };
}

// Setup request handlers
server.setRequestHandler(ListToolsRequestSchema, listToolsHandler);
server.setRequestHandler(CallToolRequestSchema, createCallToolHandler(tools));

async function main() {
  try {
    const config = getConfig();
    console.error('MCP Server started successfully');
    
    const transport = new StdioServerTransport();
    await server.connect(transport);
  } catch (error) {
    console.error('Failed to start server:', error);
    process.exit(1);
  }
}

if (import.meta.url === `file://${process.argv[1]}`) {
  main().catch(console.error);
}