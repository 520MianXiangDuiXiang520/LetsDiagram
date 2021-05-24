package versionlink

import (
	"encoding/json"
	json_diff "github.com/520MianXiangDuiXiang520/json-diff"
	"lets_diagram/src/dao"
	"log"
)

// 用于版本链的持久化

func computeNumberOfEndurance(size int) int {
	switch {
	case size < 100:
		return int(float64(size)*0.1) + 1
	case size < 500:
		return int(float64(size)*0.2) + 1
	case size < 1000:
		return int(float64(size)*0.3) + 1
	case size < 2000:
		return int(float64(size)*0.4) + 1
	default:
		return int(float64(size)*0.5) + 1
	}
}

func NewEndurance(nodes []*DiffNode, count int, canvasID uint) bool {
	canvas, err := dao.GetCanvasData(canvasID)
	if err != nil {
		// 如果持久化未完成，但用户已经删除了 canvas 可能导致 err != nil
		// 使用弱删除策略，仍然持久化
		log.Printf("%+v", err)
		return false
	}
	// 已经持久化了的版本，小于该版本的内存节点会被丢弃
	oldVersion := canvas.Version
	oldNode, err := json_diff.Unmarshal([]byte(canvas.Data))
	if err != nil {
		log.Printf("fail to endurance: %+v", err)
		return false
	}
	var newNode *json_diff.JsonNode
	newVersion := int64(0)
	for i := 0; i < count; i++ {
		diffValue := nodes[i].Value
		newVersion = nodes[i].Version
		if newVersion != oldVersion+1 {
			log.Printf("[endurance] latest %d version discard %d version", oldVersion, newVersion)
		}
		diffNode, err := json.Marshal(diffValue)
		if err != nil {
			log.Printf("fail to marshal diffNode")
			return false
		}
		diff, err := json_diff.Unmarshal(diffNode)
		if err != nil {
			log.Printf("fail to unmarshal diffNode")
			return false
		}

		newNode, err = json_diff.MergeDiffNode(oldNode, diff)
		if err != nil {
			log.Printf("can not merge diff: %v\n", err)
			return false
		}
		oldNode = newNode
		oldVersion = newVersion
	}
	err = dao.UpdateCanvasData(canvasID, newNode, newVersion)
	if err != nil {
		log.Printf("fail to update canvas: %+v", err)
		return false
	}
	return true
}
