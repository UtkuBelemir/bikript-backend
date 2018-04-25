package models

import (
	"time"
	"strings"
)

type UserInfo struct {
	Email               string          `json:"email" gorm:"column:email;primary_key;default"`
	PhoneNumber         string          `json:"phone_number" gorm:"column:phone_number;default"`
	PhoneNumberPre      string        	`json:"phone_number_pre" gorm:"column:phone_number_pre;default"`
	Password            string          `json:"password" gorm:"column:password;default"`
	IsVerified          bool            `json:"is_verified" gorm:"column:is_verified;default"`
	Referrer            string          `json:"referrer" gorm:"column:referrer;default"`
	Name                string          `json:"name" gorm:"name;default"`
	Surname             string          `json:"surname" gorm:"surname;default"`
	BirthDay            time.Time       `json:"birth_day" gorm:"birth_day;default"`
	DocumentNumber      string          `json:"document_number" gorm:"document_number;default"`
	DocumentPicturePath string          `json:"document_picture_path" gorm:"document_picture_path;default"`
	SelfiePicturePath   string          `json:"selfie_picture_path" gorm:"selfie_picture_path;default"`
	EmailVerified		bool			`json:"email_verified" gorm:"email_verified"`
	PhoneVerified		bool			`json:"phone_verified" gorm:"phone_verified"`
	TwoFaVerified		bool			`json:"two_fa_verified" gorm:"two_fa_verified"`
}
func (UserInfo) TableName() string {
	return "sc_user.users"
}

type BuyOrder struct {
	RecordId     string    `json:"record_id" gorm:"column:record_id;primary_key"`
	UserId       string    `json:"user_id" gorm:"column:user_id"`
	CoinAmount   float64   `json:"coin_amount" gorm:"column:coin_amount"`
	PricePerCoin float64   `json:"price_per_coin" gorm:"column:price_per_coin"`
	TotalPair    float64   `json:"total_pair" gorm:"column:total_pair"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;default"`
	CoinType     string    `json:"coin_type" gorm:"-"`
	PairType     string    `json:"pair_type" gorm:"-"`
}
func (b BuyOrder) TableName() string {
	return strings.ToLower("sc_" + b.CoinType + ".buy_orders_" + b.PairType)
}

type SellOrder struct {
	RecordId     string    `json:"record_id" gorm:"column:record_id;primary_key"`
	UserId       string    `json:"user_id" gorm:"column:user_id"`
	CoinAmount   float64   `json:"coin_amount" gorm:"column:coin_amount"`
	PricePerCoin float64   `json:"price_per_coin" gorm:"column:price_per_coin"`
	TotalPair    float64   `json:"total_pair" gorm:"column:total_pair"`
	CreatedAt    time.Time `json:"created_at" gorm:"column:created_at;default"`
	CoinType     string    `json:"coin_type" gorm:"-"`
	PairType     string    `json:"pair_type" gorm:"-"`
}
func (s SellOrder) TableName() string {
	return strings.ToLower("sc_" + s.CoinType + ".sell_orders_" + s.PairType)
}

type WalletAddresses struct {
	PublicAddress string    `json:"public_address" gorm:"column:public_address;primary_key"`
	UserId        string    `json:"user_id" gorm:"column:user_id"`
	LastSyncTime  time.Time `json:"last_sync_time" gorm:"column:last_sync_time;default"`
	CoinType      string    `json:"coin_type" gorm:"-"`
}
func (wa WalletAddresses) TableName() string {
	return strings.ToLower("sc_" + wa.CoinType + ".wallet_addresses")
}

type Balances struct {
	UserId     string  `json:"user_id" gorm:"column:user_id;primary_key"`
	BTCBalance float64 `json:"btc_balance" gorm:"column:btc_balance"`
	ETHBalance float64 `json:"eth_balance" gorm:"column:eth_balance"`
	TRYBalance float64 `json:"try_balance" gorm:"column:try_balance"`
}
func (Balances) TableName() string {
	return "sc_user.users_balances"
}

