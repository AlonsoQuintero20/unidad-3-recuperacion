package main

//Student: Alonso Rogelio Quintero Singh
//Email: quintero.020198@gmail.com

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

//DB la variable para la conexion a la base de datos
var DB *gorm.DB

//Celular es la estructura para la base de datos
type Celular struct {
	gorm.Model
	Precio      uint
	Descripcion string `gorm:"type:varchar(450);"`
	Marca       string `gorm:"type:varchar(450);"`
	Modelo      string `gorm:"type:varchar(450);"`
	Lanzamiento time.Time
}

func main() {
	DB, _ = gorm.Open("mysql", "root:@/moviles?charset=utf8&parseTime=True&loc=Local")

	defer DB.Close()

	DB.AutoMigrate(&Celular{})
	//var cel = Celular{Precio: 99, Descripcion: "hola", Marca: "samsumng", Modelo: "cualquiera", Lanzamiento: time.Now()}

	r := gin.Default()
	r.GET("/moviles/v1/", ObtenerCelulares)
	r.GET("/moviles/v1/:id", ObtenerCelular)
	r.POST("/moviles/v1/register", CrearCelular)
	r.PUT("/moviles/v1/:id", ActualizarCelular)
	r.DELETE("/moviles/v1/:id", EliminarCelular)

	r.Run(":8080")
}

// CrearCelular es la funcion que nos permite crear celulares
func CrearCelular(c *gin.Context) {
	var celular Celular
	c.BindJSON(&celular)

	DB.Create(&celular)
	c.JSON(200, celular)

}

// ObtenerCelulares es la funcion que nos permite obtener todos los celulares
func ObtenerCelulares(c *gin.Context) {
	var celulares []Celular

	if err := DB.Find(&celulares).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, celulares)
	}
}

// ObtenerCelular es la funcion que nos permite obtener 1 celulares mediante su Id
func ObtenerCelular(c *gin.Context) {
	id := c.Params.ByName("id")
	var celulares Celular
	if err := DB.Where("id = ?", id).First(&celulares).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, celulares)
	}
}

// EliminarCelular , funcion que permite eliminar un celular mediante su ID
func EliminarCelular(c *gin.Context) {
	id := c.Params.ByName("id")
	var celulares Celular
	d := DB.Where("id = ?", id).Delete(&celulares)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// ActualizarCelular funcion para actualizar celulares mediante id
func ActualizarCelular(c *gin.Context) {

	var celulares Celular
	id := c.Params.ByName("id")

	if err := DB.Where("id = ?", id).First(&celulares).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&celulares)

	DB.Save(&celulares)
	c.JSON(200, celulares)

}
