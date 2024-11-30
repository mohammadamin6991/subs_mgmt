CREATE TABLE sessions (
  id VARCHAR(255) PRIMARY KEY NOT NULL,
  user_email VARCHAR(255) NOT NULL,
  refresh_token VARCHAR(512) NOT NULL,
  is_revoked BOOLEAN NOT NULL DEFAULT FALSE,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  expires_at TIMESTAMP WITHOUT TIME ZONE
);
