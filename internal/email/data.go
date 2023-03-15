package emails
import (
	"os"
	"strconv"
	"strings"

	"github.com/DwarfWizzard/skillbox-diplom-go/internal/codes"
)

func GetData(path string) ([]EmailData, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	content := strings.Split(string(data), "\n")

	var result []EmailData
	for _, email := range content {
		s := strings.Split(email, ";")
		if len(s) != 3 || !codes.IsExist(s[0]) || !validateProvider(s[1]) {
			continue
		}

		deliveryTime, err := strconv.Atoi(s[2])
		if err != nil {
			continue
		}

		result = append(result, EmailData{
			Country:      s[0],
			Provider:     s[1],
			DeliveryTime: deliveryTime,
		})
	}

	return result, nil
}

func validateProvider(provider string) bool {
	return provider == "Gmail" || 
		   provider == "Yahoo" ||
		   provider == "Hotmail" ||
		   provider == "MSN" ||
		   provider == "Orange" ||
		   provider == "Comcast" ||
		   provider == "AOL" ||
		   provider == "Live" ||
		   provider == "RediffMail" ||
		   provider == "GMX" ||
		   provider == "Proton Mail" ||
		   provider == "Yandex" ||
		   provider == "Mail.ru"
}