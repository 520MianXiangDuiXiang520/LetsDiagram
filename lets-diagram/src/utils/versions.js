/*
 * DATE: 2021/04/27
 * Author: Junebao
 * Doc: versions 定义一个固定长度的版本库，用来保存历史版本数据
 *   当两个协作者同时操作画布时，将按照操作到达服务器的时间决定其版本大小。
 *   换言之，由于并发，客户端收到的 diff 版本可能是基于旧版本的，versions
 *   就用来保存这些旧版本数据。
**/

// 用来保存每一个版本号对应的数据
let versionDataMap = {}

// 版本号队列，用来控制哪些版本的数据应该被删除
let versionIDQueue = []

const MaxVersionQueueLen = 5

/**
 * 将同步到的新版本数据插入版本库中，每当客户端收到一个版本控制的 Pong 消息，
 * 在合并 diff 之后，重新 open canvas 之前都要先执行该方法保存副本。
 * @param {Number} version: 版本号 
 * @param {*} data: 该版本的数据（这些数据是经过 diff 合并后的完整数据） 
 */
export function insertVersion(version, data) {
    if (versionIDQueue.length >= MaxVersionQueueLen) {
        let rv = versionIDQueue.pop()
        delete versionDataMap[rv]
    }
    versionDataMap[version] = data;
    versionIDQueue.push(version);
}

/**
 * 获取 version 对应的版本数据，在合并 diff 之前，应该根据 Pong 中的 oldVersion
 * 获取对应的数据。
 * @param {Number} version 
 * @returns {Object} 返回该版本对应的完整数据，如果版本库中没有对应的数据，返回 undefined
 * 这种情况下应该向服务器发送心跳 Ping, 以获取最新数据
 */
export function getVersionData(version) {
    let res = versionDataMap[version]
    if (!res) {
        console.log("can not find data of this version: (" + version + ")!");
        console.log(versionDataMap);
    }
    return res
}