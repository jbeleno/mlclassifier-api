CREATE TABLE user (
    `id` INT NOT NULL AUTO_INCREMENT,
    `email` TEXT,
    `password` TEXT,
    `token` TEXT,
    `model_limit` INT,
    `created_at` DATETIME NOT NULL,
    `last_login` DATETIME,
    PRIMARY KEY (id)
);

CREATE TABLE model (
    `id` INT NOT NULL AUTO_INCREMENT,
    `id_user` INT NOT NULL,
    `url` TEXT,
    `training_limit` INT,
    `request_limit` INT,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_user) REFERENCES user(id)
);

CREATE TABLE training (
    `id` INT NOT NULL AUTO_INCREMENT,
    `id_model` INT NOT NULL,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_model) REFERENCES model(id)
);

CREATE TABLE request (
    `id` INT NOT NULL AUTO_INCREMENT,
    `id_model` INT NOT NULL,
    `frequency` TEXT,
    `value` INT,
    `created_at` DATETIME NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (id_model) REFERENCES model(id)
);