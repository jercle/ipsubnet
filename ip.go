package ipsubnet

import "strconv"

type Ip struct {
	ip          string
	networkSize uint
	subnet_mask uint
}

func SubnetCalculator(ip string, networkSize uint) *Ip {

	s := &Ip{
		ip:          ip,
		networkSize: networkSize,
		subnet_mask: 0xFFFFFFFF << uint(32-networkSize),
	}

	return s
}

func convertQuardsToInt(splits []string) []int {
	quardsInt := []int{}

	for _, quard := range splits {
		j, err := strconv.Atoi(quard)
		if err != nil {
			panic(err)
		}
		quardsInt = append(quardsInt, j)
	}

	return quardsInt
}
