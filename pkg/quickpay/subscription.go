package quickpay

type Subscription struct {
	ID              int         `json:"id"`
	ULID            string      `json:"ulid"`
	MerchantID      int         `json:"merchant_id"`
	OrderID         string      `json:"order_id"`
	Accepted        bool        `json:"accepted"`
	Type            string      `json:"type"`
	TextOnStatement string      `json:"text_on_statement"`
	BrandingID      int         `json:"branding_id"`
	Variables       any         `json:"variables"`
	Currency        string      `json:"currency"`
	State           string      `json:"state"`
	Metadata        Metadata    `json:"metadata"`
	Link            PaymentLink `json:"link"`
	ShippingAddress Address     `json:"shipping_address"`
	InvoiceAddress  Address     `json:"invoice_address"`
	Basket          []Basket    `json:"basket"`
	Shipping        Shipping    `json:"shipping"`
	Operations      []Operation `json:"operations"`
	TestMode        bool        `json:"test_mode"`
	Acquirer        string      `json:"acquirer"`
	Facilitator     string      `json:"facilitator"`
	CreatedAt       string      `json:"created_at"`
	UpdatedAt       string      `json:"updated_at"`
	RetentedAt      string      `json:"retented_at"`
	Description     string      `json:"description"`
	GroupIDs        []int       `json:"group_ids"`
	ThreedsV2       ThreedsV2   `json:"threeds_v2"`
	Unscheduled     *bool       `json:"unscheduled"`
	DeadlineAt      string      `json:"deadline_at"`
}