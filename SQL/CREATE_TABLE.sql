CREATE TABLE IF NOT EXISTS limbo(
    pk INT NOT NULL AUTO_INCREMENT,
    id UUIDOT NULL,
    PRIMARY KEY(pk),
    UNIQUE(id),

    file_name VARCHAR(255)
    file_Content TEXT
)