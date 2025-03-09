package db

import (
    "fmt"
    // "database/sql" 不能直接使用，需要使用别名 sqlpkg
    sqlpkg "database/sql"
    "myproject/config"
    // MySQL 驱动包是独立的包，需要单独导入它，所以前面加入了 _
    _ "github.com/go-sql-driver/mysql"
)

func getMySQLConfig() (string, error) {
  config, err := config.LoadConfig("config/config.json")
  if err != nil {
    return "", err
  }
  mysqlConfig := config.MySQL
  dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlConfig.Username, mysqlConfig.Password, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.Database)
  return dsn, nil
}

func QueryDB(sql string, args ...interface{}) ([]map[string]interface{}, error) {
  dsn, err := getMySQLConfig()
  if err != nil {
    return nil, err
  }
  db, err := sqlpkg.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  defer db.Close()

  rows, err := db.Query(sql, args...)
  if err != nil {
    return nil, err
  }
  defer rows.Close()

  columns, err := rows.Columns()
  if err != nil {
    return nil, err
  }

  var results []map[string]interface{}
  for rows.Next() {
    rowMap := make(map[string]interface{})
    columnPointers := make([]interface{}, len(columns))
    for i := range columns {
      columnPointers[i] = new(interface{})
    }

    if err := rows.Scan(columnPointers...); err != nil {
      return nil, err
    }

    for i, colName := range columns {
      val := columnPointers[i].(*interface{})
      rowMap[colName] = *val
    }

    results = append(results, rowMap)
  }

  if err = rows.Err(); err != nil {
    return nil, err
  }

  return results, nil
}
