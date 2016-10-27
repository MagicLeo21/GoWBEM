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
	"fmt"
)

const (
	// A general error occurred that is not covered by a more specific error code.
	ErrFailed = 1

	// Access to a CIM resource is not available to the client.
	ErrAccessDenied = 2

	// The target namespace does not exist.
	ErrInvalidNamespace = 3

	// One or more parameter values passed to the method are not valid.
	ErrInvalidParameter = 4

	// The specified class does not exist.
	ErrInvalidClass = 5

	// The requested object cannot be found. The operation can be unsupported on behalf of the WBEM server in general or on behalf of an implementation of a management profile.
	ErrNotFound = 6

	// The requested operation is not supported on behalf of the WBEM server, or on behalf of a provided class. If the operation is supported for a provided class but is not supported for particular instances of that class, then CIM_ERR_FAILED shall be used.
	ErrNotSupported = 7

	// The operation cannot be invoked on this class because it has subclasses.
	ErrClassHasChildren = 8

	// The operation cannot be invoked on this class because one or more instances of this class exist.
	ErrClassHasInstances = 9

	// The operation cannot be invoked because the specified superclass does not exist.
	ErrInvalidSuperclass = 10

	// The operation cannot be invoked because an object already exists.
	ErrAlreadyExists = 11

	// The specified property does not exist.
	ErrNoSuchProperty = 12

	// The value supplied is not compatible with the type.
	ErrTypeMismatch = 13

	// The query language is not recognized or supported.
	ErrQueryLanguageNotSupported = 14

	// The query is not valid for the specified query language.
	ErrInvalidQuery = 15

	// The extrinsic method cannot be invoked.
	ErrMethodNotAvailable = 16

	// The specified extrinsic method does not exist.
	ErrMethodNotFound = 17

	// The specified namespace is not empty.
	ErrNameSpaceNotEmpty = 20

	// The enumeration identified by the specified context cannot be found, is in a closed state, does not exist, or is otherwise invalid.
	ErrInvalidEnumerationContext = 21

	// The specified operation timeout is not supported by the WBEM server.
	ErrInvalidOperationTimeout = 22

	// The pull operation has been abandoned due to execution of a concurrent CloseEnumeration operation on the same enumeration.
	ErrPullHasBeenAbandoned = 23

	// The attempt to abandon a concurrent pull operation on the same enumeration
	// failed. The concurrent pull operation proceeds normally.
	ErrPullCannotBeAbandoned = 24

	// Using a a filter query in pulled enumerations is not supported by the WBEM server.
	ErrFilteredEnumerationNotSupported = 25

	// The WBEM server does not support continuation on error.
	ErrContinuationOnErrorNotSupported = 26

	// The WBEM server has failed the operation based upon exceeding server limits.
	ErrServerLimitsExceeded = 27

	// The WBEM server is shutting down and cannot process the operation.
	ErrServerIsShuttingDown = 28
)

