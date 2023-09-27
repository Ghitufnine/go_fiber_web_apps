package todos

import "go_fiber_web_apps/app/database"

func GetTodoQuery() []TodoModel {
	result := []TodoModel{}

	database.DBConn.Unscoped().
		Table("todo").
		Find(&result)

	return result
}

func TransactionUpdateTodoQuery(data TodoModel, id string) error {
	// Data Transaction
	tx := database.DBConn.Begin()

	if err := tx.
		Table("todo").
		Where("id = ?", id).
		Updates(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func TransactionInsertTodoQuery(data TodoModel) error {
	// Data Transaction
	tx := database.DBConn.Begin()

	if err := tx.
		Table("todo").
		Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func TransactionDeleteTodoQuery(id string) error {
	// Data Transaction
	tx := database.DBConn.Begin()

	if err := tx.
		Table("todo").
		Where("id = ?", id).
		Delete(&TodoModel{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func CheckExistTodoQuery(id string) int {
	count := 0

	database.DBConn.Unscoped().
		Table("todo").
		Where("id = ?", id).
		Count(&count)

	return count
}
