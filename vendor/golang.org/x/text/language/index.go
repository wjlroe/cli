// This file was generated by go generate; DO NOT EDIT

package language

// NumCompactTags is the number of common tags. The maximum tag is
// NumCompactTags-1.
const NumCompactTags = 710

// Size: 72 bytes, 2 elements
var specialTags = []Tag{
	{lang: 0x61, region: 0x6d, script: 0x0, pVariant: 0x5, pExt: 0xe, str: "ca-ES-valencia"},
	{lang: 0x9a, region: 0x132, script: 0x0, pVariant: 0x5, pExt: 0x5, str: "en-US-u-va-posix"},
}

var coreTags = map[uint32]uint16{
	0x0:        0,   // und
	0x00a00000: 3,   // af
	0x00a000d0: 4,   // af-NA
	0x00a0015e: 5,   // af-ZA
	0x00b00000: 6,   // agq
	0x00b00051: 7,   // agq-CM
	0x00d00000: 8,   // ak
	0x00d0007e: 9,   // ak-GH
	0x01100000: 10,  // am
	0x0110006e: 11,  // am-ET
	0x01500000: 12,  // ar
	0x01500001: 13,  // ar-001
	0x01500022: 14,  // ar-AE
	0x01500038: 15,  // ar-BH
	0x01500061: 16,  // ar-DJ
	0x01500066: 17,  // ar-DZ
	0x0150006a: 18,  // ar-EG
	0x0150006b: 19,  // ar-EH
	0x0150006c: 20,  // ar-ER
	0x01500095: 21,  // ar-IL
	0x01500099: 22,  // ar-IQ
	0x0150009f: 23,  // ar-JO
	0x015000a6: 24,  // ar-KM
	0x015000aa: 25,  // ar-KW
	0x015000ae: 26,  // ar-LB
	0x015000b7: 27,  // ar-LY
	0x015000b8: 28,  // ar-MA
	0x015000c7: 29,  // ar-MR
	0x015000df: 30,  // ar-OM
	0x015000eb: 31,  // ar-PS
	0x015000f1: 32,  // ar-QA
	0x01500106: 33,  // ar-SA
	0x01500109: 34,  // ar-SD
	0x01500113: 35,  // ar-SO
	0x01500115: 36,  // ar-SS
	0x0150011a: 37,  // ar-SY
	0x0150011e: 38,  // ar-TD
	0x01500126: 39,  // ar-TN
	0x0150015b: 40,  // ar-YE
	0x01c00000: 41,  // as
	0x01c00097: 42,  // as-IN
	0x01d00000: 43,  // asa
	0x01d0012d: 44,  // asa-TZ
	0x01f00000: 45,  // ast
	0x01f0006d: 46,  // ast-ES
	0x02400000: 47,  // az
	0x0241d000: 48,  // az-Cyrl
	0x0241d031: 49,  // az-Cyrl-AZ
	0x0244f000: 50,  // az-Latn
	0x0244f031: 51,  // az-Latn-AZ
	0x02a00000: 52,  // bas
	0x02a00051: 53,  // bas-CM
	0x02f00000: 54,  // be
	0x02f00046: 55,  // be-BY
	0x03100000: 56,  // bem
	0x0310015f: 57,  // bem-ZM
	0x03300000: 58,  // bez
	0x0330012d: 59,  // bez-TZ
	0x03800000: 60,  // bg
	0x03800037: 61,  // bg-BG
	0x04900000: 62,  // bm
	0x049000c1: 63,  // bm-ML
	0x04b00000: 64,  // bn
	0x04b00034: 65,  // bn-BD
	0x04b00097: 66,  // bn-IN
	0x04c00000: 67,  // bo
	0x04c00052: 68,  // bo-CN
	0x04c00097: 69,  // bo-IN
	0x05000000: 70,  // br
	0x05000076: 71,  // br-FR
	0x05300000: 72,  // brx
	0x05300097: 73,  // brx-IN
	0x05400000: 74,  // bs
	0x0541d000: 75,  // bs-Cyrl
	0x0541d032: 76,  // bs-Cyrl-BA
	0x0544f000: 77,  // bs-Latn
	0x0544f032: 78,  // bs-Latn-BA
	0x06100000: 79,  // ca
	0x06100021: 80,  // ca-AD
	0x0610006d: 81,  // ca-ES
	0x06100076: 82,  // ca-FR
	0x0610009c: 83,  // ca-IT
	0x06400000: 84,  // ce
	0x06400104: 85,  // ce-RU
	0x06600000: 86,  // cgg
	0x0660012f: 87,  // cgg-UG
	0x06c00000: 88,  // chr
	0x06c00132: 89,  // chr-US
	0x06f00000: 90,  // ckb
	0x06f00099: 91,  // ckb-IQ
	0x06f0009a: 92,  // ckb-IR
	0x07900000: 93,  // cs
	0x0790005d: 94,  // cs-CZ
	0x07d00000: 95,  // cu
	0x07d00104: 96,  // cu-RU
	0x07f00000: 97,  // cy
	0x07f00079: 98,  // cy-GB
	0x08000000: 99,  // da
	0x08000062: 100, // da-DK
	0x08000080: 101, // da-GL
	0x08300000: 102, // dav
	0x083000a2: 103, // dav-KE
	0x08500000: 104, // de
	0x0850002d: 105, // de-AT
	0x08500035: 106, // de-BE
	0x0850004d: 107, // de-CH
	0x0850005f: 108, // de-DE
	0x085000b0: 109, // de-LI
	0x085000b5: 110, // de-LU
	0x08800000: 111, // dje
	0x088000d2: 112, // dje-NE
	0x08b00000: 113, // dsb
	0x08b0005f: 114, // dsb-DE
	0x08e00000: 115, // dua
	0x08e00051: 116, // dua-CM
	0x09000000: 117, // dyo
	0x09000112: 118, // dyo-SN
	0x09200000: 119, // dz
	0x09200042: 120, // dz-BT
	0x09300000: 121, // ebu
	0x093000a2: 122, // ebu-KE
	0x09400000: 123, // ee
	0x0940007e: 124, // ee-GH
	0x09400120: 125, // ee-TG
	0x09900000: 126, // el
	0x0990005c: 127, // el-CY
	0x09900085: 128, // el-GR
	0x09a00000: 129, // en
	0x09a00001: 130, // en-001
	0x09a0001a: 131, // en-150
	0x09a00024: 132, // en-AG
	0x09a00025: 133, // en-AI
	0x09a0002c: 134, // en-AS
	0x09a0002d: 135, // en-AT
	0x09a0002e: 136, // en-AU
	0x09a00033: 137, // en-BB
	0x09a00035: 138, // en-BE
	0x09a00039: 139, // en-BI
	0x09a0003c: 140, // en-BM
	0x09a00041: 141, // en-BS
	0x09a00045: 142, // en-BW
	0x09a00047: 143, // en-BZ
	0x09a00048: 144, // en-CA
	0x09a00049: 145, // en-CC
	0x09a0004d: 146, // en-CH
	0x09a0004f: 147, // en-CK
	0x09a00051: 148, // en-CM
	0x09a0005b: 149, // en-CX
	0x09a0005c: 150, // en-CY
	0x09a0005f: 151, // en-DE
	0x09a00060: 152, // en-DG
	0x09a00062: 153, // en-DK
	0x09a00063: 154, // en-DM
	0x09a0006c: 155, // en-ER
	0x09a00070: 156, // en-FI
	0x09a00071: 157, // en-FJ
	0x09a00072: 158, // en-FK
	0x09a00073: 159, // en-FM
	0x09a00079: 160, // en-GB
	0x09a0007a: 161, // en-GD
	0x09a0007d: 162, // en-GG
	0x09a0007e: 163, // en-GH
	0x09a0007f: 164, // en-GI
	0x09a00081: 165, // en-GM
	0x09a00088: 166, // en-GU
	0x09a0008a: 167, // en-GY
	0x09a0008b: 168, // en-HK
	0x09a00094: 169, // en-IE
	0x09a00095: 170, // en-IL
	0x09a00096: 171, // en-IM
	0x09a00097: 172, // en-IN
	0x09a00098: 173, // en-IO
	0x09a0009d: 174, // en-JE
	0x09a0009e: 175, // en-JM
	0x09a000a2: 176, // en-KE
	0x09a000a5: 177, // en-KI
	0x09a000a7: 178, // en-KN
	0x09a000ab: 179, // en-KY
	0x09a000af: 180, // en-LC
	0x09a000b2: 181, // en-LR
	0x09a000b3: 182, // en-LS
	0x09a000bd: 183, // en-MG
	0x09a000be: 184, // en-MH
	0x09a000c4: 185, // en-MO
	0x09a000c5: 186, // en-MP
	0x09a000c8: 187, // en-MS
	0x09a000c9: 188, // en-MT
	0x09a000ca: 189, // en-MU
	0x09a000cc: 190, // en-MW
	0x09a000ce: 191, // en-MY
	0x09a000d0: 192, // en-NA
	0x09a000d3: 193, // en-NF
	0x09a000d4: 194, // en-NG
	0x09a000d7: 195, // en-NL
	0x09a000db: 196, // en-NR
	0x09a000dd: 197, // en-NU
	0x09a000de: 198, // en-NZ
	0x09a000e4: 199, // en-PG
	0x09a000e5: 200, // en-PH
	0x09a000e6: 201, // en-PK
	0x09a000e9: 202, // en-PN
	0x09a000ea: 203, // en-PR
	0x09a000ee: 204, // en-PW
	0x09a00105: 205, // en-RW
	0x09a00107: 206, // en-SB
	0x09a00108: 207, // en-SC
	0x09a00109: 208, // en-SD
	0x09a0010a: 209, // en-SE
	0x09a0010b: 210, // en-SG
	0x09a0010c: 211, // en-SH
	0x09a0010d: 212, // en-SI
	0x09a00110: 213, // en-SL
	0x09a00115: 214, // en-SS
	0x09a00119: 215, // en-SX
	0x09a0011b: 216, // en-SZ
	0x09a0011d: 217, // en-TC
	0x09a00123: 218, // en-TK
	0x09a00127: 219, // en-TO
	0x09a0012a: 220, // en-TT
	0x09a0012b: 221, // en-TV
	0x09a0012d: 222, // en-TZ
	0x09a0012f: 223, // en-UG
	0x09a00131: 224, // en-UM
	0x09a00132: 225, // en-US
	0x09a00136: 226, // en-VC
	0x09a00139: 227, // en-VG
	0x09a0013a: 228, // en-VI
	0x09a0013c: 229, // en-VU
	0x09a0013f: 230, // en-WS
	0x09a0015e: 231, // en-ZA
	0x09a0015f: 232, // en-ZM
	0x09a00161: 233, // en-ZW
	0x09b00000: 234, // eo
	0x09b00001: 235, // eo-001
	0x09c00000: 236, // es
	0x09c00003: 237, // es-003
	0x09c0001e: 238, // es-419
	0x09c0002b: 239, // es-AR
	0x09c0003e: 240, // es-BO
	0x09c00050: 241, // es-CL
	0x09c00053: 242, // es-CO
	0x09c00055: 243, // es-CR
	0x09c00058: 244, // es-CU
	0x09c00064: 245, // es-DO
	0x09c00067: 246, // es-EA
	0x09c00068: 247, // es-EC
	0x09c0006d: 248, // es-ES
	0x09c00084: 249, // es-GQ
	0x09c00087: 250, // es-GT
	0x09c0008d: 251, // es-HN
	0x09c00092: 252, // es-IC
	0x09c000cd: 253, // es-MX
	0x09c000d6: 254, // es-NI
	0x09c000e0: 255, // es-PA
	0x09c000e2: 256, // es-PE
	0x09c000e5: 257, // es-PH
	0x09c000ea: 258, // es-PR
	0x09c000ef: 259, // es-PY
	0x09c00118: 260, // es-SV
	0x09c00132: 261, // es-US
	0x09c00133: 262, // es-UY
	0x09c00138: 263, // es-VE
	0x09e00000: 264, // et
	0x09e00069: 265, // et-EE
	0x0a000000: 266, // eu
	0x0a00006d: 267, // eu-ES
	0x0a100000: 268, // ewo
	0x0a100051: 269, // ewo-CM
	0x0a300000: 270, // fa
	0x0a300023: 271, // fa-AF
	0x0a30009a: 272, // fa-IR
	0x0a500000: 273, // ff
	0x0a500051: 274, // ff-CM
	0x0a500082: 275, // ff-GN
	0x0a5000c7: 276, // ff-MR
	0x0a500112: 277, // ff-SN
	0x0a700000: 278, // fi
	0x0a700070: 279, // fi-FI
	0x0a900000: 280, // fil
	0x0a9000e5: 281, // fil-PH
	0x0ac00000: 282, // fo
	0x0ac00062: 283, // fo-DK
	0x0ac00074: 284, // fo-FO
	0x0ae00000: 285, // fr
	0x0ae00035: 286, // fr-BE
	0x0ae00036: 287, // fr-BF
	0x0ae00039: 288, // fr-BI
	0x0ae0003a: 289, // fr-BJ
	0x0ae0003b: 290, // fr-BL
	0x0ae00048: 291, // fr-CA
	0x0ae0004a: 292, // fr-CD
	0x0ae0004b: 293, // fr-CF
	0x0ae0004c: 294, // fr-CG
	0x0ae0004d: 295, // fr-CH
	0x0ae0004e: 296, // fr-CI
	0x0ae00051: 297, // fr-CM
	0x0ae00061: 298, // fr-DJ
	0x0ae00066: 299, // fr-DZ
	0x0ae00076: 300, // fr-FR
	0x0ae00078: 301, // fr-GA
	0x0ae0007c: 302, // fr-GF
	0x0ae00082: 303, // fr-GN
	0x0ae00083: 304, // fr-GP
	0x0ae00084: 305, // fr-GQ
	0x0ae0008f: 306, // fr-HT
	0x0ae000a6: 307, // fr-KM
	0x0ae000b5: 308, // fr-LU
	0x0ae000b8: 309, // fr-MA
	0x0ae000b9: 310, // fr-MC
	0x0ae000bc: 311, // fr-MF
	0x0ae000bd: 312, // fr-MG
	0x0ae000c1: 313, // fr-ML
	0x0ae000c6: 314, // fr-MQ
	0x0ae000c7: 315, // fr-MR
	0x0ae000ca: 316, // fr-MU
	0x0ae000d1: 317, // fr-NC
	0x0ae000d2: 318, // fr-NE
	0x0ae000e3: 319, // fr-PF
	0x0ae000e8: 320, // fr-PM
	0x0ae00100: 321, // fr-RE
	0x0ae00105: 322, // fr-RW
	0x0ae00108: 323, // fr-SC
	0x0ae00112: 324, // fr-SN
	0x0ae0011a: 325, // fr-SY
	0x0ae0011e: 326, // fr-TD
	0x0ae00120: 327, // fr-TG
	0x0ae00126: 328, // fr-TN
	0x0ae0013c: 329, // fr-VU
	0x0ae0013d: 330, // fr-WF
	0x0ae0015c: 331, // fr-YT
	0x0b500000: 332, // fur
	0x0b50009c: 333, // fur-IT
	0x0b800000: 334, // fy
	0x0b8000d7: 335, // fy-NL
	0x0b900000: 336, // ga
	0x0b900094: 337, // ga-IE
	0x0c100000: 338, // gd
	0x0c100079: 339, // gd-GB
	0x0c700000: 340, // gl
	0x0c70006d: 341, // gl-ES
	0x0d100000: 342, // gsw
	0x0d10004d: 343, // gsw-CH
	0x0d100076: 344, // gsw-FR
	0x0d1000b0: 345, // gsw-LI
	0x0d200000: 346, // gu
	0x0d200097: 347, // gu-IN
	0x0d600000: 348, // guz
	0x0d6000a2: 349, // guz-KE
	0x0d700000: 350, // gv
	0x0d700096: 351, // gv-IM
	0x0da00000: 352, // ha
	0x0da0007e: 353, // ha-GH
	0x0da000d2: 354, // ha-NE
	0x0da000d4: 355, // ha-NG
	0x0dc00000: 356, // haw
	0x0dc00132: 357, // haw-US
	0x0de00000: 358, // he
	0x0de00095: 359, // he-IL
	0x0df00000: 360, // hi
	0x0df00097: 361, // hi-IN
	0x0ec00000: 362, // hr
	0x0ec00032: 363, // hr-BA
	0x0ec0008e: 364, // hr-HR
	0x0ed00000: 365, // hsb
	0x0ed0005f: 366, // hsb-DE
	0x0f000000: 367, // hu
	0x0f000090: 368, // hu-HU
	0x0f100000: 369, // hy
	0x0f100027: 370, // hy-AM
	0x0f600000: 371, // id
	0x0f600093: 372, // id-ID
	0x0f800000: 373, // ig
	0x0f8000d4: 374, // ig-NG
	0x0f900000: 375, // ii
	0x0f900052: 376, // ii-CN
	0x10000000: 377, // is
	0x1000009b: 378, // is-IS
	0x10100000: 379, // it
	0x1010004d: 380, // it-CH
	0x1010009c: 381, // it-IT
	0x10100111: 382, // it-SM
	0x10500000: 383, // ja
	0x105000a0: 384, // ja-JP
	0x10700000: 385, // jgo
	0x10700051: 386, // jgo-CM
	0x10900000: 387, // jmc
	0x1090012d: 388, // jmc-TZ
	0x10e00000: 389, // ka
	0x10e0007b: 390, // ka-GE
	0x11000000: 391, // kab
	0x11000066: 392, // kab-DZ
	0x11300000: 393, // kam
	0x113000a2: 394, // kam-KE
	0x11800000: 395, // kde
	0x1180012d: 396, // kde-TZ
	0x11a00000: 397, // kea
	0x11a00059: 398, // kea-CV
	0x12500000: 399, // khq
	0x125000c1: 400, // khq-ML
	0x12800000: 401, // ki
	0x128000a2: 402, // ki-KE
	0x12c00000: 403, // kk
	0x12c000ac: 404, // kk-KZ
	0x12d00000: 405, // kkj
	0x12d00051: 406, // kkj-CM
	0x12e00000: 407, // kl
	0x12e00080: 408, // kl-GL
	0x12f00000: 409, // kln
	0x12f000a2: 410, // kln-KE
	0x13000000: 411, // km
	0x130000a4: 412, // km-KH
	0x13200000: 413, // kn
	0x13200097: 414, // kn-IN
	0x13300000: 415, // ko
	0x133000a8: 416, // ko-KP
	0x133000a9: 417, // ko-KR
	0x13500000: 418, // kok
	0x13500097: 419, // kok-IN
	0x13e00000: 420, // ks
	0x13e00097: 421, // ks-IN
	0x13f00000: 422, // ksb
	0x13f0012d: 423, // ksb-TZ
	0x14000000: 424, // ksf
	0x14000051: 425, // ksf-CM
	0x14100000: 426, // ksh
	0x1410005f: 427, // ksh-DE
	0x14700000: 428, // kw
	0x14700079: 429, // kw-GB
	0x14a00000: 430, // ky
	0x14a000a3: 431, // ky-KG
	0x14e00000: 432, // lag
	0x14e0012d: 433, // lag-TZ
	0x15100000: 434, // lb
	0x151000b5: 435, // lb-LU
	0x15700000: 436, // lg
	0x1570012f: 437, // lg-UG
	0x15e00000: 438, // lkt
	0x15e00132: 439, // lkt-US
	0x16100000: 440, // ln
	0x16100029: 441, // ln-AO
	0x1610004a: 442, // ln-CD
	0x1610004b: 443, // ln-CF
	0x1610004c: 444, // ln-CG
	0x16200000: 445, // lo
	0x162000ad: 446, // lo-LA
	0x16500000: 447, // lrc
	0x16500099: 448, // lrc-IQ
	0x1650009a: 449, // lrc-IR
	0x16600000: 450, // lt
	0x166000b4: 451, // lt-LT
	0x16800000: 452, // lu
	0x1680004a: 453, // lu-CD
	0x16a00000: 454, // luo
	0x16a000a2: 455, // luo-KE
	0x16b00000: 456, // luy
	0x16b000a2: 457, // luy-KE
	0x16d00000: 458, // lv
	0x16d000b6: 459, // lv-LV
	0x17700000: 460, // mas
	0x177000a2: 461, // mas-KE
	0x1770012d: 462, // mas-TZ
	0x17d00000: 463, // mer
	0x17d000a2: 464, // mer-KE
	0x17f00000: 465, // mfe
	0x17f000ca: 466, // mfe-MU
	0x18000000: 467, // mg
	0x180000bd: 468, // mg-MG
	0x18100000: 469, // mgh
	0x181000cf: 470, // mgh-MZ
	0x18200000: 471, // mgo
	0x18200051: 472, // mgo-CM
	0x18900000: 473, // mk
	0x189000c0: 474, // mk-MK
	0x18a00000: 475, // ml
	0x18a00097: 476, // ml-IN
	0x18c00000: 477, // mn
	0x18c000c3: 478, // mn-MN
	0x19300000: 479, // mr
	0x19300097: 480, // mr-IN
	0x19700000: 481, // ms
	0x1970003d: 482, // ms-BN
	0x197000ce: 483, // ms-MY
	0x1970010b: 484, // ms-SG
	0x19800000: 485, // mt
	0x198000c9: 486, // mt-MT
	0x19a00000: 487, // mua
	0x19a00051: 488, // mua-CM
	0x1a200000: 489, // my
	0x1a2000c2: 490, // my-MM
	0x1a600000: 491, // mzn
	0x1a60009a: 492, // mzn-IR
	0x1aa00000: 493, // naq
	0x1aa000d0: 494, // naq-NA
	0x1ab00000: 495, // nb
	0x1ab000d8: 496, // nb-NO
	0x1ab0010e: 497, // nb-SJ
	0x1ad00000: 498, // nd
	0x1ad00161: 499, // nd-ZW
	0x1b000000: 500, // ne
	0x1b000097: 501, // ne-IN
	0x1b0000d9: 502, // ne-NP
	0x1b900000: 503, // nl
	0x1b90002f: 504, // nl-AW
	0x1b900035: 505, // nl-BE
	0x1b90003f: 506, // nl-BQ
	0x1b90005a: 507, // nl-CW
	0x1b9000d7: 508, // nl-NL
	0x1b900114: 509, // nl-SR
	0x1b900119: 510, // nl-SX
	0x1ba00000: 511, // nmg
	0x1ba00051: 512, // nmg-CM
	0x1bb00000: 513, // nn
	0x1bb000d8: 514, // nn-NO
	0x1bc00000: 515, // nnh
	0x1bc00051: 516, // nnh-CM
	0x1c500000: 517, // nus
	0x1c500115: 518, // nus-SS
	0x1ca00000: 519, // nyn
	0x1ca0012f: 520, // nyn-UG
	0x1ce00000: 521, // om
	0x1ce0006e: 522, // om-ET
	0x1ce000a2: 523, // om-KE
	0x1cf00000: 524, // or
	0x1cf00097: 525, // or-IN
	0x1d000000: 526, // os
	0x1d00007b: 527, // os-GE
	0x1d000104: 528, // os-RU
	0x1d200000: 529, // pa
	0x1d205000: 530, // pa-Arab
	0x1d2050e6: 531, // pa-Arab-PK
	0x1d22e000: 532, // pa-Guru
	0x1d22e097: 533, // pa-Guru-IN
	0x1e200000: 534, // pl
	0x1e2000e7: 535, // pl-PL
	0x1e800000: 536, // prg
	0x1e800001: 537, // prg-001
	0x1e900000: 538, // ps
	0x1e900023: 539, // ps-AF
	0x1ea00000: 540, // pt
	0x1ea00029: 541, // pt-AO
	0x1ea00040: 542, // pt-BR
	0x1ea00059: 543, // pt-CV
	0x1ea00089: 544, // pt-GW
	0x1ea000c4: 545, // pt-MO
	0x1ea000cf: 546, // pt-MZ
	0x1ea000ec: 547, // pt-PT
	0x1ea00116: 548, // pt-ST
	0x1ea00124: 549, // pt-TL
	0x1ec00000: 550, // qu
	0x1ec0003e: 551, // qu-BO
	0x1ec00068: 552, // qu-EC
	0x1ec000e2: 553, // qu-PE
	0x1f700000: 554, // rm
	0x1f70004d: 555, // rm-CH
	0x1fc00000: 556, // rn
	0x1fc00039: 557, // rn-BI
	0x1fe00000: 558, // ro
	0x1fe000ba: 559, // ro-MD
	0x1fe00102: 560, // ro-RO
	0x20000000: 561, // rof
	0x2000012d: 562, // rof-TZ
	0x20200000: 563, // ru
	0x20200046: 564, // ru-BY
	0x202000a3: 565, // ru-KG
	0x202000ac: 566, // ru-KZ
	0x202000ba: 567, // ru-MD
	0x20200104: 568, // ru-RU
	0x2020012e: 569, // ru-UA
	0x20500000: 570, // rw
	0x20500105: 571, // rw-RW
	0x20600000: 572, // rwk
	0x2060012d: 573, // rwk-TZ
	0x20a00000: 574, // sah
	0x20a00104: 575, // sah-RU
	0x20b00000: 576, // saq
	0x20b000a2: 577, // saq-KE
	0x20f00000: 578, // sbp
	0x20f0012d: 579, // sbp-TZ
	0x21800000: 580, // se
	0x21800070: 581, // se-FI
	0x218000d8: 582, // se-NO
	0x2180010a: 583, // se-SE
	0x21a00000: 584, // seh
	0x21a000cf: 585, // seh-MZ
	0x21c00000: 586, // ses
	0x21c000c1: 587, // ses-ML
	0x21d00000: 588, // sg
	0x21d0004b: 589, // sg-CF
	0x22100000: 590, // shi
	0x2214f000: 591, // shi-Latn
	0x2214f0b8: 592, // shi-Latn-MA
	0x221cc000: 593, // shi-Tfng
	0x221cc0b8: 594, // shi-Tfng-MA
	0x22300000: 595, // si
	0x223000b1: 596, // si-LK
	0x22500000: 597, // sk
	0x2250010f: 598, // sk-SK
	0x22700000: 599, // sl
	0x2270010d: 600, // sl-SI
	0x22d00000: 601, // smn
	0x22d00070: 602, // smn-FI
	0x23000000: 603, // sn
	0x23000161: 604, // sn-ZW
	0x23200000: 605, // so
	0x23200061: 606, // so-DJ
	0x2320006e: 607, // so-ET
	0x232000a2: 608, // so-KE
	0x23200113: 609, // so-SO
	0x23400000: 610, // sq
	0x23400026: 611, // sq-AL
	0x234000c0: 612, // sq-MK
	0x2340014a: 613, // sq-XK
	0x23500000: 614, // sr
	0x2351d000: 615, // sr-Cyrl
	0x2351d032: 616, // sr-Cyrl-BA
	0x2351d0bb: 617, // sr-Cyrl-ME
	0x2351d103: 618, // sr-Cyrl-RS
	0x2351d14a: 619, // sr-Cyrl-XK
	0x2354f000: 620, // sr-Latn
	0x2354f032: 621, // sr-Latn-BA
	0x2354f0bb: 622, // sr-Latn-ME
	0x2354f103: 623, // sr-Latn-RS
	0x2354f14a: 624, // sr-Latn-XK
	0x24100000: 625, // sv
	0x24100030: 626, // sv-AX
	0x24100070: 627, // sv-FI
	0x2410010a: 628, // sv-SE
	0x24200000: 629, // sw
	0x2420004a: 630, // sw-CD
	0x242000a2: 631, // sw-KE
	0x2420012d: 632, // sw-TZ
	0x2420012f: 633, // sw-UG
	0x24b00000: 634, // ta
	0x24b00097: 635, // ta-IN
	0x24b000b1: 636, // ta-LK
	0x24b000ce: 637, // ta-MY
	0x24b0010b: 638, // ta-SG
	0x25200000: 639, // te
	0x25200097: 640, // te-IN
	0x25400000: 641, // teo
	0x254000a2: 642, // teo-KE
	0x2540012f: 643, // teo-UG
	0x25700000: 644, // th
	0x25700121: 645, // th-TH
	0x25b00000: 646, // ti
	0x25b0006c: 647, // ti-ER
	0x25b0006e: 648, // ti-ET
	0x25e00000: 649, // tk
	0x25e00125: 650, // tk-TM
	0x26600000: 651, // to
	0x26600127: 652, // to-TO
	0x26900000: 653, // tr
	0x2690005c: 654, // tr-CY
	0x26900129: 655, // tr-TR
	0x27800000: 656, // twq
	0x278000d2: 657, // twq-NE
	0x27b00000: 658, // tzm
	0x27b000b8: 659, // tzm-MA
	0x27d00000: 660, // ug
	0x27d00052: 661, // ug-CN
	0x27f00000: 662, // uk
	0x27f0012e: 663, // uk-UA
	0x28500000: 664, // ur
	0x28500097: 665, // ur-IN
	0x285000e6: 666, // ur-PK
	0x28600000: 667, // uz
	0x28605000: 668, // uz-Arab
	0x28605023: 669, // uz-Arab-AF
	0x2861d000: 670, // uz-Cyrl
	0x2861d134: 671, // uz-Cyrl-UZ
	0x2864f000: 672, // uz-Latn
	0x2864f134: 673, // uz-Latn-UZ
	0x28700000: 674, // vai
	0x2874f000: 675, // vai-Latn
	0x2874f0b2: 676, // vai-Latn-LR
	0x287d3000: 677, // vai-Vaii
	0x287d30b2: 678, // vai-Vaii-LR
	0x28b00000: 679, // vi
	0x28b0013b: 680, // vi-VN
	0x29000000: 681, // vo
	0x29000001: 682, // vo-001
	0x29300000: 683, // vun
	0x2930012d: 684, // vun-TZ
	0x29500000: 685, // wae
	0x2950004d: 686, // wae-CH
	0x2aa00000: 687, // xog
	0x2aa0012f: 688, // xog-UG
	0x2b000000: 689, // yav
	0x2b000051: 690, // yav-CM
	0x2b200000: 691, // yi
	0x2b200001: 692, // yi-001
	0x2b300000: 693, // yo
	0x2b30003a: 694, // yo-BJ
	0x2b3000d4: 695, // yo-NG
	0x2bb00000: 696, // zgh
	0x2bb000b8: 697, // zgh-MA
	0x2bc00000: 698, // zh
	0x2bc32000: 699, // zh-Hans
	0x2bc32052: 700, // zh-Hans-CN
	0x2bc3208b: 701, // zh-Hans-HK
	0x2bc320c4: 702, // zh-Hans-MO
	0x2bc3210b: 703, // zh-Hans-SG
	0x2bc33000: 704, // zh-Hant
	0x2bc3308b: 705, // zh-Hant-HK
	0x2bc330c4: 706, // zh-Hant-MO
	0x2bc3312c: 707, // zh-Hant-TW
	0x2be00000: 708, // zu
	0x2be0015e: 709, // zu-ZA
}

// Total table size 4328 bytes (4KiB); checksum: C9659787
