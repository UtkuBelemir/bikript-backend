package bikriptDatabase
import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	bikriptModels "./models"
	"reflect"
	"errors"
	"fmt"
)
//Connection type for creating new database connection
type Connection struct {
	DBConneciton *gorm.DB
	DBerr       error
}
func NewConnection() *Connection {
	cnn := new(Connection)
	cnn.DBConneciton, cnn.DBerr = gorm.Open("postgres", "user=repidb password=23UktuBele23 dbname=bikript sslmode=disable")
	if cnn.DBerr != nil {
		log.Println(cnn.DBerr)
		return nil
	}
	return cnn
}
//This function saves given struct to database
func (dbCnn Connection) DBSave(tData interface{}) error {
	tx := dbCnn.DBConneciton.Begin()
	if tx.Error != nil {return tx.Error}
	switch newData := tData.(type) {
	case bikriptModels.UserInfo:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.BuyOrder:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.TradeHistory:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.SellOrder:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.WalletAddresses:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.SMSQueue:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.MailQueue:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.LoginLogs:
		if err := tx.Debug().Create(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	default:
		return errors.New("Unknown type : "+reflect.TypeOf(newData).String())
	}

	return nil
}
//This function updates given struct in database
func (dbCnn Connection) DBUpdate(tData interface{}) error {
	tx := dbCnn.DBConneciton.Begin()
	if tx.Error != nil {return tx.Error}
	switch newData := tData.(type) {
	case bikriptModels.UserInfo:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.BuyOrder:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.TradeHistory:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.SellOrder:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.WalletAddresses:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.SMSQueue:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.MailQueue:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	case bikriptModels.LoginLogs:
		if err := tx.Model(&newData).Debug().Update(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
		break
	default:
		fmt.Println("NEW DATA",newData)
		return errors.New("Unknown type : "+reflect.TypeOf(newData).String())
	}

	return nil
}
//This function deletes given struct from database
func (dbCnn Connection) DBDelete(tData interface{}) error {
	tx := dbCnn.DBConneciton.Begin()
	if tx.Error != nil {return tx.Error}
	switch newData := tData.(type) {
	case bikriptModels.UserInfo:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.BuyOrder:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.TradeHistory:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.SellOrder:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.WalletAddresses:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.SMSQueue:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.MailQueue:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	case bikriptModels.LoginLogs:
		if err := tx.Debug().Delete(&newData).Error; err != nil {tx.Rollback();return err}
		if err := tx.Commit().Error; err != nil {return err}
	default:
		return errors.New("Unknown type : "+reflect.TypeOf(newData).String())
	}

	return nil
}