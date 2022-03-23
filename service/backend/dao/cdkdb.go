package dao

import (
	"github.com/txchat/dtalk/pkg/util"
	"github.com/txchat/dtalk/service/backend/model/db"
	"math"
	"strings"
	"time"
)

const (
	_InsertCdkType           = `INSERT INTO dtalk_cdk_info ( cdk_id, cdk_name, cdk_info, coin_name, exchange_rate, create_time, update_time, delete_time)VALUES( ?, ?, ?, ?, ?, ?, ?, ?)`
	_GetCdkTypes             = `SELECT * FROM dtalk_cdk_info WHERE delete_time = 0 AND coin_name LIKE ? ORDER BY create_time DESC LIMIT ?,?`
	_DeleteCdkType           = `UPDATE dtalk_cdk_info SET delete_time = ? WHERE cdk_id = ? AND delete_time = 0`
	_InsertCdk               = `INSERT INTO dtalk_cdk_list ( id, cdk_id, cdk_content, user_id, cdk_status, order_id, create_time, update_time, delete_time, exchange_time)VALUES( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`
	_GetCdks                 = `SELECT * FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND cdk_content LIKE ? ORDER BY create_time DESC LIMIT ?,?`
	_DeleteCdk               = `UPDATE dtalk_cdk_list SET delete_time = ? WHERE id = ? AND delete_time = 0`
	_UpdateCdkStatus         = `UPDATE dtalk_cdk_list SET cdk_status = ?, update_time = ? WHERE id = ? AND delete_time = 0`
	_UpdateCdkUserId         = `UPDATE dtalk_cdk_list SET user_id = ?, update_time = ? WHERE id = ? AND delete_time = 0`
	_UpdateCdkOrderId        = `UPDATE dtalk_cdk_list SET order_id = ?, cdk_status = 1, update_time = ? WHERE id = ? AND delete_time = 0`
	_GetCdkTypesCount        = `SELECT COUNT(*) FROM dtalk_cdk_info WHERE delete_time = 0 AND coin_name LIKE ?`
	_GetCdksCount            = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ?`
	_GetCdksWithUserId       = `SELECT * FROM dtalk_cdk_list WHERE delete_time = 0 AND user_id = ? AND cdk_status >= ? ORDER BY create_time DESC LIMIT ?,?`
	_GetCdksCountWithUserId  = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND user_id = ? AND cdk_status >= ?`
	_GetCdkTypesWithCoinName = `SELECT * FROM dtalk_cdk_info WHERE delete_time = 0 AND coin_name = ?`
	_GetUnusedCdksCount      = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND cdk_status = 0`
	_GetFrozenCdksCount      = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND cdk_status = 1`
	_GetUsedCdksCount        = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND cdk_status >= 2`
	_GetCdkType              = `SELECT * FROM dtalk_cdk_info WHERE delete_time = 0 AND cdk_id = ?`
	_DeleteCdks              = `UPDATE dtalk_cdk_list SET delete_time = ? WHERE delete_time = 0 AND id in `
	_DeleteCdkTypes          = `UPDATE dtalk_cdk_info SET delete_time = ? WHERE delete_time = 0 AND cdk_id in `
	_DeleteCdksByCdkIds      = `UPDATE dtalk_cdk_list SET delete_time = ? WHERE delete_time = 0 AND cdk_id in `
	_UpdateCdkType           = `UPDATE dtalk_cdk_info SET cdk_name = ?, coin_name = ?, exchange_rate = ?, update_time = ? WHERE delete_time = 0 AND cdk_id = ?`

	_GetCdksByOrderId = `SELECT * FROM dtalk_cdk_list WHERE delete_time = 0 AND order_id = ?`
	_UpdateCdksStatus = `UPDATE dtalk_cdk_list SET cdk_status = ?, update_time = ?, exchange_time = ? WHERE delete_time = 0 AND id in `
	_FrozenCdksStatus = `UPDATE dtalk_cdk_list SET cdk_status = ?, user_id = ?, order_id = ?, update_time = ? WHERE delete_time = 0 AND id in `
	_GetUnusedCdks    = `SELECT * FROM dtalk_cdk_list WHERE cdk_id = ? AND cdk_status = ? AND delete_time = 0 ORDER BY create_time ASC LIMIT 0,?`

	_CheckCdkExist                = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND cdk_content = ?`
	_GetFrozenCdks                = `SElECT * FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_status = 1`
	_CleanFrozenCdks              = `UPDATE dtalk_cdk_list SET cdk_status = ?, update_time = ? WHERE delete_time = 0 AND order_id = ?`
	_GetCdksCountByUserIdAndCdkId = `SELECT COUNT(*) FROM dtalk_cdk_list WHERE delete_time = 0 AND cdk_id = ? AND user_id = ? AND cdk_status >= 1`
)

