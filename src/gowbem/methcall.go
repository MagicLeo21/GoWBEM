//    Copyright (C) 2017  Chen Yan
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
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func newMechCall(name string) *MethodCall {
	return &MethodCall{Name: name}
}

func (methCall *MethodCall) appendLocalClassPath(namespace string, className *ClassName) {
	var ns []Namespace = []Namespace{}
	for _, sub := range strings.Split(namespace, "/") {
		ns = append(ns, Namespace{Name: sub})
	}
	methCall.LocalClassPath = &LocalClassPath{
		LocalNamespacePath: &LocalNamespacePath{ns},
		ClassName:          className,
	}
}

func (methCall *MethodCall) appendLocalInstancePath(namespace string, instanceName *InstanceName) {
	var ns []Namespace = []Namespace{}
	for _, sub := range strings.Split(namespace, "/") {
		ns = append(ns, Namespace{Name: sub})
	}
	methCall.LocalInstancePath = &LocalInstancePath{
		LocalNamespacePath: &LocalNamespacePath{ns},
		InstanceName:       instanceName,
	}
}

func (methCall *MethodCall) getObjectPathString() string {
	obj := ""
	if nil != methCall.LocalClassPath &&
		nil != methCall.LocalClassPath.ClassName &&
		nil != methCall.LocalClassPath.LocalNamespacePath {
		for i, ns := range methCall.LocalClassPath.LocalNamespacePath.Namespace {
			if 0 == i {
				obj += fmt.Sprintf("%s", ns.Name)
			} else {
				obj += fmt.Sprintf("/%s", ns.Name)
			}
		}
		obj += fmt.Sprintf(":%s", methCall.LocalClassPath.ClassName.Name)
	} else if nil != methCall.LocalInstancePath &&
		nil != methCall.LocalInstancePath.InstanceName &&
		nil != methCall.LocalInstancePath.LocalNamespacePath {
		for i, ns := range methCall.LocalInstancePath.LocalNamespacePath.Namespace {
			if 0 == i {
				obj += fmt.Sprintf("%s", ns.Name)
			} else {
				obj += fmt.Sprintf("/%s", ns.Name)
			}
		}
		obj += fmt.Sprintf(":%s", methCall.LocalInstancePath.InstanceName.ClassName)
		for i, key := range methCall.LocalInstancePath.InstanceName.KeyBinding {
			if 0 == i {
				obj += fmt.Sprintf(".%s=\"%s\"", key.Name, key.KeyValue.KeyValue)
			} else {
				obj += fmt.Sprintf(",%s=\"%s\"", key.Name, key.KeyValue.KeyValue)
			}
		}
	}
	return obj
}

func (conn *WBEMConnection) doPostMethodCall(method string, object string, content []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s://%s:%d/%s", conn.scheme, conn.host, conn.port, DefaultRequestURI), bytes.NewReader(content))
	if nil != err {
		return nil, err
	}
	req.SetBasicAuth(conn.username, conn.password)
	req.Header.Add("Content-Type", "application/xml; charset=\"utf-8\"")
	req.Header.Add("Host", fmt.Sprintf("%s:%d", conn.host, conn.port))
	req.Header.Add("Accept-Encoding", "identity")
	req.Header["TE"] = append(req.Header["TE"], "trailers")
	req.Header[HttpHdrOperation] = append(req.Header[HttpHdrOperation], "MethodCall")
	req.Header[HttpHdrMethod] = append(req.Header[HttpHdrMethod], method)
	req.Header[HttpHdrObject] = append(req.Header[HttpHdrObject], object)
	res, err := conn.httpc.Do(req)
	if nil != err {
		return nil, err
	}
	defer res.Body.Close()
	if 200 != res.StatusCode {
		err = fmt.Errorf("HTTP_ERR - %d - %s", res.StatusCode, res.Status)
		return nil, err
	}
	return ioutil.ReadAll(res.Body)
}

func (conn *WBEMConnection) methodCall(call *MethodCall) (*MethodResponse, error) {
	if nil == call {
		return nil, conn.oops(ErrFailed, "")
	}

	var cim CIM = CIM{
		CIMVersion: "2.0",
		DTDVersion: "2.0",
		Message: &Message{
			ID:              "1001",
			ProtocolVersion: "1.0",
			SimpleReq: &SimpleReq{
				MethodCall: call},
		},
	}
	raw, err := xml.Marshal(&cim)
	if nil != err {
		return nil, err
	}
	raw = append([]byte(xml.Header), raw...)
	raw, err = conn.doPostMethodCall(call.Name, call.getObjectPathString(), raw)
	if nil != err {
		return nil, err
	}
	cim = CIM{}
	err = xml.Unmarshal(raw, &cim)
	if nil != err {
		return nil, err
	}
	if nil == cim.Message || nil == cim.Message.SimpleRsp || nil == cim.Message.SimpleRsp.MethodResponse {
		return nil, conn.oops(ErrFailed, "")
	}
	return cim.Message.SimpleRsp.MethodResponse, nil
}
