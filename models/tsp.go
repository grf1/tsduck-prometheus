package models

type Tsp struct {
	Pids     []Pids     `json:"pids"`
	Services []Services `json:"services"`
	Tables   []Tables   `json:"tables"`
	Time     Time       `json:"time"`
	Ts       Ts         `json:"ts"`
}

type Pids struct {
	Audio            bool       `json:"audio"`
	Bitrate          int        `json:"bitrate"`
	Bitrate204       int        `json:"bitrate-204"`
	Description      string     `json:"description"`
	ECM              bool       `json:"ecm"`
	EMM              bool       `json:"emm"`
	Global           bool       `json:"global"`
	Id               int        `json:"id"`
	InvalidPesPrefix bool       `json:"invalid-pes-prefix"`
	IsScrambled      bool       `json:"is-scrambled"`
	Packets          pidPackets `json:"packets"`
	PES              int        `json:"pes"`
	PESStreamId      int        `json:"pes-stream-id"`
	PMT              bool       `json:"pmt"`
	SeviceCount      int        `json:"service-count"`
	Services         []int      `json:"services"`
	T2MI             bool       `json:"t2mi"`
	UnitStart        int        `json:"unit-start"`
	Unreferenced     bool       `json:"unreferenced"`
	Video            bool       `json:"video"`
}

type pidPackets struct {
	AF                int `json:"af"`
	Clear             int `json:"clear"`
	Discontinuities   int `json:"discontinuities"` //
	Duplicated        int `json:"duplicated"`      //
	InvalidScrambling int `json:"invalid-scrambling"`
	PCR               int `json:"pcr"`
	Scrambled         int `json:"scrambled"`
	Total             int `json:"total"`
}

type Services struct {
	Bitrate           int               `json:"bitrate"`
	Bitrate204        int               `json:"bitrate-204"`
	Components        serviceComponents `json:"components"`
	Id                int               `json:"id"`
	IsScrambled       bool              `json:"is-scrambled"`
	Name              string            `json:"name"`
	OriginalNetworkId int               `json:"original-network-id"`
	Packets           int               `json:"packets"`
	PcrPid            int               `json:"pcr-pid"`
	Pids              []int             `json:"pids"`
	PmtPid            int               `json:"pmt-pid"`
	Provider          string            `json:"provider"`
	SSU               bool              `json:"ssu"`
	T2MI              bool              `json:"t2mi"`
	TsId              int               `json:"tsid"`
	Type              int               `json:"type"`
	TypeName          string            `json:"string"`
}

type serviceComponents struct {
	Clear     int `json:"clear"`
	Scrambled int `json:"scrambled"`
	Total     int `json:"total"`
}

type Tables struct {
	FirstVersion     int   `json:"first-version"`
	LastVersion      int   `json:"last-version"`
	MaxRepititionMs  int   `json:"max-repetition-ms"`  //
	MaxRepititionPkt int   `json:"max-repetition-pkt"` //
	MinRepititionMs  int   `json:"min-repetition-ms"`  //
	MinRepititionPkt int   `json:"min-repetition-pkt"` //
	Pid              int   `json:"pid"`
	RepetitionMs     int   `json:"repetition-ms"`  //
	RepetitionPkt    int   `json:"repetition-pkt"` //
	Sections         int   `json:"sections"`
	Tables           int   `json:"tables"`
	Tid              int   `json:"tid"`
	TidExt           int   `json:"tid-ext"`
	Versions         []int `json:"versions"`
}

type Time struct {
	Local timeLocale `json:"local"`
	UTC   timeLocale `json:"utc"`
}

type timeLocale struct {
	System timeDatum `json:"system"`
}

type timeDatum struct {
	First timeFormat `json:"first"`
	Last  timeFormat `json:"last"`
}

type timeFormat struct {
	Date             string `json:"date"`
	SecondsSince2000 int    `json:"seconds-since-2000"`
	Time             string `json:"time"`
}

type Ts struct {
	Bitrate        int        `json:"bitrate"`
	Bitrate204     int        `json:"bitrate-204"`
	Bytes          int        `json:"bytes"`
	Duration       int        `json:"duration"`
	Id             int        `json:"id"`
	Packets        tsPackets  `json:"packets"`
	PcrBitrate     int        `json:"pcr-bitrate"`
	PcrBitrate204  int        `json:"pcr-bitrate204"`
	Pids           tsPids     `json:"pids"`
	Services       tsServices `json:"services"`
	UserBitrate    int        `json:"user-bitrate"`
	UserBitrate204 int        `json:"user-bitrate-204"`
}

type tsPackets struct {
	InvalidSyncs    int `json:"invalid-syncs"`
	SuspectIgnored  int `json:"suspect-ignored"`
	Total           int `json:"total"`
	TransportErrors int `json:"transport-errors"`
}

type tsPids struct {
	Clear        int          `json:"clear"`
	Global       tsPidsGlobal `json:"global"`
	Pcr          int          `json:"pcr"`
	Scrambled    int          `json:"scrambled"`
	Total        int          `json:"total"`
	Unreferenced int          `json:"unreferenced"`
}

type tsPidsGlobal struct {
	Bitrate     int   `json:"bitrate"`
	Bitrate204  int   `json:"bitrate-204"`
	Clear       int   `json:"clear"`
	IsScrambled bool  `json:"is-scrambled"`
	Packets     int   `json:"packets"`
	Pids        []int `json:"pids"`
	Scrambled   int   `json:"scrambled"`
}

type tsServices struct {
	Clear     int `json:"clear"`
	Scrambled int `json:"scrambled"`
	Total     int `json:"total"`
}
