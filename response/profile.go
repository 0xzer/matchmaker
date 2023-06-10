package response

import "time"

type ProfileResponse struct {
	Meta struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
	Data struct {
		Offerings struct {
			Plus struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount              int      `json:"amount,omitempty"`
					OfferType           string   `json:"offer_type,omitempty"`
					RefreshInterval     int      `json:"refresh_interval,omitempty"`
					RefreshIntervalUnit string   `json:"refresh_interval_unit,omitempty"`
					Tags                []string `json:"tags,omitempty"`
					IconURL             string   `json:"icon_url,omitempty"`
					PaymentMethods      []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
					Data struct {
						HideAds struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_ads,omitempty"`
						HideAge struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_age,omitempty"`
						HideDistance struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_distance,omitempty"`
						Like struct {
							Type string `json:"type,omitempty"`
						} `json:"like,omitempty"`
						ControlWhoSeesYou struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_sees_you,omitempty"`
						Passport struct {
							Type string `json:"type,omitempty"`
						} `json:"passport,omitempty"`
						Rewind struct {
							Type string `json:"type,omitempty"`
						} `json:"rewind,omitempty"`
						SuperboostAlcAccess struct {
							Type string `json:"type,omitempty"`
						} `json:"superboost_alc_access,omitempty"`
						ControlWhoYouSee struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_you_see,omitempty"`
					} `json:"data,omitempty"`
					Ordering struct {
						Carousel   []string `json:"carousel,omitempty"`
						PlusScreen []string `json:"plus_screen,omitempty"`
					} `json:"ordering,omitempty"`
					SubPageData struct {
						Cta      string `json:"cta,omitempty"`
						Termed   bool   `json:"termed,omitempty"`
						Sections []struct {
							Type     string `json:"type,omitempty"`
							Heading  string `json:"heading,omitempty"`
							Benefits []struct {
								Key         string `json:"key,omitempty"`
								Heading     string `json:"heading,omitempty"`
								Included    bool   `json:"included,omitempty"`
								Deeplink    string `json:"deeplink,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"benefits,omitempty"`
						} `json:"sections,omitempty"`
					} `json:"sub_page_data,omitempty"`
				} `json:"merchandising,omitempty"`
			} `json:"plus,omitempty"`
			Gold struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount              int      `json:"amount,omitempty"`
					OfferType           string   `json:"offer_type,omitempty"`
					RefreshInterval     int      `json:"refresh_interval,omitempty"`
					RefreshIntervalUnit string   `json:"refresh_interval_unit,omitempty"`
					Tags                []string `json:"tags,omitempty"`
					IconURL             string   `json:"icon_url,omitempty"`
					PaymentMethods      []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
					Data struct {
						Superlike struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"superlike,omitempty"`
						Boost struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"boost,omitempty"`
						HideAds struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_ads,omitempty"`
						HideAge struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_age,omitempty"`
						HideDistance struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_distance,omitempty"`
						Like struct {
							Type string `json:"type,omitempty"`
						} `json:"like,omitempty"`
						ControlWhoSeesYou struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_sees_you,omitempty"`
						Passport struct {
							Type string `json:"type,omitempty"`
						} `json:"passport,omitempty"`
						Rewind struct {
							Type string `json:"type,omitempty"`
						} `json:"rewind,omitempty"`
						SuperboostAlcAccess struct {
							Type string `json:"type,omitempty"`
						} `json:"superboost_alc_access,omitempty"`
						ControlWhoYouSee struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_you_see,omitempty"`
						Toppicks struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"toppicks,omitempty"`
						ToppicksAlcAccess struct {
							Type string `json:"type,omitempty"`
						} `json:"toppicks_alc_access,omitempty"`
						LikesYou struct {
							Type string `json:"type,omitempty"`
						} `json:"likes_you,omitempty"`
					} `json:"data,omitempty"`
					Ordering struct {
						Carousel   []string `json:"carousel,omitempty"`
						PlusScreen []string `json:"plus_screen,omitempty"`
					} `json:"ordering,omitempty"`
					SubPageData struct {
						Cta      string `json:"cta,omitempty"`
						Termed   bool   `json:"termed,omitempty"`
						Sections []struct {
							Type     string `json:"type,omitempty"`
							Heading  string `json:"heading,omitempty"`
							Benefits []struct {
								Key         string `json:"key,omitempty"`
								Heading     string `json:"heading,omitempty"`
								Included    bool   `json:"included,omitempty"`
								Deeplink    string `json:"deeplink,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"benefits,omitempty"`
						} `json:"sections,omitempty"`
					} `json:"sub_page_data,omitempty"`
					GoldUpsell struct {
						UpsellSkus   []string `json:"upsell_skus,omitempty"`
						UpsellLayout string   `json:"upsell_layout,omitempty"`
					} `json:"gold_upsell,omitempty"`
				} `json:"merchandising,omitempty"`
			} `json:"gold,omitempty"`
			Platinum struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount              int      `json:"amount,omitempty"`
					OfferType           string   `json:"offer_type,omitempty"`
					RefreshInterval     int      `json:"refresh_interval,omitempty"`
					RefreshIntervalUnit string   `json:"refresh_interval_unit,omitempty"`
					Tags                []string `json:"tags,omitempty"`
					IconURL             string   `json:"icon_url,omitempty"`
					PaymentMethods      []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
					Data struct {
						Superlike struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"superlike,omitempty"`
						Boost struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"boost,omitempty"`
						HideAds struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_ads,omitempty"`
						HideAge struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_age,omitempty"`
						HideDistance struct {
							Type string `json:"type,omitempty"`
						} `json:"hide_distance,omitempty"`
						Like struct {
							Type string `json:"type,omitempty"`
						} `json:"like,omitempty"`
						ControlWhoSeesYou struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_sees_you,omitempty"`
						Passport struct {
							Type string `json:"type,omitempty"`
						} `json:"passport,omitempty"`
						Rewind struct {
							Type string `json:"type,omitempty"`
						} `json:"rewind,omitempty"`
						SuperboostAlcAccess struct {
							Type string `json:"type,omitempty"`
						} `json:"superboost_alc_access,omitempty"`
						ControlWhoYouSee struct {
							Type string `json:"type,omitempty"`
						} `json:"control_who_you_see,omitempty"`
						Toppicks struct {
							Type        string `json:"type,omitempty"`
							RenewalData struct {
								Balance             int    `json:"balance,omitempty"`
								RefreshInterval     int    `json:"refresh_interval,omitempty"`
								RefreshIntervalUnit string `json:"refresh_interval_unit,omitempty"`
							} `json:"renewal_data,omitempty"`
						} `json:"toppicks,omitempty"`
						ToppicksAlcAccess struct {
							Type string `json:"type,omitempty"`
						} `json:"toppicks_alc_access,omitempty"`
						LikesYou struct {
							Type string `json:"type,omitempty"`
						} `json:"likes_you,omitempty"`
						SuperlikeAttachMessage struct {
							Type string `json:"type,omitempty"`
						} `json:"superlike_attach_message,omitempty"`
						MyLikesLookback struct {
							Type     string `json:"type,omitempty"`
							Duration int    `json:"duration,omitempty"`
						} `json:"my_likes_lookback,omitempty"`
						PriorityLikes struct {
							Type string `json:"type,omitempty"`
						} `json:"priority_likes,omitempty"`
					} `json:"data,omitempty"`
					Ordering struct {
						Carousel   []string `json:"carousel,omitempty"`
						PlusScreen []string `json:"plus_screen,omitempty"`
					} `json:"ordering,omitempty"`
					SubPageData struct {
						Cta      string `json:"cta,omitempty"`
						Termed   bool   `json:"termed,omitempty"`
						Sections []struct {
							Type     string `json:"type,omitempty"`
							Heading  string `json:"heading,omitempty"`
							Benefits []struct {
								Key         string `json:"key,omitempty"`
								Heading     string `json:"heading,omitempty"`
								Included    bool   `json:"included,omitempty"`
								Deeplink    string `json:"deeplink,omitempty"`
								Description string `json:"description,omitempty"`
							} `json:"benefits,omitempty"`
						} `json:"sections,omitempty"`
					} `json:"sub_page_data,omitempty"`
				} `json:"merchandising,omitempty"`
			} `json:"platinum,omitempty"`
			Boost struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount         int      `json:"amount,omitempty"`
					OfferType      string   `json:"offer_type,omitempty"`
					Tags           []string `json:"tags,omitempty"`
					IconURL        string   `json:"icon_url,omitempty"`
					PaymentMethods []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
					Upsell string `json:"upsell,omitempty"`
				} `json:"merchandising,omitempty"`
			} `json:"boost,omitempty"`
			Readreceipt struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount         int      `json:"amount,omitempty"`
					OfferType      string   `json:"offer_type,omitempty"`
					Tags           []string `json:"tags,omitempty"`
					IconURL        string   `json:"icon_url,omitempty"`
					PaymentMethods []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
				} `json:"merchandising,omitempty"`
			} `json:"readreceipt,omitempty"`
			Superboost struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount         int      `json:"amount,omitempty"`
					OfferType      string   `json:"offer_type,omitempty"`
					Tags           []string `json:"tags,omitempty"`
					Duration       int      `json:"duration,omitempty"`
					IconURL        string   `json:"icon_url,omitempty"`
					PaymentMethods []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
				} `json:"merchandising,omitempty"`
			} `json:"superboost,omitempty"`
			Superlike struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount         int      `json:"amount,omitempty"`
					OfferType      string   `json:"offer_type,omitempty"`
					Tags           []string `json:"tags,omitempty"`
					IconURL        string   `json:"icon_url,omitempty"`
					PaymentMethods []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
					Upsell string `json:"upsell,omitempty"`
				} `json:"merchandising,omitempty"`
			} `json:"superlike,omitempty"`
			Primetimeboost struct {
				PurchaseType string `json:"purchase_type,omitempty"`
				ProductData  []struct {
					Amount         int      `json:"amount,omitempty"`
					OfferType      string   `json:"offer_type,omitempty"`
					Tags           []string `json:"tags,omitempty"`
					IconURL        string   `json:"icon_url,omitempty"`
					PaymentMethods []struct {
						Platform   string  `json:"platform,omitempty"`
						SkuID      string  `json:"sku_id,omitempty"`
						Discount   int     `json:"discount,omitempty"`
						RequireZip bool    `json:"require_zip,omitempty"`
						Price      float64 `json:"price,omitempty"`
						IsVat      bool    `json:"is_vat,omitempty"`
						TaxRate    float64 `json:"tax_rate,omitempty"`
						Currency   string  `json:"currency,omitempty"`
					} `json:"payment_methods,omitempty"`
					OriginalProductID string `json:"original_product_id,omitempty"`
				} `json:"product_data,omitempty"`
				Merchandising struct {
				} `json:"merchandising,omitempty"`
			} `json:"primetimeboost,omitempty"`
		} `json:"offerings,omitempty"`
		FeatureAccess struct {
			LikesYou struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"likes_you,omitempty"`
			Passport struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"passport,omitempty"`
			PreferenceFilters struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"preference_filters,omitempty"`
			Rewind struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"rewind,omitempty"`
			Primetimeboost struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"primetimeboost,omitempty"`
			LikesYouAnnuity struct {
				EntryPoints string `json:"entry_points,omitempty"`
			} `json:"likes_you_annuity,omitempty"`
			MatchExtension struct {
				Unlimited bool `json:"unlimited,omitempty"`
				Amount    int  `json:"amount,omitempty"`
			} `json:"match_extension,omitempty"`
		} `json:"feature_access,omitempty"`
		Paywalls []struct {
			InstanceID           string `json:"instanceID,omitempty"`
			TemplateID           string `json:"templateID,omitempty"`
			ProductType          string `json:"productType,omitempty"`
			EntryPoint           string `json:"entryPoint,omitempty"`
			CarouselSubscription struct {
				SkuCard struct {
					IconURL struct {
						NewSub  string `json:"newSub,omitempty"`
						Upgrade string `json:"upgrade,omitempty"`
					} `json:"iconUrl,omitempty"`
					MerchandisingTextColor string `json:"merchandisingTextColor,omitempty"`
					BorderColor            struct {
						Unselected string `json:"unselected,omitempty"`
						Selected   string `json:"selected,omitempty"`
					} `json:"borderColor,omitempty"`
				} `json:"skuCard,omitempty"`
				Title struct {
					TextColor string `json:"textColor,omitempty"`
					Text      struct {
						Superlike          string `json:"superlike,omitempty"`
						Like               string `json:"like,omitempty"`
						HideAgeAndDistance string `json:"hide_age_and_distance,omitempty"`
						Toppicks           string `json:"toppicks,omitempty"`
						LikesYou           string `json:"likes_you,omitempty"`
						Default            string `json:"default,omitempty"`
						LikesYouFiltering  string `json:"likes_you_filtering,omitempty"`
						ControlWhoSeesYou  string `json:"control_who_sees_you,omitempty"`
						Passport           string `json:"passport,omitempty"`
						Rewind             string `json:"rewind,omitempty"`
						Boost              string `json:"boost,omitempty"`
						HideAds            string `json:"hide_ads,omitempty"`
						PreferenceFilters  string `json:"preference_filters,omitempty"`
						MyType             string `json:"my_type,omitempty"`
					} `json:"text,omitempty"`
				} `json:"title,omitempty"`
				Button struct {
					Text       string `json:"text,omitempty"`
					Background struct {
						Degree       int `json:"degree,omitempty"`
						GradientInfo []struct {
							Color    string  `json:"color,omitempty"`
							Position float64 `json:"position,omitempty"`
						} `json:"gradientInfo,omitempty"`
						Type string `json:"type,omitempty"`
					} `json:"background,omitempty"`
					TextColor string `json:"textColor,omitempty"`
				} `json:"button,omitempty"`
				Header struct {
					Background struct {
						Type         string  `json:"type,omitempty"`
						Degree       float64 `json:"degree,omitempty"`
						GradientInfo []struct {
							Color    string  `json:"color,omitempty"`
							Position float64 `json:"position,omitempty"`
						} `json:"gradientInfo,omitempty"`
					} `json:"background,omitempty"`
					BorderColor string `json:"borderColor,omitempty"`
					IconURL     string `json:"iconUrl,omitempty"`
				} `json:"header,omitempty"`
				DisclosureText string `json:"disclosureText,omitempty"`
				DiscountTag    string `json:"discountTag,omitempty"`
			} `json:"carouselSubscription,omitempty"`
			CarouselSubscriptionV2 struct {
				Header struct {
					Background struct {
						Type     string `json:"type,omitempty"`
						Name     string `json:"name,omitempty"`
						Fallback string `json:"fallback,omitempty"`
					} `json:"background,omitempty"`
					IconURL struct {
						Default string `json:"default,omitempty"`
						Dark    string `json:"dark,omitempty"`
					} `json:"iconUrl,omitempty"`
				} `json:"header,omitempty"`
				Title struct {
					Text struct {
						Default            string `json:"default,omitempty"`
						Like               string `json:"like,omitempty"`
						HideAgeAndDistance string `json:"hide_age_and_distance,omitempty"`
						ControlWhoSeesYou  string `json:"control_who_sees_you,omitempty"`
						Passport           string `json:"passport,omitempty"`
						Rewind             string `json:"rewind,omitempty"`
						HideAds            string `json:"hide_ads,omitempty"`
					} `json:"text,omitempty"`
					TextColor struct {
						Type     string `json:"type,omitempty"`
						Name     string `json:"name,omitempty"`
						Fallback string `json:"fallback,omitempty"`
					} `json:"textColor,omitempty"`
				} `json:"title,omitempty"`
				DisclosureText string `json:"disclosureText,omitempty"`
				DiscountTag    string `json:"discountTag,omitempty"`
				SkuCard        struct {
					BorderColor struct {
						Selected struct {
							Type     string `json:"type,omitempty"`
							Name     string `json:"name,omitempty"`
							Fallback string `json:"fallback,omitempty"`
						} `json:"selected,omitempty"`
						Unselected struct {
							Type     string `json:"type,omitempty"`
							Name     string `json:"name,omitempty"`
							Fallback string `json:"fallback,omitempty"`
						} `json:"unselected,omitempty"`
					} `json:"borderColor,omitempty"`
					MerchandisingTextColor struct {
						Type     string `json:"type,omitempty"`
						Name     string `json:"name,omitempty"`
						Fallback string `json:"fallback,omitempty"`
					} `json:"merchandisingTextColor,omitempty"`
					IconURL struct {
						NewSub struct {
							Default string `json:"default,omitempty"`
							Dark    string `json:"dark,omitempty"`
						} `json:"newSub,omitempty"`
						Upgrade struct {
							Default string `json:"default,omitempty"`
							Dark    string `json:"dark,omitempty"`
						} `json:"upgrade,omitempty"`
					} `json:"iconUrl,omitempty"`
				} `json:"skuCard,omitempty"`
				Button struct {
					Text      string `json:"text,omitempty"`
					TextColor struct {
						Type     string `json:"type,omitempty"`
						Name     string `json:"name,omitempty"`
						Fallback string `json:"fallback,omitempty"`
					} `json:"textColor,omitempty"`
					Background struct {
						Type     string `json:"type,omitempty"`
						Name     string `json:"name,omitempty"`
						Fallback string `json:"fallback,omitempty"`
					} `json:"background,omitempty"`
				} `json:"button,omitempty"`
			} `json:"carouselSubscriptionV2,omitempty"`
			CarouselB struct {
				HeroImage struct {
					Kind string `json:"kind,omitempty"`
					URL  string `json:"url,omitempty"`
				} `json:"heroImage,omitempty"`
				UpsellPrimary struct {
					Title                 string `json:"title,omitempty"`
					Subtitle              string `json:"subtitle,omitempty"`
					ProductType           string `json:"productType,omitempty"`
					Deeplink              string `json:"deeplink,omitempty"`
					HeaderBackgroundColor string `json:"headerBackgroundColor,omitempty"`
					BorderColor           string `json:"borderColor,omitempty"`
					Button                struct {
						Text       string `json:"text,omitempty"`
						TextColor  string `json:"textColor,omitempty"`
						Background struct {
							Type  string `json:"type,omitempty"`
							Color string `json:"color,omitempty"`
						} `json:"background,omitempty"`
						BorderColor string `json:"borderColor,omitempty"`
					} `json:"button,omitempty"`
					ImageURL string `json:"imageUrl,omitempty"`
				} `json:"upsellPrimary,omitempty"`
				Title       string `json:"title,omitempty"`
				Body        string `json:"body,omitempty"`
				DiscountTag string `json:"discountTag,omitempty"`
			} `json:"carouselB,omitempty"`
		} `json:"paywalls,omitempty"`
		Account struct {
			AccountPhoneNumber string `json:"account_phone_number,omitempty"`
			AppleIDLinked      bool   `json:"apple_id_linked,omitempty"`
			FacebookIDLinked   bool   `json:"facebook_id_linked,omitempty"`
			GoogleIDLinked     bool   `json:"google_id_linked,omitempty"`
			LineIDLinked       bool   `json:"line_id_linked,omitempty"`
			IsEmailVerified    bool   `json:"is_email_verified,omitempty"`
			AccountEmail       string `json:"account_email,omitempty"`
		} `json:"account,omitempty"`
		User struct {
			ID           string    `json:"_id,omitempty"`
			AgeFilterMax int       `json:"age_filter_max,omitempty"`
			AgeFilterMin int       `json:"age_filter_min,omitempty"`
			Bio          string    `json:"bio,omitempty"`
			BirthDate    time.Time `json:"birth_date,omitempty"`
			CreateDate   time.Time `json:"create_date,omitempty"`
			CrmID        string    `json:"crm_id,omitempty"`
			PosInfo      struct {
				Country struct {
					Name   string `json:"name,omitempty"`
					Cc     string `json:"cc,omitempty"`
					Alpha3 string `json:"alpha3,omitempty"`
				} `json:"country,omitempty"`
				Timezone string `json:"timezone,omitempty"`
			} `json:"pos_info,omitempty"`
			Discoverable   bool `json:"discoverable,omitempty"`
			DistanceFilter int  `json:"distance_filter,omitempty"`
			GlobalMode     struct {
				IsEnabled           bool   `json:"is_enabled,omitempty"`
				DisplayLanguage     string `json:"display_language,omitempty"`
				LanguagePreferences []struct {
					Language   string `json:"language,omitempty"`
					IsSelected bool   `json:"is_selected,omitempty"`
				} `json:"language_preferences,omitempty"`
			} `json:"global_mode,omitempty"`
			Gender      int `json:"gender,omitempty"`
			AllInGender []struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"all_in_gender,omitempty"`
			GenderFilter        int    `json:"gender_filter,omitempty"`
			ShowGenderOnProfile bool   `json:"show_gender_on_profile,omitempty"`
			Name                string `json:"name,omitempty"`
			Photos              []struct {
				ID       string `json:"id,omitempty"`
				CropInfo struct {
					User struct {
						WidthPct   int     `json:"width_pct,omitempty"`
						XOffsetPct int     `json:"x_offset_pct,omitempty"`
						HeightPct  float64 `json:"height_pct,omitempty"`
						YOffsetPct float64 `json:"y_offset_pct,omitempty"`
					} `json:"user,omitempty"`
					Algo struct {
						WidthPct   float64 `json:"width_pct,omitempty"`
						XOffsetPct float64 `json:"x_offset_pct,omitempty"`
						HeightPct  float64 `json:"height_pct,omitempty"`
						YOffsetPct float64 `json:"y_offset_pct,omitempty"`
					} `json:"algo,omitempty"`
					ProcessedByBullseye bool `json:"processed_by_bullseye,omitempty"`
					UserCustomized      bool `json:"user_customized,omitempty"`
					FacesCount          int  `json:"faces_count,omitempty"`
				} `json:"crop_info,omitempty"`
				URL            string `json:"url,omitempty"`
				FbID           string `json:"fbId,omitempty"`
				ProcessedFiles []struct {
					URL    string `json:"url,omitempty"`
					Height int    `json:"height,omitempty"`
					Width  int    `json:"width,omitempty"`
				} `json:"processedFiles,omitempty"`
				Assets    []any  `json:"assets,omitempty"`
				MediaType string `json:"media_type,omitempty"`
			} `json:"photos,omitempty"`
			PhotosProcessing      bool      `json:"photos_processing,omitempty"`
			PhotoOptimizerEnabled bool      `json:"photo_optimizer_enabled,omitempty"`
			PingTime              time.Time `json:"ping_time,omitempty"`
			Jobs                  []any     `json:"jobs,omitempty"`
			Schools               []any     `json:"schools,omitempty"`
			Badges                []struct {
				Type string `json:"type,omitempty"`
			} `json:"badges,omitempty"`
			PhoneID      string `json:"phone_id,omitempty"`
			InterestedIn []int  `json:"interested_in,omitempty"`
			Pos          struct {
				Lat float64 `json:"lat,omitempty"`
				Lon float64 `json:"lon,omitempty"`
			} `json:"pos,omitempty"`
			BillingInfo struct {
				SupportedPaymentMethods []string `json:"supported_payment_methods,omitempty"`
			} `json:"billing_info,omitempty"`
			AutoplayVideo        string `json:"autoplay_video,omitempty"`
			TopPicksDiscoverable bool   `json:"top_picks_discoverable,omitempty"`
			PhotoTaggingEnabled  bool   `json:"photo_tagging_enabled,omitempty"`
			City                 struct {
				Name   string `json:"name,omitempty"`
				Region string `json:"region,omitempty"`
			} `json:"city,omitempty"`
			ShowOrientationOnProfile bool `json:"show_orientation_on_profile,omitempty"`
			ShowSameOrientationFirst struct {
				Checked          bool `json:"checked,omitempty"`
				ShouldShowOption bool `json:"should_show_option,omitempty"`
			} `json:"show_same_orientation_first,omitempty"`
			SexualOrientations []struct {
				ID   string `json:"id,omitempty"`
				Name string `json:"name,omitempty"`
			} `json:"sexual_orientations,omitempty"`
			UserInterests struct {
				SelectedInterests []struct {
					ID   string `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"selected_interests,omitempty"`
				AvailableInterests []struct {
					ID   string `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"available_interests,omitempty"`
				MinInterests int `json:"min_interests,omitempty"`
				MaxInterests int `json:"max_interests,omitempty"`
			} `json:"user_interests,omitempty"`
			RecommendedSortDiscoverable bool   `json:"recommended_sort_discoverable,omitempty"`
			SelfieVerification          string `json:"selfie_verification,omitempty"`
			NoonlightProtected          bool   `json:"noonlight_protected,omitempty"`
			SyncSwipeEnabled            bool   `json:"sync_swipe_enabled,omitempty"`
			SparksQuizzes               []struct {
				SectionID   string `json:"section_id,omitempty"`
				SectionName string `json:"section_name,omitempty"`
				Quizzes     []struct {
					ID      string `json:"id,omitempty"`
					Name    string `json:"name,omitempty"`
					Answers []any  `json:"answers,omitempty"`
				} `json:"quizzes,omitempty"`
			} `json:"sparks_quizzes,omitempty"`
			SelectedDescriptors []struct {
				ID         string `json:"id,omitempty"`
				Name       string `json:"name,omitempty"`
				Type       string `json:"type,omitempty"`
				Visibility string `json:"visibility,omitempty"`
				IconURL    string `json:"icon_url,omitempty"`
				IconUrls   []struct {
					URL     string `json:"url,omitempty"`
					Quality string `json:"quality,omitempty"`
					Width   int    `json:"width,omitempty"`
					Height  int    `json:"height,omitempty"`
				} `json:"icon_urls,omitempty"`
				SectionID        string `json:"section_id,omitempty"`
				SectionName      string `json:"section_name,omitempty"`
				ChoiceSelections []struct {
					ID       string `json:"id,omitempty"`
					Name     string `json:"name,omitempty"`
					Style    string `json:"style,omitempty"`
					Emoji    string `json:"emoji,omitempty"`
					IconUrls []struct {
						URL     string `json:"url,omitempty"`
						Quality string `json:"quality,omitempty"`
						Width   int    `json:"width,omitempty"`
						Height  int    `json:"height,omitempty"`
					} `json:"icon_urls,omitempty"`
				} `json:"choice_selections,omitempty"`
			} `json:"selected_descriptors,omitempty"`
			RequestVerificationEnabled bool `json:"request_verification_enabled,omitempty"`
			PreferenceFilters          struct {
			} `json:"preference_filters,omitempty"`
			MmEnabled   bool `json:"mm_enabled,omitempty"`
			UserPrompts struct {
				SectionName string `json:"section_name,omitempty"`
				MaxPrompts  int    `json:"max_prompts,omitempty"`
				Prompts     []any  `json:"prompts,omitempty"`
			} `json:"user_prompts,omitempty"`
		} `json:"user,omitempty"`
		Boost struct {
			Allotment                int    `json:"allotment,omitempty"`
			AllotmentUsed            int    `json:"allotment_used,omitempty"`
			AllotmentRemaining       int    `json:"allotment_remaining,omitempty"`
			InternalRemaining        int    `json:"internal_remaining,omitempty"`
			IsSuperBoost             bool   `json:"is_super_boost,omitempty"`
			PurchasedRemaining       int    `json:"purchased_remaining,omitempty"`
			Remaining                int    `json:"remaining,omitempty"`
			Duration                 int    `json:"duration,omitempty"`
			BoostRefreshAmount       int    `json:"boost_refresh_amount,omitempty"`
			BoostRefreshInterval     int    `json:"boost_refresh_interval,omitempty"`
			BoostRefreshIntervalUnit string `json:"boost_refresh_interval_unit,omitempty"`
			CompoundBoost            []struct {
				Quantity int `json:"quantity,omitempty"`
				Duration int `json:"duration,omitempty"`
			} `json:"compound_boost,omitempty"`
			BoostEnded     bool   `json:"boost_ended,omitempty"`
			ResultViewedAt int64  `json:"result_viewed_at,omitempty"`
			BoostID        string `json:"boost_id,omitempty"`
			Multiplier     int    `json:"multiplier,omitempty"`
			ConsumedFrom   int    `json:"consumed_from,omitempty"`
		} `json:"boost,omitempty"`
		Campaigns struct {
			Campaigns []any `json:"campaigns,omitempty"`
		} `json:"campaigns,omitempty"`
		Compliance struct {
			ThirdPartyDataCollectionEnabled bool `json:"third_party_data_collection_enabled,omitempty"`
		} `json:"compliance,omitempty"`
		EmailSettings struct {
			Email         string `json:"email,omitempty"`
			EmailSettings struct {
				Promotions bool `json:"promotions,omitempty"`
				Messages   bool `json:"messages,omitempty"`
				NewMatches bool `json:"new_matches,omitempty"`
			} `json:"email_settings,omitempty"`
		} `json:"email_settings,omitempty"`
		Experiences struct {
			ShowExperiences  any   `json:"show_experiences,omitempty"`
			CompletedStories []any `json:"completed_stories,omitempty"`
			Series           struct {
				Episodes []any `json:"episodes,omitempty"`
			} `json:"series,omitempty"`
		} `json:"experiences,omitempty"`
		Instagram struct {
		} `json:"instagram,omitempty"`
		Likes struct {
			LikesRemaining int `json:"likes_remaining,omitempty"`
		} `json:"likes,omitempty"`
		Onboarding struct {
			Tutorials []struct {
				Name     string `json:"name,omitempty"`
				Selected int    `json:"selected,omitempty"`
				Assets   []struct {
					Type         string `json:"type,omitempty"`
					URL          string `json:"url,omitempty"`
					Name         string `json:"name,omitempty"`
					Age          int    `json:"age,omitempty"`
					MutualsCount int    `json:"mutuals_count,omitempty"`
				} `json:"assets,omitempty"`
			} `json:"tutorials,omitempty"`
		} `json:"onboarding,omitempty"`
		Travel struct {
			IsTraveling bool `json:"is_traveling,omitempty"`
		} `json:"travel,omitempty"`
		PlusControl struct {
			DiscoverableParty string `json:"discoverable_party,omitempty"`
			HideAds           bool   `json:"hide_ads,omitempty"`
			HideAge           bool   `json:"hide_age,omitempty"`
			HideDistance      bool   `json:"hide_distance,omitempty"`
			Blend             string `json:"blend,omitempty"`
		} `json:"plus_control,omitempty"`
		Purchase struct {
			Purchases           []any `json:"purchases,omitempty"`
			SubscriptionExpired bool  `json:"subscription_expired,omitempty"`
		} `json:"purchase,omitempty"`
		Readreceipts struct {
			InternalRemaining  int `json:"internal_remaining,omitempty"`
			PurchasedRemaining int `json:"purchased_remaining,omitempty"`
			Remaining          int `json:"remaining,omitempty"`
		} `json:"readreceipts,omitempty"`
		Spotify struct {
			SpotifyConnected  bool `json:"spotify_connected,omitempty"`
			SpotifyThemeTrack struct {
				ID    string `json:"id,omitempty"`
				Name  string `json:"name,omitempty"`
				Album struct {
					ID     string `json:"id,omitempty"`
					Name   string `json:"name,omitempty"`
					Images []struct {
						Height int    `json:"height,omitempty"`
						Width  int    `json:"width,omitempty"`
						URL    string `json:"url,omitempty"`
					} `json:"images,omitempty"`
				} `json:"album,omitempty"`
				Artists []struct {
					ID   string `json:"id,omitempty"`
					Name string `json:"name,omitempty"`
				} `json:"artists,omitempty"`
				PreviewURL string `json:"preview_url,omitempty"`
				URI        string `json:"uri,omitempty"`
			} `json:"spotify_theme_track,omitempty"`
		} `json:"spotify,omitempty"`
		SuperLikes struct {
			Remaining                    int    `json:"remaining,omitempty"`
			AlcRemaining                 int    `json:"alc_remaining,omitempty"`
			NewAlcRemaining              int    `json:"new_alc_remaining,omitempty"`
			Allotment                    int    `json:"allotment,omitempty"`
			SuperlikeRefreshAmount       int    `json:"superlike_refresh_amount,omitempty"`
			SuperlikeRefreshInterval     int    `json:"superlike_refresh_interval,omitempty"`
			SuperlikeRefreshIntervalUnit string `json:"superlike_refresh_interval_unit,omitempty"`
			ResetsAt                     any    `json:"resets_at,omitempty"`
		} `json:"super_likes,omitempty"`
		TinderU struct {
			Status   string `json:"status,omitempty"`
			StatusV2 string `json:"status_v2,omitempty"`
		} `json:"tinder_u,omitempty"`
		TopPhoto struct {
		} `json:"top_photo,omitempty"`
		Tutorials            []string `json:"tutorials,omitempty"`
		AvailableDescriptors []struct {
			Descriptors []struct {
				ID      string `json:"id,omitempty"`
				Name    string `json:"name,omitempty"`
				Prompt  string `json:"prompt,omitempty"`
				Type    string `json:"type,omitempty"`
				Choices []struct {
					ID       string `json:"id,omitempty"`
					Name     string `json:"name,omitempty"`
					Style    string `json:"style,omitempty"`
					Emoji    string `json:"emoji,omitempty"`
					IconUrls []struct {
						URL     string `json:"url,omitempty"`
						Quality string `json:"quality,omitempty"`
						Width   int    `json:"width,omitempty"`
						Height  int    `json:"height,omitempty"`
					} `json:"icon_urls,omitempty"`
				} `json:"choices,omitempty"`
				IconURL  string `json:"icon_url,omitempty"`
				IconUrls []struct {
					URL     string `json:"url,omitempty"`
					Quality string `json:"quality,omitempty"`
					Width   int    `json:"width,omitempty"`
					Height  int    `json:"height,omitempty"`
				} `json:"icon_urls,omitempty"`
				SubPrompt                   string `json:"sub_prompt,omitempty"`
				SectionID                   string `json:"section_id,omitempty"`
				SectionName                 string `json:"section_name,omitempty"`
				DiscoveryPreferencesEnabled bool   `json:"discovery_preferences_enabled,omitempty"`
			} `json:"descriptors,omitempty"`
			SectionID   string `json:"section_id,omitempty"`
			SectionName string `json:"section_name,omitempty"`
			Prompt      string `json:"prompt,omitempty"`
		} `json:"available_descriptors,omitempty"`
		ProfileMeter struct {
			PercentAchieved      int `json:"percent_achieved,omitempty"`
			IncompleteComponents []struct {
				Key         string `json:"key,omitempty"`
				RedDot      bool   `json:"red_dot,omitempty"`
				DisplayText string `json:"display_text,omitempty"`
				PillText    string `json:"pill_text,omitempty"`
			} `json:"incomplete_components,omitempty"`
		} `json:"profile_meter,omitempty"`
		MiscMerchandising struct {
			SubscriptionCardOrdering  []string `json:"subscription_card_ordering,omitempty"`
			LandingCard               string   `json:"landing_card,omitempty"`
			ControllaCarouselOrdering []string `json:"controlla_carousel_ordering,omitempty"`
		} `json:"misc_merchandising,omitempty"`
	} `json:"data,omitempty"`
}