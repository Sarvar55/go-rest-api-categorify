package controller

import (
	"categori/service"
	"categori/util"
	"github.com/gofiber/fiber/v2"
	"log"
)

type DomainInfoController struct {
	DomainService      *service.DomainService
	LocationService    *service.LocationService
}

func NewDomainInfoController(domainService *service.DomainService, locationService *service.LocationService) *DomainInfoController {
	return &DomainInfoController{
		DomainService:   domainService,
		LocationService: locationService,
	}
}

func (controller *DomainInfoController) GetDomainInfo(c *fiber.Ctx) error {
	domain := c.Params("domain")

	ipAddress, err := util.ResolveDomain(domain)
	if err != nil {
		return controller.handleError(c, "Alan adı çözümleme hatası:", "Domain adresi yanlış girilmiş olabilir")
	}

	ipDecimal, err := util.ConvertIPToDecimal(ipAddress[0])
	if err != nil {
		return controller.handleError(c, "IP dönüşüm hatası:", "Domain adresi yanlış girilmiş olabilir")
	}

	categoryDomain, err := controller.DomainService.GetLocationByDomainFromMongoDb(domain)
	if err != nil {
		return controller.handleError(c, "Konum sorgulama hatası:", "Bu lokasyonda veri bulunamadı")
	}

	if categoryDomain == nil {
		return c.JSON(fiber.Map{
			"success": "false",
			"message": "Belirtilen domain veritabanında bulunamadı",
		})
	}

	location, err := controller.LocationService.GetLocationByIPFromSQLLite(ipDecimal)
	if err != nil {
		return controller.handleError(c, "Konum sorgulama hatası:", "Bu lokasyonda veri bulunamadı")
	}

	response := fiber.Map{
		"ipaddress":    ipAddress,
		"countryName":  location.CountryName,
		"categoryName": categoryDomain.CategoryName,
		"domain":       domain,
	}

	return c.JSON(response)
}


func (controller *DomainInfoController) handleError(c *fiber.Ctx, logMessage, errorMessage string) error {
	log.Println(logMessage)
	res := fiber.Map{
		"success": "false",
		"message": errorMessage,
	}
	return c.JSON(res)
}
