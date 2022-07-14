package ip

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "net/http"
)

type ProbeIpRequest struct {
    CIDR string `json:"cidr"`
}

//type ProbeIpResponse struct {
//    Ips    []map[string]string `json:"ips"`
//}

// ProbeIp
//@Tags 扫描
//@Summary 扫描ip网段
//@Description 扫描ip网段
//@Accept json
//@Produce json
//@Success 200 {string} Hello
//@Router /api/v1/probe-ip [post]
func ProbeIp(g *gin.Context) {
    //获取request数据
    pir := ProbeIpRequest{}
    if err := g.ShouldBindJSON(&pir); err != nil {
        //if err := g.ShouldBind(&pir); err != nil {
        return
    }

    fmt.Println(&pir)

    //探测CIDR
    //ipList, err := ProbeCIDR("172.26.120.1/24")
    ipList, err := ProbeCIDR(pir.CIDR)
    if err != nil {
        panic(err)
    }

    //多协程并行处理探测
    resultList := Multi(ipList)
    fmt.Println(resultList)

    var ips []map[string]string
    for ip, ok := range resultList {
        if ok {
            tmpData := make(map[string]string)
            tmpData["ip"] = ip
            ips = append(ips, tmpData)
        }
    }

    //给出返回值
    g.JSON(http.StatusOK, ips)
}
