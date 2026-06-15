package dto
import "time"

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
