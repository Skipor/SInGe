package singe_test

import (
	"github.com/cscenter/SInGe/golang/singe"
//	"log"
	"testing"
	//	"github.com/stretchr/testify/assert"
	"bytes"
	"github.com/stretchr/testify/assert"
)

func TestNewFreeNotFails(t *testing.T) {
	s := singe.New(50)
	s.Free()
}

func TestAddNotFails(t *testing.T) {
	s := singe.New(50)
	s.Write([]byte("aaaaaaa"))
	s.Free()
}
func TestGetDictNotFails(t *testing.T) {
	s := singe.New(50)
	s.GetDict()
	s.Free()
}

func addTestSinge() singe.Singe {
	//	return singe.NewCustom(
	//		singe.DefaultMaxDict,
	//		singe.DefaultMinLen,
	//		'#',
	//		singe.DefaultAutomatonSizeLimit,
	//		singe.DefaultScoreDecreaseCoef,
	//	)
	return singe.NewCustom(
		1000,
		3,
		'#',
		10000,
		1.0,
		2,
	)

}

type dictTestCase struct {
	in          []string
	contais     []string
	notContains []string
}

var addTestTable = []dictTestCase{
	{
		[]string{
			"abacaba",
			"qwecabarty",
			"caba_cabaqwe",
			"abacaba",
			"qwecabarty",
			"caba_cabaqwe",
		},
		[]string{
			"caba",
			"qwe",
		},
		[]string{
			//"qweqwe",//todo
			//"cabacaba",//todo
		},
	},
	{
		[]string{
			"qaabaaa",
			"waabaaaa",
			"eaabaaaaa",
			"raabaaaaaa",
			"tabacabaaa",
			"qaabaaa",
			"waabaaaa",
			"eaabaaaaa",
			"raabaaaaaa",
			"tabacabaaa",
			//"aabaaaaaa",
		},
		[]string{
			"aaa",
			"aba",
		},
		[]string{
			"aca",
			"d",
		},
	},
}

func checkDict(t *testing.T, s singe.DictBuilder, testData dictTestCase) {
	dict := s.GetDict()

	t.Logf("Dict: %s", dict)
	for _, c := range testData.contais {
		assert.True(t, bytes.Contains(dict, []byte(c)))
	}
	for _, nc := range testData.notContains {
		assert.False(t, bytes.Contains(dict, []byte(nc)))
	}
}

func TestAddDoc(t *testing.T) {
	for _, testData := range addTestTable {
		s := addTestSinge()
		for _, in := range testData.in {
			s.StartDocAndWrite([]byte(in))
		}
		checkDict(t, s, testData)
		s.Free()
	}
}

func TestAddIncremental(t *testing.T) {
	for _, testData := range addTestTable {
		s := addTestSinge()
		for _, in := range testData.in {
			s.StartDocAndWrite([]byte(in))
		}
		checkDict(t, s, testData)
		s.Free()
	}
}
