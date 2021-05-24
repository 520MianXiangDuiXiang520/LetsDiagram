/*
 * DATE: 2021/04/27
 * Author: Junebao
 * Doc: connection 用来管理整个项目的 websocket 连接
 *   客户端发送给服务器的消息称为 Ping, 服务器发送给客户端的消息称为 Pong
 *   Ping 和 Pong 有不同的类型，定义在 const.js 里。
 **/

// websocket 连接管理
import * as c from "./const.js";
import { insertVersion, getVersionData } from "./versions.js";
import * as jsonpatch from "fast-json-patch/index.mjs";
import * as msgFactory from "./message.js";
let webSocket = "";

export async function newWSConn(url, onOpen, onMessage) {
  // if (webSocket != "") {
  //   return webSocket;
  // }
  if (webSocket != "") {
    webSocket.close();
  }
  // webSocket && webSocket.close();
  if (webSocket) {
    webSocket.close();
    webSocket = "";
  }
  if (!webSocket) {
    console.log("正在建立 WebSocket 连接");
    webSocket = new WebSocket(url);
    webSocket.onopen = onOpen;
    webSocket.onmessage = onMessage;
  } else {
    console.log("websocket已连接");
  }
  console.log("连接 end");
  return webSocket;
}

export function closeWebsocket(self) {
  if (webSocket) {
    webSocket.close();
  }
  webSocket = ""
  self.webSocket = ""
}

// webSocket 建立好后, 开启定时任务，发 ping 包
export function withOpen(self) {
  return function() {
    console.log("连接已打开！");
    webSocket.send(
      JSON.stringify({
        type: c.PingTypeHeartbeat,
        data: "ping",
      })
    );
    let f = sendVersionControlPing(self);
    self.timer = setInterval(f, 4);
  };
}

// 将不同类型的 pong 消息转发给特定的方法处理
export function disponse(self) {
  return function(e) {
    let msg = JSON.parse(e.data);
    let msgType = msg["type"];
    switch (msgType) {
      case c.PongTypeHeartbeatControl: {
        conductHeartbeatMessage(msg, self);
        break;
      }
      case c.PongTypeVersionContral: {
        conductVersionControlMessage(msg, self);
        break;
      }
      case c.PongTypePermissionControl: {
        conductPermissionControlMessage(msg, self);
        break;
      }
      case c.PongTypeRootCloseStopTheWorld: {
        conductSTWMessage(msg, self);
        break;
      }
      case c.PongTypeSimpleNotify: {
        conductSimpleNotifyMessage(msg["data"], self);
        break;
      }
    }
  };
}

// 向服务端发送心跳消息，获取未持久化的版本
export async function sendHeartbeatPing(self) {
  console.log("heartbeat");
  console.log(self.webSocket);
  console.log(webSocket);
  webSocket.send(
    JSON.stringify({
      type: c.PingTypeHeartbeat,
      data: "ping",
    })
  );
}

// 发送版本控制信息
export function sendVersionControlPing(self) {
  return function() {
    insertVersion(self.version, self.oldTopologyData);
    let newData = window.topology.pureData();
    let diff = jsonpatch.compare(self.oldTopologyData, newData, false);
    self.oldTopologyData = newData;
    if (self.authority > c.CanvasJurisdictionMarkReadOnly && diff.length > 0) {
      webSocket.send(
        msgFactory.createVersionControlPingMsg(self.version, diff)
      );
      self.version++;
    }
  };
}

export function createPermissionPingMessage(type, user, permission) {
  return JSON.stringify({
    type: c.PingTypePermissionControl,
    data: {
      type: type,
      user: user,
      new_permission: permission,
    },
  });
}

// 发送一个 PingPermissionApplication（向 manager 申请写权限）
export function sendWriteApplication(self, manager) {
  console.log(self);
  webSocket.send(
    createPermissionPingMessage(c.PingPermissionApplication, manager)
  );
}

// 发送允许请求的 ping
export function sendPermissionAllowed(user) {
  if (webSocket != null) {
    webSocket.send(
      createPermissionPingMessage(
        c.PingPermissionAllowed,
        user,
        c.CanvasJurisdictionMarkWrite
      )
    );
  }
}

export function sendPermissionDenied(user) {
  if (webSocket != null) {
    webSocket.send(
      createPermissionPingMessage(
        c.PingPermissionDenied,
        user,
        c.CanvasJurisdictionMarkWrite
      )
    );
  }
}

// 修改用户权限
export function sendSetPermission(user, permission) {
  if (webSocket != null) {
    webSocket.send(
      createPermissionPingMessage(c.PingPermissionSet, user, permission)
    );
  }
}

let countdownTimer = "";

