package authentication

type JwtAuth struct {
	Key                 string `json:"key"`
	Secret              string `json:"secret,omitempty"`
	PublicKey           string `json:"public_key,omitempty"`
	PrivateKey          string `json:"private_key,omitempty"`
	Algorithm           string `json:"algorithm,omitempty"`
	Exp                 int64  `json:"exp,omitempty"`
	Base64Secret        bool   `json:"base_64_secret,omitempty"`
	LifetimeGracePeriod int64  `json:"lifetime_grace_period,omitempty"`
}
