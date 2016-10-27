//    Copyright (C) 2016  Chen Yan
//
//    This program is free software: you can redistribute it and/or modify
//    it under the terms of the GNU General Public License as published by
//    the Free Software Foundation, either version 3 of the License, or
//    (at your option) any later version.
//
//    This program is distributed in the hope that it will be useful,
//    but WITHOUT ANY WARRANTY; without even the implied warranty of
//    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//    GNU General Public License for more details.
//
//    You should have received a copy of the GNU General Public License
//    along with this program.  If not, see <http://www.gnu.org/licenses/>.
//
//    Author: Chen Yan <leochenlinux@gmail.com>

package gowbem

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

const (
	SchemeHttp       string = "http"
	SchemeHttps      string = "https"
	DefaultPortHttp  int    = 5988
	DefaultPortHttps int    = 5989
	DefaultNamespace string = "root/cimv2"
)

var DefaultPortMap map[string]int = map[string]int{
	"http":  DefaultPortHttp,
	"https": DefaultPortHttps,
}

type WBEMConnection struct {
	scheme    string
	host      string
	port      int
	username  string
	password  string
	namespace string
	httpc     *http.Client
}

func NewWBEMConn(urlstr string, namespace string) (*WBEMConnection, error) {
	res, err := url.Parse(urlstr)
	if nil != err {
		return nil, err
	}

	var conn WBEMConnection
	host := strings.Split(res.Host, ":")
	for i, sub := range host {
		if 0 == i {
			conn.host = sub
		} else if 1 == i {
			conn.port, _ = strconv.Atoi(sub)
		}
	}
	conn.scheme = strings.ToLower(res.Scheme)
	if nil != res.User {
		conn.username = res.User.Username()
		conn.password, _ = res.User.Password()
	}
	conn.namespace = namespace
	if SchemeHttp != conn.scheme && SchemeHttps != conn.scheme {
		conn.scheme = SchemeHttp
	}
	if 0 == conn.port {
		conn.port = DefaultPortMap[conn.scheme]
	}
	if "" == conn.namespace {
		conn.namespace = DefaultNamespace
	}
	conn.httpc = &http.Client{
		Timeout: time.Minute,
	}
	if SchemeHttps == conn.scheme {
		conn.httpc.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
	return &conn, nil
}
