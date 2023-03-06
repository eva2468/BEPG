package models

import (
	"crypto/rand"
	"fmt"
	"io"

	"github.com/eva2468/bepg/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Header struct {
	Version   string
	Caller_ID string
	Token     string
	User_ID   string
}

type Card struct {
	gorm.Model
	Card_Header Header
	Card_no     string `gorm:""json:"card_no"`
}

type CheckBinResponse struct {
	Card_NO                     string `gorm:""json:"card_no"`
	Status                      string
	ErrorCode                   int64
	ErrorMessage                string
	QualifiedInternetPIN        string `json:"qualifiedinternetpin"`
	AuthenticationNotRequired   string `json:"authenticationnotrequired"`
	AvailableAuthMode           string `json:"availableauthmode"`
	AdditionalProductsSupported string `json:"additionalproductssupported"`
	Emi_availibility            string `json:"emi_availibility"`
	BinType                     string `json:"bintype"`
}

type Generate_Otp_Request struct {
	OTP_Header         Header
	Card_No            string `gorm:"primarykey";json:"card_no"`
	Card_Exp_date      string `json:"card_exp_date"`
	Card_CVD           string `json:"card_cvd"`
	Card_Holder_Status string `json:"card_holder_status"`
	Merchant_Name      string `json:"merchant_name"`
	Amount             string `json:"amount"`
	Mcc                string `json:"Mcc"`
	Currency_Code      string `json:"currency_code"`
	Shoper_IP_Address  string `json:"shoper_ip_address"`
	Language_Code      string `json:"language_code"`
	Request_ID         string `json:"request_id"`
}

type Generate_Otp_Response struct {
	Status         string
	ErrorCode      string
	ErrorMSG       string
	ValidityPeriod string
	TranCTX        string
}

type Verify_Otp_Request struct {
	Verify_otp_Header Header
	ShoperID          string
	RequestID         string
	TranCTX           string
	Otp               string
}

type Verify_Otp_Response struct {
	TranCTXID           string
	AuthResponseCode    string
	AuthenticationValue string
	ErrorMsg            string
}

type InitiateSI_Request struct {
	InitiateSI_Request_Header Header
	Card_No                   string `gorm:"unique";primarykey";json:"card_no"`
	Card_Exp_date             string `json:"card_expiry"`
	Card_CVD2                 string `json:"card_cvd2"`
	Card_Holder_Status        string `json:"card_holder_status"`
	Merchant_Name             string `json:"merchantName"`
	Amount                    string `json:"amount"`
	Mcc                       string `json:"mcc"`
	Card_Acceptor_ID          string `json:"card_acceptor_id"`
	IP_Address                string `json:"ipAddress"`
	Language_Code             string `json:"language_code"`
	Request_ID                string `json:"requestID"`
	StartDate                 string `json:"startDate"`
	EndDate                   string `json:"endDate"`
	NoOfSITxns                string `json:"noOfSITxns"`
	MinAmount                 string `json:"minAmount"`
	MaxAmount                 string `json:"maxAmount"`
	PreferredInitiationDate   string `json:"preferredInitiationDate"`
	AlterPreference           string `json:"alertPreference"`
	Currency_Code             string `json:"currency_code"`
}

type InitiateSI_Response struct {
	Status      string
	ErrorCode   string
	ErrorMsg    string
	Tran_ID     string
	RedirectURL string
}

type Merchant struct {
	BrowserUserAgent                  string
	IpAddress                         string
	HttpAccept                        string
	Language_Code                     string
	Transaction_Type_Indicator        string
	TID                               string
	Stan                              string
	Tran_time                         string
	Tran_Date                         string
	Mcc                               string
	Acquirer_Institution_Country_Code string
	Retrieval_Ref_Number              string
	Card_Acceptor_ID                  string
	Terminal_Owner_Name               string
	Terminal_City                     string
	Terminal_State_Code               string
	Terminal_Country_Code             string
	Merchant_Postal_Code              string
	Merchant_Telephone                string
}

type Complete_SIRegistration_Request struct {
	Complete_SIRegistration_Request_Header Header
	Merchant_SIRegistration_Request        Merchant
	Partner_ID                             string
	Tran_ID                                string
	AuthenticationMode                     string
}
type Complete_SIRegistration_Response struct {
	Status           string
	SiRegistrationID string
	ErrorCode        string
	ErrorMsg         string
}
type Bepg_authorize_Request struct {
	Bepg_authorize_Request_Header   Header
	Bepg_authorize_Request_Merchant Merchant
	Partner_ID                      string
	Auth_Amount                     string
	Authentication_Mode             string
	Authentication_Value            string
	Card_No                         string
	Card_Exp_Date                   string
	Cvd2                            string
	Order_ID                        string
	Custom1                         string
	Custom2                         string
	Custom3                         string
	Custom4                         string
	Custom5                         string
}

type Bepg_authorize_Response struct {
	Status   string
	ErroCode string
	ErrorMsg string
	Tran_ID  string
	ApprCode string
}

type Bepg_Revers_Request struct {
	Bepg_Reverse_Header  Header
	Stan                 string
	Tran_time            string
	Tran_Date            string
	Retrieval_Ref_Number string
}

type Bepg_Revers_Response struct {
	Status   string
	ErroCode string
	ErrorMsg string
}

