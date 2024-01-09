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
	InvalidEmailOrPassword       = "メールアドレスまたはパスワードが不正です。"
	PasswordNotMatch             = "パスワードが一致しません。"
	AgreeToTerms                 = "利用規約に同意してください。"
	EmailRequired                = "メールアドレスを入力してください。"
	InvalidEmail                 = "メールアドレスの形式が不正です。"
	InvalidUsername              = "ユーザー名は3文字以上20文字以下で入力してください。"
	TooShortPassword             = "パスワードは8文字以上で入力してください。"
	InvalidToken                 = "トークンが無効です。"
	// 学習
	StudyNotFound = "学習が存在しません。"
	// ユーザー
	UserNotFound      = "ユーザーが存在しません。"
	UserAlreadyExists = "ユーザーが既に存在しています。"
	InvalidPassword   = "パスワードが不正です。"
)
