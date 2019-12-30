package main

import (
    "checkssl/check"
    "checkssl/conf"
    "checkssl/notices/email"
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    config, err := conf.NewConfigInfo()
    if err != nil {
        log.Println("init config file failed: ", err)
        return
    }
    checkInfo := check.NewCheckInfo(config.Domains.Domain)
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt,syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)


    ticker := time.NewTicker(time.Second)

    for {
        select {
        case <-ticker.C:
            checkInfo.Check()
        case msg:=<- checkInfo.Ch:
            sendMail := email.NewEmail(config.Mails.Mail, msg)
            sendMail.SendMail()
            log.Println("邮件发送完成")
        case <- signals:
            //ticker.Stop()
            log.Println("stop successfully")
            return
        }
    }
}
