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
	"strings"
)

func NewIndicationFilter(hostname, namespace string) *gowbem.Instance {
	iInstance := gowbem.Instance{
		ClassName: "CIM_IndicationFilter",
		Property: []gowbem.Property{
			{
				Name:  "QueryLanguage",
				Type:  "string",
				Value: &gowbem.Value{"WQL"},
			}, {
				Name:  "Query",
				Type:  "string",
				Value: &gowbem.Value{"SELECT * FROM CIM_AlertIndication"},
			}, {
				Name:  "SourceNamespace",
				Type:  "string",
				Value: &gowbem.Value{namespace},
			}, {
				Name:  "Name",
				Type:  "string",
				Value: &gowbem.Value{"GoWbem:DefaultFilter"},
			}, {
				Name:  "CreationClassName",
				Type:  "string",
				Value: &gowbem.Value{"CIM_IndicationFilter"},
			}, {
				Name:  "SystemName",
				Type:  "string",
				Value: &gowbem.Value{hostname},
			}, {
				Name:  "SystemCreationClassName",
				Type:  "string",
				Value: &gowbem.Value{"CIM_ComputerSystem"},
			},
		},
	}
	return &iInstance
}

func NewListenerDestination(hostname, ipaddr string, port int) *gowbem.Instance {
	iInstance := gowbem.Instance{
		ClassName: "CIM_ListenerDestinationCIMXML",
		Property: []gowbem.Property{
			{
				Name:  "Destination",
				Type:  "string",
				Value: &gowbem.Value{fmt.Sprintf("http://%s:%d", ipaddr, port)},
			}, {
				Name:  "SystemCreationClassName",
				Type:  "string",
				Value: &gowbem.Value{"CIM_ComputerSystem"},
			}, {
				Name:  "SystemName",
				Type:  "string",
				Value: &gowbem.Value{hostname},
			}, {
				Name:  "CreationClassName",
				Type:  "string",
				Value: &gowbem.Value{"CIM_ListenerDestinationCIMXML"},
			}, {
				Name:  "Name",
				Type:  "string",
				Value: &gowbem.Value{fmt.Sprintf("GoWbem:%s", ipaddr)},
			},
		},
	}
	return &iInstance
}

func NewIndicationSubscription(hostname, namespace, ipaddr string, port int) *gowbem.Instance {
	ns := []gowbem.Namespace{}
	for _, sub := range strings.Split(namespace, "/") {
		ns = append(ns, gowbem.Namespace{Name: sub})
	}
	iInstance := gowbem.Instance{
		ClassName: "CIM_IndicationSubscription",
		PropertyReference: []gowbem.PropertyReference{
			{
				Name:           "Handler",
				ReferenceClass: "CIM_ListenerDestinationCIMXML",
				ValueReference: &gowbem.ValueReference{
					LocalInstancePath: &gowbem.LocalInstancePath{
						LocalNamespacePath: &gowbem.LocalNamespacePath{ns},
						InstanceName: &gowbem.InstanceName{
							ClassName: "CIM_ListenerDestinationCIMXML",
							KeyBinding: []gowbem.KeyBinding{
								{
									Name: "SystemCreationClassName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  "CIM_ComputerSystem",
									},
								}, {
									Name: "SystemName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  hostname,
									},
								}, {
									Name: "CreationClassName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  "CIM_ListenerDestinationCIMXML",
									},
								}, {
									Name: "Name",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  fmt.Sprintf("GoWbem:%s", ipaddr),
									},
								},
							},
						},
					},
				},
			}, {
				Name:           "Filter",
				ReferenceClass: "CIM_IndicationFilter",

				ValueReference: &gowbem.ValueReference{
					LocalInstancePath: &gowbem.LocalInstancePath{
						LocalNamespacePath: &gowbem.LocalNamespacePath{ns},
						InstanceName: &gowbem.InstanceName{
							ClassName: "CIM_IndicationFilter",
							KeyBinding: []gowbem.KeyBinding{
								{
									Name: "SystemCreationClassName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  "CIM_ComputerSystem",
									},
								}, {
									Name: "SystemName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  hostname,
									},
								}, {
									Name: "CreationClassName",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  "CIM_IndicationFilter",
									},
								}, {
									Name: "Name",
									KeyValue: &gowbem.KeyValue{
										ValueType: "string",
										Type:      "string",
										KeyValue:  "GoWbem:DefaultFilter",
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return &iInstance
}
