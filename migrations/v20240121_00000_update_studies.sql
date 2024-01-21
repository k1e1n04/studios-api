ALTER TABLE studies ADD COLUMN
    user_id varchar(255) NOT NULL AFTER content;

ALTER TABLE studies ADD INDEX user_id_index (user_id);

ALTER TABLE studies DROP INDEX title_index;

ALTER TABLE studies ADD INDEX user_title_index (user_id, title);

ALTER TABLE studies DROP INDEX index_number_of_review;

ALTER TABLE studies ADD INDEX user_number_of_review_index (user_id, number_of_review);