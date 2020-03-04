# Setup config file

```bash
cp config.dist.yml config.yml
```

# Run Go verbose 
```bash
go run -v ./
```

# API Call

```bash
curl --location --request POST 'HTTPS://graphql.epicgames.com/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{
  "query": "          query promotionsQuery($namespace: String!, $country: String!, $locale: String!) {            Catalog {              catalogOffers(namespace: $namespace, locale: $locale, params: {category: \"freegames\", country: $country, sortBy: \"effectiveDate\", sortDir: \"asc\"}) {                elements {                  title                  description                  id                  namespace                  categories {                    path                  }                  linkedOfferNs                  linkedOfferId                  keyImages {                    type                    url                  }                  productSlug                  promotions {                    promotionalOffers {                      promotionalOffers {                        startDate                        endDate                        discountSetting {                          discountType                          discountPercentage                        }                      }                    }                    upcomingPromotionalOffers {                      promotionalOffers {                        startDate                        endDate                        discountSetting {                          discountType                          discountPercentage                        }                      }                    }                  }                }              }            }          }        ",
  "variables": {
    "namespace": "epic",
    "country": "BE",
    "locale": "en-US"
  }
}
'
```

## Response
```json
{
    "data": {
        "Catalog": {
            "catalogOffers": {
                "elements": [
                    {
                        "title": "InnerSpace",
                        "description": "InnerSpace",
                        "id": "4e2072e53390489b8b35ddf6a1e52aa8",
                        "namespace": "epic",
                        "categories": [
                            {
                                "path": "freegames"
                            },
                            {
                                "path": "games"
                            }
                        ],
                        "linkedOfferNs": "387bc8d3398a40f7ae14de417b4acefc",
                        "linkedOfferId": "95e8b0adb366410c892f8f6fc8ac1bfa",
                        "keyImages": [
                            {
                                "type": "DieselStoreFrontWide",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_aspyr_InnerSpace_S1_StorePrimary-2560x1440-94b204bad86bd4eb0ea90675cd2d4f22.jpg"
                            },
                            {
                                "type": "DieselStoreFrontTall",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_aspyr_InnerSpace_S3_StorePortraitPromo-1280x1440-4c52c2954a5f9eddd0918eace407e1c2.jpg"
                            },
                            {
                                "type": "DieselGameBoxLogo",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_Aspyr_InnerSpace_IC1-white-827x426-ff5e261d2fe9a372f5bcabfff54f7c64.png"
                            }
                        ],
                        "productSlug": "innerspace/home",
                        "promotions": {
                            "promotionalOffers": [
                                {
                                    "promotionalOffers": [
                                        {
                                            "startDate": "2020-02-27T16:00:00.000Z",
                                            "endDate": "2020-03-05T16:00:00.000Z",
                                            "discountSetting": {
                                                "discountType": "PERCENTAGE",
                                                "discountPercentage": 0
                                            }
                                        }
                                    ]
                                }
                            ],
                            "upcomingPromotionalOffers": []
                        }
                    },
                    {
                        "title": "GoNNER",
                        "description": "GoNNER",
                        "id": "cefa7f9e970d4b93b828831d05f4e7ae",
                        "namespace": "epic",
                        "categories": [
                            {
                                "path": "freegames"
                            },
                            {
                                "path": "games"
                            }
                        ],
                        "linkedOfferNs": "109273d8719f46b9b89cb912a6649e5c",
                        "linkedOfferId": "bef777f5e5e74f6a942384638c41cc51",
                        "keyImages": [
                            {
                                "type": "DieselStoreFrontWide",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_ART_IN_HEART_GONNER_S1_DESCRIPTION-2560x1440-06254a47f26133e9c50c36b9ba5a4539.jpg"
                            },
                            {
                                "type": "DieselStoreFrontTall",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_ART_IN_HEART_GONNER_S2_DESCRIPTION-1280x1440-27d8b7fa65e5b130cbcf759a6c1204a9.jpg"
                            },
                            {
                                "type": "DieselGameBoxLogo",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_ART_IN_HEART_GONNER_IC1_DESCRIPTION-1500x500-1e41f75b20f643158b9f78a80b3ffc2c.png"
                            }
                        ],
                        "productSlug": "gonner/home",
                        "promotions": {
                            "promotionalOffers": [],
                            "upcomingPromotionalOffers": [
                                {
                                    "promotionalOffers": [
                                        {
                                            "startDate": "2020-03-05T16:00:00.000Z",
                                            "endDate": "2020-03-12T16:00:00.000Z",
                                            "discountSetting": {
                                                "discountType": "PERCENTAGE",
                                                "discountPercentage": 0
                                            }
                                        }
                                    ]
                                }
                            ]
                        }
                    },
                    {
                        "title": "Offworld Trading Company",
                        "description": "Offworld Trading Company",
                        "id": "617d128aa7ad4233918212d1e9e83eb5",
                        "namespace": "epic",
                        "categories": [
                            {
                                "path": "freegames"
                            },
                            {
                                "path": "games"
                            }
                        ],
                        "linkedOfferNs": "e8882546f28e4832af823c646aed232e",
                        "linkedOfferId": "4531b00fcf354680bf68d138b7155fda",
                        "keyImages": [
                            {
                                "type": "DieselStoreFrontTall",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_MohawkGames_OffworldTradingCompany_S4-510x680-5c715e53a4447c2dc5429fc0c76d74f6.jpg"
                            },
                            {
                                "type": "DieselStoreFrontWide",
                                "url": "https://cdn1.epicgames.com/epic/offer/EGS_MohawkGames_OffworldTradingCompany_S3-1360x766-d93f4d942f88635933d16cfc7dab4983.jpg"
                            }
                        ],
                        "productSlug": "offworld-trading-company/home",
                        "promotions": {
                            "promotionalOffers": [],
                            "upcomingPromotionalOffers": [
                                {
                                    "promotionalOffers": [
                                        {
                                            "startDate": "2020-03-05T16:00:00.000Z",
                                            "endDate": "2020-03-12T16:00:00.000Z",
                                            "discountSetting": {
                                                "discountType": "PERCENTAGE",
                                                "discountPercentage": 0
                                            }
                                        }
                                    ]
                                }
                            ]
                        }
                    }
                ]
            }
        }
    },
    "extensions": {
        "cacheControl": {
            "version": 1,
            "hints": [
                {
                    "path": [
                        "Catalog"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "categories"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "keyImages"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "promotions"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "categories"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "keyImages"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "promotions"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "categories"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "keyImages"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "promotions"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "promotions",
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "promotions",
                        "promotionalOffers",
                        0,
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "promotions",
                        "promotionalOffers",
                        0,
                        "promotionalOffers",
                        0,
                        "discountSetting"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        0,
                        "promotions",
                        "upcomingPromotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "promotions",
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "promotions",
                        "upcomingPromotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "promotions",
                        "upcomingPromotionalOffers",
                        0,
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        1,
                        "promotions",
                        "upcomingPromotionalOffers",
                        0,
                        "promotionalOffers",
                        0,
                        "discountSetting"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "promotions",
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "promotions",
                        "upcomingPromotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "promotions",
                        "upcomingPromotionalOffers",
                        0,
                        "promotionalOffers"
                    ],
                    "maxAge": 0
                },
                {
                    "path": [
                        "Catalog",
                        "catalogOffers",
                        "elements",
                        2,
                        "promotions",
                        "upcomingPromotionalOffers",
                        0,
                        "promotionalOffers",
                        0,
                        "discountSetting"
                    ],
                    "maxAge": 0
                }
            ]
        }
    }
}
```