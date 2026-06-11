package domain
import "time"

type Thread struct {
	ID		 	int
	Subject  	string
	Body 	 	string
	Author 	 	string
	Password	string
	CreatedAt 	time.Time
	BumpedAt 	time.Time
}

type Comment struct {
	ID			int
	ThreadID	int
	Body 		string
	Author		string
	Password	string
	CreatedAt	time.Time
}