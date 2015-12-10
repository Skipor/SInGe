package singe

/*
#cgo LDFLAGS: -ldictgen_c

#include <stdlib.h>
#include "dictionary_c.h"
*/
import "C"

import (
	"unsafe"
)

const (
	DefaultMaxDict            = 2 << 20
	DefaultMinLen             = 8
	DefaultStopSymbol         = byte(0)
	DefaultAutomatonSizeLimit = 2 << 25
	DefaultScoreDecreaseCoef  = 1.0
	DefaultMinDocsOccursIn    = 2
)

func UnsafeBytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

type DictBuilder interface {
	Write(data []byte)
	StartDocAndWrite(data []byte)
	GetDict() (dict []byte)
}


type Singe interface {
	DictBuilder
	Free()
}

type singe struct {
	singe      C.SInGe
	stopSymbol byte
}

func NewCustom(
	maxDict uint,
	minLen uint,
	stopSymbol byte,
	automatonSizeLimit uint,
	scoreDecreaseCoef float64,
	minDocsOccursIn uint,
) Singe {
	return singe{
		C.SInGeInit(
			C.size_t(maxDict),
			C.size_t(minLen),
			C.char(stopSymbol),
			C.size_t(automatonSizeLimit),
			C.double(scoreDecreaseCoef),
			C.size_t(minDocsOccursIn),
		),
		stopSymbol,
	}
}

func New(maxDict uint) Singe {
	return NewCustom(
		maxDict,
		DefaultMinLen,
		DefaultStopSymbol,
		DefaultAutomatonSizeLimit,
		DefaultScoreDecreaseCoef,
		DefaultMinDocsOccursIn,
	)
}

func (s singe) Free() {
	C.SInGeFree(s.singe)
}

func (s singe) Write(data []byte) {
	cData := C.CString(UnsafeBytesToString(data))
	C.SInGeAddDocument(s.singe, cData, C.size_t(len(data)))
	C.free(unsafe.Pointer(cData))
}


func (s singe) StartDocAndWrite(data []byte) { //for incremental update
	cData := C.CString(UnsafeBytesToString(data))
	C.SInGeAddDocumentViaStopSymbol(s.singe, cData, C.size_t(len(data)))
	C.free(unsafe.Pointer(cData))
}

func (s singe) GetDict() (dict []byte) {
	dict_c := C.SInGeGetDict(s.singe)
	dict = C.GoBytes(unsafe.Pointer(dict_c.data), C.int(dict_c.length))
	C.free(unsafe.Pointer(dict_c.data))
	return
}
