package timezone

// CountryTimezone Country Mapping Timezone
func CountryTimezone(country string) string {

	return countriesTz[country]

}

// ISO Country To Timezone
var countriesTz = map[string]string{
	"CC": "Indian/Cocos",
	"FR": "Europe/Paris",
	"WS": "Pacific/Apia",
	"TV": "Pacific/Funafuti",
	"FI": "Europe/Helsinki",
	"GN": "Africa/Conakry",
	"NR": "Pacific/Nauru",
	"PL": "Europe/Warsaw",
	"SB": "Pacific/Guadalcanal",
	"AQ": "Antarctica/Casey",
	"BF": "Africa/Ouagadougou",
	"FO": "Atlantic/Faroe",
	"LS": "Africa/Maseru",
	"FK": "Atlantic/Stanley",
	"MH": "Pacific/Kwajalein",
	"NP": "Asia/Kathmandu",
	"AG": "America/Antigua",
	"HU": "Europe/Budapest",
	"KE": "Africa/Nairobi",
	"GG": "Europe/Guernsey",
	"NE": "Africa/Niamey",
	"TK": "Pacific/Fakaofo",
	"BQ": "America/Kralendijk",
	"TF": "Indian/Kerguelen",
	"GU": "Pacific/Guam",
	"LU": "Europe/Luxembourg",
	"PR": "America/Puerto_Rico",
	"EG": "Africa/Cairo",
	"HT": "America/Port-au-Prince",
	"IS": "Atlantic/Reykjavik",
	"PW": "Pacific/Palau",
	"PA": "America/Panama",
	"VE": "America/Caracas",
	"AF": "Asia/Kabul",
	"CM": "Africa/Douala",
	"CD": "Africa/Kinshasa",
	"IE": "Europe/Dublin",
	"JM": "America/Jamaica",
	"BM": "Atlantic/Bermuda",
	"DO": "America/Santo_Domingo",
	"DJ": "Africa/Djibouti",
	"ER": "Africa/Asmara",
	"PS": "Asia/Gaza",
	"CV": "Atlantic/Cape_Verde",
	"EE": "Europe/Tallinn",
	"MR": "Africa/Nouakchott",
	"PK": "Asia/Karachi",
	"RS": "Europe/Belgrade",
	"SS": "Africa/Juba",
	"CR": "America/Costa_Rica",
	"LC": "America/St_Lucia",
	"ZA": "Africa/Johannesburg",
	"AR": "America/Argentina/Buenos_Aires",
	"TD": "Africa/Ndjamena",
	"GE": "Asia/Tbilisi",
	"DE": "Europe/Berlin",
	"MY": "Asia/Kuala_Lumpur",
	"YT": "Indian/Mayotte",
	"NG": "Africa/Lagos",
	"GB": "Europe/London",
	"VU": "Pacific/Efate",
	"BN": "Asia/Brunei",
	"GQ": "Africa/Malabo",
	"ET": "Africa/Addis_Ababa",
	"IL": "Asia/Jerusalem",
	"SA": "Asia/Riyadh",
	"WF": "Pacific/Wallis",
	"FJ": "Pacific/Fiji",
	"RW": "Africa/Kigali",
	"SI": "Europe/Ljubljana",
	"TG": "Africa/Lome",
	"BB": "America/Barbados",
	"BW": "Africa/Gaborone",
	"CY": "Asia/Famagusta",
	"CZ": "Europe/Prague",
	"UZ": "Asia/Samarkand",
	"AL": "Europe/Tirane",
	"AS": "Pacific/Pago_Pago",
	"KY": "America/Cayman",
	"PF": "Pacific/Gambier",
	"FM": "Pacific/Chuuk",
	"MD": "Europe/Chisinau",
	"BJ": "Africa/Porto-Novo",
	"CW": "America/Curacao",
	"VN": "Asia/Ho_Chi_Minh",
	"CA": "America/Atikokan",
	"KR": "Asia/Seoul",
	"ML": "Africa/Bamako",
	"CK": "Pacific/Rarotonga",
	"NO": "Europe/Oslo",
	"PG": "Pacific/Bougainville",
	"TJ": "Asia/Dushanbe",
	"KI": "Pacific/Enderbury",
	"MK": "Europe/Skopje",
	"CO": "America/Bogota",
	"VA": "Europe/Vatican",
	"MT": "Europe/Malta",
	"ME": "Europe/Podgorica",
	"PE": "America/Lima",
	"PT": "Atlantic/Azores",
	"CL": "America/Punta_Arenas",
	"HN": "America/Tegucigalpa",
	"KP": "Asia/Pyongyang",
	"RO": "Europe/Bucharest",
	"TR": "Europe/Istanbul",
	"AZ": "Asia/Baku",
	"BY": "Europe/Minsk",
	"KG": "Asia/Bishkek",
	"RU": "Asia/Anadyr",
	"TT": "America/Port_of_Spain",
	"AE": "Asia/Dubai",
	"SG": "Asia/Singapore",
	"KW": "Asia/Kuwait",
	"AU": "Antarctica/Macquarie",
	"BS": "America/Nassau",
	"BT": "Asia/Thimphu",
	"VI": "America/St_Thomas",
	"CU": "America/Havana",
	"PY": "America/Asuncion",
	"SM": "Europe/San_Marino",
	"UM": "Pacific/Midway",
	"JE": "Europe/Jersey",
	"NF": "Pacific/Norfolk",
	"BH": "Asia/Bahrain",
	"MG": "Indian/Antananarivo",
	"UA": "Europe/Kiev",
	"UY": "America/Montevideo",
	"AD": "Europe/Andorra",
	"AI": "America/Anguilla",
	"CF": "Africa/Bangui",
	"ZM": "Africa/Lusaka",
	"BA": "Europe/Sarajevo",
	"IM": "Europe/Isle_of_Man",
	"LB": "Asia/Beirut",
	"BE": "Europe/Brussels",
	"BZ": "America/Belize",
	"AO": "Africa/Luanda",
	"CI": "Africa/Abidjan",
	"GA": "Africa/Libreville",
	"LY": "Africa/Tripoli",
	"MP": "Pacific/Saipan",
	"KH": "Asia/Phnom_Penh",
	"KM": "Indian/Comoro",
	"MS": "America/Montserrat",
	"SN": "Africa/Dakar",
	"AW": "America/Aruba",
	"MQ": "America/Martinique",
	"SH": "Atlantic/St_Helena",
	"KN": "America/St_Kitts",
	"MN": "Asia/Choibalsan",
	"LK": "Asia/Colombo",
	"SR": "America/Paramaribo",
	"BI": "Africa/Bujumbura",
	"IN": "Asia/Kolkata",
	"MW": "Africa/Blantyre",
	"CG": "Africa/Brazzaville",
	"GT": "America/Guatemala",
	"SD": "Africa/Khartoum",
	"YE": "Asia/Aden",
	"AM": "Asia/Yerevan",
	"GH": "Africa/Accra",
	"LA": "Asia/Vientiane",
	"SJ": "Arctic/Longyearbyen",
	"CH": "Europe/Zurich",
	"GW": "Africa/Bissau",
	"JP": "Asia/Tokyo",
	"JO": "Asia/Amman",
	"NC": "Pacific/Noumea",
	"GM": "Africa/Banjul",
	"BL": "America/St_Barthelemy",
	"GS": "Atlantic/South_Georgia",
	"TM": "Asia/Ashgabat",
	"ES": "Africa/Ceuta",
	"HR": "Europe/Zagreb",
	"DK": "Europe/Copenhagen",
	"IR": "Asia/Tehran",
	"LR": "Africa/Monrovia",
	"MC": "Europe/Monaco",
	"NA": "Africa/Windhoek",
	"SK": "Europe/Bratislava",
	"ZW": "Africa/Harare",
	"SV": "America/El_Salvador",
	"OM": "Asia/Muscat",
	"TL": "Asia/Dili",
	"MO": "Asia/Macau",
	"BO": "America/La_Paz",
	"KZ": "Asia/Almaty",
	"NI": "America/Managua",
	"PH": "Asia/Manila",
	"SX": "America/Lower_Princes",
	"GF": "America/Cayenne",
	"SL": "Africa/Freetown",
	"TH": "Asia/Bangkok",
	"CX": "Indian/Christmas",
	"EC": "America/Guayaquil",
	"GP": "America/Guadeloupe",
	"MA": "Africa/Casablanca",
	"UG": "Africa/Kampala",
	"VG": "America/Tortola",
	"EH": "Africa/El_Aaiun",
	"DZ": "Africa/Algiers",
	"QA": "Asia/Qatar",
	"TN": "Africa/Tunis",
	"IO": "Indian/Chagos",
	"BG": "Europe/Sofia",
	"GD": "America/Grenada",
	"RE": "Indian/Reunion",
	"VC": "America/St_Vincent",
	"DM": "America/Dominica",
	"ST": "Africa/Sao_Tome",
	"SE": "Europe/Stockholm",
	"TW": "Asia/Taipei",
	"GR": "Europe/Athens",
	"ID": "Asia/Jakarta",
	"LV": "Europe/Riga",
	"MX": "America/Bahia_Banderas",
	"BR": "America/Araguaina",
	"SC": "Indian/Mahe",
	"TZ": "Africa/Dar_es_Salaam",
	"HK": "Asia/Hong_Kong",
	"NZ": "Pacific/Auckland",
	"MF": "America/Marigot",
	"PM": "America/Miquelon",
	"SY": "Asia/Damascus",
	"GY": "America/Guyana",
	"MZ": "Africa/Maputo",
	"NL": "Europe/Amsterdam",
	"AX": "Europe/Mariehamn",
	"CN": "Asia/Shanghai",
	"MV": "Indian/Maldives",
	"NU": "Pacific/Niue",
	"US": "America/Adak",
	"AT": "Europe/Vienna",
	"MU": "Indian/Mauritius",
	"IT": "Europe/Rome",
	"LI": "Europe/Vaduz",
	"MM": "Asia/Yangon",
	"GL": "America/Danmarkshavn",
	"IQ": "Asia/Baghdad",
	"PN": "Pacific/Pitcairn",
	"SO": "Africa/Mogadishu",
	"SZ": "Africa/Mbabane",
	"TC": "America/Grand_Turk",
	"BD": "Asia/Dhaka",
	"GI": "Europe/Gibraltar",
	"LT": "Europe/Vilnius",
	"TO": "Pacific/Tongatapu",
}