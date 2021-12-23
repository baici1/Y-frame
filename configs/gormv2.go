package configs

import (
	"fmt"
	"time"
)

type Gormv2 struct {
	UseDbType string `yaml:"UseDbType" json:"use_db_type,omitempty" mapstructure:"UseDbType"`
	Mysql     Mysql  `yaml:"Mysql" json:"mysql" mapstructure:"Mysql"`
}

type Mysql struct {
	IsInitGlobalGormMysql int           `yaml:"IsInitGlobalGormMysql" json:"is_init_global_gorm_mysql,omitempty" mapstructure:"IsInitGlobalGormMysql"`
	SlowThreshold         time.Duration `yaml:"SlowThreshold" json:"slow_threshold,omitempty" mapstructure:"SlowThreshold"`
	Write                 Write         `yaml:"Write" json:"write" mapstructure:"Write"`
	Log                   Log           `yaml:"Log" json:"log" mapstructure:"Log"`
}

type Write struct {
	Host               string        `yaml:"Host" json:"host,omitempty" mapstructure:"Host"`
	DataBase           string        `yaml:"DataBase" json:"data_base,omitempty" mapstructure:"DataBase"`
	Port               int           `yaml:"Port" json:"port,omitempty" mapstructure:"Port"`
	Prefix             string        `yaml:"Prefix" json:"prefix,omitempty" mapstructure:"Prefix"`
	SetMaxIdleConns    int           `yaml:"SetMaxIdleConns" json:"set_max_idle_conns,omitempty" mapstructure:"SetMaxIdleConns"`
	SetConnMaxLifetime time.Duration `yaml:"SetConnMaxLifetime" json:"set_conn_max_lifetime,omitempty" mapstructure:"SetConnMaxLifetime"`
	User               string        `yaml:"User" json:"user,omitempty" mapstructure:"User"`
	Pass               string        `yaml:"Pass" json:"pass,omitempty" mapstructure:"Pass"`
	Charset            string        `yaml:"Charset" json:"charset,omitempty" mapstructure:"Charset"`
	SetMaxOpenConns    int           `yaml:"SetMaxOpenConns" json:"set_max_open_conns,omitempty" mapstructure:"SetMaxOpenConns"`
}

type Log struct {
	WarnStr      string `yaml:"warnStr" json:"warn_str,omitempty" mapstructure:"WarnStr"`
	ErrStr       string `yaml:"errStr" json:"err_str,omitempty" mapstructure:"ErrStr"`
	TraceStr     string `yaml:"traceStr" json:"trace_str,omitempty" mapstructure:"TraceStr"`
	TraceWarnStr string `yaml:"traceWarnStr" json:"trace_warn_str,omitempty" mapstructure:"TraceWarnStr"`
	TraceErrStr  string `yaml:"traceErrStr" json:"trace_err_str,omitempty" mapstructure:"TraceErrStr"`
	IsOwnLogger  bool   `yaml:"IsOwnLogger" json:"is_own_logger,omitempty" mapstructure:"IsOwnLogger"`
	InfoStr      string `yaml:"infoStr" json:"info_str,omitempty" mapstructure:"InfoStr"`
}

func (m *Mysql) Dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local", m.Write.User, m.Write.Pass, m.Write.Host, m.Write.Port, m.Write.DataBase, m.Write.Charset)
}
