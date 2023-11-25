package study

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/src/adapter/infra/table"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
	"os"
	"time"
)

type StudyRepositoryImpl struct {
	db    *dynamodb.DynamoDB
	table string
}

// NewStudyRepository は StudyRepository を生成
func NewStudyRepository(db *dynamodb.DynamoDB) repository_study.StudyRepository {
	return &StudyRepositoryImpl{
		db:    db,
		table: os.Getenv("STUDY_TABLE_NAME"),
	}
}

// toStudyTableRecordFromDynamoDBItem は DynamoDBのItemから学習テーブルレコードに変換
func toStudyTableRecordFromDynamoDBItem(study *map[string]*dynamodb.AttributeValue) (*table.StudyTableRecord, error) {
	var studyTableRecord table.StudyTableRecord
	err := dynamodbattribute.UnmarshalMap(*study, &studyTableRecord)
	if err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("学習テーブルレコードの変換に失敗しました。 study: %v", study),
			err,
		)
	}
	return &studyTableRecord, nil
}

// toStudyTableRecordFromEntity は 学習エンティティから学習テーブルレコードに変換
func toStudyTableRecordFromEntity(study *model_study.StudyEntity) (*table.StudyTableRecord, error) {
	// 日付フォーマットを指定
	const dateFormat = "2006-01-02"

	// time.Time型の日付をstring型に変換する
	createdDate := study.CreatedDate.Format(dateFormat)
	updatedDate := study.UpdatedDate.Format(dateFormat)

	return &table.StudyTableRecord{
		ID:          study.ID,
		Title:       study.Title,
		Tags:        study.Tags,
		Content:     study.Content,
		CreatedDate: createdDate,
		UpdatedDate: updatedDate,
	}, nil
}

// toStudyEntity は 学習テーブルレコードを学習エンティティに変換
func toStudyEntity(study *table.StudyTableRecord) (*model_study.StudyEntity, error) {
	// 日付フォーマットを指定
	const dateFormat = "2006-01-02"

	// string型の日付をtime.Time型に変換する
	createdDate, err := time.Parse(dateFormat, study.CreatedDate)
	if err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("作成日の変換に失敗しました。 id: %s, createdDate: %s", study.ID, study.CreatedDate),
			err,
		)
	}

	updatedDate, err := time.Parse(dateFormat, study.UpdatedDate)
	if err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("更新日の変換に失敗しました。 id: %s, updatedDate: %s", study.ID, study.UpdatedDate),
			err,
		)
	}
	return &model_study.StudyEntity{
		ID:          study.ID,
		Title:       study.Title,
		Tags:        study.Tags,
		Content:     study.Content,
		CreatedDate: createdDate,
		UpdatedDate: updatedDate,
	}, nil
}

// CreateStudy はスタディを作成
func (r *StudyRepositoryImpl) CreateStudy(study *model_study.StudyEntity) error {
	studyTableRecord, err := toStudyTableRecordFromEntity(study)
	if err != nil {
		return err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item: map[string]*dynamodb.AttributeValue{
			"id":          {S: aws.String(studyTableRecord.ID)},
			"title":       {S: aws.String(studyTableRecord.Title)},
			"tags":        {S: aws.String(studyTableRecord.Tags)},
			"content":     {S: aws.String(studyTableRecord.Content)},
			"createdDate": {S: aws.String(studyTableRecord.CreatedDate)},
			"updatedDate": {S: aws.String(studyTableRecord.UpdatedDate)},
		},
	}

	_, err = r.db.PutItem(input)
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("スタディの作成に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	return nil
}

// UpdateStudy はスタディを更新
func (r *StudyRepositoryImpl) UpdateStudy(study *model_study.StudyEntity) error {
	studyTableRecord, err := toStudyTableRecordFromEntity(study)
	if err != nil {
		return err
	}
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(studyTableRecord.ID)},
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":title":       {S: aws.String(studyTableRecord.Title)},
			":tags":        {S: aws.String(studyTableRecord.Tags)},
			":content":     {S: aws.String(studyTableRecord.Content)},
			":updatedDate": {S: aws.String(studyTableRecord.UpdatedDate)},
		},
		UpdateExpression: aws.String("SET title = :title, tags = :tags, content = :content, updatedDate = :updatedDate"),
	}

	_, err = r.db.UpdateItem(input)
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("スタディの更新に失敗しました。 id: %s", study.ID),
			err,
		)
	}
	return nil
}

