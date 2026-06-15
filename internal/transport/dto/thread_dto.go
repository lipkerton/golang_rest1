package dto
import "time"

type ThreadResponse struct {
	ID 		int		`json:"id"`
	Subject		string		`json:"subject"`
	Body 		string		`json:"body"`
	Author 		string		`json:"author,omitempty"`
	CreatedAt 	time.Time	`json:"created_at"`
	BumpedAt 	time.Time	`json:"bumped_at"`
}

type CreateThreadRequest struct {
	Subject 	string		`json:"subject"`
	Body 		string		`json:"body"`
	Author 		string		`json:"author"`
	Password 	string		`json:"password"`
}

type DeleteThreadRequest struct {
	Password 	string		`json:"password"`
}
