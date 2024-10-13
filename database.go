package main

var database = map[int]Content{
	1: {
		name:      "Kimura trap",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			KIMURA: true,
		},
	},
	2: {
		name:      "Omoplata",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			OMOPLATA: true,
		},
	},
	3: {
		name:      "Americana",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: ATACK,
		categories: map[string]bool{
			AMERICANA: true,
		},
	},
	4: {
		name:      "Defesa americana",
		url:       "https://www.instagram.com/p/C_YZK-1R8Ud/",
		technique: DEFENSE,
		categories: map[string]bool{
			AMERICANA: true,
		},
	},
}
