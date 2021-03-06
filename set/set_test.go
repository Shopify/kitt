package set

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	resultSet64 Set[uint64]
	resultBool  bool
)

func ExampleFromSlice() {
	s := FromSlice([]int{1, 2, 3, 1, 2, 3})
	// need to check for each individually since set is unordered, can't rely on Values()
	fmt.Println(s.Has(1), s.Has(2), s.Has(3), s.Has(4))
	fmt.Println(s.Len())
	// Output:
	// true true true false
	// 3
}

func ExampleSet_Add() {
	s := FromSlice([]int{1, 2, 3})
	s.Add(3, 4, 5)
	fmt.Println(s.Len())
	s.Add(3)
	fmt.Println(s.Len())
	s.Add(6)
	fmt.Println(s.Len())
	// Output:
	// 5
	// 5
	// 6
}

func ExampleSet_Delete() {
	s := FromSlice([]int{1, 2, 3})
	fmt.Println(s.Len(), s.Has(3))
	s.Delete(3)
	fmt.Println(s.Len(), s.Has(3))
	// Output:
	// 3 true
	// 2 false
}

func ExampleSet_Has() {
	s := FromSlice([]int{1, 2, 3})
	fmt.Println(s.Has(1), s.Has(2), s.Has(3), s.Has(4))
	// Output: true true true false
}

func ExampleSet_Has_empty() {
	s := FromSlice([]int{})
	fmt.Println(s.Has(1), s.Has(2), s.Has(3), s.Has(4))
	// Output: false false false false
}

func ExampleSet_Len() {
	s := FromSlice([]int{1, 2, 3})
	fmt.Println(s.Len())
	// Output: 3
}

func ExampleSet_ForEach() {
	s := FromSlice([]int{1, 2, 3})
	out := make(Set[int], 0)
	s.ForEach(
		func(i int) {
			out.Add(i * 2)
		},
	)
	fmt.Println(out.Has(1), out.Has(2), out.Has(4), out.Has(6))
	// Output: false true true true
}

func ExampleSet_Values() {
	s := FromSlice([]int{1, 2, 3})
	v := s.Values()
	// note, Set is unordered so we have to sort the values first before testing against them
	sort.Ints(v)
	fmt.Println(v)
	// Output: [1 2 3]
}

func ExampleSet_Clone() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := s1.Clone()
	v1 := s1.Values()
	sort.Ints(v1)
	v2 := s2.Values()
	sort.Ints(v2)
	fmt.Println(v1, v2)

	s1.Add(999)
	v1 = s1.Values()
	sort.Ints(v1)
	fmt.Println(v1, v2)

	// Output:
	// [1 2 3] [1 2 3]
	// [1 2 3 999] [1 2 3]
}

func ExampleSet_Union() {
	s1 := FromSlice([]int{1, 2})
	s2 := FromSlice([]int{2, 3})
	u := s1.Union(s2)
	v := u.Values()
	sort.Ints(v)
	fmt.Println(v)
	// Output: [1 2 3]
}

func ExampleSet_Intersection() {
	s1 := FromSlice([]int{1, 2})
	s2 := FromSlice([]int{2, 3})
	i := s1.Intersection(s2)
	v := i.Values()
	sort.Ints(v)
	fmt.Println(v)
	// Output: [2]
}

func ExampleSet_Difference() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{2, 3, 4})

	// difference is "directional" and will return the things in s1 that are not in s2
	d1 := s1.Difference(s2)
	v1 := d1.Values()
	sort.Ints(v1)
	fmt.Println(v1)

	// and vice versa
	d2 := s2.Difference(s1)
	v2 := d2.Values()
	sort.Ints(v2)
	fmt.Println(v2)
	// Output:
	// [1]
	// [4]
}

func ExampleSet_SymmetricalDifference() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{2, 3, 4})

	// Symmetrical difference is the inverse of intersection (anything not common to s1 and s2)
	d := s1.SymmetricalDifference(s2)
	v := d.Values()
	sort.Ints(v)
	fmt.Println(v)
	// Output: [1 4]
}

func ExampleSet_Equals() {
	s1 := FromSlice([]int{1, 2, 3})
	s2 := FromSlice([]int{1, 2, 3})
	s3 := FromSlice([]int{1, 2})
	fmt.Println(s1.Equals(s2), s1.Equals(s3))
	// Output: true false
}

func TestFromSlice(t *testing.T) {
	set := FromSlice([]string{"a", "b", "c"})
	require.Len(t, set, 3)
}

func TestMakeSet(t *testing.T) {
	set := make(Set[int])
	set.Add(1, 2, 3)
	require.Len(t, set, 3)
}

