package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/mndrix/golog"
)

func recomendarCarreras(m golog.Machine, aptitud, habilidad, interes string) []string {

	query := fmt.Sprintf("carrera(Fac, Carr, %s, %s, %s).", aptitud, habilidad, interes)

	results := []string{}
	solutions := m.ProveAll(query)

	for _, sol := range solutions {
		carrera := sol.ByName_("Carr").String()
		facultad := sol.ByName_("Fac").String()
		results = append(results, fmt.Sprintf("%s (%s)", carrera, facultad))
	}
	return results
}
func cargarProlog(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type PerfilEstudiante struct {
	Aptitud   string `json:"aptitud"`
	Habilidad string `json:"habilidad"`
	Interes   string `json:"interes"`
}

type CarreraRecomendada struct {
	Facultad string `json:"facultad"`
	Carrera  string `json:"carrera"`
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
		resultados := recomendarCarreras(m, perfil.Aptitud, perfil.Habilidad, perfil.Interes)
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
