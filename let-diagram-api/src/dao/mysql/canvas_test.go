package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"lets_diagram/src/models"
	"reflect"
	"testing"
)

func TestSelectCanvasInfoByUserID(t *testing.T) {
	canvas, err := SelectCanvasInfoByUserID(uint(2), 0, 10)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(canvas)
}

func TestSelectByUserIDCanvasID(t *testing.T) {
	type args struct {
		canvasID uint
		userID   uint
	}
	tests := []struct {
		name    string
		args    args
		want    *models.UserCanvas
		wantErr bool
	}{
		{"test1", args{canvasID: uint(3), userID: uint(13)}, &models.UserCanvas{
			Model:    gorm.Model{ID: 2},
			UserID:   0,
			CanvasID: 0,
		}, false},
		{"test1", args{canvasID: uint(2), userID: uint(13)}, &models.UserCanvas{
			Model:    gorm.Model{},
			UserID:   0,
			CanvasID: 0,
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SelectByUserIDCanvasID(tt.args.canvasID, tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("SelectByUserIDCanvasID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectByUserIDCanvasID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
