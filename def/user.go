package def

type User struct {
	UserID   int64  `gorm:"column:user_id; autoIncrement"`
	Username string `gorm:"column:username"`
}
