package task

import (
	"context"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var Buf []byte

func Login(username, password string, cookies *[]*network.Cookie) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate("https://account.xiaomi.com/fe/service/login/password?_locale=zh_CN"),
		chromedp.Sleep(time.Second * 2),
		chromedp.CaptureScreenshot(&Buf),
		chromedp.SendKeys("input[name=account]", username),
		chromedp.SendKeys("input[name=password]", password),
		chromedp.Sleep(time.Second * 1),
		chromedp.Click("input[type=checkbox]"),
		chromedp.Click("button[type=submit]", chromedp.NodeVisible),
		chromedp.CaptureScreenshot(&Buf),

		chromedp.ActionFunc(func(ctx context.Context) error {
			networkCookies, err := network.GetAllCookies().Do(ctx)
			if err != nil {
				return err
			}

			*cookies = networkCookies

			return nil
		}),
		chromedp.Sleep(time.Second * 3),
		chromedp.CaptureScreenshot(&Buf),
		chromedp.Sleep(time.Second * 3),
	}
}
