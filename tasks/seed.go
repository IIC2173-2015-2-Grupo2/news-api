//usr/bin/env go run $0 "$@"; exit
package main

// for run seed, put in yout shell:
// chmod 744 script.go
// ./script.go 

import (
  "fmt"
  "os"
  "github.com/jmcvetta/neoism"
  "strconv"
  // "encoding/json"
)

func main() {

  




  // tags := ["juegos", "wii"]
  // tags := ["politica", "bachelet"]
  // tags := ["football", "deporte"]

  tagsMatrix := [3][2] string{
    {"juegos", "wii"},
    {"politica", "bachelet"},
    {"football", "deporte"},
  }

  // {name,url,summary,tagSet, providerName}

  newsMatrix:=[7][5] string{
    {"Ley de Presupuesto: La visión que han tenido los ministros de Hacienda frente al gasto fiscal","http://www.emol.com/noticias/Economia/2015/09/30/752300/Ley-de-Presupuesto-La-vision-de-los-ministros-de-Hacienda-frente-al-gasto-publico.html","El ministro de hacienda al fin hizo frente al gasto publico de Chile. Esta situ...","1","emol"},
    {"Creador de Super Mario confirma un antiguo mito sobre el juego","http://www.emol.com/noticias/Tecnologia/2015/09/11/749458/Creador-de-Super-Mario-confirma-un-antiguo-mito-sobre-el-juego.html","Un antiguo mito sobre Mario Bros es confirmado gracias a su creador, que dij...","0","emol"},
    {"Nintendo cambia su política y creará videojuegos para celulares","http://www.emol.com/noticias/tecnologia/2015/03/17/708342/nintendo-cambia-de-politica-y-creara-juegos-para-celulares.html","Al fin nintendo decidio hacer algo distinto, juegos para celulares. La medida se tomo el...","0","emol"},
    {"Super Mario Bros cumplió 30 años","http://www.cnnchile.com/noticia/2015/09/13/super-mario-bros-cumplio-30-anos-","Shigeru Miyamoto creó hace 30 años, entonces sin imaginarlo, al personaje más reconocido de la historia de videojuegos. Tan rele...","0","cnn"},
    {"Presidenta Bachelet aseguró que solidez de la economía chilena es valorada en el extranjero","http://www.cnnchile.com/noticia/2015/09/30/presidenta-bachelet-aseguro-que-solidez-de-la-economia-chilena-es-valorada-en-el-extranjero","La mandataria aseguró que la agenda del Gobierno busca recuperar el dinamismo en el área.","1","cnn"},
    {" Bachelet aseguró que solidez de la economía chilena es valorada en SQP","http://www.cnnchile.com/noticia/2015/09/30/presidenta-bachelet-aseguro-que-solidez-de-la-economia-chilena-es-valorada-en-el-extranjero","SQP es el mejor programa de politica en chile, es por esto...","1","lun"},
    {"Alguien metio un gol","http://www.lun.cl/alguien-metio-un-gol","Un gol fue hecho el dia de ayer. Esto es muy interesante y misterioso, ya que ayer...","2","lun"},

  }



  db, err := neoism.Connect("http://neo4j:12345678@localhost:7474/db/data")
    println(err)
    println(db.HrefNodeIndex)

  db.CreateNode(neoism.Props{"l": "CHAO"})


 

  var tags [2]string
  var node *neoism.Node
  for _, row := range newsMatrix {

    // element is the element from someSlice for where we are
    title := row[0]
    url := row[1]
    summary := row[2]
    tag_id, _ := strconv.Atoi(row[3])
    tags = tagsMatrix[tag_id]
    provider := row[4]


    node_, _:= db.CreateNode(neoism.Props{"url": url,"title": title, "summary": summary})
    node = node_
    node.AddLabel("NewsItem")

    providerNode, _ := db.CreateNode(neoism.Props{"name": provider}) 
    providerNode.AddLabel("NewsProvider")


    providerNode.Relate("posted", node.Id(), neoism.Props{})
  }

  for _, tag := range tags {

    // res := []struct {
    //     // `json:` tags matches column names in query
    //     A   string `json:"a.name"` 
    // }{}

    //  _ = neoism.CypherQuery{
    //     // Use backticks for long statements - Cypher is whitespace indifferent
    //     Statement: `
    //         MATCH (a:Tag { name: {tag_} }) RETURN a.name
    //     `,
    //     Parameters: neoism.Props{"tag_": tag},
    //     Result:     &res,
    // }

    // out, _ := json.Marshal(res[0])
    // println("largo")
    // println(len(res))

    tagNode, _ := db.CreateNode(neoism.Props{"name": tag})
    tagNode.AddLabel("Tag")

    tagNode.Relate("is in ",node.Id(), neoism.Props{})

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