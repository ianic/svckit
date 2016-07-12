package msgs

// Poruka za prijavu na push notifikacije
type PushNotSubscribe struct {
	// remember_token igraca koji se prijavljuje za notifikacije
	IgracId string `json:"igrac_id"`
	// Uredjaj s kojim se prijavljuje za notifikacije
	Uredjaj struct {
		// Google Cloud Messaging device id 
		GcmId   string `json:"gcm_id"`
		// Firebase Cloud Messaging device id 
		FcmId   string `json:"fcm_id"`
		// Apple Cloud Messaging device id 
		AppleId string `json:"apple_id"`
		// Da li je uredjaj aktivan, ako je false koristi se za deaktivaciju uredjaja
		Aktivan bool   `json:"aktivan"`
	} `json:"uredjaj"`
	// Notifikacije na koje se igrac pretplacuje
	Pretplate struct {
		PrivatnePoruke  bool `json:"privatne_poruke"`
		Novosti         bool `json:"novosti"`
		ListicVrednovan bool `json:"listic_vrednovan"`
	} `json:"pretplate"`
}
