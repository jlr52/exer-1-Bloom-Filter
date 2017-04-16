package main

import (
    "fmt"
    "github.com/jlr52/exer-2-bloom-filter/bloom_filter"
)


func main() {


    config := bloom_filter.NewBloomFilterConfig(100000, 4)
	filter := bloom_filter.NewBloomFilter(config)

	fmt.Println(filter.IsMember("test"))
	filter.Insert("test")
	fmt.Println(filter.IsMember("test"))
	fmt.Println(filter.IsMember("test2"))
	filter.Insert("test2")
	fmt.Println(filter.IsMember("test2"))
	fmt.Println(filter.IsMember("test3"))
	filter.Insert("test3")
	fmt.Println(filter.IsMember("test3"))
	fmt.Println(filter.IsMember("test4"))
	filter.Insert("test4")
	fmt.Println(filter.IsMember("test4"))
}