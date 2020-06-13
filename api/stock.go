package api

var gpApi = "http://dcfm.eastmoney.com/em_mutisvcexpandinterface/api/js/get?type=XGSG_LB&token=70f12f2f4f091e459a279469fe49eca5&st={purchasedate,securitycode}&sr={sortRule}&p={page}&ps={pageSize}"

type newstock struct {
	Securityshortname string `json:"securityshortname"` //股票简称
	Securitycode      string `json:"securitycode"`      //股票代码
	Purchasedate      string `json:"purchasedate"`      // 申购日期
}

type gupiao []struct {
	Financecode       string  `json:"financecode"`
	Companycode       string  `json:"companycode"`
	Securityshortname string  `json:"securityshortname"` //股票简称
	Securitycode      string  `json:"securitycode"`      //股票代码
	Subcode           string  `json:"subcode"`           // 申购代码
	Fxzl              float64 `json:"fxzl"`              // 发行总数（万股）
	Wsfxsl            float64 `json:"wsfxsl"`            // 网上发行（万股）
	Applyont          float64 `json:"applyont"`          // 申购上限（万股）
	IssuepriceMoney   float64 `json:"issuepriceMoney"`
	Issueprice        float64 `json:"issueprice"`   // 发行价格
	Purchasedate      string  `json:"purchasedate"` // 申购日期
	Lwrandate         string  `json:"lwrandate"`    // 中签号公布日
	Listingdate       string  `json:"listingdate"`  // 上市日期
	Peissuea          float64 `json:"peissuea"`     // 发行市盈率
	Lwr               float64 `json:"lwr"`          // 中签率（％）
	Cbxjrgbs          string  `json:"cbxjrgbs"`     // 询价累计报价倍数
	Sc                string  `json:"sc"`
	Mzyqgs            float64 `json:"mzyqgs"`
	Sgzs              string  `json:"sgzs"`
	ApplyontMoney     float64 `json:"applyontMoney"` // 顶格申购所需配市值（万元）
	Averagelow        string  `json:"averagelow"`
	Kb                string  `json:"kb"` //开板？待上市
	Zzf               float64 `json:"zzf"`
	Sl                float64 `json:"sl"`
	Mzyqhl            float64 `json:"mzyqhl"`
	Totaliiqrplaceoff string  `json:"totaliiqrplaceoff"`
	Jg1               string  `json:"jg1"`
	Jg2               string  `json:"jg2"`
	Jg3               string  `json:"jg3"`
	Pe1               string  `json:"pe1"`
	Pe2               string  `json:"pe2"`
	Pe3               string  `json:"pe3"`
	Bkpe              string  `json:"bkpe"`
	INDUSTRY          string  `json:"INDUSTRY"`
	Close             float64 `json:"Close"` // 最新价
	ChangePercent     float64 `json:"ChangePercent"`
	ZgsmsOrYxs        string  `json:"ZgsmsOrYxs"`
	URL               string  `json:"Url"`
	INDUSTRYPE        string  `json:"INDUSTRYPE"` //行业市盈率
	NewPrice          float64 `json:"newPrice"`   // 首日收盘价
	Wszqjkr           string  `json:"wszqjkr"`    // 中签缴款日期
	MAINBUSIN         string  `json:"MAINBUSIN"`
	Ycwssgsx          float64 `json:"ycwssgsx"`
	Ycwssgzj          float64 `json:"ycwssgzj"`
	Update            string  `json:"Update"`
	Sgrqrow           float64 `json:"sgrqrow"`
}
