## wikitools
=========

Few tools for working with wikipedia XML dumps.

* [wikicats](https://github.com/miku/wikitools#wikicats)
* [wikinorm](https://github.com/miku/wikitools#wikinorm)
* [wikitojson](https://github.com/miku/wikitools#wikitojson)
* [wikidatatojson](https://github.com/miku/wikitools#wikidatatojson)

## Installation
------------

Commands are go-getable.

    $ go get github.com/miku/wikitools/cmd/{wikicats,wikinorm,...}

## wikicats
--------

Extract category data from a wikipedia dump.

    $ wikicats -h
      -cpuprofile="": write cpu profile to file
      -filter="^file:.*|^talk:.*|^special:.*|^wikipedia...": regex for pages to skip
      -pattern="Category": word for category, e.g. in : 'Kategorie'
      -v=false: prints current program version
      -w=4: number of workers

    $ wikicats -pattern "Kategorie" dewiki-dump.xml
    Alan Smithee    Fiktive Person
    Alan Smithee    Pseudonym
    Alan Smithee    Sammelpseudonym
    Alan Smithee    Werk von Alan Smithee
    Ang Lee Ang Lee
    Ang Lee Taiwanischer Künstler
    Ang Lee Drehbuchautor
    Ang Lee Filmregisseur
    Ang Lee Oscarpreisträger
    ...

## wikinorm
--------

Extract raw authority control data from a wikipedia dump.

    $ wikinorm -h
      -cpuprofile="": write cpu profile to file
      -filter="^file:.*|^talk:.*|^special:.*|^wikipedia...": regex for pages to skip
      -pattern="Authority Control": word for Authority Control, e.g. in : 'Normdaten'
      -v=false: prints current program version
      -w=4: number of workers

    $ wikinorm -pattern "Normdaten" dewiki-dump.xml
    Alan Smithee    {{Normdaten|TYP=p|GND=123396956|VIAF=57519787}}
    Actionfilm  {{Normdaten|TYP=s|GND=4140847-0}}
    Al Pacino   {{Normdaten|TYP=p|GND=119070243|LCCN=n/85/138113|NDL=00621247}}
    Alkohole    {{Normdaten|TYP=s|GND=4141899-2}}
    Aluminium   {{Normdaten|TYP=s|GND=4001573-7|LCCN=sh/85/003956|NDL=00560358}}
    ...

## wiki to json
----------

Convert wikipedia dump into JSON without much additional parsing.

    $ wikitojson -h
      -cpuprofile="": write cpu profile to file
      -filter="^file:.*|^talk:.*|^special:.*|^wikipedia:...": regex for pages to skip
      -v=false: prints current program version
      -w=4: number of workers

    $ wikitojson dewiki-dump.xml
    ...
    {
      "title": "Felsit",
      "ctitle": "",
      "redirect": {
        "title": ""
      },
      "text": "'''Felsite''' sind allgemein helle [[Magmatisches Gestein|magmatische]] ..."
    }
    ...

## wikidata to json
--------------

Convert wikidata dump into JSON.

    $ wikidatatojson -h
      -cpuprofile="": write cpu profile to file
      -filter="^file:.*|^talk:.*|^special:.*|^wikipedia:...": regex for pages to skip
      -v=false: prints current program version
      -w=4: number of workers

    $ wikidatatojson dewikidata-dump.xml | head -1 | jsonpp
    {
      "title": "Q15",
      "ctitle": "q15",
      "redirect": {
        "title": ""
      },
      "content": {
        "aliases": {
          "ca": [
            "continent africà"
          ],
          "en": [
            "African continent"
          ],
          "ko": [
            "아불리가"
          ],
          "pt": [
            "continente africano"
          ],
          "ta": [
            "ஆபிரிக்கா"
          ],
          "vi": [
            "Phi châu",
            "Phi"
          ]
        },
        "claims": [
          {
            "g": "Q15$2858e77f-47f8-d8c5-3fc5-759d73c5773d",
            "m": [
              "value",
              1245,
              "string",
              "71"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "Q15$7870cf68-453d-0052-5a4f-bc41dc75227a",
            "m": [
              "value",
              1151,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 7486129
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$D583DF01-F7D8-4010-9EDA-C5BBE6C5A594",
            "m": [
              "value",
              31,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 5107
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$60f93cc4-4a85-f1a0-3584-d6043bd41b80",
            "m": [
              "value",
              361,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 2
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$2457B21F-EF0B-48FF-843E-77645887E6C5",
            "m": [
              "value",
              47,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 48
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$dc19045e-49dc-396d-2f77-c5bb3ff5d736",
            "m": [
              "value",
              47,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 46
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$E9B52923-02DA-4613-9043-B14AF8FBABEA",
            "m": [
              "value",
              373,
              "string",
              "Africa"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$9FDD99A4-487C-41D9-8FD7-A0A4FD1E4A75",
            "m": [
              "value",
              18,
              "string",
              "Africa satellite orthographic.jpg"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$CC452BD9-08D8-482B-8FCE-AE8F0C5F293D",
            "m": [
              "value",
              242,
              "string",
              "Africa ethnic groups 1996.jpg"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$BF765B27-3656-48DB-A47F-15EF8E358E96",
            "m": [
              "value",
              242,
              "string",
              "Afriko-meza.png"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "q15$1EE103EB-7FF7-41BB-AC7E-634DC94F058C",
            "m": [
              "value",
              625,
              "globecoordinate",
              {
                "altitude": null,
                "globe": "http://www.wikidata.org/entity/Q2",
                "latitude": 7.1880555555556,
                "longitude": 21.093611111111,
                "precision": 0.00027777777777778
              }
            ],
            "q": [],
            "rank": 1,
            "refs": [
              [
                [
                  "value",
                  143,
                  "wikibase-entityid",
                  {
                    "entity-type": "item",
                    "numeric-id": 48183
                  }
                ]
              ]
            ]
          },
          {
            "g": "Q15$96A2A45A-C2B0-4291-A4AE-D6EB97954C90",
            "m": [
              "value",
              910,
              "wikibase-entityid",
              {
                "entity-type": "item",
                "numeric-id": 5460710
              }
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "Q15$C7613A0A-26BB-4C20-9C98-20BCD8B13E47",
            "m": [
              "value",
              948,
              "string",
              "Kenya Central banner.jpg"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "Q15$BB341468-DC35-4F9C-A76A-44B166108BE9",
            "m": [
              "value",
              349,
              "string",
              "00560082"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "Q15$63174D44-2344-4CF5-A4E8-433FA07BBC51",
            "m": [
              "value",
              214,
              "string",
              "258169097"
            ],
            "q": [],
            "rank": 1,
            "refs": [
              [
                [
                  "value",
                  143,
                  "wikibase-entityid",
                  {
                    "entity-type": "item",
                    "numeric-id": 54919
                  }
                ]
              ]
            ]
          },
          {
            "g": "Q15$171B9FB9-975A-4A2A-BFC6-34782FAC4118",
            "m": [
              "value",
              998,
              "string",
              "Regional/Africa/"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          },
          {
            "g": "Q15$28FA0A6A-6EB5-4C85-9BBA-E0190582D4FE",
            "m": [
              "value",
              935,
              "string",
              "Africa - Afrique - أفريقيا"
            ],
            "q": [],
            "rank": 1,
            "refs": [
              [
                [
                  "value",
                  143,
                  "wikibase-entityid",
                  {
                    "entity-type": "item",
                    "numeric-id": 191168
                  }
                ]
              ]
            ]
          },
          {
            "g": "Q15$947D0328-A860-4456-9017-18568DA9B48C",
            "m": [
              "value",
              646,
              "string",
              "/m/0dg3n1"
            ],
            "q": [],
            "rank": 1,
            "refs": [
              [
                [
                  "value",
                  248,
                  "wikibase-entityid",
                  {
                    "entity-type": "item",
                    "numeric-id": 15241312
                  }
                ],
                [
                  "value",
                  577,
                  "time",
                  {
                    "after": 0,
                    "before": 0,
                    "calendarmodel": "http://www.wikidata.org/entity/Q1985727",
                    "precision": 11,
                    "time": "+00000002013-10-28T00:00:00Z",
                    "timezone": 0
                  }
                ]
              ]
            ]
          },
          {
            "g": "Q15$8CA395AC-2F46-448B-A76F-4F54E4D9342E",
            "m": [
              "value",
              227,
              "string",
              "4000695-5"
            ],
            "q": [],
            "rank": 1,
            "refs": [
              [
                [
                  "value",
                  143,
                  "wikibase-entityid",
                  {
                    "entity-type": "item",
                    "numeric-id": 48183
                  }
                ]
              ]
            ]
          },
          {
            "g": "Q15$a057f02b-4b35-dd18-78cb-f6795503292e",
            "m": [
              "value",
              1036,
              "string",
              "2--6"
            ],
            "q": [],
            "rank": 1,
            "refs": []
          }
        ],
        "description": {
          "ar": "قارة",
          "be-tarask": "кантынэнт",
          "ca": "continent",
          "cs": "kontinent",
          "de": "Kontinent",
          "en": "continent",
          "en-gb": "A continent",
          "eo": "kontinento",
          "es": "continente",
          "fi": "maailman toiseksi suurin maanosa",
          "fr": "continent",
          "he": "יבשת",
          "hu": "kontinens",
          "id": "benua",
          "ilo": "kontinente",
          "it": "continente",
          "ja": "大陸",
          "ko": "아시아 다음으로 면적이 넓고 인구가 많은 대륙",
          "ksh": "Kontenänt",
          "min": "benua",
          "nb": "kontinent",
          "nl": "werelddeel",
          "pl": "kontynent",
          "pt": "continente",
          "pt-br": "continente",
          "ro": "continent",
          "ru": "второй по площади континент после Евразии,омываемый ...",
          "sv": "världsdel",
          "ta": "உலகின் இரண்டாவது மிகப்பெரிய மற்றும் அதிக ...",
          "ty": "fenua ra’ituāta’a",
          "uk": "другий за площею і населенням материк у світі, після Євразії",
          "zh": "七大洲之一",
          "zh-cn": "七大洲之一",
          "zh-hans": "七大洲之一",
          "zh-hant": "七大洲之一",
          "zh-hk": "七大洲之一",
          "zh-mo": "七大洲之一",
          "zh-my": "七大洲之一",
          "zh-sg": "七大洲之一",
          "zh-tw": "七大洲之一"
        },
        "entity": [
          "item",
          15
        ],
        "label": {
          "ace": "Afrika",
          "af": "Afrika",
          "ak": "Afrika",
          "am": "አፍሪቃ",
          "an": "Africa",
          "ang": "Affrica",
          "ar": "أفريقيا",
          "arc": "ܐܦܪܝܩܐ",
          "arz": "افريقيا",
          "as": "আফ্ৰিকা",
          "ast": "África",
          "ay": "Aphrika",
          "az": "Afrika",
          "ba": "Африка",
          "bar": "Afrika",
          "bcl": "Aprika",
          "be": "Афрыка",
          "be-tarask": "Афрыка",
          "bg": "Африка",
          "bh": "अफ़्रीका",
          "bjn": "Aprika",
          "bm": "Afrika",
          "bn": "আফ্রিকা",
          "bo": "ཨ་ཧྥེ་རི་ཁ།",
          "br": "Afrika",
          "bs": "Afrika",
          "bxr": "Африка",
          "ca": "Àfrica",
          "cdo": "Hĭ-ciŭ",
          "ce": "Африка",
          "ceb": "Aprika",
          "chr": "ᎬᎿᎦᏍᏛ",
          "chy": "Mo'hetaneho'e",
          "ckb": "ئەفریقا",
          "co": "Africa",
          "crh-latn": "Afrika",
          "cs": "Afrika",
          "csb": "Afrika",
          "cu": "Афрїка",
          "cv": "Африка",
          "cy": "Affrica",
          "da": "Afrika",
          "de": "Afrika",
          "de-ch": "Afrika",
          "diq": "Afrika",
          "dsb": "Afrika",
          "dz": "ཨཕ་རི་ཀ་",
          "el": "Αφρική",
          "en": "Africa",
          "en-ca": "Africa",
          "en-gb": "Africa",
          "eo": "Afriko",
          "es": "África",
          "et": "Aafrika",
          "eu": "Afrika",
          "ext": "África",
          "fa": "آفریقا",
          "ff": "Afirik",
          "fi": "Afrikka",
          "fit": "Aafriikka",
          "fo": "Afrika",
          "fr": "Afrique",
          "frp": "Africa",
          "frr": "Afrikoo",
          "fur": "Afriche",
          "fy": "Afrika",
          "ga": "An Afraic",
          "gag": "Afrika",
          "gan": "非洲",
          "gd": "Afraga",
          "gl": "África",
          "glk": "آفریقا",
          "gn": "Afrika",
          "got": "̰̰̻̰̳͆͂̽",
          "gsw": "Afrika",
          "gu": "આફ્રિકા",
          "gv": "Yn Affrick",
          "ha": "Afirka",
          "hak": "Fî-chû",
          "haw": "ʻApelika",
          "he": "אפריקה",
          "hi": "अफ़्रीका",
          "hif": "Africa",
          "hr": "Afrika",
          "hsb": "Afrika",
          "ht": "Afrik",
          "hu": "Afrika",
          "hy": "Աֆրիկա",
          "ia": "Africa",
          "id": "Afrika",
          "ie": "Africa",
          "ig": "Afrịka",
          "ilo": "Aprika",
          "io": "Afrika",
          "is": "Afríka",
          "it": "Africa",
          "iu": "ᐊᑉᕆᖄ",
          "ja": "アフリカ",
          "jbo": "frikytu'a",
          "jv": "Afrika",
          "ka": "აფრიკა",
          "kaa": "Afrika",
          "kab": "Tafriqt",
          "kbd": "Африкэ",
          "kg": "Afelika",
          "ki": "Abĩrika",
          "kk": "Африка",
          "kl": "Afrika",
          "km": "អាហ្វ្រិក",
          "kn": "ಆಫ್ರಿಕಾ",
          "ko": "아프리카",
          "koi": "Африка",
          "krc": "Африка",
          "ks": "اَفریٖقہ",
          "ksh": "Affrekah",
          "ku": "Afrîka",
          "kv": "Африка",
          "kw": "Afrika",
          "ky": "Африка",
          "la": "Africa",
          "lad": "Afrika",
          "lb": "Afrika",
          "lez": "Африка",
          "li": "Afrika",
          "lij": "Africa",
          "lmo": "Africa",
          "ln": "Afríka",
          "lo": "ອາຟຣິກກາ",
          "lt": "Afrika",
          "ltg": "Afrika",
          "lv": "Āfrika",
          "lzh": "阿非利加洲",
          "map-bms": "Afrika",
          "mdf": "Африкань",
          "mg": "Afrika",
          "mhr": "Африка",
          "mi": "Āwherika",
          "min": "Afrika",
          "mk": "Африка",
          "ml": "ആഫ്രിക്ക",
          "mn": "Африк",
          "mr": "आफ्रिका",
          "mrj": "Африка",
          "ms": "Afrika",
          "mt": "Afrika",
          "mwl": "África",
          "my": "အာဖရိက",
          "mzn": "آفریخا",
          "na": "Aprika",
          "nah": "Africa",
          "nan": "Hui-chiu",
          "nap": "Africa",
          "nb": "Afrika",
          "nds": "Afrika",
          "nds-nl": "Afrika",
          "ne": "अफ्रीका",
          "new": "अफ्रिका",
          "nl": "Afrika",
          "nn": "Afrika",
          "nov": "Afrika",
          "nrm": "Afrique",
          "nso": "Afrika",
          "nv": "Naakaii Łizhiní Bikéyah",
          "ny": "Africa",
          "oc": "Africa",
          "om": "Afrikaa",
          "or": "ଆଫ୍ରିକା",
          "os": "Африкæ",
          "pa": "ਅਫ਼ਰੀਕਾ",
          "pag": "Afrika",
          "pam": "Aprika",
          "pap": "Afrika",
          "pcd": "Afrike",
          "pdc": "Afrikaa",
          "pih": "Afreka",
          "pl": "Afryka",
          "pms": "Àfrica",
          "pnb": "افریقہ",
          "pnt": "Αφρικήν",
          "ps": "افريقا",
          "pt": "África",
          "pt-br": "África",
          "qu": "Aphrika",
          "rm": "Africa",
          "rmy": "Afrika",
          "rn": "Bufirika",
          "ro": "Africa",
          "roa-tara": "Afriche",
          "ru": "Африка",
          "rue": "Африка",
          "rup": "Africa",
          "rw": "Afurika",
          "sa": "आफ्रिकाखण्डः",
          "sah": "Африка",
          "sc": "Àfrica",
          "scn": "Àfrica",
          "sco": "Africae",
          "se": "Afrihkká",
          "sg": "Afrîka",
          "sgs": "Afrėka",
          "sh": "Afrika",
          "si": "අප් රිකාව",
          "sk": "Afrika",
          "sl": "Afrika",
          "sm": "Aferika",
          "sn": "Africa",
          "so": "Afrika",
          "sq": "Afrika",
          "sr": "Африка",
          "srn": "Afrka",
          "ss": "Í-Afríka",
          "st": "Afrika",
          "stq": "Afrikoa",
          "su": "Afrika",
          "sv": "Afrika",
          "sw": "Afrika",
          "szl": "Afrika",
          "ta": "ஆப்பிரிக்கா",
          "te": "ఆఫ్రికా",
          "tet": "Áfrika",
          "tg": "Африқо",
          "th": "ทวีปแอฟริกา",
          "tk": "Afrika",
          "tl": "Aprika",
          "tn": "Aferika",
          "to": "ʻAfelika",
          "tokipona": "ma Apika",
          "tpi": "Aprika",
          "tr": "Afrika",
          "ts": "Afrika",
          "tt": "Африка",
          "ty": "’Āfirita",
          "ug": "ئافرىقا",
          "uk": "Африка",
          "ur": "افریقہ",
          "uz": "Afrika",
          "ve": "Afurika",
          "vec": "Àfrica",
          "vep": "Afrik",
          "vi": "châu Phi",
          "vls": "Afrika",
          "vro": "Afriga",
          "wa": "Afrike",
          "war": "Aprika",
          "wo": "Afrig",
          "wuu": "非洲",
          "xal": "Априк",
          "xmf": "აფრიკა",
          "yi": "אפריקע",
          "yo": "Áfríkà",
          "yue": "非洲",
          "zea": "Afrika",
          "zh": "非洲",
          "zh-cn": "非洲",
          "zh-hans": "非洲",
          "zh-hant": "非洲",
          "zh-hk": "非洲",
          "zh-mo": "非洲",
          "zh-my": "非洲",
          "zh-sg": "非洲",
          "zh-tw": "非洲",
          "zu": "IAfrika"
        },
        "links": {
          "acewiki": {
            "badges": [],
            "name": "Afrika"
          },
          "afwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "akwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "alswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "amwiki": {
            "badges": [],
            "name": "አፍሪቃ"
          },
          "angwiki": {
            "badges": [],
            "name": "Affrica"
          },
          "anwiki": {
            "badges": [],
            "name": "Africa"
          },
          "arcwiki": {
            "badges": [],
            "name": "ܐܦܪܝܩܐ"
          },
          "arwiki": {
            "badges": [],
            "name": "أفريقيا"
          },
          "arzwiki": {
            "badges": [],
            "name": "افريقيا"
          },
          "astwiki": {
            "badges": [],
            "name": "África"
          },
          "aswiki": {
            "badges": [],
            "name": "আফ্ৰিকা"
          },
          "aywiki": {
            "badges": [],
            "name": "Aphrika"
          },
          "azwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "barwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "bat_smgwiki": {
            "badges": [],
            "name": "Afrėka"
          },
          "bawiki": {
            "badges": [],
            "name": "Африка"
          },
          "bclwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "be_x_oldwiki": {
            "badges": [],
            "name": "Афрыка"
          },
          "bewiki": {
            "badges": [],
            "name": "Афрыка"
          },
          "bgwiki": {
            "badges": [],
            "name": "Африка"
          },
          "bhwiki": {
            "badges": [],
            "name": "अफ़्रीका"
          },
          "bjnwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "bmwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "bnwiki": {
            "badges": [],
            "name": "আফ্রিকা"
          },
          "bowiki": {
            "badges": [],
            "name": "ཨ་ཧྥེ་རི་ཁ།"
          },
          "brwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "bswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "bxrwiki": {
            "badges": [],
            "name": "Африка"
          },
          "cawiki": {
            "badges": [],
            "name": "Àfrica"
          },
          "cdowiki": {
            "badges": [],
            "name": "Hĭ-ciŭ"
          },
          "cebwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "cewiki": {
            "badges": [],
            "name": "Африка"
          },
          "chrwiki": {
            "badges": [],
            "name": "ᎬᎿᎦᏍᏛ"
          },
          "chywiki": {
            "badges": [],
            "name": "Mo'hetaneho'e"
          },
          "ckbwiki": {
            "badges": [],
            "name": "ئەفریقا"
          },
          "commonswiki": {
            "badges": [],
            "name": "Africa - Afrique - أفريقيا"
          },
          "cowiki": {
            "badges": [],
            "name": "Africa"
          },
          "crhwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "csbwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "cswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "cuwiki": {
            "badges": [],
            "name": "Афрїка"
          },
          "cvwiki": {
            "badges": [],
            "name": "Африка"
          },
          "cywiki": {
            "badges": [],
            "name": "Affrica"
          },
          "dawiki": {
            "badges": [],
            "name": "Afrika"
          },
          "dewiki": {
            "badges": [],
            "name": "Afrika"
          },
          "dewikiquote": {
            "badges": [],
            "name": "Afrika"
          },
          "dewikivoyage": {
            "badges": [],
            "name": "Afrika"
          },
          "diqwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "dsbwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "dzwiki": {
            "badges": [],
            "name": "ཨཕ་རི་ཀ་"
          },
          "elwiki": {
            "badges": [],
            "name": "Αφρική"
          },
          "elwikivoyage": {
            "badges": [],
            "name": "Αφρική"
          },
          "enwiki": {
            "badges": [],
            "name": "Africa"
          },
          "enwikiquote": {
            "badges": [],
            "name": "Africa"
          },
          "enwikivoyage": {
            "badges": [],
            "name": "Africa"
          },
          "eowiki": {
            "badges": [],
            "name": "Afriko"
          },
          "eswiki": {
            "badges": [],
            "name": "África"
          },
          "eswikiquote": {
            "badges": [],
            "name": "África"
          },
          "eswikivoyage": {
            "badges": [],
            "name": "África"
          },
          "etwiki": {
            "badges": [],
            "name": "Aafrika"
          },
          "euwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "extwiki": {
            "badges": [],
            "name": "África"
          },
          "fawiki": {
            "badges": [],
            "name": "آفریقا"
          },
          "ffwiki": {
            "badges": [],
            "name": "Afirik"
          },
          "fiu_vrowiki": {
            "badges": [],
            "name": "Afriga"
          },
          "fiwiki": {
            "badges": [],
            "name": "Afrikka"
          },
          "fowiki": {
            "badges": [],
            "name": "Afrika"
          },
          "frpwiki": {
            "badges": [],
            "name": "Africa"
          },
          "frrwiki": {
            "badges": [],
            "name": "Afrikoo"
          },
          "frwiki": {
            "badges": [],
            "name": "Afrique"
          },
          "frwikiquote": {
            "badges": [],
            "name": "Afrique"
          },
          "frwikivoyage": {
            "badges": [],
            "name": "Afrique"
          },
          "furwiki": {
            "badges": [],
            "name": "Afriche"
          },
          "fywiki": {
            "badges": [],
            "name": "Afrika"
          },
          "gagwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "ganwiki": {
            "badges": [],
            "name": "非洲"
          },
          "gawiki": {
            "badges": [],
            "name": "An Afraic"
          },
          "gdwiki": {
            "badges": [],
            "name": "Afraga"
          },
          "glkwiki": {
            "badges": [],
            "name": "آفریقا"
          },
          "glwiki": {
            "badges": [],
            "name": "África"
          },
          "gnwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "gotwiki": {
            "badges": [],
            "name": "̰̰̻̰̳͆͂̽"
          },
          "guwiki": {
            "badges": [],
            "name": "આફ્રિકા"
          },
          "gvwiki": {
            "badges": [],
            "name": "Yn Affrick"
          },
          "hakwiki": {
            "badges": [],
            "name": "Fî-chû"
          },
          "hawiki": {
            "badges": [],
            "name": "Afirka"
          },
          "hawwiki": {
            "badges": [],
            "name": "ʻApelika"
          },
          "hewiki": {
            "badges": [],
            "name": "אפריקה"
          },
          "hewikivoyage": {
            "badges": [],
            "name": "אפריקה"
          },
          "hifwiki": {
            "badges": [],
            "name": "Africa"
          },
          "hiwiki": {
            "badges": [
              "Q17437796"
            ],
            "name": "अफ़्रीका"
          },
          "hrwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "hsbwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "htwiki": {
            "badges": [],
            "name": "Afrik"
          },
          "huwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "hywiki": {
            "badges": [],
            "name": "Աֆրիկա"
          },
          "iawiki": {
            "badges": [],
            "name": "Africa"
          },
          "idwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "iewiki": {
            "badges": [],
            "name": "Africa"
          },
          "igwiki": {
            "badges": [],
            "name": "Afrịka"
          },
          "ilowiki": {
            "badges": [],
            "name": "Aprika"
          },
          "iowiki": {
            "badges": [],
            "name": "Afrika"
          },
          "iswiki": {
            "badges": [],
            "name": "Afríka"
          },
          "itwiki": {
            "badges": [],
            "name": "Africa"
          },
          "itwikiquote": {
            "badges": [],
            "name": "Africa"
          },
          "itwikivoyage": {
            "badges": [],
            "name": "Africa"
          },
          "iuwiki": {
            "badges": [],
            "name": "ᐊᑉᕆᖄ"
          },
          "jawiki": {
            "badges": [
              "Q17437798"
            ],
            "name": "アフリカ"
          },
          "jbowiki": {
            "badges": [],
            "name": "frikytu'a"
          },
          "jvwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "kaawiki": {
            "badges": [],
            "name": "Afrika"
          },
          "kabwiki": {
            "badges": [],
            "name": "Tafriqt"
          },
          "kawiki": {
            "badges": [
              "Q17437796"
            ],
            "name": "აფრიკა"
          },
          "kbdwiki": {
            "badges": [],
            "name": "Африкэ"
          },
          "kgwiki": {
            "badges": [],
            "name": "Afelika"
          },
          "kiwiki": {
            "badges": [],
            "name": "Abĩrika"
          },
          "kkwiki": {
            "badges": [],
            "name": "Африка"
          },
          "klwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "kmwiki": {
            "badges": [],
            "name": "អាហ្វ្រិក"
          },
          "knwiki": {
            "badges": [],
            "name": "ಆಫ್ರಿಕಾ"
          },
          "koiwiki": {
            "badges": [],
            "name": "Африка"
          },
          "kowiki": {
            "badges": [],
            "name": "아프리카"
          },
          "krcwiki": {
            "badges": [
              "Q17437796"
            ],
            "name": "Африка"
          },
          "kswiki": {
            "badges": [],
            "name": "اَفریٖقہ"
          },
          "kuwiki": {
            "badges": [],
            "name": "Afrîka"
          },
          "kvwiki": {
            "badges": [],
            "name": "Африка"
          },
          "kwwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "kywiki": {
            "badges": [],
            "name": "Африка"
          },
          "ladwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "lawiki": {
            "badges": [],
            "name": "Africa"
          },
          "lbwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "lezwiki": {
            "badges": [],
            "name": "Африка"
          },
          "lijwiki": {
            "badges": [],
            "name": "Africa"
          },
          "liwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "lmowiki": {
            "badges": [],
            "name": "Africa"
          },
          "lnwiki": {
            "badges": [],
            "name": "Afríka"
          },
          "lowiki": {
            "badges": [],
            "name": "ທະວີບອາຟະລິກ"
          },
          "ltgwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "ltwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "lvwiki": {
            "badges": [],
            "name": "Āfrika"
          },
          "map_bmswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "mdfwiki": {
            "badges": [],
            "name": "Африк"
          },
          "mgwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "mhrwiki": {
            "badges": [],
            "name": "Африке"
          },
          "miwiki": {
            "badges": [],
            "name": "Āwherika"
          },
          "mkwiki": {
            "badges": [],
            "name": "Африка"
          },
          "mlwiki": {
            "badges": [],
            "name": "ആഫ്രിക്ക"
          },
          "mnwiki": {
            "badges": [],
            "name": "Африк"
          },
          "mrjwiki": {
            "badges": [],
            "name": "Африка"
          },
          "mrwiki": {
            "badges": [],
            "name": "आफ्रिका"
          },
          "mswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "mtwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "mwlwiki": {
            "badges": [],
            "name": "África"
          },
          "mywiki": {
            "badges": [],
            "name": "အာဖရိက"
          },
          "mznwiki": {
            "badges": [],
            "name": "آفریخا"
          },
          "nahwiki": {
            "badges": [],
            "name": "Africa"
          },
          "napwiki": {
            "badges": [],
            "name": "Africa"
          },
          "nawiki": {
            "badges": [],
            "name": "Aprika"
          },
          "nds_nlwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "ndswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "newiki": {
            "badges": [],
            "name": "अफ्रीका"
          },
          "newwiki": {
            "badges": [],
            "name": "अफ्रिका"
          },
          "nlwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "nlwikivoyage": {
            "badges": [],
            "name": "Afrika"
          },
          "nnwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "nnwikiquote": {
            "badges": [],
            "name": "Afrika"
          },
          "novwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "nowiki": {
            "badges": [],
            "name": "Afrika"
          },
          "nrmwiki": {
            "badges": [],
            "name": "Afrique"
          },
          "nsowiki": {
            "badges": [],
            "name": "Afrika"
          },
          "nvwiki": {
            "badges": [],
            "name": "Naakaii Łizhiní Bikéyah"
          },
          "nywiki": {
            "badges": [],
            "name": "Africa"
          },
          "ocwiki": {
            "badges": [],
            "name": "Africa"
          },
          "omwiki": {
            "badges": [],
            "name": "Afrikaa"
          },
          "orwiki": {
            "badges": [],
            "name": "ଆଫ୍ରିକା"
          },
          "oswiki": {
            "badges": [],
            "name": "Африкæ"
          },
          "pagwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "pamwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "papwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "pawiki": {
            "badges": [],
            "name": "ਅਫ਼ਰੀਕਾ"
          },
          "pcdwiki": {
            "badges": [],
            "name": "Afrike"
          },
          "pdcwiki": {
            "badges": [],
            "name": "Afrikaa"
          },
          "pihwiki": {
            "badges": [],
            "name": "Afreka"
          },
          "plwiki": {
            "badges": [],
            "name": "Afryka"
          },
          "plwikiquote": {
            "badges": [],
            "name": "Afryka"
          },
          "plwikivoyage": {
            "badges": [],
            "name": "Afryka"
          },
          "pmswiki": {
            "badges": [],
            "name": "Àfrica"
          },
          "pnbwiki": {
            "badges": [],
            "name": "افریقہ"
          },
          "pntwiki": {
            "badges": [],
            "name": "Αφρικήν"
          },
          "pswiki": {
            "badges": [],
            "name": "افريقا"
          },
          "ptwiki": {
            "badges": [],
            "name": "África"
          },
          "ptwikivoyage": {
            "badges": [],
            "name": "África"
          },
          "quwiki": {
            "badges": [],
            "name": "Aphrika"
          },
          "rmwiki": {
            "badges": [],
            "name": "Africa"
          },
          "rmywiki": {
            "badges": [],
            "name": "Afrika"
          },
          "rnwiki": {
            "badges": [],
            "name": "Bufirika"
          },
          "roa_rupwiki": {
            "badges": [],
            "name": "Africa"
          },
          "roa_tarawiki": {
            "badges": [],
            "name": "Afriche"
          },
          "rowiki": {
            "badges": [],
            "name": "Africa"
          },
          "rowikivoyage": {
            "badges": [],
            "name": "Africa"
          },
          "ruewiki": {
            "badges": [],
            "name": "Африка"
          },
          "ruwiki": {
            "badges": [],
            "name": "Африка"
          },
          "ruwikiquote": {
            "badges": [],
            "name": "Африка"
          },
          "ruwikivoyage": {
            "badges": [],
            "name": "Африка"
          },
          "rwwiki": {
            "badges": [],
            "name": "Afurika"
          },
          "sahwiki": {
            "badges": [],
            "name": "Африка"
          },
          "sawiki": {
            "badges": [],
            "name": "आफ्रिकाखण्डः"
          },
          "scnwiki": {
            "badges": [],
            "name": "Àfrica"
          },
          "scowiki": {
            "badges": [],
            "name": "Africae"
          },
          "scwiki": {
            "badges": [],
            "name": "Àfrica"
          },
          "sewiki": {
            "badges": [],
            "name": "Afrihkká"
          },
          "sgwiki": {
            "badges": [],
            "name": "Afrîka"
          },
          "shwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "simplewiki": {
            "badges": [],
            "name": "Africa"
          },
          "siwiki": {
            "badges": [],
            "name": "අප්‍රිකාව"
          },
          "skwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "skwikiquote": {
            "badges": [],
            "name": "Afrika"
          },
          "slwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "smwiki": {
            "badges": [],
            "name": "Aferika"
          },
          "snwiki": {
            "badges": [],
            "name": "Africa"
          },
          "sowiki": {
            "badges": [],
            "name": "Afrika"
          },
          "sqwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "srnwiki": {
            "badges": [],
            "name": "Afrka"
          },
          "srwiki": {
            "badges": [],
            "name": "Африка"
          },
          "sswiki": {
            "badges": [],
            "name": "Í-Afríka"
          },
          "stqwiki": {
            "badges": [],
            "name": "Afrikoa"
          },
          "stwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "suwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "svwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "svwikivoyage": {
            "badges": [],
            "name": "Afrika"
          },
          "swwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "szlwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "tawiki": {
            "badges": [],
            "name": "ஆப்பிரிக்கா"
          },
          "tetwiki": {
            "badges": [],
            "name": "Áfrika"
          },
          "tewiki": {
            "badges": [],
            "name": "ఆఫ్రికా"
          },
          "tgwiki": {
            "badges": [],
            "name": "Африқо"
          },
          "thwiki": {
            "badges": [],
            "name": "ทวีปแอฟริกา"
          },
          "tkwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "tlwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "tnwiki": {
            "badges": [],
            "name": "Aferika"
          },
          "towiki": {
            "badges": [],
            "name": "ʻAfelika"
          },
          "tpiwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "trwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "tswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "ttwiki": {
            "badges": [],
            "name": "Африка"
          },
          "ugwiki": {
            "badges": [],
            "name": "ئافرىقا"
          },
          "ukwiki": {
            "badges": [],
            "name": "Африка"
          },
          "ukwikiquote": {
            "badges": [],
            "name": "Африка"
          },
          "ukwikivoyage": {
            "badges": [],
            "name": "Африка"
          },
          "urwiki": {
            "badges": [],
            "name": "افریقہ"
          },
          "uzwiki": {
            "badges": [],
            "name": "Afrika"
          },
          "vecwiki": {
            "badges": [],
            "name": "Àfrica"
          },
          "vepwiki": {
            "badges": [],
            "name": "Afrik"
          },
          "vewiki": {
            "badges": [],
            "name": "Afurika"
          },
          "viwiki": {
            "badges": [],
            "name": "Châu Phi"
          },
          "viwikivoyage": {
            "badges": [],
            "name": "Châu Phi"
          },
          "vlswiki": {
            "badges": [],
            "name": "Afrika"
          },
          "warwiki": {
            "badges": [],
            "name": "Aprika"
          },
          "wawiki": {
            "badges": [],
            "name": "Afrike"
          },
          "wowiki": {
            "badges": [],
            "name": "Afrig"
          },
          "wuuwiki": {
            "badges": [],
            "name": "非洲"
          },
          "xalwiki": {
            "badges": [],
            "name": "Априк"
          },
          "xmfwiki": {
            "badges": [],
            "name": "აფრიკა"
          },
          "yiwiki": {
            "badges": [],
            "name": "אפריקע"
          },
          "yowiki": {
            "badges": [
              "Q17437796"
            ],
            "name": "Áfríkà"
          },
          "zeawiki": {
            "badges": [],
            "name": "Afrika"
          },
          "zh_classicalwiki": {
            "badges": [],
            "name": "阿非利加洲"
          },
          "zh_min_nanwiki": {
            "badges": [],
            "name": "Hui-chiu"
          },
          "zh_yuewiki": {
            "badges": [],
            "name": "非洲"
          },
          "zhwiki": {
            "badges": [],
            "name": "非洲"
          },
          "zhwikivoyage": {
            "badges": [],
            "name": "非洲"
          },
          "zuwiki": {
            "badges": [],
            "name": "IAfrika"
          }
        }
      }
    }
    
    
   #### THANK YOU
