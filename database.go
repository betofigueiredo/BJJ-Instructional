package main

var database = []Content{
	{
		name:        "Kimura trap",
		description: "Jordan Teaches Jiujitsu Youtube channel.",
		url:         "https://www.youtube.com/watch?v=xyCakxmx-2E&ab_channel=JordanTeachesJiujitsu",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
	{
		name:        "Omoplata",
		description: "Kyra Gracie e Rayron ensinam os detalhes do Omoplata.",
		url:         "https://www.youtube.com/watch?v=umwnACtt3sU",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			OMOPLATA:   true,
		},
	},
	{
		name:        "Americana from side control!",
		description: "Slickest way to set up the Americana from side control! Matt Arroyo channel.",
		url:         "https://www.youtube.com/watch?v=F7fcbptewRQ",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			AMERICANA:  true,
		},
	},
	{
		name:        "PASSE TODAS AS GUARDAS DA ACADEMIA",
		description: "Canal Living From Jiu -Jitsu no Youtube.",
		url:         "https://www.youtube.com/watch?v=zmB2PIj5cqg",
		categories: map[string]bool{
			ATACK:           true,
			PROGRESSION:     true,
			PASSAGEM_GUARDA: true,
		},
	},
	{
		name:        "Defesa de joelho na barriga",
		description: "Mestre @jura_bjj mostra defesa de joelho na barriga.",
		url:         "https://www.instagram.com/p/C919xcfu_yn/",
		categories: map[string]bool{
			DEFENSE:        true,
			JOELHO_BARRIGA: true,
		},
	},
	{
		name:        "Saída dos cem quilos",
		description: "João Miyao analisa essa saída dos cem kgs que Luke Stewart fez na lenda Andre Galvão em uma luta de MMA.",
		url:         "https://www.instagram.com/p/C_vhmY0JrZC/",
		categories: map[string]bool{
			DEFENSE:    true,
			CEM_QUILOS: true,
		},
	},
	{
		name:        "Defesa do triângulo",
		description: "Renzo Gracie mostra defesa do triângulo.",
		url:         "https://www.instagram.com/p/C1VCimQPk_P/",
		categories: map[string]bool{
			DEFENSE:   true,
			TRIANGULO: true,
		},
	},
	{
		name:        "72 Kimura Trap Techniques",
		description: "72 Kimura Trap Techniques In Just 22 Minutes by Jason Scully - BJJ Grappling - Kimura Trap System.",
		url:         "https://www.youtube.com/watch?v=iiF0S-B_81o",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			KIMURA:     true,
		},
	},
	{
		name:        "Armlock do triângulo",
		description: "Roger Gracie ensina armlock a partir do triângulo, aula 1!",
		url:         "https://www.youtube.com/watch?v=3TJDLHO2HFM",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			ARMLOCK:    true,
		},
	},
	{
		name:        "Triângulo",
		description: "Como ter um triângulo perfeito. Canal @gabrielsouzag_.",
		url:         "https://www.youtube.com/shorts/1thpGd8axQM",
		categories: map[string]bool{
			ATACK:      true,
			SUBMISSION: true,
			TRIANGULO:  true,
		},
	},
	{
		name:        "Kata Gatame",
		description: "Kata Gatame submission from mount control. Professor Philipe Della Monica.",
		url:         "https://www.youtube.com/watch?v=qOk7rCD8r_Q",
		categories: map[string]bool{
			ATACK:       true,
			SUBMISSION:  true,
			KATA_GATAME: true,
		},
	},
	{
		name:        "Estrangulamento partindo do 100 kilos",
		description: "Estrangulamento partindo do 100 kilos com braço enrolado. Fabio Gurgel.",
		url:         "https://www.youtube.com/watch?v=zZhyQve2KTU",
		categories: map[string]bool{
			ATACK:           true,
			SUBMISSION:      true,
			ESTRANGULAMENTO: true,
		},
	},
	{
		name:        "Estrangulamento dos 100kg",
		description: "Jiu Jitsu para Iniciantes - Estrangulamento dos 100kg (imobilização lateral).",
		url:         "https://www.youtube.com/watch?v=4qpBOecIH0M",
		categories: map[string]bool{
			ATACK:           true,
			SUBMISSION:      true,
			ESTRANGULAMENTO: true,
		},
	},
}
