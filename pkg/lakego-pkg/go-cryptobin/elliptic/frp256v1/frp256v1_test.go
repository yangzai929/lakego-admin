package frp256v1

import (
    "bytes"
    "testing"
    "math/big"
    "encoding/hex"
    "crypto/rand"
    "crypto/ecdsa"
    "crypto/elliptic"
)

func fromHex(s string) []byte {
    h, _ := hex.DecodeString(s)
    return h
}

func fromDecimal(s string) []byte {
    b, _ := new(big.Int).SetString(s, 10)

    return b.Bytes()
}

func testCurve(t *testing.T, curve elliptic.Curve) {
    priv, err := ecdsa.GenerateKey(curve, rand.Reader)
    if err != nil {
        t.Fatal(err)
    }

    msg := []byte("test-data test-data test-data test-data test-data test-data test-data test-data")

    r, s, err := ecdsa.Sign(rand.Reader, priv, msg)
    if err != nil {
        t.Fatal(err)
    }

    if !ecdsa.Verify(&priv.PublicKey, msg, r, s) {
        t.Fatal("signature didn't verify.")
    }
}

func Test_All(t *testing.T) {
    t.Run("FRP256v1", func(t *testing.T) {
        testCurve(t, FRP256v1())
    })
}

func Test_ScalarBaseMult1(t *testing.T) {
    for _, td := range testKeys1 {
        key := fromDecimal(td.key)
        px := fromHex(td.px)
        py := fromHex(td.py)

        x, y := FRP256v1().ScalarBaseMult(key)

        xx := x.Bytes()
        yy := y.Bytes()

        if !bytes.Equal(xx, px) {
            t.Errorf("make x fail, got %x, want %x", xx, px)
        }
        if !bytes.Equal(yy, py) {
            t.Errorf("make y fail, got %x, want %x", yy, py)
        }
    }
}

func Test_ScalarBaseMult2(t *testing.T) {
    for _, td := range testKeys2 {
        key := fromDecimal(td.key)
        px := fromHex(td.px)
        py := fromHex(td.py)

        x, y := FRP256v1().ScalarBaseMult(key)

        xx := x.Bytes()
        yy := y.Bytes()

        if !bytes.Equal(xx, px) {
            t.Errorf("make x fail, got %x, want %x", xx, px)
        }
        if !bytes.Equal(yy, py) {
            t.Errorf("make y fail, got %x, want %x", yy, py)
        }
    }
}

type testKey struct {
    key string
    px string
    py string
}

