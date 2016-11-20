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

package main

import (
	"fmt"
	"gowbem"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func GetHostName() (string, error) {
	name, err := os.Hostname()
	if nil != err {
		return "", err
	}
	return name, nil
}

func GetDestAddr(urlstr string) (string, error) {
	res, err := url.Parse(urlstr)
	if nil != err {
		return "", err
	}
	host := strings.Split(res.Host, ":")
	return host[0], nil
}

func GetLocalIP(dest string) (string, error) {
	conn, err := net.Dial("udp", fmt.Sprintf("%s:80", dest))
	if nil != err {
		return "", err
	}
	defer conn.Close()
	return strings.Split(conn.LocalAddr().String(), ":")[0], nil
}

func ListenerHandler(writer http.ResponseWriter, req *http.Request) {
	body, _ := ioutil.ReadAll(req.Body)
	fmt.Println(string(body))
}

func SubscriptionThread(destURL, localName, srcNS, localIP string, localPort int) {
	conn, err := gowbem.NewWBEMConn(destURL)
	if nil != err {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance := NewIndicationFilter(localName, srcNS)
	fmt.Println("Creating IndicationFilter...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance = NewListenerDestination(localName, localIP, localPort)
	fmt.Println("Creating ListenerDestination...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	iInstance = NewIndicationSubscription(localName, conn.GetNamespace(), localIP, localPort)
	fmt.Println("Creating IndicationSubscription...")
	err = conn.CreateInstance(iInstance)
	if nil != err && "11 - CIM_ERR_ALREADY_EXISTS - Object already exists" != err.Error() {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Subscription Done!")
}

func ListenerThread(localIP string, localPort int) {
	fmt.Printf("Start listening on %s:%d...\n", localIP, localPort)
	http.HandleFunc("/", ListenerHandler)
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", localIP, localPort), nil)
	if nil != err {
		fmt.Println("ERR - Failed to start listener:", err.Error())
		os.Exit(1)
	}
}

func main() {
	if 2 != len(os.Args) {
		os.Exit(1)
	}

	iDestAddr, _ := GetDestAddr(os.Args[1])
	iLocalName, _ := GetHostName()
	iLocalIP, _ := GetLocalIP(iDestAddr)
	iSourceNS := "root/cimv2"
	iLocalPort := 59988

	go SubscriptionThread(os.Args[1], iLocalName, iSourceNS, iLocalIP, iLocalPort)
	ListenerThread(iLocalIP, iLocalPort)
}
