package snowflake

import (
	"errors"
	"log"
	"sync"
	"time"
)

// 雪花算法
/**
  0 | 00000000 00000000 00000000 00000000 00000000 0 / 00000000 00 / 00000000 0000
  1bit不用 |             41bit时间戳                  / 10bit工作机器ID / 12bit序列号

  1bit不用，二进制中最高位为1代表负数
  41bit代表时间戳，(2^41−1)/(1000∗60∗60∗24∗365)=69 年
  10bit工作机器ID允许分布式最大节点数为1024个（10个bit按场景划分，总计组合1024个）
  12bit序列号表示每毫秒生成的ID序号 2^12-1=4095，即0，1，...4095共4096个序列号

  实际情况下，可以根据具体场景来划分10bit工作机器ID和12bit序列号，
  比如说，把工作机器ID缩减到4位，把bit序列号扩大到18位来满足机器不足，单机可生成更多但序列号ID
*/

const (
	timestampOffset uint8 = 22 // 时间戳偏移量
	workerOffset    uint8 = 12 // 工作机器偏移量

	workerBits   int64 = 10 // 工作机器ID所占bit数
	sequenceBits int64 = 12 // 序列号所占bit数

	maxWorkerId   int64 = -1 ^ (-1 << uint64(workerBits))   // 最大工作机器ID
	maxSequenceId int64 = -1 ^ (-1 << uint64(sequenceBits)) // 最大序列号ID

	epoch int64 = 1599029955445
)

type Worker struct {
	mu         sync.Mutex
	timestamp  int64 // 上一次生成ID的时间戳
	workerId   int64 // 工作机器ID
	sequenceId int64 // 每毫秒的序列号，从0开始，最多4096个
}

func NewSnowflakeWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > maxWorkerId {
		return nil, errors.New("WorkerId big than maxWorkerId")
	}

	return &Worker{
		timestamp:  0,
		workerId:   workerId,
		sequenceId: 0,
	}, nil
}

func (w *Worker) GenerateId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()

	timestamp := w.getCurrentTime()
	if timestamp < w.timestamp {
		log.Fatal("can not generate id")
	}

	if timestamp == w.timestamp {
		w.sequenceId++
		if w.sequenceId > maxSequenceId {
			for timestamp <= w.timestamp {
				timestamp = w.getCurrentTime()
			}
			w.sequenceId = 0
			w.timestamp = timestamp
		}
	} else {
		w.sequenceId = 0
		w.timestamp = timestamp
	}

	id := (timestamp-epoch)<<timestampOffset | (w.workerId << workerOffset) | w.sequenceId
	return id
}

func (w *Worker) getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}