let countdownFunc = function(self) {
  return function() {
    if (self.seconds > 0) {
      self.seconds--;
    } else {
      clearInterval(countdownTimer);
      self.$router.push("/index");
    }
  };
};

//
export function conductSTWMessage(msg, self) {
  let t = msg["data"]["type"];
  switch (t) {
    case c.PongRootCloseSTWCrash: {
      self.stw = true;
      self.stwDialog = true;
      countdownTimer = setInterval(countdownFunc(self), 1000);
      break;
    }
    case c.PongRootCloseSTWRecovery: {
      clearInterval(countdownTimer);
      self.stw = false;
      self.stwDialog = false;
      break;
    }
  }
}

export function conductSimpleNotifyMessage(msg, self) {
  let t = msg["type"]
  switch(t) {
    case c.PongSimpleNotifyUserAdd: {
      let username = msg["user"]["username"]
      let email = msg["user"]["email"]
      self.$message.success(username + "(" + email + ") 已加入协作")
      break
    }
    case c.PongSimpleNotifyUserOut: {
      let username = msg["user"]["username"];
      let email = msg["user"]["email"];
      self.$message.info(username + "(" + email + ") 已退出协作");
      break;
    }
  }
}

// 处理心跳 pong, 不管是不是自己创建的版本，都绘制
export function conductHeartbeatMessage(msg, self) {
  let diffs = msg["data"];
  for (let i = 0; i < diffs.length; i++) {
    let diff = diffs[i]["diffs"];
    let baseVersion = diffs[i]["base_version"];
    let serverVersion = diffs[i]["version"];
    let newData = jsonpatch.applyPatch(getVersionData(baseVersion), diff)
      .newDocument;
    insertVersion(serverVersion, newData);
    window.topology.open(newData);
    self.version = serverVersion;
    self.oldTopologyData = window.topology.pureData();
  }
}

// 处理版本控制 pong，自己创建的版本不绘制
export function conductVersionControlMessage(msg, self) {
  let diffs = msg["data"];
  for (let i = 0; i < diffs.length; i++) {
    let writer_name = diffs[i]["user"]["Username"];
    if (diffs[i]["is_writer"]) {
      let diff = diffs[i]["diffs"];
      let baseVersion = diffs[i]["base_version"];
      let serverVersion = diffs[i]["version"];
      let newData = jsonpatch.applyPatch(getVersionData(baseVersion), diff)
        .newDocument;
      insertVersion(serverVersion, newData);
      window.topology.open(newData);
      self.version = serverVersion;
      self.oldTopologyData = window.topology.pureData();
      self.writer_name = writer_name;
    } else {
      self.writer_name = "您";
    }
  }
}

// 处理权限相关 pong TODO
export function conductPermissionControlMessage(msg, self) {
  let permissionData = msg["data"];
  let permissionType = permissionData["type"];
  let username = permissionData["user"]["username"];
  let email = permissionData["user"]["email"];
  let newPermission = permissionData["new_permission"];
  self.permissionPingUserID = permissionData["user"]["id"];
  switch (permissionType) {
    case c.PongPermissionAllowed: {
      // 当前用户获得了写权限
      self.$message.success(username + "(" + email + ")同意了您的请求！");
      self.authority = c.CanvasJurisdictionMarkWrite;
      self.permissionStyle = "";
      break;
    }

    case c.PongPermissionDenied: {
      self.$message.error(username + "(" + email + ")拒绝了您的请求！");
      break;
    }

    case c.PongPermissionApplication: {
      self.permissionMessage =
        username + "(" + email + ")请求获取操作权限，您是否允许？";
      self.dialogVisible = true;
      break;
    }

    case c.PongPermissionSet: {
      self.authority = newPermission;
      switch (newPermission) {
        case c.CanvasJurisdictionMarkReadOnly: {
          self.permissionStyle = "pointer-events: none";
          self.$message.info("管理员将您的权限设置为了【只读】");
          break;
        }

        case c.CanvasJurisdictionMarkWrite: {
          self.permissionStyle = "";
          self.$message.info("管理员将您的权限设置为了【读写】");
          break;
        }

        case c.CanvasJurisdictionMarkManager: {
          self.permissionStyle = "";
          self.$message.info("您已被设置为了 【管理员】");
          break;
        }
      }
      break;
    }

    case c.PongPermissionKickOut: {
      webSocket.close();
      self.$message.error("您已被踢出协作，请联系管理员！");
      self.$router.push("/index");
    }
  }
}

export function sendPermissionApplication() {
  if (this.webSocket != null) {
    this.webSocket.send(
      JSON.stringify({
        type: c.PingTypePermissionControl,
        data: {
          type: c.PingPermissionApplication,
          user: 0,
        },
      })
    );
  }
}
