package versions

type Version struct {
	Minecraft string
	ID		  string
	Protocol  int
}

var Registry = []Version{
	{Minecraft: "1.21.2", ID: "v1212", Protocol: 686},
	{Minecraft: "1.21.20", ID: "v12120", Protocol: 712},
	{Minecraft: "1.21.30", ID: "v12130", Protocol: 729},
	{Minecraft: "1.21.40", ID: "v12140", Protocol: 748},
	{Minecraft: "1.21.50", ID: "v12150", Protocol: 766},
	{Minecraft: "1.21.60", ID: "v12160", Protocol: 776},
	{Minecraft: "1.21.70", ID: "v12170", Protocol: 786},
	{Minecraft: "1.21.80", ID: "v12180", Protocol: 800},
	{Minecraft: "1.21.90", ID: "v12190", Protocol: 818},
	{Minecraft: "1.21.93", ID: "v12193", Protocol: 819},
	{Minecraft: "1.21.100", ID: "v121100", Protocol: 827},
	{Minecraft: "1.21.111", ID: "v121111", Protocol: 844},
	{Minecraft: "1.21.120", ID: "v121120", Protocol: 859},
	{Minecraft: "1.21.130", ID: "v121130", Protocol: 897},
	{Minecraft: "1.26.10", ID: "v12610", Protocol: 944},
	{Minecraft: "1.26.30", ID: "v12630", Protocol: 1001},
	{Minecraft: "1.26.40", ID: "v12640", Protocol: 2168},
}

var (
	ByMinecraft = make(map[string]Version)
	ByID        = make(map[string]Version)
	ByProtocol  = make(map[int][]Version)
)

func init() {
	for _, v := range Registry {
		ByMinecraft[v.Minecraft] = v
		ByID[v.ID] = v
		ByProtocol[v.Protocol] = append(ByProtocol[v.Protocol], v)
	}
}
