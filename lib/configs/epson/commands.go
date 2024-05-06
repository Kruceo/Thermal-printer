package epson

import "thermal-printer/lib"

// init, clear buffers
var Init = lib.CommandBytes{0x1B, 0x40}

// set character set to "TURKISH"
var SetCharacterSet = lib.CommandBytes{0x1B, 0x74, 48}

var FullCut = lib.CommandBytes{0x1D, 0x56, 0x00}

var FeedLines = lib.CommandBytes{0x1b, 0x64, 1}

// prefer using init
var ClearBuffer = lib.CommandBytes{0x10, 0x14}

var ReverseFeed = lib.CommandBytes{0x1B, 0x4B, 255}
