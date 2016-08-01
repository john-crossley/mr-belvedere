#!/usr/bin/ruby

require "json"
require "open-uri"

API_TOKEN = "05561e32-77d6-4bb3-babe-3bb620837225"

request_uri = "https://mr-belvedere-go.herokuapp.com/api/v1/jobs?api_token=#{API_TOKEN}"

def get_color_for_status(status)
  case status
  when "success"
    "green"
  when "building"
    "orange"
  else
    "red"
  end
end

begin
  buffer = open(request_uri).read
rescue Exception => e
  puts ":poop:"
  exit -1
end

results = JSON.parse(buffer, symbolize_names: true)

results.each do |result|
  puts "#{result[:name]} | color=#{get_color_for_status(result[:status])} dropdown=false"
end

puts "---"

results.each do |job|
  puts "Job: #{job[:name]} | size=12"
  puts "├ Author: #{job[:author]} | size=12"
  puts "├ Version: #{job[:version]} | size=12"
  puts "├ Branch: #{job[:branch]} | size=12 href=#{job[:branch_url]}"
  puts "└ Status: #{job[:status]} | color=#{get_color_for_status(job[:status])} size=12 href=#{job[:build_url]}"
  puts "---"
end
