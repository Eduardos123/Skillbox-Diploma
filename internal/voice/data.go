package voices

import (
	"os"
	"strconv"
	"strings"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/codes"
)

func GetData(path string) ([]VoiceCallData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := strings.Split(string(data), "\n")

	var result []VoiceCallData
	for _, voice := range content {
		s := strings.Split(voice, ";")
		if len(s) != 8 || !codes.IsExist(s[0]) ||
			(s[3] != "TransparentCalls" && s[3] != "E-Voice" && s[3] != "JustPhone") {
			continue
		}

		country := s[0]
		provider := s[3]


		load, err := strconv.Atoi(s[1])
		if err != nil {
			continue
		}

		responseTime, err := strconv.Atoi(s[2])
		if err != nil {
			continue
		}

		connStability, err := strconv.ParseFloat(s[4], 32)
		if err != nil {
			continue
		}

		purityTTFB, err := strconv.Atoi(s[5])
		if err != nil {
			continue
		}

		medianCall, err := strconv.Atoi(s[6])
		if err != nil {
			continue
		}

		unknown, err := strconv.Atoi(s[7])
		if err != nil {
			continue
		}

		result = append(result, VoiceCallData{
			Country: country,
			Load: load,
			ResponseTime: responseTime,
			Provider: provider,
			ConnectionStability: float32(connStability),
			PurityTTFB: purityTTFB,
			MedianCall: medianCall,
			Unknown: unknown,
		})
	}

	return result, nil
}