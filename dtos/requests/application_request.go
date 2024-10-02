package requests

type ApplicantionRequest struct {
	JobID uint `json:"job_id" binding:"required"`
}

type ApplicantionUpdateRequest struct {
	State string `json:"state" binding:"required"`
}
