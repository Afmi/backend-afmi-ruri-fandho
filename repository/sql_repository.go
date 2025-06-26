package repository

import (
	"fmt"
	"math"
	"strings"
	"user-test/helpers"
	"user-test/infra/database"
	"user-test/infra/logger"

	"gorm.io/gorm/clause"
)

type SQLRepository struct{}

func (sql SQLRepository) Get(model interface{}, payload helpers.Payload, preloads ...string) (interface{}, int, int, error) {
	db := database.DB.Debug()

	// Filter: Hanya ambil data yang belum dihapus
	db = db.Where("is_deleted = ?", false)

	// Apply Filters
	for _, filter := range payload.Filters {
		column := filter.Field
		value := filter.Value
		operator := filter.ComparisonOperator
		dataType := filter.Type

		// Deteksi tipe kolom
		isDateColumn := dataType == "date" || (dataType == "" && (strings.Contains(strings.ToLower(column), "date") ||
			strings.Contains(strings.ToLower(column), "tanggal") ||
			strings.HasSuffix(strings.ToLower(column), "_at")))

		isUUIDColumn := dataType == "uuid" || (dataType == "" &&
			strings.Contains(strings.ToLower(column), "id") &&
			strings.Contains(value, "-"))

		isBoolColumn := dataType == "boolean" || (dataType == "" && (strings.HasPrefix(strings.ToLower(column), "is_") ||
			strings.HasPrefix(strings.ToLower(column), "has_") ||
			value == "true" || value == "false"))

		isNumericColumn := dataType == "number" || (dataType == "" && (strings.Contains(strings.ToLower(column), "amount") ||
			strings.Contains(strings.ToLower(column), "price") ||
			strings.Contains(strings.ToLower(column), "total") ||
			strings.Contains(strings.ToLower(column), "qty") ||
			strings.Contains(strings.ToLower(column), "count")))

		// Jika type adalah string, override deteksi otomatis
		if dataType == "string" {
			isDateColumn = false
			isUUIDColumn = false
			isBoolColumn = false
			isNumericColumn = false
		}

		switch operator {
		case "=":
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) = DATE(?)", column), value)
			case isUUIDColumn:
				db = db.Where(fmt.Sprintf("%s = ?", column), value)
			case isBoolColumn:
				boolValue := strings.ToLower(value) == "true"
				db = db.Where(fmt.Sprintf("%s = ?", column), boolValue)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) = ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s = ?", column), value)
			}

		case "like":
			if isNumericColumn {
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) LIKE ?", column), "%"+value+"%")
			} else {
				db = db.Where(fmt.Sprintf("%s LIKE ?", column), "%"+value+"%")
			}

		case "ilike":
			if isNumericColumn {
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) ILIKE ?", column), "%"+value+"%")
			} else {
				db = db.Where(fmt.Sprintf("%s ILIKE ?", column), "%"+value+"%")
			}

		case ">":
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) > DATE(?)", column), value)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("%s > ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s > ?", column), value)
			}

		case "<":
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) < DATE(?)", column), value)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("%s < ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s < ?", column), value)
			}

		case ">=":
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) >= DATE(?)", column), value)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("%s >= ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s >= ?", column), value)
			}

		case "<=":
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) <= DATE(?)", column), value)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("%s <= ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s <= ?", column), value)
			}

		case "between":
			valueParts := strings.Split(value, "--")
			if len(valueParts) == 2 {
				switch {
				case isDateColumn:
					db = db.Where(fmt.Sprintf("DATE(%s) BETWEEN DATE(?) AND DATE(?)", column),
						valueParts[0], valueParts[1])
				case isNumericColumn:
					db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column),
						valueParts[0], valueParts[1])
				default:
					db = db.Where(fmt.Sprintf("%s BETWEEN ? AND ?", column),
						valueParts[0], valueParts[1])
				}
			}

		case "in":
			values := strings.Split(value, "|")
			hasEmptyString := false
			hasNull := false
			filteredValues := []string{}

			for _, v := range values {
				if v == "" {
					hasEmptyString = true
				} else if strings.ToLower(v) == "null" {
					hasNull = true
				} else {
					filteredValues = append(filteredValues, v)
				}
			}

			switch {
			case isDateColumn:
				dates := make([]string, len(filteredValues))
				for i, v := range filteredValues {
					dates[i] = fmt.Sprintf("DATE('%s')", v)
				}
				if hasEmptyString && hasNull {
					db = db.Where(fmt.Sprintf("(DATE(%s) IN (?) OR %s = '' OR %s IS NULL)", column, column, column), filteredValues)
				} else if hasEmptyString {
					db = db.Where(fmt.Sprintf("(DATE(%s) IN (?) OR %s = '')", column, column), filteredValues)
				} else if hasNull {
					db = db.Where(fmt.Sprintf("(DATE(%s) IN (?) OR %s IS NULL)", column, column), filteredValues)
				} else {
					db = db.Where(fmt.Sprintf("DATE(%s) IN (?)", column), filteredValues)
				}

			case isBoolColumn:
				boolValues := make([]bool, len(filteredValues))
				for i, v := range filteredValues {
					boolValues[i] = strings.ToLower(v) == "true"
				}
				if hasEmptyString && hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '' OR %s IS NULL)", column, column, column), boolValues)
				} else if hasEmptyString {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '')", column, column), boolValues)
				} else if hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s IS NULL)", column, column), boolValues)
				} else {
					db = db.Where(fmt.Sprintf("%s IN (?)", column), boolValues)
				}

			case isNumericColumn:
				if hasEmptyString && hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '' OR %s IS NULL)", column, column, column), filteredValues)
				} else if hasEmptyString {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '')", column, column), filteredValues)
				} else if hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s IS NULL)", column, column), filteredValues)
				} else {
					db = db.Where(fmt.Sprintf("%s IN (?)", column), filteredValues)
				}

			default:
				if hasEmptyString && hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '' OR %s IS NULL)", column, column, column), filteredValues)
				} else if hasEmptyString {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s = '')", column, column), filteredValues)
				} else if hasNull {
					db = db.Where(fmt.Sprintf("(%s IN (?) OR %s IS NULL)", column, column), filteredValues)
				} else {
					db = db.Where(fmt.Sprintf("%s IN (?)", column), filteredValues)
				}
			}

		case "isnull":
			if value == "true" {
				db = db.Where(fmt.Sprintf("%s IS NULL", column))
			} else {
				db = db.Where(fmt.Sprintf("%s IS NOT NULL", column))
			}

		case "notin":
			values := strings.Split(value, "|")
			db = db.Where(fmt.Sprintf("%s NOT IN (?)", column), values)

		default:
			switch {
			case isDateColumn:
				db = db.Where(fmt.Sprintf("DATE(%s) = DATE(?)", column), value)
			case isBoolColumn:
				boolValue := strings.ToLower(value) == "true"
				db = db.Where(fmt.Sprintf("%s = ?", column), boolValue)
			case isNumericColumn:
				db = db.Where(fmt.Sprintf("CAST(%s AS TEXT) = ?", column), value)
			default:
				db = db.Where(fmt.Sprintf("%s = ?", column), value)
			}
		}
	}

	// Apply Sorting
	for _, order := range payload.Orders {
		sortOrder := order
		if strings.ToLower(payload.SortType) == "desc" {
			sortOrder += " DESC"
		}
		db = db.Order(sortOrder)
	}

	// Hitung Total Count sebelum Pagination
	var totalCount int64
	if err := db.Model(model).Count(&totalCount).Error; err != nil {
		return nil, 0, 0, err
	}

	if payload.Length <= 0 {
		payload.Length = 10
	}
	if payload.Page <= 0 {
		payload.Page = 1
	}

	totalPages := int(math.Ceil(float64(totalCount) / float64(payload.Length)))

	if payload.Page > totalPages && totalPages > 0 {
		payload.Page = totalPages
	}

	offset := (payload.Page - 1) * payload.Length

	// Preload Relasi Dinamis
	for _, preload := range preloads {
		db = db.Preload(preload)
	}

	if err := db.Offset(offset).Limit(payload.Length).Find(model).Error; err != nil {
		return nil, 0, 0, err
	}

	return model, int(totalCount), totalPages, nil
}

