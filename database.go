package main

var database = map[int]Content{
	1: {
		name:        "Kimura trap",
		description: "Jordan Teaches Jiujitsu Youtube channel.",
		url:         "https://www.youtube.com/watch?v=xyCakxmx-2E&ab_channel=JordanTeachesJiujitsu",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
	2: {
		name:        "Omoplata",
		description: "Kyra Gracie e Rayron ensinam os detalhes do Omoplata.",
		url:         "https://www.youtube.com/watch?v=umwnACtt3sU",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			OMOPLATA:   true,
		},
	},
	3: {
		name:        "Americana from side control!",
		description: "Slickest way to set up the Americana from side control! Matt Arroyo channel.",
		url:         "https://www.youtube.com/watch?v=F7fcbptewRQ",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			AMERICANA:  true,
		},
	},
	4: {
		name:        "PASSE TODAS AS GUARDAS DA ACADEMIA",
		description: "Canal Living From Jiu -Jitsu no Youtube.",
		url:         "https://www.youtube.com/watch?v=zmB2PIj5cqg",
		categories: map[string]bool{
			ATACK:           true,
			PROGRESSION:     true,
			PASSAGEM_GUARDA: true,
		},
	},
	5: {
		name:        "Defesa de joelho na barriga",
		description: "Mestre @jura_bjj mostra defesa de joelho na barriga.",
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
	8: {
		name:        "72 Kimura Trap Techniques",
		description: "72 Kimura Trap Techniques In Just 22 Minutes by Jason Scully - BJJ Grappling - Kimura Trap System.",
		url:         "https://www.youtube.com/watch?v=iiF0S-B_81o",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
}
