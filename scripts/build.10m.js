#!/usr/bin/env /usr/local/bin/node

'use strict';

const got = require('/usr/local/lib/node_modules/got');

const API_TOKEN = 'PUT_TOKEN_HERE';
const belvEndpoint = `https://mr-belvedere.herokuapp.com/api/v1/jobs?api_token=${API_TOKEN}`;

got(belvEndpoint)
  .then(response => {
    const data = JSON.parse(response.body);

    data.map(item => {
      if (item.status == "success") {
        console.log(`${item.name} is stable :thumbsup:`)
      } else if (item.status == "failure") {
        console.log(`${item.name} is failing :thumbsdown: |color=red`)
      } else if (item.stauts == "building") {
        console.log(`${item.name} is building :construction_worker: |color=orange`)
      }
    })

  })
  .catch(() => {
    console.log('Error parsing response');
  });
