package entity

type Book struct {
	ID          int    `gorm:"primary_key:auto_increment:not null" json:"id"`
	Title       string `gorm:"type:varchar(255);not null" json:"title"`
	Description string `gorm:"type:text;" json:"description"`
	UserID      int    `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE;onDelete:CASCADE" json:"user"`
}
