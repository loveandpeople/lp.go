package samples

import (
	. "github.com/loveandpeople/lp.go/trinary"
)

var SampleAddresses = Hashes{
	"PERXVBEYBJFPNEVPJNTCLWTDVOTEFWVGKVHTGKEOYRTZWYTPXGJJGZZZ9MQMHUNYDKDNUIBWINWB9JQLD",
	"VIWGTBNSFOZDBZYRUMSFGHUJYURQHNYQMYVWGQOBNONDZRFJG9VQTAHPBMTWEEMRYIMQFRAC9VYBOLJVD",
	"OIIGC9ZBJIGFAVEJZGDPKCHPPNPSATBEFYRZYCEHQVKNQLKBZUZJVTKUXCTLGNUSSRKGBZB9BVMMOKXDW",
}

var SampleAddressesWithChecksum = Hashes{
	"PERXVBEYBJFPNEVPJNTCLWTDVOTEFWVGKVHTGKEOYRTZWYTPXGJJGZZZ9MQMHUNYDKDNUIBWINWB9JQLDXHFWXDFWW",
	"VIWGTBNSFOZDBZYRUMSFGHUJYURQHNYQMYVWGQOBNONDZRFJG9VQTAHPBMTWEEMRYIMQFRAC9VYBOLJVDBPTIELAWD",
	"OIIGC9ZBJIGFAVEJZGDPKCHPPNPSATBEFYRZYCEHQVKNQLKBZUZJVTKUXCTLGNUSSRKGBZB9BVMMOKXDWHTJRZLHBC",
}

var SampleAddress = "PERXVBEYBJFPNEVPJNTCLWTDVOTEFWVGKVHTGKEOYRTZWYTPXGJJGZZZ9MQMHUNYDKDNUIBWINWB9JQLD"
var SampleAddressWithChecksum = "PERXVBEYBJFPNEVPJNTCLWTDVOTEFWVGKVHTGKEOYRTZWYTPXGJJGZZZ9MQMHUNYDKDNUIBWINWB9JQLDXHFWXDFWW"

var SampleAddressWithInvalidChecksum = "FJHSSHBZTAKQNDTIKJYCZBOZDGSZANCZSWCNWUOCZXFADNOQSYAHEJPXRLOVPNOQFQXXGEGVDGICLMOXXCHLTDDABS"
