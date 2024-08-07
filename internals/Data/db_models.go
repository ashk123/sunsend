package Data

/*
	Special Models for using Database
*/

type Channel struct {
	ID          int
	Name        string
	Description string
	Owner       string
}

// Base Chat Structure
type Message struct {
	CID     int
	MID     int
	Author  string
	Content string
	Date    string
	Image   string
	ReplyID int
}

// type Message struct {
// 	Sender  string
// 	Date    time.Time
// 	Content string
// 	Length  int
// }
