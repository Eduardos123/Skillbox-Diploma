package service

import (
	"net/http"
	"sort"
	"strings"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/billing"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/codes"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/email"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/incedent"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/mms"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/sms"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/support"
	"github.com/DwarfWizzard/skillbox-diplom-go/internal/voice"
	"github.com/gin-gonic/gin"
)

type ResultSetT struct {
	SMS       [][]smss.SMSData                `json:"sms"`
	MMS       [][]mmss.MMSData                `json:"mms"`
	VoiceCall []voices.VoiceCallData          `json:"voice_call"`
	Email     map[string][][]emails.EmailData `json:"email"`
	Billing   billings.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Incidents []incedents.IncidentData        `json:"incident"`
}

type ResultT struct {
	Status bool       `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string     `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type service struct {
	pathToData string
	uri        string
}

func New(pathToData, uri string) *service {
	return &service{
		pathToData: pathToData,
		uri:        uri,
	}
}

func (s *service) HandleConnection(c *gin.Context) {
	result := s.getResultData()
	c.JSON(http.StatusOK, result)
}

func (s *service) getResultData() ResultT{
	result := ResultT{Status: true}

	var collectDataErrors []string
	smsData, err := smss.GetData(s.pathToData+"/sms.data")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	mmsData, err := mmss.GetData(s.uri+"/mms")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	voiceData, err := voices.GetData(s.pathToData+"/voice.data")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	emailData, err := emails.GetData(s.pathToData+"/email.data")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	billing, err := billings.GetData(s.pathToData+"/billing.data")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	supportData, err := supports.GetData(s.uri+"/support")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}
	incedentData, err := incedents.GetData(s.uri+"/accendent")
	if err != nil {
		collectDataErrors = append(collectDataErrors, err.Error())
	}

	if len(collectDataErrors) > 0 {
		errMessage := strings.Join(collectDataErrors, ";")

		result.Status = false
		result.Data = nil
		result.Error = errMessage
	}

	resultSet := &ResultSetT{}

	resultSet.SMS = prepareSMSS(smsData)

	resultSet.MMS = prepareMMSS(mmsData)

	resultSet.VoiceCall = voiceData

	resultSet.Email = prepareEmails(emailData)

	resultSet.Billing = billing

	resultSet.Support = prepareSupport(supportData)

	resultSet.Incidents = prepareIncedents(incedentData)

	result.Data = resultSet

	return result
}

func prepareSMSS(data []smss.SMSData) [][]smss.SMSData {
	var result [][]smss.SMSData

	for i, v := range data {
		data[i].Country = codes.GetName(v.Country)
	}

	smsDataSortByCountry := make([]smss.SMSData, len(data))
	smsDataSortByProvider := make([]smss.SMSData, len(data))

	copy(smsDataSortByCountry, data)
	copy(smsDataSortByProvider, data)

	sort.Slice(smsDataSortByCountry, func(i, j int) bool {
		return smsDataSortByCountry[i].Country < smsDataSortByCountry[j].Country
	})

	sort.Slice(smsDataSortByProvider, func(i, j int) bool {
		return smsDataSortByProvider[i].Provider < smsDataSortByProvider[j].Provider
	})

	result = append(result, smsDataSortByProvider, smsDataSortByCountry)

	return result
}

func prepareMMSS(data []mmss.MMSData) [][]mmss.MMSData {
	var result [][]mmss.MMSData

	for i, v := range data {
		data[i].Country = codes.GetName(v.Country)
	}

	mmsDataSortByCountry := make([]mmss.MMSData, len(data))
	mmsDataSortByProvider := make([]mmss.MMSData, len(data))

	copy(mmsDataSortByCountry, data)
	copy(mmsDataSortByProvider, data)
	sort.Slice(mmsDataSortByCountry, func(i, j int) bool {
		return mmsDataSortByCountry[i].Country < mmsDataSortByCountry[j].Country
	})

	sort.Slice(mmsDataSortByProvider, func(i, j int) bool {
		return mmsDataSortByProvider[i].Provider < mmsDataSortByProvider[j].Provider
	})

	result = append(result, mmsDataSortByProvider, mmsDataSortByCountry)

	return result
}

func prepareEmails(data []emails.EmailData) map[string][][]emails.EmailData {
	resultMap := make(map[string][][]emails.EmailData)

	emailDataByCountry := make(map[string][]emails.EmailData)
	for _, email := range data {
		emailDataByCountry[email.Country] = append(emailDataByCountry[email.Country], email)
	}

	for country, emails := range emailDataByCountry {
		sort.Slice(emails, func(i, j int) bool {
			return emails[i].DeliveryTime < emails[j].DeliveryTime
		})

		resultMap[country] = append(resultMap[country], emails[0:3], emails[len(emails)-4:])
	}

	return resultMap
}

func prepareSupport(data []supports.SupportData) []int {
	var tickets int
	for _, topic := range data {
		tickets += topic.ActiveTickets
	}

	var load int
	if tickets > 16 {
		load = 3
	} else if tickets > 9 {
		load = 2
	} else {
		load = 1
	}

	respTime := (60/18)*tickets

	return []int{load, respTime}

}

func prepareIncedents(data []incedents.IncidentData) []incedents.IncidentData {
	var result []incedents.IncidentData

	for _, incedent := range data {
		if incedent.Status == "active" {
			result = append(result, incedent)
		}
	}

	for _, incedent := range data {
		if incedent.Status == "closed" {
			result = append(result, incedent)
		}
	}

	return result
}