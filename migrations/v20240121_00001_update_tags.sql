ALTER TABLE tags ADD COLUMN
    user_id varchar(255) NOT NULL;

ALTER TABLE tags ADD INDEX user_id_index (user_id);

ALTER TABLE tags DROP INDEX name_index;

ALTER TABLE tags DROP INDEX name_unique;

ALTER TABLE tags ADD UNIQUE KEY user_id_name_unique (user_id, name);
