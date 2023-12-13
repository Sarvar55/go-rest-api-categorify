package controller

import (
	"categori/service"
	"categori/util"
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"log"
)

func GetDomainInfo(c *fiber.Ctx) error {
	domain := c.Params("domain")

	ipAddress, err := util.ResolveDomain(domain)
	if err != nil {
		log.Println("Alan adı çözümleme hatası:", err)
		res := fiber.Map{
			"success": "false",
			"message": "Domain adresi yanlış girilmiş olabilir",
		}
		return c.JSON(res)
	}

	log.Println(ipAddress)
	ipDecimal, err := util.ConvertIPToDecimal(ipAddress[0])
	log.Print(ipDecimal)
	if err != nil {
		log.Println("IP dönüşüm hatası:", err)
		res := fiber.Map{
			"success": "false",
			"message": "Domain adresi yanlış girilmiş olabilir",
		}
		return c.JSON(res)
	}

	// MongoDB'den kategori bilgisini al
	categoryDomain, err := service.GetLocationByDomainFromMongoDb(domain)
	if err != nil {
		log.Println("Konum sorgulama hatası:", err)
		res := fiber.Map{
			"success": "false",
			"message": "Bu lokasyonda veri bulunamadı",
		}
		return c.JSON(res)
	}

	if categoryDomain == nil {
		// categoryDomain nil ise, istediğiniz mesajı gönderin
		res := fiber.Map{
			"success": "false",
			"message": "Belirtilen domain veritabanında bulunamadı",
		}
		return c.JSON(res)
	}

	// SQLite'dan location bilgisini al
	location, err := service.GetLocationByIPFromSQLLite(ipDecimal)
	if err != nil {
		log.Println("Konum sorgulama hatası:", err)
		res := fiber.Map{
			"success": "false",
			"message": "Bu lokasyonda veri bulunamadı",
		}
		return c.JSON(res)
	}

	// Doğru bir şekilde response yapısını oluşturun
	response := fiber.Map{
		"ipaddress":    ipAddress,
		"countryName":  location.CountryName,
		"categoryName": categoryDomain.CategoryName,
		"domain":       domain,
	}

	return c.JSON(response)
}
