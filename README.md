# Итоговая аттестация по программе профессиональной переподготовки «Go-разработчик»

## Конфигурация
используемые модули:
```
"github.com/gin-gonic/gin"
```

## Запуск
go run cmd/main.go

## Работа с сервисом

В соотвествии с заданием  сетевая служба получает GET запрос и возвращает набор данных в формате json в соотвествии с ТЗ
```curl
GET http://127.0.0.1:8585/
```

```
http://127.0.0.1:8585/

HTTP/1.1 200 OK
Date: Wed, 01 Feb 2023 17:52:52 GMT
Content-Type: text/plain; charset=utf-8
Transfer-Encoding: chunked
```
```json
{
    "status": true,
    "data": {
        "sms": [
            [
                {
                    "Country": "Monaco",
                    "Bandwidth": "57",
                    "ResponseTime": "1194",
                    "Provider": "Kildy"
                },
                {
                    "Country": "New Zealand",
                    "Bandwidth": "74",
                    "ResponseTime": "1811",
                    "Provider": "Kildy"
                },
                {
                    "Country": "Saint Barthelemy",
                    "Bandwidth": "54",
                    "ResponseTime": "1260",
                    "Provider": "Kildy"
                },
                {
                    "Country": "Canada",
                    "Bandwidth": "51",
                    "ResponseTime": "111",
                    "Provider": "Rond"
                },
                {
                    "Country": "USA",
                    "Bandwidth": "92",
                    "ResponseTime": "1370",
                    "Provider": "Rond"
                },
                {
                    "Country": "Turkey",
                    "Bandwidth": "73",
                    "ResponseTime": "56",
                    "Provider": "Rond"
                },
                {
                    "Country": "Bulgaria",
                    "Bandwidth": "56",
                    "ResponseTime": "1892",
                    "Provider": "Rond"
                },
                {
                    "Country": "France",
                    "Bandwidth": "28",
                    "ResponseTime": "1695",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Denmark",
                    "Bandwidth": "72",
                    "ResponseTime": "360",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Spain",
                    "Bandwidth": "72",
                    "ResponseTime": "1012",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Switzerland",
                    "Bandwidth": "34",
                    "ResponseTime": "1187",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Austria",
                    "Bandwidth": "56",
                    "ResponseTime": "275",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Peru",
                    "Bandwidth": "16",
                    "ResponseTime": "296",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Russia",
                    "Bandwidth": "43",
                    "ResponseTime": "1755",
                    "Provider": "Topolo"
                },
                {
                    "Country": "United Kingdom",
                    "Bandwidth": "17",
                    "ResponseTime": "1698",
                    "Provider": "Topolo"
                }
            ],
            [
                {
                    "Country": "Austria",
                    "Bandwidth": "56",
                    "ResponseTime": "275",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Bulgaria",
                    "Bandwidth": "56",
                    "ResponseTime": "1892",
                    "Provider": "Rond"
                },
                {
                    "Country": "Canada",
                    "Bandwidth": "51",
                    "ResponseTime": "111",
                    "Provider": "Rond"
                },
                {
                    "Country": "Denmark",
                    "Bandwidth": "72",
                    "ResponseTime": "360",
                    "Provider": "Topolo"
                },
                {
                    "Country": "France",
                    "Bandwidth": "28",
                    "ResponseTime": "1695",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Monaco",
                    "Bandwidth": "57",
                    "ResponseTime": "1194",
                    "Provider": "Kildy"
                },
                {
                    "Country": "New Zealand",
                    "Bandwidth": "74",
                    "ResponseTime": "1811",
                    "Provider": "Kildy"
                },
                {
                    "Country": "Peru",
                    "Bandwidth": "16",
                    "ResponseTime": "296",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Russia",
                    "Bandwidth": "43",
                    "ResponseTime": "1755",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Saint Barthelemy",
                    "Bandwidth": "54",
                    "ResponseTime": "1260",
                    "Provider": "Kildy"
                },
                {
                    "Country": "Spain",
                    "Bandwidth": "72",
                    "ResponseTime": "1012",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Switzerland",
                    "Bandwidth": "34",
                    "ResponseTime": "1187",
                    "Provider": "Topolo"
                },
                {
                    "Country": "Turkey",
                    "Bandwidth": "73",
                    "ResponseTime": "56",
                    "Provider": "Rond"
                },
                {
                    "Country": "USA",
                    "Bandwidth": "92",
                    "ResponseTime": "1370",
                    "Provider": "Rond"
                },
                {
                    "Country": "United Kingdom",
                    "Bandwidth": "17",
                    "ResponseTime": "1698",
                    "Provider": "Topolo"
                }
            ]
        ],
        "mms": [
            [
                {
                    "country": "Monaco",
                    "provider": "Kildy",
                    "bandwidth": "94",
                    "response_time": "116"
                },
                {
                    "country": "New Zealand",
                    "provider": "Kildy",
                    "bandwidth": "33",
                    "response_time": "1178"
                },
                {
                    "country": "Saint Barthelemy",
                    "provider": "Kildy",
                    "bandwidth": "33",
                    "response_time": "592"
                },
                {
                    "country": "Canada",
                    "provider": "Rond",
                    "bandwidth": "45",
                    "response_time": "1199"
                },
                {
                    "country": "USA",
                    "provider": "Rond",
                    "bandwidth": "56",
                    "response_time": "618"
                },
                {
                    "country": "Turkey",
                    "provider": "Rond",
                    "bandwidth": "79",
                    "response_time": "1693"
                },
                {
                    "country": "Bulgaria",
                    "provider": "Rond",
                    "bandwidth": "60",
                    "response_time": "1768"
                },
                {
                    "country": "France",
                    "provider": "Topolo",
                    "bandwidth": "14",
                    "response_time": "1320"
                },
                {
                    "country": "Denmark",
                    "provider": "Topolo",
                    "bandwidth": "20",
                    "response_time": "1422"
                },
                {
                    "country": "Spain",
                    "provider": "Topolo",
                    "bandwidth": "55",
                    "response_time": "1240"
                },
                {
                    "country": "Switzerland",
                    "provider": "Topolo",
                    "bandwidth": "63",
                    "response_time": "1309"
                },
                {
                    "country": "Austria",
                    "provider": "Topolo",
                    "bandwidth": "2",
                    "response_time": "680"
                },
                {
                    "country": "Peru",
                    "provider": "Topolo",
                    "bandwidth": "45",
                    "response_time": "829"
                },
                {
                    "country": "Russia",
                    "provider": "Topolo",
                    "bandwidth": "84",
                    "response_time": "1516"
                },
                {
                    "country": "United Kingdom",
                    "provider": "Topolo",
                    "bandwidth": "98",
                    "response_time": "114"
                }
            ],
            [
                {
                    "country": "Austria",
                    "provider": "Topolo",
                    "bandwidth": "2",
                    "response_time": "680"
                },
                {
                    "country": "Bulgaria",
                    "provider": "Rond",
                    "bandwidth": "60",
                    "response_time": "1768"
                },
                {
                    "country": "Canada",
                    "provider": "Rond",
                    "bandwidth": "45",
                    "response_time": "1199"
                },
                {
                    "country": "Denmark",
                    "provider": "Topolo",
                    "bandwidth": "20",
                    "response_time": "1422"
                },
                {
                    "country": "France",
                    "provider": "Topolo",
                    "bandwidth": "14",
                    "response_time": "1320"
                },
                {
                    "country": "Monaco",
                    "provider": "Kildy",
                    "bandwidth": "94",
                    "response_time": "116"
                },
                {
                    "country": "New Zealand",
                    "provider": "Kildy",
                    "bandwidth": "33",
                    "response_time": "1178"
                },
                {
                    "country": "Peru",
                    "provider": "Topolo",
                    "bandwidth": "45",
                    "response_time": "829"
                },
                {
                    "country": "Russia",
                    "provider": "Topolo",
                    "bandwidth": "84",
                    "response_time": "1516"
                },
                {
                    "country": "Saint Barthelemy",
                    "provider": "Kildy",
                    "bandwidth": "33",
                    "response_time": "592"
                },
                {
                    "country": "Spain",
                    "provider": "Topolo",
                    "bandwidth": "55",
                    "response_time": "1240"
                },
                {
                    "country": "Switzerland",
                    "provider": "Topolo",
                    "bandwidth": "63",
                    "response_time": "1309"
                },
                {
                    "country": "Turkey",
                    "provider": "Rond",
                    "bandwidth": "79",
                    "response_time": "1693"
                },
                {
                    "country": "USA",
                    "provider": "Rond",
                    "bandwidth": "56",
                    "response_time": "618"
                },
                {
                    "country": "United Kingdom",
                    "provider": "Topolo",
                    "bandwidth": "98",
                    "response_time": "114"
                }
            ]
        ],
        "voice_call": [
            {
                "Country": "RU",
                "Load": 70,
                "ResponseTime": 1983,
                "Provider": "TransparentCalls",
                "ConnectionStability": 0.67,
                "PurityTTFB": 389,
                "MedianCall": 32,
                "Unknown": 41
            },
            {
                "Country": "US",
                "Load": 0,
                "ResponseTime": 61,
                "Provider": "E-Voice",
                "ConnectionStability": 0.89,
                "PurityTTFB": 762,
                "MedianCall": 87,
                "Unknown": 58
            },
            {
                "Country": "GB",
                "Load": 24,
                "ResponseTime": 1463,
                "Provider": "TransparentCalls",
                "ConnectionStability": 0.83,
                "PurityTTFB": 701,
                "MedianCall": 15,
                "Unknown": 55
            },
            {
                "Country": "FR",
                "Load": 20,
                "ResponseTime": 1761,
                "Provider": "TransparentCalls",
                "ConnectionStability": 0.87,
                "PurityTTFB": 809,
                "MedianCall": 22,
                "Unknown": 34
            },
            {
                "Country": "BL",
                "Load": 91,
                "ResponseTime": 420,
                "Provider": "E-Voice",
                "ConnectionStability": 0.99,
                "PurityTTFB": 70,
                "MedianCall": 79,
                "Unknown": 30
            },
            {
                "Country": "AT",
                "Load": 70,
                "ResponseTime": 957,
                "Provider": "TransparentCalls",
                "ConnectionStability": 0.88,
                "PurityTTFB": 539,
                "MedianCall": 67,
                "Unknown": 24
            },
            {
                "Country": "BG",
                "Load": 0,
                "ResponseTime": 45,
                "Provider": "E-Voice",
                "ConnectionStability": 0.65,
                "PurityTTFB": 303,
                "MedianCall": 87,
                "Unknown": 19
            },
            {
                "Country": "DK",
                "Load": 97,
                "ResponseTime": 1305,
                "Provider": "JustPhone",
                "ConnectionStability": 0.72,
                "PurityTTFB": 162,
                "MedianCall": 64,
                "Unknown": 30
            },
            {
                "Country": "CA",
                "Load": 99,
                "ResponseTime": 1855,
                "Provider": "JustPhone",
                "ConnectionStability": 0.96,
                "PurityTTFB": 764,
                "MedianCall": 19,
                "Unknown": 28
            },
            {
                "Country": "ES",
                "Load": 12,
                "ResponseTime": 1332,
                "Provider": "E-Voice",
                "ConnectionStability": 0.98,
                "PurityTTFB": 477,
                "MedianCall": 22,
                "Unknown": 39
            },
            {
                "Country": "CH",
                "Load": 61,
                "ResponseTime": 1771,
                "Provider": "JustPhone",
                "ConnectionStability": 0.75,
                "PurityTTFB": 487,
                "MedianCall": 59,
                "Unknown": 38
            },
            {
                "Country": "TR",
                "Load": 4,
                "ResponseTime": 1010,
                "Provider": "TransparentCalls",
                "ConnectionStability": 0.86,
                "PurityTTFB": 442,
                "MedianCall": 72,
                "Unknown": 36
            },
            {
                "Country": "PE",
                "Load": 3,
                "ResponseTime": 1382,
                "Provider": "JustPhone",
                "ConnectionStability": 0.64,
                "PurityTTFB": 881,
                "MedianCall": 8,
                "Unknown": 40
            },
            {
                "Country": "NZ",
                "Load": 71,
                "ResponseTime": 1134,
                "Provider": "JustPhone",
                "ConnectionStability": 0.86,
                "PurityTTFB": 105,
                "MedianCall": 72,
                "Unknown": 13
            },
            {
                "Country": "MC",
                "Load": 71,
                "ResponseTime": 774,
                "Provider": "E-Voice",
                "ConnectionStability": 0.78,
                "PurityTTFB": 294,
                "MedianCall": 28,
                "Unknown": 20
            }
        ],
        "email": {
            "AT": [
                [
                    {
                        "Country": "AT",
                        "Provider": "Gmail",
                        "DeliveryTime": 11
                    },
                    {
                        "Country": "AT",
                        "Provider": "Orange",
                        "DeliveryTime": 147
                    },
                    {
                        "Country": "AT",
                        "Provider": "RediffMail",
                        "DeliveryTime": 151
                    }
                ],
                [
                    {
                        "Country": "AT",
                        "Provider": "GMX",
                        "DeliveryTime": 357
                    },
                    {
                        "Country": "AT",
                        "Provider": "Comcast",
                        "DeliveryTime": 380
                    },
                    {
                        "Country": "AT",
                        "Provider": "Hotmail",
                        "DeliveryTime": 398
                    },
                    {
                        "Country": "AT",
                        "Provider": "Yahoo",
                        "DeliveryTime": 430
                    }
                ]
            ],
            "BG": [
                [
                    {
                        "Country": "BG",
                        "Provider": "Orange",
                        "DeliveryTime": 56
                    },
                    {
                        "Country": "BG",
                        "Provider": "MSN",
                        "DeliveryTime": 100
                    },
                    {
                        "Country": "BG",
                        "Provider": "Hotmail",
                        "DeliveryTime": 136
                    }
                ],
                [
                    {
                        "Country": "BG",
                        "Provider": "RediffMail",
                        "DeliveryTime": 494
                    },
                    {
                        "Country": "BG",
                        "Provider": "Live",
                        "DeliveryTime": 500
                    },
                    {
                        "Country": "BG",
                        "Provider": "Yahoo",
                        "DeliveryTime": 505
                    },
                    {
                        "Country": "BG",
                        "Provider": "Mail.ru",
                        "DeliveryTime": 557
                    }
                ]
            ],
            "BL": [
                [
                    {
                        "Country": "BL",
                        "Provider": "Live",
                        "DeliveryTime": 57
                    },
                    {
                        "Country": "BL",
                        "Provider": "AOL",
                        "DeliveryTime": 97
                    },
                    {
                        "Country": "BL",
                        "Provider": "Gmail",
                        "DeliveryTime": 130
                    }
                ],
                [
                    {
                        "Country": "BL",
                        "Provider": "Comcast",
                        "DeliveryTime": 506
                    },
                    {
                        "Country": "BL",
                        "Provider": "Hotmail",
                        "DeliveryTime": 513
                    },
                    {
                        "Country": "BL",
                        "Provider": "Yandex",
                        "DeliveryTime": 563
                    },
                    {
                        "Country": "BL",
                        "Provider": "Orange",
                        "DeliveryTime": 573
                    }
                ]
            ],
            "CA": [
                [
                    {
                        "Country": "CA",
                        "Provider": "Yahoo",
                        "DeliveryTime": 167
                    },
                    {
                        "Country": "CA",
                        "Provider": "Live",
                        "DeliveryTime": 178
                    },
                    {
                        "Country": "CA",
                        "Provider": "RediffMail",
                        "DeliveryTime": 209
                    }
                ],
                [
                    {
                        "Country": "CA",
                        "Provider": "AOL",
                        "DeliveryTime": 453
                    },
                    {
                        "Country": "CA",
                        "Provider": "Orange",
                        "DeliveryTime": 466
                    },
                    {
                        "Country": "CA",
                        "Provider": "Hotmail",
                        "DeliveryTime": 479
                    },
                    {
                        "Country": "CA",
                        "Provider": "Yandex",
                        "DeliveryTime": 483
                    }
                ]
            ],
            "CH": [
                [
                    {
                        "Country": "CH",
                        "Provider": "Orange",
                        "DeliveryTime": 4
                    },
                    {
                        "Country": "CH",
                        "Provider": "MSN",
                        "DeliveryTime": 70
                    },
                    {
                        "Country": "CH",
                        "Provider": "GMX",
                        "DeliveryTime": 114
                    }
                ],
                [
                    {
                        "Country": "CH",
                        "Provider": "Comcast",
                        "DeliveryTime": 447
                    },
                    {
                        "Country": "CH",
                        "Provider": "Mail.ru",
                        "DeliveryTime": 452
                    },
                    {
                        "Country": "CH",
                        "Provider": "Yahoo",
                        "DeliveryTime": 541
                    },
                    {
                        "Country": "CH",
                        "Provider": "AOL",
                        "DeliveryTime": 596
                    }
                ]
            ],
            "DK": [
                [
                    {
                        "Country": "DK",
                        "Provider": "Mail.ru",
                        "DeliveryTime": 32
                    },
                    {
                        "Country": "DK",
                        "Provider": "Gmail",
                        "DeliveryTime": 82
                    },
                    {
                        "Country": "DK",
                        "Provider": "Comcast",
                        "DeliveryTime": 114
                    }
                ],
                [
                    {
                        "Country": "DK",
                        "Provider": "Live",
                        "DeliveryTime": 410
                    },
                    {
                        "Country": "DK",
                        "Provider": "Orange",
                        "DeliveryTime": 429
                    },
                    {
                        "Country": "DK",
                        "Provider": "Yandex",
                        "DeliveryTime": 436
                    },
                    {
                        "Country": "DK",
                        "Provider": "MSN",
                        "DeliveryTime": 537
                    }
                ]
            ],
            "ES": [
                [
                    {
                        "Country": "ES",
                        "Provider": "RediffMail",
                        "DeliveryTime": 91
                    },
                    {
                        "Country": "ES",
                        "Provider": "Orange",
                        "DeliveryTime": 109
                    },
                    {
                        "Country": "ES",
                        "Provider": "GMX",
                        "DeliveryTime": 166
                    }
                ],
                [
                    {
                        "Country": "ES",
                        "Provider": "Yandex",
                        "DeliveryTime": 325
                    },
                    {
                        "Country": "ES",
                        "Provider": "Hotmail",
                        "DeliveryTime": 346
                    },
                    {
                        "Country": "ES",
                        "Provider": "Yahoo",
                        "DeliveryTime": 395
                    },
                    {
                        "Country": "ES",
                        "Provider": "AOL",
                        "DeliveryTime": 590
                    }
                ]
            ],
            "FR": [
                [
                    {
                        "Country": "FR",
                        "Provider": "Live",
                        "DeliveryTime": 7
                    },
                    {
                        "Country": "FR",
                        "Provider": "Orange",
                        "DeliveryTime": 74
                    },
                    {
                        "Country": "FR",
                        "Provider": "Comcast",
                        "DeliveryTime": 147
                    }
                ],
                [
                    {
                        "Country": "FR",
                        "Provider": "Gmail",
                        "DeliveryTime": 494
                    },
                    {
                        "Country": "FR",
                        "Provider": "GMX",
                        "DeliveryTime": 505
                    },
                    {
                        "Country": "FR",
                        "Provider": "RediffMail",
                        "DeliveryTime": 540
                    },
                    {
                        "Country": "FR",
                        "Provider": "Yahoo",
                        "DeliveryTime": 599
                    }
                ]
            ],
            "GB": [
                [
                    {
                        "Country": "GB",
                        "Provider": "Comcast",
                        "DeliveryTime": 15
                    },
                    {
                        "Country": "GB",
                        "Provider": "MSN",
                        "DeliveryTime": 16
                    },
                    {
                        "Country": "GB",
                        "Provider": "AOL",
                        "DeliveryTime": 19
                    }
                ],
                [
                    {
                        "Country": "GB",
                        "Provider": "Gmail",
                        "DeliveryTime": 309
                    },
                    {
                        "Country": "GB",
                        "Provider": "Orange",
                        "DeliveryTime": 336
                    },
                    {
                        "Country": "GB",
                        "Provider": "Yandex",
                        "DeliveryTime": 424
                    },
                    {
                        "Country": "GB",
                        "Provider": "Yahoo",
                        "DeliveryTime": 572
                    }
                ]
            ],
            "MC": [
                [
                    {
                        "Country": "MC",
                        "Provider": "Hotmail",
                        "DeliveryTime": 37
                    },
                    {
                        "Country": "MC",
                        "Provider": "Live",
                        "DeliveryTime": 46
                    },
                    {
                        "Country": "MC",
                        "Provider": "AOL",
                        "DeliveryTime": 72
                    }
                ],
                [
                    {
                        "Country": "MC",
                        "Provider": "MSN",
                        "DeliveryTime": 446
                    },
                    {
                        "Country": "MC",
                        "Provider": "Yandex",
                        "DeliveryTime": 480
                    },
                    {
                        "Country": "MC",
                        "Provider": "Gmail",
                        "DeliveryTime": 494
                    },
                    {
                        "Country": "MC",
                        "Provider": "RediffMail",
                        "DeliveryTime": 574
                    }
                ]
            ],
            "NZ": [
                [
                    {
                        "Country": "NZ",
                        "Provider": "Orange",
                        "DeliveryTime": 11
                    },
                    {
                        "Country": "NZ",
                        "Provider": "Gmail",
                        "DeliveryTime": 53
                    },
                    {
                        "Country": "NZ",
                        "Provider": "Yandex",
                        "DeliveryTime": 172
                    }
                ],
                [
                    {
                        "Country": "NZ",
                        "Provider": "AOL",
                        "DeliveryTime": 428
                    },
                    {
                        "Country": "NZ",
                        "Provider": "Live",
                        "DeliveryTime": 444
                    },
                    {
                        "Country": "NZ",
                        "Provider": "Yahoo",
                        "DeliveryTime": 465
                    },
                    {
                        "Country": "NZ",
                        "Provider": "Comcast",
                        "DeliveryTime": 530
                    }
                ]
            ],
            "PE": [
                [
                    {
                        "Country": "PE",
                        "Provider": "GMX",
                        "DeliveryTime": 11
                    },
                    {
                        "Country": "PE",
                        "Provider": "AOL",
                        "DeliveryTime": 56
                    },
                    {
                        "Country": "PE",
                        "Provider": "MSN",
                        "DeliveryTime": 80
                    }
                ],
                [
                    {
                        "Country": "PE",
                        "Provider": "Comcast",
                        "DeliveryTime": 401
                    },
                    {
                        "Country": "PE",
                        "Provider": "Mail.ru",
                        "DeliveryTime": 414
                    },
                    {
                        "Country": "PE",
                        "Provider": "Hotmail",
                        "DeliveryTime": 436
                    },
                    {
                        "Country": "PE",
                        "Provider": "Gmail",
                        "DeliveryTime": 594
                    }
                ]
            ],
            "RU": [
                [
                    {
                        "Country": "RU",
                        "Provider": "Yandex",
                        "DeliveryTime": 136
                    },
                    {
                        "Country": "RU",
                        "Provider": "Gmail",
                        "DeliveryTime": 145
                    },
                    {
                        "Country": "RU",
                        "Provider": "Hotmail",
                        "DeliveryTime": 182
                    }
                ],
                [
                    {
                        "Country": "RU",
                        "Provider": "GMX",
                        "DeliveryTime": 363
                    },
                    {
                        "Country": "RU",
                        "Provider": "AOL",
                        "DeliveryTime": 436
                    },
                    {
                        "Country": "RU",
                        "Provider": "Orange",
                        "DeliveryTime": 485
                    },
                    {
                        "Country": "RU",
                        "Provider": "Yahoo",
                        "DeliveryTime": 508
                    }
                ]
            ],
            "TR": [
                [
                    {
                        "Country": "TR",
                        "Provider": "Gmail",
                        "DeliveryTime": 37
                    },
                    {
                        "Country": "TR",
                        "Provider": "GMX",
                        "DeliveryTime": 92
                    },
                    {
                        "Country": "TR",
                        "Provider": "RediffMail",
                        "DeliveryTime": 126
                    }
                ],
                [
                    {
                        "Country": "TR",
                        "Provider": "Comcast",
                        "DeliveryTime": 549
                    },
                    {
                        "Country": "TR",
                        "Provider": "AOL",
                        "DeliveryTime": 553
                    },
                    {
                        "Country": "TR",
                        "Provider": "Yahoo",
                        "DeliveryTime": 575
                    },
                    {
                        "Country": "TR",
                        "Provider": "Yandex",
                        "DeliveryTime": 588
                    }
                ]
            ],
            "US": [
                [
                    {
                        "Country": "US",
                        "Provider": "Yahoo",
                        "DeliveryTime": 14
                    },
                    {
                        "Country": "US",
                        "Provider": "MSN",
                        "DeliveryTime": 157
                    },
                    {
                        "Country": "US",
                        "Provider": "GMX",
                        "DeliveryTime": 200
                    }
                ],
                [
                    {
                        "Country": "US",
                        "Provider": "Comcast",
                        "DeliveryTime": 331
                    },
                    {
                        "Country": "US",
                        "Provider": "Hotmail",
                        "DeliveryTime": 478
                    },
                    {
                        "Country": "US",
                        "Provider": "Live",
                        "DeliveryTime": 489
                    },
                    {
                        "Country": "US",
                        "Provider": "Orange",
                        "DeliveryTime": 555
                    }
                ]
            ]
        },
        "billing": {
            "CreateCustomer": true,
            "Purchase": true,
            "Payout": true,
            "Recurring": true,
            "FraudControl": true,
            "CheckoutPage": true
        },
        "support": [
            3,
            111
        ],
        "incident": [
            {
                "topic": "SMS delivery in EU",
                "status": "closed"
            },
            {
                "topic": "MMS connection stability",
                "status": "closed"
            },
            {
                "topic": "Voice call connection purity",
                "status": "closed"
            },
            {
                "topic": "Checkout page is down",
                "status": "closed"
            },
            {
                "topic": "Support overload",
                "status": "closed"
            },
            {
                "topic": "Buy phone number not working in US",
                "status": "closed"
            },
            {
                "topic": "API Slow latency",
                "status": "closed"
            }
        ]
    },
    "error": ""
}
```
```
Response code: 200 (OK); Time: 47ms (47 ms); Content length: 21220 bytes (21,22 kB)
```
