package http
import "time"

type ThreadResponse struct {
	ID 			int			`json:"id"`
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

type CommentResponse struct {
	ID 			int			`json:"id"`
	ThreadID	int			`json:"thread_id"`
	Body		string		`json:"body"`
	Author		string		`json:"author"`
	CreatedAt	time.Time	`json:"created_at"`
}

type CreateCommentRequest struct {
	Body		string		`json:"body"`
	Author		string		`json:"author"`
	Password	string		`json:"password"`
}

type DeleteCommentRequest struct {
	Password	string		`json:"password"`
}