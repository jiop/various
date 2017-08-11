# Golang, gRPC et radix tree

Premier billet, premier example de ce qu'on peut faire lorsqu'on s'ennuie et qu'on veut continuer a apprendre Go (autrement appele Golang). L'objectif de ce "tutoriel" sera donc de developper une application en mode console qui renvoie le code postal de la ville passee en argument.

Je veux donc un programme qui fonctionnerait de la maniere suivante :

```sh
./postalcode -city avignon
> 84140
```

Pour compliquer un peu la chose, j'ai envie d'utiliser gRPC et un radix tree. Pourquoi? Parce que. J'ai dit que je m'ennuyais.

Du coup, on va avoir une architecture client-serveur utilisant gRPC comme moyen de communication. Le serveur sera donc responsable de stocker une liste de villes / codes postaux dans un radix tree, de recevoir des requetes gRPC contenant le nom d'une ville, de chercher cette ville dans le radix tree et finalement de repondre le code postal correspondant a la ville chercher.

## Lecture du fichier CSV

Premiere etape, il faut trouver des donnees exploitables. Le site web de la poste fourni un fichier csv contenant les informations dont nous avons besoin [ici](https://datanova.legroupe.laposte.fr/explore/dataset/laposte_hexasmal/download/?format=csv). Nous ne nous interesserons qu'a la deuxieme et la troisieme colonne de ce fichier csv.

```golang
package main

import (
	"encoding/csv"
	"io"
	"os"
)

const csvFile = "csv/laposte_hexasmal.csv"

func main() {
  // Dans un premier temps, nous allons stocker les donnees dans une simple map.
  cpMap := make(map[string]string)

  // Ouvrons le fichier et assurons nous que nous allons bien le fermer a la fin du programme.
  f, err := os.Open(csvFile)
  defer f.Close()
  if err != nil {
    log.Fatal(err)
  }

  reader := csv.NewReader(f)
  // Le fichier csv de la poste utilise le caractere ; comme delimiteur.
  // Notez qu'ici il s'agit d'une rune et non pas d'une string
  reader.Comma = ';'

  // Faisons une premiere lecture pour eviter de stocker les entetes du fichier csv.
  if _, err = reader.Read(); err != nil {
    log.Fatal(err)
  }

  for {
    // Lisons les lignes une par une
    record, err := reader.Read()
    if err == io.EOF {
      // Si nous atteignons la fin du fichier, sortir de la boucle
      break
    }
    // Stockons les donnees dans la map
    cpMap[record[1]] = record[2]
  }
}
```

Dans cette premiere version, nous nous contentons de lire le fichier csv et de le stocker dans une map.

## Utilisation d'un radix tree

Une implementation de la structure de donnee est disponible [ici](https://github.com/armon/go-radix), merci aux gars de Hashicorp et a Armon Dadgar en particulier.


```golang
package main

import (
	"encoding/csv"
	"io"
	"os"
  radix "github.com/armon/go-radix"
)

const csvFile = "csv/laposte_hexasmal.csv"

func main() {
  // Initialisons le radix tree
  cpRadix := radix.New()

  // Ouvrons le fichier et assurons nous que nous allons bien le fermer a la fin du programme.
  f, err := os.Open(csvFile)
  defer f.Close()
  if err != nil {
    log.Fatal(err)
  }

  reader := csv.NewReader(f)
  // Le fichier csv de la poste utilise le caractere ; comme delimiteur.
  // Notez qu'ici il s'agit d'une rune et non pas d'une string
  reader.Comma = ';'

  // Faisons une premiere lecture pour eviter de stocker les entetes du fichier csv.
  if _, err = reader.Read(); err != nil {
    log.Fatal(err)
  }

  for {
    // Lisons les lignes une par une
    record, err := reader.Read()
    if err == io.EOF {
      // Si nous atteignons la fin du fichier, sortir de la boucle
      break
    }
    // Transformons la clef en quelquechose de plus safe (pas d'espaces dans le nom)
    cityName := strings.ToLower(strings.Replace(record[1], " ", "_", -1))

    // Stockons les donnees dans le radix tree
    cpRadix.Insert(cityName, record[2])
  }
}
```