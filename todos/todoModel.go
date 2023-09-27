package todos

type TodoModel struct {
	ID        int    `gorm:"column:id" json:"ID"`
	Title     string `gorm:"column:title" json:"title"`
	Completed int    `gorm:"column:completed" json:"completed"`
}
