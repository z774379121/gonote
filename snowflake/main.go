package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

const (
	SEQUENCE_BIT   = 12
	MACHINE_BIT    = 5
	DATACENTER_BIT = 5

	MAX_SEQUENCE_BIT   = -1 ^ (-1 << SEQUENCE_BIT)
	MAX_MACHINE_BIT    = -1 ^ (-1 << MACHINE_BIT)
	MAX_DATACENTER_BIT = -1 ^ (-1 << DATACENTER_BIT)

	// 相应的部分的左移长度
	MACHINE_LEFT    = SEQUENCE_BIT
	DATACENTER_LEFT = SEQUENCE_BIT + MACHINE_BIT
	TIMESTMP_LEFT   = DATACENTER_LEFT + DATACENTER_BIT
)

var sequence, lastStmp int64

type SnowFlake struct {
	datacenterId,
	machineId int64
}

func getNextMill() int64 {
	currStmp := time.Now().UnixNano() / 1000000
	for currStmp < lastStmp {
		currStmp = time.Now().UnixNano() / 1000000
	}
	return currStmp
}

func (this *SnowFlake) NextID() int64 {
	currStmp := getNextMill()
	// 相同毫秒内的id,序列号加一
	if currStmp == lastStmp {
		sequence = (sequence + 1) & MAX_SEQUENCE_BIT
		if sequence == 0 {
			// 该毫秒内的序列位数已经用完
			currStmp = getNextMill()
		}
	} else {
		sequence = 0
	}
	lastStmp = currStmp

	return (currStmp&0x1FFFFFFFFFF)<<TIMESTMP_LEFT | this.datacenterId<<DATACENTER_LEFT |
		this.machineId<<MACHINE_LEFT | sequence
}

func main() {
	sf := new(SnowFlake)
	// 这里表示机器码部分为0001000011
	sf.datacenterId = 2
	sf.machineId = 3
	for i := 0; i < 10000000; i++ {
		id := sf.NextID()
		fmt.Println("10进制表示：", id)
		fmt.Println("2进制原始表示：", strconv.FormatInt(id, 2))
		fmt.Println("36进制表示：", strconv.FormatInt(id, 36))
	}
	fmt.Println(math.Ceil(1.2))
	fmt.Println(math.Floor(1.2))
}
