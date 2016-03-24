#!/usr/bin/ruby

require "json"
require "open-uri"

request_uri = "https://mr-belvedere.herokuapp.com/api/v1/jobs?api_token=EMAIL ME FOR API KEY"

buffer = open(request_uri).read
results = JSON.parse(buffer, symbolize_names: true)

def parse_result(result)
  case result
  when "success"
    return "stable :ok_hand: | color=green"
  when "failure"
      return "failing :sob: | color=red"
  when "building"
    return "building :construction_worker: | color=orange"
  end
end

results.each do |result|
  status = parse_result result[:status]
  puts "#{result[:name]} is #{status} "
end


