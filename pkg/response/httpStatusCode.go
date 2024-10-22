package response

const (
	CodeSuccess = 2000
	BadRequest  = 4000
)

var msg = map[int]string{
	CodeSuccess: "Success",
	BadRequest:  "Bad request",
}
