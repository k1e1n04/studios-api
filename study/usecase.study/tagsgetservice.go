package usecase_study

import (
	"github.com/k1e1n04/studios-api/base/sharedkarnel/model/auth"
	repository_study "github.com/k1e1n04/studios-api/study/domain/repository.study"
)

// TagsGetService は タグ取得サービス
type TagsGetService struct {
	tagRepository repository_study.TagRepository
}

// NewTagsGetService は タグ取得サービスを生成
func NewTagsGetService(tagRepository repository_study.TagRepository) TagsGetService {
	return TagsGetService{
		tagRepository: tagRepository,
	}
}

// Execute は タグを取得
func (tgs *TagsGetService) Execute(name string, userID string) ([]*TagDTO, error) {
	tags, err := tgs.tagRepository.SearchTags(name, *auth.RestoreUserID(userID))
	if err != nil {
		return nil, err
	}
	tagDTOs := make([]*TagDTO, len(tags))
	for i, tag := range tags {
		tagDTOs[i] = toTagDTO(tag)
	}
	return tagDTOs, nil
}
