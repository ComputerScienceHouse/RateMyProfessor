-- Placeholder courses table creation (will need to add foreign keys later)
CREATE TABLE IF NOT EXISTS courses(
    id          VARCHAR(5) PRIMARY KEY,
    name        VARCHAR(50) NOT NULL,
    description TEXT NOT NULL,
    college     VARCHAR(100)   
);