var testKeys1 = []testKey{
    {
        "01",
        "B6B3D4C356C139EB31183D4749D423958C27D2DCAF98B70164C97A2DD98F5CFF",
        "6142E0F7C8B204911F9271F0F3ECEF8C2701C307E8E4C9E183115A1554062CFB",
    },
    {
        "02",
        "DE681B2898119885373F7EAFDDF94CA0A526794BDC8DA00E0E463860D227575E",
        "B1240D67C641B70BE151A6D456C77BE3AF2997F8858D3E07D862E37EBE0A1045",
    },
    {
        "03",
        "A13ED122D99792D3CF188FE2C964EADA77A95CE2D03BF3451BBA94DE3E967BAC",
        "D141A90C972AE0DA909A77BB02B973176238E32431CE8F378E039DEB74989A21",
    },
    {
        "04",
        "8A62BAE6BE63CE26B48368BE0B0BCA7CC6FC33B81FD685FD1085EE128F242FE4",
        "E548BB18E8E64A4F0B3852962CB58633386C2A79D995C825D434C4B5A727B1B9",
    },
    {
        "05",
        "32B3881D7C703DB9BF613F52E9C917F096839835C025DC022967BBD1C9886B6B",
        "152BF2581D33558431F325E3543EF6EA3BEB58F793669392DAC37FB6D21DEA69",
    },
    {
        "06",
        "917473C20297D8F46097D95932559BEFD24F7D6E12C95EA8BABB446D0016E368",
        "3EC7C0E04474E0F6A26927301E8D82EBB45B36960675119B42726EC10BEAE371",
    },
    {
        "07",
        "C33DF0E273A25BCCDACCDF7BCB4E4A6CCE84A1E3BCFA7E9C8C437759E5CC568E",
        "D405E69EC2F0BEA7D242BF4D743DC6AF6E5068EED23FD28975519C6821BF8E8C",
    },
    {
        "08",
        "4165FC8ABBE27583DB3FBF73A5ABE9D1CDC0CCE74C2E9277000E9C37CF6E6502",
        "A9738EF04EADBD7DB923A197DF564B66599FD6CCEFC7C7A61FCFA91F382BC4F2",
    },
    {
        "09",
        "2F13E149CD81DCB537A6E319F8958AB924D44CB9AF159EFC4815DAB2E8F085CF",
        "6F379F283A2BC73C85EA460C8D055845D65E21E6063571B67B8AE4235E03826A",
    },
    {
        "10",
        "5A5C7A63CC1B9E39096A23EDBDB910585FA6F243222022C93C80C970869C3ADD",
        "BAA2334C228B2FECAF369DBDA74D8580BF0D7332433CC63DFF3E151BC296B43A",
    },
    {
        "11",
        "990FB8C085D200757024A7CFAC7485C58A1C8479395198D364F0C8A0E4C348D9",
        "ECD80E63F368508A38865E3EF47FE85CD685B0A4076E080910B13751F03F87B7",
    },
    {
        "12",
        "24A6E34FE7900B720B9C7D16D059639259EC84597E4270071E34C8F4CE4AA8F0",
        "E70C2FEE8688B24251836F55B268451C6B01D780FD6B1455731A5E6E15F3595E",
    },
    {
        "13",
        "CDFC5A4D7449DFCEDD07A2ADE38F1395AB8925A48573FE36DD716A4D4405F0A8",
        "DB55269DB768F7B843F852AF1719CA940316A20B5760ED9F5ABFD2B22904836E",
    },
    {
        "14",
        "AA4C778DB565B117FDF8B97F6C7D8F5C2BC8FC8EE1D45E83CB2D5F869278EC93",
        "372143067F83875C493F2BF8C440DAA551261DDF4219EFAB282E842A6F216A7C",
    },
    {
        "15",
        "C14A79ED434B7275E929BE94C4EC1B6F422D6FA2D0A1FFBD49BE90E3EB3F8EFE",
        "3AB7343D8508A4B25717CB00FFF4F8F72EF6BC2FCD64BFA47D2534AF3944B648",
    },
    {
        "16",
        "ABF3E2095C8B177E1196962459355C2528D607D1C89052B18491BFC67AE3DB09",
        "CDF69F0E228FFE33AF198684EA875DE343ED920E963E9F0BAC5DC96F67B8132B",
    },
    {
        "17",
        "9259B53230E01F074F99A1F4A66008B8F60B68676F413D41F9C04222E41EB5C0",
        "5DFCF8487290353C483B51ABE397670D1A1F8960A127E7433F28E79BEE22CBA1",
    },
    {
        "18",
        "D84EC698A661149A2BDAFAC95546FADF8B886605CDEAC1F65212F06454BE6940",
        "32C3841C5A304BE2348097C85C65513AEDFD87D1DAFE6AEE9C054845A908EE25",
    },
    {
        "19",
        "BD2D2DFDED7C89F4CC3338824842C16EB980369A31F69E292097AA007B064D50",
        "3D3C1918E28690722005F8A225BD127BF925AAC8CE5A9D008C241DF1BAF8D62F",
    },
    {
        "20",
        "6B3CBE572B1A3F5FE5410BCE760FEB5093320591ACB953817BDAEC25E78AD3DA",
        "D78BBEBEF0A416D00E855CD4C6608E1927D97066AFC19ABF27064F513833D6DA",
    },
}