var ErrName map[int]string = map[int]string{
	ErrFailed:                          "CIM_ERR_FAILED",
	ErrAccessDenied:                    "CIM_ERR_ACCESS_DENIED",
	ErrInvalidNamespace:                "CIM_ERR_INVALID_NAMESPACE",
	ErrInvalidParameter:                "CIM_ERR_INVALID_PARAMETER",
	ErrInvalidClass:                    "CIM_ERR_INVALID_CLASS",
	ErrNotFound:                        "CIM_ERR_NOT_FOUND",
	ErrNotSupported:                    "CIM_ERR_NOT_SUPPORTED",
	ErrClassHasChildren:                "CIM_ERR_CLASS_HAS_CHILDREN",
	ErrClassHasInstances:               "CIM_ERR_CLASS_HAS_INSTANCES",
	ErrInvalidSuperclass:               "CIM_ERR_INVALID_SUPERCLASS",
	ErrAlreadyExists:                   "CIM_ERR_ALREADY_EXISTS",
	ErrNoSuchProperty:                  "CIM_ERR_NO_SUCH_PROPERTY",
	ErrTypeMismatch:                    "CIM_ERR_TYPE_MISMATCH",
	ErrQueryLanguageNotSupported:       "CIM_ERR_QUERY_LANGUAGE_NOT_SUPPORTED",
	ErrInvalidQuery:                    "CIM_ERR_INVALID_QUERY",
	ErrMethodNotAvailable:              "CIM_ERR_METHOD_NOT_AVAILABLE",
	ErrMethodNotFound:                  "CIM_ERR_METHOD_NOT_FOUND",
	ErrNameSpaceNotEmpty:               "CIM_ERR_NAMESPACE_NOT_EMPTY",
	ErrInvalidEnumerationContext:       "CIM_ERR_INVALID_ENUMERATION_CONTEXT",
	ErrInvalidOperationTimeout:         "CIM_ERR_INVALID_OPERATION_TIMEOUT",
	ErrPullHasBeenAbandoned:            "CIM_ERR_PULL_HAS_BEEN_ABANDONED",
	ErrPullCannotBeAbandoned:           "CIM_ERR_PULL_CANNOT_BE_ABANDONED",
	ErrFilteredEnumerationNotSupported: "CIM_ERR_FILTERED_ENUMERATION_NOT_SUPPORTED",
	ErrContinuationOnErrorNotSupported: "CIM_ERR_CONTINUATION_ON_ERROR_NOT_SUPPORTED",
	ErrServerLimitsExceeded:            "CIM_ERR_SERVER_LIMITS_EXCEEDED",
	ErrServerIsShuttingDown:            "CIM_ERR_SERVER_IS_SHUTTING_DOWN",
}

var ErrDesc map[int]string = map[int]string{
	ErrFailed:                          "A general error occurred",
	ErrAccessDenied:                    "Resource not available",
	ErrInvalidNamespace:                "The target namespace does not exist",
	ErrInvalidParameter:                "Parameter value(s) invalid",
	ErrInvalidClass:                    "The specified Class does not exist",
	ErrNotFound:                        "Requested object could not be found",
	ErrNotSupported:                    "Operation not supported",
	ErrClassHasChildren:                "Class has subclasses",
	ErrClassHasInstances:               "Class has instances",
	ErrInvalidSuperclass:               "Superclass does not exist",
	ErrAlreadyExists:                   "Object already exists",
	ErrNoSuchProperty:                  "Property does not exist",
	ErrTypeMismatch:                    "Value incompatible with type",
	ErrQueryLanguageNotSupported:       "Query language not supported",
	ErrInvalidQuery:                    "Query not valid",
	ErrMethodNotAvailable:              "Extrinsic method not executed",
	ErrMethodNotFound:                  "Extrinsic method does not exist",
	ErrNameSpaceNotEmpty:               "Namespace not empty",
	ErrInvalidEnumerationContext:       "Enumeration context is invalid",
	ErrInvalidOperationTimeout:         "Operation timeout not supported",
	ErrPullHasBeenAbandoned:            "Pull operation has been abandoned",
	ErrPullCannotBeAbandoned:           "Attempt to abandon a pull operation failed",
	ErrFilteredEnumerationNotSupported: "Filtered pulled enumeration not supported",
	ErrContinuationOnErrorNotSupported: "WBEM server does not support continuation on error",
	ErrServerLimitsExceeded:            "WBEM server limits exceeded",
	ErrServerIsShuttingDown:            "WBEM server is shutting down",
}

type CIMErr struct {
	ErrCode int
	ErrName string
	ErrDesc string
}

func (err CIMErr) Error() string {
	return fmt.Sprintf("Error  %v - %v - %v", err.ErrCode, err.ErrName, err.ErrDesc)
}

func (conn *WBEMConnection) oops(err int) error {
	if 1 > err || 18 == err || 19 == err || 28 < err {
		err = ErrFailed
	}
	return CIMErr{
		err,
		ErrName[err],
		ErrDesc[err],
	}
}
