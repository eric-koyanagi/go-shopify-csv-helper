package validateHeader

import (
	"errors"
	"fmt"
	"strings"
)

var validHeaders = []string{
	"Handle",
	"Title",
	"Body (HTML)",
	"Vendor",
	"Product Category",
	"Type",
	"Tags",
	"Published",
	"Option1 Name",
	"Option1 Value",
	"Option2 Name",
	"Option2 Value",
	"Option3 Name",
	"Option3 Value",
	"Variant SKU",
	"Variant Grams",
	"Variant Inventory Tracker",
	"Variant Inventory Qty",
	"Variant Inventory Policy",
	"Variant Fulfillment Service",
	"Variant Price",
	"Variant Compare At Price",
	"Variant Requires Shipping",
	"Variant Taxable",
	"Variant Barcode",
	"Image Src",
	"Image Position",
	"Image Alt Text",
	"Gift Card",
	"SEO Title",
	"SEO Description",
	"Google Shopping / Google Product Category",
	"Google Shopping / Gender",
	"Google Shopping / Age Group",
	"Google Shopping / MPN",
	"Google Shopping / AdWords Grouping",
	"Google Shopping / AdWords Labels",
	"Google Shopping / Condition",
	"Google Shopping / Custom Product",
	"Google Shopping / Custom Label 0",
	"Google Shopping / Custom Label 1",
	"Google Shopping / Custom Label 2",
	"Google Shopping / Custom Label 3",
	"Google Shopping / Custom Label 4",
	"Variant Image",
	"Variant Weight Unit",
	"Variant Tax Code",
	"Cost per item",
	"Price / International",
	"Compare At Price / International",
	"Status",
}

// each header must exist in the validHeaders defined above
func isValidHeader(val string) bool {
	for _, validHeader := range validHeaders {
		if strings.ToLower(val) == strings.ToLower(validHeader) {
			return true
		}
	}

	return false
}

func Validate(headers []string) (bool, error) {
	valid := true
	var err error
	for _, element := range headers {
		if !isValidHeader(element) {
			return false, errors.New(fmt.Sprintf("header \"%s\" is not valid", element))
		}
	}

	return valid, err
}