type TradeHistory struct {
	RecordId        string    `json:"record_id" gorm:"column:record_id;primary_key"`
	BuyerId         string    `json:"buyer_id" gorm:"column:buyer_id"`
	SellerId        string    `json:"seller_id" gorm:"column:seller_id"`
	CoinAmount      float64   `json:"coin_amount" gorm:"column:coin_amount"`
	PricePerCoin    float64   `json:"price_per_coin" gorm:"column:price_per_coin"`
	TotalPair       float64   `json:"total_pair" gorm:"column:total_pair"`
	TransactionDate time.Time `json:"transaction_date" gorm:"column:transaction_date;default"`
	CoinType        string    `json:"coin_type" gorm:"-"`
	PairType        string    `json:"coin_type" gorm:"-"`
}
func (th TradeHistory) TableName() string {
	return strings.ToLower("sc_" + th.CoinType + ".trade_history_" + th.PairType)
}

type MailQueue struct {
	RecordId    string    `json:"record_id" gorm:"column:record_id;primary_key"`
	SendTo      string    `json:"send_to" gorm:"column:send_to"`
	MessageText string    `json:"messagetext" gorm:"column:messagetext"`
	Status      int       `json:"status" gorm:"column:status"`
	SendDate    time.Time `json:"send_date" gorm:"column:send_date;default"`
}
func (MailQueue) TableName() string {
	return "sc_queue.mail_queue"
}

type SMSQueue struct {
	RecordId    string    `json:"record_id" gorm:"column:record_id;primary_key"`
	SendTo      string    `json:"send_to" gorm:"column:send_to"`
	MessageText string    `json:"messagetext" gorm:"column:messagetext"`
	Status      int       `json:"status" gorm:"column:status"`
	SendDate    time.Time `json:"send_date" gorm:"column:send_date;default"`
}
func (SMSQueue) TableName() string {
	return "sc_queue.sms_queue"
}

type ActivationCodes struct {
	UserId    string `json:"user_id" gorm:"column:user_id"`
	Code      string `json:"code" gorm:"column:code"`
	Type      string `json:"code_type" gorm:"column:code_type"` // sms or email
	Reason	  string `json:"reason" gorm:"column:reason"`
	ExpiresAt string `json:"expires_at" gorm:"column:expires_at;default"`
}
func (ActivationCodes) TableName() string {
	return "sc_user.activation_codes"
}

type LoginLogs struct {
	UserId string  `json:"user_id" gorm:"column:user_id;primary_key"`
	Logs   LogsMdl `json:"logins" gorm:"column:logins"`
}
func (LoginLogs) TableName() string {
	return "sc_user.login_logs"
}

type LogsMdl struct {
	LoginDate time.Time `json:"transaction_date" gorm:"column:transaction_date"`
	IPAddres  string    `json:"transaction_date" gorm:"column:transaction_date"`
	Location  string    `json:"transaction_date" gorm:"column:transaction_date"`
	Browser   string    `json:"transaction_date" gorm:"column:transaction_date"`
	Device    string    `json:"transaction_date" gorm:"column:transaction_date"`
	Status    int       `json:"transaction_date" gorm:"column:transaction_date"`
}

type EmailModel struct {
	RecordId			string	`json:"email" gorm:"column:record_id;primary_key;default"`
	EmailSubject		string	`json:"email" gorm:"column:email_subject"`
	HtmlContent			string	`json:"email" gorm:"column:html_content"`
	InsertDate			string	`json:"email" gorm:"column:insert_date"`
	PlainTextContent	string	`json:"email" gorm:"column:plain_text_content"`
	SendDate			time.Time	`json:"email" gorm:"column:send_date"`
	SendTo				string	`json:"email" gorm:"column:send_to"`
	Status				int	`json:"email" gorm:"column:status"`
}
func (em EmailModel) TableName() string {
	return "sc_queue.mail_queue"
}

