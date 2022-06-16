package models

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var TsBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_bitrate_bytes",
		Help: "The overall TS bitrate based upon 188 byte TS packet.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPcrBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pcr_bitrate_bytes",
		Help: "The PCR bitrate of the TS based upon 188 byte TS packet.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPidBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_bitrate_bytes",
		Help: "The bitrate for an individual PID based upon a 188 byte TS packet.",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidServiceCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_service_count_total",
		Help: "The total number of services within the a given PID in the TS.",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidDiscontinuity = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_discontinuities_total",
		Help: "The discontinuities per PID.",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal", "description"},
)

var TsPidDuplicated = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_pid_duplicated_total",
		Help: "The number of duplicated PIDs seen for a given PID since the start of monitoring.",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal", "description"},
)

var TsServiceBitrate = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_service_bitrate_bytes",
		Help: "The bitrate of each service carried in the TS based upon a 188 byte TS packet.",
	},
	[]string{"multicast", "label", "service_id", "ts_id", "name", "provider", "type_name", "pcr_pid", "pmt_pid"},
)

var TsPidMinRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_min_repitition_ms",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPidMaxRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_max_repitition_ms",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPidMinRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_min_repitition_pkt",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPidMaxRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_max_repitition_pkt",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPidRepititionMs = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_repitition_ms",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPidRepititionPkt = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_repitition_pkt",
		Help: "",
	},
	[]string{"multicast", "label", "pid", "pid_hexadecimal"},
)

var TsPacketInvalidSync = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_invalid_sync_total",
		Help: "The number of invalid sync packets detected.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPacketSuspectIgnored = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_suspect_ignored_total",
		Help: "The number of suspect packets ignored.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPacketTeiErrors = promauto.NewCounterVec(
	prometheus.CounterOpts{
		Name: "ts_packet_tei_count_total",
		Help: "The number of transport errors detected.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPidCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_count_total",
		Help: "The total number of pids detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPcrPidCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_pcr_count_total",
		Help: "The total number of PCR pids detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsPidUnferencedCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_pid_unreferenced_count_total",
		Help: "The total number of unreferenced pids detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
)

// Future additions
/* var TsServiceClearCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_service_clear_count_total",
		Help: "The total number of clear services detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsServiceScrambledCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_service_scrambled_count_total",
		Help: "The total number of scrambled services detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
)

var TsServiceCount = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "ts_service_count_total",
		Help: "The total number of clear and scrambled services detected in the TS.",
	},
	[]string{"multicast", "label", "ts_id"},
) */
