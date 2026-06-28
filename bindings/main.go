package torrServer

import (
	server "server"
)

func StartTorrentServer(pathdb string) {
	server.Start(pathdb, "", false, false)
}

// StartTorrentServerOnPort starts the server on a specific HTTP port. An empty
// port falls back to the server default (8090).
func StartTorrentServerOnPort(pathdb, port string) {
	server.Start(pathdb, port, false, false)
}

func WaitTorrentServer() {
	server.WaitServer()
}

func StopTorrentServer() {
	server.Stop()
}

func AddTrackers(trackers string) {
	server.AddTrackers(trackers)
}
