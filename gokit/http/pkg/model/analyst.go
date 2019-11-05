package model

import (
)

/*
收藏
 */
// 查询用户是否存在
func GetCollectByYrdUidStockId(yrdUid string, stockId int) (*AnalystCollect, bool, error) {
	var collect AnalystCollect
	if result, err := db.GetEngine().Where("yrd_uid=? and stock_id=?", yrdUid, stockId).Get(&collect); err != nil {
		return nil, result, err
	} else {
		return &collect, result, nil
	}
}

// 更新数据
func UpdateCollectByYrdUidStockId(yrdUid string, stockId int, bean interface{}) (int64, error) {
	return db.GetEngine().Where("yrd_uid=? and stock_id=?", yrdUid, stockId).Cols(
		`collection`).Update(bean)
}

// 查询收藏股票信息总数
func GetCollectCount(tabType, yrdUid string) (int64, error) {
	s := Count{}
	if _, err := db.GetEngine().SQL(`select count(1) as s from analyst_stocksdata as a,analyst_collect as b where b.yrd_uid = ? and b.stock_id = a.id and a.srating_name=? and b.collection=?`,
		yrdUid, tabType, config.CollectYes,
	).Get(&s); err != nil {
		return 0, err
	} else {
		return s.S, err
	}
}

// 查询收藏股票信息
func GetCollectStockByDateYrdUid(tabType, yrdUid string, page, pageSize int64) ([]Collect, error) {
	stocks := make([]Collect, 0)
	err := db.GetEngine().SQL(`select * from analyst_stocksdata as a,analyst_collect as b where b.yrd_uid = ? and b.stock_id = a.id and a.srating_name=?  and b.collection=? order by b.updated_at desc limit ?,?`,
		yrdUid, tabType, config.CollectYes, int((page-1)*pageSize), int(pageSize),
	).Find(&stocks)
	return stocks, err
}

/*
用户
 */
// 查询用户是否存在
func GetUserByYrdUid(yrdUid string) (*AnalystUser, bool, error) {
	var user AnalystUser
	if result, err := db.GetEngine().Where("yrd_uid=?", yrdUid).Get(&user); err != nil {
		return nil, result, err
	} else {
		return &user, result, nil
	}
}

// 查询股票by id
func GetUserById(id int) (*AnalystStocksdata, bool, error) {
	var stock AnalystStocksdata
	if result, err := db.GetEngine().Id(id).Get(&stock); err != nil {
		return nil, result, err
	} else {
		return &stock, result, nil
	}
}

/*
股票数据
 */
// 查询股票信息
func GetStockByDate(date string, tabType string) ([]AnalystStocksdata, error) {
	stocks := make([]AnalystStocksdata, 0)
	err := db.GetEngine().Where("date=? and srating_name=?", date, tabType).Find(&stocks)
	return stocks, err
}

// 根据日期和code查询股票信息
func GetStockByDateCodeSratingNameStockType(date, code, tabType string) (*AnalystStocksdata, bool, error) {
	var stock AnalystStocksdata
	if result, err := db.GetEngine().Where("date=? and code=? and srating_name=?", date, code, tabType).Get(&stock); err != nil {
		return nil, result, err
	} else {
		return &stock, result, nil
	}
}
// 根据日期和industryCode查询股票信息
func GetStockByDateIndustryCodeSratingNameStockType(date, industryCode, tabType string) (*AnalystStocksdata, bool, error) {
	var stock AnalystStocksdata
	if result, err := db.GetEngine().Where("date=? and industry_code=? and srating_name=?", date, industryCode, tabType).Get(&stock); err != nil {
		return nil, result, err
	} else {
		return &stock, result, nil
	}
}

// 查询股票信息
func GetStockByDateYrdUid(date, yrdUid, tabType string, page, pageSize int64) ([]Collect, error) {
	stocks := make([]Collect, 0)
	err := db.GetEngine().SQL(`select * from analyst_stocksdata as a left join analyst_collect as b on b.yrd_uid = ? and b.stock_id = a.id where date = ? and srating_name=? order by potential_gain desc limit ?,?`,
		yrdUid, date, tabType, int((page-1)*pageSize), int(pageSize),
	).Find(&stocks)
	return stocks, err
}
