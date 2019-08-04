package main

import (
	"platform2.0-go-challenge/src/datalayer/repositories/audienceRepository"
	"platform2.0-go-challenge/src/datalayer/repositories/chartRepository"
	"platform2.0-go-challenge/src/datalayer/repositories/insightRepository"
	"platform2.0-go-challenge/src/datalayer/repositories/userRepository"
	"platform2.0-go-challenge/src/helpers/drivers"
	"platform2.0-go-challenge/src/servicelayer/services/assetService"
	"platform2.0-go-challenge/src/servicelayer/services/userService"
	"platform2.0-go-challenge/src/servicelayer/validators/audienceValidator"
	"platform2.0-go-challenge/src/servicelayer/validators/chartValidator"
	"platform2.0-go-challenge/src/servicelayer/validators/insightValidator"
	"platform2.0-go-challenge/src/servicelayer/validators/userValidator"
	"platform2.0-go-challenge/src/weblayer/controllers/assetController"
	"platform2.0-go-challenge/src/weblayer/controllers/userController"
)

var AssetController *assetController.AssetController
var UserController *userController.UserController

func wireDependencies() {
	db := drivers.ConnectPostgresDB()

	var audienceRepository audienceRepository.IAudienceRepository = audienceRepository.NewAudienceRepository(db)
	var chartRepository chartRepository.IChartRepository = chartRepository.NewChartRepository(db)
	var insightRepository insightRepository.IInsightRepository = insightRepository.NewInsightRepository(db)
	var userRepository userRepository.IUserRepository = userRepository.NewUserRepository(db)

	var audienceValidator audienceValidator.IAudienceValidator = audienceValidator.NewAudienceValidator()
	var chartValidator chartValidator.IChartValidator = chartValidator.NewChartValidator()
	var insightValidator insightValidator.IInsightValidator = insightValidator.NewInsightValidator()
	var userValidator userValidator.IUserValidator = userValidator.NewUserValidator()

	var assetService assetService.IAssetService = assetService.NewAssetService(
		audienceRepository,
		chartRepository,
		insightRepository,
		audienceValidator,
		chartValidator,
		insightValidator,
	)

	var userService userService.IUserService = userService.NewUserService(
		userRepository,
		userValidator,
	)

	AssetController = assetController.NewAssetController(assetService)
	UserController = userController.NewUserController(userService)
}
