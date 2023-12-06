type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// gorm.Model 的定义
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type User struct {
	Name string `gorm:"<-:create"`          // 允许读和创建
	Name string `gorm:"<-:update"`          // 允许读和更新
	Name string `gorm:"<-"`                 // 允许读和写（创建和更新）
	Name string `gorm:"<-:false"`           // 允许读，禁止写
	Name string `gorm:"->"`                 // 只读（除非有自定义配置，否则禁止写）
	Name string `gorm:"->;<-:create"`       // 允许读和写
	Name string `gorm:"->:false;<-:create"` // 仅创建（禁止从 db 读）
	Name string `gorm:"-"`                  // 通过 struct 读写会忽略该字段
}
   