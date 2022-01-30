package invoice

type InvoiceLine struct {
	Amount           int64                        `json:"amount"`
	Currency         string                       `json:"currency"`
	Description      string                       `json:"description"`
	Discountable     bool                         `json:"discountable"`
	Discounts        int64                        `json:"discounts"`
	DiscountAmounts  []*InvoiceLineDiscountAmount `json:"discount_amounts"`
	ID               string                       `json:"id"`
	InvoiceItem      string                       `json:"invoice_item"`
	Livemode         bool                         `json:"livemode"`
	Metadata         map[string]string            `json:"metadata"`
	Object           string                       `json:"object"`
	Period           *Period                      `json:"period"`
	Plan             int64                        `json:"plan"`
	Price            *Price                       `json:"price"`
	Proration        bool                         `json:"proration"`
	Quantity         int64                        `json:"quantity"`
	Subscription     string                       `json:"subscription"`
	SubscriptionItem string                       `json:"subscription_item"`
	TaxAmounts       []*InvoiceTaxAmount          `json:"tax_amounts"`
	TaxRates         int64                        `json:"tax_rates"`
	Type             string                       `json:"type"`
	UnifiedProration bool                         `json:"unified_proration"`
}

type Price struct {
	Active            bool              `json:"active"`
	BillingScheme     string            `json:"billing_scheme"`
	Created           int64             `json:"created"`
	Currency          string            `json:"currency"`
	Deleted           bool              `json:"deleted"`
	ID                string            `json:"id"`
	Livemode          bool              `json:"livemode"`
	LookupKey         string            `json:"lookup_key"`
	Metadata          map[string]string `json:"metadata"`
	Nickname          string            `json:"nickname"`
	Object            string            `json:"object"`
	Product           string            `json:"product"`
	Recurring         int64             `json:"recurring"`
	TiersMode         int64             `json:"tiers_mode"`
	TransformQuantity int64             `json:"transform_quantity"`
	Type              string            `json:"type"`
	UnitAmount        int64             `json:"unit_amount"`
	UnitAmountDecimal float64           `json:"unit_amount_decimal,string"`
}

type InvoiceTaxAmount struct {
	Amount    int64  `json:"amount"`
	Inclusive bool   `json:"inclusive"`
	TaxRate   string `json:"tax_rate"`
}

type InvoiceLineDiscountAmount struct {
	Amount   int64  `json:"amount"`
	Discount string `json:"discount"`
}

type Period struct {
	End   int64 `json:"end"`
	Start int64 `json:"start"`
}
