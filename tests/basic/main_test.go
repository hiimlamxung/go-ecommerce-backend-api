// File phải có đuôi _test.go
package basic

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Function Test phải có chữ test ở đầu
func TestAddOne(t *testing.T) {
	// var (
	// 	input  = 1
	// 	output = 3
	// )
	// actual := AddOne(input)
	// if actual != output {
	// 	t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	// }

	// Cách test khác dùng thư viện testify
	assert.Equal(t, AddOne(2), 4, "AddOne(2) should be 3")
	assert.NotEqual(t, 2, 2)
	assert.Nil(t, nil, nil)
}

func TestAssert(t *testing.T) {
	assert.Equal(t, 2, 3)
	fmt.Println("After Assert")
}

func TestRequire(t *testing.T) {
	require.Equal(t, 2, 3) // Require nghĩa là Nếu không thỏa mãn thì sẽ dừng test và trả về lỗi
	fmt.Println("After require")
}

// Phân tích độ bao phủ code, run: "go test ./ -coverprofile=coverage.out" sẽ sinh ra file coverage.out
// File này hơi khó nhìn, có thể ren ra file html để dễ nhìn hơn: go tool cover -html coverage -o coverage.html

func TestAddOne2(t *testing.T) {
	var (
		input  = 1
		output = 3
	)
	actual := AddOne2(input)
	if actual != output {
		t.Errorf("AddOne(%d), input %d, actual = %d", input, output, actual)
	}
}
