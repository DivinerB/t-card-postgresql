package requests

type JobRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Geoposition string `json:"geoposition" binding:"required"`
	Salary      int    `json:"salary" binding:"required"`
	Money       string `json:"money" binding:"required"`
}
