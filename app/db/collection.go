package db

//Collections .
type Collections struct {
	// Cursor Cursor `json:"cursor"`
	// Ok     int    `json:"ok"`
}

//Collection .
type Collection struct {
	Name  string `json:"name"`
	Types string `json:"type"`
}

// //Cursor .
// type Cursor struct {
// 	ID         string       `json:"id"`
// 	Ns         string       `json:"ns"`
// 	FirstBatch []Collection `json:"firstBatch"`
// }
