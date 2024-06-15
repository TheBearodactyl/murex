package ansi

var constants = map[string][]byte{
	// ASCII control keys
	"^@": {0},
	"^A": {1},
	"^B": {2},
	"^C": {3},
	"^D": {4},
	"^E": {5},
	"^F": {6},
	"^G": {7},
	"^H": {8},
	"^I": {9},
	"^J": {10},
	"^K": {11},
	"^L": {12},
	"^M": {13},
	"^N": {14},
	"^O": {15},
	"^P": {16},
	"^Q": {17},
	"^R": {18},
	"^S": {19},
	"^T": {20},
	"^U": {21},
	"^V": {22},
	"^W": {23},
	"^X": {24},
	"^Y": {25},
	"^Z": {26},
	"^[": {27},
	`^/`: {28},
	"^]": {29},
	"^^": {30},
	"^_": {31},

	"^?": {127},

	// ASCII control codes
	"EOF":      {4},
	"EOT":      {4},
	"BELL":     {7},
	"BS-ISO":   {8},
	"LF":       {10},
	"CR":       {13},
	"CRLF":     {13, 10},
	"ESC":      {27},
	"ESCAPE":   {27},
	"BS-ASCII": {127},

	// ASCII escape sequences
	"CURSOR-UP":         {27, 91, 65},
	"CURSOR-DOWN":       {27, 91, 66},
	"CURSOR-FORWARDS":   {27, 91, 67},
	"CURSOR-BACKWARDS":  {27, 91, 68},
	"CURSOR-HOME":       {27, 91, 72},
	"CURSOR-HOME-VT100": {27, 91, 49, 126},
	"CURSOR-END":        {27, 91, 70},
	"CURSOR-END-VT100":  {27, 91, 52, 126},

	"INS":    {27, 91, 50, 126},
	"INSERT": {27, 91, 50, 126},
	"DEL":    {27, 91, 51, 126},
	"DELETE": {27, 91, 51, 126},

	// Function keys
	"F1-VT100": {27, 79, 80},
	"F2-VT100": {27, 79, 81},
	"F3-VT100": {27, 79, 82},
	"F4-VT100": {27, 79, 83},
	"F1-VT220": {27, 91, 49, 49, 126},
	"F2-VT220": {27, 91, 49, 50, 126},
	"F3-VT220": {27, 91, 49, 51, 126},
	"F4-VT220": {27, 91, 49, 52, 126},
	"F1":       {27, 79, 80},
	"F2":       {27, 79, 81},
	"F3":       {27, 79, 82},
	"F4":       {27, 79, 83},
	"F5":       {27, 91, 49, 53, 126},
	"F6":       {27, 91, 49, 55, 126},
	"F7":       {27, 91, 49, 56, 126},
	"F8":       {27, 91, 49, 57, 126},
	"F9":       {27, 91, 50, 48, 126},
	"F10":      {27, 91, 50, 49, 126},
	"F11":      {27, 91, 50, 51, 126},
	"F12":      {27, 91, 50, 52, 126},

	// alt-num
	"ALT-0": {27, 48},
	"ALT-1": {27, 49},
	"ALT-2": {27, 50},
	"ALT-3": {27, 51},
	"ALT-4": {27, 52},
	"ALT-5": {27, 53},
	"ALT-6": {27, 54},
	"ALT-7": {27, 55},
	"ALT-8": {27, 56},
	"ALT-9": {27, 57},

	// control sequence
	"CSI": {27, 91},
}

var sgr = map[string][]byte{
	// text effects
	"RESET":      {27, 91, 48, 109},
	"BOLD":       {27, 91, 49, 109},
	"ITALIC":     {27, 91, 51, 109}, // Not commonly supported in terminals
	"UNDERSCORE": {27, 91, 52, 109},
	"UNDERLINE":  {27, 91, 52, 109},
	"UNDEROFF":   {27, 91, '2', '4', 109},
	"BLINK":      {27, 91, 53, 109},
	"INVERT":     {27, 91, 55, 109},

	"ALT-FONT-1": {27, 91, 49, 49, 109}, // Not commonly supported in terminals
	"ALT-FONT-2": {27, 91, 49, 50, 109}, // Not commonly supported in terminals
	"ALT-FONT-3": {27, 91, 49, 51, 109}, // Not commonly supported in terminals
	"ALT-FONT-4": {27, 91, 49, 52, 109}, // Not commonly supported in terminals
	"ALT-FONT-5": {27, 91, 49, 53, 109}, // Not commonly supported in terminals
	"ALT-FONT-6": {27, 91, 49, 54, 109}, // Not commonly supported in terminals
	"ALT-FONT-7": {27, 91, 49, 55, 109}, // Not commonly supported in terminals
	"ALT-FONT-8": {27, 91, 49, 56, 109}, // Not commonly supported in terminals
	"ALT-FONT-9": {27, 91, 49, 57, 109}, // Not commonly supported in terminals
	"FRAKTUR":    {27, 91, 50, 48, 109}, // Not commonly supported in terminals

	// text colours
	"BLACK":   {27, 91, 51, 48, 109},
	"RED":     {27, 91, 51, 49, 109},
	"GREEN":   {27, 91, 51, 50, 109},
	"YELLOW":  {27, 91, 51, 51, 109},
	"BLUE":    {27, 91, 51, 52, 109},
	"MAGENTA": {27, 91, 51, 53, 109},
	"CYAN":    {27, 91, 51, 54, 109},
	"WHITE":   {27, 91, 51, 55, 109},

	"BLACK-BRIGHT":   {27, 91, 49, 59, 51, 48, 109},
	"RED-BRIGHT":     {27, 91, 49, 59, 51, 48, 109},
	"GREEN-BRIGHT":   {27, 91, 49, 59, 51, 48, 109},
	"YELLOW-BRIGHT":  {27, 91, 49, 59, 51, 48, 109},
	"BLUE-BRIGHT":    {27, 91, 49, 59, 51, 48, 109},
	"MAGENTA-BRIGHT": {27, 91, 49, 59, 51, 48, 109},
	"CYAN-BRIGHT":    {27, 91, 49, 59, 51, 48, 109},
	"WHITE-BRIGHT":   {27, 91, 49, 59, 51, 48, 109},

	// background colours
	"BG-BLACK":   {27, 91, 52, 48, 109},
	"BG-RED":     {27, 91, 52, 49, 109},
	"BG-GREEN":   {27, 91, 52, 50, 109},
	"BG-YELLOW":  {27, 91, 52, 51, 109},
	"BG-BLUE":    {27, 91, 52, 52, 109},
	"BG-MAGENTA": {27, 91, 52, 53, 109},
	"BG-CYAN":    {27, 91, 52, 54, 109},
	"BG-WHITE":   {27, 91, 52, 55, 109},

	"BG-BLACK-BRIGHT":   {27, 91, 49, 59, 52, 48, 109},
	"BG-RED-BRIGHT":     {27, 91, 49, 59, 52, 48, 109},
	"BG-GREEN-BRIGHT":   {27, 91, 49, 59, 52, 48, 109},
	"BG-YELLOW-BRIGHT":  {27, 91, 49, 59, 52, 48, 109},
	"BG-BLUE-BRIGHT":    {27, 91, 49, 59, 52, 48, 109},
	"BG-MAGENTA-BRIGHT": {27, 91, 49, 59, 52, 48, 109},
	"BG-CYAN-BRIGHT":    {27, 91, 49, 59, 52, 48, 109},
	"BG-WHITE-BRIGHT":   {27, 91, 49, 59, 52, 48, 109},
}
