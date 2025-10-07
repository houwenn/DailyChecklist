package models

import (
	"DailyChecklist/dao"
	"time"
)

// Todo Model
/*
CREATE TABLE `todo` (
    `id` BIGINT(20) NOT NULL AUTO_INCREMENT,
    `title` VARCHAR(255) NOT NULL,
    `status` TINYINT(1) DEFAULT 0,
    `created_at` DATETIME NOT NULL,
    `updated_at` DATETIME NOT NULL,
    `is_archived` TINYINT(1) DEFAULT 0,
    `archived_at` DATETIME,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
*/
type Todo struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	Status     bool       `json:"status"`
	IsArchived bool       `json:"is_archived"` // 是否已存档
	CreatedAt  time.Time  `json:"created_at"`  // 创建时间
	UpdatedAt  time.Time  `json:"updated_at"`  // 更新时间
	ArchivedAt *time.Time `json:"archived_at"` // 存档时间
}

/*
	Todo这个Model的增删改查操作都放在这里
*/
// CreateATodo 创建todo
func CreateATodo(todo *Todo) (err error) {
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()
	err = dao.DB.Create(&todo).Error
	return
}

func GetATodo(id string) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Debug().Where("id=?", id).First(todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(todo).Error
	return
}

func DeleteATodo(id string) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}

// ArchiveATodo 存档todo
func ArchiveATodo(id string) (err error) {
	now := time.Now()
	err = dao.DB.Model(&Todo{}).Where("id=?", id).Updates(map[string]interface{}{
		"is_archived": true,
		"archived_at": now,
	}).Error
	return
}

// GetArchivedTodos 获取所有存档的todo
func GetArchivedTodos() (todoList []*Todo, err error) {
	if err = dao.DB.Where("is_archived = ?", true).Order("archived_at DESC").Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetActiveTodos 获取所有未存档的todo
func GetActiveTodos() (todoList []*Todo, err error) {
	if err = dao.DB.Where("is_archived = ?", false).Find(&todoList).Error; err != nil {
		return nil, err
	}
	return
}

// GetTodoStats 获取统计数据
func GetTodoStats() (stats map[string]interface{}, err error) {
	var totalCount, completedCount, archivedCount int64

	// 总数量
	dao.DB.Model(&Todo{}).Where("is_archived = ?", false).Count(&totalCount)

	// 已完成数量
	dao.DB.Model(&Todo{}).Where("is_archived = ? AND status = ?", false, true).Count(&completedCount)

	// 已存档数量
	dao.DB.Model(&Todo{}).Where("is_archived = ?", true).Count(&archivedCount)

	completionRate := 0.0
	if totalCount > 0 {
		completionRate = float64(completedCount) / float64(totalCount) * 100
	}

	stats = map[string]interface{}{
		"total_count":     totalCount,
		"completed_count": completedCount,
		"archived_count":  archivedCount,
		"completion_rate": completionRate,
	}

	return stats, nil
}

// GetArchivedTodoStats 获取存档统计数据
func GetArchivedTodoStats() (stats map[string]interface{}, err error) {
	var totalCount, completedCount int64

	// 总存档数量
	dao.DB.Model(&Todo{}).Where("is_archived = ?", true).Count(&totalCount)

	// 存档中已完成的数量
	dao.DB.Model(&Todo{}).Where("is_archived = ? AND status = ?", true, true).Count(&completedCount)

	completionRate := 0.0
	if totalCount > 0 {
		completionRate = float64(completedCount) / float64(totalCount) * 100
	}

	stats = map[string]interface{}{
		"total_count":     totalCount,
		"completed_count": completedCount,
		"completion_rate": completionRate,
		"archived_count":  totalCount, // 保持兼容性
	}

	return stats, nil
}

// GetTodoHistoryStats 获取历史统计数据
func GetTodoHistoryStats(period string) (stats map[string]interface{}, err error) {
	var labels []string
	var data []map[string]interface{}

	today := time.Now()

	// 根据时间段生成标签和数据
	switch period {
	case "week":
		// 最近7天
		for i := 6; i >= 0; i-- {
			date := today.AddDate(0, 0, -i)
			startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
			endOfDay := startOfDay.AddDate(0, 0, 1)

			var count int64
			dao.DB.Model(&Todo{}).Where("created_at >= ? AND created_at < ?", startOfDay, endOfDay).Count(&count)

			labels = append(labels, date.Format("01/02"))
			data = append(data, map[string]interface{}{
				"date":  date.Format("2006-01-02"),
				"count": count,
			})
		}
	case "month":
		// 最近30天
		for i := 29; i >= 0; i-- {
			date := today.AddDate(0, 0, -i)
			startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
			endOfDay := startOfDay.AddDate(0, 0, 1)

			var count int64
			dao.DB.Model(&Todo{}).Where("created_at >= ? AND created_at < ?", startOfDay, endOfDay).Count(&count)

			labels = append(labels, date.Format("01/02"))
			data = append(data, map[string]interface{}{
				"date":  date.Format("2006-01-02"),
				"count": count,
			})
		}
	case "year":
		// 最近12个月
		for i := 11; i >= 0; i-- {
			date := today.AddDate(0, -i, 0)
			startOfMonth := time.Date(date.Year(), date.Month(), 1, 0, 0, 0, 0, date.Location())
			endOfMonth := startOfMonth.AddDate(0, 1, 0)

			var count int64
			dao.DB.Model(&Todo{}).Where("created_at >= ? AND created_at < ?", startOfMonth, endOfMonth).Count(&count)

			labels = append(labels, date.Format("2006/01"))
			data = append(data, map[string]interface{}{
				"date":  date.Format("2006-01"),
				"count": count,
			})
		}
	}

	stats = map[string]interface{}{
		"labels": labels,
		"data":   data,
		"period": period,
	}

	return stats, nil
}
