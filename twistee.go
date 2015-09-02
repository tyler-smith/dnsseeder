package main

import (
	"net"
	"time"

	"github.com/btcsuite/btcd/wire"
)

// Twistee struct contains details on one twister client
type twistee struct {
	na           *wire.NetAddress // holds ip address & port details
	lastConnect  time.Time        // last time we sucessfully connected to this client
	lastTry      time.Time        // last time we tried to connect to this client
	crawlStart   time.Time        // time when we started the last crawl
	statusTime   time.Time        // time the status was last updated
	crawlActive  bool             // are we currently crawling this client
	connectFails uint32           // number of times we have failed to connect to this client
	statusStr    string           // string with last error or OK details
	version      int32            // remote client protocol version
	strVersion   string           // remote client user agent
	services     wire.ServiceFlag // remote client supported services
	lastBlock    int32            // remote client last block
	status       uint32           // rg,cg,wg,ng
	rating       uint32           // if it reaches 100 then we mark them statusNG
	nonstdIP     net.IP           // if not using the default port then this is the encoded ip containing the actual port
	dnsType      uint32           // what dns type this client is
}

// status2str will return the string description of the status
func (tw twistee) status2str() string {
	switch tw.status {
	case statusRG:
		return "statusRG"
	case statusCG:
		return "statusCG"
	case statusWG:
		return "statusWG"
	case statusNG:
		return "statusNG"
	default:
		return "Unknown"
	}
}

// dns2str will return the string description of the dns type
func (tw twistee) dns2str() string {
	switch tw.dnsType {
	case dnsV4Std:
		return "v4 standard port"
	case dnsV4Non:
		return "v4 non-standard port"
	case dnsV6Std:
		return "v6 standard port"
	case dnsV6Non:
		return "v6 non-standard port"
	default:
		return "Unknown DNS Type"
	}
}

/*

 */
