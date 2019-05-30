package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

func isEqual(actual, expected *PortMap) bool {

	if len((*actual)["tcp"]) != len((*expected)["tcp"]) {
		return false
	}

	if len((*actual)["udp"]) != len((*expected)["udp"]) {
		return false
	}

	for ix, x := range (*actual)["tcp"] {
		if (*expected)["tcp"][ix] != x {
			return false
		}
	}

	for ix, x := range (*actual)["udp"] {
		if (*expected)["udp"][ix] != x {
			return false
		}
	}

	return true
}

func Test(t *testing.T) {

	var mConfig MasterConfig
	mb, err := ioutil.ReadFile("testdata/test.json")
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(mb, &mConfig); err != nil {
		panic(err)
	}

	host := "10.10.10.1"

	t.Run("Server config", func(t *testing.T) {
		expected := &PortMap{
			"tcp": []int{5000},
			"udp": []int{6000, 6001},
		}

		actual := spitServerConfig(&mConfig, host)
		if !isEqual(actual, expected) {
			t.Errorf("%+v and %+v are not equal", actual, expected)
		}
	})

	t.Run("Client config", func(t *testing.T) {
		expected := map[string]PortMap{
			"8.8.8.8": {
				"tcp": []int{},
				"udp": []int{},
			},
			"13.10.12.15": {
				"tcp": []int{},
				"udp": []int{},
			},
			"10.10.10.2": {
				"tcp": []int{5000},
				"udp": []int{6000, 6001},
			},
		}

		actual := spitClientConfig(&mConfig, host)

		for i, a := range actual {
			e := expected[i]
			if !isEqual(&a, &e) {
				t.Errorf("%+v and %+v are not equal", actual, expected)
			}
		}
	})
}
