package quickpay

type Operation struct {
	ID                   int    `json:"id"`
	Type                 string `json:"type"`
	Amount               int    `json:"amount"`
	Pending              bool   `json:"pending"`
	QPStatusCode         string `json:"qp_status_code"`
	QPStatusMsg          string `json:"qp_status_msg"`
	AQStatusCode         string `json:"aq_status_code"`
	AQStatusMsg          string `json:"aq_status_msg"`
	Data                 any    `json:"data"`
	CallbackURL          string `json:"callback_url"`
	CallbackSuccess      bool   `json:"callback_success"`
	CallbackResponseCode string `json:"callback_response_code"`
	CallbackDuration     string `json:"callback_duration"`
	Acquirer             string `json:"acquirer"`
	SecureStatus3D       string `json:"3d_secure_status"`
	CallbackAt           string `json:"callback_at"`
	CreatedAt            string `json:"created_at"`
}