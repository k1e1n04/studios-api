CREATE TABLE IF NOT EXISTS studies (
    id VARCHAR(36) NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date DATE NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

ALTER TABLE studies ADD INDEX title_index (title);

CREATE TABLE IF NOT EXISTS tags (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    UNIQUE KEY name_unique (name)
);

ALTER TABLE tags ADD INDEX name_index (name);

CREATE TABLE IF NOT EXISTS study_tags (
    study_id VARCHAR(36) NOT NULL,
    tag_id INT NOT NULL,
    PRIMARY KEY (study_id, tag_id)
);

ALTER TABLE study_tags ADD INDEX idx_study_id (study_id);
ALTER TABLE study_tags ADD INDEX idx_tag_id (tag_id);