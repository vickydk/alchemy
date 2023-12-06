BEGIN;

CREATE TABLE user (
    id BIGINT NOT NULL AUTO_INCREMENT,
    email VARCHAR (128) NOT NULL DEFAULT '' UNIQUE,
    password VARCHAR (128) NOT NULL DEFAULT '',
    name VARCHAR (32) NOT NULL DEFAULT '',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

INSERT
INTO
    user (email,
          password,
          name)
VALUES ('admin@test.com',
        '$2a$10$z0aFtC6R0SU/AVaBLyW4cezo/r5wBdOH.7dsrgweAndBkijcwHE1m',
        'Admin');


COMMIT;