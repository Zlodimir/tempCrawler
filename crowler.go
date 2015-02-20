package main
import (
    "log"
    "github.com/PuerkitoBio/goquery"
    "strconv"
    "time"
   _"github.com/lib/pq"
    "github.com/astaxie/beego/orm"
  //"database/sql"
  //"strings"
  //"fmt"
  //_ "github.com/bmizerany/pq"
)

type Gismeteo struct {
    Resource_name string
    Url string
    Temp int
    Date time.Time
}

type Accuweather struct {
    Resource_name string
    Url string
    Temp int
    Date time.Time
}

type Temperaturer interface {
    Temperature() (int, error)
}

func (a *Accuweather) Temperature() (int, error) {
    doc, err := goquery.NewDocument(a.Url)
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
    a.Temp = int_temperature
    a.Date = time.Now()
    return int_temperature, err
}

func (g *Gismeteo) Temperature() (int, error) {
    doc, err := goquery.NewDocument(g.Url)
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
    g.Temp = int_temperature
    g.Date = time.Now()
    return int_temperature, err
}

func (a Accuweather) String() (string) {
    return a.Date.Format(time.RFC822) + " - " + a.Resource_name + " " + strconv.Itoa(a.Temp) + " °C"
}

func (*a Accuweather) Save(orm) (int) {
    
}

func (g Gismeteo) String() (string) {
    return g.Date.Format(time.RFC822) + " - " + g.Resource_name + " " + strconv.Itoa(g.Temp) + " °C"
}

type Temperature struct {
    Id          int
    Date        time.Time
    Source      int
    Value       int
}

func init() {
    // register model
    orm.RegisterModel(new(Temperature))
    // set default database
    orm.RegisterDriver("postgres", orm.DR_Postgres)
    orm.RegisterDataBase("default", "postgres", "dbname=tempdb user=postgres password=123456")
}

func main() {
    orm.Debug = true
    o := orm.NewOrm()
    o.Using("default")

    // Database alias.
    name := "default"
    // Drop table and re-create.
    force :=  false
    // Print log.
    verbose := true

    // Error.
    err := orm.RunSyncdb(name, force, verbose)
    if err != nil {
        log.Println(err)
    }
    /*
    t := Temperature{Id: 2}

    err = o.Read(&t)
    
    if err == orm.ErrNoRows {
        log.Println("No result found.")
    } else if err == orm.ErrMissPK {
        log.Println("No primary key found.")
    } else {
        log.Println(t.Id, t.Value)
    }
    */
    a := Accuweather{Resource_name: "www.accuweather.com", Url: "http://www.accuweather.com/ru/ru/saratov/295382/current-weather/295382"}
    //a.Temperature()
    log.Print(a)

    g := Gismeteo{Resource_name: "www.gismeteo.ru", Url: "http://www.gismeteo.ru/city/daily/5032/"}
    //g.Temperature()
    log.Print(g)
    
    var temp Temperature
    temp.Source = 1
    temp.Date = g.Date
    temp.Value = g.Temp

    id, err := o.Insert(&temp)
    if err == nil {
        log.Println(id)
    } else {
        log.Print(id)
    }

    temp.Source = 2
    temp.Date = a.Date
    temp.Value = a.Temp

    id, err = o.Insert(&temp)
    if err == nil {
        log.Println(id)
    } else {
        log.Print(id)
    }
}