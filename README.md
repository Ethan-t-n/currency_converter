# Currency Converter

The Currency Converter is a simple command-line program written in Go that allows users to convert amounts between different currencies using real-time exchange rates.

## Features

- Fetches the latest exchange rates from an external API.
- Displays the available currencies.
- Prompts the user to enter the amount and the target currency for conversion.
- Performs the currency conversion based on the provided inputs.
- Prints the original amount, converted amount, and target currency.

## Prerequisites

- Go programming language (version 1.13 or higher)
- Internet connection (to fetch exchange rate data from the API)

## Getting Started

1. Clone the repository or download the source code files to your local machine.

2. Open a terminal and navigate to the project directory.

3. Build the program by running the following command:

   ```terminal
   go build
4. Run the program with the following command: 
   
   ```terminal
   go run currency_converter.go   
5. Follow the prompts in the terminal to enter the amount and target currency for conversion.

6. The program will display the converted amount along with the original amount and target currency.

## Currency guide

```terminal
AOA - Angolan Kwanza
ARS - Argentine Peso
AUD - Australian Dollar
AWG - Aruban Florin
AZN - Azerbaijani Manat
BAM - Bosnia-Herzegovina Convertible Mark
BBD - Barbadian Dollar
BMD - Bermudan Dollar
BND - Brunei Dollar
BOB - Bolivian Boliviano
BRL - Brazilian Real
BSD - Bahamian Dollar
BTN - Bhutanese Ngultrum
BWP - Botswanan Pula
BYN - Belarusian Ruble
BZD - Belize Dollar
CAD - Canadian Dollar
CDF - Congolese Franc
CHF - Swiss Franc
CLP - Chilean Peso
CNY - Chinese Yuan
COP - Colombian Peso
CRC - Costa Rican Colón
CVE - Cape Verdean Escudo
CZK - Czech Koruna
DJF - Djiboutian Franc
DKK - Danish Krone
DOP - Dominican Peso
DZD - Algerian Dinar
EGP - Egyptian Pound
ERN - Eritrean Nakfa
ETB - Ethiopian Birr
EUR - Euro
FJD - Fijian Dollar
FKP - Falkland Islands Pound
FOK - Fictitious Currency
GEL - Georgian Lari
GGP - Guernsey Pound
GHS - Ghanaian Cedi
GIP - Gibraltar Pound
GMD - Gambian Dalasi
GNF - Guinean Franc
GTQ - Guatemalan Quetzal
GYD - Guyanaese Dollar
HKD - Hong Kong Dollar
HNL - Honduran Lempira
HRK - Croatian Kuna
HTG - Haitian Gourde
HUF - Hungarian Forint
IDR - Indonesian Rupiah
ILS - Israeli Shekel
IMP - Manx Pound
INR - Indian Rupee
IQD - Iraqi Dinar
IRR - Iranian Rial
ISK - Icelandic Króna
JEP - Jersey Pound
JMD - Jamaican Dollar
JOD - Jordanian Dinar
JPY - Japanese Yen
KES - Kenyan Shilling
KGS - Kyrgystani Som
KHR - Cambodian Riel
KID - Kiribati Dollar
KMF - Comorian Franc
KRW - South Korean Won
KWD - Kuwaiti Dinar
KYD - Cayman Islands Dollar
KZT - Kazakhstani Tenge
LAK - Laotian Kip
LBP - Lebanese Pound
LKR - Sri Lankan Rupee
LRD - Liberian Dollar
LSL - Lesotho Loti
LTL - Lithuanian Litas
LVL - Latvian Lats
LYD - Libyan Dinar
MAD - Moroccan Dirham
MDL - Moldovan Leu
MKD - Macedonian Denar
MMK - Myanma Kyat
MNT - Mongolian Tögrög
MOP - Macanese Pataca
MGA - Malagasy Ariary
MWK - Malawian Kwacha
MYR - Malaysian Ringgit
MZN - Mozambican Metical
NAD - Namibian Dollar
NGN - Nigerian Naira
NIO - Nicaraguan Córdoba
NOK - Norwegian Krone
NPR - Nepalese Rupee
NZD - New Zealand Dollar
OMR - Omani Rial
PAB - Panamanian Balboa
PEN - Peruvian Sol
PGK - Papua New Guinean Kina
PHP - Philippine Peso
PKR - Pakistani Rupee
PLN - Polish Złoty
PYG - Paraguayan Guarani
QAR - Qatari Riyal
RON - Romanian Leu
RSD - Serbian Dinar
RUB - Russian Ruble
RWF - Rwandan Franc
SAR - Saudi Riyal
SBD - Solomon Islands Dollar
SCR - Seychellois Rupee
SDG - Sudanese Pound
SEK - Swedish Krona
SGD - Singapore Dollar
SHP - St. Helena Pound
SLL - Sierra Leonean Leone
SOS - Somali Shilling
SRD - Surinamese Dollar
SSP - South Sudanese Pound
STN - São Tomé & Príncipe Dobra
SYP - Syrian Pound
SZL - Eswatini Lilangeni
THB - Thai Baht
TJS - Tajikistani Somoni
TOP - Tongan Paʻanga
TTD - Trinidad & Tobago Dollar
TVD - Tuvaluan Dollar
TWD - New Taiwan Dollar
TZS - Tanzanian Shilling
UAH - Ukrainian Hryvnia
UGX - Ugandan Shilling
USD - United States Dollar
UYU - Uruguayan Peso
UZS - Uzbekistani Som
VES - Venezuelan Bolívar
VND - Vietnamese Đồng
VUV - Vanuatu Vatu
WST - Samoan Tala
XAF - Central African CFA Franc
XCD - East Caribbean Dollar
XDR - Special Drawing Rights
XOF - West African CFA Franc
XPF - CFP Franc
YER - Yemeni Rial
ZAR - South African Rand
ZMW - Zambian Kwacha
ZWL - Zimbabwean Dollar
```
## Testing

The currency converter program includes a unit test to verify the correctness of the conversion logic. To run the test, follow these steps:

1. Open a terminal and navigate to the project directory.
2. Execute the following command:

   ```terminal
   go test
The test will run, and the test results will be displayed in the terminal.

## Dependencies
This program uses the following external dependencies:
- github.com/go-resty/resty/v2: A Go HTTP client library for making API requests.
- github.com/stretchr/testify/assert: A testing library for Go that provides helpful assertion functions.

This program fetches the latest exchange rates from the ExchangeRate-API service.

## Contributing 
Contributions to the Currency Converter program are welcome! If you find any issues or have suggestions for improvements, please feel free to submit a pull request or open an issue on the GitHub repository.
