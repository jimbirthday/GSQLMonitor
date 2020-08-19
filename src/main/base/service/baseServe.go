package service

import (
	"math"
	"middleProject/src/main/base/domain"
	"middleProject/src/main/comm"
	"strconv"
)

func DataBaseStatus() *domain.DBInfo {
	dbInfo := new(domain.DBInfo)
	dbInfo.DataBaseName = comm.DBDatabase
	dbInfo.Port = comm.PORT

	_, err := comm.Db.Query(comm.DB_CHECK)
	if err == nil {
		dbInfo.Status = "Running"
	} else {
		dbInfo.Status = "Stop"
	}

	upResults, err := comm.Db.Query(comm.DB_UPTIME)
	if err == nil {
		for _, res := range upResults {
			if v, ok := res["Value"]; ok {
				dbInfo.Uptime = string(v)
			}
		}
	}

	engineResults, err := comm.DbInfo.Query(comm.DB_ENGINE)
	if err == nil {
		for _, res := range engineResults {
			if v, ok := res["Value"]; ok {
				dbInfo.Engine = string(v)
			}
		}

	}

	sizeResults, _ := comm.DbInfo.Query(comm.DB_DATABASE_SIZE)
	if err == nil {
		for _, res := range sizeResults {
			if v, ok := res["data"]; ok {
				dbInfo.DataSize = string(v)
			}
		}
	}

	return dbInfo
}

//查看MySQl连接数
func Processlist() (list []*domain.ProcesslistEntity, err error) {
	results, err := comm.Db.Query(comm.DB_PROCESS_LIST)
	if err != nil {
		return nil, err
	}
	entities := make([]*domain.ProcesslistEntity, len(results))
	for k, res := range results {
		d := new(domain.ProcesslistEntity)
		if v, ok := res["Id"]; ok {
			d.Id = string(v)
		}
		if v, ok := res["User"]; ok {
			d.User = string(v)
		}
		if v, ok := res["Host"]; ok {
			d.Host = string(v)
		}
		if v, ok := res["db"]; ok {
			d.BaseName = string(v)
		}
		if v, ok := res["Command"]; ok {
			d.Command = string(v)
		}
		if v, ok := res["Time"]; ok {
			d.Time = string(v)
		}
		if v, ok := res["State"]; ok {
			d.State = string(v)
		}
		if v, ok := res["Info"]; ok {
			d.Info = string(v)
		}
		entities[k] = d
	}

	return entities, nil
}

//查询数据库QPS【每秒处理的查询数】
func QPS() (list []*domain.QuestionsEntity, err error) {
	qpsResults, err := comm.Db.Query(comm.DB_QPS_QUESTIONS)
	if err != nil {
		return nil, err
	}
	qpsList := make([]*domain.QuestionsEntity, len(qpsResults))
	for k, res := range qpsResults {
		d := new(domain.QuestionsEntity)
		if v, ok := res["Variable_name"]; ok {
			d.VariableName = string(v)
		}
		if v, ok := res["Value"]; ok {
			parseInt, _ := strconv.ParseFloat(string(v), 64)
			d.Value = parseInt
		}
		qpsList[k] = d
	}

	upResults, err := comm.Db.Query(comm.DB_UPTIME)
	if err != nil {
		return nil, err
	}
	upList := make([]*domain.QuestionsEntity, len(upResults))
	for k, res := range upResults {
		d := new(domain.QuestionsEntity)
		if v, ok := res["Variable_name"]; ok {
			d.VariableName = string(v)
		}
		if v, ok := res["Value"]; ok {
			parseInt, _ := strconv.ParseFloat(string(v), 64)
			d.Value = parseInt
		}
		upList[k] = d
	}

	QPSList := make([]*domain.QuestionsEntity, len(qpsList))
	for k, res := range qpsList {
		d := new(domain.QuestionsEntity)
		d.VariableName = res.VariableName
		upEntity := upList[k]
		d.Value = res.Value / upEntity.Value
		QPSList[k] = d
	}

	return QPSList, nil
}

