package main

import (
	"fmt"
	"transcookie/pkg/request"
)

func main() {
	cookieStr := "_aegis_cas=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJuYmYiOjE2Njc5NjM2MzUsInN1YiI6ImxpdXlhbmFuMTEiLCJkZXRhaWwiOiKUxFx1MDAwNlx1MDAwZWPlM5FcdTAwMTV2h1x1MDAwNOt86WPWXHUwMDA1U1xmiFxyOJG6yMS_VZU-PTJcItqUetuA5vNcdTAwMTDzolxcJknTXHUwMDBiQ62sMkvqM7wpXHUwMDA1XHUwMDFiXHUwMDA1eW_1MaiNxppKI1x1MDAwZcbzw1x1MDAxOMq1XHUwMDE0PTraTuiTz3xq6TdMXHUwMDEw7WHk0Y3PtX5kOnp9mVv757ib8ihPkqrUXHUwMDE5X0JeaELUVLSw-87LV12kpZZcdTAwMTehXGaq61x1MDAxZpGYXfuxypJcblx1MDAwNOH_csG7Z2QnaPHnj3lcdTAwMWXCXYtcdTAwMDS8TXSysZecvdbsjrBFaDEnnptMgWrFqorZxHRDyT82JFwvXHUwMDE3av_OXHUwMDAxP8xcdTAwMTFbXdlgYf1ZV8KgXoGWr13kM1x1MDAxZFx1MDAxZjLmp1x1MDAwMFK9x8tcdTAwMTFYynKlVCH3QNdoXHUwMDEw9pz7M0_bTl_kI6nXaDaPJtWDxTXIPemCXHUwMDBiXHUwMDFkOsOeqbPa-T87PrV2XHUwMDBiQTZNjnPnXHUwMDAyxcPDXHUwMDFk26Dy7Vx1MDAxNshcdTAwMTfugbE3jUT_tDNI571YXHUwMDA1aFx1MDAwYk-5gkJ8XHUwMDFhu-RRXHUwMDE3Q-k5ftjwXHUwMDFjvKiuIX6zkVx1MDAxMapxOtrdZq3r6utcdTAwMGI0tFx1MDAxMqJK4VssmlwiW5PVTVx1MDAwNiaIUWS6OrA1P5X5u3ldxVx1MDAxZNv5WlAzI1l7m9DJxr29ZN1dlXhErydN-FTfNjFBY1x1MDAxM6VZXHUwMDFht4l1XFxcdTAwMTnlXGKPuHaNQPRcdTAwN2b4pUSKjzijjkL1xnptXHUwMDEwXHUwMDAy8lx1MDAwM_qFWNuHorbxPtwnakV5_9C4vnPx4lSC9HVcdTAwMWNcdTAwMGVcdTAwMDQ7dZLT-4yNr0hhfVx1MDAxZO5cdTAwMTGSkmNCqUdd-CgwfFxmSlBcdTAwMDQrVlCn_ZqBXHUwMDAw7CdpaVx1MDAxNVwvJMPW9Vx1MDA3ZlxiWdLf5pgpXHTi1Fx1MDAwNXrKc1wvl6dE6GnsheDJ61HUfc7qOUVHs3b1JKuNY6lcdTAwMTPJhjpcdTAwMDFcdTAwMWEtaN0xW-Y0iafnsc1riLH0grL3XFyP8dGM-oG1XHUwMDFmcUKozbHKMHmAdLuO7qNALqJcdTAwMWRcdTAwMWLsbIKqy7Hxj5I5XHUwMDFlbpZcdTAwMWHDyi2AzLqo0sGHXHUwMDEylPan_SytqphcdTAwMDZ3a9BlW_FyW1t662B5nDBmQGfoRdgu8oBcdTAwMDWXU-aBK-GRMyvD26psrlwiXHUwMDEyLMBcdTAwMTIhrPdZIiwiZXhwIjoxNjY4MjI2NDM1LCJpc3MiOiJNSS1JTkZPU0VDIiwiYXVkIjoicXMubWl0dm9zLnR2LnhpYW9taS5zcnYiLCJjIjowLCJ0eXAiOiJjYXMifQ.nuXEDwkCemC9OfBZAYGHzzzPLxf1b91MSjCLjlh9RAk51YftP74O5_UNhdP-v-DtDdGoQXXguOiHgqFHl0Jn8g; user=liuyanan11; Jwt-Token=eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpc3MiOiIxMjcuMC4wLjE6OTAwMSIsImlhdCI6MTY2ODE0OTQ0MywiZXhwIjoxNjY4MTY3NDQzLCJ1c2VybmFtZSI6ImxpdXlhbmFuMTEifQ.yPqNHcF2N5pP7ffzAfs7avwF2Gs5Gb4uYHab8ruwB_c; sidebarStatus=0"
	method := "GET"
	url := "http://qs.mitvos.tv.xiaomi.srv/#/tvParameter/detail"

	cookies := request.MakeCookies(cookieStr)
	req, err := request.MakeRequest(method, url, "", nil, nil, cookies)
	if err != nil {
		fmt.Errorf("make request failed. error: %v", err)
		return
	}

	resp, err := request.DoWithTimeout(req, 10, 10240)
	if err != nil {
		fmt.Errorf("request failed. error: %v", err)
	}

	fmt.Println(string(resp))
}
