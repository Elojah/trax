import Pino from "pino";

type Config = {
  web_client_url: string;
  api_url: string;
  log_level: string;
  google_client_id: string;
}

export const config: Config = JSON.parse(`__CONFIG__`);

export const logger = Pino({
  name: 'trax-client',
  level: config.log_level,
});
