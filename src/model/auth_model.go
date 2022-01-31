package model

type SigninWithUsername struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (swu SigninWithUsername) Validate() error {
	if err := validate.Struct(swu); err != nil {
		return err
	}

	return nil
}
