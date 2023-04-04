package DBModels

type FileFolder struct {
	ID             int    `db:"ID" json:"ID"`
	FileFolderName string `db:"FileFolderName" json:"FileFolderName"`
	ParentFolderID int    `db:"ParentFolderID" json:"ParentFolderID"`
	FileStoreID    int    `db:"FileStoreID" json:"FileStoreID"`
	Time           string `db:"Time" json:"Time"`
}
