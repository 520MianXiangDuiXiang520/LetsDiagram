package redis

import (
    "fmt"
    "github.com/520MianXiangDuiXiang520/GoTools/dao"
    "github.com/garyburd/redigo/redis"
    `lets_diagram/src/core/access`
    "log"
    "strconv"
)

func getUCPKey(canvasID uint) string {
	return UserCanvasPermissionsKey + strconv.Itoa(int(canvasID))
}

// SetWriterPermissions 用来将 user 对 canvas 的权限设置为可写
// 如果该 canvas 已经设置了 writer, 则 user 对 canvas 被降为 readOnly
func SetWriterOrReadPermissions(canvasID, userID uint) error {
	var err error
	r := dao.GetRedisConn()
	defer func() {
		if err != nil {
			_, err := r.Do("DISCARD")
			if err != nil {
				log.Printf("DISCARD error")
			}
		}
		r.Close()
	}()
	_, err = r.Do("MULTI")
	if err != nil {
		log.Println("MULTI error")
	}
	writer, err := getWriterPermission(r, canvasID)
	if err != nil {
		return err
	}
	if writer == 0 {
		err := setPermissions(r, canvasID, userID, access.CanvasJurisdictionMarkWrite)
		if err != nil {
			return err
		}
		return nil
	}
	err = setPermissions(r, canvasID, userID, access.CanvasJurisdictionMarkReadOnly)
	if err != nil {
		return err
	}
	reply, err := r.Do("EXEC")
	if err != nil {
		log.Printf("EXEC error")
	}
	fmt.Println(reply)
	return nil
}

func setPermissions(redis redis.Conn, canvasID, userID uint, per access.CanvasJurisdictionMark) error {
	var err error
	_, err = redis.Do("HSET", getUCPKey(canvasID), userID, per)
	if err != nil {
		return err
	}
	if per == access.CanvasJurisdictionMarkWrite {
		_, err = redis.Do("HSET", getUCPKey(canvasID), UserCanvasPermissionsKeyWriter, userID)
		if err != nil {
			return err
		}
	}
	return nil
}

// SetPermissions 将某个用户对 canvas 的操作权限设置为 per
func SetPermissions(canvasID, userID uint, per access.CanvasJurisdictionMark) error {
	var err error
	r := dao.GetRedisConn()
	defer func() {
		if err != nil {
			_, _ = r.Do("DISCARD")
		} else {
			_, _ = r.Do("EXEC")
		}
		r.Close()
	}()
	// 开启事务
	_, _ = r.Do("MULTI")

	return setPermissions(r, canvasID, userID, per)
}

func getPermissions(conn redis.Conn, canvasID, userID uint) (access.CanvasJurisdictionMark, error) {
	reply, err := redis.String(conn.Do("HGET", getUCPKey(canvasID), userID))
	if err != nil {
		return 0, err
	}
	mark, err := strconv.Atoi(reply)
	if err != nil {
		return 0, fmt.Errorf("GetPermissions want get a type of int but got %v \n", reply)
	}
	switch mark {
	case 1:
		return access.CanvasJurisdictionMarkReadOnly, nil
	case 2:
		return access.CanvasJurisdictionMarkWrite, nil
	default:
		return 0, fmt.Errorf("mark(%d) not a CanvasJurisdictionMark", mark)
	}
}

// GetPermissions 查询 redis, 获得 user 在 canvas 中的权限
func GetPermissions(canvasID, userID uint) (access.CanvasJurisdictionMark, error) {
	conn := dao.GetRedisConn()
	defer conn.Close()
	return getPermissions(conn, canvasID, userID)
}

// getWriterPermission 通过查询 redis 返回 canvas 的写权限由谁持有
func getWriterPermission(conn redis.Conn, canvasID uint) (uint, error) {
	reply, err := redis.String(conn.Do("HGET", getUCPKey(canvasID), UserCanvasPermissionsKeyWriter))
	if err != nil {
		return 0, err
	}
	if reply == "QUEUED" {
		return 0, nil
	}
	u, err := strconv.Atoi(reply)
	if err != nil {
		return 0, fmt.Errorf("getWriterPermission want get a type of int, but got %v", reply)
	}
	return uint(u), nil
}

func GetWriterPermission(canvasID uint) (uint, error) {
	conn := dao.GetRedisConn()
	defer conn.Close()
	return getWriterPermission(conn, canvasID)
}

// 将写权限移交给 toUser
func CirculationWriterPermissions(canvasID, UserID uint) error {
	conn := dao.GetRedisConn()
	var err error

	defer func() {
		if err != nil {
			_, _ = conn.Do("DISCARD")
		} else {
			_, _ = conn.Do("EXEC")
		}
		conn.Close()

	}()
	// 开启事务
	_, _ = conn.Do("MULTI")
	writer, err := getWriterPermission(conn, canvasID)
	if err != nil {
		return err
	}
	err = SetPermissions(canvasID, writer, access.CanvasJurisdictionMarkReadOnly)
	if err != nil {
		return err
	}
	err = SetPermissions(canvasID, UserID, access.CanvasJurisdictionMarkWrite)
	if err != nil {
		return err
	}
	return nil
}
