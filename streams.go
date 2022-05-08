package main

import (
//	"database/sql"
	"errors"
	"log"
	"time"
	"github.com/deepch/vdk/format/rtspv2"
//	_ "github.com/go-sql-driver/mysql"
)

var (
	ErrorStreamExitNoVideoOnStream = errors.New("Stream Exit No Video On Stream")
	ErrorStreamExitRtspDisconnect  = errors.New("Stream Exit Rtsp Disconnect")
	ErrorStreamExitNoViewer        = errors.New("Stream Exit On Demand No Viewer")
)


func serveStreams() { 
	
	// db, err := sql.Open("mysql", "mulelocal:!Y4dfg@wer@tcp(127.0.0.1:3307)/sdcd_app")
	// if err != nil {
	// 	log.Fatal("Unable to open connection to db")
	// }
	// defer db.Close()

	// results, err := db.Query("select * from device")
	// if err != nil {
	// 	log.Fatal("Error when fetching device table rows:", err)
	// }
	// defer results.Close()
	// devices := make([]*Device, 0)
	// for results.Next() {
	// 	device :=new(Device)
	// 	if err := results.Scan(&device.id, &device.name, &device.ip,  &device.onDemand, &device.protocol, &device.port, &device.username, &device.password); 
	// 	err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	devices = append(devices, device)

	// }

	// for _, device := range devices {
	// 	go RTSPWorkerLoop(device.name, device.protocol + "://" + device.username + ":" + device.password + "@" + device.ip + ":" + device.port, device.onDemand)
	// }

	for k, v := range Config.Streams {
		if !v.OnDemand {
			go RTSPWorkerLoop(k, v.URL, v.OnDemand)
			log.Fatal(k);
		}
	}

}
func RTSPWorkerLoop(name, url string, OnDemand bool) { 
	defer Config.RunUnlock(name)
	for {
		log.Println("Stream Try Connect", name)
		err := RTSPWorker(name, url, OnDemand)
		if err != nil {
			log.Println(err)
		}
		if OnDemand && !Config.HasViewer(name) {
			log.Println(ErrorStreamExitNoViewer)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
func RTSPWorker(name, url string, OnDemand bool) error {
	keyTest := time.NewTimer(20 * time.Second)
	clientTest := time.NewTimer(20 * time.Second)
	RTSPClient, err := rtspv2.Dial(rtspv2.RTSPClientOptions{URL: url, DisableAudio: false, DialTimeout: 3 * time.Second, ReadWriteTimeout: 3 * time.Second, Debug: false})
	if err != nil {
		return err
	}
	defer RTSPClient.Close()
	if RTSPClient.CodecData != nil {
		Config.coAd(name, RTSPClient.CodecData)
	}
	var AudioOnly bool
	if len(RTSPClient.CodecData) == 1 && RTSPClient.CodecData[0].Type().IsAudio() {
		AudioOnly = true
	}
	for {
		select {
		case <-clientTest.C:
			if OnDemand && !Config.HasViewer(name) {
				return ErrorStreamExitNoViewer
			}
		case <-keyTest.C:
			return ErrorStreamExitNoVideoOnStream
		case signals := <-RTSPClient.Signals:
			switch signals {
			case rtspv2.SignalCodecUpdate:
				Config.coAd(name, RTSPClient.CodecData)
			case rtspv2.SignalStreamRTPStop:
				return ErrorStreamExitRtspDisconnect
			}
		case packetAV := <-RTSPClient.OutgoingPacketQueue:
			if AudioOnly || packetAV.IsKeyFrame {
				keyTest.Reset(20 * time.Second)
			}
			Config.cast(name, *packetAV)
		}
	}
}
