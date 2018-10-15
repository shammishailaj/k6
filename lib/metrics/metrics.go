/*
 *
 * k6 - a next-generation load testing tool
 * Copyright (C) 2017 Load Impact
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as
 * published by the Free Software Foundation, either version 3 of the
 * License, or (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 */

package metrics

import (
	"github.com/loadimpact/k6/stats"
)

//TODO: refactor this, using non thread-safe global variables seems like a bad idea for various reasons...

var (
	// We are certain no errors will happen here
	newMetric = func(name string, typ stats.MetricType, t ...stats.ValueType) *stats.Metric {
		r, err := stats.New(name, typ, t...)
		if err != nil {
			// if they do happen we have done something very wrong
			panic(err)
		}
		return r
	}

	// Engine-emitted.
	VUs               = newMetric("vus", stats.Gauge)
	VUsMax            = newMetric("vus_max", stats.Gauge)
	Iterations        = newMetric("iterations", stats.Counter)
	IterationDuration = newMetric("iteration_duration", stats.Trend, stats.Time)
	Errors            = newMetric("errors", stats.Counter)

	// Runner-emitted.
	Checks        = newMetric("checks", stats.Rate)
	GroupDuration = newMetric("group_duration", stats.Trend, stats.Time)

	// HTTP-related.
	HTTPReqs              = newMetric("http_reqs", stats.Counter)
	HTTPReqDuration       = newMetric("http_req_duration", stats.Trend, stats.Time)
	HTTPReqBlocked        = newMetric("http_req_blocked", stats.Trend, stats.Time)
	HTTPReqConnecting     = newMetric("http_req_connecting", stats.Trend, stats.Time)
	HTTPReqTLSHandshaking = newMetric("http_req_tls_handshaking", stats.Trend, stats.Time)
	HTTPReqSending        = newMetric("http_req_sending", stats.Trend, stats.Time)
	HTTPReqWaiting        = newMetric("http_req_waiting", stats.Trend, stats.Time)
	HTTPReqReceiving      = newMetric("http_req_receiving", stats.Trend, stats.Time)

	// Websocket-related
	WSSessions         = newMetric("ws_sessions", stats.Counter)
	WSMessagesSent     = newMetric("ws_msgs_sent", stats.Counter)
	WSMessagesReceived = newMetric("ws_msgs_received", stats.Counter)
	WSPing             = newMetric("ws_ping", stats.Trend)
	WSSessionDuration  = newMetric("ws_session_duration", stats.Trend, stats.Time)
	WSConnecting       = newMetric("ws_connecting", stats.Trend, stats.Time)

	// Network-related; used for future protocols as well.
	DataSent     = newMetric("data_sent", stats.Counter, stats.Data)
	DataReceived = newMetric("data_received", stats.Counter, stats.Data)
)
