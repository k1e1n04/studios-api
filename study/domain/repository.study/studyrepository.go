package repository_study

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/customerrors"
	"github.com/k1e1n04/studios-api/src/adapter/infra/table"
	model_study "github.com/k1e1n04/studios-api/study/domain/model.study"
	"os"
	"time"
)

type StudyRepository struct {
	db    *dynamodb.DynamoDB
	table string
}

func NewStudyRepository(db *dynamodb.DynamoDB) *StudyRepository {
	return &StudyRepository{
		db:    db,
		table: os.Getenv("STUDY_TABLE_NAME"),
	}
}

// toStudyTableRecordFromDynamoDBItem は DynamoDBのItemから学習テーブルレコードに変換
func toStudyTableRecordFromDynamoDBItem(study *map[string]*dynamodb.AttributeValue) (*table.StudyTableRecord, error) {
	var studyTableRecord table.StudyTableRecord
	err := dynamodbattribute.UnmarshalMap(*study, &studyTableRecord)
	if err != nil {
		return nil, err
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
func (r *StudyRepository) CreateStudy(study *model_study.StudyEntity) error {
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
func (r *StudyRepository) UpdateStudy(study *model_study.StudyEntity) error {
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
func (r *StudyRepository) DeleteStudy(id string) error {
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
func (r *StudyRepository) GetStudyByID(id string) (*model_study.StudyEntity, error) {
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
func (r *StudyRepository) GetStudiesByTitleOrTags(title string, tags string, limit int, exclusiveStartKey string) ([]*model_study.StudyEntity, string, error) {
	var queryInput *dynamodb.QueryInput
	var lastEvaluatedKeyMap map[string]*dynamodb.AttributeValue

	// ExclusiveStartKey を設定（ページネーションの開始点）
	if exclusiveStartKey != "" {
		lastEvaluatedKeyMap = map[string]*dynamodb.AttributeValue{
			"ID": {S: aws.String(exclusiveStartKey)},
		}
	}

	// タイトルやタグによってクエリを構築
	if title != "" {
		queryInput = &dynamodb.QueryInput{
			TableName:                 aws.String(r.table),
			IndexName:                 aws.String("TitleIndex"),
			KeyConditionExpression:    aws.String("Title = :title"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":title": {S: aws.String(title)}},
			Limit:                     aws.Int64(int64(limit)),
			ExclusiveStartKey:         lastEvaluatedKeyMap,
		}
	} else if tags != "" {
		queryInput = &dynamodb.QueryInput{
			TableName:                 aws.String(r.table),
			IndexName:                 aws.String("TagsIndex"),
			KeyConditionExpression:    aws.String("Tags = :tags"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{":tags": {S: aws.String(tags)}},
			Limit:                     aws.Int64(int64(limit)),
			ExclusiveStartKey:         lastEvaluatedKeyMap,
		}
	} else {
		// ここで全件取得のための設定を行う（省略）
		queryInput = &dynamodb.QueryInput{
			TableName:         aws.String(r.table),
			Limit:             aws.Int64(int64(limit)),
			ExclusiveStartKey: lastEvaluatedKeyMap,
		}
	}

	// クエリの実行
	result, err := r.db.Query(queryInput)
	if err != nil {
		return nil, "", customerrors.NewInternalServerError(
			fmt.Sprintf("スタディの取得に失敗しました。 title: %s, tags: %s", title, tags),
			err,
		)
	}

	// 結果の変換
	var studies []*model_study.StudyEntity
	for _, item := range result.Items {
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

	// LastEvaluatedKey を次のページネーションキーとして設定
	var nextExclusiveStartKey string
	if result.LastEvaluatedKey != nil {
		nextExclusiveStartKey = *result.LastEvaluatedKey["ID"].S
	}

	return studies, nextExclusiveStartKey, nil
}
