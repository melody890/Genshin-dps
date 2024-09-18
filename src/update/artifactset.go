package update

//func UpdateArtifactSetStatus(cfg character2.Charcfg) [4][attributes.EndStatType]float64 {
//	var states [4][attributes.EndStatType]float64
//	//var state [attributes.EndStatType]float64
//	for _, row := range cfg.Artifacts {
//		countMap := make(map[string]int)
//		NowArtifact := noblesse.Mistsplitter{}
//		// 统计每个元素的出现次数
//		for _, item := range row {
//			countMap[item]++
//		}
//
//		// 输出出现次数为 2、3 或 4 以上的元素
//		for item, count := range countMap {
//			if item != "" && count >= 2 {
//				NowArtifact.ApplyTwoSetStates(cfg)
//			}
//		}
//	}
//	os.Exit(0)
//	return states
//}
