package utils

import (
	"chill/internal/log"
	"chill/internal/sms/sms"
	"strings"

	"github.com/alibabacloud-go/tea/tea"

	dyvmsapiclient "github.com/alibabacloud-go/dyvmsapi-20170525/client"
	rpc "github.com/alibabacloud-go/tea-rpc/client"
)

func SendSms(tencentMsg *sms.TencentSms, types string, phones []string, tplID int64) {
	var tels []sms.SmsTel
	for _, nPhone := range phones {
		var nationCode, phone string
		tmpPhone := strings.Split(nPhone, "-")
		if len(tmpPhone) == 1 {
			phone = tmpPhone[0]
		} else if len(tmpPhone) > 1 {
			nationCode = tmpPhone[0]
			phone = tmpPhone[1]
		} else {
			log.Errorw("电话解析失败", "用户", nPhone)
			return
		}
		tels = append(tels, sms.SmsTel{
			Mobile:     phone,
			Nationcode: nationCode,
		})
	}
	p := &sms.MultiParams{Params: []string{types}, Tel: tels, TplId: tplID}
	result, err := tencentMsg.GetSendMultiSms().Fetch(p)
	if err != nil {
		log.Errorw("发短信失败", "用户", phones, "error", err.Error())
	}

	for _, item := range result.Detail {
		if item.Errmsg != "OK" {
			log.Errorw("发短信失败", "电话", item.Mobile, "返回结果:", result)
		} else {
			log.Infow("发短信成功", "电话", item.Mobile)
		}
	}
}

func SendPhone(aliPhone *dyvmsapiclient.Client, msg string, phones []string,
	phoneShowNumber, phoneTtsCode string) {

	for _, nPhone := range phones {
		var nationCode, phone string
		tmpPhone := strings.Split(nPhone, "-")
		if len(tmpPhone) == 1 {
			phone = tmpPhone[0]
		} else if len(tmpPhone) > 1 {
			nationCode = tmpPhone[0]
			phone = tmpPhone[1]
		} else {
			log.Errorw("电话解析失败", "用户", nPhone)
			return
		}
		request := &dyvmsapiclient.SingleCallByTtsRequest{
			// 被叫显号，若您使用的模板为公共号池号码外呼模板，则该字段值必须为空；
			// 若您使用的模板为专属号码外呼模板，则必须传入已购买的号码，仅支持一个号码，您可以在语音服务控制台上查看已购买的号码。
			CalledShowNumber: tea.String(phoneShowNumber),
			// 被叫号码。仅支持中国内地号码。一次请求仅支持一个被叫号。
			CalledNumber: tea.String(nationCode + phone),
			// 语音文件的语音ID。
			TtsCode:  tea.String(phoneTtsCode),
			TtsParam: tea.String("{\"node\":\"" + msg + "\"}"),
		}
		response, _err := aliPhone.SingleCallByTts(request)
		if _err != nil {
			log.Errorw("打电话失败", "用户:", phone, "error", _err.Error())
			continue
		}
		if *response.Code != "OK" {
			log.Errorw("打电话失败", "返回数据", response, "用户", phone)
			continue
		}
		log.Infow("打电话成功", "返回数据:", response, "用户:", phone)
	}
}

func CreatePhoneClient(accessKeyId *string, accessKeySecret *string, regionId *string) (_result *dyvmsapiclient.Client, _err error) {
	config := &rpc.Config{}
	// 您的AccessKey ID
	config.AccessKeyId = accessKeyId
	// 您的AccessKey Secret
	config.AccessKeySecret = accessKeySecret
	// 您的可用区ID
	config.RegionId = regionId
	_result = &dyvmsapiclient.Client{}
	_result, _err = dyvmsapiclient.NewClient(config)
	return _result, _err
}
