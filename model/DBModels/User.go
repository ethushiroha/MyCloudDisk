package DBModels

type User struct {
	UserID      int    `db:"UserID" json:"UserID"`
	Username    string `db:"Username" json:"Username"`
	Password    string `db:"Password" json:"Password"`
	FileStoreID int    `db:"FileStoreID" json:"FileStoreID"`
}
