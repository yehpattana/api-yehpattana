package seeding

// import (
// 	"log"
// 	"strconv"
// 	"strings"
// 	"time"

// 	"github.com/yehpattana/api-yehpattana/modules/data/entities"
// 	"github.com/tealeg/xlsx"
// 	"gorm.io/gorm"
// )

// func SeedingProductsFromExcelToDB(db *gorm.DB) {
// 	// Define the mapping (Excel column name -> Struct field name)
// 	columnMapping := map[string]string{
// 		// products table
// 		"Product Group":    "product_group",
// 		"Product Class":    "product_class",
// 		"1st Season":       "season",
// 		"Launch Date":      "launch_date",
// 		"Description 1":    "name",
// 		"Description 2":    "description",
// 		"Collection":       "collection",
// 		"Category":         "category",
// 		"Brand":            "brand",
// 		"Club/ Generic":    "is_club", // ต้องเขียน fx แปลงค่าจาก str เป็น bool
// 		"Club's Name":      "club_name",
// 		"Age Group":        "gender",
// 		"Pack Size":        "pack_size",
// 		"Fabric Details":   "fabric_content",
// 		"Fabric Type":      "fabric_type",
// 		"Weight":           "weight",
// 		"Current Supplier": "current_supplier",
// 		"Remarks":          "remark",
// 		"Edit By":          "edited_by",
// 		"Edit Date":        "updated_at",
// 		"Create By":        "created_by",
// 		"Create Date":      "created_at",

// 		// product_variants table
// 		"Product Code":   "product_code",
// 		"Master Product": "master_code",

// 		// product_items table
// 		"Size Range": "size", // ต้องเขียน fx แมพค่าว่ามีไซส์อะไรบ้าง แล้วสร้างตามนั้น

// 	}

// 	// Load the Excel file
// 	filePath := "path_to_your_excel_file.xlsx"
// 	xlFile, err := xlsx.OpenFile(filePath)
// 	if err != nil {
// 		log.Fatalf("Failed to open Excel file: %v", err)
// 	}

// 	sheet := xlFile.Sheets[0]
// 	headers := make(map[int]string)
// 	for rowIndex, row := range sheet.Rows {
// 		if rowIndex == 0 {
// 			for cellIndex, cell := range row.Cells {
// 				headers[cellIndex] = columnMapping[cell.String()]
// 			}
// 		} else {
// 			var product entities.Product
// 			var variant entities.ProductVaraint
// 			var item entities.Size

// 			for cellIndex, cell := range row.Cells {
// 				fieldName, ok := headers[cellIndex]
// 				if !ok {
// 					continue
// 				}
// 				value := cell.String()
// 				switch fieldName {
// 				case "product_group":
// 					product.ProductGroup = value
// 				case "product_class":
// 					product.ProductClass = value
// 				case "season":
// 					product.Season = value
// 				case "launch_date":
// 					date, err := convertLaunchDateStringToIsoTime(value)
// 					if err != nil {
// 						log.Fatalf("Failed to convert launch date string to ISO time: %v", err)
// 					}
// 					product.LaunchDate = date
// 				case "name":
// 					product.Name = value
// 				case "description":
// 					product.Description = value
// 				case "collection":
// 					product.Collection = value
// 				case "category":
// 					product.Category = value
// 				case "brand":
// 					product.Brand = value
// 				case "is_club":
// 					product.IsClub = isClubStringConvertor(value)
// 				case "club_name":
// 					product.ClubName = value
// 				case "gender":
// 					product.Gender = value
// 				case "pack_size":
// 					product.PackSize = value
// 				case "fabric_content":
// 					product.FabricContent = value
// 				case "fabric_type":
// 					product.FabricType = value
// 				case "weight":
// 					product.Weight = convertStringToFloat32(value)
// 				case "current_supplier":
// 					product.CurrentSupplier = value
// 				case "remark":
// 					product.Remark = value
// 				case "edited_by":
// 					product.EditedBy = value
// 				case "updated_at":
// 					product.UpdatedAt = value
// 				case "created_by":
// 					product.CreatedBy = value
// 				case "created_at":
// 					product.CreatedAt = value
// 				case "product_code":
// 					variant.ProductCode = value
// 				case "master_code":
// 					variant.MasterCode = value
// 					// TODO write fx to map size range
// 				case "size":
// 					item.Size = value
// 				}
// 			}

// 			// TODO pls modifiy code below this part to fit project =============
// 			// Insert the product into the database
// 			if err := db.Create(&product).Error; err != nil {
// 				log.Printf("Failed to insert product: %v", err)
// 			}

// 			// Insert the product variant into the database
// 			if err := db.Create(&variant).Error; err != nil {
// 				log.Printf("Failed to insert product variant: %v", err)
// 			}

// 			// Insert the product item into the database
// 			if err := db.Create(&item).Error; err != nil {
// 				log.Printf("Failed to insert product item: %v", err)
// 				// }
// 				// ===================================================================
// 			}
// 		}

// 		// Commit the transaction
// 		if err := db.Commit().Error; err != nil {
// 			log.Fatalf("Failed to commit the transaction: %v", err)
// 		}

// 	}
// }

// func isClubStringConvertor(value string) bool {
// 	return value == "Club"
// }

// func convertStringToFloat32(value string) float32 {
// 	// convert string to float32
// 	value = strings.Replace(value, ",", "", -1)
// 	floatValue, err := strconv.ParseFloat(value, 32)
// 	if err != nil {
// 		log.Fatalf("Failed to convert string to float32: %v", err)
// 	}

// 	return float32(floatValue)
// }

// func convertLaunchDateStringToIsoTime(value string) (string, error) {
// 	// Define the layout of the input date string
// 	const inputLayout = "1/2/2006"

// 	// Parse the date string
// 	t, err := time.Parse(inputLayout, value)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Format the time to ISO 8601
// 	return t.Format(time.RFC3339), nil
// }
