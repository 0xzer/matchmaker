package response

type BucketResponse struct {
	Meta struct {
		Status int `json:"status,omitempty"`
	} `json:"meta,omitempty"`
	Data struct {
		Buckets struct {
			CorePhotoQuality                string `json:"core_photo_quality,omitempty"`
			IdentityReviveMigration         string `json:"identity_revive_migration,omitempty"`
			TinderRules                     string `json:"tinder_rules,omitempty"`
			ConsentManagementPlatformGlobal string `json:"consent_management_platform_global,omitempty"`
			TwoFactorAuth                   string `json:"two_factor_auth,omitempty"`
			OnboardingObsidianV1            string `json:"onboarding_obsidian_v1,omitempty"`
			HouseRulesUpdate                string `json:"house_rules_update,omitempty"`
			PushAuth                        string `json:"push_auth,omitempty"`
			BirthdayPageInOnboarding        string `json:"birthday_page_in_onboarding,omitempty"`
			EventsSdkMetricsPreAuth         string `json:"events_sdk_metrics_pre_auth,omitempty"`
			ProfileUserMigration            string `json:"profile_user_migration,omitempty"`
			PreAuthMarketingEmailConsent    string `json:"pre_auth_marketing_email_consent,omitempty"`
			UserInterestsAvailable          string `json:"user_interests_available,omitempty"`
			PostOnboardingObsidianV11       string `json:"post_onboarding_obsidian_v1.1,omitempty"`
			PhoneUpdate                     string `json:"phone_update,omitempty"`
			AuthOptions                     string `json:"auth_options,omitempty"`
			RelationshipIntentOnboarding    string `json:"relationship_intent_onboarding,omitempty"`
			UnderageAppealsOnboarding       string `json:"underage_appeals_onboarding,omitempty"`
			HubblePreonboarding             string `json:"hubble_preonboarding,omitempty"`
			AccountRecovery                 string `json:"account_recovery,omitempty"`
		} `json:"buckets,omitempty"`
		Levers struct {
			AuthV2 struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"auth_v2,omitempty"`
			OnboardingAgeSettingsVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.age_settings_variant,omitempty"`
			AppPinDiskCacheSizeLimit struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"app.pin_disk_cache_size_limit,omitempty"`
			OnboardingBirthdayCelebrationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.birthday_celebration_enabled,omitempty"`
			IdentitySmsDeviceCheckEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.sms_device_check_enabled,omitempty"`
			OnboardingIsPostOnboardingObsidianV1Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_post_onboarding_obsidian_v1_enabled,omitempty"`
			IdentityAuthV3PhoneUpdateEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.auth_v3_phone_update_enabled,omitempty"`
			IdentityIsPreOnboardingHubbleInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.is_pre_onboarding_hubble_instrumentation_enabled,omitempty"`
			ProfilePhotoQualityV2Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"profile.photo_quality_v2_enabled,omitempty"`
			AuthIsPreOnboardingHubbleInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"auth.is_pre_onboarding_hubble_instrumentation_enabled,omitempty"`
			IsThreadSafeAppStateEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"is_thread_safe_app_state_enabled,omitempty"`
			OnboardingIsAgeSettingsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_age_settings_enabled,omitempty"`
			TrustSelfieVerificationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"trust.selfie_verification_enabled,omitempty"`
			AppIsIosTwelveIssueReportingEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.is_ios_twelve_issue_reporting_enabled,omitempty"`
			ContentcreatorIsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"contentcreator.is_enabled,omitempty"`
			OnboardingRelationshipIntentSkipEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.relationship_intent_skip_enabled,omitempty"`
			CoreVaultEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"core.vault_enabled,omitempty"`
			EventsHubbleEventTrackerIsEnabledAndroid struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.hubble_event_tracker_is_enabled_android,omitempty"`
			AppNetworkPerfV2Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.network_perf_v2_enabled,omitempty"`
			ConsentManagementAppsflyerIsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"consent_management.appsflyer_is_enabled,omitempty"`
			SelfieV2BiometricConsentDescriptionTwoWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_biometric_consent_description_two_web,omitempty"`
			SelfieV2ModalUnverifyDescription struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_modal_unverify_description,omitempty"`
			SelfieV2BiometricConsentDescriptionTwo struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_biometric_consent_description_two,omitempty"`
			OnboardingIsNudityPromptEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_nudity_prompt_enabled,omitempty"`
			EventsIsGrpcPublishEnabledAndroid struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_grpc_publish_enabled_android,omitempty"`
			IdentitySignInWithGoogleOneClickEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.sign_in_with_google_one_click_enabled,omitempty"`
			RecsIsRecsRefactorEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"recs.is_recs_refactor_enabled,omitempty"`
			OnboardingIsMultiPhotoDiscoverArchEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_multi_photo_discover_arch_enabled,omitempty"`
			EventsIsDataLossInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_data_loss_instrumentation_enabled,omitempty"`
			OnboardingRefactorV1Batch1 struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.refactor_v1_batch_1,omitempty"`
			IdentityMarriageSupportProjectJapanDisclaimerEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.marriage_support_project_japan_disclaimer_enabled,omitempty"`
			OnboardingIsNudityDetectionEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_nudity_detection_enabled,omitempty"`
			AppIsNotificationWindowRefactorEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.is_notification_window_refactor_enabled,omitempty"`
			IdentityAuthOptions struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"identity.auth_options,omitempty"`
			SelfieV2ModalUnverifyDescriptionWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_modal_unverify_description_web,omitempty"`
			IdentitySuppressSplashExperiment struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.suppress_splash_experiment,omitempty"`
			EventsHubbleEventTrackerIsEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.hubble_event_tracker_is_enabled_ios,omitempty"`
			AppRemoveIntrospectiveEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.remove_introspective_enabled,omitempty"`
			EventsSdkIsEnabledWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_is_enabled_web,omitempty"`
			ExperimentalPreauthNoTreatment struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"experimental.preauth_no_treatment,omitempty"`
			EventsIsLoggerEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_logger_enabled_ios,omitempty"`
			TrustSelfieIntroVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"trust.selfie_intro_variant,omitempty"`
			OnboardingRefereeRedemptionEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.referee_redemption_enabled,omitempty"`
			OnboardingMinPhotoCountVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.min_photo_count_variant,omitempty"`
			EventsTrackerIsEnabledAndroid struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.tracker_is_enabled_android,omitempty"`
			EventsTrackerIsEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.tracker_is_enabled_ios,omitempty"`
			RegulatoryAffairsPreAuthMarketingEmailConsentType struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"regulatory_affairs.pre_auth_marketing_email_consent_type,omitempty"`
			OnboardingSkipPassionScreenEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.skip_passion_screen_enabled,omitempty"`
			ProfileProfileUserMigrationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"profile.profile_user_migration_enabled,omitempty"`
			EventsSdkKillswitchWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_killswitch_web,omitempty"`
			IdentityCaptchaGatedSmsTesting struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.captcha_gated_sms_testing,omitempty"`
			AppDeletePhotosInteractorVersion struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"app.delete_photos_interactor_version,omitempty"`
			IdentityNewHouseRulesDesign struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.new_house_rules_design,omitempty"`
			TrustSelfieVerificationFacemapEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"trust.selfie_verification_facemap_enabled,omitempty"`
			AppIsBugsnagEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.is_bugsnag_enabled,omitempty"`
			AgeVerificationPostOnboardingEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"age_verification.post_onboarding_enabled,omitempty"`
			EventsSdkIsEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_is_enabled_ios,omitempty"`
			OnboardingBirthdayConfirmationModal struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.birthday_confirmation_modal,omitempty"`
			OnboardingInterestStepEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.interest_step_enabled,omitempty"`
			OnboardingImageModerationURL struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"onboarding.image_moderation_url,omitempty"`
			OnboardingIsHubbleEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_hubble_enabled,omitempty"`
			IdentityAuthV3AccountRecoveryEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.auth_v3_account_recovery_enabled,omitempty"`
			EventsIsJSONEventTransportEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_json_event_transport_enabled,omitempty"`
			DevonTestLever struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"devon_test_lever,omitempty"`
			AppPinMemoryCacheSizeLimit struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"app.pin_memory_cache_size_limit,omitempty"`
			EventsIsAutomatedHubbleEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_automated_hubble_enabled,omitempty"`
			ProfileCustomGenderAvailable struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"profile.custom_gender_available,omitempty"`
			EventsIsLegacyMigrationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_legacy_migration_enabled,omitempty"`
			OnboardingMultiphotoRedesignWithTips struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.multiphoto_redesign_with_tips,omitempty"`
			OnboardingIsV2AnalyticsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_v2_analytics_enabled,omitempty"`
			OnboardingSmileDetectionVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.smile_detection_variant,omitempty"`
			OnboardingIsHubbleInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_hubble_instrumentation_enabled,omitempty"`
			OnboardingNameConfirmationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.name_confirmation_enabled,omitempty"`
			IdentityEmailMarketingOptOutEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.email_marketing_opt_out_enabled,omitempty"`
			OnboardingDistanceFilter struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.distance_filter,omitempty"`
			TrustBanScreensV2Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"trust.ban_screens_v2_enabled,omitempty"`
			ObsidianDarkModeBackgroundSwap struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"obsidian.dark_mode_background_swap,omitempty"`
			EventsSdkIsLoggerEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_is_logger_enabled_ios,omitempty"`
			IdentityGoogleAuthVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"identity.google_auth_variant,omitempty"`
			EventsSdkKillswitchAndroid struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_killswitch_android,omitempty"`
			OnboardingRefactorV1Batch2 struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.refactor_v1_batch_2,omitempty"`
			EventsSdkIsEnabledAndroid struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_is_enabled_android,omitempty"`
			ProfilePhotoUploadResolutionThreshold struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"profile.photo_upload_resolution_threshold,omitempty"`
			IdentityTwoFactorAuthEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.two_factor_auth_enabled,omitempty"`
			CaliforniaCcCheckbox struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"california_cc_checkbox,omitempty"`
			OnboardingAllInEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.all_in_enabled,omitempty"`
			EventsSdkErrorLogSamplingRate struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"events.sdk_error_log_sampling_rate,omitempty"`
			OnboardingDescriptorsV1Variant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.descriptors_v1_variant,omitempty"`
			OnboardingProfileElementsInterestsMinSelection struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.profile_elements_interests_min_selection,omitempty"`
			AppIsIosThirteenIssueReportingEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.is_ios_thirteen_issue_reporting_enabled,omitempty"`
			AppRetryMetaOnAuthRefresh struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.retry_meta_on_auth_refresh,omitempty"`
			AnalyticsFireworksToRapidfireMigrationTreatmentGroup struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"analytics.fireworks_to_rapidfire_migration_treatment_group,omitempty"`
			OnboardingIsOnboardingObsidianV1Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_onboarding_obsidian_v1_enabled,omitempty"`
			ConsentManagementIsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"consent_management.is_enabled,omitempty"`
			EventsLogErrorsWithSdkWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.log_errors_with_sdk_web,omitempty"`
			TrustAgeVerificationPostOnboardingInJapanEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"trust.age_verification_post_onboarding_in_japan_enabled,omitempty"`
			OnboardingAgeGateJumioAppealEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.age_gate_jumio_appeal_enabled,omitempty"`
			EventsSdkKillswitchIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.sdk_killswitch_ios,omitempty"`
			EventsIsGrpcPublishEnabledIos struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_grpc_publish_enabled_ios,omitempty"`
			SelfieVerificationPostOnboardingEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"selfie_verification_post_onboarding_enabled,omitempty"`
			EventsTrackerIsEnabledWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.tracker_is_enabled_web,omitempty"`
			IdentityTaglineUpdate struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.tagline_update,omitempty"`
			AppBugsnagNonFatalAllowPercentage struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"app.bugsnag_non_fatal_allow_percentage,omitempty"`
			IdentityAutoSignInEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.auto_sign_in_enabled,omitempty"`
			IdentityPushAuthEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.push_auth_enabled,omitempty"`
			IdentitySmsRequiresCaptchaVerification struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.sms_requires_captcha_verification,omitempty"`
			AppShouldProactivelyRegisterLevers struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.should_proactively_register_levers,omitempty"`
			SelfieV2BiometricConsentDescriptionOne struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_biometric_consent_description_one,omitempty"`
			SelfieV2BiometricConsentDescriptionOneWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"selfie_v2_biometric_consent_description_one_web,omitempty"`
			IdentityOneTapVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"identity.one_tap_variant,omitempty"`
			IdentityTokenGatewayServiceCutover struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"identity.token_gateway_service_cutover,omitempty"`
			EventsPushStatusPollerIsActive struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.push_status_poller_is_active,omitempty"`
			IsDataLossInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"is_data_loss_instrumentation_enabled,omitempty"`
			ConsentManagementCmpSdkDefaultValue struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"consent_management.cmp_sdk_default_value,omitempty"`
			OnboardingRefactorV1Batch3 struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.refactor_v1_batch_3,omitempty"`
			OnboardingProfileElementsInterestsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.profile_elements_interests_enabled,omitempty"`
			BugsnagIsEnabledWeb struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"bugsnag.is_enabled_web,omitempty"`
			OnboardingIsPreOnboardingObsidianV1Enabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_pre_onboarding_obsidian_v1_enabled,omitempty"`
			TrustSelfieConsentVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"trust.selfie_consent_variant,omitempty"`
			PlatformBugsnagEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"platform.bugsnag_enabled,omitempty"`
			ProfilePhotoUploadCompressionQuality struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"profile.photo_upload_compression_quality,omitempty"`
			BugsnagScopeFilters struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"bugsnag.scope_filters,omitempty"`
			OnboardingIsFaceFeaturesDetectionEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.is_face_features_detection_enabled,omitempty"`
			AppIsVcDeallocationCheckEnabledInProduction struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"app.is_vc_deallocation_check_enabled_in_production,omitempty"`
			GooglePhase2Cutover struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"google_phase_2_cutover,omitempty"`
			EventsIsPreOnboardingHubbleInstrumentationEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"events.is_pre_onboarding_hubble_instrumentation_enabled,omitempty"`
			OnboardingProfileElementsInterestsMaxSelection struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.profile_elements_interests_max_selection,omitempty"`
			AnalyticsSendEtlEventsAsProto struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"analytics.send_etl_events_as_proto,omitempty"`
			TrustSelfieVerificationPostOnboardingEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"trust.selfie_verification_post_onboarding_enabled,omitempty"`
			SessionIsPaginatedUpdatesEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"session.is_paginated_updates_enabled,omitempty"`
			AppInAppUpdateVersion struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"app.in_app_update_version,omitempty"`
			OnboardingRelationshipIntentEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"onboarding.relationship_intent_enabled,omitempty"`
			IdentitySignInWithGoogleEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value string `json:"value,omitempty"`
			} `json:"identity.sign_in_with_google_enabled,omitempty"`
			OnboardingImageModerationThreshold struct {
				PID   string  `json:"p_id,omitempty"`
				Value float64 `json:"value,omitempty"`
			} `json:"onboarding.image_moderation_threshold,omitempty"`
			PlatformAppdynamicsEnabled struct {
				PID   string `json:"p_id,omitempty"`
				Value bool   `json:"value,omitempty"`
			} `json:"platform.appdynamics_enabled,omitempty"`
			TrustSelfieV2BiometricConsentScreenVersion struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"trust.selfie_v2_biometric_consent_screen_version,omitempty"`
			EventsPushStatusPoller struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"events.push_status_poller,omitempty"`
			OnboardingFaceDetectionVariant struct {
				PID   string `json:"p_id,omitempty"`
				Value int    `json:"value,omitempty"`
			} `json:"onboarding.face_detection_variant,omitempty"`
			ProfilePhotoUploadAspectRatio struct {
				PID   string  `json:"p_id,omitempty"`
				Value float64 `json:"value,omitempty"`
			} `json:"profile.photo_upload_aspect_ratio,omitempty"`
		} `json:"levers,omitempty"`
		Routes []struct {
			Method       string `json:"method,omitempty"`
			Endpoint     string `json:"endpoint,omitempty"`
			SamplingRate int    `json:"sampling_rate,omitempty"`
		} `json:"routes,omitempty"`
	} `json:"data,omitempty"`
}