// DeleteStudy はスタディを削除
func (r *StudyRepositoryImpl) DeleteStudy(id string) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	_, err := r.db.DeleteItem(input)
	if err != nil {
		return customerrors.NewInternalServerError(
			fmt.Sprintf("スタディの削除に失敗しました。 id: %s", id),
			err,
		)
	}
	return nil
}

// GetStudyByID はIDでスタディを取得
func (r *StudyRepositoryImpl) GetStudyByID(id string) (*model_study.StudyEntity, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {S: aws.String(id)},
		},
	}

	result, err := r.db.GetItem(input)
	if err != nil {
		return nil, customerrors.NewInternalServerError(
			fmt.Sprintf("スタディの取得に失敗しました。 id: %s", id),
			err,
		)
	}

	studyTableRecord, err := toStudyTableRecordFromDynamoDBItem(&result.Item)
	if err != nil {
		return nil, err
	}
	study, err := toStudyEntity(studyTableRecord)
	if err != nil {
		return nil, err
	}
	return study, nil
}

// GetStudiesByTitleOrTags はタイトルまたはタグでスタディを検索し、GSIを使用して全体を検索する
func (r *StudyRepositoryImpl) GetStudiesByTitleOrTags(title string, tags string, limit int, exclusiveStartKey string) ([]*model_study.StudyEntity, string, error) {
	var lastEvaluatedKeyMap map[string]*dynamodb.AttributeValue

	if exclusiveStartKey != "" {
		lastEvaluatedKeyMap = map[string]*dynamodb.AttributeValue{"ID": {S: aws.String(exclusiveStartKey)}}
	}

	var studies []*model_study.StudyEntity
	var items []map[string]*dynamodb.AttributeValue
	var lastEvaluatedKey map[string]*dynamodb.AttributeValue
	var err error

	if title != "" || tags != "" {
		queryInput := &dynamodb.QueryInput{
			TableName:         aws.String(r.table),
			Limit:             aws.Int64(int64(limit)),
			ExclusiveStartKey: lastEvaluatedKeyMap,
		}

		if title != "" {
			queryInput.IndexName = aws.String("TitleIndex")
			queryInput.KeyConditionExpression = aws.String("title = :title")
			queryInput.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{":title": {S: aws.String(title)}}
		} else if tags != "" {
			queryInput.IndexName = aws.String("TagsIndex")
			queryInput.KeyConditionExpression = aws.String("tags = :tags")
			queryInput.ExpressionAttributeValues = map[string]*dynamodb.AttributeValue{":tags": {S: aws.String(tags)}}
		}

		var result *dynamodb.QueryOutput
		result, err = r.db.Query(queryInput)
		if err != nil {
			return nil, "", customerrors.NewInternalServerError(
				fmt.Sprintf("スタディの検索に失敗しました。 title: %s, tags: %s", title, tags),
				err,
			)
		}

		items = result.Items
		lastEvaluatedKey = result.LastEvaluatedKey
	} else {
		scanInput := &dynamodb.ScanInput{
			TableName:         aws.String(r.table),
			Limit:             aws.Int64(int64(limit)),
			ExclusiveStartKey: lastEvaluatedKeyMap,
		}

		var result *dynamodb.ScanOutput
		result, err = r.db.Scan(scanInput)
		if err != nil {
			return nil, "", customerrors.NewInternalServerError(
				fmt.Sprintf("スタディの検索に失敗しました。 title: %s, tags: %s", title, tags),
				err,
			)
		}

		items = result.Items
		lastEvaluatedKey = result.LastEvaluatedKey
	}

	for _, item := range items {
		studyTableRecord, err := toStudyTableRecordFromDynamoDBItem(&item)
		if err != nil {
			return nil, "", err
		}
		study, err := toStudyEntity(studyTableRecord)
		if err != nil {
			return nil, "", err
		}
		studies = append(studies, study)
	}

	var nextExclusiveStartKey string
	if lastEvaluatedKey != nil {
		nextExclusiveStartKey = *lastEvaluatedKey["ID"].S
	}

	return studies, nextExclusiveStartKey, nil
}
