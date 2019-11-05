package model


// 插入记录
func InsertData(bean interface{}) (int64, error) {
	return db.GetEngine().InsertOne(bean)
}

// 删除记录
func DeleteData(bean interface{}, where string, args ...interface{}) (int64, error) {
	return db.GetEngine().Where(where, args...).Delete(bean)
}
