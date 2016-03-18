#!/usr/bin/env /usr/local/bin/node

'use strict';

const got = require('/usr/local/lib/node_modules/got');

const API_TOKEN = 'PUT_TOKEN_HERE';
const belvEndpoint = `https://mr-belvedere.herokuapp.com/api/v1/jobs?api_token=${API_TOKEN}`;

got(belvEndpoint)
  .then(response => {
    const data = JSON.parse(response.body);

    data.map(item => {
      if (item.success) {
        console.log(`${item.name} is stable :thumbsup:`)
      } else {
        console.log(`${item.name} is failing :thumbsdown: |color=red`)
      }
    })

  })
  .catch(() => {
    console.log('Error parsing response');
  });
