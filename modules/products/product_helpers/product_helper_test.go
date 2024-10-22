package producthelpers

import (
	"testing"

	"github.com/natersland/b2b-e-commerce-api/modules/data/entities"
	"github.com/stretchr/testify/assert"
)

func TestCollectAndSortSizes(t *testing.T) {

	t.Run("should return sorted sizes when have many sizes", func(t *testing.T) {
		// Arrange
		productID := "1"
		productVariants := []entities.Stock{
			{
				ProductId: "1",
				Id:        "1",
				Size:      "L",
			},
			{
				ProductId: "1",
				Id:        "2",
				Size:      "XS",
			},
			{
				ProductId: "1",
				Id:        "3",
				Size:      "M",
			},
			{
				ProductId: "2",
				Id:        "4",
				Size:      "XL",
			},
		}

		// Act
		result := CollectAndSortSizes(productID, productVariants)

		// Assert
		assert.Equal(t, "XS - L", result)
	},
	)

	t.Run("should return sorted sizes when have one size", func(t *testing.T) {
		// Arrange
		productID := "1"
		productVariants := []entities.Stock{
			{
				ProductId: "1",
				Id:        "1",
				Size:      "M",
			},
			{
				ProductId: "2",
				Id:        "2",
				Size:      "L",
			},
		}

		// Act
		result := CollectAndSortSizes(productID, productVariants)

		// Assert
		assert.Equal(t, "M", result)
	},
	)

	t.Run("should return - when have no size", func(t *testing.T) {
		// Arrange
		productID := "1"
		productVariants := []entities.Stock{}

		// Act
		result := CollectAndSortSizes(productID, productVariants)

		// Assert
		assert.Equal(t, "-", result)
	},
	)
}

func TestCollectColorCodes(t *testing.T) {

	t.Run("should return color codes", func(t *testing.T) {
		// Arrange
		productID := "1"

		productVariants := []entities.Product{
			{
				MasterCode: "1",
				Id:         "1",
				ColorCode:  "#8386BD",
			},
			{
				MasterCode: "2",
				Id:         "2",
				ColorCode:  "#B7FF2B",
			},
			{
				MasterCode: "1",
				Id:         "3",
				ColorCode:  "#700000",
			},
		}

		// Act
		result := CollectColorCodes(productID, productVariants)

		// Assert
		assert.Equal(t, []string{"#8386BD", "#700000"}, result)
	},
	)

	t.Run("should return null when have no color codes", func(t *testing.T) {
		// Arrange
		productID := "1"

		productVariants := []entities.Product{
			{
				MasterCode: "2",
				Id:         "2",
				ColorCode:  "#B7FF2B",
			},
		}

		// Act
		result := CollectColorCodes(productID, productVariants)

		// Assert
		assert.Nil(t, result)
	})
}
