#!/usr/bin/env /usr/local/bin/node

'use strict';

const request = require('https').get;

const API_TOKEN = 'PUT_TOKEN_HERE';
const belvEndpoint = `https://mr-belvedere.herokuapp.com/api/v1/jobs?api_token=${API_TOKEN}`;

request(belvEndpoint, res => {
  res.setEncoding('utf8');
  res.on('data', data => {

    const response = JSON.parse(data);
    response.map(item => {
      if (item.status == "success") {
        console.log(`${item.name} is stable :thumbsup: | color=green`)
      } else if (item.status == "failure") {
        console.log(`${item.name} is failing :thumbsdown: | color=red`)
      } else if (item.status == "building") {
        console.log(`${item.name} is building :construction_worker: | color=orange`)
      }
    })

  })
}).on('error', e => {
  console.log('Error fetching builds')
});
