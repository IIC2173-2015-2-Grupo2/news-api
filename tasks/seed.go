//usr/bin/env go run $0 "$@"; exit
package tasks

// Remember change db connection. default is localhost
// for run seed, put in yout shell:
// chmod 744 script.go
// ./script.go

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jmcvetta/neoism"
	// "encoding/json"
)

func getOrCreateNode(db *neoism.Database, label string, key string, value string, props neoism.Props) *neoism.Node {
	res := []struct {
		// `json:` tags matches column names in query
		N neoism.Node
	}{}

	//MATCH (n:Tag { name: 'deporte' }) RETURN n
	st := "MATCH (n:" + label + " { " + key + ": '" + value + "' }) RETURN n"
	println(st)

	cq0 := neoism.CypherQuery{
		// Use backticks for long statements - Cypher is whitespace indifferent
		Statement:  st,
		Parameters: neoism.Props{},
		Result:     &res,
	}

	db.Cypher(&cq0)

	// println("largo")
	// println(len(res))
	// tagName, err := n1.Property("name")
	// println(tagName)
	// println(err)

	var node *neoism.Node

	if len(res) > 0 { // existe el tag
		println("si")
		node = &res[0].N
		node.Db = db
	} else {
		println("no")
		node, _ = db.CreateNode(props)
		node.AddLabel(label)
	}

	return node
}

func main() {

	// tags := ["juegos", "wii"]
	// tags := ["politica", "bachelet"]
	// tags := ["football", "deporte"]
	tagsMatrix := [3][5]string{
		{"juegos", "wii", "nintendo", "entretenimiento", "japon"},
		{"politica", "bachelet", "gobierno", "nueva mayoria", "michelle"},
		{"football", "deporte", "correr", "gol", "delantero"},
	}

	// {name,url,summary,tagSet, providerName}

	newsMatrix := [7][6]string{
		{"Ley de Presupuesto: La visión que han tenido los ministros de Hacienda frente al gasto fiscal", "http://www.emol.com/noticias/Economia/2015/09/30/752300/Ley-de-Presupuesto-La-vision-de-los-ministros-de-Hacienda-frente-al-gasto-publico.html", "El ministro de hacienda al fin hizo frente al gasto publico de Chile. Esta situ...", "1", "emol", "https://dl.dropboxusercontent.com/u/35902651/arqui/moneda.jpeg"},
		{"Creador de Super Mario confirma un antiguo mito sobre el juego", "http://www.emol.com/noticias/Tecnologia/2015/09/11/749458/Creador-de-Super-Mario-confirma-un-antiguo-mito-sobre-el-juego.html", "Un antiguo mito sobre Mario Bros es confirmado gracias a su creador, que dij...", "0", "emol", "https://dl.dropboxusercontent.com/u/35902651/arqui/mario.jpeg"},
		{"Nintendo cambia su política y creará videojuegos para celulares", "http://www.emol.com/noticias/tecnologia/2015/03/17/708342/nintendo-cambia-de-politica-y-creara-juegos-para-celulares.html", "Al fin nintendo decidio hacer algo distinto, juegos para celulares. La medida se tomo el...", "0", "emol", "https://dl.dropboxusercontent.com/u/35902651/arqui/nintendo.png"},
		{"Super Mario Bros cumplió 30 años", "http://www.cnnchile.com/noticia/2015/09/13/super-mario-bros-cumplio-30-anos-", "Shigeru Miyamoto creó hace 30 años, entonces sin imaginarlo, al personaje más reconocido de la historia de videojuegos. Tan rele...", "0", "cnn", ""},
		{"Presidenta Bachelet aseguró que solidez de la economía chilena es valorada en el extranjero", "http://www.cnnchile.com/noticia/2015/09/30/presidenta-bachelet-aseguro-que-solidez-de-la-economia-chilena-es-valorada-en-el-extranjero", "La mandataria aseguró que la agenda del Gobierno busca recuperar el dinamismo en el área.", "1", "cnn", "https://dl.dropboxusercontent.com/u/35902651/arqui/bachelet.jpeg"},
		{" Bachelet aseguró que solidez de la economía chilena es valorada en SQP", "http://www.lun.com/noticia/2015/09/30/presidenta-bachelet-aseguro-que-solidez-de-la-economia-chilena-es-valorada-en-el-extranjero", "SQP es el mejor programa de politica en chile, es por esto...", "1", "lun", "https://dl.dropboxusercontent.com/u/35902651/arqui/sqp.jpeg"},
		{"Alguien metio un gol", "http://www.lun.cl/alguien-metio-un-gol", "Un gol fue hecho el dia de ayer. Esto es muy interesante y misterioso, ya que ayer...", "2", "lun", "https://dl.dropboxusercontent.com/u/35902651/arqui/gol.jpeg"},
	}

	//Localhost
	// db, err := neoism.Connect("http://neo4j:12345678@localhost:7474/db/data")

	//Server
	db, err := neoism.Connect("http://neo4j:7c38caaee73a5564a3183c0970118725189ef64e9a565c982edb10e4388f43df@arqui7.ing.puc.cl:80/db/data")

	println(err)
	println(db.HrefNodeIndex)

	// db.CreateNode(neoism.Props{"l": "CHAO"})

	var tags [5]string

	for i := 0; i < 50; i++ {

		for _, row := range newsMatrix {

			println("row")

			// element is the element from someSlice for where we are
			title := row[0] + strconv.Itoa(i)
			url := row[1] + strconv.Itoa(i)
			summary := row[2] + strconv.Itoa(i)
			tag_id, _ := strconv.Atoi(row[3])
			tags = tagsMatrix[tag_id]
			provider := row[4]
			image := row[5]

			newItem := getOrCreateNode(db, "NewsItem", "url", url, neoism.Props{"url": url, "title": title, "summary": summary, "image": image})

			providerNode := getOrCreateNode(db, "NewsProvider", "name", provider, neoism.Props{"name": provider})

			providerNode.Relate("posted", newItem.Id(), neoism.Props{})

			for _, tag := range tags {
				num := i / 2
				tagNode := getOrCreateNode(db, "Tag", "name", tag+strconv.Itoa(num), neoism.Props{"name": tag + strconv.Itoa(num)})

				tagNode.Relate("is in ", newItem.Id(), neoism.Props{})

			}

		}
	}

	fmt.Println("Hello world!")
	cwd, _ := os.Getwd()
	fmt.Println("cwd:", cwd)
	fmt.Println("args:", os.Args[1:])

	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
	// db.CreateNode(neoism.Props{"name": "Captain Kirk"})
}
