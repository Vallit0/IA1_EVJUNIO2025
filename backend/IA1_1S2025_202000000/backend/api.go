package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mndrix/golog"
)

func recomendarCarreras(m golog.Machine, aptitud, habilidad1, interes1, habilidad2, interes2 string) []CarreraRecomendada {
	query := "carrera(Fac, Carr, Apt, Hab, Int)."
	solutions := m.ProveAll(query)

	results := []CarreraRecomendada{}

	for _, sol := range solutions {
		apt := sol.ByName_("Apt").String()
		hab := sol.ByName_("Hab").String()
		ints := sol.ByName_("Int").String()

		matchCount := 0
		if apt == aptitud {
			matchCount++
		}
		if hab == habilidad1 || hab == habilidad2 {
			matchCount++
		}
		if ints == interes1 || ints == interes2 {
			matchCount++
		}
		if habilidad1 == ints || habilidad2 == apt || interes1 == hab || interes2 == hab {
			matchCount++ // peso extra si hay cruce interesante
		}

		matchPercent := float64(matchCount) / 5.0 * 100.0

		results = append(results, CarreraRecomendada{
			Facultad: sol.ByName_("Fac").String(),
			Carrera:  sol.ByName_("Carr").String(),
			Match:    matchPercent,
		})
	}
	return results
}

// Agregar numeros para recomendacion
func cargarProlog(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type PerfilEstudiante struct {
	Aptitud    string `json:"aptitud"`
	Habilidad  string `json:"habilidad"`
	Interes    string `json:"interes"`
	Habilidad2 string `json:"habilidad2"`
	Interes2   string `json:"interes2"`
}

type CarreraRecomendada struct {
	Facultad string  `json:"facultad"`
	Carrera  string  `json:"carrera"`
	Match    float64 `json:"match"`
}

func main() {
	app := fiber.New()
	// resolver problema de CORS
	app.Use(func(c *fiber.Ctx) error {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "Content-Type")

		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusOK)
		}

		return c.Next()
	})

	programa := cargarProlog("./prolog/conocimiento.pl")
	m := golog.NewMachine()
	m = m.Consult(programa)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Servidor UniMatch funcionando 🧠")
	})

	app.Post("/recomendar", func(c *fiber.Ctx) error {
		var perfil PerfilEstudiante
		if err := c.BodyParser(&perfil); err != nil {
			return c.Status(400).SendString("Error de entrada.")
		}
		fmt.Println("Perfil recibido:", perfil)
		resultados := recomendarCarreras(m, perfil.Aptitud, perfil.Habilidad, perfil.Interes, perfil.Interes2, perfil.Habilidad2)
		if len(resultados) == 0 {
			return c.JSON(fiber.Map{"mensaje": "No se encontraron coincidencias."})
		}

		return c.JSON(fiber.Map{
			"recomendaciones": resultados,
		})
	})
	fmt.Println("Servidor UniMatch iniciado en el puerto 8080")
	app.Listen(":8080")
}