type Check_EMIAvailibility_Request struct {
	Check_EMIAvailibility_Header Header
	Card_No                      string
	Amount                       string
	Currency_Code                string
	Mcc                          string
	Terminal_Owner_Name          string
}

type EmiDetails struct {
	Tenure        string
	RateOfInt     string
	EmiAmount     string
	ProcessingFee string
}

type Check_EMIAvailibility_Response struct {
	Status   string
	ErorCode string
	ErrorMsg string
	Emi      []EmiDetails
}

type DeRegister_SI_Request struct {
	DeRegister_SI_Header Header
	Card_Acceptor_ID     string
	Tran_ID              string
	Authentication_Mode  string
	Card_No              string
	Stan                 string
	Tran_Date            string
	Tran_time            string
	Retrieval_Ref_Number string
	SiRegistratioID      string
}

type DeRegister_SI_Response struct {
	Status    string
	ErrorCode string
	ErrorMsg  string
}

type Intimate_SI_Request struct {
	InitiateSI_Header         Header
	Card_Acceptor_Id          string
	Tran_Id                   string
	Card_No                   string
	Stan                      string
	Tran_Date                 string
	Tran_time                 string
	Retrieval_Ref_Number      string
	SiRegistrationID          string
	SiAmount                  string
	SiAlertPreference         string
	SiPreferredInitiationDate string
}

type Intimate_SI_Response struct {
	Status   string
	ErroCode string
	ErorMsg  string
}

type Deregister_SI_Advice_Request struct {
	DeRegister_SI_Advice_Header Header
	SiRegistrationID            string
}

type Deregister_SI_Advice_Response struct {
	Status  string
	ErorMsg string
}

/*
type Card_check struct {
	Card_no string `gorm:""json:"card_no"`
	Status  int64  `json:"status"`
}
*/

func init() {
	fmt.Println("Initiated.....!!!!!")
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Card{})
	db.AutoMigrate(&CheckBinResponse{})
	db.AutoMigrate(&Generate_Otp_Request{})
	db.AutoMigrate(&InitiateSI_Request{})
	//db.AutoMigrate(&Card_check{})

}

type ModelInterface interface {
	//CheckBin(*Card) *Card
	FuncForTest(string) string
}

func (card_holder *Card) CheckBin() *Card {

	db.NewRecord(card_holder)
	db.Create(&card_holder)
	return card_holder
}

func GetCardStatus(card_No string) (*CheckBinResponse, *gorm.DB) {

	var GetCard CheckBinResponse
	db := db.Where("card_no=?", card_No).Find(&GetCard)
	if GetCard.QualifiedInternetPIN == "1" {
		GetCard.Status = "Success"
		GetCard.ErrorMessage = "nil"
	} else {
		GetCard.Status = "failure"
		GetCard.ErrorCode = 1
		GetCard.ErrorMessage = "bin not enabled"
	}
	return &GetCard, db
}

func Generate_Otp(Gen_OTP *Generate_Otp_Request) Generate_Otp_Response {

	var Res_otp Generate_Otp_Response
	db.NewRecord(Gen_OTP)
	db.Create(&Gen_OTP)
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	var max int = 6
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	Res_otp.Status = "success"
	Res_otp.ErrorCode = "0"
	Res_otp.ErrorMSG = "nil"
	Res_otp.TranCTX = "0"

	return Res_otp

}

func Verify_Otp(Gen_OTP *Verify_Otp_Request) Verify_Otp_Response {

	var Res_otp Verify_Otp_Response
	if Gen_OTP.Otp == "2468" {
		Res_otp.TranCTXID = Gen_OTP.TranCTX
		Res_otp.ErrorMsg = ""
		Res_otp.AuthResponseCode = "ACCU000"
		Res_otp.AuthenticationValue = "MDEwMTIzNTQ2NzA5YWJjZGVmMDEyMzU0NjcwOWFiY2RlZjAxMjM1NDY3MDlhYmNkZWYwMTIzNTQ2NzA5YWJjZGVm100403210110040321011004032101"

	}

	return Res_otp

}

func InitiateSI(New_SI *InitiateSI_Request) *InitiateSI_Response {

	var res InitiateSI_Response
	db.NewRecord(New_SI)
	db.Create(&New_SI)
	res.Status = "success"
	res.ErrorCode = "0"
	res.ErrorMsg = ""
	res.Tran_ID = "102345678910234567891023456789"
	res.RedirectURL = "https://testnpci/redirect_otp/Home/IssuerReg?AccuCardholderId=89172389132&amp;AccuGuid=6298312e-c54c-1954-ba32-a9211038d150&amp;AccuHkey=5629y50g-e743-0022-5i2b-9aw8de632896"
	return &res
}

func GetEmiStatus(card_No string) (*Check_EMIAvailibility_Response, *gorm.DB) {

	var GetCard CheckBinResponse
	var res Check_EMIAvailibility_Response
	db := db.Where("card_no=?", card_No).Find(&GetCard)
	if GetCard.Emi_availibility == "1" {
		res.Status = "Success"
		res.ErrorMsg = "nil"
		res.ErorCode = "0"
	} else {
		res.Status = "failure"
		res.ErorCode = "1"
		res.ErrorMsg = "bin not enabled"
	}
	return &res, db
}
