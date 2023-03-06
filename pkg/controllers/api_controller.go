package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/eva2468/bepg/pkg/models"
	"github.com/eva2468/bepg/pkg/utils"
)

var CardHolder models.Card

/*
func CheckBin_Auth(w http.ResponseWriter, r *http.Request) {

	CreateCard := &models.Card{}
	version := r.Header.Get("version")
	caller_Id := r.Header.Get("caller_Id")
	token := r.Header.Get("token")
	user_Id := r.Header.Get("user_Id")
	utils.ParseBody(r, CreateCard)
	b := CreateCard.CheckBin()
	CardDetail := GetCardAuth(b.Card_no) // For checking the validity of the card
	if CardDetail.Status == 1 {
		res, _ := json.Marshal(CardDetail)
		w.Header().Set("version", version)
		w.Header().Set("caller_Id", caller_Id)
		w.Header().Set("token", token)
		w.Header().Set("user_Id", user_Id)
		w.WriteHeader(http.StatusOK)
		w.Write(res)

	} else {
		w.Header().Set("version", version)
		w.Header().Set("caller_Id", caller_Id)
		w.Header().Set("token", token)
		w.Header().Set("user_Id", user_Id)
		w.WriteHeader(http.StatusConflict)

	}

}*/

/*
func GetCardAuth(card_no string) *models.Card_check {

	card_No, err := strconv.ParseInt(card_no, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	CardDetail, _ := models.GetCardStatus(card_No)
	return (CardDetail)

}
*/
func Bepg_checkbin(w http.ResponseWriter, r *http.Request) {
	CreateCard := &models.Card{}
	Res_header := &models.Header{}
	utils.ParseBody(r, CreateCard)
	Res_header.Version = r.Header.Get("version")
	Res_header.Caller_ID = r.Header.Get("caller_Id")
	Res_header.Token = r.Header.Get("token")
	Res_header.User_ID = r.Header.Get("user_Id")
	b := CreateCard.CheckBin()
	CardDetail, _ := models.GetCardStatus(b.Card_no)
	response := struct {
		Status                      string `json:"status"`
		Errorcode                   int64  `json:"errocode"`
		Errormessage                string `json:"errormessage"`
		Qualifiedinternetpin        string `json:"qualified_internetpin"`
		Authenticationnotrequired   string `json:"authenticationNotRequired"`
		Availableauthmode           string `json:"available_AuthMode"`
		AdditionalProductsSupported string `json:"additionalProductsSupported"`
		Bintype                     string `json:"binType"`
	}{
		Status:                    CardDetail.Status,
		Errorcode:                 CardDetail.ErrorCode,
		Errormessage:              CardDetail.ErrorMessage,
		Qualifiedinternetpin:      CardDetail.QualifiedInternetPIN,
		Authenticationnotrequired: CardDetail.AuthenticationNotRequired,
		Availableauthmode:         CardDetail.AvailableAuthMode,
		Bintype:                   CardDetail.BinType,
	}
	res, _ := json.Marshal(response)
	w.Write(res)

}

func Generate_otp(w http.ResponseWriter, r *http.Request) {

	Gen_OTP := &models.Generate_Otp_Request{}
	Res_header := &models.Header{}
	Res_header.Version = r.Header.Get("version")
	Res_header.Caller_ID = r.Header.Get("caller_Id")
	Res_header.Token = r.Header.Get("token")
	Res_header.User_ID = r.Header.Get("user_Id")
	utils.ParseBody(r, Gen_OTP)
	b := models.Generate_Otp(Gen_OTP)
	/* otp */
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Verify_otp(w http.ResponseWriter, r *http.Request) {
	Ver_OTP := &models.Verify_Otp_Request{}
	Res_header := &models.Header{}
	Res_header.Version = r.Header.Get("version")
	Res_header.Caller_ID = r.Header.Get("caller_Id")
	Res_header.Token = r.Header.Get("token")
	Res_header.User_ID = r.Header.Get("user_Id")
	utils.ParseBody(r, Ver_OTP)
	b := models.Verify_Otp(Ver_OTP)

	/*
	 verify
	*/
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func Bepg_authorize(w http.ResponseWriter, r *http.Request) {

	//CreateCard := &models.Card{}
	version := r.Header.Get("version")
	caller_Id := r.Header.Get("caller_Id")
	token := r.Header.Get("token")
	user_Id := r.Header.Get("user_Id")
	//utils.ParseBody(r, CreateCard)
	//b := models.Generate_Otp()
	/* otp */
	//fmt.Println(b)
	w.Header().Set("version", version)
	w.Header().Set("caller_Id", caller_Id)
	w.Header().Set("token", token)
	w.Header().Set("user_Id", user_Id)
	w.WriteHeader(http.StatusOK)

}

func Bepg_reverse(w http.ResponseWriter, r *http.Request) {

	//CreateCard := &models.Card{}
	version := r.Header.Get("version")
	caller_Id := r.Header.Get("caller_Id")
	token := r.Header.Get("token")
	user_Id := r.Header.Get("user_Id")
	//utils.ParseBody(r, CreateCard)
	//b := models.Generate_Otp()
	/* otp */
	//fmt.Println(b)
	w.Header().Set("version", version)
	w.Header().Set("caller_Id", caller_Id)
	w.Header().Set("token", token)
	w.Header().Set("user_Id", user_Id)
	w.WriteHeader(http.StatusOK)

}
func Initiate_SIRegistration(w http.ResponseWriter, r *http.Request) {
	CreateSI := &models.InitiateSI_Request{}
	Res_header := &models.Header{}
	utils.ParseBody(r, CreateSI)
	Res_header.Version = r.Header.Get("version")
	Res_header.Caller_ID = r.Header.Get("caller_Id")
	Res_header.Token = r.Header.Get("token")
	Res_header.User_ID = r.Header.Get("user_Id")
	b := models.InitiateSI(CreateSI)
	res, _ := json.Marshal(b)
	w.Write(res)

}

func Check_EMIAvailability(w http.ResponseWriter, r *http.Request) {
	Emi_check := &models.Check_EMIAvailibility_Request{}
	Res_header := &models.Header{}
	utils.ParseBody(r, Emi_check)
	Res_header.Version = r.Header.Get("version")
	Res_header.Caller_ID = r.Header.Get("caller_Id")
	Res_header.Token = r.Header.Get("token")
	Res_header.User_ID = r.Header.Get("user_Id")
	b, _ := models.GetEmiStatus(Emi_check.Card_No)
	res, _ := json.Marshal(b)
	w.Write(res)

}

func complete_SIRegistration(w http.ResponseWriter, r *http.Request) {

}

func deregister_SI(w http.ResponseWriter, r *http.Request) {

}
