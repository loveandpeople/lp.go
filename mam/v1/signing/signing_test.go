package signing_test

import (
	. "github.com/loveandpeople/lp.go/consts"
	. "github.com/loveandpeople/lp.go/curl"
	. "github.com/loveandpeople/lp.go/mam/v1/signing"
	. "github.com/loveandpeople/lp.go/trinary"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Signing", func() {

	const seed = "ZLNM9UHJWKTTDEZOTH9CXDEIFUJQCIACDPJIXPOWBDW9LTBHC9AQRIXTIHYLIIURLZCXNSTGNIVC9ISVB"
	const hashSignature = "B9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9YB9SB9YB9YB9YB9YB9YB9YB9YB9YB9E"

	subseeds := []Trytes{
		"DRUVUYIMTSVUXGUCZA9BRWOGZIEJXE9LXAPOO9OVVWMCYJYJYSZALBHACPPIQBGRCNDWLLBXLOAYWLCRU",
		"FGQTVMPSJWOWYSPIJLIXNJOPIOTRMDGCZYYSJXPMMTSKURBSJ9TWODQPQD9USBQV9TQZTTFCUKQOBSXMQ",
		"GWIUZGY9KS9MHYKXGYUJBXCOWVNLCDLIDLUQECFKFICMOFPDHIOROCA99IH9SUUS9HYYPUTIOGBBYOBOL",
		"RTBWWVOGAMDS9ZSHWJOOPBJEJMTCBVXXFTVPUJDDZRDWIFOGUJDFY9L9JNGUAGAXGLYVIYU9XAWDVEGDG",
		"ZNXXANFDTCMTK9JFCKDZXODRNGDCOKDSXQC9JJAFLMOLVHESSN9WELIDFKXVDANFTWQJPABYKBZE9BOXH",
	}

	keys := []Trytes{
		"HWJVLIDXSNPTMROGNJPBMF9BZBLGJECQZWWWTLAGLGAMLHRWFOXCTZDWMTQSYEVKWBESSWKEWV9AYZHFIRWZRXUZGZQN9KLDLYQCKCGRAXHNFIEJQWFFIHVXDWJIBYLA9BLKFFUQFRBWIXTHDLYSR9KJQWRZUQBFEOJPL9RDDNJCZHAWHCZXRLKYCWCUIODROQA9HDFZFJDJXR9IEILQSYLKXZ9UA9FATTDXOGQL9TOLIHOEAPXYCRKIJNQPTJDOQRKBWYIFIGBQDDMGEU99GSAEVXJ9UTWYTLXHDNRZICLBCIIEGUYEBPAHYY9RIFQUTRVNXSHVKANGTWCVTIWOO9BUPRLLDQ9ZYGKSJUNNTGAFPINPTJXYUTXCLXFGFGDZSOFQ9CTEV9YBXPKKJ9QJFOJHIWGXEMCURNVTJBMKMZ9EYTZ9O9MLZTCHIHFBFZJAT9PEAHGN9RKKCBEBHYLWNSLLHDEAPYFCVYQPXGNORCHXONSDXWVYKDZVIBQWYKSAAQNW9PUTRLGP9B9HEZGHDIHDCOKWDLEITQVM9S9BALPQRJIZFMFXT9FHGEHXJWYSKWYFXOYBIPJUCVJVEAUBULEKKGRIEODONQLU9CQLPDQBBKWIQMDK9ICOXDCYEQUUFMIFLZKUPMOL99LOCLDYEQXYGIKJMCCXN9OTJSSRVLUJNGLKYPJPCDPTYLRQOTTCVHTVGVPGCXCM9AXKPKGFPTX9KFOEWDMJAOSSJ9OKUJXADWZIRHJJQGJ9XZBBEWATROGUIFUEF9JVHU9NF9KQRHTTZIOQPYZEBDMVJZTDKKHAWFNRSWIPDJRIVGIJPUQSKEFBWIGLEFPJSSWQTDVCDOFPYDALFZHZJFKP9SVFERWAZUZSOELYWRDOWHFNXWVVFQTSHQWD9TVICIXQUCKUIPLOLSYJ9SJRMQNSJNPTEEKELNAY9WCTRQSHIHQFYDSV9FTHMUBRFJLCEP9HUBANVOSUVZNVUGAPUZXAOOFMWQWXWNIRMIALUZSSNSWKNQYUOUCPQ99XRTKX9XRVTGVZQVWDSLQPYLJNBZYYTRWAXLGIBHCAKJISJLICHZKZ99PAGHAJVGAOEVAZFOEHOMYEORCHBXSIKBFTANGBNDZTDMKGKUDAREITODOK9NITUDTRQNFWYRRAKGHXDCAHMVXWHCKZWBSAUHQDSENADYUYEGXUISMGAZOIAUPGVSYGUFAZLRASZBLXLRCCGU9YTXF9DFHRHWBE9YZHDOLLJI9OCGUWJDXWOOQZYEHDNGMAVGPJWPJDMQIWDVXRRKPMFJQACJADWXTRTNOSUNCUZD9NIRYCKJXOVQUJPHXPCJMUZAPDIPXQXJGEZWXGJSBSFEVX9REVNPEUMBDEFSXMFECLSD9AHZAXTNUHEHYRQMSBIBCMMSMBNCCNFYNJLCICPTMHATFZNNHBQPNNJJNVFLYUWZIBQUKISWCIVNKBRYUFHSZXNMXSMZGFCPQUMYUSJIKEVKLAQZLEN9VKIEVZUTPKCLJMWETXIUZJRWZGCACNUEOIQDVIZVDHKJMXJERUF9QVRCFOTXGMWMTKZRQCKY9TJD9KMAK9WUKRCWUKHSRQKIYHANCOSOUUGXUIEG9ZOXLVXBELUMBBSCVITQVCNYGOUIHXGZKQ9RMLUSESKQIWHPHEUHIZORWKWZJXPOEWMYAVJYGGHEUQOMIWVALCSNJPSCJUUGAJPNWNPR9NEFYXVNSMXBHRNWOJOKZ9KZCVPBCNVKZRBNVFGZMNBEDEQYSKWLARPJAIBWYGASNUSBCEPGKWVNSRDO9CBDTJXIBUVEPKGKYCRWSABH9MAATCWZYVFDCOUPONSSJPHNDDXVTXYJPXKPQAVHSURLGAYRQRCDAGXFNVYJPUQZS9QKYSQ9RIECWWBCEVCORSLUIEWUIOSOGKYWCQCDEKBEPZDVUBEBMYRSOTTICZGQDHYJFKAOHWKKUCRRGFLINY9XLMXGQZDGGVSXDBPENVCGMMVAVOZVDKOBAULIIYBXPAFKAVOKMQSYHZDFSHLFYSMXIBOZHF9OMMQUJTZGHGO9KOCRR9EUUVH9GFVAKFMBLCCZCPXDCMPRTIIFUZLEDDSLGGWBZQNMEEOCDEGJBBWSPITOHGISBDVMPHGCHNTC9LBKIURKNXXWEANRCF9HLGLXTCINCDDTAXUXOEHFVLEAX",
		"TAWZAGFJNXPWGYPUDNQXDSTSQAYTWYL9GWDAJ9RKXMKKLUBYZIDFQXGELRAZRDILGMOSNTOIQAG9IMSHMSNGRAADHJNBHXRHKQFPMZUYMVACVVRCWBSQPX9ZDOCVIMQLLUSCPPEOWQLLDEDVQFCMD9TCJFVRKYZJRZ9SXYWMRBMETPWKYUGXEJNHBTZSMLIBZXUMVNJSUHLXLLGTJI9XYPGZJKHPGQZQIORXFLMNGJZNOOKCZRD9XGLC9ICZPIEGLMDRKE9UBT9QLMBQUHLWTXJKYOWWSAJASAGOVJBCDTAIMRBATYJFDMCNMBMDS99DRNZCKKVKECVNRZKYDNXNIZ9LTMLFLIRNPYABTFJXVZOTMUSVPXUBQJZVKCDARRCVMDUJCQYGFUHDINGMPGUYQVLCTGLBMVGQYCXHZMCCYRHFQHKOGYBZVJNHTJGHYA9NZF9CSOUDOGYTDFDYZBKIFOMMYWNRSIBRZJGYPXIPPDGNGW9SQC9RBBR9UNVPYQDWSCCDNTDLA9ZLZGXOBZSXRFYXJWTVKKLWPORSOHQBIZVFAYZYLOUGLWYLRBIZAVWBPZAJXKIM9ZTQXUPNL99HPMZXRILQJ9WMGWDLSPXPYDNAWGZGKOCVUCXGCVQLVOZZYBTHKNQVEEBJE9HGAHKBBXFYKARKHWEJFHTUQAMHHCKNXKG9HANIAKDNXDEJOCODQGVIZRKHVDSSKQEKPCFEATXQREFH9HKXLCIFJNXYHC9DXGOEMLKGDMQAQYIPYUFXQGCANKXTFLGVSYNSFKQMQDLKWNVTD9MCRBMKRRWSYNGVKQDRGHWDUROOPMYKVDAVLCWSPZALUCVEIJOCEIQECYAXBJRFMCEDFTPHB9TJZRBSGTGDLHXVVTJZVNLT9HCXSOSJOND9MSEJ9APQSLQMQXSLIHQQDRJBCHNSKUMHNJGUHNFRAMQTNOXVLDSNXWZAPXNEVKYMSBYGRXVCYFODMKLQKJBVPPDNIIRDTCBXUEFJHPOWLLPKOICSTXM9LTHKCYJPHLLFUGCDSQNYVMXYYHJSWLKPSWGNERBMFNNNXIOUVMXZCDLOGPHXDCIDEZZLELWZEDDWTESLXMGJBGZS9KJTDTUMYXV9DRC9XXDATUTXTMNTDEB9HJXCJXRSLRFPOZQLEHL9DVCDEVVQALRFQQJBOYQUBHCRNTXVJKJNAWZINPXYOUCJJIFODPXYUZH9VQZQKKRK9QZJ9UVOGVYTCGUDLUC9VRWIXKKJXNNKPHJCJYIJGUEGGGYBCQSCHWHXUDIHJWCGYRPUPYKPFAGXS9AMF9PDGITCKHRNEHPMHYFGZCWVFJSTUACBKAWWURMITBKAGUWDEUIUTJNW9DKLWLVLGTIDPYEAVA9ISGQOJBFQAZKMJEQUOMKDHAHQJY9RPD9JOUZZ9KAZKKERQBGUBHKYPDV9LCFFLGRKBNENEVLSJCAJSECLMRQMZMSWQZN9KZLGIBNEQYIAWR9YPENYHXEZIHTXOIYZWAHFJDHMWHYWNPIUXZXSZVGPAAMILI9CKCQDSUHCSQCSDWUVVKUEJAIJWHBLSKSTVKSRFITZNAVPBZNJN99ZWUZBOIUVFYYHMLFBRD9EXGGGQ9HOVREWCRLRLGTFHIYGIADTHUERGP9BGH9KAGIBDVJ9SXPFGOFCGIUZUQDQKYQDBYOJZWNDVSHSZMUPPT9DBCTDZSXEWEVURFBBVBOZRDWTTEBLOXEXYTKWIP9BKNNMZKS9UTFUJJEMGINQ9QITQ99DOUZIRJWAPDTHKXCRHZBSHHQVZPSZXIDGVHBF9ARQVICLBQZVAWMSQUZTWLIUEXMZJKSVFIBEMEBPJDURNMJDHYAXURWRIRNVFJZORVGOTMGAATOFICLSPICQUMCHYFYGVFDXYELQEHPNDZIXFDMHHAZMBZYMZRTSHVTQYWYCYGJFHAIUWXZIIWXCQPZUSALYUOCBXTKFOAGRIOEGZUNHYESLFJEOW9XXFTFZWUTXUVFOHZYYQEWX9SZZLYVFVNJYHEQDILOGVXX9OEROTBECUEABSFWVWJIVZVOHMQSSOFTZYPHUKKWNJTSKFTTCVGXLMIHRWOBCCAFXCBIITXYSRGYYQYXAXLYAHEDJGYSLWIGUVPCQGN9E9NSMANNHVCTIVYMFALUEGWINFNLTTHJMVWUUGYVPIXDADQVJCMD",
		"GANISBRBXOLXDYUVAHLGEEJBTDESQQTWGULLZ9MCFLQLPWWCGQCAVVRUCAWJPRIWLNXBVTUFJXQGSNJMUHHIOBECNYDSOXZHRNMXXMWSBGFWLXHJVZRSU9XZWZVMQVZHKKOWNREFSANMSXUJHKIRUSOZJUQSXQDLCNNX9RNAFPOTOYEIORUTYTCUG9TINCTXBHMLVIASPICSLXCQYSQARWGTWRVOGWKIIFORUJRFRNMLFMUAOTETPYP9RARTKIGGCRHVGMIMAFFVEMITAXTXRVMGTDYAMWUOSDX9DN9FYV9KSRKHPYYGYQWDHCCFWPFIBXDJOVBBIXWTSYOMPBUMUS9FZVOLBBMKLHNFSQWZ9FJNWITKFUQOAJBRP9OQIFY9SGTRQZGWVAIYBOWKGWN9DGQVPNUQQPKVVQZIHAPGSGBKVL9VWZKKX9TDPQIKKWHP9OKGVDLNCTBYOONYDHWQKIF9BFYLBPUVKKJOHYQDCYPHSDOUSQNKSQCSL9XICJYFPZNEZTRBWUCFVCSKUDVPKZPSQUELRVROZVUDVAPWXKFCXMMCRKNNC9FBWWQBUUZSEVYXQRBJLBCQBUGMOIFFSZNFVKWO9WGJXLNCQWSH9L9VVLSRIXMTPVMKXBSFBAWSFAOHMVLSYRSKLTPNCOL9AZDVSGKPLBRBSYSGXZVDQMQQNDIIEWVWZAGYFOZCXJGSJLFIVHKF99SQOOYCUYAKIMCWSBZ9QFUTZLNJSIEYEPPJHDBGKLFCOKTNIXAZTGLODTVMYCOOICTNLOB9RXVVQOIXXLNXMVSBU9YDVNZZHCKVZXUEPGYFVTMNKJTJGFBBPMPFWIQ9AHCXMPN9PHU9FFGFGINFHLNOEIALYDPFHJQCJCCPDNXZQXLEWIMLZGRBXHGPJFFAHGNWXQGRKLZOMBDMOLRQQWAYYBNXSVCLE9ILFGZRV9YTEQP9LQIOHGPIOC9EFIJGQUGOHNJBKXHXOSLVOKZVYLWSFEPZC9JJLDBXCNWSBBWOASPOKSDLO99FILDUQLMGQSESKJGBYQPGZITGOWPATPVPFXBSGQQQDTZJPNDVTLOAHBMHGFLGZQHY9UQABUWNQALSGEPGXPDHTDK99WBANZURIJ99WVGNJVXZWDWISXK9ZYCSEEHIE9VEMDQGSTH9ORPWYRXDWJUONBAZRJ9CDIMVSHJBLCONLGSIEWYEUSTSZRJBIMCGKBIHL9XIMNYVGKCOTJOEGRFLVUASMXDSSLKKJAEOMXBWNYNDXQMEHBXSPUEPRMPWXYRHLIMCEEGTBSXUXXQOBBGQZEPWEBUNNCGUPVNJEPVPSNAAFY9RPVIDZNSXTCATVHGZE9NRUXXJVLNRVLKZFSGRVDXSPZFTIFKJMQVDBKDKHWEGIOX9UQKBZSKFIVJAI9FJNCMUI9ESLSGSBBJUYNOHVZAFZREQM9EHC9XXSOVXIS9SD9OMWWPPT9XNPHSVFAFUEMXNDURPYSYKQBCOPDPMJZOQGFAIBIICM9BPOEHDMBTBRGQSCKSUBFK9JHKRSWIVYQARHPIMTDQMDZIU9KMUFEH9RYBCVMYBYSARZAFQJSEAYDNZHWCLIRFHZEGBKVYWLAPMKSRJUVIQUUBTOQIGBZLALFNSSRTFHHVJUOVGBP9GJ9SYCDT9LN9AZG9HSKYMTMYWGCYSGPEEDHS9EYVTHKCPDRIOYRUHZIHQPJVARAJDIMVMZQFLBU9ETZUDEIKUYCJLTZLTQLSUAAGREGIFTOEJLBCUEQYBVQM9CIJGXWBKFCQ9XKBEDVMKSRTIUVDIVZNZUHHOKHNPDILXUQQOVWKDPVSMPQGVWIKOL9ASSBRUMDCEOFSF9ADAKJTXMRHWMFKOYFJUEVBCFXDTRPGKXSL9TWRKPEWVKKKPBJYDJQKPLFIKOOSAGCXGFMISDASXYAKAUZYKAYFUAGCANVLWFL9YREVSIIXGHUYVPRDSMEWGPRQAOWVDTRMVQJTRAUNCLXU9A9EEULLLKATARLXHWAQFULBCCCJSCZH9DEAJABTTKWYD9E9VWKAAUOKBGKIZEYBRIPHDFJA99VARW9ZAGVK9MKDRHKAXTLGYGZTRKO99DBOZJRZMXSD9TJXQVRALZUKQFGAYRBJWAABJZPSWHTYQDXDMXGGFYZRSLWCBXWBQRUREXMPMMHVVPDPUUTPQS9UOYWFEKK9",
		"WYUDWRWLKAVEUGBLPXDSRAZBRFMCHSJIUHRZFQSLEIMFRITBBC9QEDOGGUMJXUIQSGC9UO9RKPZUGKBMZHZBLIJSYMIFYMBXFMMCJCDGXNHXPP9FSHWDFZKGCIQDHY9WLOGY9XKMKJCGOLEZOLFHMNAIYBVXFNUXJYBGZ9IJLYGQHUVYYHMRUCJENQNJFKIEDSKWOSDMJEJXXLVPRBQGNCAARGDUTRSQMMOLSHKRJRJJOJIZVVAHSZGLRHGSHTBCKZWFOHQOJJAJVSLHJFNXMZBDRLKYAC9FPLZLRANUYTUKRZMGHDQ9QYESWIBTDAHUJWXGALPKIFBGOYPCTHNVSOEBBLT99JUHFPKY9QHDMGRLEMSUUAXAEPPXEXPMDTTUABBIABGOIHV9LHJWEBXECEGDILEQCAWQONTUO9BMBRIZUYFTCQOSERMHNPDVGHXWJWIPRJVYOMXDLINAW9K9QKBUJWAFNSBZCBIPKOHYSQTP9OYIONJGEMXOSJDPMHMJ9IARWJWDS9YC9GFDOWUIKRDBEAXCASFLFFLJLWVC9DHMSYNHPCQVHRFRYKIASIOGKKZVXYMLYHDKOAXYGXWIZEE9LADXVMQIADYFLPRSFZBRJORU9MJLC9KKKBQYERS9OZPKZMFAZJRFHTJDSVSIAF9RCCYKFCNYZBVATBPAAWJ9FDKXQARWPCZI99QSTLVAVJFVTIKVMJOZZBV9EAUMZPDCDPJSNHYNVVNPINNGUYVZLXDYXI9MRFUKTENHUGTEI9DIMVHEXWGLDUCBYGHWQGYCBALSIQQYMFVUNGZFPT9TQPYNCM9UV9WWTEOOUPIIGQELACGHOBPMIPCZYXSVQWBYHQFKBSCJFWIQLIKDUBXKUWNESMGWHFTXWUZHNMSRKRYOLDROO9GTCJPSJZXPHLZLKKQZJRLOVZUIJDIPGISKUQDPKEHLMJ9RRESAKSOWYLCWZWQZEQOMESAIKGJNNIVTWLVOSXJCBNBJVYYPBKJIOFXOFNUCL9B9UDTEU9EJQBJ9SYUGXSNZUUQGECEKOP9OIXFFJQGNROGMCTZSTRPVRNMCAYNDHZKYAHLYSCOI9FGCDHWOLOCDCGTBUFCSAWLVHKZUR9IQZAMMSM9EGDPSKVYHXSPRUDUBENUD99URN9GREJHDEETQCZ9EDMBX9DDSGOMVGYRCS9MKVPINSUGM9WPKEPIFO9APJSQOE9VGRPPCWSWZDMQV9PSAQRZ9IDYRAIJQPSRGQUGAD9SWLCHYXQFATIDGIX9TPUAFVMQJDMLYRFBJWEOZSA99CLLVGVRGPEESJRWETOEHABEFBITRHFKR9BSVRQUITNOIPCQBPINVJDLEEGLKQVPLPFZPFFSZ9GJNCWHWZIVDYIMKPZEXDFXS9CMFPBWZGXALVZUHBCPKRZSOFH9ZFJHNBDREYULIVRXKHZZ9BUJ9BFDPPJJJHXDWVNBQQWXASPTHTYMNZFZZYJXHO9AMPTSSLT9PFIYPDGXNZOQZFTMKAZQB9IMX9ITZXHJNDPHIW9NVRLMLMBZQWFHUQ9NMZDEKKEFJWVYUXXECNQVCQYAVVGPWUDWABQUDRVLCDSFNBWLFLWAMVYSPBGSZSBNOGFZTYEDAD9DGIXVBNIAQKJJBKKAO9SKERNIFPHGEYXMKRQYLVJDJLRYOLBFH9WUOUQTRDNMZNYVVBRWNDES9MLIJLJIFVVXIQTGVV9ORNKTTGTAYHERHZAAKTKVIDSCDFXQRUJPARCGIQICAWZAWOAU99ASDYAPXIFWNJYWRMXSRPRMLKXOTHENMHFLRTFS9UDDBAICRYLPVDYBMDMHXUB9QXVHDNACQYBQPVXMLXQHUXPTFXQNOCVNTEXPWHOXCDVNCUW9XZKGJSDZDZIUYEQTRWBDPY9YIKWOAGGOFXBNMUUGBLVYQFJOEDDMQRWPORAYRQKFLEYNLFRCGNH99MYGEQQLQWWPYFDKAEJCEMOMJJLKUISXOJURFGPVVQQKWMHKP9ZVLRYYDQNUOEPGRIAASZRWJHKLWUIAHMHYMIPNETTBWC9O9EIUETEDJVZXSRZZVEODLUF9SWNE9FPKRCW9GSYJBPZSLKUFKUDOOHU9EWHRZPTVYPENMPVXRVIAGTYZWEZ9VWXFFIICRE9DENHXRTJOCAGTFIAVLYSURHLQCXRMQRSSNMFSRODZPYLM",
		"SFYXQVLGFFCGKMDLIFVPGK9PG99PVPGFPKQWQXNLMEXTZHTPAF9BFLQSAOXSZAHEYUNEBGXWDCXLUHIWPEZTXNGOFCMPOJXXWRKWOTATLWPHRQ9XFWBVSPZLBOCI9MOQAGRGFPHRTL9DDXIGYQ9YYPRDIRSSIPDR9KMZMXGBQSKZAZMGCJRWPXOEULXVQQLMFCUKKEBDHQAVNCLMDXPHSMTUDBMJVJ9ASBFSHIIASSMPAJVICSUQYJYTCVBZJDSVJNUJCCLZPXCIGCVAMZBLAHPXYMIWLUFRTFATVWJQUFMBJWVRYSQNGWULVDCOAYRJGPGDOJLYDRHXWHSHVRDZEBGEQUONMGOUSCOZPEREERAFQFIDLSFOODFEGAQBBAQMTHMFYCDLUOHZVGCIUEBSWXJNPPZPDYVZRIWICXTQQAWHZGRMCEOLS9VSSBJDAXBSISASNDFKMGWMQQZJWUQQWEVRAQRQDLASBTLNVKZTQKQE9MCASWVEX9RP9IUEUKKOZBPXTUDETTYXOCCCZJQLDXCE9XKODQZYBBEGPOPZAQZOKMLUIWZDCCXIWEJCIEXVTPDGYMD9RIKBHDQWKFLYBSBWXNCCNTKIIUXZOTC99QISHKHAQMFKY9RJIKRBGQ9UDNCOLGASCDHZYRQWVHZBTUEYLDGYYIBVZCUBMHYQPAGCTKUSCWEJPNCABWLYGCAISYZSMPLQIMZQCQPVC9OMBYGAROWTMOYSXTJIJGTWJHFZGQEBRTTVEOIQYWFYPKDSCGAKAZRDQMMSB9FRQ9LYRPKPPMIGJHUR9KLBEKRNADDCBRTDKEQSRQFKMK9IWAMHRMFGAXRZWRBLGWMMHVCARZVILZLCS9F9ILFAPXVFMGCXKCOWLWGRBNYJPVBXUQGGIUWVGQBWTPGDYSTGMWRWBDIOKXRSHWCPKSCMBIJGRPPGOKAGVRFCZLJYRJTLDEZWBY9ONATRTADWOLBXOOHNFQQSLAHXDDRPEZHTTXANRVWURPVOYKXHGRDKEZR9XKGAWPECMYWCCRHJTESXFTUXSHWQGCIGE9XGVHTINBYC9HWTNPUPURBR9FVZWVROJZVRFKXLVYCMBWZXQR9WMIIQKDFTLDGSAIOYFDJTHVUYICFNVBHCPCNNGPPYJNOW9PZKENJWSUKQCKLZEEVAHSBOEQDOXGLOFYX9MBBWOUKVSPYLP9FSPVYMHSJIRESJ99UEZAXBOCJ99A9HLWFTYGJJDRUMQGH9SFCGZZNAOZBESKDSJ99KBRWMVZGCFUVORYFWTTU9EJEXVPKBOYTNRTHQFGGFEFDE9NUXCF9QUYCXRVKKJ9BMSHMOZRROMNDLDLYWSDAHCAHXG9HLLOPD9KINRI9WHHYJZYWTYBDWDBMBSSD9EIDWQL9EUXPQZGNODLHUZSZWOFTWZBUWUEX9TJOZUJVEYFUGWWKFJIQHZXIFRXLFLEOZRBLGQHPBW9B9UCPROLSESZZAJVC9TDNNVPVEXKRWOCOJIMZOYHSDJFLFNEBEHFDNOKIJURCYAKLJQYZSSKJVMUB9FUOQR9HKYRUUIBMPMLELHJCSAHIUYVYPHNEDUZNXBDDSVD9QVPJAXDOGYZXLCDYKCRCXCQZCNJCZVAXOKW9GXLOMAYBPNVRZKIIX9QHZITIGWENQUZDQJVYMGTK9PWFZUMSPRGYQW999I9ISMPZ9DKBWDZIEZJSWBXUDPVDFMHNB9NYQYNUQWGQSXHRDDSG9HXGBNOCNCKWPAEVPDNMADGTIAXSQWDGLYLIHIOAKYQW9KRBEPIUGCYFNRRIYCEOTPGDBFOOBNMCHHYGUHNHKR9ERXIELFQTYB9ZQDIWRCJNGCODGV9YFMUDJGEUNOYWMWSQHPUPCWHUCEZAZZAXYADBDDZ9ZUWXKMRESQ9TDQF9MPUYGMHTROEOM9KUOADFITEXS9R9JOOHBUZRKHP9FPRU9VJZXLXINOCHBBUJWENLQIXTQHHJBKSMMLLKTNQKPIN9YUYQHQPHVOBBQODHZMOV9VHEPXAGTZIIGCVRJJVCB9FARZXWUTSICRGXBMQVMVANMMEWNFTNBVVVEOLFCODBNLUHZLMLLLLVOPRPWUUPBNFTWJSFYKLOIFSAABAHR9SWVHUEJSSCYRZCLHBVEKUXTIUWQMUYGAAUWHKIGSMRPCZVTKFFWIEDAOPX9WRCRNBR",
	}

	digests := []Trytes{
		"CPJIVRRPXR9GALMWKJEYDNCVQNAAMTXRDACBTJKKLQOCBSBKJWAJOPDQHBWJCVQN9WOFIRNVDHDJXNGVC",
		"YYLMTCDLVRSRGFKTKYGMKOIYSTANUEPJYLMN9R9E9JIPGHURVFDEVNM9FXXDJ9NXBRWNWJLPDHVZPYFDG",
		"HNGTVLYPQQY9D9CLIBPIY9VLDHWFTSEMZWBXOJVPXQLQHRQTLZMNJRHFYS9RQPTGQEDLKQGJWKHSFWEHO",
		"CJDGHUWLZXWVEEV9EVTWWNJJBJMELKIWMLTPBREIGLXUWXQ9KMFTDEAAVWNBFXTMVBNB9VUDKQYF9PRPA",
		"VIOFEPBHHZOFZIXQGPYHXECZRAGYHPEVBGDOVCRZPWXC9DUNFG9PHAWAYDPTIBRWKVRABXLSDBIMCDQOH",
	}

	addresses := []Trytes{
		"SKCPWXDBWKCGIHILLGAXUFFXGMDNVVZYLKZUJIZDCYFDEWDNOIMZEITSQLEOLKBXV9WXOKFKBXIZSGSYL",
		"BNXKSMVDRLWHYMWFOFUUPMVXHK9IQLZWRCJSLUUIWFVFZRKFHWSXAVZVSLMEAJOVKWGMOYBOXJASKHJBW",
		"KVBWQSNZRYRSPCTEKQJYZBMLHQCFX99SEIQKJVBVXFPSQQJSUQXCAJSTQSZ9ACKOYIEXH9GPHFCHMUEQS",
		"OKWAMANQEUATHUQNUS9ZMCOKMISQTEFBZOJVCRUWVMRINQSRNCHGOITCARCVTBGSAMEPYGPXAIMJMRVIS",
		"FDHVVUQAEDHQPMNYDJLNIUVZEQDPTU9DOKCWK9MAEAM9JHTCUFUHVZDXYFHUWUXGVUXYUBLKGXSRXBJUU",
	}

	signatureFragments := []Trytes{
		"HGHPPLSOXAXGXYZTPVVMM9WALFW9WCMYETKMXJPZTXUXJHSXECKKSCRGHHIWVUMTUFGVPXJJHGMHBZVRACVLRKGQSRHLWLDHGP9RHGHPRGNVCBWPLHCXAJESPJTGGRWZUEIZRUREXCJ9XJBGSZZVQLMCQJISOHKKOFLSGCDRZNUMHATOHYKBHNLCPU9EOJHEBYQZVULFTN9GWPOMIYVPGR9TNYXFUTED9IUPTORUIMYMUBOWUVDCRMCGZWMKSKPHSTCIRPIBNJJGBRMDSQKPSB9LXMRQNKPQHGXMAAFVRUNS9STMFDYQZQMIXXHCSPOCKQJZDOOQRVXYGVELDPCS9ZVSQBIEXPXMIOOCSTEJJJJZUWAMGHCZRONRWLKZQ9UNWEWSUOXVNUJJBZTBFWFFZTWZHDGSTKTPTWKVLVBND9AHWEHFAORKTOVTCHQDUJHRSKRG99BWIXJGEZVWSYSEBHPVFZZTCCDPDALNGOKXJVZUWRXVAPNQSWQHZKRVMIQJBEQKLJAQVLKJUWBYPMEOIQDYDQEX9OOSRBXSQKJRPWEHXLFSEXEXAUSXYEQRFIZNNMDQSYIQUJAPND9OFKT9AWGQPKYNSSXRSHZNHVOCLWJE9E9DODOUFVLZQLJPNRQZFOFBNMQNFSGASMJYOPPSEGZWEDOSPVUJLPKWM99DCIFWIARHAQDIQEVOIZQNFEUPUBCQNGAVDUGLYEXJNZBDJKJMHQICYILUWCJVRNUOIWBOEHJBQVOFIDWYOLFKCSQEYWYZL9YXDCAULSMBSFHBZWWCDXZAYMPOCRWYYCFFGHH9AVWNUYNEEWWLROPRFDUDOBWAFRVFDVR9KLZGHHBFJBWQVHRMVPTTTXGGRLKEJGIOVTSATGNUTVHKISCYEECQPOQNYXKMCGJCVSXYNDPKQIVRPHUCHXTPFTYSMXVJZPHGWFKIURGWMRSMZFEOZBGQBZHURDSKZVWDSZNPCNLCVCFZAHVITYTJXOUVINWVDUZRSSTMTIYVQ9VXBHNBELRUIRYTSHPBFSPJBWDJNDZK9TKLNAXVNHDMTWBNNEBDF9XTXZDBGGLXMLW9YNNGC9YTVCPKNF9VAAMUENNJU9DLILQNRWMGTVIEZRNAQWYFLWDWZGNPDFCTCOVBEQNWVLLZZXJPIRBQ9VMQGYYNJEXCX9GNFDREEIGLTJMSBLHV9OSACLJKRGAVBEOFMQPQNNYQPBBYLCLTSOEBPZCJNQ9QHCWCRTLGZLJBPJSMHWQZMOKNDGZTBNRTMTQLGAODZSCQCFWGXYDDEHFLUVGMPZYTVHJYHSFSCAONZNVPPXMRBMSWGK9NHQKAUUYMIQHQSSFQDIDVKEHXBWITMCNTAUWJGEXJSGHFIXYMNEAERTRSLEJOGMAMWLYWRTKALYPBNJCPYGDKBCFPNFGZLKGPBLSJLDYDIYTHBRD9PKUFILVXBPMQIVGGTUPEDBQXWNMWLZFLEZQWWAWGCZBADQHEBMEOBDQRRDNQ9ZQCEMNQABFRTFZUIEL9MJXKYWTRIREUQUBOM9NANGBMUGMDNHNIIZPORLQGAK9WAYNRHQRTHRREQWBWGUCWBIHBLFZLFXQMURPDQQLUGBITEJNYL9WJXDPYBEREVUEORHGQRYO9AYBRW9PDPPBQDNMOPJNSUZQLRMJE9YFCSSEHJBRSDDPSDMOF9HBMSYJFTN9X9PHLHQPURKTQYMFDBDTNDNGFXIEDJZBUFCBZ9MAZTYPREKTS99DMLKJPHZE9Z9MOPD9SDEUTJOTIOANGKXLLPB9AIFBSGRFLYIKPLSLKSWDTXAEGKMAFIWLVANHOIBEMLGQWAWNLOFIRCHYZEBOSMZL9EDM9ECSYWGTHIZAFWOQWROMBXMOEIKSZJO9GEDEHSKYCMQICHNGTFOEAGZWIJYCHNHEFZQZQPETINBFFJPMGZYW9HSUFZLQFDMWPECWJHRECAWFZZBSQSGBASBCAWQPELHYFKV9WUWJOISBRULLDLMWLXGSYETLKXTQPVPIWUEYCGKLUOTDMVXJYGHHBMYARWNDFPMGFVOPCAUW9GHFPJZPS9ESPCTFOWYHQWIYAJGCCOJWLQVQZPSIXJFJ9TE9W9RFXYOCKJ9NTGWBTNQLAYTSQASRWAZLFVBMBEGIDSESPQSUIOOVCGVKMVWHIJJVFPMWRIVVSJPRWGKSIQHE",
		"QQLRUSXTIPYVNGBVMCB9EBLFEXJNRIPIUSL9IFHHPZQOLEMJNKCAKDYKMURZJXAWSXIFDVZYWQEXXVLQHZVRXKUMQUXTTTGAFCIGKJUTLZKBTISABSJJ9PORLIM9FHSPLNXJPXBAXGV9TCYUWBEOPYSJOEDJGZKHHRVRTKDFGFFKW9ZFWJPDJBPPXYRWNYJNQWIOTCCVBIEPUVEPHYFUUOFVPHGYPGSOBMJAMRRVEGGEQLLNN9UGSHEQNELMTGEVLNHVBBXKM9PBHSLECSFBXAGMKV9OXSDGKRYNQLPDRCKXSJWQBDPCIONUDMSTNQUCLQOABVBKQYGWURBALTRSXSEV9PZ9KXUADPSNXGXTKSFCVOTPSVKPVMNNZHRHEMVWEYNXURXPJBKOUXZFETZROQAMTYRKKLFLJZKSTQGWKBWDHQPXNDMNKL9LNJKUAJPTVFUMN9XXBC9ZDATKTKMXXGFNGLZINUQBSHXGXJQRNJCVCTODRVNYRNDVCRGEDTRPQDEKBWASWAGSO9NTFFYKOKCNNSTXQLM9ZCWOLPSKMVNMMVKUHQDFAXBCXLDCLLPOLMSIFSCAMOIYQN9GHXCG9NDJCAVNRIAUZMXCFXYYRLML9BAVXWFIRHLLFCOTIQQBCJXKDTWYYKVJOVFS9RRTLEJTEHKOXXQUQSDEVPPDAZEFGURWSXVNUFTIXWA9JLOYJCEGMGZMMAU9BAHRZGLXFXONINQXYYJFVIVKWAGGL9OALF9K9GQIVTLPAGASBWDHNAXVVEVUYEEMOFKZNHZVSYNRGIBNOCODKKXPK9HDO9RY99VKGYQNFPUOYVGQHNDSNNX9ZSMPXYLWAVZOCWXWBAZTFLYMAERJYPKVDDULQIWNLCECNYXYNOGBIMOMZSIALBMXDTYX9OND9LHVRVPQA9GDAYSGMYPCPACRALBUUHBEELBZGHPNRLZH9MHSEHPGBSLHNTI9MAYTMQQWLEKQOYF9WC99LVMJBCAOECMBSHSQDVKDZMVJEQENVPKTQPZWCVTHQUNCUCRQQFGPZRNVROZY9DZSALRDXSCREONXHDSPDTTGTYUBYVTYJRTGHXKVUGFK9EZDORXOJUNZJMKZHIFXXDRHYGBAEH9OIZEJCWMMPMWHSWXLXLUZVCYPPZIOHNVSJVFLKZTIGAEYMF9UZWHJO9ZFJNRRVRUWZJNCXAWPNNCLPQUCSZTZAFJWKMZCINPLYWWLYYEKXTKKPMGFQLOJNNIFLOCPSTKECNJTNRYQCGYYGOTCKHFCFHXWTEAQJZXNTRMUEMUNRNTDLBNESOBMQTMLVLE9YTNPBROZZTHQEPSWXIAULSZVQGBUCEKTMCGNFUVQXSE9WXHHXYCFYXJGBAPH9DVKOLPYJJIVP9BLFQPTGCTQNWMFGHUAVCHBRPL9SYRDAQOGYLXOOBZCNYNN9WEZREFGWQ9EEYDHZE9SKYC99KZANQJ9SIFSCLHRXWICULLLTGFSHRENUWSE9CSWTDUMAZLMRUCXZVV9KWJNZQIALFTGTRSTBBGVCI99FQPPTZMOTMLIAO9XXNSNVCFEHLZJMSE9RDEVLLHGBIXUNRLBSNOPANANXPYSHPILQPFGAOSTCGSBMJW99DTOUCYCAISXRCU9AVGVV9YUKYIKAC9RQNWGIUQRHVUYIBASZUJEBYAFBTXSJZKWLCGTBEGOPJRYWPNCRRBPWOIKCZAMOC9JMAEHGVDWU9RCEVBSDLKWQHYHFICDVVCEXEVCSUCBGZWJ9RGWDUAYYHI9MXRHDEFDWIQDA9JSEDUZSJXWAWEWFXBAKLMYECASZMYWLKMASBVYTMPXFSWX9BGUN9XQJAELDRQ99JHRCUUEPANCZLGBMQCMFKQDSXRBCEDNBWTNPCOYFRWOVWEB9YCG9MKGQD9AYBANZ9JMOAGUTWHPIIDYSRKNLEMEJRYKYBJWHVRXJIVJUKUOUBCJHKTOFAJKKUAXSWCFNKYVQRWIFYJKVDBGVMLTFENHWTKIZNEJFZIZCIO9GZPYQIFOXBKYJOJGZLMGQQKZHSCWIKKPFFXNFWVBMDFBALPEBTFOEOOAWOUYWJNFDGZSVTHHBOUJ9OJWZJUQNYKEUZ9YUHWUHDWUIQEBSXUJNL9EKRLATWQVPHDFFXIMKTZEEMOJLPCLHCZPDANGMKEHRYB9AUAYJEMSZHUJOJQHXTC",
		"IDWUZDSXBJUYAVMLDFRUECQFFYIFBABQUWNTGSPNFLLZIFSSLQIJZSZSJVTZMTHTLBPMFGFCFESIZFYJFSXUMDCWWNIWXTDQNOIWEFVDPTHNFLPOCACAHRRNRVCBPYXPKZMJNNTD9PRKTMTNZVQVCQXXBYTQOGSMVTUV9LV9PDHVAXQF9GGRUNUYQ9BJPOEPGYVWVKBATTDAQJTJRNFHENMKLCHJLGBEZBSOWOKCSUJOKTPKEUQ9TLWOCTDXQHMYGOOZCBQXZFTFELYACSUSYJ9SBFDODFNOGFGCTO9SJKHNMRZIKGCEAZHYUIINAPYYFGWBEGCQHQWKBUHDIJXRIBWMLQSUGFI9S9BRAQTCOZWCQQJIZMAEOIIWGQCRTUIFTQEECXODWDWGGPEWXMTGTZGLXU9VDPYUUKTLP9KSITTZYUMJDKLFRXVLIZGIW9ACMMQIRGDXCKSNWUDUWRAPHELF9QEESLFTVFGRTMKIEWZXBSJPSXT9TRLWFKCFRK9YGEZZSASKJRRHGKKJZANBYFTFOTFCRQTMOOYP9FSTUEEMHVDBKJN9QPDEENNKKUFNPRPD9SWSNKSLIY9RDOLO9QMMIKJGDCVNBJC9LNOKXDH9HCESYTEXZNV9QXXMFIUESAHTQALAUTNZZIOZVZIZOLIYDXOMRFKQOJRFSOKLMMKAFBSTSOKOJRJZCNAUCZODU9XAEXDMOSJGSPPULSOYRWGPKUWMBNYCG9GZACHDHNCPMHQPXVAERYUPVXNNCISHJOUERHRVLSVXAYXBYJYMNQRARTPDULWJKLXMAIAOVXYKHQF9DFIHB9SKY9RREMOCHFGEIQGMVNNXGVFAPYISSIZ9EVWZSOTUHGWANETTS9RSOGJDWFICIFJYBQO9SVKLGTRDMIMEROMTVAV9AJHI99IDMKJZGQSFFRPWYKRG9NZTEWLGCBLYDA9XVCGUCKOIUHPKZYRWBEAYQUDTQDLXJJFUSJWWOIEBESLVRLQZHIIPUKLYFHHKPHISMLNVWAJDJOEBPAJOXCEQZBOEJPFCKIRTQMQKWIDHXDVWBZXBFMG9ZKJIFCNTLF9TZJZPSGQQSOXGRGGJFRXFEQXSEDGQGEOREQNCFPFNUKLKXOJSZATXFWCBDGYL9HQESRXXNFAPJWTFMYYYKAJMLXQDNBWICJEGDNWXTFPIDHD9WQPQHGERNYRPBCL9UWKYM9NEEMKUYBONLILGULTDNLQI9BSSL9ENAHADPEGTUHTENTOTE9DXZ9IN9WFPXUFAAUF9EUEEXRPQNSUUHPDSFZADKNYMLZALB9FKMDIHCNOMIMNXWXYVPZKQUQRXJVSMGHWDLKGZHMEZOHDJSDHUKYOYUTOHMGNTPL9MCIOGU9NPISCDWWSOTLEGANSE9RRHMFHVGVSIZBLGNEGRMWNOYBABWKPEJWNFCHTRUPSBY9EBBH9YMZOCRVWFCIJKSIRRXMMLBLNGHVJFLTFGVRXLZYIME9JDFU9ZLAHOVKYCRQT9PKOURKRSQEDZINYAHDFQAQCTM9WKNOFCNWKUCPR9HAMUEMGJXCNB9IYUGSERFJNZXWW9EVQI9YJDXUIFKKGTBZQMVLMFFCZFKAREWAIUWSBLVCLRQVJ9RJRVSLXPRZRAOWILFYGITFS9NDWFJOXCZ99NZMKEL9WVV9INFSXWISSFRIVSYVWKMRVMJNUZFVCFUJELLTRYPVRNTGP9ORYBXI9VKMHKUAJKVGJUKFD9KTLGUVCKHTLT9ZAAUTEIKAEYRVSAJGTIIUQNBIHTNPIPUJYXVYDGPZWPHOQGQPYQIUHCBQZBCMMUMAHQHB9LERZXZ9ZM9PJQCAFS9LINHFMBRTYCMOTZDVLMMOKPPEEQAPKHRIEYPHZPMXLX9MJOMWMKJLXBRUWMLO9GV9WXDMTCT9RQEJLYCSS9LHGUVEY9KFLFRJJFZRUHFLNN9AMAEWCNGRCBXYC9QAEBFRZMJSKAVV9UZISAFMKZB9DTBRXOZNBVPQJJFXRSWVYVFMQILQRSBVYCMFISQOKNQBWCWHDZAZZIKMJUYPBNHYKXYDSX9LQKODMMFFXDLYHBPOJPBFIBBSZDXTNLZHQ9IVWSVVOJHZUNDG9DSOAFN9IFGRTXBSLLHSXYQRUHPP9JICZOPPHZORMRPAEIGONBUMRGCZDQVUPBJHWSOWBHZOFVFYA",
		"UIBLFGYUAKWVTHIT9BFE9OU9VPKICTNTMJXRWUCTZCAQFQGKQEYQYMUNWZZDMFVEZCMRUWDZSSWTMYJDEPAIJCSRAFHBPZPCRGRSBXUZJSWZ9HLVBETDUHDUKBIWZVJSPWWKBGCPA9RTILLZUBT9O9LNCNIOXMZGHYWHZHVWMUQCOXYACRRUKCPEG9AAJK9HJZLQIN9ATRGXHGOOCANXIDSZZJTYTOUXQBQWBNCT9HLINGNJTDAA9XCAWASKWJKFGHZEYQAFKCJJBZDAFZVBUTRUNEPLJY9XHOMQEJKMCHADPBMIRCJKEQOOXNTGCTCCYIGBTVMEZJYAAJQLP9MZQPARCTAXLTXYBPRSTRLTPJYUXYSKXEZXGWBPADNMSOWJL9FQSCAHJGGCBUMCAVJPJNU9XDAITWUNWWFKTUNFOBEXIUDDVUMUXTXXNODMKLE9FSMAD9EOXERAKKMDTBPRBLBLTVNQHUGBHODYZJKEMCRDYRJMUGFEBEFPLZCONETTLJBUEVZBJNQW9HRNWETZQSXZFIHOWPOHQY9NVQQI9MBIPYMOFHNEMLQVESSTOZYNVROUFBOMLIFYZEQSAL9GTIIZV99ZVDPYCQVPQTZBZCQQVCIISBSEDAZVKWAQUWRWSZBOVKECRKTGFOLVCGK9CDQUUMTWSZSTYEJWXUBKP99GRMTGEZPYFRFDUUOTTWXI9SKEXSTDMWMVKI9IABZOXMZFIRM9IREPJDFFERWNZVRWGSPCXMRIARDLOZYPUGHCOSEFBCW9IQOZUTRFD9YNFXHYVETFDFQKSHZJ9UFPPBAWMROTF9BXAFVXZCSGBSIPEIXESYOGNVSRZRUVXIWCRTKJEKIKEVBDEFZGDITSYDSMQYFOX9QLJOZPQNKX9EQPUVGN9OREB9ZAHAAJRNDIQ9PYGNQZDKMMVDVXSTEWUUXFG9EUZGQISVZNIAIBVFFPXSHNSHFGWWYAQN9IMMHLGM9PBERWHIDLJN9AJKHRQRRNGIUHRHQIXMUVUBJ99DS9AFMTDRV9BHYMMWSYAUPHETVUKJVIBFSD9VOFOCKDPNVDVXMHYMKMONBSOLZ9PRDIUCEBJQ9JOVKBQMXDCTOQTIJSLSWCPWDVTBBGEJAYXRHDC9TZWSV9RALREJHXMNAOZ9ICYJFSOKJOOLMJCSRTYPAMSRWCKOHLZMYECGVHVHGZ9IESRMK9FSLTHCIJZMIFXXPUECLIM9BUHPKHEIGEYAZZMCGCPXARC9ZPHJGKTITGYWLGLHPGYIGIVCDVDVCNCPQXCKPAJSCZSDRJCDADKUJFOXDNTFKVKBPXPGGOZWMLLZQOTJFALQIKGVVXZZGUWDIQOZQXTXYYPFMXOKYHHGWPEZPNZNFMXGT9SYTXA9CKAFGLWXJMP9MSQPRAIIWEPSVNSUFTNFQCVXCGH9AREHUDIJRVIXI9KZZDMUKAMDBJMZAEKTMWMNZPEQHG9HEZIJSEDGWIFZBKSBPLHEIQBERPV9OEUVMWUHEYGKYCJQXWTWXCVHGAWLBIZCZYGCMLHGWGOYVXTWRTKUVTUT9DAZJEAH9RAPRHXIYGGAHYCHQRQARWPDOOLBRNWYQSWJDHQIJCJBCPIOQOUJXSVXHTJC9ACKNIERBONLFRPTRZIMWKAYSSVKNGUOCADQUWQKLSSHNWQQQPOHFGGFRZICIXCPTQQGII9RKQIKOMRPEVFOOQZ9XNS9ROMSBUFMOKHLEAESRMDQJNINGTDKLWJNHSFBSXVGDGH9SYCJ9NGCQBWCLSAIMG9VHYWDHJCRRQHBXHYDCFLPRTQBEIHUATDUHENLYFNFR99NPWFHFQNBLBRHPPFBBOEBZLXSUKWHMGCHVKQFIDYTYVPDPVDBOPWKWLVOJQSVU99HFVKLNXLRZFB9EKX9EKDABJMLU9VRNBFYLISFGTBEYRGYOMDWCDXSAYPYTDDSC9DJLKCLI9WAZAGARLKFLMOFZSNVA9HMYJAQTN9FXGHEQYUHSZYZDJGGHNCONQLYKTWFBKVYJVCHELDHTWNPHVODGREXRP9MNVNTZNGAVANMOWURUTFICWCXEAVTNCWRQBHZDL9NPFZCXOEZDMAMREQNZXNNBNTUW9CISWY9MSVWPBQRWUDVHMOLWJXSDUOUSQYKICIYHXEQQJQPGFPULQPBMDYYNYBNPSDOZBFIZEKOVDZKJ",
		"QPJBAAIBVAEI9IPQAVBYDAGWGINPAVVIOJCGTMIYNEBVLSF9MWVMVOCJWMDWABODXSLZRUL9YTABMVVKRARVUVBDMHPNINAMNJCIEILEAMIJTOZNRS9TAUMAMHWW9QLQP9U9SZTSYUWRRCQI9UZHZFAWCWXRTXULHHQIMTQHGWRRLDZEBPDJVPCYHHYBOHKBHUHRLFNDZASR9MXXCLUSTNZHOHGWCPSI9DUHKLTIUZHPFI9NCEQSKOEURUEFVGCNNMYQWLGNFTYPXOQPVTOSQODXYZURJDCNECGJFTLXYFGZIALCMBBYWLGBFIWJXNYRLRQCWJYHVWCJTQBNDXXMEGMBD9SFXVXDLUQ9WH9GLG9RBL9BJLLWGUH9XWWLMDYDIYGS9OE9EOK9ALZAKCIMIJWJSIUZQJUQEZGKKYLHUNUQNVNUHDSXIDWOVCTWOAU9OIWG9NKDDCBGSCJR9AVBEANMCOHMYTYOGTLXBUHVKS9LRIJDZSTAFROPMLGDNQGAKUUOYJZMKJMOGUCYFOKOBDEDZSYXHA9WSLCZF9SIQTAZHGILNQIFFZPVFVPQQVFLYUSEECEBWAQHHDGPH9BTIWCEYOJCJHFU9FAZKTAHKB9RDVHXOVYNXK9CAACBZLCDTANMHXTNZMYSHNZTBRYJMQRLNITLVD9VGZNOF9ZKTARVHDEUA9XXUORYTPCBA9MLYATWZ9VHOXQOTRROXXFPKDUHXBQLBOBKGPUYCSCQVOGXRWTTDEPBWTVFGWSKUMZ9SYRRKCGLHNDJLYQNKORRZUXSIEIETPZULUDHECEDNHKEEICXEX9TXYEFFMTUAMIHOMOIVQGNTTV9LHMXAFBWCDBAWYNCVBUBPYIKKUCSTUEKACTFHEKSMVECQBHAAQGDJJCVLCJXPVGSYFNQKQNZDWHORVGHTZUCFSOUQLZLPZUILKMQGYDQJXXS9MBL9NGQNBZWYFKBHEXWEHWNLFGRHTOIKBXOZGBNHXATOZ9SRXZLEDQSYOV9HDYHZCLYPMISSKEXTICPVPCMLJRXDJQVSJVPUHQHOHIZYOLUILCRWPUIKTVFFJQXIWGCCHFWONCWMJEHUCKOKPYENDYUBPPU9IUTOOILMOUGKYKPW9BBVKCGWKGYNVYKAUHDPSVEGDGMSPXOIULAIDMCUXSNILIKJEI9HHZSLKSITCTGGTXHOYRNMWGTHQRXKCYQCGTMQLABZOVUMTY9PNML9KRSFA9VYRWGIXFRHRPDMIJYSCUKVOASHEDZPNPFRPIFGILY9VWYEXIGFLABSYUXFQMDMS9SWFOHQLULSRQK9PUNGUINZPDDGZRUNQCIWK9OU99XEPD9DCZYVUVCXREKPPCQNOORVUKYWKLJCJTJZNSLVLBBSIESOSRXKCVABA9UYIETASXVXVJLWXWTB9GEIUOLL9WJCYQWVML9KIXZUFYTUR9DQ9HEPNOYYCBDDBWKNGAQNACXHCUDSXSBUSAQLKYRMWQMTXHFGLKLTRIJQTQXGFL9S9QXFWFOVGUENHBYYQYYUOG9UAFCIFWUNIIQGHAIGBWRYWBFA9GKATNWCAVGHDXBPUJFALALPIUBLZURGPPFKZJHIMSUWSFAECUIXHJ9UUMCFTCSUIZQHXSHFNTTAQOYBREPNEBWWKTRSRZAURCALYTTWTFYPXTUYRMOAKEZLAQXJ9ZQBDFOWKEP9MKEPIVUFRACVALGXGCKIVSLNBGKHHHGKNLIZTE9QZPGNWMBHGCEBRFGMY9ICBFMG9WGNCJSKLPUFIFCUFMMIA9LWLSVXRXLNDASTLKRPHLBFCRURWVQER9FYTEGRJTFHGOGPCOEKADRRM9CFUHGTSVGSCGVHLSOBQDADE99LECUZEBLGMUPVEZQGJLKJEKEVISPGIAUNPFBSRAV99QZRKITENHSJQCUYTAPAZEXWVONIXIHZFHGDZEKVLTKRDAVXJA9GRZUXZ9JBUEKYXDYYFIIL9ZPUPRIICGADGDEEMUZE9TMLVCP9DYQKEOMWVZW9UGFSRPEKQUFBETKNTVNSGHOOLRUJDMYUHZKNPUWLSSQNHS9E9RZQOCEXIRTOJEMETIFSQNFPXYHJZUZGWZIYWZGAJMEHENHBIKLILPHWZJFMBECBUQJLJGTDSASKH9R9GTATVYSOZIMZLKKYEGAWDFZZKEZAJMEKNPZGWDTOBA",
	}

	signatureDigests := []Trytes{
		"CPJIVRRPXR9GALMWKJEYDNCVQNAAMTXRDACBTJKKLQOCBSBKJWAJOPDQHBWJCVQN9WOFIRNVDHDJXNGVC",
		"YYLMTCDLVRSRGFKTKYGMKOIYSTANUEPJYLMN9R9E9JIPGHURVFDEVNM9FXXDJ9NXBRWNWJLPDHVZPYFDG",
		"HNGTVLYPQQY9D9CLIBPIY9VLDHWFTSEMZWBXOJVPXQLQHRQTLZMNJRHFYS9RQPTGQEDLKQGJWKHSFWEHO",
		"CJDGHUWLZXWVEEV9EVTWWNJJBJMELKIWMLTPBREIGLXUWXQ9KMFTDEAAVWNBFXTMVBNB9VUDKQYF9PRPA",
		"VIOFEPBHHZOFZIXQGPYHXECZRAGYHPEVBGDOVCRZPWXC9DUNFG9PHAWAYDPTIBRWKVRABXLSDBIMCDQOH",
	}

	Context("Subseed()", func() {
		It("returns the correct subseed", func() {
			for i := 0; i < len(subseeds); i++ {
				subseed, err := Subseed(seed, uint64(i), NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(subseed)).To(Equal(subseeds[i]))
			}
		})

		It("returns an error for invalid seed", func() {
			_, err := Subseed("asdf", 0)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("Key()", func() {
		It("returns the correct key", func() {
			for i := 0; i < len(keys); i++ {
				key, err := Key(MustTrytesToTrits(subseeds[i]), SecurityLevelLow, NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(key)).To(Equal(keys[i]))
			}
		})
	})

	Context("Digests()", func() {
		It("returns the correct digests", func() {
			for i := 0; i < len(digests); i++ {
				digs, err := Digests(MustTrytesToTrits(keys[i]), NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(digs)).To(Equal(digests[i]))
			}
		})
	})

	Context("Address()", func() {
		It("returns the correct address", func() {
			for i := 0; i < len(addresses); i++ {
				addr, err := Address(MustTrytesToTrits(digests[i]), NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(addr)).To(Equal(addresses[i]))
			}
		})
	})

	Context("SignatureFragment()", func() {
		It("should return the correct signature fragments", func() {
			for i := 0; i < len(signatureFragments); i++ {
				sigFragment, err := SignatureFragment(MustTrytesToTrits(hashSignature), MustTrytesToTrits(keys[i]), 0, NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(sigFragment)).To(Equal(signatureFragments[i]))
			}
		})
	})

	Context("Digest()", func() {
		It("should return the correct signature digests", func() {
			for i := 0; i < len(signatureDigests); i++ {
				sigDigest, err := Digest(MustTrytesToTrits(hashSignature), MustTrytesToTrits(signatureFragments[i]), 0, NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(MustTritsToTrytes(sigDigest)).To(Equal(signatureDigests[i]))
			}
		})
	})

	Context("ValidateSignatures()", func() {
		It("should return true for valid signature", func() {
			for i := 0; i < len(signatureDigests); i++ {
				valid, err := ValidateSignatures(addresses[i], signatureFragments[i], hashSignature, NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(valid).To(BeTrue())
			}
		})

		It("should return false for invalid signature", func() {
			for i := 0; i < len(signatureDigests); i++ {
				valid, err := ValidateSignatures(addresses[i], signatureFragments[i], "BBBBHANFEOTRSIPCLG9MIPENDFPLQQUGSBLBHMKZ9XVCUSWIKJOOHSPWJAXVLPTAKMPURYAYD9ONODVOW", NewCurlP81())
				Expect(err).ToNot(HaveOccurred())
				Expect(valid).To(BeFalse())
			}
		})
	})
})
