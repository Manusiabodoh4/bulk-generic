package models

type Users struct {
	ID       string `json:"id,omitempty" gorm:"primaryKey;autoIncrement:true"`
	Tipe     string `json:"tipe"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
}

type RegisterRequest struct {
	Tipe     string `json:"tipe"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
}

type GetDetailRequest struct {
	ID string `json:"id,omitempty"`
}

type UpdateUsersRequest struct {
	Tipe     string `json:"tipe"`
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
	Port     int    `json:"port"`
}
