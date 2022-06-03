package admin


type AppointmentListReq struct {
	Day string
}


type SigninAppointmentListReq struct {
	Day string `json:"day" binding:"required"`
}

