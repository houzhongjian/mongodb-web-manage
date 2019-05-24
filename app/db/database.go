package db

//Databases .
type Databases struct {
	Database  []Database `json:"databases"`
	TotalSize int        `json:"totalSize"`
	Ok        int        `json:"ok"`
}

//Database .
type Database struct {
	Name       string `json:"name"`
	SizeOnDisk int    `json:"sizeOnDisk"`
	Empty      bool   `json:"empty"`
}
