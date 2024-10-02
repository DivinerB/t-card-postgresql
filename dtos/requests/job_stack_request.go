package requests

type JobStackRequest struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo"`
}
