CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username varchar(32) UNIQUE NOT NULL,
  password_hash varchar(100) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE scores (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  score NUMERIC(5,2) NOT NULL,
  accuracy NUMERIC(5,2) NOT NULL,
  duration BIGINT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);