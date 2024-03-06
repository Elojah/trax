import fs from 'fs';

const configFile = process.env.CONFIG_FILE;

if (!configFile) {
  throw new Error('CONFIG_FILE not set');
}

const configData = fs.readFileSync(configFile, 'utf-8');

if (!configData) {
  throw new Error('CONFIG_FILE is empty');
}

type Config = {
  web_client_url:string;
  api_url:string;
}

export const config:Config = JSON.parse(configData);
export const logger = require('pino')()
