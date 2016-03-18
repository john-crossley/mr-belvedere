#!/usr/bin/ruby

require "json"
require "open-uri"

request_uri = "https://mr-belvedere.herokuapp.com/api/v1/jobs?api_token=EMAIL ME FOR API KEY"

buffer = open(request_uri).read
results = JSON.parse(buffer, symbolize_names: true)

results.each do |result|
  status = result[:success] == true ? "stable :ok_hand:" : "failing :sob:"
  working = "#{result[:name]} is #{status} "
  puts working += "| color=#{result[:success] ? 'green' : 'red'}"
end
