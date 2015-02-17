package main
import (
  "log"
  "github.com/PuerkitoBio/goquery"
  "strconv"
  //"strings"
  //"fmt"
)

func getTempGismeteo() (int, error) {
    doc, err := goquery.NewDocument("http://www.gismeteo.ru/city/daily/5032/")
        if err != nil {
        log.Fatal(err)
    }
    //take first temperature in Celsius
    temperature := doc.Find(".temp dd").First().Text()
    //sign of number
    sign := temperature[0:3]
    //cut temperature
    temperature = temperature[3:len(temperature) - 3]
    //convert string to int
    int_temperature, err := strconv.Atoi(temperature)
    if err != nil {
        log.Fatal(err)
    }
    //if minus
    if sign == "−" {
        int_temperature = int_temperature * -1
        //log.Print("Value: ", int_temperature)
    }
    return int_temperature, err
}

func getAccuweather() (int, error) {
    doc, err := goquery.NewDocument("http://www.accuweather.com/ru/ru/saratov/295382/current-weather/295382")
        if err != nil {
        log.Fatal(err)
    }
    //take first temperature in Celsius
    temperature := doc.Find(".temp")
    temp := temperature.Last().Text()
    temp = temp[0:len(temp) - 2]
    int_temperature, err := strconv.Atoi(temp)
    if err != nil {
        log.Fatal(err)
    }
    //log.Print("Value: ",  int_temperature)
    return int_temperature, err
}

func main() {
  getTempGismeteo()
  getAccuweather()
}