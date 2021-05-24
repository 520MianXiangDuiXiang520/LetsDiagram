package src

import (
	"github.com/520MianXiangDuiXiang520/GoTools/json"
	path2 "github.com/520MianXiangDuiXiang520/GoTools/path"
	"path"
	"runtime"
	"sync"
	"time"
)

type Setting struct {
	MySQLSetting   MySQLConn      `json:"mysql_setting"`
	TokenSetting   TokenSetting   `json:"token_setting"`
	CORSSetting    CORSSetting    `json:"cors_setting"`
	MongoDBConn    MongoDBConn    `json:"mongodb_conn"`
	DefaultSetting DefaultSetting `json:"default_setting"`
	RedisSetting   RedisSetting   `json:"redis_setting"`
}

type RedisSetting struct {
	Host      string        `json:"host"`
	Password  string        `json:"password"`
	Port      int           `json:"port"`
	MIdleConn int           `json:"max_idle_conn"` // 最大空闲连接数
	MOpenConn int           `json:"max_open_conn"` // 最大打开连接数
	MLifetime time.Duration `json:"max_lifetime"`  // 连接超时时间
}

type DefaultSetting struct {
	DefaultPage      int `json:"default_page"`
	DefaultPageSize  int `json:"default_page_size"`
	DefaultMaxPCU    int `json:"default_max_pcu"`    // 同一 Canvas 最多允许的同时在线协作人数
	CooperateCodeLen int `json:"cooperate_code_len"` // 协作码的长度
}

type MySQLConn struct {
	Engine    string        `json:"engine"`
	DBName    string        `json:"db_name"`
	User      string        `json:"user"`
	Password  string        `json:"password"`
	Host      string        `json:"host"`
	Port      int           `json:"port"`
	MIdleConn int           `json:"max_idle_conn"` // 最大空闲连接数
	MOpenConn int           `json:"max_open_conn"` // 最大打开连接数
	MLifetime time.Duration `json:"max_lifetime"`  // 连接超时时间
	LogMode   bool          `json:"log_mode"`
}

type MongoDBConn struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Database string `json:"database"`
}

type CORSSetting struct {
	AllowedList []string `json:"allowed_list"`
}

type TokenSetting struct {
	ExpireTime int64 `json:"expire_time"`
}

var setting *Setting
var settingLock sync.Mutex

func InitSetting(filePath string) {
	defer func() {
		if e := recover(); e != nil {
			settingLock.Unlock()
		}
	}()
	filename := filePath
	if !path2.IsAbs(filePath) {
		_, currently, _, _ := runtime.Caller(1)
		filename = path.Join(path.Dir(currently), filePath)
	}
	if setting == nil {
		settingLock.Lock()
		if setting == nil {
			err := json.FromFileLoadToObj(&setting, filename)
			if err != nil {
				panic("read setting error!")
			}
		}
		settingLock.Unlock()
	}
}

func GetSetting() *Setting {
	if setting == nil {
		panic("setting Uninitialized！")
	}
	return setting
}
