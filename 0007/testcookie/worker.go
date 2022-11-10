package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	task "testcookie/pkg/chrome"
	"testcookie/pkg/headless"
	"time"
)

func main() {
	ctx := context.Background()
	headlessBrowser, err := headless.New(ctx)
	if err != nil {
		fmt.Println("error:", err.Error())
		return
	}
	defer headlessBrowser.Close()

	cookies := make([]*network.Cookie, 0)
	loginCtx, loginCtxCancel := chromedp.NewContext(headlessBrowser.Context)
	defer loginCtxCancel()

	fmt.Println("lynbklk, here 1")
	go func() {
		for {
			time.Sleep(time.Second * 1)
			ioutil.WriteFile("./log.png", task.Buf, 0644)
		}
	}()

	err = chromedp.Run(loginCtx, task.Login("13366038505", "lyn9012241318", &cookies))
	if err != nil {
		fmt.Println("error:", err.Error())
	}
	cookiesJSON, _ := json.Marshal(cookies)
	fmt.Println("get cookies:", string(cookiesJSON))
}
