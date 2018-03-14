package WifiCassandra

import (
	"KLTN/data-stream-backend/Cassandra"
	"KLTN/data-stream-backend/Data"
	"KLTN/data-stream-backend/Stream"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	getstream "github.com/GetStream/stream-go"
	"github.com/gocql/gocql"
)

func InsertWifiTable(wifiData Data.Data) {
	gocqlUUID := gocql.TimeUUID()
	// theGeom := wifiData.TheGeom
	err := Cassandra.Session.Query(`
		INSERT INTO wifi_data (id, activated, bbl, bin, boro,
		borocd, borocode, boroct2010, boroname, city, coundist, ct2010,
		doitt_id, lat, location, location_t, lon, name, ntacode, ntaname,
		objectid, postcode, provider, remarks, sourceid, ssid, the_geom,
		type, x, y)
		VALUES (?, ? ,?, ?, ?,
			?, ?, ?, ?, ?, ?, ?,
			?, ?, ?, ?, ? ,? ,?, ?,
			?, ?, ?, ?, ?, ?, ?,
			?, ?, ?)`,
		gocqlUUID, wifiData.Activated, wifiData.Bbl, wifiData.Bin, wifiData.Boro,
		wifiData.Borocd, wifiData.Borocode, wifiData.Boroct2010, wifiData.Boroname,
		wifiData.City, wifiData.Coundist, wifiData.Ct2010, wifiData.DoittID, wifiData.Lat,
		wifiData.Location, wifiData.LocationT, wifiData.Lon, wifiData.Name, wifiData.Ntacode,
		wifiData.Ntaname, wifiData.Objectid, wifiData.Postcode, wifiData.Provider, wifiData.Remarks,
		wifiData.Sourceid, wifiData.Ssid, wifiData.TheGeom, wifiData.Type, wifiData.X,
		wifiData.Y).Exec()
	if err != nil {
		panic(err)
	} else {
		insertWifiDataToStream(wifiData, gocqlUUID)
		log.Println("Insert successlly")
	}
}

func insertWifiDataToStream(wifiData Data.Data, uuId gocql.UUID) {
	globalMessages, err := Stream.Client.FlatFeed("wifiData", "global")
	if err == nil {
		log.Println("Insert to getStream.io")
		activity, err := globalMessages.AddActivity(&getstream.Activity{
			Actor:  fmt.Sprintf(uuId.String()),
			Verb:   "post",
			Object: fmt.Sprintf("object:%s", uuId.String()),
			MetaData: map[string]string{
				// add as many custom keys/values here as you like
				"activated":  fmt.Sprintf(wifiData.Activated.String()),
				"bbl":        wifiData.Bbl,
				"bin":        wifiData.Bin,
				"boro":       wifiData.Boro,
				"borocd":     wifiData.Borocd,
				"borocode":   wifiData.Borocode,
				"boroct2010": wifiData.Boroct2010,
				"boroname":   wifiData.Boroname,
				"city":       wifiData.City,
				"coundist":   wifiData.Coundist,
				"ct2010":     wifiData.Ct2010,
				"doitt_id":   wifiData.DoittID,
				"lat":        wifiData.Lat,
				"location":   wifiData.Location,
				"location_t": wifiData.LocationT,
				"lon":        wifiData.Lon,
				"name":       wifiData.Name,
				"ntacode":    wifiData.Ntacode,
				"ntaname":    wifiData.Ntaname,
				"objectid":   wifiData.Objectid,
				"postcode":   wifiData.Postcode,
				"provider":   wifiData.Provider,
				"remarks":    wifiData.Remarks,
				"sourceid":   wifiData.Sourceid,
				"ssid":       wifiData.Ssid,
				"the_geom":   fmt.Sprintf("object:%s", wifiData.TheGeom),
				"type":       wifiData.Type,
				"x":          wifiData.X,
				"y":          wifiData.Y,
			},
		})
		log.Println("activity", activity)
		if err != nil {
			log.Println(err)
		}

	}
}

func GetAllWifi(w http.ResponseWriter, r *http.Request) {
	var dataList []Data.Data
	// m := map[string]interface{}{}

	globalMessages, err := Stream.Client.FlatFeed("wifiData", "global")
	// fetch from Stream
	if err == nil {
		activities, err := globalMessages.Activities(&getstream.GetFlatFeedInput{
			Offset: 0,
		})

		if err == nil {
			fmt.Println("Fetching activities from Stream")
			for _, activity := range activities.Activities {
				fmt.Println(activity)
				// id, _ := gocql.ParseUUID(activity.Actor)
				activated, _ := time.Parse(time.RFC3339, activity.MetaData["activated"])
				dataList = append(dataList, Data.Data{
					Activated:  activated,
					Bbl:        activity.MetaData["bbl"],
					Bin:        activity.MetaData["bin"],
					Boro:       activity.MetaData["boro"],
					Borocd:     activity.MetaData["borocd"],
					Borocode:   activity.MetaData["borocode"],
					Boroct2010: activity.MetaData["boroct2010"],
					Boroname:   activity.MetaData["boroname"],
					City:       activity.MetaData["city"],
					Coundist:   activity.MetaData["coundist"],
					Ct2010:     activity.MetaData["ct2010"],
					DoittID:    activity.MetaData["doitt_id"],
					Lat:        activity.MetaData["lat"],
					Location:   activity.MetaData["location"],
					LocationT:  activity.MetaData["location_t"],
					Lon:        activity.MetaData["lon"],
					Name:       activity.MetaData["name"],
					Ntacode:    activity.MetaData["ntacode"],
					Ntaname:    activity.MetaData["ntaname"],
					Objectid:   activity.MetaData["objectid"],
					Postcode:   activity.MetaData["postcode"],
					Provider:   activity.MetaData["provider"],
					Remarks:    activity.MetaData["remarks"],
					Sourceid:   activity.MetaData["sourceid"],
					Ssid:       activity.MetaData["ssid"],
					Type:       activity.MetaData["type"],
					X:          activity.MetaData["x"],
					Y:          activity.MetaData["y"],
				})
			}
		}
	}
	json.NewEncoder(w).Encode(dataList)
}

func PostWifiData(w http.ResponseWriter, r *http.Request) {
	url := "https://data.cityofnewyork.us/api/id/varh-9tsp.json"
	dataList, err := Data.GetDataByUrl(url)
	if err != nil {
		panic(err)
	} else {
		log.Println("Connect cityofnewyork successlly")
	}
	for index := range dataList {
		InsertWifiTable(dataList[index])
	}
	json.NewEncoder(w).Encode(dataList)
}

func DeleteWifiData(w http.ResponseWriter, r *http.Request) {
	globalMessages, err := Stream.Client.FlatFeed("wifiData", "global")
	if err == nil {
		activities, err := globalMessages.Activities(&getstream.GetFlatFeedInput{
			Offset: 0,
		})
		if err == nil {
			fmt.Println("Fetching activities from Stream")
			for _, activity := range activities.Activities {
				fmt.Println(activity.Actor)
				errDeleteDB := deleteWifiDataToCassandraDb(activity.Actor)
				if errDeleteDB == nil {
					fmt.Println("Delete to Streamio", activity.Actor)
					globalMessages.RemoveActivity(activity)
				}
			}
		}
	}
	json.NewEncoder(w).Encode("200")
}

func deleteWifiDataToCassandraDb(uuId string) error {
	err := Cassandra.Session.Query(`DELETE FROM wifi_data WHERE id = ?`,
		uuId).Exec()
	if err != nil {
		panic(err)
		return err
	}
	log.Println("Delete successlly", uuId)
	return nil
}
