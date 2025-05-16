CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username varchar(32) UNIQUE NOT NULL,
  password_hash varchar(255) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE messages (
  id SERIAL PRIMARY KEY,
  user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  message TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW()
);