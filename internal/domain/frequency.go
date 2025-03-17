package domain

type Frequency string

const (
	Daily   Frequency = "daily"
	Weekly  Frequency = "weekly"
	Monthly Frequency = "monthly"
)

func IsValidFrequency(freq string) bool {
	return freq == string(Daily) || freq == string(Weekly) || freq == string(Monthly)
}
