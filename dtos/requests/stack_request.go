package requests

type StackRequest struct {
	Name string `json:"name" binding:"required"`
	Logo string `json:"logo"`
}