func (sql SQLRepository) GetList(model interface{}, field string, value any, preload ...string) (interface{}, error) {
	database.DB = database.DB.Debug()

	query := database.DB.Where("is_deleted = false")

	if field != "" && value != nil {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	if len(preload) > 0 {
		for _, p := range preload {
			query = query.Preload(p)
		}
	} else {
		query = query.Preload(clause.Associations)
	}

	err := query.Find(model).Error
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (sql SQLRepository) GetOne(model interface{}, field string, value any, preload ...string) (interface{}, error) {
	database.DB = database.DB.Debug()

	query := database.DB.Where("is_deleted = false")

	if field != "" && value != nil {
		query = query.Where(fmt.Sprintf("%s = ?", field), value)
	}

	if len(preload) > 0 {
		for _, p := range preload {
			query = query.Preload(p)
		}
	} else {
		query = query.Preload(clause.Associations)
	}

	err := query.First(model).Error
	if err != nil {
		logger.Errorf("error, data not found: %v", err)
		return nil, err
	}

	return model, nil
}

func (sql SQLRepository) Save(model interface{}, returnModel bool) (interface{}, error) {
	database.DB = database.DB.Debug()
	err := database.DB.Create(model).Error
	if err != nil {
		logger.Errorf("error, not save data %v", err)
		return nil, err
	}
	if returnModel {
		return model, nil
	}
	return nil, nil
}

func (sql SQLRepository) Updates(model interface{}, field string, value any) error {
	db := database.DB.Debug()
	err := db.Model(model).Where(fmt.Sprintf("%s = ?", field), value).Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}

func (sql SQLRepository) Update(model interface{}, field string, value any, data map[string]interface{}) error {
	db := database.DB.Debug()
	err := db.Model(model).Where(fmt.Sprintf("%s = ?", field), value).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func (sql SQLRepository) Delete(model interface{}, field string, value any) error {
	// Ensure that the model is a pointer to a struct
	if err := database.DB.Model(model).Where(fmt.Sprintf("%s = ?", field), value).Updates(model).Error; err != nil {
		return err
	}
	return nil
}

func (sql SQLRepository) HardDelete(model interface{}, field string, value any) error {
	// Pastikan model adalah pointer ke struct
	if err := database.DB.Where(fmt.Sprintf("%s = ?", field), value).Delete(model).Error; err != nil {
		return err
	}
	return nil
}

func (sql SQLRepository) Reject(model interface{}, field string, value any) error {
	var db = database.DB

	if err := db.Model(model).Where(fmt.Sprintf("%s = ?", field), value).
		Update("process_status", "fbeae42e-05bf-48ca-852e-4d77cfbd1f28").Error; err != nil {
		return err
	}
	return nil

}

func (sql SQLRepository) Unreceived(model interface{}, field string, value any) error {
	db := database.DB.Debug()
	err := db.Model(model).Where(fmt.Sprintf("%s = ?", field), value).Select("*").Updates(model).Error
	if err != nil {
		return err
	}
	return nil
}
