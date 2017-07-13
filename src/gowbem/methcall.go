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
	"encoding/xml"
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
	raw, err = conn.doPostMethodCall(call.Name, raw)
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
