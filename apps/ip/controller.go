package ip

import (
    "github.com/3th1nk/cidr"
    "github.com/go-ping/ping"
    "time"
)

// ProbeCIDR 获取指定网段的所有IP地址
func ProbeCIDR(CIDR string) ([]string, error) {

    var ips []string

    c, err := cidr.ParseCIDR(CIDR)
    if err != nil {
        return nil, err
    }

    if err := c.ForEachIP(func(ip string) error {
        ips = append(ips, ip)
        return nil
    }); err != nil {
        return nil, err
    }

    return ips, nil
}

// PingIP TODO 这里的探测功能实现的过于简陋，等待后面有时间时，回头来重写完善一下
// 使用ping探测检查ping是否被占用
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
        //utils.Log().Infof("发送数据包=%v，接收数据包=%v，接收数据包重复=%v，丢失数据包=%v，Rtts=%v，目标地址=%v，IP=%v，最大Rtt=%v，最小Rtt=%v，平均Rtt=%v，标准差Rtt=%v",
        //    stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss,
        //    stats.Rtts, stats.Addr, stats.IPAddr, stats.MaxRtt, stats.MinRtt, stats.AvgRtt, stats.StdDevRtt)

        //被占用的ip，返回false
        return false, nil
    } else {
        //fmt.Println("闲置：", ip)

        //空闲可用的ip，返回true
        return true, nil
    }
}
