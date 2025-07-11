package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("parameter: %s (type: %s) is required", name, typ)
}

//Create Openenig

type CreateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}

func (r *CreateOpeningRequest) Validate() error {
	if (r.Role == "" &&	r.Company == "" && r.Location == "" &&	r.Link == "" && 	r.Remote == nil && 	r.Salary <=0) {
		return fmt.Errorf("request body is empty or malformed")
	}
	if r.Role == "" {
		return errParamIsRequired("role", "string")
	}
	if r.Company == "" {
		return errParamIsRequired("company", "string")
	}
	if r.Location == "" {
		return errParamIsRequired("location", "string")
	}
	if r.Link == "" {
		return errParamIsRequired("link", "string")
	}
	// force user to set remote
	if r.Remote == nil {
		return errParamIsRequired("remote", "boolean")
	}
	if r.Salary <=0 {
		return errParamIsRequired("salary", "int64")
	}
	return nil
}



type UpdateOpeningRequest struct {
	Role string `json:"role"`
	Company string `json:"company"`
	Location string `json:"location"`
	Remote *bool `json:"remote"`
	Link string `json:"link"`
	Salary int64 `json:"salary"`
}
func (r *UpdateOpeningRequest) Validate() error {
	if (r.Role != "" ||	r.Company != "" || r.Location != "" ||	r.Link != "" || 	r.Remote != nil || 	r.Salary !=0) {
		return nil
	}
	
	return fmt.Errorf("at least one valid field must be set to update the opening")
}