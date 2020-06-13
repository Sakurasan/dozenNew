package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"
)

var (
	Time2DozenNew = NewdozenNewkzz()
	Dozenchan     = make(chan []newkzz)
)

type dozenNewkzz map[string][]newkzz

func NewdozenNewkzz() dozenNewkzz {
	return make(dozenNewkzz)
}

func (t dozenNewkzz) AddNewKzz(n newkzz) {
	tt, _ := time.ParseInLocation("2006-01-02", n.STARTDATE[:10], time.Local)
	day := tt.Format("20060102")

	if _, ok := t[day]; ok {
		for _, kzz := range t[day] {
			if kzz.BONDCODE == n.BONDCODE {
				return
			}
		}
		t[day] = append(t[day], n)
	} else {
		t[day] = append(t[day], n)
	}

}

func (t dozenNewkzz) RenewDzenNewkzz() {
	today := time.Now().AddDate(0, 0, -1).Format("20060102")
	if _, ok := t[today]; ok {
		delete(t, today)
	}

}

type newkzz struct {
	BONDCODE  string `json:"BONDCODE"`  //债券代码
	SNAME     string `json:"SNAME"`     //债券简称
	STARTDATE string `json:"STARTDATE"` //申购日期
}

func Newnewkzz(a, b, c string) newkzz {
	return newkzz{BONDCODE: a, SNAME: b, STARTDATE: c}
}

type kzz []struct {
	BONDCODE          string      `json:"BONDCODE"` //债券代码
	TEXCH             string      `json:"TEXCH"`    //? "sh" : "sz"
	BCODE             string      `json:"BCODE"`
	SNAME             string      `json:"SNAME"`      //债券简称
	STARTDATE         string      `json:"STARTDATE"`  //申购日期
	CORRESCODE        string      `json:"CORRESCODE"` //申购代码
	CORRESNAME        string      `json:"CORRESNAME"`
	SWAPSCODE         string      `json:"SWAPSCODE"`         //正股代码
	SECURITYSHORTNAME string      `json:"SECURITYSHORTNAME"` //正股简称
	PARVALUE          float64     `json:"PARVALUE"`
	ISSUEPRICE        float64     `json:"ISSUEPRICE"`
	SWAPPRICE         float64     `json:"SWAPPRICE"`
	ZQHDATE           string      `json:"ZQHDATE"`  //中签号 发布日
	LUCKRATE          interface{} `json:"LUCKRATE"` //中签率(%)
	LISTDATE          string      `json:"LISTDATE"` //上市 时间
	DELISTDATE        string      `json:"DELISTDATE"`
	MEMO              string      `json:"MEMO"`
	ZGJ               string      `json:"ZGJ"`
	MarketType        string      `json:"MarketType"`
	HSCFJ             string      `json:"HSCFJ"`
	QSCFJ             string      `json:"QSCFJ"`
	PB                float64     `json:"PB"`
	GDYXSTARTDATE     string      `json:"GDYX_STARTDATE"` //股权登记日
	NOTICEDATESH      string      `json:"NOTICEDATE_SH"`
	STARTDATESH       string      `json:"STARTDATE_SH"`
	PRICESH           string      `json:"PRICE_SH"`
	RECORDATESH       string      `json:"RECORDATE_SH"`
	EXECREASONSH      string      `json:"EXECREASON_SH"`
	CALLSTYPESH       string      `json:"CALLSTYPE_SH"`
	NOTICEDATEHS      string      `json:"NOTICEDATE_HS"`
	STARTDATEHS       string      `json:"STARTDATE_HS"`
	ENDATEHS          string      `json:"ENDATE_HS"`
	PRICEHS           string      `json:"PRICE_HS"`
	EXECREASONHS      string      `json:"EXECREASON_HS"`
	LIMITBUYIPUB      string      `json:"LIMITBUYIPUB"`     //申购上限 (万元)
	ZGJHQ             string      `json:"ZGJ_HQ"`           //正股价
	ZGJZGJ            string      `json:"ZGJZGJ"`           //转股价
	ZGJZGJJZ          string      `json:"ZGJZGJJZ"`         //转股价值=100*正股价/转股价
	ZQNEW             string      `json:"ZQNEW"`            //债现价
	YJL               string      `json:"YJL"`              //转股溢价率=100%*债现价/转股价值-100%
	FSTPLACVALPERSTK  string      `json:"FSTPLACVALPERSTK"` //每股配售额<i title=\"股权登记日持股的原股东可以优先按照每股配售额优先配售可转债，单位：元/股
	AISSUEVOL         string      `json:"AISSUEVOL"`        //发行规模 (亿元)
}

var (
	tokenReg = regexp.MustCompile(`token=(.*)&c`)
	kzzUrl   = "http://data.eastmoney.com/kzz/"
	// token=70f12f2f4f091e459a279469fe49eca5
	kzzApi = "http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get?type=KZZ_LB2.0&token=70f12f2f4f091e459a279469fe49eca5&cmd=&st=STARTDATE&sr=-1&p=1&ps={1}"
	// &cmd=&st={sortType}&sr={sortRule}&p={page}&ps={pageSize}"
	// &cmd=&st=STARTDATE&sr=-1&p=1&ps={1}
)

func GetkzzUrlToken() string {
	res, _ := http.DefaultClient.Get(kzzUrl)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	token := tokenReg.FindString(string(body))
	fmt.Println(token)
	return token[6 : len(token)-2]
}

func Getkzz() kzz {
	res, _ := http.DefaultClient.Get(kzzApi)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	// fmt.Println(string(body))

	kzzs := kzz{}
	_ = json.Unmarshal(body, &kzzs)
	fmt.Println(len(kzzs))

	newkzzs := kzz{}
	for _, k := range kzzs {
		switch k.LUCKRATE.(type) {
		case string:
			if k.LUCKRATE.(string) == "-" {
				fmt.Println(k.BONDCODE, k.SNAME, k.STARTDATE[:10])
				newkzzs = append(newkzzs, k)
			}
		}
	}
	return newkzzs
}

func Publisher(c chan []newkzz) {
	k := Getkzz()
	for _, v := range k {
		Time2DozenNew.AddNewKzz(newkzz{BONDCODE: v.BONDCODE, SNAME: v.SNAME, STARTDATE: v.STARTDATE[:10]})
	}
	Time2DozenNew.RenewDzenNewkzz()

	if have, ok := Time2DozenNew[time.Now().Format("20060102")]; ok {
		c <- have
		// } else {
		// 	c <- []newkzz{newkzz{BONDCODE: "2333", SNAME: "测试布丁", STARTDATE: "2333"}}
	}
	return
}
