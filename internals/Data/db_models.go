package Data

type Channel struct {
	ID          int
	Name        string
	Description string
	Owner       string
}

// Base Chat Structure
type Chat struct {
	CID     int
	MID     int
	Author  string
	Content string
	Date    string
	ReplyID int
}
