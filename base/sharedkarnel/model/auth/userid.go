package auth

// UserID は ユーザーID
type UserID struct {
	// Value は 値
	Value string
}

// RestoreUserID は ユーザーIDを復元
func RestoreUserID(value string) *UserID {
	return &UserID{
		Value: value,
	}
}
