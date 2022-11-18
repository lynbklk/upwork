package task

import (
	"bufio"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"os"
	"time"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

var res string

func VisitWeb(url string, cookies []*network.Cookie) chromedp.Tasks {
	return chromedp.Tasks{
		//ActionFunc是一个适配器，允许使用普通函数作为操作。
		chromedp.ActionFunc(func(ctx context.Context) error {
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			for _, cookie := range cookies {
				err := network.SetCookie(cookie.Name, cookie.Value).
					WithExpires(&expr).
					WithDomain("http://qs.mitvos.tv.xiaomi.srv/").
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return err
				}
			}
			return nil
		}),
		chromedp.Navigate(url),
	}
}

func DoCrawler() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Sleep(1 * time.Second),
		chromedp.WaitVisible("#app > div > div.main-container.hasTagsView > section > div > div.el-table.el-table--fit.el-table--border.el-table--scrollable-x.el-table--enable-row-hover.el-table--enable-row-transition.el-table--medium > div.el-table__body-wrapper.is-scrolling-left > table"),
		chromedp.Sleep(1 * time.Second),
		chromedp.OuterHTML(`tbody`, &res, chromedp.ByQuery),
	}
}

func Next() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click("#app > div > div.main-container.hasTagsView > section > div > div.pagination-container > div > button.btn-next", chromedp.ByQuery),
	}
}

func WriteTXT(txt string) {
	f, err := os.OpenFile("1.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Println("os Create error: ", err)
		return
	}
	defer f.Close()

	bw := bufio.NewWriter(f)
	bw.WriteString(txt + "\n")
	bw.Flush()
}
