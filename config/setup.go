package config

import (
	"github.com/RedisLabs/redisearch-go/redisearch"
	"log"
	"fmt"
)

func Setup(c *redisearch.Client){

	indexOptions := redisearch.Options{
		NoSave:          false,
		NoFieldFlags:    false,
		NoFrequencies:   true,
		NoOffsetVectors: true,
//		Stopwords:       []string{"0"},
	}
	fmt.Println("setting up index")
	sc := redisearch.NewSchema(indexOptions).
		AddField(redisearch.NewNumericField("hb"))
/*.
		AddField(redisearch.NewTextFieldOptions("ha",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("geo",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("uid",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewNumericField("da")).
		AddField(redisearch.NewTextFieldOptions("co",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("di",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("co",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("ge",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("lc",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("mo",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewNumericField("hb")).
		AddField(redisearch.NewNumericField("bd")).
		AddField(redisearch.NewTextFieldOptions("li",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.Field{ Name:    "l", Type:    redisearch.GeoField, Options: nil, }).
		AddField(redisearch.NewTextFieldOptions("fgen",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewNumericField("fmin")).
		AddField(redisearch.NewNumericField("fmax")).
		AddField(redisearch.NewNumericField("lld")).
		AddField(redisearch.NewNumericField("lmd")).
		AddField(redisearch.NewNumericField("lpd")).
		AddField(redisearch.NewNumericField("lfd")).
		AddField(redisearch.NewNumericField("lmd")).
		AddField(redisearch.NewNumericField("lpd")).
		AddField(redisearch.NewNumericField("lfd")).
		AddField(redisearch.NewNumericField("lbd")).
		AddField(redisearch.NewNumericField("hq")).
		AddField(redisearch.NewTextFieldOptions("ed",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("hq",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("jo",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("or",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("ws",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewTextFieldOptions("rs",redisearch.TextFieldOptions{NoStem: true})).
		AddField(redisearch.NewNumericField("pme"))
*/
	if err := c.CreateIndex(sc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("done setting up index")

}