ALTER TABLE studies ADD COLUMN number_of_review int NOT NULL DEFAULT 0 AFTER content;
ALTER TABLE studies ADD INDEX index_number_of_review (number_of_review);