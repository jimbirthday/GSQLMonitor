package comm

var (
	//连接数
	DB_PROCESS_LIST = "show processlist"
	//查询数
	DB_QPS_QUESTIONS = "SHOW GLOBAL STATUS LIKE 'Questions'"
	//启动时间
	DB_UPTIME = "SHOW GLOBAL STATUS LIKE 'Uptime'"
	//提交数
	DB_TPS_COM_COMMIT = " SHOW GLOBAL STATUS LIKE 'Com_commit'"
	//回滚数
	DB_TPS_COM_ROLLBACK = " SHOW GLOBAL STATUS LIKE 'Com_rollback'"
	//key buffer 状态
	DB_KEY_BUFFER = " show status like 'Key%'"
	//innodb 缓冲池状态
	DB_INNODB_BUFFER = " show status like 'innodb_buffer_pool_read%';"
	//Query Cache命中率
	DB_QUERY_CACHE = " show status like 'Qcache%';"
	//Thread Cache 命中率
	DB_THREAD_CACHE = " show status like 'Thread%';"
	DB_CONNECTIONS  = " show status like 'Connections';"
	//锁状态
	DB_LOCK = " show status like '%lock%';"
	//慢sql状态
	DB_SLOW_SQL_INFO = " show variables like 'slow_query%'"
	//慢sql时长
	DB_SLOW_SQL_TIME = " show variables like '%long_query_time%'"
	//慢sql条数
	DB_SLOW_SQL_COUNTS = " SHOW GLOBAL STATUS LIKE '%Slow_queries%';"
	//查看默认引擎
	DB_ENGINE = "show variables like 'default_storage_engine';"

	DB_DATABASE_SIZE = " select concat(round(sum(DATA_LENGTH/1024/1024),2),'MB') as data from TABLES;"

	DB_CHECK= " select 1=1;"

)