func (d *Dao) InsertCdkType(cdkType *db.CdkType) error {
	_, _, err := d.conn.Exec(_InsertCdkType, cdkType.CdkId, cdkType.CdkName, cdkType.CdkInfo, cdkType.CoinName, cdkType.ExchangeRate, cdkType.CreateTime, cdkType.UpdateTime, cdkType.DeleteTime)
	return err
}

func (d *Dao) InsertCdk(cdk *db.Cdk) error {
	_, _, err := d.conn.Exec(_InsertCdk, cdk.Id, cdk.CdkId, cdk.CdkContent, cdk.UserId, cdk.CdkStatus, cdk.OrderId, cdk.CreateTime, cdk.UpdateTime, cdk.DeleteTime, cdk.ExchangeTime)
	return err
}

func (d *Dao) DeleteCdkType(cdkId int64, deleteTime int64) error {
	_, _, err := d.conn.Exec(_DeleteCdkType, deleteTime, cdkId)
	return err
}

func (d *Dao) DeleteCdk(id int64, deleteTime int64) error {
	_, _, err := d.conn.Exec(_DeleteCdk, deleteTime, id)
	return err
}

func (d *Dao) DeleteCdks(ids []int64, deleteTime int64) error {
	idStrs := make([]string, 0, len(ids))
	for _, id := range ids {
		idStrs = append(idStrs, util.ToString(id))
	}
	idStr := "(" + strings.Join(idStrs, ",") + ")"
	execSql := _DeleteCdks + idStr
	_, _, err := d.conn.Exec(execSql, deleteTime)

	return err
}

func (d *Dao) DeleteCdkTypes(cdkIds []int64) error {
	idStrs := make([]string, 0, len(cdkIds))
	for _, id := range cdkIds {
		idStrs = append(idStrs, util.ToString(id))
	}
	idStr := "(" + strings.Join(idStrs, ",") + ")"
	execSql := _DeleteCdkTypes + idStr
	_, _, err := d.conn.Exec(execSql, d.getNowTime())

	return err
}

func (d *Dao) DeleteCdksByCdkIds(cdkIds []int64) error {
	idStrs := make([]string, 0, len(cdkIds))
	for _, id := range cdkIds {
		idStrs = append(idStrs, util.ToString(id))
	}
	idStr := "(" + strings.Join(idStrs, ",") + ")"
	execSql := _DeleteCdksByCdkIds + idStr
	_, _, err := d.conn.Exec(execSql, d.getNowTime())

	return err
}

func (d *Dao) UpdateCdkType(cdkId int64, cdkName, coinName string, exchangeRate int64) error {
	_, _, err := d.conn.Exec(_UpdateCdkType, cdkName, coinName, exchangeRate, d.getNowTime(), cdkId)
	return err
}

func (d *Dao) UpdateCdkStatus(id int64, cdkStatus int32) error {
	_, _, err := d.conn.Exec(_UpdateCdkStatus, cdkStatus, d.getNowTime(), id)
	return err
}

func (d *Dao) UpdateCdkUserId(id int64, userId int64) error {
	_, _, err := d.conn.Exec(_UpdateCdkUserId, userId, d.getNowTime(), id)
	return err
}

func (d *Dao) UpdateCdkOrderId(id int64, orderId int64) error {
	_, _, err := d.conn.Exec(_UpdateCdkOrderId, orderId, d.getNowTime(), id)
	return err
}

func (d *Dao) GetCdkType(cdkId int64) (*db.CdkType, error) {
	var records []map[string]string
	records, err := d.conn.Query(_GetCdkType, cdkId)
	if err != nil {
		return nil, err
	}
	response, err := d.ToCdkTypes(records)
	if err != nil {
		return nil, err
	}

	if len(response) == 0 {
		return nil, nil
	}

	return &response[0], nil
}

