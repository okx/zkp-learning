package ntt

const (
	p = 998244353 //2^23*7*17
)

var (
	pMinus2Arr = []int{3, 3, 3, 13, 29, 281, 349} //998244351=3^3*13*29*281*349
	inverse2   = modInverse(2)
)

func NTT(xn []int) []int {
	N := len(xn)
	gnn1 := GNN(N)
	gnn := make([]int, N)
	for i := 0; i < N; i++ {
		gnn[i] = 1
	}

	XN := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			XN[i] = add(XN[i], multiply(xn[j], gnn[j]))
			gnn[j] = multiply(gnn[j], gnn1[j])
		}
	}
	return XN
}

func INTT(XN []int) []int {
	N := len(XN)
	gnn1Inverse := GNNInverse(N)
	gnn := make([]int, N)
	for i := 0; i < N; i++ {
		gnn[i] = 1
	}

	NInverse := modInverse(N)
	xn := make([]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			xn[i] = add(xn[i], multiply(XN[j], gnn[j]))
			gnn[j] = multiply(gnn[j], gnn1Inverse[j])
		}
		xn[i] = multiply(xn[i], NInverse)
	}

	return xn
}

func FNTT(xn []int) []int {
	return nttInterval(xn, 0, 1)
}

func IFNTT(XN []int) []int {
	return inttInterval(XN, 0, 1)
}

func nttInterval(xn []int, start, interval int) []int {
	if len(xn) < interval*2 {
		panic("xn length should be 2^N")
	}
	if len(xn) == interval*2 {
		XN := make([]int, 2)
		XN[0] = (xn[start] + multiply(xn[start+interval], 1)) % p
		XN[1] = (xn[start] + multiply(xn[start+interval], p-1)) % p
		return XN
	}

	evenSubXN := nttInterval(xn, start, interval*2)
	oddSubXN := nttInterval(xn, start+interval, interval*2)

	N := len(xn) / interval
	gnn1 := GNN(N)
	XN := make([]int, N)
	for i := range evenSubXN {
		XN[i] = add(evenSubXN[i], multiply(gnn1[i], oddSubXN[i]))
		XN[i+N/2] = subtract(evenSubXN[i], multiply(gnn1[i], oddSubXN[i]))
	}
	return XN
}

func inttInterval(XN []int, start, interval int) []int {
	if len(XN) < interval*2 {
		panic("XN length should be 2^N")
	}
	if len(XN) == interval*2 {
		gnn1Inverse := GNNInverse(2)
		xn := make([]int, 2)
		xn[0] = add(XN[start], multiply(XN[start+interval], gnn1Inverse[0]))
		xn[0] = multiply(xn[0], inverse2)
		xn[1] = add(XN[start], multiply(XN[start+interval], gnn1Inverse[1]))
		xn[1] = multiply(xn[1], inverse2)

		return xn
	}

	evenSubxn := inttInterval(XN, start, interval*2)
	oddSubxn := inttInterval(XN, start+interval, interval*2)

	N := len(XN) / interval
	gnn1Inverse := GNNInverse(N)
	xn := make([]int, N)
	for i := range evenSubxn {
		xn[i] = add(evenSubxn[i], multiply(gnn1Inverse[i], oddSubxn[i]))
		xn[i] = divide(xn[i], 2)
		xn[i+N/2] = subtract(evenSubxn[i], multiply(gnn1Inverse[i], oddSubxn[i]))
		xn[i+N/2] = divide(xn[i+N/2], 2)
	}
	return xn
}

func GNN(N int) []int {
	gn := GN(N)
	gns := make([]int, N)
	gns[0] = 1
	for i := 1; i < N; i++ {
		gns[i] = multiply(gns[i-1], gn)
	}
	return gns
}

func GNNInverse(N int) []int {
	gn := GN(N)
	gnInverse := modInverse(gn)
	gns := make([]int, N)
	gns[0] = 1
	for i := 1; i < N; i++ {
		gns[i] = multiply(gns[i-1], gnInverse)
	}
	return gns
}

func GN(N int) int {
	n := 1
	logN := 0
	for i := 0; ; i++ {
		n *= 2
		logN++
		if n == N {
			break
		}
		if n > N {
			panic("N should be 2^n")
		}
	}
	if logN > 22 {
		panic("N is too big")
	}
	g := 3
	for i := 0; i < 23-logN; i++ {
		g = multiply(g, g)
	}
	g = pow(g, 7)
	g = pow(g, 17)
	return g
}

func add(a, b int) int {
	return (a + b) % p
}

func subtract(a, b int) int {
	return add(a, p-b)
}

func multiply(a, b int) int {
	return (a * b) % p
}

func divide(a, b int) int {
	return multiply(a, modInverse(b))
}

func pow(base, exp int) int {
	n := 1
	for i := 0; i < exp; i++ {
		n = multiply(n, base)
	}
	return n
}

func modInverse(element int) int {
	n := element
	for _, exp := range pMinus2Arr {
		n = pow(n, exp)
	}
	return n
}
