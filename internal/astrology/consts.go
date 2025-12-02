package astrology

import (
	"database/sql"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type status int

const (
	isSingle status = iota
	inRelationship
	isMarried
)

const (
	dateForRecovery = "01/01/6969"
)

type Player struct {
	status     status
	zodiacSign *widget.Label
	dob        string
}

type config struct {
	DB              *sql.DB
	Path            string
	PathDB          string
	RickyWall       *[]byte
	RickyAudioBytes *[]byte
	Window          fyne.Window
	LogsCh          chan string
	RestartCh       chan string
}

var (
	cfg               config
	player            Player
	navigationButtons *fyne.Container
)

// Roasting lines
var Roasts = map[string]map[status]string{
	"aries": {
		isSingle:       "Single because dating you feels like babysitting a lit fuse — exhausting and unsafe.",
		inRelationship: "In a relationship but fights like their hobby is starting wars and losing medals.",
		isMarried:      "Married yet still screaming like a toddler who found a microphone and no shame.",
	},
	"taurus": {
		isSingle:       "Single because you move like molasses and charm left the building years ago.",
		inRelationship: "In a relationship but stubborn enough to turn love into a daily hostage situation.",
		isMarried:      "Married but so rooted in your nonsense your partner files complaints to the plants.",
	},
	"gemini": {
		isSingle:       "Single because dating you is like subscribing to chaos with unexpected subscription fees.",
		inRelationship: "In a relationship but shows up emotionally on random days and ghosts the rest.",
		isMarried:      "Married but your spouse gets a new person each morning and customer service is closed.",
	},
	"cancer": {
		isSingle:       "Single because you cry through the whole trailer and ruin the movie for everyone.",
		inRelationship: "In a relationship but weaponizes tears like it's a full-time income source.",
		isMarried:      "Married yet constantly offended — even the air gets accused of betrayal.",
	},
	"leo": {
		isSingle:       "Single because your ego gobbles the room and leaves nothing for anyone else.",
		inRelationship: "In a relationship but needs applause for breathing; exhausting, narcissistic, permanent show.",
		isMarried:      "Married but treats the vows like a solo performance nobody asked to watch.",
	},
	"virgo": {
		isSingle:       "Single because you critique dating profiles like you’re handing out failing grades.",
		inRelationship: "In a relationship but corrects love like it's a homework assignment you won't sign off on.",
		isMarried:      "Married but nags so much you could charge rent for the toxic atmosphere.",
	},
	"libra": {
		isSingle:       "Single because you take forever to choose and everyone left hungry for an answer.",
		inRelationship: "In a relationship but indecisive so fights end in awkward silence and confusion.",
		isMarried:      "Married but still unsure if yes meant yes or an accident — regret pending.",
	},
	"scorpio": {
		isSingle:       "Single because your 'mystery' is just moodiness with a side of paranoia.",
		inRelationship: "In a relationship but jealous of shadows and rumors and things that never existed.",
		isMarried:      "Married but broods like revenge is a hobby and trust is a myth.",
	},
	"sagittarius": {
		isSingle:       "Single because you sprint at the first hint of 'forever' like it’s a forest fire.",
		inRelationship: "In a relationship but treats loyalty like a prank you forgot to stop playing.",
		isMarried:      "Married but planning exits like they collect frequent flyer miles for running away.",
	},
	"capricorn": {
		isSingle:       "Single because your heart is locked in a vault and you lost the key years ago.",
		inRelationship: "In a relationship but turns romance into a lecture — boring and soul-sucking.",
		isMarried:      "Married but colder than a closed office on a holiday; affection sold separately.",
	},
	"aquarius": {
		isSingle:       "Single because your 'unique' vibe feels like someone forgot to human you properly.",
		inRelationship: "In a relationship but emotionally unplugged — partner wonders if they live with a concept.",
		isMarried:      "Married but lives in fantasy land where everyone tolerates your weird like it’s charming.",
	},
	"pisces": {
		isSingle:       "Single because you're busy dating clouds while the rest of us live on land.",
		inRelationship: "In a relationship but drifts away so often your partner needs a search party.",
		isMarried:      "Married but half your marriage is a soap opera in your head where you star and suffer.",
	},
}
