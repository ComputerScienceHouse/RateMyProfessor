-- Placeholder reviews table

CREATE TABLE IF NOT EXISTS reviews(
    id  INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    quality  FLOAT UNSIGNED NOT NULL,
    difficulty  FLOAT UNSIGNED NOT NULL,
    date INT NOT NULL,
    course VARCHAR(5) NOT NULL,
    professor INT UNSIGNED NOT NULL,
    user VARCHAR(50) -- idk how it's formatted from ldap
);