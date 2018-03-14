import { Component, OnInit } from '@angular/core';
import { Http, Response, Headers, RequestOptions } from '@angular/http';
import {Observable} from 'rxjs';
import { element } from 'protractor';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
 
  title = 'app';
  baseUrl = 'http://localhost:8080/get-all-wifi';

  wifiDataList: any=[];

  constructor(
    private http: Http
  ){

  }

  ngOnInit() {
    this.http.get(this.baseUrl).subscribe(data => {
      let list =[];
      list = JSON.parse(data['_body']);
      this.wifiDataList = list.map(wifiData => {
        wifiData.lat = Number(wifiData.lat);
        wifiData.lon = Number(wifiData.lon);
        return wifiData;
      })
    });
  }

  isNumber(number){
    if(number instanceof Number)
      return true;
    return false;
  }
  

}

