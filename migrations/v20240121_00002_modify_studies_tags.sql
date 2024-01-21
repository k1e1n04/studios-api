ALTER TABLE studies
    MODIFY id varchar(26) NOT NULL COMMENT '学習ID',
    MODIFY title varchar(255) NOT NULL COMMENT 'タイトル',
    MODIFY content text NOT NULL COMMENT '内容',
    MODIFY user_id varchar(255) NOT NULL COMMENT 'ユーザーID',
    MODIFY number_of_review int DEFAULT 0 NOT NULL COMMENT '復習回数',
    MODIFY created_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '作成日時',
    MODIFY updated_date timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時';

ALTER TABLE study_tags
    MODIFY study_id varchar(26) NOT NULL COMMENT '学習ID',
    MODIFY tag_id varchar(26) NOT NULL COMMENT 'タグID';

ALTER TABLE tags
    MODIFY id varchar(26) NOT NULL COMMENT 'タグID',
    MODIFY name varchar(255) NOT NULL COMMENT '名前',
    MODIFY user_id varchar(255) NOT NULL COMMENT 'ユーザーID';

