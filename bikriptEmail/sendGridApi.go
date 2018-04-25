// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package bikriptEmail

import (
	"fmt"
	"os"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	bikriptDB "../bikriptDatabase"
	bikriptModels "../bikriptDatabase/models"
	"time"
)

const senderName = "Bikript Bitcoin Borsası A.Ş."
const senderMail = "noreply@bikript.com"
func CheckEmail(dbCnn *bikriptDB.Connection){
	var unsentMails []bikriptModels.EmailModel
	for {
		unsentMails = nil
		dbCnn.DBConneciton.Where("status = ?",0).Find(&unsentMails)
		for _, oneMail := range unsentMails{
			err := SendEmail(oneMail.EmailSubject,oneMail.SendTo,oneMail.SendTo,oneMail.PlainTextContent,oneMail.HtmlContent)
			if err != nil{
				dbCnn.DBConneciton.Model(bikriptModels.EmailModel{}).Where("record_id = ?",oneMail.RecordId).Updates(bikriptModels.EmailModel{Status:99})
				fmt.Println("Mail didn't send",oneMail.Status," to user ",oneMail.SendTo)
			}else{
				dbCnn.DBConneciton.Model(bikriptModels.EmailModel{}).Where("record_id = ?",oneMail.RecordId).Updates(bikriptModels.EmailModel{Status:1,SendDate:time.Now()})
			}
		}
		time.Sleep(30 * time.Second)
	}
}
func SendEmail(subject string,receiverName string,receiverEmail string,plainTextContent string,htmlContent string) error{
	from := mail.NewEmail(senderName, senderMail)
	to := mail.NewEmail(receiverName, receiverEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	_, err := client.Send(message)
	if err != nil {
		fmt.Println("EROOR")
		return err
	} else {
		fmt.Println("gonderdim ok")
		return nil
	}
}