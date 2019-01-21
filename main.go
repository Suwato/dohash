package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
	"fmt"
	"encoding/csv"
	"io"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
)

const version = "0.1.0"

func main() {

	var file string
	var algorithm string
	var stretching string
	var salt string

	app := cli.NewApp()
	app.Name = "dohash"
	app.Version = version
	app.Usage = "指定されたCSVの1列目をハッシュ化します。"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "file, f",
			Usage:       "ハッシュ化したいcsvのpathを指定してください。",
			Destination: &file,
		},
		cli.StringFlag{
			Name:        "algorithm, a",
			Usage:       "ハッシュ化のアルゴリズムを指定してください。sha256 と sha512 に対応しています。",
			Value:       "sha256",
			Destination: &algorithm,
		},
		cli.StringFlag{
			Name:        "stretching",
			Usage:       "ストレッチングの回数を指定してください。",
			Value:       "10",
			Destination: &stretching,
		},
		cli.StringFlag{
			Name:        "salt",
			Usage:       "saltを指定してください。",
			Destination: &salt,
		},
	}

	app.Action = func(c *cli.Context) error {
		if file == "" {
			return fmt.Errorf("plese set file path")
		}

		readFr, err := os.Open(file)

		if err != nil {
			return err
		}

		defer readFr.Close()

		r := csv.NewReader(readFr)

		writeFr, err := os.OpenFile("hash.csv", os.O_WRONLY|os.O_CREATE, 0600)

		if err != nil {
			return err
		}

		defer writeFr.Close()

		w := csv.NewWriter(writeFr)

		for {
			r, err := r.Read()

			if err == io.EOF {
				break
			} else if err != nil {
				return err
			}

			hash, err := hash(r[0], algorithm, salt)

			if err != nil {
				fmt.Errorf("error:%s", err)
			}

			w.Write([]string{hash})

		}

		w.Flush()

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func hash(s string, algorithm string, salt string) (string, error) {
	if algorithm == "sha256" {
		sum256 := sha256.Sum256([]byte(s + salt))
		return hex.EncodeToString(sum256[:]), nil
	}

	if algorithm == "sha512" {
		sum512 := sha512.Sum512([]byte(s + salt))
		return hex.EncodeToString(sum512[:]), nil
	}

	return "", fmt.Errorf("unknown algorithm")

}
