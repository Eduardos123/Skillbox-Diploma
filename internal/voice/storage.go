package voices

type VoiceCallData struct {
	Country             string
	Load                int
	ResponseTime        int
	Provider            string
	ConnectionStability float32
	PurityTTFB          int
	MedianCall          int
	Unknown             int
}
