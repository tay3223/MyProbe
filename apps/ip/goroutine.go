package ip

import "sync"

type Result struct {
    ip string
    ok bool
}

const num = 500

func Multi(list []string) map[string]bool {
    ipCh := make(chan string)     //信道
    resultCh := make(chan Result) //信道

    go func() {
        for _, ip := range list {
            ipCh <- ip
        }
        close(ipCh)
    }()

    go func() {
        var wg sync.WaitGroup
        for i := 0; i < num; i++ {
            wg.Add(1)
            go func() {
                defer wg.Done()
                for ip := range ipCh {
                    ok, _ := PingIP(ip)
                    resultCh <- Result{ip: ip, ok: ok}
                }
            }()
        }
        wg.Wait()
        close(resultCh)
    }()

    ret := map[string]bool{}
    for result := range resultCh {
        ret[result.ip] = result.ok
    }
    return ret
}
