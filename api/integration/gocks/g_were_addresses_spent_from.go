package gocks

import (
	. "github.com/loveandpeople/lp.go/api"
	. "github.com/loveandpeople/lp.go/api/integration/samples"
	. "github.com/loveandpeople/lp.go/trinary"
	"gopkg.in/h2non/gock.v1"
)

func init() {
	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(WereAddressesSpentFromCommand{
			Command:   Command{WereAddressesSpentFromCmd},
			Addresses: SampleAddresses,
		}).
		Reply(200).
		JSON(WereAddressesSpentFromResponse{
			States: []bool{true, false, false},
		})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(WereAddressesSpentFromCommand{
			Command:   Command{WereAddressesSpentFromCmd},
			Addresses: Hashes{SampleAddresses[0]},
		}).
		Reply(200).
		JSON(WereAddressesSpentFromResponse{
			States: []bool{true},
		})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(WereAddressesSpentFromCommand{
			Command:   Command{WereAddressesSpentFromCmd},
			Addresses: Hashes{SampleAddresses[1]},
		}).
		Reply(200).
		JSON(WereAddressesSpentFromResponse{
			States: []bool{false},
		})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(WereAddressesSpentFromCommand{
			Command:   Command{WereAddressesSpentFromCmd},
			Addresses: Hashes{SampleAddresses[2]},
		}).
		Reply(200).
		JSON(WereAddressesSpentFromResponse{
			States: []bool{false},
		})

	gock.New(DefaultLocalIRIURI).
		Persist().
		Post("/").
		MatchType("json").
		JSON(WereAddressesSpentFromCommand{
			Command:   Command{WereAddressesSpentFromCmd},
			Addresses: SampleAddresses[1:],
		}).
		Reply(200).
		JSON(WereAddressesSpentFromResponse{
			States: []bool{false, false},
		})

}
