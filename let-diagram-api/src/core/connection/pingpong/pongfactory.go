package pingpong

import (
    `lets_diagram/src/core/access`
    "lets_diagram/src/models"
    `lets_diagram/src/models/nottable`
)

func GetPermissionPong(user *models.User, t PongPermissionControlType,
    newPer access.CanvasJurisdictionMark) PongMessage {
	msg := PongPermissionData{
		Type: t,
		User: &nottable.SimpleUser{
			Username: user.Username,
			Email:    user.Email,
			ID:       user.ID,
		},
		NewPermission: newPer,
	}
	return PongMessage{
		Type: PongTypePermissionControl,
		Data: msg,
	}
}

func GetSTWPong(t PongRootCloseSTWType) PongMessage {
    return PongMessage{
        Type: PongTypeRootCloseStopTheWorld,
        Data: PongSTWData{Type: t},
    }
}

func GetSimpleNotifyPong(t PongSimpleNotifyType, user *models.User) PongMessage {
    switch t {
    case PongSimpleNotifyUserAdd:
        return PongMessage{Type: PongTypeSimpleNotify, Data: PongSimpleNotifyUserAddData{
            Type: t,
            User: &nottable.SimpleUser{
                ID:       user.ID,
                Username: user.Username,
                Email:    user.Email,
            },
        }}
    case PongSimpleNotifyUserOut:
        return PongMessage{Type: PongTypeSimpleNotify, Data: PongSimpleNotifyUserOutData{
            Type: t,
            User: &nottable.SimpleUser{
                ID:       user.ID,
                Username: user.Username,
                Email:    user.Email,
            },
        }}
    }
    return PongMessage{}
}
