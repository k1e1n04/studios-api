CREATE TABLE studies_review_setting (
    first_review_interval int NOT NULL COMMENT '一回目の復習までの期間',
    second_review_interval int NOT NULL COMMENT '二回目の復習までの期間',
    third_review_interval int NOT NULL COMMENT '三回目の復習までの期間'
);
INSERT INTO studies_review_setting VALUES (1, 3, 7);