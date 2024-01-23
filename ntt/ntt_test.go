package ntt

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestModInverse(t *testing.T) {
	for i := 1; i < 15; i++ {
		res := modInverse(i)
		require.Equal(t, 1, res*i%p)
	}
}

func TestCreateGN(t *testing.T) {
	N := 8
	gn1 := GN(N)
	gn := gn1
	for i := 0; i < N; i++ {
		fmt.Println(gn)
		gn = gn * gn1 % p
	}
	require.Equal(t, gn1, gn)
}

func TestNTT(t *testing.T) {
	const N = 32
	xn := make([]int, 0, N)
	for i := 1; i <= N; i++ {
		xn = append(xn, i)
	}
	XN := NTT4Step(xn, 4)
	fmt.Println("XN:", XN)

	XN2 := NTT4StepOpt(xn, 4)
	require.Equal(t, XN, XN2)
	//[36 894301004 346334868 201631260 998244349 796613085 651909477 103943341]
	require.Equal(t, xn, INTT(XN))
}

func TestINTT(t *testing.T) {
	XN := []int{1, 0, p - 1, 0}
	xn := INTT(XN)
	fmt.Println(xn)
	require.Equal(t, XN, NTT(xn))
}

func TestFNTT(t *testing.T) {
	xn := []int{1, 2, 3, 8, 5, 6, 7, 8}
	require.Equal(t, NTT(xn), FNTT(xn))
}

func TestIFNTT(t *testing.T) {
	XN := []int{1, 2, 3, 4, 5, 6, 7, 8}
	require.Equal(t, INTT(XN), IFNTT(XN))
}

const benchN = 1 << 15

// time cost: 12016409125ns
func BenchmarkNTT(b *testing.B) {
	xn := make([]int, benchN)
	for i := 0; i < len(xn); i++ {
		xn[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		INTT(NTT(xn))
	}
}

// time cost: 1425133167ns
func BenchmarkFNTT(b *testing.B) {
	xn := make([]int, benchN)
	for i := 0; i < len(xn); i++ {
		xn[i] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IFNTT(FNTT(xn))
	}
}

func TestRearrage(t *testing.T) {
	xn := []int{1, 2, 3, 4, 5, 6, 7, 8}
	xn2 := rearrage(xn, 2, 4)
	fmt.Println(xn2)

	xn3 := transpose(xn2)
	fmt.Println(xn3)
}

func TestTransposeInSlice(t *testing.T) {
	xn := []int{1, 2, 3, 4, 5, 6, 7, 8}
	xn2 := transposeInSlice(xn, 4)
	fmt.Println(xn2)
}
