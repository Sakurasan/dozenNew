package main

import (
	"dozenNew/api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	notiUrl = ""

	TextTpl = "打新日历丨往死里打"
	DespTpl = `***
# 📅打新日历
{time}
{code}										
今日有新债，别忘了

{kzz}
{code}
</br></br>

***

> **往死里打 🐣**`
)

type dozenNewTpl struct {
	Text string `json:"text,omitempty"`
	Desp string `json:"desp,omitempty"`
}

func pushDozenNew(kzz string) bool {
	DespTpl = strings.Replace(DespTpl, "{time}", "## "+time.Now().AddDate(0, 0, -13).Format("2006-01-02"), -1)
	DespTpl = strings.Replace(DespTpl, "{kzz}", kzz, -1)
	mv := make(url.Values)
	mv.Add("text", TextTpl)
	mv.Add("desp", strings.Replace(DespTpl, "{code}", "```", -1))

	// params := fmt.Sprintf("text=%s&desp=%s", url.QueryEscape(TextTpl), url.QueryEscape(DespTpl))
	payload := strings.NewReader(mv.Encode())

	req, _ := http.NewRequest("POST", notiUrl, payload)

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("cache-control", "no-cache")

	// dumpreq, _ := httputil.DumpRequest(req, true)
	// fmt.Println(string(dumpreq))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println("通知请求失败")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	m := make(map[string]interface{})
	_ = json.Unmarshal(body, &m)

	if m["errno"] == 0 {
		return true
	}
	return false

}

func main() {
	if ftqq := os.Getenv("ftqq"); ftqq != "" {
		notiUrl = ftqq
	}

	go func() {
		api.Publisher(api.Dozenchan)
		for {
			now := time.Now()
			// 计算下一个零点
			next := now.Add(time.Hour * 24)
			// next := now.Add(time.Second * 30)
			next = time.Date(next.Year(), next.Month(), next.Day(), 12, 0, 0, 0, next.Location())
			fmt.Println(next.Sub(now))
			t := time.NewTimer(next.Sub(now))
			<-t.C
			api.Publisher(api.Dozenchan)
		}
	}()

	go func() {
		for {
			select {
			case c := <-api.Dozenchan:
				var kzz string
				for _, v := range c {
					kzz += fmt.Sprintln(v.BONDCODE, v.SNAME, v.STARTDATE)
				}
				notry := pushDozenNew(kzz)
				if !notry {
					for i := 0; i < 3; i++ {
						time.Sleep(5 * time.Minute)
						if ok := pushDozenNew(kzz); ok {
							break
						}

					}
				}
			}
		}
	}()
	select {}
}
