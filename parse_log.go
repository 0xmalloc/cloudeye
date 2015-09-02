package main

import (
	"encoding/json"
	//"errors"
	"fmt"
	"github.com/ActiveState/tail"
	fb "github.com/huandu/facebook"
	"github.com/influxdb/influxdb/client"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	//"strconv"
	"time"
)

type ConfigStruct struct {
	Metricname string   `json:"metricname"`
	Value      string   `json:"value"`
	C_type     string   `json:"type"`
	Time       string   `json:"time"`
	Tags       []string `json:"tags"`
}

type InfluxdbConf struct {
	Host     string `json:host`
	Port     int    `json:port`
	Database string `json:database`
	User     string `json:user`
	Pwd      string `json:pwd`
}

type ConfigArr struct {
	Metrics          []ConfigStruct `json:"metrics"`
	Filepath         string         `json:"filepath"`
	Backend_influxdb InfluxdbConf   `json:"backend_influxdb"`
	con              *client.Client //global conn
}

var config ConfigArr

func write_influxdb(con *client.Client, conf InfluxdbConf, pts []client.Point) {

	bps := client.BatchPoints{
		Points:          pts,
		Database:        conf.Database,
		RetentionPolicy: "default",
	}
	_, err := con.Write(bps)
	if err != nil {
		log.Fatal(err)
	}
}

func init_fluxdb(conf InfluxdbConf) (pcon *client.Client, perr error) {
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", conf.Host, conf.Port))
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	influxconf := client.Config{
		URL:      *u,
		Username: os.Getenv(conf.User),
		Password: os.Getenv(conf.Pwd),
	}

	con, err := client.NewClient(influxconf)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	dur, ver, err := con.Ping()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Printf("Happy as a Hippo! %v, %s", dur, ver)
	return con, nil
}

func readconf(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ReadFile: ", err.Error())
		return err
	}
	if err := json.Unmarshal([]byte(bytes), &config); err != nil {
		fmt.Println("Unmarshal: ", err.Error())
		return err
	}
	return nil
}

func processmetric(c ConfigStruct, log fb.Result) error {
	//fmt.Println("metricname:value|type|@smaple_rate|#tag1:value1,tagO2:value2")
	s := fmt.Sprintf("%s:%f|%s|@smaple_rate|", c.Metricname, log[c.Value], c.C_type)
	influxtags := make(map[string]string)

	for _, val := range c.Tags {
		influxtags[val] = log[val].(string)
		/*if index == 0 {
			t := string("#") + string(val) + string(":") + log[val].(string)
			s += t
		} else {
			t := string(",") + string(val) + string(":") + log[val].(string)
			s += t
		}*/
	}
	fmt.Println(s)
	/*logtime, err := strconv.ParseInt(log[c.Time].(string), 10, 0)
	if err != nil {
		return errors.New("strconv.ParseInt(log[c.Time] failed")
	}*/
	p := client.Point{
		Measurement: c.Metricname,
		Tags:        influxtags,
		Fields: map[string]interface{}{
			"value": log[c.Value],
		},
		//Time:      time.Unix(int64(log[c.Time].(float64)), 0),
		Time:      time.Now(),
		Precision: "s",
	}
	var pts = make([]client.Point, 1)
	pts[0] = p
	fmt.Println(pts)
	write_influxdb(config.con, config.Backend_influxdb, pts)
	return nil
}

func processlog(s string) error {

	var r fb.Result
	json.Unmarshal([]byte(s), &r)
	//fmt.Println(r)
	for _, val := range config.Metrics {
		//fmt.Println(index)
		processmetric(val, r)
	}
	return nil
}

func main() {
	log.Println("Started!")
	readconf("./conf/parse.conf")
	fmt.Println(config)
	con, err := init_fluxdb(config.Backend_influxdb)
	if err != nil {
		fmt.Println("init_fluxdb error")
	}
	config.con = con
	t, _ := tail.TailFile(config.Filepath, tail.Config{Follow: true})
	for line := range t.Lines {
		processlog(line.Text)
	}
}
