/*
 * OVH API - EU
 *
 * Build your own OVH world.
 *
 * OpenAPI spec version: 1.0.0
 * Contact: api-subscribe@ml.ovh.net
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package types

import (
	"time"
)

// BillingRefund Details about a Refund
type BillingRefund struct {
	Date *time.Time `json:"date,omitempty"`

	OrderID int64 `json:"orderId,omitempty"`

	OriginalBillID string `json:"originalBillId,omitempty"`

	Password string `json:"password,omitempty"`

	PdfURL string `json:"pdfUrl,omitempty"`

	PriceWithTax *OrderPrice `json:"priceWithTax,omitempty"`

	PriceWithoutTax *OrderPrice `json:"priceWithoutTax,omitempty"`

	RefundID string `json:"refundId,omitempty"`

	Tax *OrderPrice `json:"tax,omitempty"`

	URL string `json:"url,omitempty"`
}
