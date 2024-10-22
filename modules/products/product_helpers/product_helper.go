package producthelpers

import (
	"strconv"

	"github.com/yehpattana/api-yehpattana/modules/data/entities"
)

// import (
// 	"github.com/yehpattana/api-yehpattana/modules/data/entities"
// )

func CheckIsValidSize(size string) bool {
	validSize := []string{"2XS", "XS", "S", "M", "L", "XL", "2XL", "3XL", "4XL", "5XL", "Free Size", "Other"}
	isValidSize := false
	for _, validSizeValue := range validSize {
		if size == validSizeValue {
			isValidSize = true
			break
		}
	}

	return isValidSize
}

func CheckIsValidStatus(status string) bool {
	validStatus := []string{"available", "hidden", "out_of_stock"}
	isValidStatus := false
	for _, validStatusValue := range validStatus {
		if status == validStatusValue {
			isValidStatus = true
			break
		}
	}

	return isValidStatus
}

// // Define the order of sizes
var SizeOrder = map[string]int{
	"2XS":       0,
	"XS":        1,
	"S":         2,
	"M":         3,
	"L":         4,
	"XL":        5,
	"2XL":       6,
	"3XL":       7,
	"4XL":       8,
	"5XL":       9,
	"Free Size": 10,
	"Other":     11,
}
var ProductOrder = map[string]int{
	"shirts": 0,
	"pants":  1,
	"socks":  2,
	"":       3,
	" ":      4,
}

// SizeSorter implements sort.Interface for []string based on the sizeOrder
type SizeSorter []string

func (s SizeSorter) Len() int {
	return len(s)
}

func (s SizeSorter) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SizeSorter) Less(i, j int) bool {
	return SizeOrder[s[i]] < SizeOrder[s[j]]
}

// CollectAndSortSizes collects and sorts sizes for a given product
func CollectAndSortSizes(productId string, stockItem []entities.Stock) string {
	println("stockItem: ", stockItem)
	// Collect sizes for the given product
	sizeSet := make(map[string]struct{})
	println("sizeSet: ", len(sizeSet))
	for _, item := range stockItem {
		if item.ProductId == productId {
			sizeSet[item.Size] = struct{}{}
		}
	}

	// If there are no sizes, return "-"
	if len(sizeSet) == 0 {
		return "-"
	}

	// If there's only one size, return that size directly
	if len(sizeSet) == 1 {
		for size := range sizeSet {
			return size
		}
	}

	// Find the smallest and largest sizes
	smallestSize := ""
	largestSize := ""
	for size := range sizeSet {
		if smallestSize == "" || SizeOrder[size] < SizeOrder[smallestSize] {
			smallestSize = size
		}
		if largestSize == "" || SizeOrder[size] > SizeOrder[largestSize] {
			largestSize = size
		}
	}

	// Construct the result string
	result := smallestSize + " - " + largestSize
	return result
}

func CollectColorCodes(masterCode string, productVariants []entities.Product) []string {
	var colorCodes []string
	for _, variant := range productVariants {
		if variant.MasterCode == masterCode {
			colorCodes = append(colorCodes, variant.ColorCode)
		}
	}
	return colorCodes
}

func CollectPriceRange(masterCode string, product []entities.Product) string {
	var minPrice float64
	var maxPrice float64
	for _, variant := range product {
		if variant.MasterCode == masterCode {
			if minPrice == 0 || variant.Price < minPrice {
				minPrice = float64(variant.Price)
			}
			if maxPrice == 0 || variant.Price > maxPrice {
				maxPrice = float64(variant.Price)
			}
		}
	}
	if minPrice == 0 && maxPrice == 0 {
		return convertFloatToString(minPrice)
	}
	if minPrice == maxPrice {
		return convertFloatToString(minPrice)
	}
	return convertFloatToString(minPrice) + " - " + convertFloatToString(maxPrice)
}

func convertFloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
