package structures

type CourseInfo struct {
	Name string
	URL  string
}

var CourseMap = map[string][]CourseInfo{
	"ИПТИП": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/064/0klu1q7g97pj4gqzmf3ygq9yali343o6/IPTIP_1-kurs_vesna_09.04.24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/2cc/08mm2ctolr4kh8u78kb2akgfqkczc0im/IPTIP_2-kurs_vesna_25.03.24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/68c/s5szzvc5rk30899ks9p5pa1i1aqcil5z/IPTIP_3-kurs_vesna_01.04.24.xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/8ab/r84zhklvwc7fwbqmhwomebkz0e8xqsac/IPTIP_4-kurs_vesna_11.03.24.xlsx"},
		{"5 Курс", "https://webservices.mirea.ru/upload/iblock/e85/ccv8mcfkpk4yiyurk1tb01ewm55ie1kt/IPTIP_5-kurs_vesna_01.04.24.xlsx"},
	},
	"ИТУ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/d88/8lbnze9xuhhb9k2u4nhoq8eowhxs2662/ITU_1-kurs_vesna_19.04.24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/a31/yjms9j8n0uz9h9bk69fg52w69mar7bn5/ITU_2-kurs_vesna_19.04.24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/e3a/4e8mchlxn9epaqq93iufk69bvuicffom/ITU_3-kurs_vesna_12.04.24.xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/6f1/ryg1dqbixzh4jbgo88py9rjfp9u9q3cl/ITU_4-kurs_vesna_29.03.24.xlsx"},
	},
	"ИИТ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/d1e/t6rsp67gqg3bw3m0goxijumni76fdshd/IIT_1-kurs_vesna_19.04.24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/d17/qt3qybpiqptjy08ea98iixpxd6bflkkg/IIT_2-kurs_vesna_15.03.24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/d21/u93ld61vc9t9s5l7wu4gwm4yxeg15yem/IIT_3-kurs_vesna_15.04.24.xlsx"},
	},
	"ИИИ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/89d/117rl6cfvngdsuhpnhysjio2x3v4gvp1/III_1-kurs_vesna_16.04.24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/e07/on9970ee91tnnvgsl3za5624f1pyftgm/III_2-kurs_vesna_16.04.24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/4c7/vwotny4vb288h1epjl6m84ng9tmc2op3/III_3-kurs_vesna_16.04.24.xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/cb7/yjsehvb6xso57khwmg5542z4l4g73d64/III_4-kurs_vesna_08.04.24.xlsx"},
		{"5 Курс", "https://webservices.mirea.ru/upload/iblock/c85/rnhf4tajxjmt0a65jxbey0qom2tovzns/III_5-kurs_vesna_11.03.24.xlsx"},
	},
	"ИКБ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/ccf/lhiypx63fzewr1h8rtd28hn36na1o264/IKTST_1_k_vesna_23_24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/8cb/2gw5xaud9mt9c9i4dabeirjhec6g5p75/IKTST_2_k_vesna_23_24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/c9c/80ljm36r304aw8f1uv46wc2es9tqaxzp/IKTST_3_k_vesna_23_24.xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/ac3/gr2y7u2z7l5w2v4y0df96icv8qzxxhzj/IKTST_4_k_vesna_23_24.xlsx"},
		{"5 Курс", "https://webservices.mirea.ru/upload/iblock/57f/ik20y2hflyeyp169iomh9b8cogfbkelo/IKTST_5_k_vesna_23_24.xlsx"},
	},
	"ИРИ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/639/jgfi576kuaqtqegwgzevjp5u8dgsb96j/IRI_1-kurs_vesna_10.04.24.xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/cfa/1q3xa4hcvwsrg2ucjurdepyo1anpd36x/IRI_2-kurs_vesna_09.04.24.xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/f3d/f3y1zgkfv9in3wl5kef0h7zv945z4130/IRI_3-kurs_vesna_11.04.24.xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/e20/lbw3ezlhjg9wag8dvcoamswvm83rn4sh/IRI_4-kurs_vesna_10.04.24.xlsx"},
		{"5 Курс", "https://webservices.mirea.ru/upload/iblock/6d3/047bq0p9wzet5tcj89faohtyhnkhwezf/IRI_5-kurs_vesna_08.04.24.xlsx"},
	},
	"ИТХТ": {
		{"1 Курс", "https://webservices.mirea.ru/upload/iblock/a7d/v3y2nm6c01jyix5pf1wcz1ne0xar41vf/ITKHT_bak_1-kurs_vesna_23_24-uch.g..xlsx"},
		{"2 Курс", "https://webservices.mirea.ru/upload/iblock/556/7ghw8ta7nujz4jtkg7pl16w5c9weanqq/ITKHT_bak_2_kurs_vesna_23_24_uch_g..xlsx"},
		{"3 Курс", "https://webservices.mirea.ru/upload/iblock/f3e/stftalolcu1x0trd586067x5864hwovz/ITKHT_bak_3_kurs_vesna_23_24-uch.g..xlsx"},
		{"4 Курс", "https://webservices.mirea.ru/upload/iblock/b48/khwuazpe6wj5v3fn7b2ztpals9a02wjg/ITKHT_bak_4_kurs_vesna_23_24-uch.g..xlsx"},
	},
}
