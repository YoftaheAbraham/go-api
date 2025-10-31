package controllers

type ControllerResponse struct {
	Message string
}

func HandleReq() ControllerResponse {
	return ControllerResponse{Message: "Success"}
}
