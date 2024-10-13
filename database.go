package main

var database = map[int]Content{
	1: {
		name: "Kimura trap",
		url:  "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
	2: {
		name: "Omoplata",
		url:  "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			OMOPLATA:   true,
		},
	},
	3: {
		name: "Americana",
		url:  "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			AMERICANA:  true,
		},
	},
	4: {
		name: "Passagem de guarda",
		url:  "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:           true,
			PROGRESSION:     true,
			PASSAGEM_GUARDA: true,
		},
	},
	5: {
		name: "Defesa americana",
		url:  "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			DEFENSE:   true,
			AMERICANA: true,
		},
	},
}