//查看数据库TPS 【每秒处理的事务数】
func TPS() (list []*domain.QuestionsEntity, err error) {
	commitResults, err := comm.Db.Query(comm.DB_TPS_COM_COMMIT)
	if err != nil {
		return nil, err
	}
	commitList := make([]*domain.QuestionsEntity, len(commitResults))
	for k, res := range commitResults {
		d := new(domain.QuestionsEntity)
		if v, ok := res["Variable_name"]; ok {
			d.VariableName = string(v)
		}
		if v, ok := res["Value"]; ok {
			parseInt, _ := strconv.ParseFloat(string(v), 64)
			d.Value = parseInt
		}
		commitList[k] = d
	}

	rollbackResults, err := comm.Db.Query(comm.DB_TPS_COM_ROLLBACK)
	if err != nil {
		return nil, err
	}
	rollbackList := make([]*domain.QuestionsEntity, len(rollbackResults))
	for k, res := range rollbackResults {
		d := new(domain.QuestionsEntity)
		if v, ok := res["Variable_name"]; ok {
			d.VariableName = string(v)
		}
		if v, ok := res["Value"]; ok {
			parseInt, _ := strconv.ParseFloat(string(v), 64)
			d.Value = parseInt
		}
		rollbackList[k] = d
	}

	upResults, err := comm.Db.Query(comm.DB_UPTIME)
	if err != nil {
		return nil, err
	}
	upList := make([]*domain.QuestionsEntity, len(upResults))
	for k, res := range upResults {
		d := new(domain.QuestionsEntity)
		if v, ok := res["Variable_name"]; ok {
			d.VariableName = string(v)
		}
		if v, ok := res["Value"]; ok {
			parseInt, _ := strconv.ParseFloat(string(v), 64)
			d.Value = parseInt
		}
		upList[k] = d
	}

	TPSList := make([]*domain.QuestionsEntity, len(rollbackList))
	for k, res := range upList {
		d := new(domain.QuestionsEntity)
		d.VariableName = res.VariableName
		commEntity := commitList[k]
		rollEntity := rollbackList[k]
		d.Value = (commEntity.Value + rollEntity.Value) / res.Value
		TPSList[k] = d
	}
	return TPSList, nil
}