func TestSetOperations(t *testing.T) {
	set := FromSlice([]int{1, 2, 3})
	require.Len(t, set, 3)
	require.True(t, set.Has(1))
	require.True(t, set.Has(2))
	require.True(t, set.Has(3))
	require.False(t, set.Has(4))

	set.Add(1)
	require.Len(t, set, 3)
	set.Add(-1)
	require.Len(t, set, 4)
	set.Delete(1, -1)
	require.Len(t, set, 2)
	require.False(t, set.Has(1))
	require.False(t, set.Has(-1))

	set.Delete(100)
	require.Len(t, set, 2)

	set.Delete(2, 3, 50000)
	require.Len(t, set, 0)

	set2 := FromSlice([]int{1, 1, 1, 1, 1})
	require.Len(t, set2, 1)
}

func TestSetOperationsStrings(t *testing.T) {
	set := FromSlice([]string{"a", "b", "c"})
	require.Len(t, set, 3)
	require.True(t, set.Has("a"))
	require.True(t, set.Has("b"))
	require.True(t, set.Has("c"))
	require.False(t, set.Has("f"))

	set.Add("foobar")
	require.Len(t, set, 4)
	set.Add("")
	require.Len(t, set, 5)
}

func TestEquals(t *testing.T) {
	set := FromSlice([]int{3, 2, 1})
	set2 := FromSlice([]int{1, 2, 3})
	require.True(t, set.Equals(set2))

	emptySet1 := FromSlice([]int{})
	emptySet2 := FromSlice([]int{})
	require.True(t, emptySet1.Equals(emptySet2))

	set3 := FromSlice([]int{1, 2, 3})
	set4 := FromSlice([]int{1, 2, 3, 4})
	require.False(t, set3.Equals(set4))
	require.False(t, set4.Equals(set3))

	set5 := FromSlice([]int{2, 3, 4})
	require.False(t, set.Equals(set5))
}

func TestUnion(t *testing.T) {
	set1 := FromSlice([]int{1, 2, 3})
	set2 := FromSlice([]int{3, 4, 5})

	union := set1.Union(set2)
	require.Len(t, union, 5)

	test := FromSlice([]int{1, 2, 3, 4, 5})
	require.True(t, union.Equals(test))
}

func TestIntersection(t *testing.T) {
	set1 := FromSlice([]int{1, 2, 3})
	set2 := FromSlice([]int{2, 3, 4})

	intersection := set1.Intersection(set2)
	require.Len(t, intersection, 2)

	test := FromSlice([]int{2, 3})
	require.True(t, intersection.Equals(test))
}

func TestDifference(t *testing.T) {
	var a, b, diff, test Set[int]

	a = FromSlice([]int{1, 2, 3})
	b = FromSlice([]int{2, 3, 4})
	diff = a.Difference(b)
	require.Len(t, diff, 1)
	test = FromSlice([]int{1})
	require.True(t, diff.Equals(test))

	// test inverse
	diff = b.Difference(a)
	require.Len(t, diff, 1)
	test = FromSlice([]int{4})
	require.True(t, diff.Equals(test))

	// test what might be a false assumption by a naive user
	// remember, Difference is unidirectional! See SymmetricalDifference
	test = FromSlice([]int{1, 4})
	require.False(t, diff.Equals(test))
}

func TestSymmetricalDifference(t *testing.T) {
	var a, b, diff, test Set[int]

	a = FromSlice([]int{1, 2, 3})
	b = FromSlice([]int{2, 3, 4})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 2)
	test = FromSlice([]int{1, 4})
	require.True(t, diff.Equals(test))

	a = FromSlice([]int{1, 2})
	b = FromSlice([]int{2, 1})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 0)

	a = FromSlice([]int{1, 2})
	b = FromSlice([]int{3, 4})
	diff = a.SymmetricalDifference(b)
	require.Len(t, diff, 4)
	test = FromSlice([]int{1, 2, 3, 4})
	require.True(t, diff.Equals(test))
}

func BenchmarkFromSlice(b *testing.B) {
	s := make(Set[uint64])
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	for n := 0; n < b.N; n++ {
		s = FromSlice(slice)
	}
	resultSet64 = s
}

func BenchmarkAddByRangePreallocated(b *testing.B) {
	s := make(Set[uint64])
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	for n := 0; n < b.N; n++ {
		set := make(Set[uint64], len(slice))
		for _, val := range slice {
			set.Add(val)
		}
		s = set
	}
	resultSet64 = s
}

func BenchmarkAddByRangeUnallocated(b *testing.B) {
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	for n := 0; n < b.N; n++ {
		set := make(Set[uint64])
		for _, val := range slice {
			set.Add(val)
		}
	}
}

