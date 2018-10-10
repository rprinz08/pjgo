package main

import (
	"flag"
	"os"
	"pjgo/pjsua2"
	"strings"
	"time"

	logging "github.com/op/go-logging"
)

// Create main go logger
var log = logging.MustGetLogger("pjgo")

// Shows usage informations
func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

// Initialize application
func initialize() {
	// Customize logger.
	logBackend := logging.NewLogBackend(os.Stderr, "", 0)
	// Log format string. Everything except the message has a custom color
	// which is dependent on the log level. Many fields have a custom output
	// formatting too, eg. the time returns the hour down to the millisecond.
	logFormat := logging.MustStringFormatter(
		`%{color}%{time:2006.01.02 15:04:05.000} %{program} %{pid:05d} %{id:04x} %{level:.3s} ▶%{color:reset} %{message}`,
	)
	logBackendFormatter := logging.NewBackendFormatter(logBackend, logFormat)
	logging.SetBackend(logBackendFormatter)

	// Process command line args
	flag.Usage = usage
	flag.Parse()
}

// Terminate application
func terminate(exitCode int) {
	log.Infof("pjgo golang pjsip binding test stopped")
	os.Exit(exitCode)
}

// MyAccount extends pjsau2.Account to get notifications etc.
type MyAccount struct {
	pjsua2.Account
}

// onRegState callback called by pjsip when account Sip registration changes
func (a MyAccount) onRegState(prm pjsua2.OnRegStateParam) {
	log.Infof("*** on registration state (%s) (%s)",
		prm.GetCode(), prm.GetReason())
}

// NewMyAccount constructs a new MyAccount
func NewMyAccount() *MyAccount {
	p := MyAccount{}
	p.Account = pjsua2.NewAccount()
	return &p
}

// MyLogger extends base pjsua2 logger to create a customized logger
type MyLogger struct {
	pjsua2.LogWriter
}

// NewMyLogger constructs a new custom logger instance
func NewMyLogger() *MyLogger {
	p := MyLogger{}
	p.LogWriter = pjsua2.NewDirectorLogWriter(p)
	return &p
}

// Write gets calld by pjsip to output new log entries
func (l MyLogger) Write(entry pjsua2.LogEntry) {
	msg := strings.TrimSpace(entry.GetMsg())
	switch pjLogLevel := entry.GetLevel(); pjLogLevel {
	case 0: // fatal
		log.Fatalf("%s", msg)
	case 1: // error
		log.Errorf("%s", msg)
	case 2: // warning
		log.Warningf("%s", msg)
	case 3: // info
		log.Infof("%s", msg)
	case 4: // debug
		log.Debugf("%s", msg)
	case 5: // debug
		log.Debugf("%s", msg)
	case 6: // debug
		log.Debugf("%s", msg)
	}
}

// Main routine
func main() {
	initialize()
	log.Infof("pjgo golang pjsip binding test started")

	// initialize endpoint
	ep := pjsua2.NewEndpoint()
	ep.LibCreate()
	epCfg := pjsua2.NewEpConfig()

	// configure logging
	logCfg := pjsua2.NewLogConfig()
	logCfg.SetWriter(NewMyLogger())
	logCfg.SetDecor(
		uint(pjsua2.PJ_LOG_HAS_SENDER) |
			uint(pjsua2.PJ_LOG_HAS_THREAD_ID))
	logCfg.SetLevel(6)
	epCfg.SetLogConfig(logCfg)

	ep.LibInit(epCfg)

	// Create SIP transport.
	tCfg := pjsua2.NewTransportConfig()
	tCfg.SetPort(5060)
	ep.TransportCreate(pjsua2.PJSIP_TRANSPORT_UDP, tCfg)

	// Start the library (worker threads etc)
	ep.LibStart()

	// Configure an account configuration
	aCfg := pjsua2.NewAccountConfig()
	aCfg.SetIdUri("sip:test@pjsip.org")
	aRegCfg := pjsua2.NewAccountRegConfig()
	aRegCfg.SetRegistrarUri("sip:pjsip.org")
	aCfg.SetRegConfig(aRegCfg)
	cred := pjsua2.NewAuthCredInfo("digest", "*", "test", 0, "secret")
	aSipCfg := pjsua2.NewAccountSipConfig()
	credVec := pjsua2.NewAuthCredInfoVector(int64(1))
	credVec.Add(cred)
	aSipCfg.SetAuthCreds(credVec)
	aCfg.SetSipConfig(aSipCfg)

	// Create the account
	acc := NewMyAccount()
	acc.Create(aCfg)

	// Here we don't have anything else to do ...
	time.Sleep(5 * time.Second)

	// terminate
	ep.LibDestroy()

	exitCode := 0
	terminate(exitCode)
}
