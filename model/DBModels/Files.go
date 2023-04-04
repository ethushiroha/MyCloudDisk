package DBModels

const (
	TEXT  = 1
	IMAGE = 2
	OTHER = 4
)

type MyFile struct {
	ID             int    `db:"ID" json:"ID"`
	FileName       string `db:"FileName" json:"FileName"`
	FileHash       string `db:"FileHash" json:"FileHash"`
	FileStoreID    int    `db:"FileStoreID" json:"FileStoreID"`
	FilePath       string `db:"FilePath" json:"FilePath"`
	DownloadNum    int    `db:"DownloadNum" json:"DownloadNum"`
	UpdateTime     string `db:"UpdateTime" json:"UpdateTime"`
	ParentFolderID int    `db:"ParentFolderID" json:"ParentFolderID"`
	Size           int64  `db:"Size" json:"Size"`
	SizeStr        string `db:"SizeStr" json:"SizeStr"`
	Type           int    `db:"Type" json:"Type"`
	Postfix        string `db:"Postfix" json:"Postfix"`
}
