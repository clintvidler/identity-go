CREATE TABLE IF NOT EXISTS users (
  id serial PRIMARY KEY,
  email VARCHAR (320) UNIQUE NOT NULL,
  password VARCHAR (120) NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS tokens (
  id serial PRIMARY KEY,
  token VARCHAR (1200) NOT NULL,
  app_id VARCHAR (60) NOT NULL DEFAULT 0,
  user_id int NOT NULL REFERENCES users (id),
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS pending_registration (
  id serial PRIMARY KEY,
  key VARCHAR (32) UNIQUE NOT NULL,
  email VARCHAR (320) UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS pending_reset_password (
  id serial PRIMARY KEY,
  key VARCHAR (32) UNIQUE NOT NULL,
  -- user_id int UNIQUE NOT NULL REFERENCES users (id),
  -- email VARCHAR (320) UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE
);

CREATE TABLE IF NOT EXISTS pending_update_email (
  id serial PRIMARY KEY,
  key VARCHAR (32) UNIQUE NOT NULL,
  user_id int UNIQUE NOT NULL REFERENCES users (id),
  email VARCHAR (320) UNIQUE NOT NULL,
  created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE,
  expired_at TIMESTAMP WITH TIME ZONE
);