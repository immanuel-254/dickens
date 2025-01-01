-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY,
    surname TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    email TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY,
    user_id INTEGER,
    name TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) 
      REFERENCES users (id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
    
);

CREATE TABLE IF NOT EXISTS blogs (
	id INTEGER PRIMARY KEY,
    user_id INTEGER,
    title TEXT NOT NULL,
    body TEXT NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) 
      REFERENCES users (id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
);

CREATE TABLE IF NOT EXISTS category_blogs (
   category_id INTEGER NOT NULL,
   blog_id INTEGER NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (category_id, blog_id),
   FOREIGN KEY (category_id) 
      REFERENCES categories (id) 
        ON DELETE CASCADE 
        ON UPDATE NO ACTION,
   FOREIGN KEY (blog_id) 
      REFERENCES blogs (id) 
        ON DELETE CASCADE 
        ON UPDATE NO ACTION
);


Create Table IF NOT EXISTS logs(
    id INTEGER PRIMARY KEY,
    db_table TEXT NOT NULL,
    action TEXT NOT NULL,
    object_id INTEGER NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

CREATE TABLE IF NOT EXISTS profiles(
	id INTEGER PRIMARY KEY,
    user_id INTEGER,
    username TEXT NOT NULL,
	image TEXT,
	bio TEXT,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY (user_id) 
      REFERENCES users (id) 
         ON DELETE CASCADE 
         ON UPDATE NO ACTION
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS category_blogs;
DROP TABLE IF EXISTS blogs;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS profiles;
DROP TABLE IF EXISTS logs;
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
