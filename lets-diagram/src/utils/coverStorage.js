/*
 * DATE: 2021/06/04
 * Author: Junebao
 * Doc: 提供缓存 cover 的方法，限定只缓存 MaxCoverNumber 个
 **/

import storage from "good-storage";

export const MaxCoverNumber = 20;

/**
 * Cover
 */
const CoverQueueKey = "let-diagram:coverQueue";

function getCoverStorageKey(id) {
  return "let-diageam:cover:" + id;
}

/**
 * 将一个 cover 加入到本地缓存
 * @param {Number} id 画布ID
 * @param {String} value 画布封面 Base64 编码
 */
export function set(id, value) {
  if (!storage.has(CoverQueueKey)) {
    storage.clear();
    let coverQueue = [];
    coverQueue.push(id);
    storage.set(CoverQueueKey, coverQueue)
    storage.set(getCoverStorageKey(id), value);
  }
  let coverQueue = storage.get(CoverQueueKey);
  if (coverQueue.length >= MaxCoverNumber) {
      let removeID = coverQueue.shift();
      storage.remove(getCoverStorageKey(removeID));
  }
  coverQueue.push(id);
  storage.set(getCoverStorageKey(id), value);
  storage.set(CoverQueueKey, coverQueue);
}

export function get(id) {
    return storage.get(getCoverStorageKey(id))
}

export function has(id) {
    return storage.has(getCoverStorageKey(id))
}

export function remove(id) {
    if (!has(id)) {
        return
    }
      if (!storage.has(CoverQueueKey)) {
        storage.clear();
        let coverQueue = [];
        storage.set(CoverQueueKey, coverQueue);
      }
      let coverQueue = storage.get(CoverQueueKey);
      let newQueue = []
      for (let i = 0; i < coverQueue.length; i++) {
          if (coverQueue[i] != id) {
              newQueue.push(coverQueue[i])
          }
      }
      storage.set(CoverQueueKey, newQueue);
      storage.remove(getCoverStorageKey(id))
}

export function clear() {
    storage.clear()
}