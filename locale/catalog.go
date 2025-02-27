// Code generated by running "go generate" in golang.org/x/text. DO NOT EDIT.

package locale

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
	"golang.org/x/text/message/catalog"
)

type dictionary struct {
	index []uint32
	data  string
}

func (d *dictionary) Lookup(key string) (data string, ok bool) {
	p, ok := messageKeyToIndex[key]
	if !ok {
		return "", false
	}
	start, end := d.index[p], d.index[p+1]
	if start == end {
		return "", false
	}
	return d.data[start:end], true
}

func init() {
	dict := map[string]catalog.Dictionary{
		"en_US": &dictionary{index: en_USIndex, data: en_USData},
		"ja_JP": &dictionary{index: ja_JPIndex, data: ja_JPData},
		"tr_TR": &dictionary{index: tr_TRIndex, data: tr_TRData},
	}
	fallback := language.MustParse("en-US")
	cat, err := catalog.NewFromMap(dict, catalog.Fallback(fallback))
	if err != nil {
		panic(err)
	}
	message.DefaultCatalog = cat
}

var messageKeyToIndex = map[string]int{
	"*bump*":                56,
	"*burp*":                58,
	"*click*":               67,
	"*creak*":               72,
	"*munch*":               59,
	"*pmub*":                57,
	"*snarf*":               63,
	"*thud*":                65,
	"*whump*":               64,
	"...and so you return.": 27,
	"Archetype is a collection of attributes and traits":          7,
	"Brains determines spell damage and ability to by-pass traps": 10,
	"Confirm deletion":                                2,
	"Confirm your password to register.":              29,
	"Create a new hero or select a previous one.":     6,
	"Funk determines luck and area of effect bonuses": 11,
	"It's already closed.":                            69,
	"It's already open.":                              73,
	"It's locked.":                                    70,
	"Swole determines damage and health":              8,
	"There is nothing there to open.":                 74,
	"Traits are unique modifiers to an archetype":     12,
	"You can't apply that.":                           60,
	"You can't pick that up.":                         61,
	"You can't reach that.":                           62,
	"You can't use %s.":                               36,
	"You close the door.":                             68,
	"You don't have that item.":                       35,
	"You kick at the air.":                            66,
	"You open the door.":                              71,
	"You're too full to eat that.":                    37,
	"Zooms determines speed and dodge":                9,
	"archetype must be selected":                      14,
	"back":                                            20,
	"bad password":                                    48,
	"character does not exist":                        50,
	"character exists":                                49,
	"character name":                                  3,
	"confirm":                                         17,
	"create":                                          4,
	"delete":                                          1,
	"heavy":                                           34,
	"it is already closed":                            39,
	"it is already open":                              38,
	"join":                                            0,
	"light":                                           32,
	"logged in!":                                      28,
	"logging in...":                                   25,
	"login":                                           18,
	"login. you will be prompted to register if username does not exist.": 21,
	"logout":                     5,
	"medium":                     33,
	"melee":                      45,
	"missing slot":               40,
	"missing slots":              41,
	"name cannot be empty":       75,
	"name must not be empty":     13,
	"no such archetype":          51,
	"no such fixture":            54,
	"no such tile":               55,
	"no such user":               52,
	"not logged in":              47,
	"password":                   16,
	"password must not be empty": 23,
	"passwords must match":       24,
	"petty":                      31,
	"range":                      43,
	"register":                   19,
	"registering...":             26,
	"thrown":                     44,
	"unarmed":                    46,
	"used slots":                 42,
	"user exists":                53,
	"username":                   15,
	"username must not be empty": 22,
	"world name":                 30,
}