func (d *Dao) GetCdkTypes(coinName string, page int64, size int64) ([]db.CdkType, int64, int64, error) {
	var records []map[string]string
	offset := page * size
	if coinName == "" {
		coinName = "%"
	}
	records, err := d.conn.Query(_GetCdkTypes, coinName, offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	response, err := d.ToCdkTypes(records)
	if err != nil {
		return nil, 0, 0, err
	}
	if coinName == "" {
		coinName = "%"
	}
	recordsCount, err := d.conn.Query(_GetCdkTypesCount, coinName)
	if err != nil {
		return nil, 0, 0, err
	}
	totalElements := util.ToInt64(recordsCount[0]["COUNT(*)"])
	totalPages := int64(math.Ceil(float64(totalElements) / float64(size)))
	return response, totalElements, totalPages, nil
}

func (d *Dao) GetCdks(cdkId int64, cdkContent string, page int64, size int64) ([]db.Cdk, int64, int64, error) {
	var records []map[string]string
	offset := page * size
	cdkContent += "%"
	records, err := d.conn.Query(_GetCdks, cdkId, cdkContent, offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	response, err := d.ToCdks(records)
	if err != nil {
		return nil, 0, 0, err
	}
	recordsCount, err := d.conn.Query(_GetCdksCount, cdkId)
	if err != nil {
		return nil, 0, 0, err
	}
	totalElements := util.ToInt64(recordsCount[0]["COUNT(*)"])
	totalPages := int64(math.Ceil(float64(totalElements) / float64(size)))
	return response, totalElements, totalPages, nil
}

func (d *Dao) GetCdksWithUserId(userId string, page int64, size int64) ([]db.Cdk, int64, int64, error) {
	var records []map[string]string
	offset := page * size
	records, err := d.conn.Query(_GetCdksWithUserId, userId, db.CdkFrozen, offset, size)
	if err != nil {
		return nil, 0, 0, err
	}
	response, err := d.ToCdks(records)
	if err != nil {
		return nil, 0, 0, err
	}
	recordsCount, err := d.conn.Query(_GetCdksCountWithUserId, userId, db.CdkFrozen)
	if err != nil {
		return nil, 0, 0, err
	}
	totalElements := util.ToInt64(recordsCount[0]["COUNT(*)"])
	totalPages := int64(math.Ceil(float64(totalElements) / float64(size)))
	return response, totalElements, totalPages, nil
}

func (d *Dao) GetCdkTypesWithCoinName(coinName string) ([]db.CdkType, error) {
	var records []map[string]string
	records, err := d.conn.Query(_GetCdkTypesWithCoinName, coinName)
	if err != nil {
		return nil, err
	}
	response, err := d.ToCdkTypes(records)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *Dao) GetCdksCount(cdkId int64) (int64, int64, int64, error) {
	res, err := d.conn.Query(_GetUnusedCdksCount, cdkId)
	if err != nil {
		return 0, 0, 0, err
	}
	Unused := util.ToInt64(res[0]["COUNT(*)"])

	res, err = d.conn.Query(_GetFrozenCdksCount, cdkId)
	if err != nil {
		return 0, 0, 0, err

	}
	Frozen := util.ToInt64(res[0]["COUNT(*)"])

	res, err = d.conn.Query(_GetUsedCdksCount, cdkId)
	if err != nil {
		return 0, 0, 0, err

	}
	Used := util.ToInt64(res[0]["COUNT(*)"])

	return Unused, Frozen, Used, nil

}

func (d *Dao) GetUnusedCdksCount(cdkId string) (int64, error) {
	recordsCount, err := d.conn.Query(_GetUnusedCdksCount, cdkId)
	if err != nil {
		return 0, err
	}
	response := util.ToInt64(recordsCount[0]["COUNT(*)"])
	return response, nil
}

func (d *Dao) GetFrozenCdksCount(cdkId string) (int64, error) {
	recordsCount, err := d.conn.Query(_GetFrozenCdksCount, cdkId)
	if err != nil {
		return 0, err
	}
	response := util.ToInt64(recordsCount[0]["COUNT(*)"])
	return response, nil
}

func (d *Dao) GetUsedCdksCount(cdkId string) (int64, error) {
	recordsCount, err := d.conn.Query(_GetUsedCdksCount, cdkId)
	if err != nil {
		return 0, err
	}
	response := util.ToInt64(recordsCount[0]["COUNT(*)"])
	return response, nil
}

func (d *Dao) ToCdkTypes(records []map[string]string) ([]db.CdkType, error) {
	cdkTypes := make([]db.CdkType, len(records), len(records))
	for i, record := range records {
		cdkType := db.CdkType{
			CdkId:        util.ToInt64(record["cdk_id"]),
			CdkName:      record["cdk_name"],
			CdkInfo:      record["cdl_info"],
			CoinName:     record["coin_name"],
			ExchangeRate: util.ToInt64(record["exchange_rate"]),
			TimeInfo: db.TimeInfo{
				CreateTime: util.ToInt64(record["create_time"]),
				UpdateTime: util.ToInt64(record["update_time"]),
				DeleteTime: util.ToInt64(record["delete_time"]),
			},
		}

		cdkTypes[i] = cdkType
	}
	return cdkTypes, nil
}

func (d *Dao) ToCdks(records []map[string]string) ([]db.Cdk, error) {
	cdks := make([]db.Cdk, len(records), len(records))
	for i, record := range records {
		cdk := db.Cdk{
			Id:         util.ToInt64(record["id"]),
			CdkId:      util.ToInt64(record["cdk_id"]),
			CdkContent: record["cdk_content"],
			UserId:     record["user_id"],
			CdkStatus:  util.ToInt32(record["cdk_status"]),
			OrderId:    util.ToInt64(record["order_id"]),
			TimeInfo: db.TimeInfo{
				CreateTime: util.ToInt64(record["create_time"]),
				UpdateTime: util.ToInt64(record["update_time"]),
				DeleteTime: util.ToInt64(record["delete_time"]),
			},
			ExchangeTime: util.ToInt64(record["exchange_time"]),
		}

		//fmt.Println(cdk)

		cdks[i] = cdk
	}
	return cdks, nil
}

func (d *Dao) UpdateCdksStatus(ids []int64, status int64) error {
	idStrs := make([]string, 0, len(ids))
	for _, id := range ids {
		idStrs = append(idStrs, util.ToString(id))
	}
	idStr := "(" + strings.Join(idStrs, ",") + ")"

	execSql := _UpdateCdksStatus + idStr
	var exchangeTime int64
	if status == db.CdkUsed || status == db.CdkExchange {
		exchangeTime = d.getNowTime()
	}
	_, _, err := d.conn.Exec(execSql, status, d.getNowTime(), exchangeTime)

	return err
}

func (d *Dao) GetCdksByOrderId(orderId int64) ([]db.Cdk, error) {
	var records []map[string]string
	records, err := d.conn.Query(_GetCdksByOrderId, orderId)
	if err != nil {
		return nil, err
	}
	response, err := d.ToCdks(records)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (d *Dao) GetUnusedCdks(cdkId int64, number int64) ([]db.Cdk, error) {
	var records []map[string]string
	records, err := d.conn.Query(_GetUnusedCdks, cdkId, db.CdkUnused, number)
	if err != nil {
		return nil, err
	}
	response, err := d.ToCdks(records)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *Dao) FrozenCdksStatus(ids []int64, userId string, orderId int64) error {
	idStrs := make([]string, 0, len(ids))
	for _, id := range ids {
		idStrs = append(idStrs, util.ToString(id))
	}
	idStr := "(" + strings.Join(idStrs, ",") + ")"

	execSql := _FrozenCdksStatus + idStr

	_, _, err := d.conn.Exec(execSql, db.CdkFrozen, userId, orderId, d.getNowTime())
	return err
}

func (d *Dao) getNowTime() int64 {
	return time.Now().UnixNano() / 1e6
}

func (d *Dao) CheckCdkExist(cdkId int64, cdkContent string) (bool, error) {
	recordsCount, err := d.conn.Query(_CheckCdkExist, cdkId, cdkContent)
	if err != nil {
		return false, err
	}
	response := util.ToInt64(recordsCount[0]["COUNT(*)"])
	return response >= 1, nil
}

func (d *Dao) GetFrozenCdks() ([]db.Cdk, error) {
	var records []map[string]string
	records, err := d.conn.Query(_GetFrozenCdks)
	if err != nil {
		return nil, err
	}
	response, err := d.ToCdks(records)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (d *Dao) CleanFrozenCdks(orderId int64) error {
	_, _, err := d.conn.Exec(_CleanFrozenCdks, db.CdkUnused, d.getNowTime(), orderId)
	if err != nil {
		return err
	}

	return nil
}

func (d *Dao) GetCdksCountByUserIdAndCdkId(cdkId int64, userId string) (int64, error) {
	recordsCount, err := d.conn.Query(_GetCdksCountByUserIdAndCdkId, cdkId, userId)
	if err != nil {
		return 0, err
	}
	response := util.ToInt64(recordsCount[0]["COUNT(*)"])
	return response, nil
}