func BenchmarkDeleteByRange(b *testing.B) {
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	set := FromSlice(slice)
	for n := 0; n < b.N; n++ {
		for _, value := range slice {
			set.Delete(value)
		}
	}
}

func BenchmarkDeleteFullSlice(b *testing.B) {
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	set := FromSlice(slice)
	for n := 0; n < b.N; n++ {
		set.Delete(slice...)
	}
}

func BenchmarkHas(b *testing.B) {
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	set := FromSlice(slice)
	for n := 0; n < b.N; n++ {
		for _, value := range slice {
			set.Has(value)
		}
	}
}

func BenchmarkValues(b *testing.B) {
	slice, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	set := FromSlice(slice)
	for n := 0; n < b.N; n++ {
		set.Values()
	}
}

func BenchmarkEquals(b *testing.B) {
	var r bool
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	// use the same slice so we are comparing apples to apples
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice1)
	for n := 0; n < b.N; n++ {
		r = s1.Equals(s2)
	}
	resultBool = r
}

func BenchmarkEqualsDualDifference(b *testing.B) {
	var r bool
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice1)
	for n := 0; n < b.N; n++ {
		r = s1.Difference(s2).Len() == 0 && s2.Difference(s1).Len() == 0
	}
	resultBool = r
}

func BenchmarkEqualsSequential(b *testing.B) {
	var r bool
	slice := generateSequentialSliceUint64Ascending(10_000)
	s1 := FromSlice(slice)
	s2 := FromSlice(slice)

	for n := 0; n < b.N; n++ {
		r = s1.Equals(s2)
	}
	resultBool = r
}

func BenchmarkEqualsDualDifferenceSequential(b *testing.B) {
	var r bool
	slice := generateSequentialSliceUint64Ascending(10_000)
	s1 := FromSlice(slice)
	s2 := FromSlice(slice)

	for n := 0; n < b.N; n++ {
		r = s1.Difference(s2).Len() == 0 && s2.Difference(s1).Len() == 0
	}
	resultBool = r
}

func BenchmarkEqualsSequentialDesc(b *testing.B) {
	var r bool
	slice := generateSequentialSliceUint64Descending(10_000)
	s1 := FromSlice(slice)
	s2 := FromSlice(slice)

	for n := 0; n < b.N; n++ {
		r = s1.Equals(s2)
	}
	resultBool = r
}

func BenchmarkEqualsDualDifferenceSequentialDesc(b *testing.B) {
	var r bool
	slice := generateSequentialSliceUint64Descending(10_000)
	s1 := FromSlice(slice)
	s2 := FromSlice(slice)

	for n := 0; n < b.N; n++ {
		r = s1.Difference(s2).Len() == 0 && s2.Difference(s1).Len() == 0
	}
	resultBool = r
}

func BenchmarkDifference(b *testing.B) {
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	slice2, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice2)
	for n := 0; n < b.N; n++ {
		resultSet64 = s1.Difference(s2)
	}
}

func BenchmarkSymmetricalDiff(b *testing.B) {
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	slice2, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice2)
	for n := 0; n < b.N; n++ {
		resultBool = len(s1.SymmetricalDifference(s2)) == 0
	}
}

func BenchmarkIntersection(b *testing.B) {
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	slice2, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice2)
	for n := 0; n < b.N; n++ {
		resultSet64 = s1.Intersection(s2)
	}
}

func BenchmarkUnion(b *testing.B) {
	slice1, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	slice2, err := generateRandomSliceUint64(10_000, 1_000_000)
	if err != nil {
		b.Errorf("Error in generating random slice")
	}
	s1 := FromSlice(slice1)
	s2 := FromSlice(slice2)
	for n := 0; n < b.N; n++ {
		resultSet64 = s1.Union(s2)
	}
}

func generateSequentialSliceUint64Ascending(elements int) []uint64 {
	inputSlice := make([]uint64, elements)
	for i := 0; i < elements; i++ {
		inputSlice[i] = uint64(i)
	}
	return inputSlice
}

func generateSequentialSliceUint64Descending(elements int) []uint64 {
	inputSlice := make([]uint64, elements)
	for i := elements - 1; i >= 0; i-- {
		inputSlice[i] = uint64(i)
	}
	return inputSlice
}

func generateRandomSliceUint64(elements int, max int64) ([]uint64, error) {
	inputSlice := make([]uint64, elements)
	for i := 0; i < elements; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(max))
		if err != nil {
			return nil, err
		}
		inputSlice[i] = num.Uint64()
	}
	return inputSlice, nil
}
