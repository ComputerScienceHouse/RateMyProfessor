-- Placeholder reviews table

CREATE TABLE IF NOT EXISTS reviews(
    id  INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    rating  INT UNSIGNED NOT NULL,
    date VARCHAR(30) NOT NULL,
    course VARCHAR(5) NOT NULL,
    professor INT UNSIGNED NOT NULL,
    user VARCHAR(50) -- idk how it's formatted from ldap
);