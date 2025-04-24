package logs

import log "github.com/calmw/clog"

var Clog *log.Logger

func InitLog() {
	clog := log.Root()
	clog.Debug("Starting daily financial report...")
	Clog = &clog
}
