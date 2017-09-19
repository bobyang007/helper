package boundaryh

var ucdSentenceTests = []ucdSentenceTest{
	{[]rune{0x1, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x1, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x1, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0xd, 0x1}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x1}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0xd}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0xd}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0xd, 0x308, 0xa}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x85}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x85}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x9}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x9}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x61}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x61}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x41}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x41}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x1bb}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x1bb}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x30}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x30}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x2e}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x2e}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x21}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x21}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x22}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x22}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x2c}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x2c}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0xad}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0xad}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xd, 0x300}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xd, 0x308, 0x300}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x1}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x1}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0xd}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0xd}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0xa}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0xa}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x85}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x85}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x9}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x9}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x61}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x61}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x41}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x41}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x1bb}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x1bb}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x30}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x30}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x2e}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x2e}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x21}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x21}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x22}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x22}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x2c}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x2c}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0xad}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0xad}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0xa, 0x300}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0xa, 0x308, 0x300}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x1}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x1}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0xd}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0xd}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0xa}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0xa}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x85}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x85}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x9}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x9}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x61}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x61}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x41}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x41}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x1bb}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x1bb}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x30}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x30}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x2e}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x2e}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x21}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x21}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x22}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x22}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x2c}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x2c}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0xad}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0xad}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x85, 0x300}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x85, 0x308, 0x300}, []Boundary{{0, 1}, {1, 3}}},
	{[]rune{0x9, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x9, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x9, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x61, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x61, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x41, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x41, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x1bb, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x1bb, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x30, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x30, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x1}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x2e, 0x308, 0x1}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x2e, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x41}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x2e, 0x308, 0x41}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x2e, 0x1bb}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x2e, 0x308, 0x1bb}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x2e, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x2e, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x2e, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x1}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x21, 0x308, 0x1}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x61}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x21, 0x308, 0x61}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0x41}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x21, 0x308, 0x41}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0x1bb}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x21, 0x308, 0x1bb}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0x30}, []Boundary{{0, 1}, {1, 2}}},
	{[]rune{0x21, 0x308, 0x30}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x21, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x21, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x22, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x22, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x2c, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x2c, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0xad, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0xad, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x1}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x1}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0xd}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0xd}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0xa}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0xa}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x85}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x85}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x9}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x9}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x61}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x61}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x41}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x41}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x1bb}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x1bb}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x30}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x30}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x2e}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x2e}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x21}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x21}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x22}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x22}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x2c}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x2c}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0xad}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0xad}, []Boundary{{0, 3}}},
	{[]rune{0x300, 0x300}, []Boundary{{0, 2}}},
	{[]rune{0x300, 0x308, 0x300}, []Boundary{{0, 3}}},
	{[]rune{0xd, 0xa, 0x61, 0xa, 0x308}, []Boundary{{0, 2}, {2, 4}, {4, 5}}},
	{[]rune{0x61, 0x308}, []Boundary{{0, 2}}},
	{[]rune{0x20, 0x200d, 0x646}, []Boundary{{0, 3}}},
	{[]rune{0x646, 0x200d, 0x20}, []Boundary{{0, 3}}},
	{[]rune{0x28, 0x22, 0x47, 0x6f, 0x2e, 0x22, 0x29, 0x20, 0x28, 0x48, 0x65, 0x20, 0x64, 0x69, 0x64, 0x2e, 0x29}, []Boundary{{0, 8}, {8, 17}}},
	{[]rune{0x28, 0x201c, 0x47, 0x6f, 0x3f, 0x201d, 0x29, 0x20, 0x28, 0x48, 0x65, 0x20, 0x64, 0x69, 0x64, 0x2e, 0x29}, []Boundary{{0, 8}, {8, 17}}},
	{[]rune{0x55, 0x2e, 0x53, 0x2e, 0x41, 0x300, 0x2e, 0x20, 0x69, 0x73}, []Boundary{{0, 10}}},
	{[]rune{0x55, 0x2e, 0x53, 0x2e, 0x41, 0x300, 0x3f, 0x20, 0x48, 0x65}, []Boundary{{0, 8}, {8, 10}}},
	{[]rune{0x55, 0x2e, 0x53, 0x2e, 0x41, 0x300, 0x2e}, []Boundary{{0, 7}}},
	{[]rune{0x33, 0x2e, 0x34}, []Boundary{{0, 3}}},
	{[]rune{0x63, 0x2e, 0x64}, []Boundary{{0, 3}}},
	{[]rune{0x43, 0x2e, 0x64}, []Boundary{{0, 3}}},
	{[]rune{0x63, 0x2e, 0x44}, []Boundary{{0, 3}}},
	{[]rune{0x43, 0x2e, 0x44}, []Boundary{{0, 3}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x74, 0x68, 0x65}, []Boundary{{0, 10}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x54, 0x68, 0x65}, []Boundary{{0, 7}, {7, 10}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x2018, 0x28, 0x74, 0x68, 0x65}, []Boundary{{0, 12}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x2018, 0x28, 0x54, 0x68, 0x65}, []Boundary{{0, 7}, {7, 12}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x308, 0x74, 0x68, 0x65}, []Boundary{{0, 11}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0xa0, 0x308, 0x54, 0x68, 0x65}, []Boundary{{0, 8}, {8, 11}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0x2019, 0x308, 0x54, 0x68, 0x65}, []Boundary{{0, 7}, {7, 10}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x29, 0xa, 0x308, 0x54, 0x68, 0x65}, []Boundary{{0, 6}, {6, 10}}},
	{[]rune{0x74, 0x68, 0x65, 0x20, 0x72, 0x65, 0x73, 0x70, 0x2e, 0x20, 0x6c, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x20, 0x61, 0x72, 0x65}, []Boundary{{0, 21}}},
	{[]rune{0x5b57, 0x2e, 0x5b57}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x5b83}, []Boundary{{0, 4}, {4, 5}}},
	{[]rune{0x65, 0x74, 0x63, 0x2e, 0x3002}, []Boundary{{0, 5}}},
	{[]rune{0x5b57, 0x3002, 0x5b83}, []Boundary{{0, 2}, {2, 3}}},
	{[]rune{0x21, 0x20, 0x20}, []Boundary{{0, 3}}},
	{[]rune{0x2060, 0x28, 0x2060, 0x22, 0x2060, 0x47, 0x2060, 0x6f, 0x2060, 0x2e, 0x2060, 0x22, 0x2060, 0x29, 0x2060, 0x20, 0x2060, 0x28, 0x2060, 0x48, 0x2060, 0x65, 0x2060, 0x20, 0x2060, 0x64, 0x2060, 0x69, 0x2060, 0x64, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2060}, []Boundary{{0, 17}, {17, 36}}},
	{[]rune{0x2060, 0x28, 0x2060, 0x201c, 0x2060, 0x47, 0x2060, 0x6f, 0x2060, 0x3f, 0x2060, 0x201d, 0x2060, 0x29, 0x2060, 0x20, 0x2060, 0x28, 0x2060, 0x48, 0x2060, 0x65, 0x2060, 0x20, 0x2060, 0x64, 0x2060, 0x69, 0x2060, 0x64, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2060}, []Boundary{{0, 17}, {17, 36}}},
	{[]rune{0x2060, 0x55, 0x2060, 0x2e, 0x2060, 0x53, 0x2060, 0x2e, 0x2060, 0x41, 0x2060, 0x300, 0x2e, 0x2060, 0x20, 0x2060, 0x69, 0x2060, 0x73, 0x2060, 0x2060}, []Boundary{{0, 21}}},
	{[]rune{0x2060, 0x55, 0x2060, 0x2e, 0x2060, 0x53, 0x2060, 0x2e, 0x2060, 0x41, 0x2060, 0x300, 0x3f, 0x2060, 0x20, 0x2060, 0x48, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 16}, {16, 21}}},
	{[]rune{0x2060, 0x55, 0x2060, 0x2e, 0x2060, 0x53, 0x2060, 0x2e, 0x2060, 0x41, 0x2060, 0x300, 0x2e, 0x2060, 0x2060}, []Boundary{{0, 15}}},
	{[]rune{0x2060, 0x33, 0x2060, 0x2e, 0x2060, 0x34, 0x2060, 0x2060}, []Boundary{{0, 8}}},
	{[]rune{0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x64, 0x2060, 0x2060}, []Boundary{{0, 8}}},
	{[]rune{0x2060, 0x43, 0x2060, 0x2e, 0x2060, 0x64, 0x2060, 0x2060}, []Boundary{{0, 8}}},
	{[]rune{0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x44, 0x2060, 0x2060}, []Boundary{{0, 8}}},
	{[]rune{0x2060, 0x43, 0x2060, 0x2e, 0x2060, 0x44, 0x2060, 0x2060}, []Boundary{{0, 8}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x74, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 22}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x54, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 15}, {15, 22}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x2018, 0x2060, 0x28, 0x2060, 0x74, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 26}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x2018, 0x2060, 0x28, 0x2060, 0x54, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 15}, {15, 26}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x308, 0x74, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 23}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0xa0, 0x2060, 0x308, 0x54, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 16}, {16, 23}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0x2019, 0x2060, 0x308, 0x54, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 14}, {14, 21}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x29, 0x2060, 0xa, 0x2060, 0x308, 0x2060, 0x54, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 12}, {12, 22}}},
	{[]rune{0x2060, 0x74, 0x2060, 0x68, 0x2060, 0x65, 0x2060, 0x20, 0x2060, 0x72, 0x2060, 0x65, 0x2060, 0x73, 0x2060, 0x70, 0x2060, 0x2e, 0x2060, 0x20, 0x2060, 0x6c, 0x2060, 0x65, 0x2060, 0x61, 0x2060, 0x64, 0x2060, 0x65, 0x2060, 0x72, 0x2060, 0x73, 0x2060, 0x20, 0x2060, 0x61, 0x2060, 0x72, 0x2060, 0x65, 0x2060, 0x2060}, []Boundary{{0, 44}}},
	{[]rune{0x2060, 0x5b57, 0x2060, 0x2e, 0x2060, 0x5b57, 0x2060, 0x2060}, []Boundary{{0, 5}, {5, 8}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x5b83, 0x2060, 0x2060}, []Boundary{{0, 9}, {9, 12}}},
	{[]rune{0x2060, 0x65, 0x2060, 0x74, 0x2060, 0x63, 0x2060, 0x2e, 0x2060, 0x3002, 0x2060, 0x2060}, []Boundary{{0, 12}}},
	{[]rune{0x2060, 0x5b57, 0x2060, 0x3002, 0x2060, 0x5b83, 0x2060, 0x2060}, []Boundary{{0, 5}, {5, 8}}},
	{[]rune{0x2060, 0x21, 0x2060, 0x20, 0x2060, 0x20, 0x2060, 0x2060}, []Boundary{{0, 8}}},
}