var en_USIndex = []uint32{ // 77 elements
	// Entry 0 - 1F
	0x00000000, 0x00000005, 0x0000000c, 0x0000001d,
	0x0000002c, 0x00000033, 0x0000003a, 0x00000066,
	0x00000099, 0x000000bc, 0x000000dd, 0x00000119,
	0x00000149, 0x00000175, 0x0000018c, 0x000001a7,
	0x000001b0, 0x000001b9, 0x000001c1, 0x000001c7,
	0x000001d0, 0x000001d5, 0x00000219, 0x00000234,
	0x0000024f, 0x00000264, 0x00000272, 0x00000281,
	0x00000297, 0x000002a2, 0x000002c5, 0x000002d0,
	// Entry 20 - 3F
	0x000002d6, 0x000002dc, 0x000002e3, 0x000002e9,
	0x00000303, 0x00000318, 0x00000335, 0x00000348,
	0x0000035d, 0x0000036a, 0x00000378, 0x00000383,
	0x00000389, 0x00000390, 0x00000396, 0x0000039e,
	0x000003ac, 0x000003b9, 0x000003ca, 0x000003e3,
	0x000003f5, 0x00000402, 0x0000040e, 0x0000041e,
	0x0000042b, 0x00000432, 0x00000439, 0x00000440,
	0x00000448, 0x0000045e, 0x00000476, 0x0000048c,
	// Entry 40 - 5F
	0x00000494, 0x0000049c, 0x000004a3, 0x000004b8,
	0x000004c0, 0x000004d4, 0x000004e9, 0x000004f6,
	0x00000509, 0x00000511, 0x00000524, 0x00000544,
	0x00000559,
} // Size: 332 bytes

const en_USData string = "" + // Size: 1369 bytes
	"\x02join\x02delete\x02Confirm deletion\x02character name\x02create\x02lo" +
	"gout\x02Create a new hero or select a previous one.\x02Archetype is a co" +
	"llection of attributes and traits\x02Swole determines damage and health" +
	"\x02Zooms determines speed and dodge\x02Brains determines spell damage a" +
	"nd ability to by-pass traps\x02Funk determines luck and area of effect b" +
	"onuses\x02Traits are unique modifiers to an archetype\x02name must not b" +
	"e empty\x02archetype must be selected\x02username\x02password\x02confirm" +
	"\x02login\x02register\x02back\x02login. you will be prompted to register" +
	" if username does not exist.\x02username must not be empty\x02password m" +
	"ust not be empty\x02passwords must match\x02logging in...\x02registering" +
	"...\x02...and so you return.\x02logged in!\x02Confirm your password to r" +
	"egister.\x02world name\x02petty\x02light\x02medium\x02heavy\x02You don't" +
	" have that item.\x02You can't use %[1]s.\x02You're too full to eat that." +
	"\x02it is already open\x02it is already closed\x02missing slot\x02missin" +
	"g slots\x02used slots\x02range\x02thrown\x02melee\x02unarmed\x02not logg" +
	"ed in\x02bad password\x02character exists\x02character does not exist" +
	"\x02no such archetype\x02no such user\x02user exists\x02no such fixture" +
	"\x02no such tile\x02*bump*\x02*pmub*\x02*burp*\x02*munch*\x02You can't a" +
	"pply that.\x02You can't pick that up.\x02You can't reach that.\x02*snarf" +
	"*\x02*whump*\x02*thud*\x02You kick at the air.\x02*click*\x02You close t" +
	"he door.\x02It's already closed.\x02It's locked.\x02You open the door." +
	"\x02*creak*\x02It's already open.\x02There is nothing there to open.\x02" +
	"name cannot be empty"

var ja_JPIndex = []uint32{ // 77 elements
	// Entry 0 - 1F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 20 - 3F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 40 - 5F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000,
} // Size: 332 bytes

const ja_JPData string = ""

var tr_TRIndex = []uint32{ // 77 elements
	// Entry 0 - 1F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 20 - 3F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	// Entry 40 - 5F
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000, 0x00000000, 0x00000000, 0x00000000,
	0x00000000,
} // Size: 332 bytes

const tr_TRData string = ""

// Total table size 2365 bytes (2KiB); checksum: 7B3534F5