var testKeys2 = []testKey{
    {
        "112233445566778899",
        "CC326C5058CC9A69BDF973FC9C884D2CE0FB14A345D75A42B9EDB4E6ADC86354",
        "C62A2257D0E3E8A6493E626EC3EA9D523BF647DEEDDE79E916FD69845DF9D60C",
    },
    {
        "112233445566778899112233445566778899",
        "37F36537B7C11B67BB28BBE176195157C45B35436A041B8EECA6C890A0FD52AC",
        "83F3246923B4D29F14B62292A32AA6AB05EC840877FF2B23EE96B0CB7DC313C1",
    },
    {
        "1769805277975163035253775930842367129093741786725376786007349332653323812656658291413435033257677579095366632521448854141275926144187294499863933403633025023",
        "0D3A00D4A5C235246E833EE73028521F6959CBFD57F879031C620804A5A15EE4",
        "8607FB8E3DC827E5FBFD3C233B1F62245B2FAA9458DB10605D4D9719BF347320",
    },
    {
        "104748400337157462316262627929132596317243790506798133267698218707528750292682889221414310155907963824712114916552440160880550666043997030661040721887239",
        "E152362AF7609A04E311D2B591538EC07211C79668838A25DD52FEFFC0A1FD6C",
        "8CFC47460B9A3A66EE8CE497D02C5D57F2AB4B08C7652AB746FC0DC73AAE39C1",
    },
    {
        "6703903865078345888141381651430168039496664077350965054288133126549307058741788671148197429777343936466127575938031786147409472627479702469884214509568000",
        "AE597AD61FF4489367D4BD4132CCFD738E53C347AA463FFB5EA193713612530C",
        "BDAF81342A5ABF8B9A62CA88D52C5B6F6873678B6FEB0B991C2E16E32FDEB141",
    },
    {
        "1675925643682395305404517165643562251880026958780896531698856737024179880343339878336382412050263431942974939646683480906434632963478257639757341102436352",
        "6D68D7E26CA83876F061A5DA7C98211B7295105D95A68809C607C3FA779A0804",
        "C422AE5AB150FFCB160D6B83D500E48074021BC618800E7CA2F97273C6E67358",
    },
    {
        "12785133382149415221402495202586701798620696169446772599038235721862338692190156163951558963856959059232381602864743924427451786769515154396810706943",
        "872741C0C28892A29F34534B3BA89ABB6382B033DE51E483DA70495A55256C33",
        "D1AC425C776ADF4392B8CD0F5C0F74D550B9CD491F0BA5FF1A5C4E9F5B6329A3",
    },
    {
        "214524875832249255872206855495734426889477529336261655255492425273322727861341825677722947375406711676372335314043071600934941615185418540320233184489636351",
        "0BBDC745E6DE704B17131349B4B5B4A74C3F1B010A54F1C7EC609D96C9D26F56",
        "EBE2E6BEA745FA35F13F6A426F08889423B03CF8BD5E71841E4503E8B70D9CDF",
    },
    {
        "51140486275567859131139077890835526884648461857823088348651153840508287621366854506831244746531272246620295123104269565867055949378266395604768784399",
        "4030688B3171D7822BD5F927E085E90731C30601166A9FEFC4C9910B65AC4438",
        "058D624C7710EF1B32212785F78E90122E4F5E25775C3F702D2389DA9F9BD9EC",
    },
    {
        "6651529716025206881035279952881520627841152247212784520914425039312606120198879080839643311347169019249080198239408356563413447402270445462102068592377843",
        "222D3259F3438A91CC662C357D35B9C9C431A50E4ACFBBC84EE558229077CE81",
        "BE913139F0196E38E485CCA8F41F27D9F56D3FB403D9262D23ECA130B6D4228A",
    },
    {
        "3224551824613232232537680077946818660156835288778087344805370397811379731631671254853846826682273677870214778462237171365140390183770226853329363961324241919",
        "BD49C1629FA9DCC53B8081C2E3D67FF2C485F098BFDBCBAAEF356D62C8B937DC",
        "44D4FA85987422B84DFDD4C661E9093F95B83ACBC103A5A55E723218A61877DC",
    },
    {
        "12486613128442885430380874043991285080254917488396284953815149251315412600634581539066663092297612040669978017623587752845409653167277021864132608",
        "0783015F0DB4715136551412DE3D34377BDC912EBACC659B73AF67DBCA7971FA",
        "A3B84AB91CC6626F90FE059D35727E271C00D34A57FA2423084805BB198A6FED",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005429",
        "E08BE120BE2B8213CE99CCE06696936E0600DAA24BD5C4D856C5D4B239187B04",
        "18C3F112709CF6FA6FBDFB3F6FDC64E05B590A9172AA70DCD948D050C31D71A7",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005430",
        "A8D03B6C783CF9303A5F7C4F93893BCB358E4137163BD9AE0D69AB5D786C89BB",
        "63FB806035A05C6E21D7892AF2096F409D01D23E5EE6B198FDE656052C56BC1C",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005431",
        "717026BD0CF7690784464A5ED334B340F164C39D7E54E528974E060B9A1142AD",
        "D225B8A4826E44EC10546F07C7283796E0F2A2419C3187A4B018EC116565D0A7",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005432",
        "97A507F8CED10A2CDC401206D7333CCFC94BB10A2F0875313A602AECD3548A15",
        "BE70BCE00A9C467E38B8F1FE2ED2B15B1FD659063EB09829680026C3D303B389",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005433",
        "553ACFEDF88CB67BA41AFF7AFC10E4033DA5D26CCF37C7A92368C934D8EF1A34",
        "1E894D7A75FF89A7EAD8B0DA5120F1895AD502FD00D14DC5005A6D75FE9DEF0B",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005434",
        "8C6125EA1D79EA756B3E0A13F4D9212F7263D76AAFE912F420052973B4228CDD",
        "1836D6E2DD52BD45FED4CF30472D71C2A1F19743EA5A4133F136D0BB6F6B5627",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005435",
        "38F743EAE098EEF0A1DD876AB112331C136A097509CAA088018E0FB8D168AB30",
        "6CB3B6207F9860ECB83D4CC84E2E3CFD4A1736413C57F05610E0D52F39FC5C6A",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005436",
        "6D321411C2E4603E571119B563FDB84A39C2907A5BAD35A3F5B0BB5F4E874884",
        "D0CAEE3B69D88FCAC801D6DEDB772C5A6752FA2E70B7C8462C212DB8767105B9",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005437",
        "B36D6FAB2E0A74FC883CD9409838E6488BC7B5A025BB776B987E3030B1C19BA8",
        "D7858284D60BD9FEB433FEA325F820FFD9ED4B702C500AC903B07B17375DC321",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005438",
        "4F8A6197415B349FAA6F716D6FA731C6BB28963EC7F8FB1E3B4CAD2F95B87FFE",
        "637B012FAEE8D7A75E31600215337B66680E88C87B4BEF37882F7F9CD648534B",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005439",
        "9C19D1C857C7A24BCF45DDE3DA816C9310EA064D1722CA21ED7E9FA156273F76",
        "6748E610801EE0CC435B7C7684E186095062F0B0F945797CBA1EFB9F8A33162D",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005440",
        "8AAA80C25B0D1C4F9C7ED55490DD1443516A92E78ABCDF5491BC0FFA210714DA",
        "5EB51EAB682572B08A5EA638F2389E10D3CB72407A5B8D8722E714DA20D398AE",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005441",
        "8F9A9567680E53CD4948E044079747E6CDD08F22F155813DD5B05421A84A484E",
        "7FFB2F1EF0BB440BB042069CDCBA4EA8BF85310B4F58C264A794C36648D3A413",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005442",
        "A0202EA29D108949C419E8CEC44AEDBB186EB20703DE96C30E7B81F777EF4CAA",
        "C558C3FA566D841E5A4320A8AC26B5A33D968FCA0983EE74756A4550A539F062",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005443",
        "36F3BB462B064195BFE09D4A42E062D8CCEA8F5F2F0B88D4B9475092C939C9EB",
        "5D93510DEA301705972BF4A1D354A89369131FBEE8797461A1F0EE712FA16F46",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005444",
        "03DA216B079C4F3C2A5994AAAEA1E5317F75FAC2603BE93A3E4A0D62E2173914",
        "A0B9AE3D930ED404F4297BFCD85BCA2343BA00C81FC2B52720EAE72CD1A0CD28",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005445",
        "2F05C2086D70439F241CB519E0DDDAE74665AC2E8C6E37606BE1F5866623C77A",
        "E5D726A483CF4659F3EE31E34516921BF55405ADB5891177E6D8921D4CA3A58D",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005446",
        "893EF9F7B2A14F17501A5B446A5B5D62B1B2E94ED7162717D612ED6E1CEE4165",
        "B628E67C197A2A672C558764B69D79C2F1F7ECD55EBB3C870E854D5BA861FA07",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005447",
        "8EB2D128495074CC7149E19DB2A5CA5A04960EC735A8A6FB77EBC74B2A6BBBD5",
        "E623CA71BFAA48B7E8972EAFACE6561144A5A9CEE366769C8C9E786FDC6EB407",
    },
    {
        "6864797660130609714981900799081393217269435300143305409394463459185543183397655394245057746333217197532963996371363321113864768612440380340372808892707005448",
        "891DC951F8F41D80A124A6675058D305BB97F70E3BB2C13BA8056EDD15323006",
        "428DF0068E6BB7DF3B03AA7765B45422995F9AEE89CA2CA1EC400B99A4C331EC",
    },
}