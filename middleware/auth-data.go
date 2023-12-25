package middleware

import "encoding/json"

type AuthData struct {
	UserID           uint   `json:"userId"`
	RoleID           uint   `json:"roleId"`
	RoleOriginalName string `json:"roleOriginalName"`
	Phone            string `json:"phone"`
	Scopes           string `json:"scopes"`
}

func (authData *AuthData) LoadFromMap(m map[string]interface{}) error {
	data, err := json.Marshal(m)
	if err == nil {
		err = json.Unmarshal(data, authData)
	}
	return err
}
