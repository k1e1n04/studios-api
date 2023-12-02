package base

const (
	InternalServerError = "サーバーエラーが発生しました。"
	BadRequestError     = "リクエストが不正です。"
	ConflictError       = "他のユーザーによって操作が行われたため、処理を中断しました。"
	InvalidJSONError    = "JSONの形式が不正です。"
	InvalidPageNumber   = "ページ番号が不正です。"
	InvalidLimit        = "ページサイズが不正です。"
	InvalidSize         = "サイズが不正です。"
	// 認証
	AuthenticationHeaderRequired = "Authorizationヘッダーは必須です。"
	InvalidAuthenticationHeader  = "Authorizationヘッダーの形式が正しくありません。"
	InvalidAPIKey                = "APIキーが無効です。"
	// 学習
	StudyNotFound = "学習が存在しません。"
)
