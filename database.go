package main

var database = map[int]Content{
	1: {
		name:        "Kimura trap",
		description: "",
		url:         "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
	2: {
		name:        "Omoplata",
		description: "",
		url:         "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			OMOPLATA:   true,
		},
	},
	3: {
		name:        "Americana",
		description: "",
		url:         "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			AMERICANA:  true,
		},
	},
	4: {
		name:        "Passagem de guarda",
		description: "",
		url:         "https://www.instagram.com/p/C_YZK-1R8Ud/",
		categories: map[string]bool{
			ATACK:           true,
			PROGRESSION:     true,
			PASSAGEM_GUARDA: true,
		},
	},
	5: {
		name:        "Defesa de joelho na barriga",
		description: "",
		url:         "https://www.instagram.com/p/C919xcfu_yn/",
		categories: map[string]bool{
			DEFENSE:        true,
			JOELHO_BARRIGA: true,
		},
	},
	6: {
		name:        "Saída dos cem quilos",
		description: "João Miyao analisa essa saída dos cem kgs que Luke Stewart fez na lenda Andre Galvão em uma luta de MMA.",
		url:         "https://www.instagram.com/p/C_vhmY0JrZC/",
		categories: map[string]bool{
			DEFENSE:    true,
			CEM_QUILOS: true,
		},
	},
	7: {
		name:        "Defesa do triângulo",
		description: "Renzo Gracie mostra defesa do triângulo.",
		url:         "https://www.instagram.com/p/C1VCimQPk_P/",
		categories: map[string]bool{
			DEFENSE:   true,
			TRIANGULO: true,
		},
	},
}
