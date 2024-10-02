package requests

type ApplicantionRequest struct {
	JobID uint `json:"job_id" binding:"required"`
}
