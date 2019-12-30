package check

import (
    "fmt"
    "net/http"
    "sync"
    "time"
)

var (
    wg sync.WaitGroup
)

type CheckInfo struct {
    Address []string
    Ch   chan string
}

func (c *CheckInfo) Request(url string ) (msg string) {
    tr := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    30 * time.Second,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}
    resp, _ := client.Get(fmt.Sprintf("https://%s", url))
    defer resp.Body.Close()
    Now := time.Now()
    //MaxTTL := resp.TLS.PeerCertificates[0].NotAfter.Format("2006-01-02 15:04:05")
    MaxTTL := resp.TLS.PeerCertificates[0].NotAfter
    sub := MaxTTL.Sub(Now)
    switch {
    case sub.Hours()/24 < 20:
        msg = fmt.Sprintf("%s 证书还有二十天过期，请及时更新证书，避免造成不必要的损失....\n", url)
        return
        //log.Println("证书还有三天过期，请及时更新证书，避免造成不必要的损失....")
    case sub.Hours()/24 < 10:
        msg = fmt.Sprintf("%s 证书还有十天过期，请及时更新证书，避免造成不必要的损失.... \n", url)
        return
        // log.Println("证书还有十天过期，请及时更新证书，避免造成不必要的损失....")
    case sub.Hours()/24 < 3:
        msg = fmt.Sprintf("%s 证书还有三天过期，请及时更新证书，避免造成不必要的损失.... \n", url)
        return
        // log.Println("证书还有二十天过期，请及时更新证书，避免造成不必要的损失....")
    }
    //for _, tlsValue := range resp.TLS.PeerCertificates{
    //   //log.Println(tlsValue.NotBefore)
    //   log.Println(tlsValue.NotAfter)
    //}
    return
}

func (c *CheckInfo) Check() {
    for _,url := range  c.Address{
        fmt.Println(url)
       msg := c.Request(url)
       c.Ch <- msg
    }
    return
}
func NewCheckInfo(address []string) *CheckInfo {
    return &CheckInfo{
        Address: address,
        Ch: make(chan string,4),
    }
}
