/*
 * DATE: 2021/04/27
 * Author: Junebao
 * Doc: message 用来产生不同类型的 ping 消息包
 **/

import * as c from "./const.js";

/**
 * 生成一个用于版本控制的 ping 消息
 * @param {Number} baseVersion 该 diff 是基于哪一版本的数据产生的 
 * @param {*} diffs 
 * @returns 返回 json 序列化后的消息
 */
export function createVersionControlPingMsg(baseVersion, diffs) {
  return JSON.stringify({
    type: c.PingTypeVersionControl,
    data: {
      base_version: baseVersion,
      diffs: diffs,
    },
  });
}
