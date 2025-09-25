/**
 * Get recent articles globally
 */

import fs from 'fs';
import path from 'path';
import os from 'os';

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

export async function get_articles(tag, author, favorited, offset, limit) {
  try {
    const config = getConfig();
    const params = new URLSearchParams();
      if (tag) params.append("tag", tag);
      if (author) params.append("author", author);
      if (favorited) params.append("favorited", favorited);
      if (offset) params.append("offset", offset);
      if (limit) params.append("limit", limit);
    const queryString = params.toString();
    const finalUrl = queryString ? `${url}?${queryString}` : url;
    
    const url = `${config.baseURL}/api/unknown`;
    
    const response = await fetch(finalUrl, {
      method: 'GET',
      headers: {
        'Authorization': `Bearer ${config.bearerToken}`,
        'Accept': 'application/json'
      }
    });
    
    if (!response.ok) {
      return `Failed to format JSON: ${response.status} ${response.statusText}`;
    }
    
    try {
      const result = await response.json();
      return JSON.stringify(result, null, 2);
    } catch (e) {
      return await response.text();
    }
    
  } catch (error) {
    return `Request failed: ${error.message}`;
  }
}

export function createGetArticlesTool() {
  return {
    definition: {
      name: 'get-articles',
      description: 'Get recent articles globally',
      inputSchema: {
        type: 'object',
        properties: {
          tag: {
            type: 'string',
            description: 'Filter by tag'
          },
          author: {
            type: 'string',
            description: 'Filter by author (username)'
          },
          favorited: {
            type: 'string',
            description: 'Filter by favorites of a user (username)'
          },
          offset: {
            type: 'number',
            description: 'The number of items to skip before starting to collect the result set.'
          },
          limit: {
            type: 'number',
            description: 'The numbers of items to return.'
          }
        },
        required: []
      }
    },
    handler: get_articles
  };
}