package redis

import (
    "fmt"
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    `lets_diagram/src`
    "log"
    "sync"
    "testing"
)

func init() {
	src.InitSetting("../../../config/setting.json")
	err := dao.InitRedisPool(src.GetSetting().RedisSetting)
	if err != nil {
		log.Fatalf("fail to init redis connection: %v", err)
	}
}

func TestCirculationWriterPermissions(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(30)
	for i := 0; i < 30; i++ {
		ii := i
		go func() {
			err := SetWriterOrReadPermissions(uint(3), uint(ii))
			if err != nil {
				t.Errorf("SetWriterOrReadPermissions() got an error: %v", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	w, err := GetWriterPermission(13)
	if err != nil {
		t.Errorf("GetWriterPermission() got an error: %v", err)
	}
	fmt.Println(w)
	err = CirculationWriterPermissions(uint(13), uint(31))
	if err != nil {
		t.Errorf("CirculationWriterPermissions() got an error: %v", err)
	}
	writer, err := GetWriterPermission(13)
	if err != nil {
		t.Errorf("GetWriterPermission() got an error: %v", err)
	}
	if writer != 31 {
		t.Errorf("GetWriterPermission() want %d, got %d", 31, writer)
	}

}
