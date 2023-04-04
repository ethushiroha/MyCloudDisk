package DBModels

type FileStore struct {
	FileStoreID int   `db:"StoreID" json:"StoreID"`
	UserID      int   `db:"UserID" json:"UserID"`
	CurrentSize int64 `db:"CurrentSize" json:"CurrentSize"`
	MaxSize     int64 `db:"MaxSize" json:"MaxSize"`
	TopFolderID int   `db:"TopFolderID" json:"TopFolderID"`
}
