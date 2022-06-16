package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"tsduck-prometheus/models"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func updatePidValues(multicast, label string, pid models.Pids) {
	// PID Bitrate
	models.TsPidBitrate.WithLabelValues(multicast, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Set(float64(pid.Bitrate))
	// PID Service Count
	models.TsPidServiceCount.WithLabelValues(multicast, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Set(float64(pid.SeviceCount))
	// PID Discontinuities Count
	if pid.Packets.Discontinuities >= 1 {
		sum := 0
		for sum < pid.Packets.Discontinuities {
			sum += 1
			models.TsPidDiscontinuity.WithLabelValues(multicast, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}
	// Duplicate PID Count
	if pid.Packets.Duplicated >= 1 {
		sum := 0
		for sum < pid.Packets.Duplicated {
			sum += 1
			models.TsPidDuplicated.WithLabelValues(multicast, label, fmt.Sprint(pid.Id), strconv.FormatInt(int64(pid.Id), 16), pid.Description).Inc()
		}
	}
}

func updateServiceValues(multicast, label string, service models.Services) {
	// TS bitrate per service
	models.TsServiceBitrate.WithLabelValues(multicast, label, strconv.Itoa(service.Id), strconv.Itoa(service.TsId), service.Name, service.Provider, service.TypeName, strconv.Itoa(service.PcrPid), strconv.Itoa(service.PmtPid)).Set(float64(service.Bitrate))
}

func updateTsValues(multicast, label string, ts models.Ts) {
	// Total TS Bitrate (188byte/pkt)
	models.TsBitrate.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Set(float64(ts.Bitrate))
	// PCR TS Bitrate (188byte/pkt)
	models.TsPcrBitrate.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Set(float64(ts.PcrBitrate))
	// TS Packet Invalid Sync Count
	if ts.Packets.InvalidSyncs >= 1 {
		sum := 0
		for sum < ts.Packets.InvalidSyncs {
			sum += 1
			models.TsPacketInvalidSync.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Packet Suspect Ignored Count
	if ts.Packets.SuspectIgnored >= 1 {
		sum := 0
		for sum < ts.Packets.SuspectIgnored {
			sum += 1
			models.TsPacketSuspectIgnored.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Packet Transport Error Count
	if ts.Packets.TransportErrors >= 1 {
		sum := 0
		for sum < ts.Packets.TransportErrors {
			sum += 1
			models.TsPacketTeiErrors.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Inc()
		}
	}
	// TS Total PID Count
	models.TsPidCount.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Total))
	// TS PCR PID Count
	models.TsPcrPidCount.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Pcr))
	// TS Unreferenced PID Count
	models.TsPidUnferencedCount.WithLabelValues(multicast, label, strconv.Itoa(ts.Id)).Set(float64(ts.Pids.Unreferenced))
	// TsServiceClearCount (future)
	// TsServiceScrambledCount (future)
	// TsServiceCount (future)
}

func updateTableValues(multicast, label string, tables models.Tables) {
	// PID Max Repitition MS
	models.TsPidMaxRepititionMs.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MaxRepititionMs))
	// PID Max Repitition Pkt
	models.TsPidMaxRepititionPkt.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MaxRepititionPkt))
	// PID Min Repitition MS
	models.TsPidMinRepititionMs.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MinRepititionMs))
	// PID In Repitition Pkt
	models.TsPidMinRepititionPkt.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.MinRepititionPkt))
	// PID Repitition MS
	models.TsPidRepititionMs.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.RepetitionMs))
	// PID Repitition Pkt
	models.TsPidRepititionPkt.WithLabelValues(multicast, label, fmt.Sprint(tables.Pid), strconv.FormatInt(int64(tables.Pid), 16)).Set(float64(tables.RepetitionPkt))
}

func readTspOut(multicast, label string, scanner *bufio.Scanner) {
	var tsp models.Tsp

	for scanner.Scan() {
		s := scanner.Text()
		// Hacky way of removing the beginning of the JSON line which prepends TSDucks JSON output
		t := strings.Replace(s, "* analyze: ", "", -1)
		// Unmarshal JSON into TSP model
		json.Unmarshal([]byte(t), &tsp)

		// Update PID metrics
		for _, pid := range tsp.Pids {
			go updatePidValues(multicast, label, pid)
		}

		// Update Service metrics
		for _, service := range tsp.Services {
			go updateServiceValues(multicast, label, service)
		}

		// Update TS metrics
		go updateTsValues(multicast, label, tsp.Ts)

		// Update TS table metrics
		for _, table := range tsp.Tables {
			go updateTableValues(multicast, label, table)
		}
	}
}

func launchTsp(multicast, source_address, label string) {
	// Launch TSP
	cmd := exec.Command("tsp", "-I", "ip", multicast, "-s", source_address, "-P", "analyze", "-i", "1", "--json-line", "-O", "drop")
	// TSDuck outputs to stderr
	cmdReader, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Started monitoring for %v\n", label)
	}

	// Create buffer to read the stderr output
	scanner := bufio.NewScanner(cmdReader)

	go func() {
		// Read Output Buffer
		readTspOut(multicast, label, scanner)
	}()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Parse Args
	args := os.Args[1:]
	if len(args) < 1 {
		log.Fatal("Not enough arguments! Plase parse at least one input e.g. 225.0.0.1:20000,172.0.0.1,My_Service")
	}

	for _, item := range args {
		s := strings.Split(item, ",")
		if len(s) < 3 {
			log.Fatal("Not enough arguments! Required format is multicast:port,source,label, e.g. 225.0.0.1:20000,172.0.0.1,My_Service")
		}
		// Launch TSDuck (tsp subprocess)
		go launchTsp(s[0], s[1], s[2])
	}

	r := prometheus.NewRegistry()
	// Register prometheus metrics
	r.MustRegister(models.TsBitrate)
	r.MustRegister(models.TsPcrBitrate)
	r.MustRegister(models.TsPidBitrate)
	r.MustRegister(models.TsPidServiceCount)
	r.MustRegister(models.TsPidDiscontinuity)
	r.MustRegister(models.TsPidDuplicated)
	r.MustRegister(models.TsServiceBitrate)
	r.MustRegister(models.TsPidMaxRepititionMs)
	r.MustRegister(models.TsPidMaxRepititionPkt)
	r.MustRegister(models.TsPidMinRepititionMs)
	r.MustRegister(models.TsPidMinRepititionPkt)
	r.MustRegister(models.TsPidRepititionMs)
	r.MustRegister(models.TsPidRepititionPkt)
	r.MustRegister(models.TsPacketInvalidSync)
	r.MustRegister(models.TsPacketSuspectIgnored)
	r.MustRegister(models.TsPacketTeiErrors)
	r.MustRegister(models.TsPidCount)
	r.MustRegister(models.TsPcrPidCount)
	r.MustRegister(models.TsPidUnferencedCount)

	handler := promhttp.HandlerFor(r, promhttp.HandlerOpts{})

	http.Handle("/metrics", handler)
	http.ListenAndServe(":8000", nil)
}
