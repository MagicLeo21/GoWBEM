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

type CIMErr struct {
	ErrCode int
	ErrName string
	ErrDesc string
}

func errName(err int) string {
	switch err {
	case ErrFailed:
		return "CIM_ERR_FAILED"
	case ErrAccessDenied:
		return "CIM_ERR_ACCESS_DENIED"
	case ErrInvalidNamespace:
		return "CIM_ERR_INVALID_NAMESPACE"
	case ErrInvalidParameter:
		return "CIM_ERR_INVALID_PARAMETER"
	case ErrInvalidClass:
		return "CIM_ERR_INVALID_CLASS"
	case ErrNotFound:
		return "CIM_ERR_NOT_FOUND"
	case ErrNotSupported:
		return "CIM_ERR_NOT_SUPPORTED"
	case ErrClassHasChildren:
		return "CIM_ERR_CLASS_HAS_CHILDREN"
	case ErrClassHasInstances:
		return "CIM_ERR_CLASS_HAS_INSTANCES"
	case ErrInvalidSuperclass:
		return "CIM_ERR_INVALID_SUPERCLASS"
	case ErrAlreadyExists:
		return "CIM_ERR_ALREADY_EXISTS"
	case ErrNoSuchProperty:
		return "CIM_ERR_NO_SUCH_PROPERTY"
	case ErrTypeMismatch:
		return "CIM_ERR_TYPE_MISMATCH"
	case ErrQueryLanguageNotSupported:
		return "CIM_ERR_QUERY_LANGUAGE_NOT_SUPPORTED"
	case ErrInvalidQuery:
		return "CIM_ERR_INVALID_QUERY"
	case ErrMethodNotAvailable:
		return "CIM_ERR_METHOD_NOT_AVAILABLE"
	case ErrMethodNotFound:
		return "CIM_ERR_METHOD_NOT_FOUND"
	case ErrNameSpaceNotEmpty:
		return "CIM_ERR_NAMESPACE_NOT_EMPTY"
	case ErrInvalidEnumerationContext:
		return "CIM_ERR_INVALID_ENUMERATION_CONTEXT"
	case ErrInvalidOperationTimeout:
		return "CIM_ERR_INVALID_OPERATION_TIMEOUT"
	case ErrPullHasBeenAbandoned:
		return "CIM_ERR_PULL_HAS_BEEN_ABANDONED"
	case ErrPullCannotBeAbandoned:
		return "CIM_ERR_PULL_CANNOT_BE_ABANDONED"
	case ErrFilteredEnumerationNotSupported:
		return "CIM_ERR_FILTERED_ENUMERATION_NOT_SUPPORTED"
	case ErrContinuationOnErrorNotSupported:
		return "CIM_ERR_CONTINUATION_ON_ERROR_NOT_SUPPORTED"
	case ErrServerLimitsExceeded:
		return "CIM_ERR_SERVER_LIMITS_EXCEEDED"
	case ErrServerIsShuttingDown:
		return "CIM_ERR_SERVER_IS_SHUTTING_DOWN"
	default:
	}
	return "CIM_ERR_FAILED"
}

func errDesc(err int) string {
	switch err {
	case ErrFailed:
		return "A general error occurred"
	case ErrAccessDenied:
		return "Resource not available"
	case ErrInvalidNamespace:
		return "The target namespace does not exist"
	case ErrInvalidParameter:
		return "Parameter value(s) invalid"
	case ErrInvalidClass:
		return "The specified Class does not exist"
	case ErrNotFound:
		return "Requested object could not be found"
	case ErrNotSupported:
		return "Operation not supported"
	case ErrClassHasChildren:
		return "Class has subclasses"
	case ErrClassHasInstances:
		return "Class has instances"
	case ErrInvalidSuperclass:
		return "Superclass does not exist"
	case ErrAlreadyExists:
		return "Object already exists"
	case ErrNoSuchProperty:
		return "Property does not exist"
	case ErrTypeMismatch:
		return "Value incompatible with type"
	case ErrQueryLanguageNotSupported:
		return "Query language not supported"
	case ErrInvalidQuery:
		return "Query not valid"
	case ErrMethodNotAvailable:
		return "Extrinsic method not executed"
	case ErrMethodNotFound:
		return "Extrinsic method does not exist"
	case ErrNameSpaceNotEmpty:
		return "Namespace not empty"
	case ErrInvalidEnumerationContext:
		return "Enumeration context is invalid"
	case ErrInvalidOperationTimeout:
		return "Operation timeout not supported"
	case ErrPullHasBeenAbandoned:
		return "Pull operation has been abandoned"
	case ErrPullCannotBeAbandoned:
		return "Attempt to abandon a pull operation failed"
	case ErrFilteredEnumerationNotSupported:
		return "Filtered pulled enumeration not supported"
	case ErrContinuationOnErrorNotSupported:
		return "WBEM server does not support continuation on error"
	case ErrServerLimitsExceeded:
		return "WBEM server limits exceeded"
	case ErrServerIsShuttingDown:
		return "WBEM server is shutting down"
	default:
	}
	return "A general error occurred"
}

func (err CIMErr) Error() string {
	return fmt.Sprintf("%v - %v - %v", err.ErrCode, err.ErrName, err.ErrDesc)
}

func (conn *WBEMConnection) oops(err int) error {
	if 1 > err || 18 == err || 19 == err || 28 < err {
		err = ErrFailed
	}
	return CIMErr{
		err,
		errName(err),
		errDesc(err),
	}
}
