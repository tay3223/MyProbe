package main

import (
    "github.com/go-ping/ping"
    "ipProbe/utils"
    "time"
)

func PingIP(ip string) (bool, error) {
    pinger, err := ping.NewPinger(ip)
    if err != nil {
        panic(err)
        return false, nil
    }

    pinger.Count = 1
    pinger.Timeout = 1 * time.Second
    err = pinger.Run()
    stats := pinger.Statistics()
    if stats.PacketsRecv != 0 && len(stats.Rtts) != 0 {
        utils.Log().Infof("发送数据包=%v，接收数据包=%v，接收数据包重复=%v，丢失数据包=%v，Rtts=%v，目标地址=%v，IP=%v，最大Rtt=%v，最小Rtt=%v，平均Rtt=%v，标准差Rtt=%v",
            stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss,
            stats.Rtts, stats.Addr, stats.IPAddr, stats.MaxRtt, stats.MinRtt, stats.AvgRtt, stats.StdDevRtt)

        //被占用的ip，返回false
        return false, nil
    } else {
        //fmt.Println("闲置：", ip)

        //空闲可用的ip，返回true
        return true, nil
    }
}

func main() {

}
