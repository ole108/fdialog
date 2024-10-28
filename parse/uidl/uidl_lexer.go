// Code generated from UIDL.g4 by ANTLR 4.13.2. DO NOT EDIT.

package uidl

import (
	"fmt"
	"github.com/antlr4-go/antlr/v4"
	"sync"
	"unicode"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = sync.Once{}
var _ = unicode.IsLetter

type UIDLLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var UIDLLexerLexerStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	ChannelNames           []string
	ModeNames              []string
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func uidllexerLexerInit() {
	staticData := &UIDLLexerLexerStaticData
	staticData.ChannelNames = []string{
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
	}
	staticData.ModeNames = []string{
		"DEFAULT_MODE",
	}
	staticData.LiteralNames = []string{
		"", "'version'", "'v'", "'{'", "'}'", "'('", "')'", "'='",
	}
	staticData.SymbolicNames = []string{
		"", "", "", "", "", "", "", "", "Bool", "DoubleQuotedString", "BackQuotedString",
		"Identifier", "Natural", "Float", "Int", "Semicolon", "Comma", "WhiteSpace",
	}
	staticData.RuleNames = []string{
		"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "Bool", "DoubleQuotedString",
		"BackQuotedString", "EscapedChar", "UnicodeChar", "HexDigit", "SafeCodepoint",
		"Identifier", "Natural", "Float", "Int", "Exponent", "Semicolon", "Comma",
		"WhiteSpace", "Space", "Comment",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 0, 17, 189, 6, -1, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2,
		4, 7, 4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2,
		10, 7, 10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15,
		7, 15, 2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7,
		20, 2, 21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 2, 1, 2, 1, 3, 1, 3, 1, 4, 1, 4, 1,
		5, 1, 5, 1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 3, 7, 79, 8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 84, 8, 8, 10, 8, 12, 8, 87,
		9, 8, 1, 8, 1, 8, 1, 9, 1, 9, 5, 9, 93, 8, 9, 10, 9, 12, 9, 96, 9, 9, 1,
		9, 1, 9, 1, 10, 1, 10, 1, 10, 3, 10, 103, 8, 10, 1, 11, 1, 11, 1, 11, 1,
		11, 1, 11, 1, 11, 1, 12, 1, 12, 1, 13, 1, 13, 1, 14, 1, 14, 5, 14, 117,
		8, 14, 10, 14, 12, 14, 120, 9, 14, 1, 15, 1, 15, 5, 15, 124, 8, 15, 10,
		15, 12, 15, 127, 9, 15, 1, 16, 1, 16, 1, 16, 4, 16, 132, 8, 16, 11, 16,
		12, 16, 133, 1, 16, 3, 16, 137, 8, 16, 1, 17, 3, 17, 140, 8, 17, 1, 17,
		1, 17, 3, 17, 144, 8, 17, 1, 18, 1, 18, 3, 18, 148, 8, 18, 1, 18, 4, 18,
		151, 8, 18, 11, 18, 12, 18, 152, 1, 19, 3, 19, 156, 8, 19, 1, 19, 1, 19,
		3, 19, 160, 8, 19, 1, 20, 3, 20, 163, 8, 20, 1, 20, 1, 20, 3, 20, 167,
		8, 20, 1, 21, 1, 21, 4, 21, 171, 8, 21, 11, 21, 12, 21, 172, 1, 22, 4,
		22, 176, 8, 22, 11, 22, 12, 22, 177, 1, 23, 1, 23, 5, 23, 182, 8, 23, 10,
		23, 12, 23, 185, 9, 23, 1, 23, 3, 23, 188, 8, 23, 0, 0, 24, 1, 1, 3, 2,
		5, 3, 7, 4, 9, 5, 11, 6, 13, 7, 15, 8, 17, 9, 19, 10, 21, 0, 23, 0, 25,
		0, 27, 0, 29, 11, 31, 12, 33, 13, 35, 14, 37, 0, 39, 15, 41, 16, 43, 17,
		45, 0, 47, 0, 1, 0, 13, 1, 0, 96, 96, 7, 0, 34, 34, 92, 92, 98, 98, 102,
		102, 110, 110, 114, 114, 116, 116, 3, 0, 48, 57, 65, 70, 97, 102, 3, 0,
		0, 31, 34, 34, 92, 92, 733, 0, 65, 90, 95, 95, 97, 122, 170, 170, 181,
		181, 186, 186, 192, 214, 216, 246, 248, 705, 710, 721, 736, 740, 748, 748,
		750, 750, 837, 837, 880, 884, 886, 887, 890, 893, 895, 895, 902, 902, 904,
		906, 908, 908, 910, 929, 931, 1013, 1015, 1153, 1162, 1327, 1329, 1366,
		1369, 1369, 1376, 1416, 1456, 1469, 1471, 1471, 1473, 1474, 1476, 1477,
		1479, 1479, 1488, 1514, 1519, 1522, 1552, 1562, 1568, 1623, 1625, 1631,
		1646, 1747, 1749, 1756, 1761, 1768, 1773, 1775, 1786, 1788, 1791, 1791,
		1808, 1855, 1869, 1969, 1994, 2026, 2036, 2037, 2042, 2042, 2048, 2071,
		2074, 2092, 2112, 2136, 2144, 2154, 2160, 2183, 2185, 2190, 2208, 2249,
		2260, 2271, 2275, 2281, 2288, 2363, 2365, 2380, 2382, 2384, 2389, 2403,
		2417, 2435, 2437, 2444, 2447, 2448, 2451, 2472, 2474, 2480, 2482, 2482,
		2486, 2489, 2493, 2500, 2503, 2504, 2507, 2508, 2510, 2510, 2519, 2519,
		2524, 2525, 2527, 2531, 2544, 2545, 2556, 2556, 2561, 2563, 2565, 2570,
		2575, 2576, 2579, 2600, 2602, 2608, 2610, 2611, 2613, 2614, 2616, 2617,
		2622, 2626, 2631, 2632, 2635, 2636, 2641, 2641, 2649, 2652, 2654, 2654,
		2672, 2677, 2689, 2691, 2693, 2701, 2703, 2705, 2707, 2728, 2730, 2736,
		2738, 2739, 2741, 2745, 2749, 2757, 2759, 2761, 2763, 2764, 2768, 2768,
		2784, 2787, 2809, 2812, 2817, 2819, 2821, 2828, 2831, 2832, 2835, 2856,
		2858, 2864, 2866, 2867, 2869, 2873, 2877, 2884, 2887, 2888, 2891, 2892,
		2902, 2903, 2908, 2909, 2911, 2915, 2929, 2929, 2946, 2947, 2949, 2954,
		2958, 2960, 2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975, 2979, 2980,
		2984, 2986, 2990, 3001, 3006, 3010, 3014, 3016, 3018, 3020, 3024, 3024,
		3031, 3031, 3072, 3084, 3086, 3088, 3090, 3112, 3114, 3129, 3133, 3140,
		3142, 3144, 3146, 3148, 3157, 3158, 3160, 3162, 3165, 3165, 3168, 3171,
		3200, 3203, 3205, 3212, 3214, 3216, 3218, 3240, 3242, 3251, 3253, 3257,
		3261, 3268, 3270, 3272, 3274, 3276, 3285, 3286, 3293, 3294, 3296, 3299,
		3313, 3315, 3328, 3340, 3342, 3344, 3346, 3386, 3389, 3396, 3398, 3400,
		3402, 3404, 3406, 3406, 3412, 3415, 3423, 3427, 3450, 3455, 3457, 3459,
		3461, 3478, 3482, 3505, 3507, 3515, 3517, 3517, 3520, 3526, 3535, 3540,
		3542, 3542, 3544, 3551, 3570, 3571, 3585, 3642, 3648, 3654, 3661, 3661,
		3713, 3714, 3716, 3716, 3718, 3722, 3724, 3747, 3749, 3749, 3751, 3769,
		3771, 3773, 3776, 3780, 3782, 3782, 3789, 3789, 3804, 3807, 3840, 3840,
		3904, 3911, 3913, 3948, 3953, 3971, 3976, 3991, 3993, 4028, 4096, 4150,
		4152, 4152, 4155, 4159, 4176, 4239, 4250, 4253, 4256, 4293, 4295, 4295,
		4301, 4301, 4304, 4346, 4348, 4680, 4682, 4685, 4688, 4694, 4696, 4696,
		4698, 4701, 4704, 4744, 4746, 4749, 4752, 4784, 4786, 4789, 4792, 4798,
		4800, 4800, 4802, 4805, 4808, 4822, 4824, 4880, 4882, 4885, 4888, 4954,
		4992, 5007, 5024, 5109, 5112, 5117, 5121, 5740, 5743, 5759, 5761, 5786,
		5792, 5866, 5870, 5880, 5888, 5907, 5919, 5939, 5952, 5971, 5984, 5996,
		5998, 6000, 6002, 6003, 6016, 6067, 6070, 6088, 6103, 6103, 6108, 6108,
		6176, 6264, 6272, 6314, 6320, 6389, 6400, 6430, 6432, 6443, 6448, 6456,
		6480, 6509, 6512, 6516, 6528, 6571, 6576, 6601, 6656, 6683, 6688, 6750,
		6753, 6772, 6823, 6823, 6847, 6848, 6860, 6862, 6912, 6963, 6965, 6979,
		6981, 6988, 7040, 7081, 7084, 7087, 7098, 7141, 7143, 7153, 7168, 7222,
		7245, 7247, 7258, 7293, 7296, 7304, 7312, 7354, 7357, 7359, 7401, 7404,
		7406, 7411, 7413, 7414, 7418, 7418, 7424, 7615, 7655, 7668, 7680, 7957,
		7960, 7965, 7968, 8005, 8008, 8013, 8016, 8023, 8025, 8025, 8027, 8027,
		8029, 8029, 8031, 8061, 8064, 8116, 8118, 8124, 8126, 8126, 8130, 8132,
		8134, 8140, 8144, 8147, 8150, 8155, 8160, 8172, 8178, 8180, 8182, 8188,
		8305, 8305, 8319, 8319, 8336, 8348, 8450, 8450, 8455, 8455, 8458, 8467,
		8469, 8469, 8473, 8477, 8484, 8484, 8486, 8486, 8488, 8488, 8490, 8493,
		8495, 8505, 8508, 8511, 8517, 8521, 8526, 8526, 8544, 8584, 9398, 9449,
		11264, 11492, 11499, 11502, 11506, 11507, 11520, 11557, 11559, 11559, 11565,
		11565, 11568, 11623, 11631, 11631, 11648, 11670, 11680, 11686, 11688, 11694,
		11696, 11702, 11704, 11710, 11712, 11718, 11720, 11726, 11728, 11734, 11736,
		11742, 11744, 11775, 11823, 11823, 12293, 12295, 12321, 12329, 12337, 12341,
		12344, 12348, 12353, 12438, 12445, 12447, 12449, 12538, 12540, 12543, 12549,
		12591, 12593, 12686, 12704, 12735, 12784, 12799, 13312, 19903, 19968, 42124,
		42192, 42237, 42240, 42508, 42512, 42527, 42538, 42539, 42560, 42606, 42612,
		42619, 42623, 42735, 42775, 42783, 42786, 42888, 42891, 42954, 42960, 42961,
		42963, 42963, 42965, 42969, 42994, 43013, 43015, 43047, 43072, 43123, 43136,
		43203, 43205, 43205, 43250, 43255, 43259, 43259, 43261, 43263, 43274, 43306,
		43312, 43346, 43360, 43388, 43392, 43442, 43444, 43455, 43471, 43471, 43488,
		43503, 43514, 43518, 43520, 43574, 43584, 43597, 43616, 43638, 43642, 43710,
		43712, 43712, 43714, 43714, 43739, 43741, 43744, 43759, 43762, 43765, 43777,
		43782, 43785, 43790, 43793, 43798, 43808, 43814, 43816, 43822, 43824, 43866,
		43868, 43881, 43888, 44010, 44032, 55203, 55216, 55238, 55243, 55291, 63744,
		64109, 64112, 64217, 64256, 64262, 64275, 64279, 64285, 64296, 64298, 64310,
		64312, 64316, 64318, 64318, 64320, 64321, 64323, 64324, 64326, 64433, 64467,
		64829, 64848, 64911, 64914, 64967, 65008, 65019, 65136, 65140, 65142, 65276,
		65313, 65338, 65345, 65370, 65382, 65470, 65474, 65479, 65482, 65487, 65490,
		65495, 65498, 65500, 65536, 65547, 65549, 65574, 65576, 65594, 65596, 65597,
		65599, 65613, 65616, 65629, 65664, 65786, 65856, 65908, 66176, 66204, 66208,
		66256, 66304, 66335, 66349, 66378, 66384, 66426, 66432, 66461, 66464, 66499,
		66504, 66511, 66513, 66517, 66560, 66717, 66736, 66771, 66776, 66811, 66816,
		66855, 66864, 66915, 66928, 66938, 66940, 66954, 66956, 66962, 66964, 66965,
		66967, 66977, 66979, 66993, 66995, 67001, 67003, 67004, 67072, 67382, 67392,
		67413, 67424, 67431, 67456, 67461, 67463, 67504, 67506, 67514, 67584, 67589,
		67592, 67592, 67594, 67637, 67639, 67640, 67644, 67644, 67647, 67669, 67680,
		67702, 67712, 67742, 67808, 67826, 67828, 67829, 67840, 67861, 67872, 67897,
		67968, 68023, 68030, 68031, 68096, 68099, 68101, 68102, 68108, 68115, 68117,
		68119, 68121, 68149, 68192, 68220, 68224, 68252, 68288, 68295, 68297, 68324,
		68352, 68405, 68416, 68437, 68448, 68466, 68480, 68497, 68608, 68680, 68736,
		68786, 68800, 68850, 68864, 68903, 69248, 69289, 69291, 69292, 69296, 69297,
		69376, 69404, 69415, 69415, 69424, 69445, 69488, 69505, 69552, 69572, 69600,
		69622, 69632, 69701, 69745, 69749, 69760, 69816, 69826, 69826, 69840, 69864,
		69888, 69938, 69956, 69959, 69968, 70002, 70006, 70006, 70016, 70079, 70081,
		70084, 70094, 70095, 70106, 70106, 70108, 70108, 70144, 70161, 70163, 70196,
		70199, 70199, 70206, 70209, 70272, 70278, 70280, 70280, 70282, 70285, 70287,
		70301, 70303, 70312, 70320, 70376, 70400, 70403, 70405, 70412, 70415, 70416,
		70419, 70440, 70442, 70448, 70450, 70451, 70453, 70457, 70461, 70468, 70471,
		70472, 70475, 70476, 70480, 70480, 70487, 70487, 70493, 70499, 70656, 70721,
		70723, 70725, 70727, 70730, 70751, 70753, 70784, 70849, 70852, 70853, 70855,
		70855, 71040, 71093, 71096, 71102, 71128, 71133, 71168, 71230, 71232, 71232,
		71236, 71236, 71296, 71349, 71352, 71352, 71424, 71450, 71453, 71466, 71488,
		71494, 71680, 71736, 71840, 71903, 71935, 71942, 71945, 71945, 71948, 71955,
		71957, 71958, 71960, 71989, 71991, 71992, 71995, 71996, 71999, 72002, 72096,
		72103, 72106, 72151, 72154, 72159, 72161, 72161, 72163, 72164, 72192, 72242,
		72245, 72254, 72272, 72343, 72349, 72349, 72368, 72440, 72704, 72712, 72714,
		72758, 72760, 72766, 72768, 72768, 72818, 72847, 72850, 72871, 72873, 72886,
		72960, 72966, 72968, 72969, 72971, 73014, 73018, 73018, 73020, 73021, 73023,
		73025, 73027, 73027, 73030, 73031, 73056, 73061, 73063, 73064, 73066, 73102,
		73104, 73105, 73107, 73110, 73112, 73112, 73440, 73462, 73472, 73488, 73490,
		73530, 73534, 73536, 73648, 73648, 73728, 74649, 74752, 74862, 74880, 75075,
		77712, 77808, 77824, 78895, 78913, 78918, 82944, 83526, 92160, 92728, 92736,
		92766, 92784, 92862, 92880, 92909, 92928, 92975, 92992, 92995, 93027, 93047,
		93053, 93071, 93760, 93823, 93952, 94026, 94031, 94087, 94095, 94111, 94176,
		94177, 94179, 94179, 94192, 94193, 94208, 100343, 100352, 101589, 101632,
		101640, 110576, 110579, 110581, 110587, 110589, 110590, 110592, 110882,
		110898, 110898, 110928, 110930, 110933, 110933, 110948, 110951, 110960,
		111355, 113664, 113770, 113776, 113788, 113792, 113800, 113808, 113817,
		113822, 113822, 119808, 119892, 119894, 119964, 119966, 119967, 119970,
		119970, 119973, 119974, 119977, 119980, 119982, 119993, 119995, 119995,
		119997, 120003, 120005, 120069, 120071, 120074, 120077, 120084, 120086,
		120092, 120094, 120121, 120123, 120126, 120128, 120132, 120134, 120134,
		120138, 120144, 120146, 120485, 120488, 120512, 120514, 120538, 120540,
		120570, 120572, 120596, 120598, 120628, 120630, 120654, 120656, 120686,
		120688, 120712, 120714, 120744, 120746, 120770, 120772, 120779, 122624,
		122654, 122661, 122666, 122880, 122886, 122888, 122904, 122907, 122913,
		122915, 122916, 122918, 122922, 122928, 122989, 123023, 123023, 123136,
		123180, 123191, 123197, 123214, 123214, 123536, 123565, 123584, 123627,
		124112, 124139, 124896, 124902, 124904, 124907, 124909, 124910, 124912,
		124926, 124928, 125124, 125184, 125251, 125255, 125255, 125259, 125259,
		126464, 126467, 126469, 126495, 126497, 126498, 126500, 126500, 126503,
		126503, 126505, 126514, 126516, 126519, 126521, 126521, 126523, 126523,
		126530, 126530, 126535, 126535, 126537, 126537, 126539, 126539, 126541,
		126543, 126545, 126546, 126548, 126548, 126551, 126551, 126553, 126553,
		126555, 126555, 126557, 126557, 126559, 126559, 126561, 126562, 126564,
		126564, 126567, 126570, 126572, 126578, 126580, 126583, 126585, 126588,
		126590, 126590, 126592, 126601, 126603, 126619, 126625, 126627, 126629,
		126633, 126635, 126651, 127280, 127305, 127312, 127337, 127344, 127369,
		131072, 173791, 173824, 177977, 177984, 178205, 178208, 183969, 183984,
		191456, 194560, 195101, 196608, 201546, 201552, 205743, 773, 0, 48, 57,
		65, 90, 95, 95, 97, 122, 170, 170, 181, 181, 186, 186, 192, 214, 216, 246,
		248, 705, 710, 721, 736, 740, 748, 748, 750, 750, 837, 837, 880, 884, 886,
		887, 890, 893, 895, 895, 902, 902, 904, 906, 908, 908, 910, 929, 931, 1013,
		1015, 1153, 1162, 1327, 1329, 1366, 1369, 1369, 1376, 1416, 1456, 1469,
		1471, 1471, 1473, 1474, 1476, 1477, 1479, 1479, 1488, 1514, 1519, 1522,
		1552, 1562, 1568, 1623, 1625, 1641, 1646, 1747, 1749, 1756, 1761, 1768,
		1773, 1788, 1791, 1791, 1808, 1855, 1869, 1969, 1984, 2026, 2036, 2037,
		2042, 2042, 2048, 2071, 2074, 2092, 2112, 2136, 2144, 2154, 2160, 2183,
		2185, 2190, 2208, 2249, 2260, 2271, 2275, 2281, 2288, 2363, 2365, 2380,
		2382, 2384, 2389, 2403, 2406, 2415, 2417, 2435, 2437, 2444, 2447, 2448,
		2451, 2472, 2474, 2480, 2482, 2482, 2486, 2489, 2493, 2500, 2503, 2504,
		2507, 2508, 2510, 2510, 2519, 2519, 2524, 2525, 2527, 2531, 2534, 2545,
		2556, 2556, 2561, 2563, 2565, 2570, 2575, 2576, 2579, 2600, 2602, 2608,
		2610, 2611, 2613, 2614, 2616, 2617, 2622, 2626, 2631, 2632, 2635, 2636,
		2641, 2641, 2649, 2652, 2654, 2654, 2662, 2677, 2689, 2691, 2693, 2701,
		2703, 2705, 2707, 2728, 2730, 2736, 2738, 2739, 2741, 2745, 2749, 2757,
		2759, 2761, 2763, 2764, 2768, 2768, 2784, 2787, 2790, 2799, 2809, 2812,
		2817, 2819, 2821, 2828, 2831, 2832, 2835, 2856, 2858, 2864, 2866, 2867,
		2869, 2873, 2877, 2884, 2887, 2888, 2891, 2892, 2902, 2903, 2908, 2909,
		2911, 2915, 2918, 2927, 2929, 2929, 2946, 2947, 2949, 2954, 2958, 2960,
		2962, 2965, 2969, 2970, 2972, 2972, 2974, 2975, 2979, 2980, 2984, 2986,
		2990, 3001, 3006, 3010, 3014, 3016, 3018, 3020, 3024, 3024, 3031, 3031,
		3046, 3055, 3072, 3084, 3086, 3088, 3090, 3112, 3114, 3129, 3133, 3140,
		3142, 3144, 3146, 3148, 3157, 3158, 3160, 3162, 3165, 3165, 3168, 3171,
		3174, 3183, 3200, 3203, 3205, 3212, 3214, 3216, 3218, 3240, 3242, 3251,
		3253, 3257, 3261, 3268, 3270, 3272, 3274, 3276, 3285, 3286, 3293, 3294,
		3296, 3299, 3302, 3311, 3313, 3315, 3328, 3340, 3342, 3344, 3346, 3386,
		3389, 3396, 3398, 3400, 3402, 3404, 3406, 3406, 3412, 3415, 3423, 3427,
		3430, 3439, 3450, 3455, 3457, 3459, 3461, 3478, 3482, 3505, 3507, 3515,
		3517, 3517, 3520, 3526, 3535, 3540, 3542, 3542, 3544, 3551, 3558, 3567,
		3570, 3571, 3585, 3642, 3648, 3654, 3661, 3661, 3664, 3673, 3713, 3714,
		3716, 3716, 3718, 3722, 3724, 3747, 3749, 3749, 3751, 3769, 3771, 3773,
		3776, 3780, 3782, 3782, 3789, 3789, 3792, 3801, 3804, 3807, 3840, 3840,
		3872, 3881, 3904, 3911, 3913, 3948, 3953, 3971, 3976, 3991, 3993, 4028,
		4096, 4150, 4152, 4152, 4155, 4169, 4176, 4253, 4256, 4293, 4295, 4295,
		4301, 4301, 4304, 4346, 4348, 4680, 4682, 4685, 4688, 4694, 4696, 4696,
		4698, 4701, 4704, 4744, 4746, 4749, 4752, 4784, 4786, 4789, 4792, 4798,
		4800, 4800, 4802, 4805, 4808, 4822, 4824, 4880, 4882, 4885, 4888, 4954,
		4992, 5007, 5024, 5109, 5112, 5117, 5121, 5740, 5743, 5759, 5761, 5786,
		5792, 5866, 5870, 5880, 5888, 5907, 5919, 5939, 5952, 5971, 5984, 5996,
		5998, 6000, 6002, 6003, 6016, 6067, 6070, 6088, 6103, 6103, 6108, 6108,
		6112, 6121, 6160, 6169, 6176, 6264, 6272, 6314, 6320, 6389, 6400, 6430,
		6432, 6443, 6448, 6456, 6470, 6509, 6512, 6516, 6528, 6571, 6576, 6601,
		6608, 6617, 6656, 6683, 6688, 6750, 6753, 6772, 6784, 6793, 6800, 6809,
		6823, 6823, 6847, 6848, 6860, 6862, 6912, 6963, 6965, 6979, 6981, 6988,
		6992, 7001, 7040, 7081, 7084, 7141, 7143, 7153, 7168, 7222, 7232, 7241,
		7245, 7293, 7296, 7304, 7312, 7354, 7357, 7359, 7401, 7404, 7406, 7411,
		7413, 7414, 7418, 7418, 7424, 7615, 7655, 7668, 7680, 7957, 7960, 7965,
		7968, 8005, 8008, 8013, 8016, 8023, 8025, 8025, 8027, 8027, 8029, 8029,
		8031, 8061, 8064, 8116, 8118, 8124, 8126, 8126, 8130, 8132, 8134, 8140,
		8144, 8147, 8150, 8155, 8160, 8172, 8178, 8180, 8182, 8188, 8305, 8305,
		8319, 8319, 8336, 8348, 8450, 8450, 8455, 8455, 8458, 8467, 8469, 8469,
		8473, 8477, 8484, 8484, 8486, 8486, 8488, 8488, 8490, 8493, 8495, 8505,
		8508, 8511, 8517, 8521, 8526, 8526, 8544, 8584, 9398, 9449, 11264, 11492,
		11499, 11502, 11506, 11507, 11520, 11557, 11559, 11559, 11565, 11565, 11568,
		11623, 11631, 11631, 11648, 11670, 11680, 11686, 11688, 11694, 11696, 11702,
		11704, 11710, 11712, 11718, 11720, 11726, 11728, 11734, 11736, 11742, 11744,
		11775, 11823, 11823, 12293, 12295, 12321, 12329, 12337, 12341, 12344, 12348,
		12353, 12438, 12445, 12447, 12449, 12538, 12540, 12543, 12549, 12591, 12593,
		12686, 12704, 12735, 12784, 12799, 13312, 19903, 19968, 42124, 42192, 42237,
		42240, 42508, 42512, 42539, 42560, 42606, 42612, 42619, 42623, 42735, 42775,
		42783, 42786, 42888, 42891, 42954, 42960, 42961, 42963, 42963, 42965, 42969,
		42994, 43013, 43015, 43047, 43072, 43123, 43136, 43203, 43205, 43205, 43216,
		43225, 43250, 43255, 43259, 43259, 43261, 43306, 43312, 43346, 43360, 43388,
		43392, 43442, 43444, 43455, 43471, 43481, 43488, 43518, 43520, 43574, 43584,
		43597, 43600, 43609, 43616, 43638, 43642, 43710, 43712, 43712, 43714, 43714,
		43739, 43741, 43744, 43759, 43762, 43765, 43777, 43782, 43785, 43790, 43793,
		43798, 43808, 43814, 43816, 43822, 43824, 43866, 43868, 43881, 43888, 44010,
		44016, 44025, 44032, 55203, 55216, 55238, 55243, 55291, 63744, 64109, 64112,
		64217, 64256, 64262, 64275, 64279, 64285, 64296, 64298, 64310, 64312, 64316,
		64318, 64318, 64320, 64321, 64323, 64324, 64326, 64433, 64467, 64829, 64848,
		64911, 64914, 64967, 65008, 65019, 65136, 65140, 65142, 65276, 65296, 65305,
		65313, 65338, 65345, 65370, 65382, 65470, 65474, 65479, 65482, 65487, 65490,
		65495, 65498, 65500, 65536, 65547, 65549, 65574, 65576, 65594, 65596, 65597,
		65599, 65613, 65616, 65629, 65664, 65786, 65856, 65908, 66176, 66204, 66208,
		66256, 66304, 66335, 66349, 66378, 66384, 66426, 66432, 66461, 66464, 66499,
		66504, 66511, 66513, 66517, 66560, 66717, 66720, 66729, 66736, 66771, 66776,
		66811, 66816, 66855, 66864, 66915, 66928, 66938, 66940, 66954, 66956, 66962,
		66964, 66965, 66967, 66977, 66979, 66993, 66995, 67001, 67003, 67004, 67072,
		67382, 67392, 67413, 67424, 67431, 67456, 67461, 67463, 67504, 67506, 67514,
		67584, 67589, 67592, 67592, 67594, 67637, 67639, 67640, 67644, 67644, 67647,
		67669, 67680, 67702, 67712, 67742, 67808, 67826, 67828, 67829, 67840, 67861,
		67872, 67897, 67968, 68023, 68030, 68031, 68096, 68099, 68101, 68102, 68108,
		68115, 68117, 68119, 68121, 68149, 68192, 68220, 68224, 68252, 68288, 68295,
		68297, 68324, 68352, 68405, 68416, 68437, 68448, 68466, 68480, 68497, 68608,
		68680, 68736, 68786, 68800, 68850, 68864, 68903, 68912, 68921, 69248, 69289,
		69291, 69292, 69296, 69297, 69376, 69404, 69415, 69415, 69424, 69445, 69488,
		69505, 69552, 69572, 69600, 69622, 69632, 69701, 69734, 69743, 69745, 69749,
		69760, 69816, 69826, 69826, 69840, 69864, 69872, 69881, 69888, 69938, 69942,
		69951, 69956, 69959, 69968, 70002, 70006, 70006, 70016, 70079, 70081, 70084,
		70094, 70106, 70108, 70108, 70144, 70161, 70163, 70196, 70199, 70199, 70206,
		70209, 70272, 70278, 70280, 70280, 70282, 70285, 70287, 70301, 70303, 70312,
		70320, 70376, 70384, 70393, 70400, 70403, 70405, 70412, 70415, 70416, 70419,
		70440, 70442, 70448, 70450, 70451, 70453, 70457, 70461, 70468, 70471, 70472,
		70475, 70476, 70480, 70480, 70487, 70487, 70493, 70499, 70656, 70721, 70723,
		70725, 70727, 70730, 70736, 70745, 70751, 70753, 70784, 70849, 70852, 70853,
		70855, 70855, 70864, 70873, 71040, 71093, 71096, 71102, 71128, 71133, 71168,
		71230, 71232, 71232, 71236, 71236, 71248, 71257, 71296, 71349, 71352, 71352,
		71360, 71369, 71424, 71450, 71453, 71466, 71472, 71481, 71488, 71494, 71680,
		71736, 71840, 71913, 71935, 71942, 71945, 71945, 71948, 71955, 71957, 71958,
		71960, 71989, 71991, 71992, 71995, 71996, 71999, 72002, 72016, 72025, 72096,
		72103, 72106, 72151, 72154, 72159, 72161, 72161, 72163, 72164, 72192, 72242,
		72245, 72254, 72272, 72343, 72349, 72349, 72368, 72440, 72704, 72712, 72714,
		72758, 72760, 72766, 72768, 72768, 72784, 72793, 72818, 72847, 72850, 72871,
		72873, 72886, 72960, 72966, 72968, 72969, 72971, 73014, 73018, 73018, 73020,
		73021, 73023, 73025, 73027, 73027, 73030, 73031, 73040, 73049, 73056, 73061,
		73063, 73064, 73066, 73102, 73104, 73105, 73107, 73110, 73112, 73112, 73120,
		73129, 73440, 73462, 73472, 73488, 73490, 73530, 73534, 73536, 73552, 73561,
		73648, 73648, 73728, 74649, 74752, 74862, 74880, 75075, 77712, 77808, 77824,
		78895, 78913, 78918, 82944, 83526, 92160, 92728, 92736, 92766, 92768, 92777,
		92784, 92862, 92864, 92873, 92880, 92909, 92928, 92975, 92992, 92995, 93008,
		93017, 93027, 93047, 93053, 93071, 93760, 93823, 93952, 94026, 94031, 94087,
		94095, 94111, 94176, 94177, 94179, 94179, 94192, 94193, 94208, 100343,
		100352, 101589, 101632, 101640, 110576, 110579, 110581, 110587, 110589,
		110590, 110592, 110882, 110898, 110898, 110928, 110930, 110933, 110933,
		110948, 110951, 110960, 111355, 113664, 113770, 113776, 113788, 113792,
		113800, 113808, 113817, 113822, 113822, 119808, 119892, 119894, 119964,
		119966, 119967, 119970, 119970, 119973, 119974, 119977, 119980, 119982,
		119993, 119995, 119995, 119997, 120003, 120005, 120069, 120071, 120074,
		120077, 120084, 120086, 120092, 120094, 120121, 120123, 120126, 120128,
		120132, 120134, 120134, 120138, 120144, 120146, 120485, 120488, 120512,
		120514, 120538, 120540, 120570, 120572, 120596, 120598, 120628, 120630,
		120654, 120656, 120686, 120688, 120712, 120714, 120744, 120746, 120770,
		120772, 120779, 120782, 120831, 122624, 122654, 122661, 122666, 122880,
		122886, 122888, 122904, 122907, 122913, 122915, 122916, 122918, 122922,
		122928, 122989, 123023, 123023, 123136, 123180, 123191, 123197, 123200,
		123209, 123214, 123214, 123536, 123565, 123584, 123627, 123632, 123641,
		124112, 124139, 124144, 124153, 124896, 124902, 124904, 124907, 124909,
		124910, 124912, 124926, 124928, 125124, 125184, 125251, 125255, 125255,
		125259, 125259, 125264, 125273, 126464, 126467, 126469, 126495, 126497,
		126498, 126500, 126500, 126503, 126503, 126505, 126514, 126516, 126519,
		126521, 126521, 126523, 126523, 126530, 126530, 126535, 126535, 126537,
		126537, 126539, 126539, 126541, 126543, 126545, 126546, 126548, 126548,
		126551, 126551, 126553, 126553, 126555, 126555, 126557, 126557, 126559,
		126559, 126561, 126562, 126564, 126564, 126567, 126570, 126572, 126578,
		126580, 126583, 126585, 126588, 126590, 126590, 126592, 126601, 126603,
		126619, 126625, 126627, 126629, 126633, 126635, 126651, 127280, 127305,
		127312, 127337, 127344, 127369, 130032, 130041, 131072, 173791, 173824,
		177977, 177984, 178205, 178208, 183969, 183984, 191456, 194560, 195101,
		196608, 201546, 201552, 205743, 1, 0, 49, 57, 1, 0, 48, 57, 2, 0, 43, 43,
		45, 45, 2, 0, 69, 69, 101, 101, 10, 0, 9, 13, 32, 32, 133, 133, 160, 160,
		5760, 5760, 8192, 8202, 8232, 8233, 8239, 8239, 8287, 8287, 12288, 12288,
		1, 0, 10, 10, 1, 1, 10, 10, 202, 0, 1, 1, 0, 0, 0, 0, 3, 1, 0, 0, 0, 0,
		5, 1, 0, 0, 0, 0, 7, 1, 0, 0, 0, 0, 9, 1, 0, 0, 0, 0, 11, 1, 0, 0, 0, 0,
		13, 1, 0, 0, 0, 0, 15, 1, 0, 0, 0, 0, 17, 1, 0, 0, 0, 0, 19, 1, 0, 0, 0,
		0, 29, 1, 0, 0, 0, 0, 31, 1, 0, 0, 0, 0, 33, 1, 0, 0, 0, 0, 35, 1, 0, 0,
		0, 0, 39, 1, 0, 0, 0, 0, 41, 1, 0, 0, 0, 0, 43, 1, 0, 0, 0, 1, 49, 1, 0,
		0, 0, 3, 57, 1, 0, 0, 0, 5, 59, 1, 0, 0, 0, 7, 61, 1, 0, 0, 0, 9, 63, 1,
		0, 0, 0, 11, 65, 1, 0, 0, 0, 13, 67, 1, 0, 0, 0, 15, 78, 1, 0, 0, 0, 17,
		80, 1, 0, 0, 0, 19, 90, 1, 0, 0, 0, 21, 99, 1, 0, 0, 0, 23, 104, 1, 0,
		0, 0, 25, 110, 1, 0, 0, 0, 27, 112, 1, 0, 0, 0, 29, 114, 1, 0, 0, 0, 31,
		121, 1, 0, 0, 0, 33, 128, 1, 0, 0, 0, 35, 139, 1, 0, 0, 0, 37, 145, 1,
		0, 0, 0, 39, 155, 1, 0, 0, 0, 41, 162, 1, 0, 0, 0, 43, 170, 1, 0, 0, 0,
		45, 175, 1, 0, 0, 0, 47, 179, 1, 0, 0, 0, 49, 50, 5, 118, 0, 0, 50, 51,
		5, 101, 0, 0, 51, 52, 5, 114, 0, 0, 52, 53, 5, 115, 0, 0, 53, 54, 5, 105,
		0, 0, 54, 55, 5, 111, 0, 0, 55, 56, 5, 110, 0, 0, 56, 2, 1, 0, 0, 0, 57,
		58, 5, 118, 0, 0, 58, 4, 1, 0, 0, 0, 59, 60, 5, 123, 0, 0, 60, 6, 1, 0,
		0, 0, 61, 62, 5, 125, 0, 0, 62, 8, 1, 0, 0, 0, 63, 64, 5, 40, 0, 0, 64,
		10, 1, 0, 0, 0, 65, 66, 5, 41, 0, 0, 66, 12, 1, 0, 0, 0, 67, 68, 5, 61,
		0, 0, 68, 14, 1, 0, 0, 0, 69, 70, 5, 116, 0, 0, 70, 71, 5, 114, 0, 0, 71,
		72, 5, 117, 0, 0, 72, 79, 5, 101, 0, 0, 73, 74, 5, 102, 0, 0, 74, 75, 5,
		97, 0, 0, 75, 76, 5, 108, 0, 0, 76, 77, 5, 115, 0, 0, 77, 79, 5, 101, 0,
		0, 78, 69, 1, 0, 0, 0, 78, 73, 1, 0, 0, 0, 79, 16, 1, 0, 0, 0, 80, 85,
		5, 34, 0, 0, 81, 84, 3, 21, 10, 0, 82, 84, 3, 27, 13, 0, 83, 81, 1, 0,
		0, 0, 83, 82, 1, 0, 0, 0, 84, 87, 1, 0, 0, 0, 85, 83, 1, 0, 0, 0, 85, 86,
		1, 0, 0, 0, 86, 88, 1, 0, 0, 0, 87, 85, 1, 0, 0, 0, 88, 89, 5, 34, 0, 0,
		89, 18, 1, 0, 0, 0, 90, 94, 5, 96, 0, 0, 91, 93, 8, 0, 0, 0, 92, 91, 1,
		0, 0, 0, 93, 96, 1, 0, 0, 0, 94, 92, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95,
		97, 1, 0, 0, 0, 96, 94, 1, 0, 0, 0, 97, 98, 5, 96, 0, 0, 98, 20, 1, 0,
		0, 0, 99, 102, 5, 92, 0, 0, 100, 103, 7, 1, 0, 0, 101, 103, 3, 23, 11,
		0, 102, 100, 1, 0, 0, 0, 102, 101, 1, 0, 0, 0, 103, 22, 1, 0, 0, 0, 104,
		105, 5, 117, 0, 0, 105, 106, 3, 25, 12, 0, 106, 107, 3, 25, 12, 0, 107,
		108, 3, 25, 12, 0, 108, 109, 3, 25, 12, 0, 109, 24, 1, 0, 0, 0, 110, 111,
		7, 2, 0, 0, 111, 26, 1, 0, 0, 0, 112, 113, 8, 3, 0, 0, 113, 28, 1, 0, 0,
		0, 114, 118, 7, 4, 0, 0, 115, 117, 7, 5, 0, 0, 116, 115, 1, 0, 0, 0, 117,
		120, 1, 0, 0, 0, 118, 116, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 30, 1,
		0, 0, 0, 120, 118, 1, 0, 0, 0, 121, 125, 7, 6, 0, 0, 122, 124, 7, 7, 0,
		0, 123, 122, 1, 0, 0, 0, 124, 127, 1, 0, 0, 0, 125, 123, 1, 0, 0, 0, 125,
		126, 1, 0, 0, 0, 126, 32, 1, 0, 0, 0, 127, 125, 1, 0, 0, 0, 128, 129, 3,
		35, 17, 0, 129, 131, 5, 46, 0, 0, 130, 132, 7, 7, 0, 0, 131, 130, 1, 0,
		0, 0, 132, 133, 1, 0, 0, 0, 133, 131, 1, 0, 0, 0, 133, 134, 1, 0, 0, 0,
		134, 136, 1, 0, 0, 0, 135, 137, 3, 37, 18, 0, 136, 135, 1, 0, 0, 0, 136,
		137, 1, 0, 0, 0, 137, 34, 1, 0, 0, 0, 138, 140, 7, 8, 0, 0, 139, 138, 1,
		0, 0, 0, 139, 140, 1, 0, 0, 0, 140, 143, 1, 0, 0, 0, 141, 144, 5, 48, 0,
		0, 142, 144, 3, 31, 15, 0, 143, 141, 1, 0, 0, 0, 143, 142, 1, 0, 0, 0,
		144, 36, 1, 0, 0, 0, 145, 147, 7, 9, 0, 0, 146, 148, 7, 8, 0, 0, 147, 146,
		1, 0, 0, 0, 147, 148, 1, 0, 0, 0, 148, 150, 1, 0, 0, 0, 149, 151, 7, 7,
		0, 0, 150, 149, 1, 0, 0, 0, 151, 152, 1, 0, 0, 0, 152, 150, 1, 0, 0, 0,
		152, 153, 1, 0, 0, 0, 153, 38, 1, 0, 0, 0, 154, 156, 3, 43, 21, 0, 155,
		154, 1, 0, 0, 0, 155, 156, 1, 0, 0, 0, 156, 157, 1, 0, 0, 0, 157, 159,
		5, 59, 0, 0, 158, 160, 3, 43, 21, 0, 159, 158, 1, 0, 0, 0, 159, 160, 1,
		0, 0, 0, 160, 40, 1, 0, 0, 0, 161, 163, 3, 43, 21, 0, 162, 161, 1, 0, 0,
		0, 162, 163, 1, 0, 0, 0, 163, 164, 1, 0, 0, 0, 164, 166, 5, 44, 0, 0, 165,
		167, 3, 43, 21, 0, 166, 165, 1, 0, 0, 0, 166, 167, 1, 0, 0, 0, 167, 42,
		1, 0, 0, 0, 168, 171, 3, 45, 22, 0, 169, 171, 3, 47, 23, 0, 170, 168, 1,
		0, 0, 0, 170, 169, 1, 0, 0, 0, 171, 172, 1, 0, 0, 0, 172, 170, 1, 0, 0,
		0, 172, 173, 1, 0, 0, 0, 173, 44, 1, 0, 0, 0, 174, 176, 7, 10, 0, 0, 175,
		174, 1, 0, 0, 0, 176, 177, 1, 0, 0, 0, 177, 175, 1, 0, 0, 0, 177, 178,
		1, 0, 0, 0, 178, 46, 1, 0, 0, 0, 179, 183, 5, 35, 0, 0, 180, 182, 8, 11,
		0, 0, 181, 180, 1, 0, 0, 0, 182, 185, 1, 0, 0, 0, 183, 181, 1, 0, 0, 0,
		183, 184, 1, 0, 0, 0, 184, 187, 1, 0, 0, 0, 185, 183, 1, 0, 0, 0, 186,
		188, 7, 12, 0, 0, 187, 186, 1, 0, 0, 0, 188, 48, 1, 0, 0, 0, 23, 0, 78,
		83, 85, 94, 102, 118, 125, 133, 136, 139, 143, 147, 152, 155, 159, 162,
		166, 170, 172, 177, 183, 187, 0,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// UIDLLexerInit initializes any static state used to implement UIDLLexer. By default the
// static state used to implement the lexer is lazily initialized during the first call to
// NewUIDLLexer(). You can call this function if you wish to initialize the static state ahead
// of time.
func UIDLLexerInit() {
	staticData := &UIDLLexerLexerStaticData
	staticData.once.Do(uidllexerLexerInit)
}

// NewUIDLLexer produces a new lexer instance for the optional input antlr.CharStream.
func NewUIDLLexer(input antlr.CharStream) *UIDLLexer {
	UIDLLexerInit()
	l := new(UIDLLexer)
	l.BaseLexer = antlr.NewBaseLexer(input)
	staticData := &UIDLLexerLexerStaticData
	l.Interpreter = antlr.NewLexerATNSimulator(l, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	l.channelNames = staticData.ChannelNames
	l.modeNames = staticData.ModeNames
	l.RuleNames = staticData.RuleNames
	l.LiteralNames = staticData.LiteralNames
	l.SymbolicNames = staticData.SymbolicNames
	l.GrammarFileName = "UIDL.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// UIDLLexer tokens.
const (
	UIDLLexerT__0               = 1
	UIDLLexerT__1               = 2
	UIDLLexerT__2               = 3
	UIDLLexerT__3               = 4
	UIDLLexerT__4               = 5
	UIDLLexerT__5               = 6
	UIDLLexerT__6               = 7
	UIDLLexerBool               = 8
	UIDLLexerDoubleQuotedString = 9
	UIDLLexerBackQuotedString   = 10
	UIDLLexerIdentifier         = 11
	UIDLLexerNatural            = 12
	UIDLLexerFloat              = 13
	UIDLLexerInt                = 14
	UIDLLexerSemicolon          = 15
	UIDLLexerComma              = 16
	UIDLLexerWhiteSpace         = 17
)
