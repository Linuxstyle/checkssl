package email

import (
    "fmt"
    "log"
    "net/smtp"
)

type Email struct {
    To []string
    Msg string
}


func (e *Email)SendMail(){
    auth := smtp.PlainAuth("","123@qq.com","1223456","smtp.exmail.qq.com")
    msg := []byte(fmt.Sprintf("From: 123@qq.com\r\n"+"To: %s\r\n"+"Subject: Check Domain SSL\r\n"+"\r\n"+"%s \r\n",e.To,e.Msg))
    err := smtp.SendMail("smtp.exmail.qq.com:25",auth,"123@qq.com",e.To,msg)
    if err != nil {
        log.Println("Send Mail Failed: ",err)
        return
    }
}


func NewEmail(to []string,msg string )*Email{
    return &Email{
        To: to,
        Msg: msg,
    }
}
