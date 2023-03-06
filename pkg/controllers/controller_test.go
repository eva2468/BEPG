package controllers

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_Bepg_checkbin(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/bepg_checkbin.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/bepg_chekbin", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Bepg_checkbin(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"status":"Success","errocode":0,"errormessage":"nil","qualified_internetpin":"1","authenticationNotRequired":"","available_AuthMode":"","additionalProductsSupported":"","binType":""}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func Test_Bepg_generate_Otp(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/generate_otp.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/generate_otp", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Generate_otp(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"Status":"success","ErrorCode":"0","ErrorMSG":"nil","ValidityPeriod":"","TranCTX":"0"}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func Test_Bepg_verify_Otp(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/verify_otp.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/verify_otp", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Verify_otp(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"TranCTXID":"100000000000000000000000025253","AuthResponseCode":"ACCU000","AuthenticationValue":"MDEwMTIzNTQ2NzA5YWJjZGVmMDEyMzU0NjcwOWFiY2RlZjAxMjM1NDY3MDlhYmNkZWYwMTIzNTQ2NzA5YWJjZGVm100403210110040321011004032101","ErrorMsg":""}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func Test_Bepg_Initiate_SI(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/Initiate_SI.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/initiate_SIRegistration", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Initiate_SIRegistration(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"Status":"success","ErrorCode":"0","ErrorMsg":"","Tran_ID":"102345678910234567891023456789","RedirectURL":"https://testnpci/redirect_otp/Home/IssuerReg?AccuCardholderId=89172389132\u0026amp;AccuGuid=6298312e-c54c-1954-ba32-a9211038d150\u0026amp;AccuHkey=5629y50g-e743-0022-5i2b-9aw8de632896"}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}
func Test_Bepg_CompleteRegistration_SI(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/CompleteRegister_SI.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/complete_SIRegistration", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Complete_SIRegistration(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"Status":"success","SiRegistrationID":"1","ErrorCode":"0","ErrorMsg":"nil"}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func Test_Bepg_Check_EMI(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/check_emi.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/check_EMIAvailability", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Check_EMIAvailability(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"Status":"Success","ErorCode":"0","ErrorMsg":"nil","Emi":null}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}

func Test_Bepg_Deregister_SI(t *testing.T) {
	data, _ := ioutil.ReadFile("test_json/Deregister_SI.json")
	req := httptest.NewRequest("POST", "http://localhost:9090/deregister_SI", bytes.NewReader(data))
	w := httptest.NewRecorder()

	// Call the handler function to handle the request
	Deregister_SI(w, req)

	// Check the response status code and body
	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", resp.StatusCode, http.StatusOK)
	}
	expected := `{"Status":"Success","ErrorCode":"0","ErrorMsg":"nil"}`
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("unexpected body: got %v, want %v", string(body), expected)
	}
}