//key_buffer_read_hits = (1-key_reads / key_read_requests) * 100%
//key_buffer_write_hits = (1-key_writes / key_write_requests) * 100%
//key Buffer 命中率
func BufferCache() (list []*domain.KeyBufferEntity, err error) {
	commitResults, err := comm.Db.Query(comm.DB_KEY_BUFFER)
	if err != nil {
		return nil, err
	}
	commitList := make([]*domain.KeyBufferEntity, 2)
	var keyReads = 0.0
	var keyReadRequests = 0.0
	var keyWrites = 0.0
	var keyWriteRequests = 0.0
	for _, res := range commitResults {
		if v, ok := res["Variable_name"]; ok {
			if string(v) == "Key_reads" {
				keyReads, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
			if string(v) == "Key_read_requests" {
				keyReadRequests, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
			if string(v) == "Key_writes" {
				keyWrites, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
			if string(v) == "Key_write_requests" {
				keyWriteRequests, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
		}

	}
	keyReadEntity := new(domain.KeyBufferEntity)
	keyReadEntity.Name = "key_buffer_read_hits"
	r := keyReads / keyReadRequests
	if n := math.IsNaN(r); n {
		keyReadEntity.Value = 0
	} else {
		if r == 1 {
			keyReadEntity.Value = 100
		} else {
			keyReadEntity.Value = (1 - r) * 100
		}
	}

	keyWriteEntity := new(domain.KeyBufferEntity)
	keyWriteEntity.Name = "key_buffer_write_hits"
	w := keyWrites / keyWriteRequests
	if n := math.IsNaN(w); n {
		keyWriteEntity.Value = 0
	} else {
		if w == 1 {
			keyWriteEntity.Value = 100
		} else {
			keyWriteEntity.Value = (1 - w) * 100
		}
	}
	commitList[0] = keyReadEntity
	commitList[1] = keyWriteEntity

	return commitList, nil

}

//InnoDB Buffer命中率
//innodb_buffer_read_hits = (1 - innodb_buffer_pool_reads / innodb_buffer_pool_read_requests) * 100%
func InnoDBBufferCache() (list *domain.KeyBufferEntity, err error) {
	commitResults, err := comm.Db.Query(comm.DB_INNODB_BUFFER)
	if err != nil {
		return nil, err
	}
	var reads = 0.0
	var requests = 0.0
	for _, res := range commitResults {
		if v, ok := res["Variable_name"]; ok {
			if string(v) == "Innodb_buffer_pool_reads" {
				reads, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
			if string(v) == "Innodb_buffer_pool_read_requests" {
				requests, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
		}

	}
	keyReadEntity := new(domain.KeyBufferEntity)
	keyReadEntity.Name = "innodb_buffer_read_hits"
	r := reads / requests
	if n := math.IsNaN(r); n {
		keyReadEntity.Value = 0
	} else {
		keyReadEntity.Value = (1 - r) * 100

	}
	return keyReadEntity, nil

}

//Query Cache命中率
//Query_cache_hits = (Qcahce_hits / (Qcache_hits + Qcache_inserts )) * 100%;
func QueryCache() (list *domain.KeyBufferEntity, err error) {
	commitResults, err := comm.Db.Query(comm.DB_QUERY_CACHE)
	if err != nil {
		return nil, err
	}
	var qcacheHits = 0.0
	var QcacheInsert = 0.0
	for _, res := range commitResults {
		if v, ok := res["Variable_name"]; ok {
			if string(v) == "Qcache_hits" {
				qcacheHits, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
			if string(v) == "Qcache_inserts" {
				QcacheInsert, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
		}
	}
	keyReadEntity := new(domain.KeyBufferEntity)
	keyReadEntity.Name = "Query_cache_hits"
	r := qcacheHits / (qcacheHits + QcacheInsert)
	if n := math.IsNaN(r); n {
		keyReadEntity.Value = 0
	} else {
		keyReadEntity.Value = (1 - r) * 100

	}
	return keyReadEntity, nil

}

//Thread Cache 命中率
//Thread_cache_hits = (1 - Threads_created / connections ) * 100%
func ThreadCache() (list *domain.KeyBufferEntity, err error) {
	commitResults, err := comm.Db.Query(comm.DB_THREAD_CACHE)
	if err != nil {
		return nil, err
	}
	var thread = 0.0
	for _, res := range commitResults {
		if v, ok := res["Variable_name"]; ok {
			if string(v) == "Threads_created" {
				thread, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
		}
	}

	connResults, err := comm.Db.Query(comm.DB_CONNECTIONS)
	if err != nil {
		return nil, err
	}
	var conn = 0.0
	for _, res := range connResults {
		if v, ok := res["Variable_name"]; ok {
			if string(v) == "Connections" {
				conn, _ = strconv.ParseFloat(string(res["Value"]), 64)
			}
		}
	}

	keyReadEntity := new(domain.KeyBufferEntity)
	keyReadEntity.Name = "Query_cache_hits"
	r := thread / conn
	if n := math.IsNaN(r); n {
		keyReadEntity.Value = 0
	} else {
		keyReadEntity.Value = (1 - r) * 100

	}
	return keyReadEntity, nil

}

//锁定状态
func LockStatus() (list []*domain.KeyBufferEntity, err error) {
	results, err := comm.Db.Query(comm.DB_LOCK)
	if err != nil {
		return nil, err
	}
	entities := make([]*domain.KeyBufferEntity, len(results))
	for k, res := range results {
		d := new(domain.KeyBufferEntity)
		if v, ok := res["Variable_name"]; ok {
			d.Name = string(v)
		}
		if v, ok := res["Value"]; ok {
			conn, _ := strconv.ParseFloat(string(v), 64)
			d.Value = conn
		}
		entities[k] = d
	}
	return entities, nil

}

func SlowSQl(sql string) (list []map[string]string, err error) {
	commitResults, err := comm.Db.QueryString(" explain " + sql)
	if err != nil {
		return nil, err
	}
	return commitResults, nil

}
